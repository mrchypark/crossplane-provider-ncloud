package nks

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds NKS resource configuration.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("ncloud_nks_cluster", func(r *ujconfig.Resource) {
		r.ShortGroup = "nks"
		r.Kind = "NksCluster"
		r.UseAsync = true
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: "ncloud_vpc"}
		r.References["lb_private_subnet_no"] = ujconfig.Reference{TerraformName: "ncloud_subnet"}
		r.References["lb_public_subnet_no"] = ujconfig.Reference{TerraformName: "ncloud_subnet"}
		r.References["login_key_name"] = ujconfig.Reference{TerraformName: "ncloud_login_key"}
		r.References["subnet_no_list"] = ujconfig.Reference{
			TerraformName:     "ncloud_subnet",
			RefFieldName:      "SubnetNoRefs",
			SelectorFieldName: "SubnetNoSelector",
		}
		ujconfig.MoveToStatus(r.TerraformResource, "uuid", "endpoint")
	})
	p.AddResourceConfigurator("ncloud_nks_node_pool", func(r *ujconfig.Resource) {
		r.ShortGroup = "nks"
		r.Kind = "NksNodePool"
		r.UseAsync = true
		r.References["cluster_uuid"] = ujconfig.Reference{TerraformName: "ncloud_nks_cluster"}
	})
}
