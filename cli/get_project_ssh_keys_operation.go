package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/lsh/client/ssh_keys"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

// makeOperationSSHKeysGetProjectSSHKeysCmd returns a cmd to handle operation getProjectSshKeys
func makeOperationSSHKeysGetProjectSSHKeysCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "list",
		Short: `List all SSH Keys in the project. These keys can be used to access servers after deploy and reinstall actions.`,
		RunE:  runOperationSSHKeysGetProjectSSHKeys,
	}

	if err := registerOperationSSHKeysGetProjectSSHKeysParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSSHKeysGetProjectSSHKeys uses cmd flags to call endpoint api
func runOperationSSHKeysGetProjectSSHKeys(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := ssh_keys.NewGetProjectSSHKeysParams()
	if err, _ := retrieveOperationSSHKeysGetProjectSSHKeysProjectIDOrSlugFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.SSHKeys.GetProjectSSHKeys(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		response.Render()
	}
	return nil
}

// registerOperationSSHKeysGetProjectSSHKeysParamFlags registers all flags needed to fill params
func registerOperationSSHKeysGetProjectSSHKeysParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSSHKeysGetProjectSSHKeysProjectIDOrSlugParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSSHKeysGetProjectSSHKeysProjectIDOrSlugParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationSSHKeysGetProjectSSHKeysProjectIDOrSlugFlag(m *ssh_keys.GetProjectSSHKeysParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
