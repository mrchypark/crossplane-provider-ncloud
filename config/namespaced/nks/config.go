package nks

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	cluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/nks"
)

// Configure adds NKS resource configuration.
func Configure(p *ujconfig.Provider) {
	cluster.Configure(p)
}
