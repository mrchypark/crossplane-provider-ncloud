package database

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	cluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/database"
)

// Configure adds database and cache resource configuration.
func Configure(p *ujconfig.Provider) {
	cluster.Configure(p)
}
