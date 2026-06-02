package source

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

const shortGroupSource = "source"

// Configure adds Source* resource configuration.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("ncloud_sourcebuild_project", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupSource
		r.Kind = "SourcebuildProject"
		r.UseAsync = true
	})
	p.AddResourceConfigurator("ncloud_sourcecommit_repository", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupSource
		r.Kind = "SourcecommitRepository"
	})
	p.AddResourceConfigurator("ncloud_sourcedeploy_project", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupSource
		r.Kind = "SourcedeployProject"
	})
	p.AddResourceConfigurator("ncloud_sourcedeploy_project_stage", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupSource
		r.Kind = "SourcedeployProjectStage"
		r.References["project_id"] = ujconfig.Reference{TerraformName: "ncloud_sourcedeploy_project"}
	})
	p.AddResourceConfigurator("ncloud_sourcedeploy_project_stage_scenario", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupSource
		r.Kind = "SourcedeployProjectStageScenario"
		r.References["project_id"] = ujconfig.Reference{TerraformName: "ncloud_sourcedeploy_project"}
		r.References["stage_id"] = ujconfig.Reference{TerraformName: "ncloud_sourcedeploy_project_stage"}
	})
	p.AddResourceConfigurator("ncloud_sourcepipeline_project", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupSource
		r.Kind = "SourcepipelineProject"
		r.UseAsync = true
	})
}
