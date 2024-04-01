package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/lsh/client/servers"
	"github.com/latitudesh/lsh/cmd/lsh"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

// makeOperationServersGetServerCmd returns a cmd to handle operation getServer
func makeOperationServersGetServerCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "get",
		Short: `Returns a server that belongs to the team.`,
		RunE:  runOperationServersGetServer,
	}

	if err := registerOperationServersGetServerParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServersGetServer uses cmd flags to call endpoint api
func runOperationServersGetServer(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := servers.NewGetServerParams()
	if err, _ := retrieveOperationServersGetServerExtraFieldsServersFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServerServerIDFlag(params, "", cmd); err != nil {
		return err
	}
	if lsh.DryRun {

		lsh.LogDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Servers.GetServer(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !lsh.Debug {
		utils.Render(response.GetData())
	}
	return nil
}

// registerOperationServersGetServerParamFlags registers all flags needed to fill params
func registerOperationServersGetServerParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServersGetServerExtraFieldsServersParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServerServerIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServersGetServerExtraFieldsServersParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	extraFieldsServersDescription := `The ` + "`" + `credentials` + "`" + ` are provided as extra attributes that is lazy loaded. To request it, just set ` + "`" + `extra_fields[servers]=credentials` + "`" + ` in the query string.`

	var extraFieldsServersFlagName string
	if cmdPrefix == "" {
		extraFieldsServersFlagName = "extra_fields[servers]"
	} else {
		extraFieldsServersFlagName = fmt.Sprintf("%v.extra_fields[servers]", cmdPrefix)
	}

	var extraFieldsServersFlagDefault string

	_ = cmd.PersistentFlags().String(extraFieldsServersFlagName, extraFieldsServersFlagDefault, extraFieldsServersDescription)

	return nil
}
func registerOperationServersGetServerServerIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	serverIdDescription := `Required. The Server ID`

	var serverIdFlagName string
	if cmdPrefix == "" {
		serverIdFlagName = "id"
	} else {
		serverIdFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var serverIdFlagDefault string

	_ = cmd.PersistentFlags().String(serverIdFlagName, serverIdFlagDefault, serverIdDescription)
	cmd.MarkPersistentFlagRequired(serverIdFlagName)

	return nil
}

func retrieveOperationServersGetServerExtraFieldsServersFlag(m *servers.GetServerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("extra_fields[servers]") {

		var extraFieldsServersFlagName string
		if cmdPrefix == "" {
			extraFieldsServersFlagName = "extra_fields[servers]"
		} else {
			extraFieldsServersFlagName = fmt.Sprintf("%v.extra_fields[servers]", cmdPrefix)
		}

		extraFieldsServersFlagValue, err := cmd.Flags().GetString(extraFieldsServersFlagName)
		if err != nil {
			return err, false
		}
		m.ExtraFieldsServers = &extraFieldsServersFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServerServerIDFlag(m *servers.GetServerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var serverIdFlagName string
		if cmdPrefix == "" {
			serverIdFlagName = "id"
		} else {
			serverIdFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		serverIdFlagValue, err := cmd.Flags().GetString(serverIdFlagName)
		if err != nil {
			return err, false
		}
		m.ServerID = serverIdFlagValue

	}
	return nil, retAdded
}
