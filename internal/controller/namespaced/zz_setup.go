// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cdssconfiggroup "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/analytics/cdssconfiggroup"
	hadoop "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/analytics/hadoop"
	sescluster "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/analytics/sescluster"
	autoscalinggroup "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/autoscaling/autoscalinggroup"
	autoscalingpolicy "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/autoscaling/autoscalingpolicy"
	autoscalingschedule "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/autoscaling/autoscalingschedule"
	accesscontrolgroup "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/compute/accesscontrolgroup"
	accesscontrolgrouprule "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/compute/accesscontrolgrouprule"
	blockstorage "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/compute/blockstorage"
	blockstoragesnapshot "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/compute/blockstoragesnapshot"
	initscript "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/compute/initscript"
	launchconfiguration "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/compute/launchconfiguration"
	loginkey "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/compute/loginkey"
	networkinterface "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/compute/networkinterface"
	placementgroup "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/compute/placementgroup"
	publicip "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/compute/publicip"
	server "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/compute/server"
	mongodb "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/database/mongodb"
	mongodbusers "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/database/mongodbusers"
	mssql "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/database/mssql"
	mysql "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/database/mysql"
	mysqldatabases "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/database/mysqldatabases"
	mysqlrecovery "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/database/mysqlrecovery"
	mysqlslave "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/database/mysqlslave"
	mysqlusers "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/database/mysqlusers"
	postgresql "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/database/postgresql"
	postgresqldatabases "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/database/postgresqldatabases"
	postgresqlreadreplica "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/database/postgresqlreadreplica"
	postgresqlusers "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/database/postgresqlusers"
	redis "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/database/redis"
	redisconfiggroup "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/database/redisconfiggroup"
	lb "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/loadbalancer/lb"
	lblistener "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/loadbalancer/lblistener"
	lbtargetgroup "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/loadbalancer/lbtargetgroup"
	lbtargetgroupattachment "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/loadbalancer/lbtargetgroupattachment"
	nasvolume "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/nas/nasvolume"
	natgateway "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/network/natgateway"
	networkacl "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/network/networkacl"
	networkacldenyallowgroup "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/network/networkacldenyallowgroup"
	networkaclrule "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/network/networkaclrule"
	route "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/network/route"
	routetable "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/network/routetable"
	routetableassociation "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/network/routetableassociation"
	subnet "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/network/subnet"
	vpc "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/network/vpc"
	vpcpeering "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/network/vpcpeering"
	nkscluster "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/nks/nkscluster"
	nksnodepool "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/nks/nksnodepool"
	objectstoragebucket "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/objectstorage/objectstoragebucket"
	objectstoragebucketacl "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/objectstorage/objectstoragebucketacl"
	objectstorageobject "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/objectstorage/objectstorageobject"
	objectstorageobjectacl "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/objectstorage/objectstorageobjectacl"
	objectstorageobjectcopy "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/objectstorage/objectstorageobjectcopy"
	providerconfig "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/providerconfig"
	sourcebuildproject "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/source/sourcebuildproject"
	sourcecommitrepository "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/source/sourcecommitrepository"
	sourcedeployproject "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/source/sourcedeployproject"
	sourcedeployprojectstage "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/source/sourcedeployprojectstage"
	sourcedeployprojectstagescenario "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/source/sourcedeployprojectstagescenario"
	sourcepipelineproject "github.com/mrchypark/crossplane-provider-ncloud/internal/controller/namespaced/source/sourcepipelineproject"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cdssconfiggroup.Setup,
		hadoop.Setup,
		sescluster.Setup,
		autoscalinggroup.Setup,
		autoscalingpolicy.Setup,
		autoscalingschedule.Setup,
		accesscontrolgroup.Setup,
		accesscontrolgrouprule.Setup,
		blockstorage.Setup,
		blockstoragesnapshot.Setup,
		initscript.Setup,
		launchconfiguration.Setup,
		loginkey.Setup,
		networkinterface.Setup,
		placementgroup.Setup,
		publicip.Setup,
		server.Setup,
		mongodb.Setup,
		mongodbusers.Setup,
		mssql.Setup,
		mysql.Setup,
		mysqldatabases.Setup,
		mysqlrecovery.Setup,
		mysqlslave.Setup,
		mysqlusers.Setup,
		postgresql.Setup,
		postgresqldatabases.Setup,
		postgresqlreadreplica.Setup,
		postgresqlusers.Setup,
		redis.Setup,
		redisconfiggroup.Setup,
		lb.Setup,
		lblistener.Setup,
		lbtargetgroup.Setup,
		lbtargetgroupattachment.Setup,
		nasvolume.Setup,
		natgateway.Setup,
		networkacl.Setup,
		networkacldenyallowgroup.Setup,
		networkaclrule.Setup,
		route.Setup,
		routetable.Setup,
		routetableassociation.Setup,
		subnet.Setup,
		vpc.Setup,
		vpcpeering.Setup,
		nkscluster.Setup,
		nksnodepool.Setup,
		objectstoragebucket.Setup,
		objectstoragebucketacl.Setup,
		objectstorageobject.Setup,
		objectstorageobjectacl.Setup,
		objectstorageobjectcopy.Setup,
		providerconfig.Setup,
		sourcebuildproject.Setup,
		sourcecommitrepository.Setup,
		sourcedeployproject.Setup,
		sourcedeployprojectstage.Setup,
		sourcedeployprojectstagescenario.Setup,
		sourcepipelineproject.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cdssconfiggroup.SetupGated,
		hadoop.SetupGated,
		sescluster.SetupGated,
		autoscalinggroup.SetupGated,
		autoscalingpolicy.SetupGated,
		autoscalingschedule.SetupGated,
		accesscontrolgroup.SetupGated,
		accesscontrolgrouprule.SetupGated,
		blockstorage.SetupGated,
		blockstoragesnapshot.SetupGated,
		initscript.SetupGated,
		launchconfiguration.SetupGated,
		loginkey.SetupGated,
		networkinterface.SetupGated,
		placementgroup.SetupGated,
		publicip.SetupGated,
		server.SetupGated,
		mongodb.SetupGated,
		mongodbusers.SetupGated,
		mssql.SetupGated,
		mysql.SetupGated,
		mysqldatabases.SetupGated,
		mysqlrecovery.SetupGated,
		mysqlslave.SetupGated,
		mysqlusers.SetupGated,
		postgresql.SetupGated,
		postgresqldatabases.SetupGated,
		postgresqlreadreplica.SetupGated,
		postgresqlusers.SetupGated,
		redis.SetupGated,
		redisconfiggroup.SetupGated,
		lb.SetupGated,
		lblistener.SetupGated,
		lbtargetgroup.SetupGated,
		lbtargetgroupattachment.SetupGated,
		nasvolume.SetupGated,
		natgateway.SetupGated,
		networkacl.SetupGated,
		networkacldenyallowgroup.SetupGated,
		networkaclrule.SetupGated,
		route.SetupGated,
		routetable.SetupGated,
		routetableassociation.SetupGated,
		subnet.SetupGated,
		vpc.SetupGated,
		vpcpeering.SetupGated,
		nkscluster.SetupGated,
		nksnodepool.SetupGated,
		objectstoragebucket.SetupGated,
		objectstoragebucketacl.SetupGated,
		objectstorageobject.SetupGated,
		objectstorageobjectacl.SetupGated,
		objectstorageobjectcopy.SetupGated,
		providerconfig.SetupGated,
		sourcebuildproject.SetupGated,
		sourcecommitrepository.SetupGated,
		sourcedeployproject.SetupGated,
		sourcedeployprojectstage.SetupGated,
		sourcedeployprojectstagescenario.SetupGated,
		sourcepipelineproject.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
