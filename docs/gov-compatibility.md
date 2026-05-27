# Ncloud Gov Compatibility

This provider exposes `site` on ProviderConfig and passes it directly to the Terraform provider. Set `site: gov` for NAVER Cloud Platform Government. The Terraform provider maps this to the government API gateway and uses a government Object Storage endpoint pattern for buckets.

## Provider Settings

| Setting | Public | Gov | Fin | Notes |
| --- | --- | --- | --- | --- |
| `region` | Required | Required | Required | Region values are validated by the Terraform provider at runtime. |
| `site` | `public` | `gov` | `fin` | Defaults to `public` when omitted. |
| `support_vpc` | Always true | Always true | Always true | Classic is not supported by Terraform provider `v4.0.5`. |
| credentials JSON | Required | Required | Required | Secret key `credentials` must contain `access_key` and `secret_key`. |

## Resource Coverage

| API Group | Resources | Gov Status |
| --- | --- | --- |
| `network.ncloud.m.crossplane.io` | `Vpc`, `Subnet`, `NetworkACL`, `NetworkACLDenyAllowGroup`, `NetworkACLRule`, `RouteTable`, `Route`, `RouteTableAssociation`, `NatGateway`, `VpcPeering` | VPC smoke-tested with gov credentials. Other resources need service-specific acceptance coverage. |
| `compute.ncloud.m.crossplane.io` | `LoginKey`, `InitScript`, `Server`, `NetworkInterface`, `PublicIP`, `BlockStorage`, `BlockStorageSnapshot`, `LaunchConfiguration`, `PlacementGroup`, `AccessControlGroup`, `AccessControlGroupRule` | Generated and apply-tested. Full creation needs image/spec/key/subnet inputs and cost controls. |
| `loadbalancer.ncloud.m.crossplane.io` | `Lb`, `LbListener`, `LbTargetGroup`, `LbTargetGroupAttachment` | Generated and apply-tested. Full creation needs subnet and target dependency coverage. |
| `objectstorage.ncloud.m.crossplane.io` | `ObjectstorageBucket`, `ObjectstorageBucketACL`, `ObjectstorageObject`, `ObjectstorageObjectACL`, `ObjectstorageObjectCopy` | Generated and apply-tested. Actual gov S3 API create currently requires compatible Object Storage credentials/permissions. |
| `autoscaling.ncloud.m.crossplane.io` | `AutoScalingGroup`, `AutoScalingPolicy`, `AutoScalingSchedule` | Generated and apply-tested. Full creation can create billable servers. |
| `database.ncloud.m.crossplane.io` | `Mongodb`, `MongodbUsers`, `Mssql`, `Mysql`, `MysqlDatabases`, `MysqlRecovery`, `MysqlSlave`, `MysqlUsers`, `Postgresql`, `PostgresqlDatabases`, `PostgresqlReadReplica`, `PostgresqlUsers`, `Redis`, `RedisConfigGroup` | Generated and apply-tested. Gov field differences require dedicated acceptance coverage. |
| `analytics.ncloud.m.crossplane.io` | `CdssCluster`, `CdssConfigGroup`, `Hadoop`, `SesCluster` | Generated and apply-tested. Full creation is expensive and needs dedicated fixtures. |
| `nks.ncloud.m.crossplane.io` | `NksCluster`, `NksNodePool` | Generated and apply-tested. Full creation is expensive and needs subnet/node-pool fixtures. |
| `nas.ncloud.m.crossplane.io` | `NasVolume` | Generated and apply-tested. Full creation is billable. |
| `source.ncloud.m.crossplane.io` | `SourcebuildProject`, `SourcecommitRepository`, `SourcedeployProject`, `SourcedeployProjectStage`, `SourcedeployProjectStageScenario`, `SourcepipelineProject` | Generated and apply-tested. DevTools are expected to be public-site only unless Ncloud documents gov support. |

Managed resources are namespaced only and use Crossplane v2 groups ending in
`ncloud.m.crossplane.io`. They may use either a namespace-local
`ncloud.m.crossplane.io/v1beta1` `ProviderConfig` or a shared
`ClusterProviderConfig`.

## Known Gov Differences To Respect Later

- DevTools resources such as SourceBuild, SourceCommit, SourceDeploy, and SourcePipeline are documented as public-site only, so gov tests should skip them unless Ncloud documents gov support.
- PostgreSQL has gov-specific field gaps, including storage encryption and multi-zone related fields.
- MySQL slave/recovery multi-zone subnet behavior differs in gov.
- Redis user fields differ for gov and should be added only with dedicated acceptance coverage.
- NKS has mixed site-specific behavior and should be added in a separate phase.

## Acceptance Checks

Run public and gov smoke tests separately. For gov:

```console
UPTEST_NCLOUD_REGION=KR \
UPTEST_NCLOUD_SITE=gov \
UPTEST_CLOUD_CREDENTIALS='{"access_key":"...","secret_key":"..."}' \
make e2e
```

The first gov acceptance path should create and delete a minimal VPC, Subnet, Server, PublicIP, NatGateway, Route, Load Balancer, and Object Storage bucket.
The repository includes minimal VPC and Object Storage manifests at
`examples/namespaced/network/vpc.yaml`, and
`examples/namespaced/objectstorage/objectstoragebucket.yaml`. These examples can be used
for public and gov by changing only the setup script environment values.
