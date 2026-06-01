package analytics

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds analytics and data platform resource configuration.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("ncloud_cdss_config_group", func(r *ujconfig.Resource) {
		r.ShortGroup = "analytics"
		r.Kind = "CdssConfigGroup"
	})
	p.AddResourceConfigurator("ncloud_hadoop", func(r *ujconfig.Resource) {
		r.ShortGroup = "analytics"
		r.Kind = "Hadoop"
		r.UseAsync = true
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: "ncloud_vpc"}
		r.References["edge_node_subnet_no"] = ujconfig.Reference{TerraformName: "ncloud_subnet"}
		r.References["master_node_subnet_no"] = ujconfig.Reference{TerraformName: "ncloud_subnet"}
		r.References["worker_node_subnet_no"] = ujconfig.Reference{TerraformName: "ncloud_subnet"}
		r.References["login_key_name"] = ujconfig.Reference{TerraformName: "ncloud_login_key"}
		r.References["bucket_name"] = ujconfig.Reference{TerraformName: "ncloud_objectstorage_bucket"}
	})
	p.AddResourceConfigurator("ncloud_ses_cluster", func(r *ujconfig.Resource) {
		r.ShortGroup = "analytics"
		r.Kind = "SesCluster"
		r.UseAsync = true
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: "ncloud_vpc"}
		r.References["login_key_name"] = ujconfig.Reference{TerraformName: "ncloud_login_key"}
	})
}
