package subaccount

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	cluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/subaccount"
)

// Configure adds Sub Account resource configuration.
func Configure(p *ujconfig.Provider) {
	cluster.Configure(p)
}
