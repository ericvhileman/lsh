// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/out_of_band"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationOutOfBandCreateServerOutOfBandCmd returns a cmd to handle operation createServerOutOfBand
func makeOperationOutOfBandCreateServerOutOfBandCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "create-server-out-of-band",
		Short: ``,
		RunE:  runOperationOutOfBandCreateServerOutOfBand,
	}

	if err := registerOperationOutOfBandCreateServerOutOfBandParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationOutOfBandCreateServerOutOfBand uses cmd flags to call endpoint api
func runOperationOutOfBandCreateServerOutOfBand(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := out_of_band.NewCreateServerOutOfBandParams()
	if err, _ := retrieveOperationOutOfBandCreateServerOutOfBandAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationOutOfBandCreateServerOutOfBandBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationOutOfBandCreateServerOutOfBandServerIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationOutOfBandCreateServerOutOfBandResult(appCli.OutOfBand.CreateServerOutOfBand(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationOutOfBandCreateServerOutOfBandParamFlags registers all flags needed to fill params
func registerOperationOutOfBandCreateServerOutOfBandParamFlags(cmd *cobra.Command) error {
	if err := registerOperationOutOfBandCreateServerOutOfBandAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationOutOfBandCreateServerOutOfBandBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationOutOfBandCreateServerOutOfBandServerIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationOutOfBandCreateServerOutOfBandAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationOutOfBandCreateServerOutOfBandBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelCreateServerOutOfBandBodyFlags(0, "createServerOutOfBandBody", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationOutOfBandCreateServerOutOfBandServerIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	serverIdDescription := `Required. `

	var serverIdFlagName string
	if cmdPrefix == "" {
		serverIdFlagName = "server_id"
	} else {
		serverIdFlagName = fmt.Sprintf("%v.server_id", cmdPrefix)
	}

	var serverIdFlagDefault string

	_ = cmd.PersistentFlags().String(serverIdFlagName, serverIdFlagDefault, serverIdDescription)

	return nil
}

func retrieveOperationOutOfBandCreateServerOutOfBandAPIVersionFlag(m *out_of_band.CreateServerOutOfBandParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationOutOfBandCreateServerOutOfBandBodyFlag(m *out_of_band.CreateServerOutOfBandParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := out_of_band.CreateServerOutOfBandBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in CreateServerOutOfBandBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = out_of_band.CreateServerOutOfBandBody{}
	}
	err, added := retrieveModelCreateServerOutOfBandBodyFlags(0, &bodyValueModel, "createServerOutOfBandBody", cmd)
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
func retrieveOperationOutOfBandCreateServerOutOfBandServerIDFlag(m *out_of_band.CreateServerOutOfBandParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("server_id") {

		var serverIdFlagName string
		if cmdPrefix == "" {
			serverIdFlagName = "server_id"
		} else {
			serverIdFlagName = fmt.Sprintf("%v.server_id", cmdPrefix)
		}

		serverIdFlagValue, err := cmd.Flags().GetString(serverIdFlagName)
		if err != nil {
			return err, false
		}
		m.ServerID = serverIdFlagValue

	}
	return nil, retAdded
}

// parseOperationOutOfBandCreateServerOutOfBandResult parses request result and return the string content
func parseOperationOutOfBandCreateServerOutOfBandResult(resp0 *out_of_band.CreateServerOutOfBandCreated, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*out_of_band.CreateServerOutOfBandCreated)
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
		resp1, ok := iResp1.(*out_of_band.CreateServerOutOfBandForbidden)
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
		resp2, ok := iResp2.(*out_of_band.CreateServerOutOfBandNotFound)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning createServerOutOfBandNotAcceptable is not supported

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
func registerModelCreateServerOutOfBandBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateServerOutOfBandBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateServerOutOfBandBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelCreateServerOutOfBandParamsBodyDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateServerOutOfBandBodyFlags(depth int, m *out_of_band.CreateServerOutOfBandBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveCreateServerOutOfBandBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveCreateServerOutOfBandBodyDataFlags(depth int, m *out_of_band.CreateServerOutOfBandBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data CreateServerOutOfBandParamsBodyData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &out_of_band.CreateServerOutOfBandParamsBodyData{}
	}

	err, dataAdded := retrieveModelCreateServerOutOfBandParamsBodyDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
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
func registerModelCreateServerOutOfBandParamsBodyDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateServerOutOfBandParamsBodyDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateServerOutOfBandParamsBodyDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateServerOutOfBandParamsBodyDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelCreateServerOutOfBandParamsBodyDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateServerOutOfBandParamsBodyDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["out_of_band"]. Required. `

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
			if err := json.Unmarshal([]byte(`["out_of_band"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateServerOutOfBandParamsBodyDataFlags(depth int, m *out_of_band.CreateServerOutOfBandParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveCreateServerOutOfBandParamsBodyDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, typeAdded := retrieveCreateServerOutOfBandParamsBodyDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveCreateServerOutOfBandParamsBodyDataAttributesFlags(depth int, m *out_of_band.CreateServerOutOfBandParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes CreateServerOutOfBandParamsBodyDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &out_of_band.CreateServerOutOfBandParamsBodyDataAttributes{}
	}

	err, attributesAdded := retrieveModelCreateServerOutOfBandParamsBodyDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveCreateServerOutOfBandParamsBodyDataTypeFlags(depth int, m *out_of_band.CreateServerOutOfBandParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func registerModelCreateServerOutOfBandParamsBodyDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateServerOutOfBandParamsBodyDataAttributesSSHKeyID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateServerOutOfBandParamsBodyDataAttributesSSHKeyID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sshKeyIdDescription := `SSH Key ID to set for out of band`

	var sshKeyIdFlagName string
	if cmdPrefix == "" {
		sshKeyIdFlagName = "ssh_key_id"
	} else {
		sshKeyIdFlagName = fmt.Sprintf("%v.ssh_key_id", cmdPrefix)
	}

	var sshKeyIdFlagDefault string

	_ = cmd.PersistentFlags().String(sshKeyIdFlagName, sshKeyIdFlagDefault, sshKeyIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateServerOutOfBandParamsBodyDataAttributesFlags(depth int, m *out_of_band.CreateServerOutOfBandParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, sshKeyIdAdded := retrieveCreateServerOutOfBandParamsBodyDataAttributesSSHKeyIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sshKeyIdAdded

	return nil, retAdded
}

func retrieveCreateServerOutOfBandParamsBodyDataAttributesSSHKeyIDFlags(depth int, m *out_of_band.CreateServerOutOfBandParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sshKeyIdFlagName := fmt.Sprintf("%v.ssh_key_id", cmdPrefix)
	if cmd.Flags().Changed(sshKeyIdFlagName) {

		var sshKeyIdFlagName string
		if cmdPrefix == "" {
			sshKeyIdFlagName = "ssh_key_id"
		} else {
			sshKeyIdFlagName = fmt.Sprintf("%v.ssh_key_id", cmdPrefix)
		}

		sshKeyIdFlagValue, err := cmd.Flags().GetString(sshKeyIdFlagName)
		if err != nil {
			return err, false
		}
		m.SSHKeyID = sshKeyIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
