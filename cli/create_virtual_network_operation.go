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

// makeOperationVirtualNetworksCreateVirtualNetworkCmd returns a cmd to handle operation createVirtualNetwork
func makeOperationVirtualNetworksCreateVirtualNetworkCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "create-virtual-network",
		Short: `Creates a new Virtual Network.
`,
		RunE: runOperationVirtualNetworksCreateVirtualNetwork,
	}

	if err := registerOperationVirtualNetworksCreateVirtualNetworkParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVirtualNetworksCreateVirtualNetwork uses cmd flags to call endpoint api
func runOperationVirtualNetworksCreateVirtualNetwork(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := virtual_networks.NewCreateVirtualNetworkParams()
	if err, _ := retrieveOperationVirtualNetworksCreateVirtualNetworkAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationVirtualNetworksCreateVirtualNetworkBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationVirtualNetworksCreateVirtualNetworkResult(appCli.VirtualNetworks.CreateVirtualNetwork(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		utils.PrintOutput(msgStr)
	}
	return nil
}

// registerOperationVirtualNetworksCreateVirtualNetworkParamFlags registers all flags needed to fill params
func registerOperationVirtualNetworksCreateVirtualNetworkParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVirtualNetworksCreateVirtualNetworkAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationVirtualNetworksCreateVirtualNetworkBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVirtualNetworksCreateVirtualNetworkAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationVirtualNetworksCreateVirtualNetworkBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelCreateVirtualNetworkBodyFlags(0, "createVirtualNetworkBody", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationVirtualNetworksCreateVirtualNetworkAPIVersionFlag(m *virtual_networks.CreateVirtualNetworkParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationVirtualNetworksCreateVirtualNetworkBodyFlag(m *virtual_networks.CreateVirtualNetworkParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := virtual_networks.CreateVirtualNetworkBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in CreateVirtualNetworkBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = virtual_networks.CreateVirtualNetworkBody{}
	}
	err, added := retrieveModelCreateVirtualNetworkBodyFlags(0, &bodyValueModel, "createVirtualNetworkBody", cmd)
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

// parseOperationVirtualNetworksCreateVirtualNetworkResult parses request result and return the string content
func parseOperationVirtualNetworksCreateVirtualNetworkResult(resp0 *virtual_networks.CreateVirtualNetworkCreated, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*virtual_networks.CreateVirtualNetworkCreated)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*virtual_networks.CreateVirtualNetworkUnprocessableEntity)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		parsedErrorResponse, err := api.ParseErrorResponse(respErr)

		if err != nil {
			return "", err
		}

		return parsedErrorResponse, nil
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
func registerModelCreateVirtualNetworkBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateVirtualNetworkBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateVirtualNetworkBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelCreateVirtualNetworkParamsBodyDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateVirtualNetworkBodyFlags(depth int, m *virtual_networks.CreateVirtualNetworkBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveCreateVirtualNetworkBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveCreateVirtualNetworkBodyDataFlags(depth int, m *virtual_networks.CreateVirtualNetworkBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data CreateVirtualNetworkParamsBodyData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &virtual_networks.CreateVirtualNetworkParamsBodyData{}
	}

	err, dataAdded := retrieveModelCreateVirtualNetworkParamsBodyDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
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
func registerModelCreateVirtualNetworkParamsBodyDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateVirtualNetworkParamsBodyDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateVirtualNetworkParamsBodyDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateVirtualNetworkParamsBodyDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelCreateVirtualNetworkParamsBodyDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateVirtualNetworkParamsBodyDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["virtual_network"]. Required. `

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

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
func retrieveModelCreateVirtualNetworkParamsBodyDataFlags(depth int, m *virtual_networks.CreateVirtualNetworkParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveCreateVirtualNetworkParamsBodyDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, typeAdded := retrieveCreateVirtualNetworkParamsBodyDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveCreateVirtualNetworkParamsBodyDataAttributesFlags(depth int, m *virtual_networks.CreateVirtualNetworkParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes CreateVirtualNetworkParamsBodyDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &virtual_networks.CreateVirtualNetworkParamsBodyDataAttributes{}
	}

	err, attributesAdded := retrieveModelCreateVirtualNetworkParamsBodyDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveCreateVirtualNetworkParamsBodyDataTypeFlags(depth int, m *virtual_networks.CreateVirtualNetworkParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

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
func registerModelCreateVirtualNetworkParamsBodyDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateVirtualNetworkParamsBodyDataAttributesDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateVirtualNetworkParamsBodyDataAttributesProject(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateVirtualNetworkParamsBodyDataAttributesSite(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateVirtualNetworkParamsBodyDataAttributesDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := `Required. `

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)
	cmd.MarkPersistentFlagRequired(descriptionFlagName)

	return nil
}

func registerCreateVirtualNetworkParamsBodyDataAttributesProject(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	projectDescription := `Required. Project ID or slug`

	var projectFlagName string
	if cmdPrefix == "" {
		projectFlagName = "project"
	} else {
		projectFlagName = fmt.Sprintf("%v.project", cmdPrefix)
	}

	var projectFlagDefault string

	_ = cmd.PersistentFlags().String(projectFlagName, projectFlagDefault, projectDescription)
	cmd.MarkPersistentFlagRequired(projectFlagName)

	return nil
}

func registerCreateVirtualNetworkParamsBodyDataAttributesSite(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	siteDescription := `Enum: ["ASH","BGT","BUE","CHI","DAL","FRA","LAX","LON","MEX","MEX2","MIA","MIA2","NYC","SAN","SAN2","SAO","SAO2","SYD","TYO","TYO2"]. Site ID or slug`

	var siteFlagName string
	if cmdPrefix == "" {
		siteFlagName = "site"
	} else {
		siteFlagName = fmt.Sprintf("%v.site", cmdPrefix)
	}

	var siteFlagDefault string

	_ = cmd.PersistentFlags().String(siteFlagName, siteFlagDefault, siteDescription)
	cmd.MarkPersistentFlagRequired(siteFlagName)

	if err := cmd.RegisterFlagCompletionFunc(siteFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["ASH","BGT","BUE","CHI","DAL","FRA","LAX","LON","MEX","MEX2","MIA","MIA2","NYC","SAN","SAN2","SAO","SAO2","SYD","TYO","TYO2"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateVirtualNetworkParamsBodyDataAttributesFlags(depth int, m *virtual_networks.CreateVirtualNetworkParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, descriptionAdded := retrieveCreateVirtualNetworkParamsBodyDataAttributesDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, projectAdded := retrieveCreateVirtualNetworkParamsBodyDataAttributesProjectFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || projectAdded

	err, siteAdded := retrieveCreateVirtualNetworkParamsBodyDataAttributesSiteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || siteAdded

	return nil, retAdded
}

func retrieveCreateVirtualNetworkParamsBodyDataAttributesDescriptionFlags(depth int, m *virtual_networks.CreateVirtualNetworkParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = &descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateVirtualNetworkParamsBodyDataAttributesProjectFlags(depth int, m *virtual_networks.CreateVirtualNetworkParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	projectFlagName := fmt.Sprintf("%v.project", cmdPrefix)
	if cmd.Flags().Changed(projectFlagName) {

		var projectFlagName string
		if cmdPrefix == "" {
			projectFlagName = "project"
		} else {
			projectFlagName = fmt.Sprintf("%v.project", cmdPrefix)
		}

		projectFlagValue, err := cmd.Flags().GetString(projectFlagName)
		if err != nil {
			return err, false
		}
		m.Project = &projectFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateVirtualNetworkParamsBodyDataAttributesSiteFlags(depth int, m *virtual_networks.CreateVirtualNetworkParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	siteFlagName := fmt.Sprintf("%v.site", cmdPrefix)
	if cmd.Flags().Changed(siteFlagName) {

		var siteFlagName string
		if cmdPrefix == "" {
			siteFlagName = "site"
		} else {
			siteFlagName = fmt.Sprintf("%v.site", cmdPrefix)
		}

		siteFlagValue, err := cmd.Flags().GetString(siteFlagName)
		if err != nil {
			return err, false
		}
		m.Site = siteFlagValue

		retAdded = true
	}

	return nil, retAdded
}
