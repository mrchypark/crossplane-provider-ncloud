package network

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortGroupNetwork = "network"
	tfVPC             = "ncloud_vpc"
)

// Configure adds VPC and network resource configuration.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("ncloud_vpc", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupNetwork
		r.Kind = "Vpc"
		ujconfig.MoveToStatus(r.TerraformResource, "vpc_no", "default_access_control_group_no", "default_network_acl_no", "default_private_route_table_no", "default_public_route_table_no")
	})
	p.AddResourceConfigurator("ncloud_subnet", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupNetwork
		r.Kind = "Subnet"
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: tfVPC}
		r.References["network_acl_no"] = ujconfig.Reference{TerraformName: "ncloud_network_acl"}
		ujconfig.MoveToStatus(r.TerraformResource, "subnet_no")
	})
	p.AddResourceConfigurator("ncloud_network_acl", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupNetwork
		r.Kind = "NetworkACL"
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: tfVPC}
		ujconfig.MoveToStatus(r.TerraformResource, "network_acl_no", "is_default")
	})
	p.AddResourceConfigurator("ncloud_network_acl_deny_allow_group", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupNetwork
		r.Kind = "NetworkACLDenyAllowGroup"
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: tfVPC}
	})
	p.AddResourceConfigurator("ncloud_network_acl_rule", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupNetwork
		r.Kind = "NetworkACLRule"
		r.References["network_acl_no"] = ujconfig.Reference{TerraformName: "ncloud_network_acl"}
	})
	p.AddResourceConfigurator("ncloud_route_table", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupNetwork
		r.Kind = "RouteTable"
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: tfVPC}
		ujconfig.MoveToStatus(r.TerraformResource, "route_table_no", "is_default")
	})
	p.AddResourceConfigurator("ncloud_route", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupNetwork
		r.Kind = "Route"
		r.References["route_table_no"] = ujconfig.Reference{TerraformName: "ncloud_route_table"}
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: tfVPC}
	})
	p.AddResourceConfigurator("ncloud_route_table_association", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupNetwork
		r.Kind = "RouteTableAssociation"
		r.References["route_table_no"] = ujconfig.Reference{TerraformName: "ncloud_route_table"}
		r.References["subnet_no"] = ujconfig.Reference{TerraformName: "ncloud_subnet"}
	})
	p.AddResourceConfigurator("ncloud_nat_gateway", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupNetwork
		r.Kind = "NatGateway"
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: tfVPC}
		r.References["subnet_no"] = ujconfig.Reference{TerraformName: "ncloud_subnet"}
		r.References["public_ip_no"] = ujconfig.Reference{TerraformName: "ncloud_public_ip"}
		ujconfig.MoveToStatus(r.TerraformResource, "nat_gateway_no", "private_ip", "public_ip", "subnet_name")
	})
	p.AddResourceConfigurator("ncloud_vpc_peering", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupNetwork
		r.Kind = "VpcPeering"
		r.References["source_vpc_no"] = ujconfig.Reference{TerraformName: tfVPC}
		r.References["target_vpc_no"] = ujconfig.Reference{TerraformName: tfVPC}
		ujconfig.MoveToStatus(r.TerraformResource, "vpc_peering_no", "target_vpc_name")
	})
}
