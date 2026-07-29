# crossplane-provider-ncloud

`crossplane-provider-ncloud` is an Upjet-based Crossplane provider for NAVER Cloud Platform.

This implementation targets Terraform provider `NaverCloudPlatform/ncloud` `v4.0.6-csap.1` and generates Crossplane v2-style namespaced managed resources for every managed resource exposed by the Terraform provider.

`v4.0.6-csap.1` is a build of [`hiidy/terraform-provider-ncloud`](https://github.com/hiidy/terraform-provider-ncloud/releases/tag/v4.0.6-csap.1) — upstream `v4.0.6` plus the Sub Account resources from [upstream PR #594](https://github.com/NaverCloudPlatform/terraform-provider-ncloud/pull/594), which has not been merged yet. The Terraform Registry does not carry it, so the provider is installed from that release into a local filesystem mirror and verified against the digests pinned in the `Makefile`. Once #594 is released upstream, set `TERRAFORM_PROVIDER_USE_MIRROR=false`, point the `TERRAFORM_PROVIDER_*` addresses back at `NaverCloudPlatform`, and regenerate.

## ProviderConfig

Credentials are read from a Secret whose `credentials` key contains JSON:

```json
{
  "access_key": "NCLOUD_ACCESS_KEY",
  "secret_key": "NCLOUD_SECRET_KEY"
}
```

The ProviderConfig supplies region and site. Shared credentials can use a cluster-wide `ClusterProviderConfig`:

```yaml
apiVersion: ncloud.crossplane.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default
spec:
  region: KR
  site: public
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: provider-secret
      key: credentials
```

Supported `site` values are `public`, `gov`, and `fin`. The provider always configures Terraform with `support_vpc: true` because the current Terraform provider no longer supports the classic environment.

## Development

Run the generation pipeline:

```console
make generate
```

Run tests:

```console
go test ./...
```

Build the provider binary:

```console
make build
```

Build the Crossplane package:

```console
make xpkg.build
```
