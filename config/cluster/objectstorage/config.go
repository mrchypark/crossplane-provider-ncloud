package objectstorage

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

const shortGroupObjectStorage = "objectstorage"

// Configure adds Object Storage resource configuration.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("ncloud_objectstorage_bucket", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupObjectStorage
		r.Kind = "ObjectstorageBucket"
		ujconfig.MoveToStatus(r.TerraformResource, "creation_date")
	})
	p.AddResourceConfigurator("ncloud_objectstorage_bucket_acl", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupObjectStorage
		r.Kind = "ObjectstorageBucketACL"
		r.References["bucket_name"] = ujconfig.Reference{TerraformName: "ncloud_objectstorage_bucket"}
	})
	p.AddResourceConfigurator("ncloud_objectstorage_object", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupObjectStorage
		r.Kind = "ObjectstorageObject"
		r.References["bucket"] = ujconfig.Reference{TerraformName: "ncloud_objectstorage_bucket"}
	})
	p.AddResourceConfigurator("ncloud_objectstorage_object_acl", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupObjectStorage
		r.Kind = "ObjectstorageObjectACL"
		r.References["object_id"] = ujconfig.Reference{TerraformName: "ncloud_objectstorage_object"}
	})
	p.AddResourceConfigurator("ncloud_objectstorage_object_copy", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroupObjectStorage
		r.Kind = "ObjectstorageObjectCopy"
		r.References["bucket"] = ujconfig.Reference{TerraformName: "ncloud_objectstorage_bucket"}
	})
}
