// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/lsh/client/virtual_networks"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationVirtualNetworksUpdateVirtualNetworkCmd returns a cmd to handle operation updateVirtualNetwork
func makeOperationVirtualNetworksUpdateVirtualNetworkCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "update",
		Short: `Update a Virtual Network.`,
		RunE:  runOperationVirtualNetworksUpdateVirtualNetwork,
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

	response, err := appCli.VirtualNetworks.UpdateVirtualNetwork(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetPayload())
	}
	return nil
}

// registerOperationVirtualNetworksUpdateVirtualNetworkParamFlags registers all flags needed to fill params
func registerOperationVirtualNetworksUpdateVirtualNetworkParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVirtualNetworksUpdateVirtualNetworkBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationVirtualNetworksUpdateVirtualNetworkVirtualNetworkIDParamFlags("", cmd); err != nil {
		return err
	}
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
		virtualNetworkIdFlagName = "id"
	} else {
		virtualNetworkIdFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var virtualNetworkIdFlagDefault string

	_ = cmd.PersistentFlags().String(virtualNetworkIdFlagName, virtualNetworkIdFlagDefault, virtualNetworkIdDescription)
	cmd.MarkPersistentFlagRequired(virtualNetworkIdFlagName)

	return nil
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
	if cmd.Flags().Changed("id") {

		var virtualNetworkIdFlagName string
		if cmdPrefix == "" {
			virtualNetworkIdFlagName = "id"
		} else {
			virtualNetworkIdFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		virtualNetworkIdFlagValue, err := cmd.Flags().GetString(virtualNetworkIdFlagName)
		if err != nil {
			return err, false
		}
		m.VirtualNetworkID = virtualNetworkIdFlagValue
		m.Body.Data.ID = &m.VirtualNetworkID

	}
	return nil, retAdded
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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUpdateVirtualNetworkParamsBodyDataFlags(depth int, m *virtual_networks.UpdateVirtualNetworkParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveUpdateVirtualNetworkParamsBodyDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

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
