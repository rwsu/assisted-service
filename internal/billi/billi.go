package billi

import (
	"context"
	"os"
	"time"

	hiveext "github.com/openshift/assisted-service/api/hiveextension/v1beta1"
	aiv1beta1 "github.com/openshift/assisted-service/api/v1beta1"
	"github.com/openshift/assisted-service/internal/bminventory"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/controller/controllers"
	"github.com/openshift/assisted-service/restapi/operations/installer"
	hivev1 "github.com/openshift/hive/apis/hive/v1"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/yaml"
)

func LoadZTPmanifests(logger logrus.FieldLogger, installerInternals bminventory.InstallerInternals, db *gorm.DB) {
	logger.Info("Loading cluster definitions")

	// pull secret
	secretData, err := os.ReadFile("/data/pull-secret.yaml")
	if err != nil {
		logger.Fatal(err)
	}
	var secret corev1.Secret
	if err := yaml.Unmarshal(secretData, &secret); err != nil {
		logger.Fatal(err)
	}
	pullSecret := secret.StringData[".dockerconfigjson"]

	// ClusterDeployment
	cdData, err := os.ReadFile("/data/cluster-deployment.yaml")
	if err != nil {
		logger.Fatal(err)
	}
	var cd hivev1.ClusterDeployment
	if err := yaml.Unmarshal(cdData, &cd); err != nil {
		logger.Fatal(err)
	}

	// AgentClusterInstall
	aciData, err := os.ReadFile("/data/agent-cluster-install.yaml")
	if err != nil {
		logger.Fatal(err)
	}
	var aci hiveext.AgentClusterInstall
	if err := yaml.Unmarshal(aciData, &aci); err != nil {
		logger.Fatal(err)
	}

	clusterKey := types.NamespacedName{
		Namespace: cd.Name,
		Name:      cd.Namespace,
	}

	ctx := context.Background()
	releaseImageVersion := "4.10.0-rc.1"
	releaseImageCPUArch := "x86_64"
	clusterParams := controllers.CreateClusterParams(&cd, &aci, pullSecret, releaseImageVersion, releaseImageCPUArch)
	cluster, err := installerInternals.RegisterClusterInternal(ctx, &clusterKey, installer.V2RegisterClusterParams{
		NewClusterParams: clusterParams,
	}, false)
	if err != nil {
		logger.Fatalln(err)
	}

	// TODO: RegisterClusterInternal does not set api_vip in the database.
	// How is api_vip populated normally? For now manually update the db.
	updates := map[string]interface{}{}
	updates["api_vip"] = aci.Spec.APIVIP
	updates["trigger_monitor_timestamp"] = time.Now()
	err = db.Model(&common.Cluster{}).Where("id = ?", cluster.ID.String()).Updates(updates).Error
	if err != nil {
		logger.Fatal(err)
	}

	// InfraEnv
	infraEnvData, err := os.ReadFile("/data/infraenv.yaml")
	if err != nil {
		logger.Fatal(err)
	}
	var infraEnv aiv1beta1.InfraEnv
	if err := yaml.Unmarshal(infraEnvData, &infraEnv); err != nil {
		logger.Fatal(err)
	}

	infraEnvKey := types.NamespacedName{
		Namespace: infraEnv.Namespace,
		Name:      infraEnv.Name,
	}

	infraEnvParams := controllers.CreateInfraEnvParams(&infraEnv, cluster, &infraEnvKey, "full-iso", pullSecret)
	_, err = installerInternals.RegisterInfraEnvInternal(ctx, &infraEnvKey, infraEnvParams)
	if err != nil {
		logger.Fatalln(err)
	}
	logger.Info("Finished loading cluster definitions")
}
