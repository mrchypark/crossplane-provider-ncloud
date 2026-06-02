package network

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	cluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/network"
)

// Configure adds VPC and network resource configuration.
func Configure(p *ujconfig.Provider) {
	cluster.Configure(p)
}
