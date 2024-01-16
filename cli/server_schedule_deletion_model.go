// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/latitudesh/cli/models"

	"github.com/spf13/cobra"
)

// Schema cli for ServerScheduleDeletion

// register flags to command
func registerModelServerScheduleDeletionFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServerScheduleDeletionData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServerScheduleDeletionData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelServerScheduleDeletionDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServerScheduleDeletionFlags(depth int, m *models.ServerScheduleDeletion, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveServerScheduleDeletionDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveServerScheduleDeletionDataFlags(depth int, m *models.ServerScheduleDeletion, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data ServerScheduleDeletionData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.ServerScheduleDeletionData{}
	}

	err, dataAdded := retrieveModelServerScheduleDeletionDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}

// Extra schema cli for ServerScheduleDeletionData

// register flags to command
func registerModelServerScheduleDeletionDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServerScheduleDeletionDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerScheduleDeletionDataID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerScheduleDeletionDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServerScheduleDeletionDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelServerScheduleDeletionDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerServerScheduleDeletionDataID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerServerScheduleDeletionDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["schedule_deletion"]. `

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["schedule_deletion"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServerScheduleDeletionDataFlags(depth int, m *models.ServerScheduleDeletionData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveServerScheduleDeletionDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrieveServerScheduleDeletionDataIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, typeAdded := retrieveServerScheduleDeletionDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveServerScheduleDeletionDataAttributesFlags(depth int, m *models.ServerScheduleDeletionData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes ServerScheduleDeletionDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &models.ServerScheduleDeletionDataAttributes{}
	}

	err, attributesAdded := retrieveModelServerScheduleDeletionDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveServerScheduleDeletionDataIDFlags(depth int, m *models.ServerScheduleDeletionData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerScheduleDeletionDataTypeFlags(depth int, m *models.ServerScheduleDeletionData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for ServerScheduleDeletionDataAttributes

// register flags to command
func registerModelServerScheduleDeletionDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServerScheduleDeletionDataAttributesScheduledDeletionAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerScheduleDeletionDataAttributesServerID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServerScheduleDeletionDataAttributesScheduledDeletionAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	scheduledDeletionAtDescription := ``

	var scheduledDeletionAtFlagName string
	if cmdPrefix == "" {
		scheduledDeletionAtFlagName = "scheduled_deletion_at"
	} else {
		scheduledDeletionAtFlagName = fmt.Sprintf("%v.scheduled_deletion_at", cmdPrefix)
	}

	var scheduledDeletionAtFlagDefault string

	_ = cmd.PersistentFlags().String(scheduledDeletionAtFlagName, scheduledDeletionAtFlagDefault, scheduledDeletionAtDescription)

	return nil
}

func registerServerScheduleDeletionDataAttributesServerID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	serverIdDescription := ``

	var serverIdFlagName string
	if cmdPrefix == "" {
		serverIdFlagName = "server_id"
	} else {
		serverIdFlagName = fmt.Sprintf("%v.server_id", cmdPrefix)
	}

	var serverIdFlagDefault string

	_ = cmd.PersistentFlags().String(serverIdFlagName, serverIdFlagDefault, serverIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServerScheduleDeletionDataAttributesFlags(depth int, m *models.ServerScheduleDeletionDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, scheduledDeletionAtAdded := retrieveServerScheduleDeletionDataAttributesScheduledDeletionAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scheduledDeletionAtAdded

	err, serverIdAdded := retrieveServerScheduleDeletionDataAttributesServerIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || serverIdAdded

	return nil, retAdded
}

func retrieveServerScheduleDeletionDataAttributesScheduledDeletionAtFlags(depth int, m *models.ServerScheduleDeletionDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	scheduledDeletionAtFlagName := fmt.Sprintf("%v.scheduled_deletion_at", cmdPrefix)
	if cmd.Flags().Changed(scheduledDeletionAtFlagName) {

		var scheduledDeletionAtFlagName string
		if cmdPrefix == "" {
			scheduledDeletionAtFlagName = "scheduled_deletion_at"
		} else {
			scheduledDeletionAtFlagName = fmt.Sprintf("%v.scheduled_deletion_at", cmdPrefix)
		}

		scheduledDeletionAtFlagValue, err := cmd.Flags().GetString(scheduledDeletionAtFlagName)
		if err != nil {
			return err, false
		}
		m.ScheduledDeletionAt = scheduledDeletionAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerScheduleDeletionDataAttributesServerIDFlags(depth int, m *models.ServerScheduleDeletionDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	serverIdFlagName := fmt.Sprintf("%v.server_id", cmdPrefix)
	if cmd.Flags().Changed(serverIdFlagName) {

		var serverIdFlagName string
		if cmdPrefix == "" {
			serverIdFlagName = "server_id"
		} else {
			serverIdFlagName = fmt.Sprintf("%v.server_id", cmdPrefix)
		}

		serverIdFlagValue, err := cmd.Flags().GetString(serverIdFlagName)
		if err != nil {
			return err, false
		}
		m.ServerID = serverIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}