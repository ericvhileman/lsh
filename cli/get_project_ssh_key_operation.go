// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/ssh_keys"
	"github.com/latitudesh/cli/internal/utils"
	"github.com/latitudesh/cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSSHKeysGetProjectSSHKeyCmd returns a cmd to handle operation getProjectSshKey
func makeOperationSSHKeysGetProjectSSHKeyCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "get-project-ssh-key",
		Short: `List all SSH Keys in the project. These keys can be used to access servers after deploy and reinstall actions.
`,
		RunE: runOperationSSHKeysGetProjectSSHKey,
	}

	if err := registerOperationSSHKeysGetProjectSSHKeyParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSSHKeysGetProjectSSHKey uses cmd flags to call endpoint api
func runOperationSSHKeysGetProjectSSHKey(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := ssh_keys.NewGetProjectSSHKeyParams()
	if err, _ := retrieveOperationSSHKeysGetProjectSSHKeyAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSSHKeysGetProjectSSHKeyProjectIDOrSlugFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSSHKeysGetProjectSSHKeySSHKeyIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSSHKeysGetProjectSSHKeyResult(appCli.SSHKeys.GetProjectSSHKey(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(utils.PrettifyJson(msgStr))
	}
	return nil
}

// registerOperationSSHKeysGetProjectSSHKeyParamFlags registers all flags needed to fill params
func registerOperationSSHKeysGetProjectSSHKeyParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSSHKeysGetProjectSSHKeyAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSSHKeysGetProjectSSHKeyProjectIDOrSlugParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSSHKeysGetProjectSSHKeySSHKeyIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSSHKeysGetProjectSSHKeyAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSSHKeysGetProjectSSHKeyProjectIDOrSlugParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSSHKeysGetProjectSSHKeySSHKeyIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	sshKeyIdDescription := `Required. `

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

func retrieveOperationSSHKeysGetProjectSSHKeyAPIVersionFlag(m *ssh_keys.GetProjectSSHKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSSHKeysGetProjectSSHKeyProjectIDOrSlugFlag(m *ssh_keys.GetProjectSSHKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSSHKeysGetProjectSSHKeySSHKeyIDFlag(m *ssh_keys.GetProjectSSHKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("ssh_key_id") {

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

	}
	return nil, retAdded
}

// parseOperationSSHKeysGetProjectSSHKeyResult parses request result and return the string content
func parseOperationSSHKeysGetProjectSSHKeyResult(resp0 *ssh_keys.GetProjectSSHKeyOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*ssh_keys.GetProjectSSHKeyOK)
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
func registerModelGetProjectSSHKeyOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetProjectSSHKeyOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetProjectSSHKeyOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelSSHKeyDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetProjectSSHKeyOKBodyFlags(depth int, m *ssh_keys.GetProjectSSHKeyOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveGetProjectSSHKeyOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveGetProjectSSHKeyOKBodyDataFlags(depth int, m *ssh_keys.GetProjectSSHKeyOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data models.SSHKeyData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.SSHKeyData{}
	}

	err, dataAdded := retrieveModelSSHKeyDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}