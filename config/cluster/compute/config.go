package compute

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds compute resource configuration.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("ncloud_login_key", func(r *ujconfig.Resource) {
		r.ShortGroup = "compute"
		r.Kind = "LoginKey"
	})
	p.AddResourceConfigurator("ncloud_init_script", func(r *ujconfig.Resource) {
		r.ShortGroup = "compute"
		r.Kind = "InitScript"
	})
	p.AddResourceConfigurator("ncloud_server", func(r *ujconfig.Resource) {
		r.ShortGroup = "compute"
		r.Kind = "Server"
		r.UseAsync = true
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: "ncloud_vpc"}
		r.References["subnet_no"] = ujconfig.Reference{TerraformName: "ncloud_subnet"}
		r.References["init_script_no"] = ujconfig.Reference{TerraformName: "ncloud_init_script"}
		r.References["network_interface.network_interface_no"] = ujconfig.Reference{TerraformName: "ncloud_network_interface"}
		ujconfig.MoveToStatus(r.TerraformResource, "instance_no", "private_ip", "public_ip", "port_forwarding_public_ip", "port_forwarding_external_port", "port_forwarding_internal_port")
	})
	p.AddResourceConfigurator("ncloud_network_interface", func(r *ujconfig.Resource) {
		r.ShortGroup = "compute"
		r.Kind = "NetworkInterface"
		r.References["subnet_no"] = ujconfig.Reference{TerraformName: "ncloud_subnet"}
		r.References["server_instance_no"] = ujconfig.Reference{TerraformName: "ncloud_server"}
		r.References["access_control_groups"] = ujconfig.Reference{
			TerraformName:     "ncloud_access_control_group",
			RefFieldName:      "AccessControlGroupRefs",
			SelectorFieldName: "AccessControlGroupSelector",
		}
		ujconfig.MoveToStatus(r.TerraformResource, "network_interface_no", "is_default", "status")
	})
	p.AddResourceConfigurator("ncloud_public_ip", func(r *ujconfig.Resource) {
		r.ShortGroup = "compute"
		r.Kind = "PublicIP"
		r.References["server_instance_no"] = ujconfig.Reference{TerraformName: "ncloud_server"}
		r.References["instance_no"] = ujconfig.Reference{TerraformName: "ncloud_server"}
		ujconfig.MoveToStatus(r.TerraformResource, "public_ip_no", "public_ip")
	})
	p.AddResourceConfigurator("ncloud_block_storage", func(r *ujconfig.Resource) {
		r.ShortGroup = "compute"
		r.Kind = "BlockStorage"
		r.UseAsync = true
		r.References["server_instance_no"] = ujconfig.Reference{TerraformName: "ncloud_server"}
		ujconfig.MoveToStatus(r.TerraformResource, "block_storage_no", "server_name", "status")
	})
	p.AddResourceConfigurator("ncloud_block_storage_snapshot", func(r *ujconfig.Resource) {
		r.ShortGroup = "compute"
		r.Kind = "BlockStorageSnapshot"
		r.UseAsync = true
		r.References["block_storage_instance_no"] = ujconfig.Reference{TerraformName: "ncloud_block_storage"}
		ujconfig.MoveToStatus(r.TerraformResource, "block_storage_snapshot_no", "status")
	})
	p.AddResourceConfigurator("ncloud_launch_configuration", func(r *ujconfig.Resource) {
		r.ShortGroup = "compute"
		r.Kind = "LaunchConfiguration"
		r.References["login_key_name"] = ujconfig.Reference{TerraformName: "ncloud_login_key"}
	})
	p.AddResourceConfigurator("ncloud_placement_group", func(r *ujconfig.Resource) {
		r.ShortGroup = "compute"
		r.Kind = "PlacementGroup"
	})
	p.AddResourceConfigurator("ncloud_access_control_group", func(r *ujconfig.Resource) {
		r.ShortGroup = "compute"
		r.Kind = "AccessControlGroup"
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: "ncloud_vpc"}
	})
	p.AddResourceConfigurator("ncloud_access_control_group_rule", func(r *ujconfig.Resource) {
		r.ShortGroup = "compute"
		r.Kind = "AccessControlGroupRule"
		r.References["access_control_group_no"] = ujconfig.Reference{TerraformName: "ncloud_access_control_group"}
	})
}
