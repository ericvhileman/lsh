// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for IPAddresses

// register flags to command
func registerModelIPAddressesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerIPAddressesData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerIPAddressesData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data []*IPAddress array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelIPAddressesFlags(depth int, m *models.IPAddresses, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveIPAddressesDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveIPAddressesDataFlags(depth int, m *models.IPAddresses, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// warning: data array type []*IPAddress is not supported by go-swagger cli yet
	}

	return nil, retAdded
}