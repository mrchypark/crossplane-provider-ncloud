package objectstorage

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	cluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/objectstorage"
)

// Configure adds Object Storage resource configuration.
func Configure(p *ujconfig.Provider) {
	cluster.Configure(p)
}
