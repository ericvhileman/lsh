// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/projects"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationProjectsDeleteProjectCmd returns a cmd to handle operation deleteProject
func makeOperationProjectsDeleteProjectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "delete-project",
		Short: ``,
		RunE:  runOperationProjectsDeleteProject,
	}

	if err := registerOperationProjectsDeleteProjectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationProjectsDeleteProject uses cmd flags to call endpoint api
func runOperationProjectsDeleteProject(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := projects.NewDeleteProjectParams()
	if err, _ := retrieveOperationProjectsDeleteProjectAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectsDeleteProjectIDOrSlugFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationProjectsDeleteProjectResult(appCli.Projects.DeleteProject(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationProjectsDeleteProjectParamFlags registers all flags needed to fill params
func registerOperationProjectsDeleteProjectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationProjectsDeleteProjectAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectsDeleteProjectIDOrSlugParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationProjectsDeleteProjectAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationProjectsDeleteProjectIDOrSlugParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idOrSlugDescription := `Required. The project ID or Slug`

	var idOrSlugFlagName string
	if cmdPrefix == "" {
		idOrSlugFlagName = "id_or_slug"
	} else {
		idOrSlugFlagName = fmt.Sprintf("%v.id_or_slug", cmdPrefix)
	}

	var idOrSlugFlagDefault string

	_ = cmd.PersistentFlags().String(idOrSlugFlagName, idOrSlugFlagDefault, idOrSlugDescription)

	return nil
}

func retrieveOperationProjectsDeleteProjectAPIVersionFlag(m *projects.DeleteProjectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationProjectsDeleteProjectIDOrSlugFlag(m *projects.DeleteProjectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id_or_slug") {

		var idOrSlugFlagName string
		if cmdPrefix == "" {
			idOrSlugFlagName = "id_or_slug"
		} else {
			idOrSlugFlagName = fmt.Sprintf("%v.id_or_slug", cmdPrefix)
		}

		idOrSlugFlagValue, err := cmd.Flags().GetString(idOrSlugFlagName)
		if err != nil {
			return err, false
		}
		m.IDOrSlug = idOrSlugFlagValue

	}
	return nil, retAdded
}

// parseOperationProjectsDeleteProjectResult parses request result and return the string content
func parseOperationProjectsDeleteProjectResult(resp0 *projects.DeleteProjectNoContent, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning deleteProjectNoContent is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*projects.DeleteProjectForbidden)
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
		resp2, ok := iResp2.(*projects.DeleteProjectNotFound)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*projects.DeleteProjectNotAcceptable)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response deleteProjectNoContent is not supported by go-swagger cli yet.

	return "", nil
}
