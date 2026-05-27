package autoscaling

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds auto scaling resource configuration.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("ncloud_auto_scaling_group", func(r *ujconfig.Resource) {
		r.ShortGroup = "autoscaling"
		r.Kind = "AutoScalingGroup"
		r.UseAsync = true
		r.References["launch_configuration_no"] = ujconfig.Reference{TerraformName: "ncloud_launch_configuration"}
		r.References["subnet_no"] = ujconfig.Reference{TerraformName: "ncloud_subnet"}
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: "ncloud_vpc"}
		r.References["access_control_group_no_list"] = ujconfig.Reference{
			TerraformName:     "ncloud_access_control_group",
			RefFieldName:      "AccessControlGroupNoRefs",
			SelectorFieldName: "AccessControlGroupNoSelector",
		}
		r.References["target_group_list"] = ujconfig.Reference{
			TerraformName:     "ncloud_lb_target_group",
			RefFieldName:      "TargetGroupRefs",
			SelectorFieldName: "TargetGroupSelector",
		}
	})
	p.AddResourceConfigurator("ncloud_auto_scaling_policy", func(r *ujconfig.Resource) {
		r.ShortGroup = "autoscaling"
		r.Kind = "AutoScalingPolicy"
		r.References["auto_scaling_group_no"] = ujconfig.Reference{TerraformName: "ncloud_auto_scaling_group"}
	})
	p.AddResourceConfigurator("ncloud_auto_scaling_schedule", func(r *ujconfig.Resource) {
		r.ShortGroup = "autoscaling"
		r.Kind = "AutoScalingSchedule"
		r.References["auto_scaling_group_no"] = ujconfig.Reference{TerraformName: "ncloud_auto_scaling_group"}
	})
}
