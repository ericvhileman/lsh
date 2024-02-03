package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/lsh/client/ssh_keys"
	"github.com/latitudesh/lsh/internal/utils"
	"github.com/latitudesh/lsh/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSSHKeysGetProjectSSHKeyCmd returns a cmd to handle operation getProjectSshKey
func makeOperationSSHKeysGetProjectSSHKeyCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "get",
		Short: `Returns a SSH Key in the project. These keys can be used to access servers after deploy and reinstall actions.`,
		RunE:  runOperationSSHKeysGetProjectSSHKey,
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

	response, err := appCli.SSHKeys.GetProjectSSHKey(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetPayload())
	}
	return nil
}

// registerOperationSSHKeysGetProjectSSHKeyParamFlags registers all flags needed to fill params
func registerOperationSSHKeysGetProjectSSHKeyParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSSHKeysGetProjectSSHKeyProjectIDOrSlugParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSSHKeysGetProjectSSHKeySSHKeyIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSSHKeysGetProjectSSHKeyProjectIDOrSlugParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	projectIdOrSlugDescription := `Project Id or Slug (Required).`

	var projectIdOrSlugFlagName string
	if cmdPrefix == "" {
		projectIdOrSlugFlagName = "project"
	} else {
		projectIdOrSlugFlagName = fmt.Sprintf("%v.project", cmdPrefix)
	}

	var projectIdOrSlugFlagDefault string

	_ = cmd.PersistentFlags().String(projectIdOrSlugFlagName, projectIdOrSlugFlagDefault, projectIdOrSlugDescription)
	cmd.MarkPersistentFlagRequired(projectIdOrSlugFlagName)

	return nil
}
func registerOperationSSHKeysGetProjectSSHKeySSHKeyIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	sshKeyIdDescription := `Required. `

	var sshKeyIdFlagName string
	if cmdPrefix == "" {
		sshKeyIdFlagName = "id"
	} else {
		sshKeyIdFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var sshKeyIdFlagDefault string

	_ = cmd.PersistentFlags().String(sshKeyIdFlagName, sshKeyIdFlagDefault, sshKeyIdDescription)
	cmd.MarkPersistentFlagRequired(sshKeyIdFlagName)

	return nil
}

func retrieveOperationSSHKeysGetProjectSSHKeyProjectIDOrSlugFlag(m *ssh_keys.GetProjectSSHKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("project") {

		var projectIdOrSlugFlagName string
		if cmdPrefix == "" {
			projectIdOrSlugFlagName = "project"
		} else {
			projectIdOrSlugFlagName = fmt.Sprintf("%v.project", cmdPrefix)
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
	if cmd.Flags().Changed("id") {

		var sshKeyIdFlagName string
		if cmdPrefix == "" {
			sshKeyIdFlagName = "id"
		} else {
			sshKeyIdFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		sshKeyIdFlagValue, err := cmd.Flags().GetString(sshKeyIdFlagName)
		if err != nil {
			return err, false
		}
		m.SSHKeyID = sshKeyIdFlagValue

	}
	return nil, retAdded
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
