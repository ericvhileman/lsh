// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/v_p_n_sessions"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationVpnSessionsGetVpnSessionsCmd returns a cmd to handle operation getVpnSessions
func makeOperationVpnSessionsGetVpnSessionsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "get-vpn-sessions",
		Short: ``,
		RunE:  runOperationVpnSessionsGetVpnSessions,
	}

	if err := registerOperationVpnSessionsGetVpnSessionsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVpnSessionsGetVpnSessions uses cmd flags to call endpoint api
func runOperationVpnSessionsGetVpnSessions(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := v_p_n_sessions.NewGetVpnSessionsParams()
	if err, _ := retrieveOperationVpnSessionsGetVpnSessionsAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationVpnSessionsGetVpnSessionsFilterLocationFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationVpnSessionsGetVpnSessionsResult(appCli.VpnSessions.GetVpnSessions(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(utils.PrettifyJson(msgStr))
	}
	return nil
}

// registerOperationVpnSessionsGetVpnSessionsParamFlags registers all flags needed to fill params
func registerOperationVpnSessionsGetVpnSessionsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVpnSessionsGetVpnSessionsAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationVpnSessionsGetVpnSessionsFilterLocationParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVpnSessionsGetVpnSessionsAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationVpnSessionsGetVpnSessionsFilterLocationParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterLocationDescription := `Enum: ["ASH","BGT","BUE","CHI","DAL","FRA","LAX","LON","MEX","MEX2","MIA","MIA2","NYC","SAN","SAN2","SAO","SAO2","SYD","TYO","TYO2"]. `

	var filterLocationFlagName string
	if cmdPrefix == "" {
		filterLocationFlagName = "filter[location]"
	} else {
		filterLocationFlagName = fmt.Sprintf("%v.filter[location]", cmdPrefix)
	}

	var filterLocationFlagDefault string

	_ = cmd.PersistentFlags().String(filterLocationFlagName, filterLocationFlagDefault, filterLocationDescription)

	if err := cmd.RegisterFlagCompletionFunc(filterLocationFlagName,
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

func retrieveOperationVpnSessionsGetVpnSessionsAPIVersionFlag(m *v_p_n_sessions.GetVpnSessionsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationVpnSessionsGetVpnSessionsFilterLocationFlag(m *v_p_n_sessions.GetVpnSessionsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[location]") {

		var filterLocationFlagName string
		if cmdPrefix == "" {
			filterLocationFlagName = "filter[location]"
		} else {
			filterLocationFlagName = fmt.Sprintf("%v.filter[location]", cmdPrefix)
		}

		filterLocationFlagValue, err := cmd.Flags().GetString(filterLocationFlagName)
		if err != nil {
			return err, false
		}
		m.FilterLocation = &filterLocationFlagValue

	}
	return nil, retAdded
}

// parseOperationVpnSessionsGetVpnSessionsResult parses request result and return the string content
func parseOperationVpnSessionsGetVpnSessionsResult(resp0 *v_p_n_sessions.GetVpnSessionsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*v_p_n_sessions.GetVpnSessionsOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning getVpnSessionsUnprocessableEntity is not supported

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
func registerModelGetVpnSessionsOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetVpnSessionsOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGetVpnSessionsOKBodyMeta(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetVpnSessionsOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data []*models.VpnSessionDataWithPassword array type is not supported by go-swagger cli yet

	return nil
}

func registerGetVpnSessionsOKBodyMeta(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: meta interface{} map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetVpnSessionsOKBodyFlags(depth int, m *v_p_n_sessions.GetVpnSessionsOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveGetVpnSessionsOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	err, metaAdded := retrieveGetVpnSessionsOKBodyMetaFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || metaAdded

	return nil, retAdded
}

func retrieveGetVpnSessionsOKBodyDataFlags(depth int, m *v_p_n_sessions.GetVpnSessionsOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// warning: data array type []*models.VpnSessionDataWithPassword is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveGetVpnSessionsOKBodyMetaFlags(depth int, m *v_p_n_sessions.GetVpnSessionsOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	metaFlagName := fmt.Sprintf("%v.meta", cmdPrefix)
	if cmd.Flags().Changed(metaFlagName) {
		// warning: meta map type interface{} is not supported by go-swagger cli yet
	}

	return nil, retAdded
}