package database

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortGroupDatabase = "database"
	tfSubnet           = "ncloud_subnet"
)

// Configure adds database and cache resource configuration.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("ncloud_mongodb", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupDatabase
		r.Kind = "Mongodb"
		r.UseAsync = true
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: "ncloud_vpc"}
		r.References["subnet_no"] = ujconfig.Reference{TerraformName: tfSubnet}
	})
	p.AddResourceConfigurator("ncloud_mongodb_users", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupDatabase
		r.Kind = "MongodbUsers"
		r.References["id"] = ujconfig.Reference{TerraformName: "ncloud_mongodb"}
	})
	p.AddResourceConfigurator("ncloud_mssql", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupDatabase
		r.Kind = "Mssql"
		r.UseAsync = true
		r.References["subnet_no"] = ujconfig.Reference{TerraformName: tfSubnet}
	})
	p.AddResourceConfigurator("ncloud_mysql", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupDatabase
		r.Kind = "Mysql"
		r.UseAsync = true
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: "ncloud_vpc"}
		r.References["subnet_no"] = ujconfig.Reference{TerraformName: tfSubnet}
	})
	p.AddResourceConfigurator("ncloud_mysql_databases", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupDatabase
		r.Kind = "MysqlDatabases"
		r.References["mysql_instance_no"] = ujconfig.Reference{TerraformName: "ncloud_mysql"}
	})
	p.AddResourceConfigurator("ncloud_mysql_recovery", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupDatabase
		r.Kind = "MysqlRecovery"
		r.UseAsync = true
		r.References["mysql_instance_no"] = ujconfig.Reference{TerraformName: "ncloud_mysql"}
	})
	p.AddResourceConfigurator("ncloud_mysql_slave", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupDatabase
		r.Kind = "MysqlSlave"
		r.UseAsync = true
		r.References["mysql_instance_no"] = ujconfig.Reference{TerraformName: "ncloud_mysql"}
	})
	p.AddResourceConfigurator("ncloud_mysql_users", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupDatabase
		r.Kind = "MysqlUsers"
		r.References["mysql_instance_no"] = ujconfig.Reference{TerraformName: "ncloud_mysql"}
	})
	p.AddResourceConfigurator("ncloud_postgresql", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupDatabase
		r.Kind = "Postgresql"
		r.UseAsync = true
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: "ncloud_vpc"}
		r.References["subnet_no"] = ujconfig.Reference{TerraformName: tfSubnet}
	})
	p.AddResourceConfigurator("ncloud_postgresql_databases", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupDatabase
		r.Kind = "PostgresqlDatabases"
		r.References["id"] = ujconfig.Reference{TerraformName: "ncloud_postgresql"}
	})
	p.AddResourceConfigurator("ncloud_postgresql_read_replica", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupDatabase
		r.Kind = "PostgresqlReadReplica"
		r.UseAsync = true
		r.References["postgresql_instance_no"] = ujconfig.Reference{TerraformName: "ncloud_postgresql"}
	})
	p.AddResourceConfigurator("ncloud_postgresql_users", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupDatabase
		r.Kind = "PostgresqlUsers"
		r.References["id"] = ujconfig.Reference{TerraformName: "ncloud_postgresql"}
	})
	p.AddResourceConfigurator("ncloud_redis", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupDatabase
		r.Kind = "Redis"
		r.UseAsync = true
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: "ncloud_vpc"}
		r.References["subnet_no"] = ujconfig.Reference{TerraformName: tfSubnet}
		r.References["config_group_no"] = ujconfig.Reference{TerraformName: "ncloud_redis_config_group"}
	})
	p.AddResourceConfigurator("ncloud_redis_config_group", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupDatabase
		r.Kind = "RedisConfigGroup"
	})
}
