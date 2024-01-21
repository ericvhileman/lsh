// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/virtual_network_assignments"
	"github.com/latitudesh/cli/internal/api"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkCmd returns a cmd to handle operation assignServerVirtualNetwork
func makeOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "assign-server-virtual-network",
		Short: ``,
		RunE:  runOperationVirtualNetworkAssignmentsAssignServerVirtualNetwork,
	}

	if err := registerOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVirtualNetworkAssignmentsAssignServerVirtualNetwork uses cmd flags to call endpoint api
func runOperationVirtualNetworkAssignmentsAssignServerVirtualNetwork(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := virtual_network_assignments.NewAssignServerVirtualNetworkParams()
	if err, _ := retrieveOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkResult(appCli.VirtualNetworkAssignments.AssignServerVirtualNetwork(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		utils.PrintOutput(msgStr)
	}
	return nil
}

// registerOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkParamFlags registers all flags needed to fill params
func registerOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelAssignServerVirtualNetworkBodyFlags(0, "assignServerVirtualNetworkBody", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkAPIVersionFlag(m *virtual_network_assignments.AssignServerVirtualNetworkParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkBodyFlag(m *virtual_network_assignments.AssignServerVirtualNetworkParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := virtual_network_assignments.AssignServerVirtualNetworkBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in AssignServerVirtualNetworkBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = virtual_network_assignments.AssignServerVirtualNetworkBody{}
	}
	err, added := retrieveModelAssignServerVirtualNetworkBodyFlags(0, &bodyValueModel, "assignServerVirtualNetworkBody", cmd)
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

// parseOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkResult parses request result and return the string content
func parseOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkResult(resp0 *virtual_network_assignments.AssignServerVirtualNetworkCreated, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*virtual_network_assignments.AssignServerVirtualNetworkCreated)
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
		resp1, ok := iResp1.(*virtual_network_assignments.AssignServerVirtualNetworkForbidden)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*virtual_network_assignments.AssignServerVirtualNetworkUnprocessableEntity)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		notFoundErrorMessage, err := api.ParseErrorResponse(respErr)

		if err != nil {
			return "", err
		}

		if len(notFoundErrorMessage) > 0 {
			return notFoundErrorMessage, nil
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
func registerModelAssignServerVirtualNetworkBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAssignServerVirtualNetworkBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAssignServerVirtualNetworkBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelAssignServerVirtualNetworkParamsBodyDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAssignServerVirtualNetworkBodyFlags(depth int, m *virtual_network_assignments.AssignServerVirtualNetworkBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveAssignServerVirtualNetworkBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveAssignServerVirtualNetworkBodyDataFlags(depth int, m *virtual_network_assignments.AssignServerVirtualNetworkBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data AssignServerVirtualNetworkParamsBodyData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &virtual_network_assignments.AssignServerVirtualNetworkParamsBodyData{}
	}

	err, dataAdded := retrieveModelAssignServerVirtualNetworkParamsBodyDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
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
func registerModelAssignServerVirtualNetworkParamsBodyDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAssignServerVirtualNetworkParamsBodyDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAssignServerVirtualNetworkParamsBodyDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAssignServerVirtualNetworkParamsBodyDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelAssignServerVirtualNetworkParamsBodyDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerAssignServerVirtualNetworkParamsBodyDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["virtual_network_assignment"]. Required. `

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
			if err := json.Unmarshal([]byte(`["virtual_network_assignment"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAssignServerVirtualNetworkParamsBodyDataFlags(depth int, m *virtual_network_assignments.AssignServerVirtualNetworkParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveAssignServerVirtualNetworkParamsBodyDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, typeAdded := retrieveAssignServerVirtualNetworkParamsBodyDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveAssignServerVirtualNetworkParamsBodyDataAttributesFlags(depth int, m *virtual_network_assignments.AssignServerVirtualNetworkParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes AssignServerVirtualNetworkParamsBodyDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &virtual_network_assignments.AssignServerVirtualNetworkParamsBodyDataAttributes{}
	}

	err, attributesAdded := retrieveModelAssignServerVirtualNetworkParamsBodyDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveAssignServerVirtualNetworkParamsBodyDataTypeFlags(depth int, m *virtual_network_assignments.AssignServerVirtualNetworkParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func registerModelAssignServerVirtualNetworkParamsBodyDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAssignServerVirtualNetworkParamsBodyDataAttributesServerID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAssignServerVirtualNetworkParamsBodyDataAttributesVirtualNetworkID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAssignServerVirtualNetworkParamsBodyDataAttributesServerID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	serverIdDescription := `Required. `

	var serverIdFlagName string
	if cmdPrefix == "" {
		serverIdFlagName = "server_id"
	} else {
		serverIdFlagName = fmt.Sprintf("%v.server_id", cmdPrefix)
	}

	var serverIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(serverIdFlagName, serverIdFlagDefault, serverIdDescription)
	cmd.MarkPersistentFlagRequired(serverIdFlagName)

	return nil
}

func registerAssignServerVirtualNetworkParamsBodyDataAttributesVirtualNetworkID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	virtualNetworkIdDescription := `Required. `

	var virtualNetworkIdFlagName string
	if cmdPrefix == "" {
		virtualNetworkIdFlagName = "virtual_network_id"
	} else {
		virtualNetworkIdFlagName = fmt.Sprintf("%v.virtual_network_id", cmdPrefix)
	}

	var virtualNetworkIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(virtualNetworkIdFlagName, virtualNetworkIdFlagDefault, virtualNetworkIdDescription)
	cmd.MarkPersistentFlagRequired(virtualNetworkIdFlagName)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAssignServerVirtualNetworkParamsBodyDataAttributesFlags(depth int, m *virtual_network_assignments.AssignServerVirtualNetworkParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, serverIdAdded := retrieveAssignServerVirtualNetworkParamsBodyDataAttributesServerIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || serverIdAdded

	err, virtualNetworkIdAdded := retrieveAssignServerVirtualNetworkParamsBodyDataAttributesVirtualNetworkIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || virtualNetworkIdAdded

	return nil, retAdded
}

func retrieveAssignServerVirtualNetworkParamsBodyDataAttributesServerIDFlags(depth int, m *virtual_network_assignments.AssignServerVirtualNetworkParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	serverIdFlagName := fmt.Sprintf("%v.server_id", cmdPrefix)
	if cmd.Flags().Changed(serverIdFlagName) {

		var serverIdFlagName string
		if cmdPrefix == "" {
			serverIdFlagName = "server_id"
		} else {
			serverIdFlagName = fmt.Sprintf("%v.server_id", cmdPrefix)
		}

		serverIdFlagValue, err := cmd.Flags().GetInt64(serverIdFlagName)
		if err != nil {
			return err, false
		}
		m.ServerID = &serverIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAssignServerVirtualNetworkParamsBodyDataAttributesVirtualNetworkIDFlags(depth int, m *virtual_network_assignments.AssignServerVirtualNetworkParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	virtualNetworkIdFlagName := fmt.Sprintf("%v.virtual_network_id", cmdPrefix)
	if cmd.Flags().Changed(virtualNetworkIdFlagName) {

		var virtualNetworkIdFlagName string
		if cmdPrefix == "" {
			virtualNetworkIdFlagName = "virtual_network_id"
		} else {
			virtualNetworkIdFlagName = fmt.Sprintf("%v.virtual_network_id", cmdPrefix)
		}

		virtualNetworkIdFlagValue, err := cmd.Flags().GetInt64(virtualNetworkIdFlagName)
		if err != nil {
			return err, false
		}
		m.VirtualNetworkID = &virtualNetworkIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
