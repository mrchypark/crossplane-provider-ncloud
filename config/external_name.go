package config

import (
	"context"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"ncloud_access_control_group":                config.IdentifierFromProvider,
	"ncloud_access_control_group_rule":           config.IdentifierFromProvider,
	"ncloud_auto_scaling_group":                  config.IdentifierFromProvider,
	"ncloud_auto_scaling_policy":                 config.IdentifierFromProvider,
	"ncloud_auto_scaling_schedule":               config.IdentifierFromProvider,
	"ncloud_block_storage":                       config.IdentifierFromProvider,
	"ncloud_block_storage_snapshot":              config.IdentifierFromProvider,
	"ncloud_cdss_cluster":                        config.IdentifierFromProvider,
	"ncloud_cdss_config_group":                   config.IdentifierFromProvider,
	"ncloud_hadoop":                              config.IdentifierFromProvider,
	"ncloud_init_script":                         identifierFromProviderWithPlaceholder("0"),
	"ncloud_launch_configuration":                config.IdentifierFromProvider,
	"ncloud_lb":                                  config.IdentifierFromProvider,
	"ncloud_lb_listener":                         config.IdentifierFromProvider,
	"ncloud_lb_target_group":                     config.IdentifierFromProvider,
	"ncloud_lb_target_group_attachment":          config.IdentifierFromProvider,
	"ncloud_login_key":                           config.IdentifierFromProvider,
	"ncloud_mongodb":                             config.IdentifierFromProvider,
	"ncloud_mongodb_users":                       config.IdentifierFromProvider,
	"ncloud_mssql":                               config.IdentifierFromProvider,
	"ncloud_mysql":                               config.IdentifierFromProvider,
	"ncloud_mysql_databases":                     config.IdentifierFromProvider,
	"ncloud_mysql_recovery":                      config.IdentifierFromProvider,
	"ncloud_mysql_slave":                         config.IdentifierFromProvider,
	"ncloud_mysql_users":                         config.IdentifierFromProvider,
	"ncloud_nas_volume":                          config.IdentifierFromProvider,
	"ncloud_nat_gateway":                         config.IdentifierFromProvider,
	"ncloud_network_acl":                         config.IdentifierFromProvider,
	"ncloud_network_acl_deny_allow_group":        config.IdentifierFromProvider,
	"ncloud_network_acl_rule":                    config.IdentifierFromProvider,
	"ncloud_network_interface":                   config.IdentifierFromProvider,
	"ncloud_nks_cluster":                         config.IdentifierFromProvider,
	"ncloud_nks_node_pool":                       config.IdentifierFromProvider,
	"ncloud_objectstorage_bucket":                config.IdentifierFromProvider,
	"ncloud_objectstorage_bucket_acl":            config.IdentifierFromProvider,
	"ncloud_objectstorage_object":                config.IdentifierFromProvider,
	"ncloud_objectstorage_object_acl":            config.IdentifierFromProvider,
	"ncloud_objectstorage_object_copy":           config.IdentifierFromProvider,
	"ncloud_placement_group":                     config.IdentifierFromProvider,
	"ncloud_postgresql":                          config.IdentifierFromProvider,
	"ncloud_postgresql_databases":                config.IdentifierFromProvider,
	"ncloud_postgresql_read_replica":             config.IdentifierFromProvider,
	"ncloud_postgresql_users":                    config.IdentifierFromProvider,
	"ncloud_public_ip":                           config.IdentifierFromProvider,
	"ncloud_redis":                               config.IdentifierFromProvider,
	"ncloud_redis_config_group":                  config.IdentifierFromProvider,
	"ncloud_route":                               config.IdentifierFromProvider,
	"ncloud_route_table":                         config.IdentifierFromProvider,
	"ncloud_route_table_association":             config.IdentifierFromProvider,
	"ncloud_server":                              config.IdentifierFromProvider,
	"ncloud_ses_cluster":                         config.IdentifierFromProvider,
	"ncloud_sourcebuild_project":                 config.IdentifierFromProvider,
	"ncloud_sourcecommit_repository":             config.IdentifierFromProvider,
	"ncloud_sourcedeploy_project":                config.IdentifierFromProvider,
	"ncloud_sourcedeploy_project_stage":          config.IdentifierFromProvider,
	"ncloud_sourcedeploy_project_stage_scenario": config.IdentifierFromProvider,
	"ncloud_sourcepipeline_project":              config.IdentifierFromProvider,
	"ncloud_subnet":                              config.IdentifierFromProvider,
	"ncloud_vpc":                                 identifierFromProviderWithPlaceholder("0"),
	"ncloud_vpc_peering":                         config.IdentifierFromProvider,
}

func identifierFromProviderWithPlaceholder(placeholder string) config.ExternalName {
	return config.NewExternalNameFrom(config.IdentifierFromProvider,
		config.WithGetIDFn(func(_ config.GetIDFn, _ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
			if externalName == "" {
				return placeholder, nil
			}
			return externalName, nil
		}))
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
