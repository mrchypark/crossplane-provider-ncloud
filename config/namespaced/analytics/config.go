package analytics

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	cluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/analytics"
)

// Configure adds analytics and data platform resource configuration.
func Configure(p *ujconfig.Provider) {
	cluster.Configure(p)
}
