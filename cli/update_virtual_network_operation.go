// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/virtual_networks"
	"github.com/latitudesh/cli/internal/api"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationVirtualNetworksUpdateVirtualNetworkCmd returns a cmd to handle operation updateVirtualNetwork
func makeOperationVirtualNetworksUpdateVirtualNetworkCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "update",
		Short: `Update a Virtual Network.`,
		RunE: runOperationVirtualNetworksUpdateVirtualNetwork,
	}

	if err := registerOperationVirtualNetworksUpdateVirtualNetworkParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVirtualNetworksUpdateVirtualNetwork uses cmd flags to call endpoint api
func runOperationVirtualNetworksUpdateVirtualNetwork(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := virtual_networks.NewUpdateVirtualNetworkParams()
	if err, _ := retrieveOperationVirtualNetworksUpdateVirtualNetworkAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationVirtualNetworksUpdateVirtualNetworkBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationVirtualNetworksUpdateVirtualNetworkVirtualNetworkIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	result, err := appCli.VirtualNetworks.UpdateVirtualNetwork(params, nil)
	if err != nil {
		api.RenderErrorOutput(err)
		return nil
	}

	msgStr, err := parseOperationVirtualNetworksUpdateVirtualNetworkResult(result)
	if err != nil {
		return err
	}
	if !debug {

		utils.PrintOutput(msgStr)
	}
	return nil
}

// registerOperationVirtualNetworksUpdateVirtualNetworkParamFlags registers all flags needed to fill params
func registerOperationVirtualNetworksUpdateVirtualNetworkParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVirtualNetworksUpdateVirtualNetworkAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationVirtualNetworksUpdateVirtualNetworkBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationVirtualNetworksUpdateVirtualNetworkVirtualNetworkIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVirtualNetworksUpdateVirtualNetworkAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationVirtualNetworksUpdateVirtualNetworkBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelUpdateVirtualNetworkBodyFlags(0, "", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationVirtualNetworksUpdateVirtualNetworkVirtualNetworkIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	virtualNetworkIdDescription := `Required. The Virtual Network ID`

	var virtualNetworkIdFlagName string
	if cmdPrefix == "" {
		virtualNetworkIdFlagName = "virtual_network_id"
	} else {
		virtualNetworkIdFlagName = fmt.Sprintf("%v.virtual_network_id", cmdPrefix)
	}

	var virtualNetworkIdFlagDefault string

	_ = cmd.PersistentFlags().String(virtualNetworkIdFlagName, virtualNetworkIdFlagDefault, virtualNetworkIdDescription)
	cmd.MarkPersistentFlagRequired(virtualNetworkIdFlagName)

	return nil
}

func retrieveOperationVirtualNetworksUpdateVirtualNetworkAPIVersionFlag(m *virtual_networks.UpdateVirtualNetworkParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationVirtualNetworksUpdateVirtualNetworkBodyFlag(m *virtual_networks.UpdateVirtualNetworkParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := virtual_networks.UpdateVirtualNetworkBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in UpdateVirtualNetworkBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = virtual_networks.UpdateVirtualNetworkBody{}
	}
	err, added := retrieveModelUpdateVirtualNetworkBodyFlags(0, &bodyValueModel, "", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Body = bodyValueModel
	}
	if dryRun && debug {

		bodyValueDebugBytes, err := json.Marshal(m.Body)
		if err != nil {
			return err, false
		}
		logDebugf("Body dry-run payload: %v", string(bodyValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationVirtualNetworksUpdateVirtualNetworkVirtualNetworkIDFlag(m *virtual_networks.UpdateVirtualNetworkParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("virtual_network_id") {

		var virtualNetworkIdFlagName string
		if cmdPrefix == "" {
			virtualNetworkIdFlagName = "virtual_network_id"
		} else {
			virtualNetworkIdFlagName = fmt.Sprintf("%v.virtual_network_id", cmdPrefix)
		}

		virtualNetworkIdFlagValue, err := cmd.Flags().GetString(virtualNetworkIdFlagName)
		if err != nil {
			return err, false
		}
		m.VirtualNetworkID = virtualNetworkIdFlagValue

	}
	return nil, retAdded
}

// parseOperationVirtualNetworksUpdateVirtualNetworkResult parses request result and return the string content
func parseOperationVirtualNetworksUpdateVirtualNetworkResult(resp0 *virtual_networks.UpdateVirtualNetworkOK) (string, error) {
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
func registerModelUpdateVirtualNetworkBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUpdateVirtualNetworkBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUpdateVirtualNetworkBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName = ""

	if err := registerModelUpdateVirtualNetworkParamsBodyDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUpdateVirtualNetworkBodyFlags(depth int, m *virtual_networks.UpdateVirtualNetworkBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveUpdateVirtualNetworkBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveUpdateVirtualNetworkBodyDataFlags(depth int, m *virtual_networks.UpdateVirtualNetworkBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%vdata", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data UpdateVirtualNetworkParamsBodyData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &virtual_networks.UpdateVirtualNetworkParamsBodyData{}
	}

	dataFlagName = ""
	err, dataAdded := retrieveModelUpdateVirtualNetworkParamsBodyDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}

// register flags to command
func registerModelUpdateVirtualNetworkParamsBodyDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUpdateVirtualNetworkParamsBodyDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateVirtualNetworkParamsBodyDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUpdateVirtualNetworkParamsBodyDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName = ""

	if err := registerModelUpdateVirtualNetworkParamsBodyDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerUpdateVirtualNetworkParamsBodyDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["virtual_network"]. Required. `

	var typeFlagName = "type"

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)
	cmd.MarkPersistentFlagRequired(typeFlagName)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["virtual_network"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUpdateVirtualNetworkParamsBodyDataFlags(depth int, m *virtual_networks.UpdateVirtualNetworkParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveUpdateVirtualNetworkParamsBodyDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, typeAdded := retrieveUpdateVirtualNetworkParamsBodyDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveUpdateVirtualNetworkParamsBodyDataAttributesFlags(depth int, m *virtual_networks.UpdateVirtualNetworkParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%vattributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes UpdateVirtualNetworkParamsBodyDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &virtual_networks.UpdateVirtualNetworkParamsBodyDataAttributes{}
	}

	attributesFlagName = ""
	err, attributesAdded := retrieveModelUpdateVirtualNetworkParamsBodyDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveUpdateVirtualNetworkParamsBodyDataTypeFlags(depth int, m *virtual_networks.UpdateVirtualNetworkParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	var typeFlagName = "type"
	if cmd.Flags().Changed(typeFlagName) {


		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = &typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// register flags to command
func registerModelUpdateVirtualNetworkParamsBodyDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUpdateVirtualNetworkParamsBodyDataAttributesDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUpdateVirtualNetworkParamsBodyDataAttributesDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := ``

	var descriptionFlagName = "description"

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUpdateVirtualNetworkParamsBodyDataAttributesFlags(depth int, m *virtual_networks.UpdateVirtualNetworkParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, descriptionAdded := retrieveUpdateVirtualNetworkParamsBodyDataAttributesDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	return nil, retAdded
}

func retrieveUpdateVirtualNetworkParamsBodyDataAttributesDescriptionFlags(depth int, m *virtual_networks.UpdateVirtualNetworkParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	var descriptionFlagName = "description"
	if cmd.Flags().Changed(descriptionFlagName) {

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}
