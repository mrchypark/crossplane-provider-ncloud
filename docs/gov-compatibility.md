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
| `network.ncloud.m.crossplane.io` | `Vpc`, `Subnet`, `NetworkACL`, `NetworkACLDenyAllowGroup`, `NetworkACLRule`, `RouteTable`, `Route`, `RouteTableAssociation`, `NatGateway`, `VpcPeering` | Real gov creation tested and deleted in the local `kind-provider-ncloud-test` cluster. |
| `compute.ncloud.m.crossplane.io` | `LoginKey`, `InitScript`, `Server`, `NetworkInterface`, `PublicIP`, `BlockStorage`, `BlockStorageSnapshot`, `LaunchConfiguration`, `PlacementGroup`, `AccessControlGroup`, `AccessControlGroupRule` | Real gov creation tested and deleted. |
| `loadbalancer.ncloud.m.crossplane.io` | `Lb`, `LbListener`, `LbTargetGroup`, `LbTargetGroupAttachment` | Real gov creation tested and deleted after a local Terraform provider read-only field patch for `ncloud_lb`. |
| `objectstorage.ncloud.m.crossplane.io` | `ObjectstorageBucket`, `ObjectstorageBucketACL`, `ObjectstorageObject`, `ObjectstorageObjectACL`, `ObjectstorageObjectCopy` | Real gov creation tested and deleted after Object Storage service subscription. Object/object-copy require `sourceSecretRef` because Terraform `source` is sensitive. |
| `autoscaling.ncloud.m.crossplane.io` | `AutoScalingGroup`, `AutoScalingPolicy`, `AutoScalingSchedule` | Real gov creation tested and deleted. |
| `database.ncloud.m.crossplane.io` | `Mongodb`, `MongodbUsers`, `Mssql`, `Mysql`, `MysqlDatabases`, `MysqlRecovery`, `MysqlSlave`, `MysqlUsers`, `Postgresql`, `PostgresqlDatabases`, `PostgresqlReadReplica`, `PostgresqlUsers`, `Redis`, `RedisConfigGroup` | Real gov creation tested for Redis, MongoDB, MSSQL, MySQL, and PostgreSQL families after Cloud Log Analytics subscription. `MysqlSlave` and `PostgresqlReadReplica` need high-availability parent fixtures. `MysqlRecovery` needs point-in-time restore data. |
| `analytics.ncloud.m.crossplane.io` | `CdssConfigGroup`, `Hadoop`, `SesCluster` | `CdssConfigGroup` and `Hadoop` were real-created. `CdssCluster` is intentionally excluded from the provider after Gov API rejected multiple generation-code combinations. `SesCluster` reaches Gov create after Cloud Log Analytics subscription, but Gov rejects tested generation-code combinations. |
| `nks.ncloud.m.crossplane.io` | `NksCluster`, `NksNodePool` | Real gov creation tested and deleted after local Terraform provider read/not-found handling patches. |
| `nas.ncloud.m.crossplane.io` | `NasVolume` | Real gov creation tested and deleted. |
| `source.ncloud.m.crossplane.io` | `SourcebuildProject`, `SourcecommitRepository`, `SourcedeployProject`, `SourcedeployProjectStage`, `SourcedeployProjectStageScenario`, `SourcepipelineProject` | Real gov creation tested and deleted. SourceCommit, SourceBuild, SourcePipeline required local Terraform provider read/not-found and empty `linked_tasks` handling patches. |

Managed resources are namespaced only and use Crossplane v2 groups ending in
`ncloud.m.crossplane.io`. They may use either a namespace-local
`ncloud.m.crossplane.io/v1beta1` `ProviderConfig` or a shared
`ClusterProviderConfig`.

## Known Gov Differences To Respect Later

- SourceCommit, SourceBuild, SourceDeploy, and SourcePipeline are reachable in gov with the tested account, even though older planning notes expected public-only behavior. Keep gov coverage enabled for these resources.
- Object Storage Gov uses endpoint `https://kr.object.gov-ncloudstorage.com` and signing region `gov-standard` for `KR`. Upstream Terraform provider `v4.0.5` signs S3 requests with the provider `region` value, so a native provider patch is needed for Gov Object Storage.
- Object Storage returns `InvalidAccessKeyId` when the API key is not present in the Object Storage credential database. After Object Storage service subscription, the same preflight returned `objectstorage_s3_listbuckets_http=200`.
- Cloud Log Analytics is a hard precondition for Cloud DB for MySQL, PostgreSQL, MongoDB, MSSQL, and Search Engine Service creation in gov. After console subscription, the local preflight returned `cla_capacity_http=200 body_bytes=287 code=0 message=SUCCESS`.
- Hadoop Gov creation requires catalog-compatible image and engine codes. The tested fixture used `SW.VCHDP.OS.LNX64.ROCKY.0810.HDP.B050` with engine `HADOOP2.2`.
- Search Engine Service reached the Gov create API after Cloud Log Analytics subscription, but Gov returned code `10140` for tested Rocky 8.6 and CentOS 7.8 generation/catalog combinations.
- PostgreSQL has gov-specific field gaps, including storage encryption and multi-zone related fields.
- MySQL slave/recovery multi-zone subnet behavior differs in gov.
- Redis user fields differ for gov and should be added only with dedicated acceptance coverage.
- NKS has mixed site-specific behavior and should be added in a separate phase.

## Local Acceptance Snapshot

As of 2026-06-01, local Gov acceptance covers all 59 generated managed
resource kinds after excluding `CdssCluster`:

| Result | Count | Notes |
| --- | ---: | --- |
| Ready | 55 | Real-created in gov and reached `Ready=True`. Most were deleted; `Hadoop` cleanup can remain pending while Ncloud reports an internal setup operation. |
| Not ready | 4 | `MysqlSlave`, `MysqlRecovery`, `PostgresqlReadReplica`, and `SesCluster` need additional fixtures or Gov catalog clarification. |
| Missing | 0 | Every generated kind has an explicit result row. |

The latest cleanup check showed `Hadoop` deletion pending while Ncloud reports
`settingUp/SETUP`, with its dependent VPC/Subnet/LoginKey/NetworkACL/Object
Storage resources intentionally retained until the Hadoop external resource is
gone. The billing snapshot hash did not change.

Before running Gov acceptance, run the local gov preflight:

```console
scripts/gov-acceptance-preflight.sh
```

It validates account-level prerequisites without creating managed resources. The
latest run returned `objectstorage_s3_listbuckets_http=200 error_code=none` and
`cla_capacity_http=200 body_bytes=287 code=0 message=SUCCESS`, so Object
Storage and Cloud Log Analytics are available for the tested account.

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
