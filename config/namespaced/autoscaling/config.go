package autoscaling

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	cluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/autoscaling"
)

// Configure adds auto scaling resource configuration.
func Configure(p *ujconfig.Provider) {
	cluster.Configure(p)
}
