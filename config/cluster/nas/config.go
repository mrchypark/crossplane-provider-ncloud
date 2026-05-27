package nas

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds NAS resource configuration.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("ncloud_nas_volume", func(r *ujconfig.Resource) {
		r.ShortGroup = "nas"
		r.Kind = "NasVolume"
		r.UseAsync = true
	})
}
