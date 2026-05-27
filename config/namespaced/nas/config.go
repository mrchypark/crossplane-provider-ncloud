package nas

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	cluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/nas"
)

// Configure adds NAS resource configuration.
func Configure(p *ujconfig.Provider) {
	cluster.Configure(p)
}
