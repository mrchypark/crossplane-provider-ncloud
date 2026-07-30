package subaccount

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

const shortGroupSubAccount = "subaccount"

// Configure adds Sub Account resource configuration.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("ncloud_subaccount", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupSubAccount
		r.Kind = "Subaccount"
	})
	p.AddResourceConfigurator("ncloud_subaccount_access_key", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupSubAccount
		r.Kind = "SubaccountAccessKey"
		r.References["sub_account_id"] = ujconfig.Reference{TerraformName: "ncloud_subaccount"}
	})
}
