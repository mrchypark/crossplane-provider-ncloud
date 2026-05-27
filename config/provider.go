package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	analyticsCluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/analytics"
	autoscalingCluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/autoscaling"
	computeCluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/compute"
	databaseCluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/database"
	loadbalancerCluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/loadbalancer"
	nasCluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/nas"
	networkCluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/network"
	nksCluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/nks"
	objectstorageCluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/objectstorage"
	sourceCluster "github.com/mrchypark/crossplane-provider-ncloud/config/cluster/source"
	analyticsNamespaced "github.com/mrchypark/crossplane-provider-ncloud/config/namespaced/analytics"
	autoscalingNamespaced "github.com/mrchypark/crossplane-provider-ncloud/config/namespaced/autoscaling"
	computeNamespaced "github.com/mrchypark/crossplane-provider-ncloud/config/namespaced/compute"
	databaseNamespaced "github.com/mrchypark/crossplane-provider-ncloud/config/namespaced/database"
	loadbalancerNamespaced "github.com/mrchypark/crossplane-provider-ncloud/config/namespaced/loadbalancer"
	nasNamespaced "github.com/mrchypark/crossplane-provider-ncloud/config/namespaced/nas"
	networkNamespaced "github.com/mrchypark/crossplane-provider-ncloud/config/namespaced/network"
	nksNamespaced "github.com/mrchypark/crossplane-provider-ncloud/config/namespaced/nks"
	objectstorageNamespaced "github.com/mrchypark/crossplane-provider-ncloud/config/namespaced/objectstorage"
	sourceNamespaced "github.com/mrchypark/crossplane-provider-ncloud/config/namespaced/source"
)

const (
	resourcePrefix = "ncloud"
	modulePath     = "github.com/mrchypark/crossplane-provider-ncloud"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("ncloud.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		analyticsCluster.Configure,
		autoscalingCluster.Configure,
		networkCluster.Configure,
		computeCluster.Configure,
		databaseCluster.Configure,
		loadbalancerCluster.Configure,
		nasCluster.Configure,
		nksCluster.Configure,
		objectstorageCluster.Configure,
		sourceCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("ncloud.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		analyticsNamespaced.Configure,
		autoscalingNamespaced.Configure,
		networkNamespaced.Configure,
		computeNamespaced.Configure,
		databaseNamespaced.Configure,
		loadbalancerNamespaced.Configure,
		nasNamespaced.Configure,
		nksNamespaced.Configure,
		objectstorageNamespaced.Configure,
		sourceNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
