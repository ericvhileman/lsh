// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/v_p_n_sessions"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationVpnSessionsDeleteVpnSessionCmd returns a cmd to handle operation deleteVpnSession
func makeOperationVpnSessionsDeleteVpnSessionCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "delete-vpn-session",
		Short: `Deletes an existing VPN Session.
`,
		RunE: runOperationVpnSessionsDeleteVpnSession,
	}

	if err := registerOperationVpnSessionsDeleteVpnSessionParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVpnSessionsDeleteVpnSession uses cmd flags to call endpoint api
func runOperationVpnSessionsDeleteVpnSession(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := v_p_n_sessions.NewDeleteVpnSessionParams()
	if err, _ := retrieveOperationVpnSessionsDeleteVpnSessionAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationVpnSessionsDeleteVpnSessionVpnSessionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationVpnSessionsDeleteVpnSessionResult(appCli.VpnSessions.DeleteVpnSession(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationVpnSessionsDeleteVpnSessionParamFlags registers all flags needed to fill params
func registerOperationVpnSessionsDeleteVpnSessionParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVpnSessionsDeleteVpnSessionAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationVpnSessionsDeleteVpnSessionVpnSessionIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVpnSessionsDeleteVpnSessionAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationVpnSessionsDeleteVpnSessionVpnSessionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	vpnSessionIdDescription := `Required. `

	var vpnSessionIdFlagName string
	if cmdPrefix == "" {
		vpnSessionIdFlagName = "vpn_session_id"
	} else {
		vpnSessionIdFlagName = fmt.Sprintf("%v.vpn_session_id", cmdPrefix)
	}

	var vpnSessionIdFlagDefault string

	_ = cmd.PersistentFlags().String(vpnSessionIdFlagName, vpnSessionIdFlagDefault, vpnSessionIdDescription)

	return nil
}

func retrieveOperationVpnSessionsDeleteVpnSessionAPIVersionFlag(m *v_p_n_sessions.DeleteVpnSessionParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationVpnSessionsDeleteVpnSessionVpnSessionIDFlag(m *v_p_n_sessions.DeleteVpnSessionParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("vpn_session_id") {

		var vpnSessionIdFlagName string
		if cmdPrefix == "" {
			vpnSessionIdFlagName = "vpn_session_id"
		} else {
			vpnSessionIdFlagName = fmt.Sprintf("%v.vpn_session_id", cmdPrefix)
		}

		vpnSessionIdFlagValue, err := cmd.Flags().GetString(vpnSessionIdFlagName)
		if err != nil {
			return err, false
		}
		m.VpnSessionID = vpnSessionIdFlagValue

	}
	return nil, retAdded
}

// parseOperationVpnSessionsDeleteVpnSessionResult parses request result and return the string content
func parseOperationVpnSessionsDeleteVpnSessionResult(resp0 *v_p_n_sessions.DeleteVpnSessionNoContent, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning deleteVpnSessionNoContent is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*v_p_n_sessions.DeleteVpnSessionNotFound)
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

	// warning: non schema response deleteVpnSessionNoContent is not supported by go-swagger cli yet.

	return "", nil
}
