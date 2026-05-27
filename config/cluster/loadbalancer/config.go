package loadbalancer

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds load balancer resource configuration.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("ncloud_lb", func(r *ujconfig.Resource) {
		r.ShortGroup = "loadbalancer"
		r.Kind = "Lb"
		r.UseAsync = true
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: "ncloud_vpc"}
		r.References["subnet_no_list"] = ujconfig.Reference{
			TerraformName:     "ncloud_subnet",
			RefFieldName:      "SubnetNoRefs",
			SelectorFieldName: "SubnetNoSelector",
		}
		ujconfig.MoveToStatus(r.TerraformResource, "load_balancer_no", "domain", "ip_list", "listener_no_list")
	})
	p.AddResourceConfigurator("ncloud_lb_listener", func(r *ujconfig.Resource) {
		r.ShortGroup = "loadbalancer"
		r.Kind = "LbListener"
		r.References["load_balancer_no"] = ujconfig.Reference{TerraformName: "ncloud_lb"}
		r.References["target_group_no"] = ujconfig.Reference{TerraformName: "ncloud_lb_target_group"}
		ujconfig.MoveToStatus(r.TerraformResource, "listener_no", "rule_no_list")
	})
	p.AddResourceConfigurator("ncloud_lb_target_group", func(r *ujconfig.Resource) {
		r.ShortGroup = "loadbalancer"
		r.Kind = "LbTargetGroup"
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: "ncloud_vpc"}
		r.References["target_no_list"] = ujconfig.Reference{
			TerraformName:     "ncloud_server",
			RefFieldName:      "TargetNoRefs",
			SelectorFieldName: "TargetNoSelector",
		}
		ujconfig.MoveToStatus(r.TerraformResource, "target_group_no", "load_balancer_instance_no")
	})
	p.AddResourceConfigurator("ncloud_lb_target_group_attachment", func(r *ujconfig.Resource) {
		r.ShortGroup = "loadbalancer"
		r.Kind = "LbTargetGroupAttachment"
		r.References["target_group_no"] = ujconfig.Reference{TerraformName: "ncloud_lb_target_group"}
		r.References["target_no_list"] = ujconfig.Reference{
			TerraformName:     "ncloud_server",
			RefFieldName:      "TargetNoRefs",
			SelectorFieldName: "TargetNoSelector",
		}
	})
}
