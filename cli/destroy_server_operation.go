// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/lsh/client/servers"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

// makeOperationServersDestroyServerCmd returns a cmd to handle operation destroyServer
func makeOperationServersDestroyServerCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "destroy",
		Short: `Remove a server.`,
		RunE:  runOperationServersDestroyServer,
	}

	if err := registerOperationServersDestroyServerParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServersDestroyServer uses cmd flags to call endpoint api
func runOperationServersDestroyServer(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := servers.NewDestroyServerParams()
	if err, _ := retrieveOperationServersDestroyServerServerIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Servers.DestroyServer(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		response.Render()
	}
	return nil
}

// registerOperationServersDestroyServerParamFlags registers all flags needed to fill params
func registerOperationServersDestroyServerParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServersDestroyServerServerIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServersDestroyServerServerIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	serverIdDescription := `Required. The server ID`

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

func retrieveOperationServersDestroyServerServerIDFlag(m *servers.DestroyServerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
