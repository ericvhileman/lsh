// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/virtual_networks"
	"github.com/latitudesh/cli/internal/utils"
	"github.com/latitudesh/cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationVirtualNetworksGetVirtualNetworkCmd returns a cmd to handle operation getVirtualNetwork
func makeOperationVirtualNetworksGetVirtualNetworkCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "get-virtual-network",
		Short: `Retrieve a Virtual Network.
`,
		RunE: runOperationVirtualNetworksGetVirtualNetwork,
	}

	if err := registerOperationVirtualNetworksGetVirtualNetworkParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVirtualNetworksGetVirtualNetwork uses cmd flags to call endpoint api
func runOperationVirtualNetworksGetVirtualNetwork(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := virtual_networks.NewGetVirtualNetworkParams()
	if err, _ := retrieveOperationVirtualNetworksGetVirtualNetworkAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationVirtualNetworksGetVirtualNetworkIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationVirtualNetworksGetVirtualNetworkResult(appCli.VirtualNetworks.GetVirtualNetwork(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(utils.PrettifyJson(msgStr))
	}
	return nil
}

// registerOperationVirtualNetworksGetVirtualNetworkParamFlags registers all flags needed to fill params
func registerOperationVirtualNetworksGetVirtualNetworkParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVirtualNetworksGetVirtualNetworkAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationVirtualNetworksGetVirtualNetworkIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVirtualNetworksGetVirtualNetworkAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationVirtualNetworksGetVirtualNetworkIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. `

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func retrieveOperationVirtualNetworksGetVirtualNetworkAPIVersionFlag(m *virtual_networks.GetVirtualNetworkParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationVirtualNetworksGetVirtualNetworkIDFlag(m *virtual_networks.GetVirtualNetworkParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

	}
	return nil, retAdded
}

// parseOperationVirtualNetworksGetVirtualNetworkResult parses request result and return the string content
func parseOperationVirtualNetworksGetVirtualNetworkResult(resp0 *virtual_networks.GetVirtualNetworkOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*virtual_networks.GetVirtualNetworkOK)
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

// register flags to command
func registerModelGetVirtualNetworkOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetVirtualNetworkOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetVirtualNetworkOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelVirtualNetworkFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetVirtualNetworkOKBodyFlags(depth int, m *virtual_networks.GetVirtualNetworkOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveGetVirtualNetworkOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveGetVirtualNetworkOKBodyDataFlags(depth int, m *virtual_networks.GetVirtualNetworkOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data models.VirtualNetwork is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.VirtualNetwork{}
	}

	err, dataAdded := retrieveModelVirtualNetworkFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}