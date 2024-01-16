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

// makeOperationOutOfBandGetServerOutOfBandCmd returns a cmd to handle operation getServerOutOfBand
func makeOperationOutOfBandGetServerOutOfBandCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "get-server-out-of-band",
		Short: ``,
		RunE:  runOperationOutOfBandGetServerOutOfBand,
	}

	if err := registerOperationOutOfBandGetServerOutOfBandParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationOutOfBandGetServerOutOfBand uses cmd flags to call endpoint api
func runOperationOutOfBandGetServerOutOfBand(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := out_of_band.NewGetServerOutOfBandParams()
	if err, _ := retrieveOperationOutOfBandGetServerOutOfBandAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationOutOfBandGetServerOutOfBandServerIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationOutOfBandGetServerOutOfBandResult(appCli.OutOfBand.GetServerOutOfBand(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationOutOfBandGetServerOutOfBandParamFlags registers all flags needed to fill params
func registerOperationOutOfBandGetServerOutOfBandParamFlags(cmd *cobra.Command) error {
	if err := registerOperationOutOfBandGetServerOutOfBandAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationOutOfBandGetServerOutOfBandServerIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationOutOfBandGetServerOutOfBandAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationOutOfBandGetServerOutOfBandServerIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	serverIdDescription := `Required. The Server ID`

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

func retrieveOperationOutOfBandGetServerOutOfBandAPIVersionFlag(m *out_of_band.GetServerOutOfBandParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationOutOfBandGetServerOutOfBandServerIDFlag(m *out_of_band.GetServerOutOfBandParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationOutOfBandGetServerOutOfBandResult parses request result and return the string content
func parseOperationOutOfBandGetServerOutOfBandResult(resp0 *out_of_band.GetServerOutOfBandOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*out_of_band.GetServerOutOfBandOK)
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
		resp1, ok := iResp1.(*out_of_band.GetServerOutOfBandNotFound)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
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
