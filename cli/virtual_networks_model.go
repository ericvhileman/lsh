// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for VirtualNetworks

// register flags to command
func registerModelVirtualNetworksFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVirtualNetworksData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVirtualNetworksMeta(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVirtualNetworksData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data []*VirtualNetwork array type is not supported by go-swagger cli yet

	return nil
}

func registerVirtualNetworksMeta(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: meta interface{} map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVirtualNetworksFlags(depth int, m *models.VirtualNetworks, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveVirtualNetworksDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	err, metaAdded := retrieveVirtualNetworksMetaFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || metaAdded

	return nil, retAdded
}

func retrieveVirtualNetworksDataFlags(depth int, m *models.VirtualNetworks, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// warning: data array type []*VirtualNetwork is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveVirtualNetworksMetaFlags(depth int, m *models.VirtualNetworks, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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