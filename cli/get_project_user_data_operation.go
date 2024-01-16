// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/user_data"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationUserDataGetProjectUserDataCmd returns a cmd to handle operation getProjectUserData
func makeOperationUserDataGetProjectUserDataCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "get-project-user-data",
		Short: `Get User Data in the project. These scripts can be used to configure servers with user data.
`,
		RunE: runOperationUserDataGetProjectUserData,
	}

	if err := registerOperationUserDataGetProjectUserDataParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationUserDataGetProjectUserData uses cmd flags to call endpoint api
func runOperationUserDataGetProjectUserData(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := user_data.NewGetProjectUserDataParams()
	if err, _ := retrieveOperationUserDataGetProjectUserDataAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUserDataGetProjectUserDataExtraFieldsUserDataFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUserDataGetProjectUserDataProjectIDOrSlugFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUserDataGetProjectUserDataUserDataIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationUserDataGetProjectUserDataResult(appCli.UserData.GetProjectUserData(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(utils.PrettifyJson(msgStr))
	}
	return nil
}

// registerOperationUserDataGetProjectUserDataParamFlags registers all flags needed to fill params
func registerOperationUserDataGetProjectUserDataParamFlags(cmd *cobra.Command) error {
	if err := registerOperationUserDataGetProjectUserDataAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUserDataGetProjectUserDataExtraFieldsUserDataParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUserDataGetProjectUserDataProjectIDOrSlugParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUserDataGetProjectUserDataUserDataIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationUserDataGetProjectUserDataAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationUserDataGetProjectUserDataExtraFieldsUserDataParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	extraFieldsUserDataDescription := `The ` + "`" + `decoded_content` + "`" + ` is provided as an extra attribute that shows content in decoded form.`

	var extraFieldsUserDataFlagName string
	if cmdPrefix == "" {
		extraFieldsUserDataFlagName = "extra_fields[user_data]"
	} else {
		extraFieldsUserDataFlagName = fmt.Sprintf("%v.extra_fields[user_data]", cmdPrefix)
	}

	var extraFieldsUserDataFlagDefault string = "decoded_content"

	_ = cmd.PersistentFlags().String(extraFieldsUserDataFlagName, extraFieldsUserDataFlagDefault, extraFieldsUserDataDescription)

	return nil
}
func registerOperationUserDataGetProjectUserDataProjectIDOrSlugParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	projectIdOrSlugDescription := `Required. `

	var projectIdOrSlugFlagName string
	if cmdPrefix == "" {
		projectIdOrSlugFlagName = "project_id_or_slug"
	} else {
		projectIdOrSlugFlagName = fmt.Sprintf("%v.project_id_or_slug", cmdPrefix)
	}

	var projectIdOrSlugFlagDefault string

	_ = cmd.PersistentFlags().String(projectIdOrSlugFlagName, projectIdOrSlugFlagDefault, projectIdOrSlugDescription)

	return nil
}
func registerOperationUserDataGetProjectUserDataUserDataIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	userDataIdDescription := `Required. `

	var userDataIdFlagName string
	if cmdPrefix == "" {
		userDataIdFlagName = "user_data_id"
	} else {
		userDataIdFlagName = fmt.Sprintf("%v.user_data_id", cmdPrefix)
	}

	var userDataIdFlagDefault string

	_ = cmd.PersistentFlags().String(userDataIdFlagName, userDataIdFlagDefault, userDataIdDescription)

	return nil
}

func retrieveOperationUserDataGetProjectUserDataAPIVersionFlag(m *user_data.GetProjectUserDataParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationUserDataGetProjectUserDataExtraFieldsUserDataFlag(m *user_data.GetProjectUserDataParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("extra_fields[user_data]") {

		var extraFieldsUserDataFlagName string
		if cmdPrefix == "" {
			extraFieldsUserDataFlagName = "extra_fields[user_data]"
		} else {
			extraFieldsUserDataFlagName = fmt.Sprintf("%v.extra_fields[user_data]", cmdPrefix)
		}

		extraFieldsUserDataFlagValue, err := cmd.Flags().GetString(extraFieldsUserDataFlagName)
		if err != nil {
			return err, false
		}
		m.ExtraFieldsUserData = &extraFieldsUserDataFlagValue

	}
	return nil, retAdded
}
func retrieveOperationUserDataGetProjectUserDataProjectIDOrSlugFlag(m *user_data.GetProjectUserDataParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("project_id_or_slug") {

		var projectIdOrSlugFlagName string
		if cmdPrefix == "" {
			projectIdOrSlugFlagName = "project_id_or_slug"
		} else {
			projectIdOrSlugFlagName = fmt.Sprintf("%v.project_id_or_slug", cmdPrefix)
		}

		projectIdOrSlugFlagValue, err := cmd.Flags().GetString(projectIdOrSlugFlagName)
		if err != nil {
			return err, false
		}
		m.ProjectIDOrSlug = projectIdOrSlugFlagValue

	}
	return nil, retAdded
}
func retrieveOperationUserDataGetProjectUserDataUserDataIDFlag(m *user_data.GetProjectUserDataParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("user_data_id") {

		var userDataIdFlagName string
		if cmdPrefix == "" {
			userDataIdFlagName = "user_data_id"
		} else {
			userDataIdFlagName = fmt.Sprintf("%v.user_data_id", cmdPrefix)
		}

		userDataIdFlagValue, err := cmd.Flags().GetString(userDataIdFlagName)
		if err != nil {
			return err, false
		}
		m.UserDataID = userDataIdFlagValue

	}
	return nil, retAdded
}

// parseOperationUserDataGetProjectUserDataResult parses request result and return the string content
func parseOperationUserDataGetProjectUserDataResult(resp0 *user_data.GetProjectUserDataOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*user_data.GetProjectUserDataOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning getProjectUserDataNotFound is not supported

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