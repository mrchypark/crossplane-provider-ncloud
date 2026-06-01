package config

import (
	"context"
	"fmt"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"ncloud_access_control_group":                identifierFromProviderWithPlaceholder("0"),
	"ncloud_access_control_group_rule":           identifierFromProviderFromField("access_control_group_no", "accessControlGroupNo", "0"),
	"ncloud_auto_scaling_group":                  identifierFromProviderWithPlaceholder("0"),
	"ncloud_auto_scaling_policy":                 identifierFromProviderWithPlaceholder("0"),
	"ncloud_auto_scaling_schedule":               identifierFromProviderWithPlaceholder("0"),
	"ncloud_block_storage":                       identifierFromProviderWithPlaceholder("0"),
	"ncloud_block_storage_snapshot":              identifierFromProviderWithPlaceholder("0"),
	"ncloud_cdss_cluster":                        identifierFromProviderWithPlaceholder("0"),
	"ncloud_cdss_config_group":                   identifierFromProviderFromFields("0:{{ .first }}", "kafka_version_code", "kafkaVersionCode"),
	"ncloud_hadoop":                              identifierFromProviderWithPlaceholder("0"),
	"ncloud_init_script":                         identifierFromProviderWithPlaceholder("0"),
	"ncloud_launch_configuration":                identifierFromProviderWithPlaceholder("0"),
	"ncloud_lb":                                  config.FrameworkResourceWithComputedIdentifier("load_balancer_no", "0"),
	"ncloud_lb_listener":                         identifierFromProviderFromFields("{{ .first }}:0", "load_balancer_no", "loadBalancerNo"),
	"ncloud_lb_target_group":                     identifierFromProviderWithPlaceholder("0"),
	"ncloud_lb_target_group_attachment":          identifierFromProviderFromField("target_group_no", "targetGroupNo", "0"),
	"ncloud_login_key":                           identifierFromProviderWithPlaceholder("0"),
	"ncloud_mongodb":                             identifierFromProviderWithPlaceholder("0"),
	"ncloud_mongodb_users":                       identifierFromProviderWithPlaceholder("0"),
	"ncloud_mssql":                               identifierFromProviderWithPlaceholder("0"),
	"ncloud_mysql":                               identifierFromProviderWithPlaceholder("0"),
	"ncloud_mysql_databases":                     identifierFromProviderWithPlaceholder("0"),
	"ncloud_mysql_recovery":                      identifierFromProviderWithPlaceholder("0"),
	"ncloud_mysql_slave":                         identifierFromProviderWithPlaceholder("0"),
	"ncloud_mysql_users":                         identifierFromProviderWithPlaceholder("0"),
	"ncloud_nas_volume":                          identifierFromProviderWithPlaceholder("0"),
	"ncloud_nat_gateway":                         identifierFromProviderWithPlaceholder("0"),
	"ncloud_network_acl":                         identifierFromProviderWithPlaceholder("0"),
	"ncloud_network_acl_deny_allow_group":        identifierFromProviderWithPlaceholder("0"),
	"ncloud_network_acl_rule":                    identifierFromProviderFromField("network_acl_no", "networkAclNo", "0"),
	"ncloud_network_interface":                   identifierFromProviderWithPlaceholder("0"),
	"ncloud_nks_cluster":                         identifierFromProviderWithPlaceholder("00000000-0000-0000-0000-000000000000"),
	"ncloud_nks_node_pool":                       identifierFromProviderFromFields("{{ .first }}:{{ .second }}", "cluster_uuid", "clusterUuid", "node_pool_name", "nodePoolName"),
	"ncloud_objectstorage_bucket":                identifierFromProviderFromField("bucket_name", "bucketName", "0"),
	"ncloud_objectstorage_bucket_acl":            identifierFromProviderFromField("bucket_name", "bucketName", "0"),
	"ncloud_objectstorage_object":                identifierFromProviderFromFields("{{ .first }}/{{ .second }}", "bucket", "bucket", "key", "key"),
	"ncloud_objectstorage_object_acl":            identifierFromProviderFromField("object_id", "objectId", "0"),
	"ncloud_objectstorage_object_copy":           identifierFromProviderFromFields("{{ .first }}/{{ .second }}", "bucket", "bucket", "key", "key"),
	"ncloud_placement_group":                     identifierFromProviderWithPlaceholder("0"),
	"ncloud_postgresql":                          identifierFromProviderWithPlaceholder("0"),
	"ncloud_postgresql_databases":                identifierFromProviderWithPlaceholder("0"),
	"ncloud_postgresql_read_replica":             identifierFromProviderWithPlaceholder("0"),
	"ncloud_postgresql_users":                    identifierFromProviderWithPlaceholder("0"),
	"ncloud_public_ip":                           identifierFromProviderWithPlaceholder("0"),
	"ncloud_redis":                               identifierFromProviderWithPlaceholder("999999999"),
	"ncloud_redis_config_group":                  identifierFromProviderWithPlaceholder("0"),
	"ncloud_route":                               identifierFromProviderFromFields("{{ .first }}:{{ .second }}", "route_table_no", "routeTableNo", "destination_cidr_block", "destinationCidrBlock"),
	"ncloud_route_table":                         identifierFromProviderWithPlaceholder("0"),
	"ncloud_route_table_association":             identifierFromProviderFromFields("{{ .first }}:{{ .second }}", "route_table_no", "routeTableNo", "subnet_no", "subnetNo"),
	"ncloud_server":                              identifierFromProviderWithPlaceholder("0"),
	"ncloud_ses_cluster":                         identifierFromProviderWithPlaceholder("0"),
	"ncloud_sourcebuild_project":                 identifierFromProviderWithPlaceholder("0"),
	"ncloud_sourcecommit_repository":             identifierFromProviderWithPlaceholder("0"),
	"ncloud_sourcedeploy_project":                identifierFromProviderWithPlaceholder("0"),
	"ncloud_sourcedeploy_project_stage":          identifierFromProviderFromFields("{{ .first }}:0", "project_id", "projectId"),
	"ncloud_sourcedeploy_project_stage_scenario": identifierFromProviderWithPlaceholder("0"),
	"ncloud_sourcepipeline_project":              identifierFromProviderWithPlaceholder("0"),
	"ncloud_subnet":                              identifierFromProviderWithPlaceholder("0"),
	"ncloud_vpc":                                 identifierFromProviderWithPlaceholder("0"),
	"ncloud_vpc_peering":                         identifierFromProviderWithPlaceholder("0"),
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

func identifierFromProviderFromField(snakeName, camelName, placeholder string) config.ExternalName {
	return config.NewExternalNameFrom(config.IdentifierFromProvider,
		config.WithGetIDFn(func(_ config.GetIDFn, _ context.Context, externalName string, parameters map[string]any, _ map[string]any) (string, error) {
			if externalName != "" {
				return externalName, nil
			}
			if value := stringParameter(parameters, snakeName, camelName); value != "" {
				return value, nil
			}
			return placeholder, nil
		}))
}

func identifierFromProviderFromFields(format string, names ...string) config.ExternalName {
	return config.NewExternalNameFrom(config.IdentifierFromProvider,
		config.WithGetIDFn(func(_ config.GetIDFn, _ context.Context, externalName string, parameters map[string]any, _ map[string]any) (string, error) {
			if externalName != "" {
				return externalName, nil
			}
			first := "0"
			second := "0"
			if len(names) >= 2 {
				if value := stringParameter(parameters, names[0], names[1]); value != "" {
					first = value
				}
			}
			if len(names) >= 4 {
				if value := stringParameter(parameters, names[2], names[3]); value != "" {
					second = value
				}
			}
			switch format {
			case "{{ .first }}:0":
				return fmt.Sprintf("%s:0", first), nil
			case "0:{{ .first }}":
				return fmt.Sprintf("0:%s", first), nil
			case "{{ .first }}:{{ .second }}":
				return fmt.Sprintf("%s:%s", first, second), nil
			case "{{ .first }}/{{ .second }}":
				return fmt.Sprintf("%s/%s", first, second), nil
			default:
				return "0", nil
			}
		}))
}

func stringParameter(parameters map[string]any, names ...string) string {
	for _, name := range names {
		if value, ok := parameters[name]; ok {
			switch v := value.(type) {
			case string:
				return v
			case *string:
				if v != nil {
					return *v
				}
			default:
				if v != nil {
					return fmt.Sprint(v)
				}
			}
		}
	}
	return ""
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
