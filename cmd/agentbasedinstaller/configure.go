package agentbasedinstaller

import (
	"context"
	"fmt"

	"github.com/go-openapi/strfmt"
	hiveext "github.com/openshift/assisted-service/api/hiveextension/v1beta1"
	"github.com/openshift/assisted-service/client"
	"github.com/openshift/assisted-service/client/installer"
	"github.com/openshift/assisted-service/models"
	errorutil "github.com/openshift/assisted-service/pkg/error"
	log "github.com/sirupsen/logrus"
)

func AlignHostRolesToControlPlaneReplicas(ctx context.Context, log *log.Logger, bmInventory *client.AssistedInstall, agentClusterInstallPath string, infraEnvID strfmt.UUID) error {
	var aci hiveext.AgentClusterInstall
	if aciErr := getFileData(agentClusterInstallPath, &aci); aciErr != nil {
		return aciErr
	}

	log.Infof("RWSU control plane replicas %v", aci.Spec.ProvisionRequirements.ControlPlaneAgents)
	// Roles are properly auto assigned in HA (3 control plane nodes) and SNO topologies
	if aci.Spec.ProvisionRequirements.ControlPlaneAgents == 3 || aci.Spec.ProvisionRequirements.ControlPlaneAgents == 1 {

		return nil
	}

	hostList, err := bmInventory.Installer.V2ListHosts(ctx, installer.NewV2ListHostsParams().WithInfraEnvID(infraEnvID))
	if err != nil {
		return fmt.Errorf("Failed to list hosts: %w", errorutil.GetAssistedError(err))
	}

	// Count number of hosts already assigned as master
	numHostsAlreadyAssignedToControlPlane := 0
	for _, host := range hostList.Payload {
		log.Infof("RWSU Host %v currently has role %v", host.ID, host.Role)
		if host.Role == "master" {
			numHostsAlreadyAssignedToControlPlane++
		}
	}

	numHostsToAssignedToControlPlane := aci.Spec.ProvisionRequirements.ControlPlaneAgents - numHostsAlreadyAssignedToControlPlane
	log.Infof("RWSU numHostsAlreadyAssignedToControlPlane %v", numHostsAlreadyAssignedToControlPlane)
	log.Infof("RWSU numHostsToAssignedToControlPlane %v", numHostsToAssignedToControlPlane)
	if numHostsToAssignedToControlPlane == 0 {
		return nil
	}

	log.Infof("Assigning host roles to match number of desired control plane replicas of %v", aci.Spec.ProvisionRequirements.ControlPlaneAgents)

	for _, host := range hostList.Payload {
		if numHostsToAssignedToControlPlane == 0 {
			return nil
		}
		if len(host.Inventory) == 0 {
			log.Info("Inventory information not yet available")
			return nil
		}

		inventory := &models.Inventory{}
		err := inventory.UnmarshalBinary([]byte(host.Inventory))
		if err != nil {
			return fmt.Errorf("failed to unmarshal host inventory: %w", err)
		}

		log.Infof("Host %v currently has role %v", host.ID, host.Role)
		controlPlaneRole := "master"
		if host.Role != models.HostRole(controlPlaneRole) {
			updateParams := &models.HostUpdateParams{}
			changed := applyRole(log, host, inventory, &controlPlaneRole, updateParams)
			if changed {
				err := updateHost(ctx, bmInventory, inventory, host, updateParams)
				if err != nil {
					return err
				}
			}
			log.Infof("Host %v assigned to role %v ", host.ID, controlPlaneRole)
			numHostsToAssignedToControlPlane--
		}
	}
	return nil
}
