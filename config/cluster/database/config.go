package database

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds database and cache resource configuration.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("ncloud_mongodb", func(r *ujconfig.Resource) {
		r.ShortGroup = "database"
		r.Kind = "Mongodb"
		r.UseAsync = true
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: "ncloud_vpc"}
		r.References["subnet_no"] = ujconfig.Reference{TerraformName: "ncloud_subnet"}
	})
	p.AddResourceConfigurator("ncloud_mongodb_users", func(r *ujconfig.Resource) {
		r.ShortGroup = "database"
		r.Kind = "MongodbUsers"
		r.References["id"] = ujconfig.Reference{TerraformName: "ncloud_mongodb"}
	})
	p.AddResourceConfigurator("ncloud_mssql", func(r *ujconfig.Resource) {
		r.ShortGroup = "database"
		r.Kind = "Mssql"
		r.UseAsync = true
		r.References["subnet_no"] = ujconfig.Reference{TerraformName: "ncloud_subnet"}
	})
	p.AddResourceConfigurator("ncloud_mysql", func(r *ujconfig.Resource) {
		r.ShortGroup = "database"
		r.Kind = "Mysql"
		r.UseAsync = true
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: "ncloud_vpc"}
		r.References["subnet_no"] = ujconfig.Reference{TerraformName: "ncloud_subnet"}
	})
	p.AddResourceConfigurator("ncloud_mysql_databases", func(r *ujconfig.Resource) {
		r.ShortGroup = "database"
		r.Kind = "MysqlDatabases"
		r.References["mysql_instance_no"] = ujconfig.Reference{TerraformName: "ncloud_mysql"}
	})
	p.AddResourceConfigurator("ncloud_mysql_recovery", func(r *ujconfig.Resource) {
		r.ShortGroup = "database"
		r.Kind = "MysqlRecovery"
		r.UseAsync = true
		r.References["mysql_instance_no"] = ujconfig.Reference{TerraformName: "ncloud_mysql"}
	})
	p.AddResourceConfigurator("ncloud_mysql_slave", func(r *ujconfig.Resource) {
		r.ShortGroup = "database"
		r.Kind = "MysqlSlave"
		r.UseAsync = true
		r.References["mysql_instance_no"] = ujconfig.Reference{TerraformName: "ncloud_mysql"}
	})
	p.AddResourceConfigurator("ncloud_mysql_users", func(r *ujconfig.Resource) {
		r.ShortGroup = "database"
		r.Kind = "MysqlUsers"
		r.References["mysql_instance_no"] = ujconfig.Reference{TerraformName: "ncloud_mysql"}
	})
	p.AddResourceConfigurator("ncloud_postgresql", func(r *ujconfig.Resource) {
		r.ShortGroup = "database"
		r.Kind = "Postgresql"
		r.UseAsync = true
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: "ncloud_vpc"}
		r.References["subnet_no"] = ujconfig.Reference{TerraformName: "ncloud_subnet"}
	})
	p.AddResourceConfigurator("ncloud_postgresql_databases", func(r *ujconfig.Resource) {
		r.ShortGroup = "database"
		r.Kind = "PostgresqlDatabases"
		r.References["id"] = ujconfig.Reference{TerraformName: "ncloud_postgresql"}
	})
	p.AddResourceConfigurator("ncloud_postgresql_read_replica", func(r *ujconfig.Resource) {
		r.ShortGroup = "database"
		r.Kind = "PostgresqlReadReplica"
		r.UseAsync = true
		r.References["postgresql_instance_no"] = ujconfig.Reference{TerraformName: "ncloud_postgresql"}
	})
	p.AddResourceConfigurator("ncloud_postgresql_users", func(r *ujconfig.Resource) {
		r.ShortGroup = "database"
		r.Kind = "PostgresqlUsers"
		r.References["id"] = ujconfig.Reference{TerraformName: "ncloud_postgresql"}
	})
	p.AddResourceConfigurator("ncloud_redis", func(r *ujconfig.Resource) {
		r.ShortGroup = "database"
		r.Kind = "Redis"
		r.UseAsync = true
		r.References["vpc_no"] = ujconfig.Reference{TerraformName: "ncloud_vpc"}
		r.References["subnet_no"] = ujconfig.Reference{TerraformName: "ncloud_subnet"}
		r.References["config_group_no"] = ujconfig.Reference{TerraformName: "ncloud_redis_config_group"}
	})
	p.AddResourceConfigurator("ncloud_redis_config_group", func(r *ujconfig.Resource) {
		r.ShortGroup = "database"
		r.Kind = "RedisConfigGroup"
	})
}
