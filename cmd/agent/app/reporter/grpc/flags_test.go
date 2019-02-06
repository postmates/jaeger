// Copyright (c) 2018 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpc

import (
	"flag"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBindFlags(t *testing.T) {
	v := viper.New()

	tests := []struct {
		cOpts    []string
		expected *Options
	}{
		{cOpts: []string{"--reporter.grpc.host-port=localhost:1111"},
			expected: &Options{CollectorHostPort: []string{"localhost:1111"}}},
		{cOpts: []string{"--reporter.grpc.host-port=localhost:1111,localhost:2222"},
			expected: &Options{CollectorHostPort: []string{"localhost:1111", "localhost:2222"}}},
		{cOpts: []string{"--reporter.grpc.tls", "--reporter.grpc.tls.ca=/tmp/myca", "--reporter.grpc.tls.server-name=testserver"},
			expected: &Options{TLS: true, TLSCA: "/tmp/myca", TLSServerName: "testserver"}},
	}
	for _, test := range tests {
		command := cobra.Command{}
		flags := &flag.FlagSet{}
		AddFlags(flags)
		command.PersistentFlags().AddGoFlagSet(flags)
		v.BindPFlags(command.PersistentFlags())

		err := command.ParseFlags(test.cOpts)
		require.NoError(t, err)
		b := new(Options).InitFromViper(v)
		assert.Equal(t, test.expected, b)
	}
}
