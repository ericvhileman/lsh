// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/operating_systems"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationOperatingSystemsGetPlansOperatingSystemCmd returns a cmd to handle operation getPlansOperatingSystem
func makeOperationOperatingSystemsGetPlansOperatingSystemCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "get-plans-operating-system",
		Short: `Lists all operating systems available to deploy and reinstall.
`,
		RunE: runOperationOperatingSystemsGetPlansOperatingSystem,
	}

	if err := registerOperationOperatingSystemsGetPlansOperatingSystemParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationOperatingSystemsGetPlansOperatingSystem uses cmd flags to call endpoint api
func runOperationOperatingSystemsGetPlansOperatingSystem(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := operating_systems.NewGetPlansOperatingSystemParams()
	if err, _ := retrieveOperationOperatingSystemsGetPlansOperatingSystemAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationOperatingSystemsGetPlansOperatingSystemResult(appCli.OperatingSystems.GetPlansOperatingSystem(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(utils.PrettifyJson(msgStr))
	}
	return nil
}

// registerOperationOperatingSystemsGetPlansOperatingSystemParamFlags registers all flags needed to fill params
func registerOperationOperatingSystemsGetPlansOperatingSystemParamFlags(cmd *cobra.Command) error {
	if err := registerOperationOperatingSystemsGetPlansOperatingSystemAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationOperatingSystemsGetPlansOperatingSystemAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationOperatingSystemsGetPlansOperatingSystemAPIVersionFlag(m *operating_systems.GetPlansOperatingSystemParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationOperatingSystemsGetPlansOperatingSystemResult parses request result and return the string content
func parseOperationOperatingSystemsGetPlansOperatingSystemResult(resp0 *operating_systems.GetPlansOperatingSystemOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*operating_systems.GetPlansOperatingSystemOK)
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
func registerModelGetPlansOperatingSystemOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetPlansOperatingSystemOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetPlansOperatingSystemOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data []*models.OperatingSystems array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetPlansOperatingSystemOKBodyFlags(depth int, m *operating_systems.GetPlansOperatingSystemOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveGetPlansOperatingSystemOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveGetPlansOperatingSystemOKBodyDataFlags(depth int, m *operating_systems.GetPlansOperatingSystemOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// warning: data array type []*models.OperatingSystems is not supported by go-swagger cli yet
	}

	return nil, retAdded
}