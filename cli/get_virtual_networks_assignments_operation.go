// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/virtual_network_assignments"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsCmd returns a cmd to handle operation getVirtualNetworksAssignments
func makeOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "get-virtual-networks-assignments",
		Short: `Returns a list of all servers assigned to virtual networks.
`,
		RunE: runOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignments,
	}

	if err := registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignments uses cmd flags to call endpoint api
func runOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignments(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := virtual_network_assignments.NewGetVirtualNetworksAssignmentsParams()
	if err, _ := retrieveOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterServerFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterVidFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterVirtualNetworkIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsResult(appCli.VirtualNetworkAssignments.GetVirtualNetworksAssignments(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsParamFlags registers all flags needed to fill params
func registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterServerParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterVidParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterVirtualNetworkIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterServerParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterServerDescription := `The server ID to filter by`

	var filterServerFlagName string
	if cmdPrefix == "" {
		filterServerFlagName = "filter[server]"
	} else {
		filterServerFlagName = fmt.Sprintf("%v.filter[server]", cmdPrefix)
	}

	var filterServerFlagDefault string

	_ = cmd.PersistentFlags().String(filterServerFlagName, filterServerFlagDefault, filterServerDescription)

	return nil
}
func registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterVidParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterVidDescription := `The vlan ID to filter by`

	var filterVidFlagName string
	if cmdPrefix == "" {
		filterVidFlagName = "filter[vid]"
	} else {
		filterVidFlagName = fmt.Sprintf("%v.filter[vid]", cmdPrefix)
	}

	var filterVidFlagDefault string

	_ = cmd.PersistentFlags().String(filterVidFlagName, filterVidFlagDefault, filterVidDescription)

	return nil
}
func registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterVirtualNetworkIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterVirtualNetworkIdDescription := `The virtual network ID to filter by`

	var filterVirtualNetworkIdFlagName string
	if cmdPrefix == "" {
		filterVirtualNetworkIdFlagName = "filter[virtual_network_id]"
	} else {
		filterVirtualNetworkIdFlagName = fmt.Sprintf("%v.filter[virtual_network_id]", cmdPrefix)
	}

	var filterVirtualNetworkIdFlagDefault string

	_ = cmd.PersistentFlags().String(filterVirtualNetworkIdFlagName, filterVirtualNetworkIdFlagDefault, filterVirtualNetworkIdDescription)

	return nil
}

func retrieveOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsAPIVersionFlag(m *virtual_network_assignments.GetVirtualNetworksAssignmentsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterServerFlag(m *virtual_network_assignments.GetVirtualNetworksAssignmentsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[server]") {

		var filterServerFlagName string
		if cmdPrefix == "" {
			filterServerFlagName = "filter[server]"
		} else {
			filterServerFlagName = fmt.Sprintf("%v.filter[server]", cmdPrefix)
		}

		filterServerFlagValue, err := cmd.Flags().GetString(filterServerFlagName)
		if err != nil {
			return err, false
		}
		m.FilterServer = &filterServerFlagValue

	}
	return nil, retAdded
}
func retrieveOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterVidFlag(m *virtual_network_assignments.GetVirtualNetworksAssignmentsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[vid]") {

		var filterVidFlagName string
		if cmdPrefix == "" {
			filterVidFlagName = "filter[vid]"
		} else {
			filterVidFlagName = fmt.Sprintf("%v.filter[vid]", cmdPrefix)
		}

		filterVidFlagValue, err := cmd.Flags().GetString(filterVidFlagName)
		if err != nil {
			return err, false
		}
		m.FilterVid = &filterVidFlagValue

	}
	return nil, retAdded
}
func retrieveOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterVirtualNetworkIDFlag(m *virtual_network_assignments.GetVirtualNetworksAssignmentsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[virtual_network_id]") {

		var filterVirtualNetworkIdFlagName string
		if cmdPrefix == "" {
			filterVirtualNetworkIdFlagName = "filter[virtual_network_id]"
		} else {
			filterVirtualNetworkIdFlagName = fmt.Sprintf("%v.filter[virtual_network_id]", cmdPrefix)
		}

		filterVirtualNetworkIdFlagValue, err := cmd.Flags().GetString(filterVirtualNetworkIdFlagName)
		if err != nil {
			return err, false
		}
		m.FilterVirtualNetworkID = &filterVirtualNetworkIdFlagValue

	}
	return nil, retAdded
}

// parseOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsResult parses request result and return the string content
func parseOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsResult(resp0 *virtual_network_assignments.GetVirtualNetworksAssignmentsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*virtual_network_assignments.GetVirtualNetworksAssignmentsOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
