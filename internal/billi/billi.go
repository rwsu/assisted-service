package billi

import (
	"context"

	hiveext "github.com/openshift/assisted-service/api/hiveextension/v1beta1"
	"github.com/openshift/assisted-service/internal/controller/controllers"
	hivev1 "github.com/openshift/hive/apis/hive/v1"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func LoadZTPmanifests(logger logrus.FieldLogger, cdReconciler controllers.ClusterDeploymentsReconciler, infraEnvReconciler controllers.InfraEnvReconciler) {
	//
	// TODO: replace this block with reading the ClusterDeployment
	// and AgentClusterInstall CRs from the filesystem.
	clusterName := "cluster0"
	pullSecretName := "pullSecretName"
	aciName := "aciName"
	clusterDeploymentName := "cluster0Deployment"
	namespace := "cluster0"

	spec := hivev1.ClusterDeploymentSpec{
		BaseDomain:  "hive.example.com",
		ClusterName: clusterName,
		PullSecretRef: &corev1.LocalObjectReference{
			Name: pullSecretName,
		},
		ClusterInstallRef: &hivev1.ClusterInstallLocalReference{
			Group:   hiveext.Group,
			Version: hiveext.Version,
			Kind:    "AgentClusterInstall",
			Name:    aciName,
		},
	}

	cd := &hivev1.ClusterDeployment{
		Spec: spec,
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterDeployment",
			APIVersion: "hive.openshift.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterDeploymentName,
			Namespace: namespace,
		},
	}

	aci := &hiveext.AgentClusterInstall{}

	nName := types.NamespacedName{
		Namespace: "cluster0",
		Name:      "cluster0",
	}
	// end of TODO
	//

	// pullSecret := "change-me"
	ctx := context.Background()

	logger.Info("RWSU before billi create InfraEnv")
	// infraEnvReconciler.create
	logger.Info("RWSU after billi create InfraEnv")

	logger.Info("RWSU before billi create Cluster")
	cdReconciler.CreateNewCluster(ctx, logger, nName, cd, aci)
	logger.Info("RWSU after billi create Cluster")
}
