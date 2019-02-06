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
	"strings"

	"github.com/spf13/viper"
)

const (
	gRPCPrefix             = "reporter.grpc."
	collectorHostPort      = gRPCPrefix + "host-port"
	collectorTLS           = gRPCPrefix + "tls"
	collectorTLSCA         = gRPCPrefix + "tls.ca"
	collectorTLSServerName = gRPCPrefix + "tls.server-name"
)

// Options Struct to hold configurations
type Options struct {
	// CollectorHostPort is list of host:port Jaeger Collectors.
	CollectorHostPort []string
	TLS               bool
	TLSCA             string
	TLSServerName     string
}

// AddFlags adds flags for Options.
func AddFlags(flags *flag.FlagSet) {
	flags.String(collectorHostPort, "", "(experimental) Comma-separated string representing host:port of a static list of collectors to connect to directly.")
	flags.Bool(collectorTLS, false, "(experimental) Enable TLS.")
	flags.String(collectorTLSCA, "", "(experimental) Path to a TLS CA file. (default use the systems truststore)")
	flags.String(collectorTLSServerName, "", "(experimental) Override the TLS server name.")
}

// InitFromViper initializes Options with properties retrieved from Viper.
func (o *Options) InitFromViper(v *viper.Viper) *Options {
	hostPorts := v.GetString(collectorHostPort)
	if hostPorts != "" {
		o.CollectorHostPort = strings.Split(hostPorts, ",")
	}

	o.TLS = v.GetBool(collectorTLS)
	o.TLSCA = v.GetString(collectorTLSCA)
	o.TLSServerName = v.GetString(collectorTLSServerName)
	return o
}
