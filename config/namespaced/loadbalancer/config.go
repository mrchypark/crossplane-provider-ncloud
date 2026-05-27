package loadbalancer

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	cluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/loadbalancer"
)

// Configure adds load balancer resource configuration.
func Configure(p *ujconfig.Provider) {
	cluster.Configure(p)
}
