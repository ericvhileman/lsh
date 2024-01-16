// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/members"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationMembersDestroyTeamMemberCmd returns a cmd to handle operation destroyTeamMember
func makeOperationMembersDestroyTeamMemberCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "destroy-team-member",
		Short: ``,
		RunE:  runOperationMembersDestroyTeamMember,
	}

	if err := registerOperationMembersDestroyTeamMemberParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationMembersDestroyTeamMember uses cmd flags to call endpoint api
func runOperationMembersDestroyTeamMember(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := members.NewDestroyTeamMemberParams()
	if err, _ := retrieveOperationMembersDestroyTeamMemberAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationMembersDestroyTeamMemberUserIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationMembersDestroyTeamMemberResult(appCli.Members.DestroyTeamMember(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(utils.PrettifyJson(msgStr))
	}
	return nil
}

// registerOperationMembersDestroyTeamMemberParamFlags registers all flags needed to fill params
func registerOperationMembersDestroyTeamMemberParamFlags(cmd *cobra.Command) error {
	if err := registerOperationMembersDestroyTeamMemberAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationMembersDestroyTeamMemberUserIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationMembersDestroyTeamMemberAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationMembersDestroyTeamMemberUserIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	userIdDescription := `Required. The user ID`

	var userIdFlagName string
	if cmdPrefix == "" {
		userIdFlagName = "user_id"
	} else {
		userIdFlagName = fmt.Sprintf("%v.user_id", cmdPrefix)
	}

	var userIdFlagDefault string

	_ = cmd.PersistentFlags().String(userIdFlagName, userIdFlagDefault, userIdDescription)

	return nil
}

func retrieveOperationMembersDestroyTeamMemberAPIVersionFlag(m *members.DestroyTeamMemberParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationMembersDestroyTeamMemberUserIDFlag(m *members.DestroyTeamMemberParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("user_id") {

		var userIdFlagName string
		if cmdPrefix == "" {
			userIdFlagName = "user_id"
		} else {
			userIdFlagName = fmt.Sprintf("%v.user_id", cmdPrefix)
		}

		userIdFlagValue, err := cmd.Flags().GetString(userIdFlagName)
		if err != nil {
			return err, false
		}
		m.UserID = userIdFlagValue

	}
	return nil, retAdded
}

// parseOperationMembersDestroyTeamMemberResult parses request result and return the string content
func parseOperationMembersDestroyTeamMemberResult(resp0 *members.DestroyTeamMemberOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning destroyTeamMemberOK is not supported

		// Non schema case: warning destroyTeamMemberForbidden is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*members.DestroyTeamMemberNotFound)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response destroyTeamMemberOK is not supported by go-swagger cli yet.

	return "", nil
}