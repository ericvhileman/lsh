// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/cli/client/virtual_network_assignments"
	"github.com/latitudesh/cli/internal/api"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/spf13/cobra"
)

// makeOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsCmd returns a cmd to handle operation deleteVirtualNetworksAssignments
func makeOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "destroy",
		Short: `Allow you to remove a Virtual Network assignment.`,
		RunE:  runOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignments,
	}

	if err := registerOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignments uses cmd flags to call endpoint api
func runOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignments(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := virtual_network_assignments.NewDeleteVirtualNetworksAssignmentsParams()
	if err, _ := retrieveOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsAssignmentIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	result, err := appCli.VirtualNetworkAssignments.DeleteVirtualNetworksAssignments(params, nil)
	if err != nil {
		api.RenderErrorOutput(err)
		return nil
	}

	msgStr, err := parseOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsResult(result)
	if err != nil {
		return err
	}
	if !debug {

		utils.PrintOutput(msgStr)
	}
	return nil
}

// registerOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsParamFlags registers all flags needed to fill params
func registerOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsAssignmentIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	apiVersionDescription := ``

	var apiVersionFlagName string
	if cmdPrefix == "" {
		apiVersionFlagName = "API-Version"
	} else {
		apiVersionFlagName = fmt.Sprintf("%v.API-Version", cmdPrefix)
	}

	var apiVersionFlagDefault string = "2023-06-01"

	_ = cmd.PersistentFlags().String(apiVersionFlagName, apiVersionFlagDefault, apiVersionDescription)

	return nil
}
func registerOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsAssignmentIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	assignmentIdDescription := `Required. `

	var assignmentIdFlagName string
	if cmdPrefix == "" {
		assignmentIdFlagName = "assignment_id"
	} else {
		assignmentIdFlagName = fmt.Sprintf("%v.assignment_id", cmdPrefix)
	}

	var assignmentIdFlagDefault string

	_ = cmd.PersistentFlags().String(assignmentIdFlagName, assignmentIdFlagDefault, assignmentIdDescription)
	cmd.MarkPersistentFlagRequired(assignmentIdFlagName)

	return nil
}

func retrieveOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsAPIVersionFlag(m *virtual_network_assignments.DeleteVirtualNetworksAssignmentsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("API-Version") {

		var apiVersionFlagName string
		if cmdPrefix == "" {
			apiVersionFlagName = "API-Version"
		} else {
			apiVersionFlagName = fmt.Sprintf("%v.API-Version", cmdPrefix)
		}

		apiVersionFlagValue, err := cmd.Flags().GetString(apiVersionFlagName)
		if err != nil {
			return err, false
		}
		m.APIVersion = &apiVersionFlagValue

	}
	return nil, retAdded
}
func retrieveOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsAssignmentIDFlag(m *virtual_network_assignments.DeleteVirtualNetworksAssignmentsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("assignment_id") {

		var assignmentIdFlagName string
		if cmdPrefix == "" {
			assignmentIdFlagName = "assignment_id"
		} else {
			assignmentIdFlagName = fmt.Sprintf("%v.assignment_id", cmdPrefix)
		}

		assignmentIdFlagValue, err := cmd.Flags().GetString(assignmentIdFlagName)
		if err != nil {
			return err, false
		}
		m.AssignmentID = assignmentIdFlagValue

	}
	return nil, retAdded
}

// parseOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsResult parses request result and return the string content
func parseOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsResult(resp0 *virtual_network_assignments.DeleteVirtualNetworksAssignmentsNoContent) (string, error) {
	// warning: non schema response deleteVirtualNetworksAssignmentsNoContent is not supported by go-swagger cli yet.

	return "", nil
}
