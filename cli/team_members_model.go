// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/latitudesh/cli/models"

	"github.com/spf13/cobra"
)

// Schema cli for TeamMembers

// register flags to command
func registerModelTeamMembersFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTeamMembersData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTeamMembersData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data []*TeamMembersDataItems0 array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTeamMembersFlags(depth int, m *models.TeamMembers, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveTeamMembersDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveTeamMembersDataFlags(depth int, m *models.TeamMembers, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// warning: data array type []*TeamMembersDataItems0 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

// Extra schema cli for TeamMembersDataItems0

// register flags to command
func registerModelTeamMembersDataItems0Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTeamMembersDataItems0Email(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTeamMembersDataItems0FirstName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTeamMembersDataItems0LastName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTeamMembersDataItems0MfaEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTeamMembersDataItems0Role(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTeamMembersDataItems0Email(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	emailDescription := ``

	var emailFlagName string
	if cmdPrefix == "" {
		emailFlagName = "email"
	} else {
		emailFlagName = fmt.Sprintf("%v.email", cmdPrefix)
	}

	var emailFlagDefault string

	_ = cmd.PersistentFlags().String(emailFlagName, emailFlagDefault, emailDescription)

	return nil
}

func registerTeamMembersDataItems0FirstName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	firstNameDescription := ``

	var firstNameFlagName string
	if cmdPrefix == "" {
		firstNameFlagName = "first_name"
	} else {
		firstNameFlagName = fmt.Sprintf("%v.first_name", cmdPrefix)
	}

	var firstNameFlagDefault string

	_ = cmd.PersistentFlags().String(firstNameFlagName, firstNameFlagDefault, firstNameDescription)

	return nil
}

func registerTeamMembersDataItems0LastName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	lastNameDescription := ``

	var lastNameFlagName string
	if cmdPrefix == "" {
		lastNameFlagName = "last_name"
	} else {
		lastNameFlagName = fmt.Sprintf("%v.last_name", cmdPrefix)
	}

	var lastNameFlagDefault string

	_ = cmd.PersistentFlags().String(lastNameFlagName, lastNameFlagDefault, lastNameDescription)

	return nil
}

func registerTeamMembersDataItems0MfaEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	mfaEnabledDescription := ``

	var mfaEnabledFlagName string
	if cmdPrefix == "" {
		mfaEnabledFlagName = "mfa_enabled"
	} else {
		mfaEnabledFlagName = fmt.Sprintf("%v.mfa_enabled", cmdPrefix)
	}

	var mfaEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(mfaEnabledFlagName, mfaEnabledFlagDefault, mfaEnabledDescription)

	return nil
}

func registerTeamMembersDataItems0Role(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var roleFlagName string
	if cmdPrefix == "" {
		roleFlagName = "role"
	} else {
		roleFlagName = fmt.Sprintf("%v.role", cmdPrefix)
	}

	if err := registerModelTeamMembersDataItems0RoleFlags(depth+1, roleFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTeamMembersDataItems0Flags(depth int, m *models.TeamMembersDataItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, emailAdded := retrieveTeamMembersDataItems0EmailFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || emailAdded

	err, firstNameAdded := retrieveTeamMembersDataItems0FirstNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || firstNameAdded

	err, lastNameAdded := retrieveTeamMembersDataItems0LastNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lastNameAdded

	err, mfaEnabledAdded := retrieveTeamMembersDataItems0MfaEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || mfaEnabledAdded

	err, roleAdded := retrieveTeamMembersDataItems0RoleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || roleAdded

	return nil, retAdded
}

func retrieveTeamMembersDataItems0EmailFlags(depth int, m *models.TeamMembersDataItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	emailFlagName := fmt.Sprintf("%v.email", cmdPrefix)
	if cmd.Flags().Changed(emailFlagName) {

		var emailFlagName string
		if cmdPrefix == "" {
			emailFlagName = "email"
		} else {
			emailFlagName = fmt.Sprintf("%v.email", cmdPrefix)
		}

		emailFlagValue, err := cmd.Flags().GetString(emailFlagName)
		if err != nil {
			return err, false
		}
		m.Email = emailFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTeamMembersDataItems0FirstNameFlags(depth int, m *models.TeamMembersDataItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	firstNameFlagName := fmt.Sprintf("%v.first_name", cmdPrefix)
	if cmd.Flags().Changed(firstNameFlagName) {

		var firstNameFlagName string
		if cmdPrefix == "" {
			firstNameFlagName = "first_name"
		} else {
			firstNameFlagName = fmt.Sprintf("%v.first_name", cmdPrefix)
		}

		firstNameFlagValue, err := cmd.Flags().GetString(firstNameFlagName)
		if err != nil {
			return err, false
		}
		m.FirstName = firstNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTeamMembersDataItems0LastNameFlags(depth int, m *models.TeamMembersDataItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	lastNameFlagName := fmt.Sprintf("%v.last_name", cmdPrefix)
	if cmd.Flags().Changed(lastNameFlagName) {

		var lastNameFlagName string
		if cmdPrefix == "" {
			lastNameFlagName = "last_name"
		} else {
			lastNameFlagName = fmt.Sprintf("%v.last_name", cmdPrefix)
		}

		lastNameFlagValue, err := cmd.Flags().GetString(lastNameFlagName)
		if err != nil {
			return err, false
		}
		m.LastName = lastNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTeamMembersDataItems0MfaEnabledFlags(depth int, m *models.TeamMembersDataItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	mfaEnabledFlagName := fmt.Sprintf("%v.mfa_enabled", cmdPrefix)
	if cmd.Flags().Changed(mfaEnabledFlagName) {

		var mfaEnabledFlagName string
		if cmdPrefix == "" {
			mfaEnabledFlagName = "mfa_enabled"
		} else {
			mfaEnabledFlagName = fmt.Sprintf("%v.mfa_enabled", cmdPrefix)
		}

		mfaEnabledFlagValue, err := cmd.Flags().GetBool(mfaEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.MfaEnabled = mfaEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTeamMembersDataItems0RoleFlags(depth int, m *models.TeamMembersDataItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	roleFlagName := fmt.Sprintf("%v.role", cmdPrefix)
	if cmd.Flags().Changed(roleFlagName) {
		// info: complex object role TeamMembersDataItems0Role is retrieved outside this Changed() block
	}
	roleFlagValue := m.Role
	if swag.IsZero(roleFlagValue) {
		roleFlagValue = &models.TeamMembersDataItems0Role{}
	}

	err, roleAdded := retrieveModelTeamMembersDataItems0RoleFlags(depth+1, roleFlagValue, roleFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || roleAdded
	if roleAdded {
		m.Role = roleFlagValue
	}

	return nil, retAdded
}

// Extra schema cli for TeamMembersDataItems0Role

// register flags to command
func registerModelTeamMembersDataItems0RoleFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTeamMembersDataItems0RoleCreatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTeamMembersDataItems0RoleID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTeamMembersDataItems0RoleName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTeamMembersDataItems0RoleUpdatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTeamMembersDataItems0RoleCreatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	createdAtDescription := ``

	var createdAtFlagName string
	if cmdPrefix == "" {
		createdAtFlagName = "created_at"
	} else {
		createdAtFlagName = fmt.Sprintf("%v.created_at", cmdPrefix)
	}

	var createdAtFlagDefault string

	_ = cmd.PersistentFlags().String(createdAtFlagName, createdAtFlagDefault, createdAtDescription)

	return nil
}

func registerTeamMembersDataItems0RoleID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerTeamMembersDataItems0RoleName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerTeamMembersDataItems0RoleUpdatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	updatedAtDescription := ``

	var updatedAtFlagName string
	if cmdPrefix == "" {
		updatedAtFlagName = "updated_at"
	} else {
		updatedAtFlagName = fmt.Sprintf("%v.updated_at", cmdPrefix)
	}

	var updatedAtFlagDefault string

	_ = cmd.PersistentFlags().String(updatedAtFlagName, updatedAtFlagDefault, updatedAtDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTeamMembersDataItems0RoleFlags(depth int, m *models.TeamMembersDataItems0Role, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, createdAtAdded := retrieveTeamMembersDataItems0RoleCreatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAtAdded

	err, idAdded := retrieveTeamMembersDataItems0RoleIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, nameAdded := retrieveTeamMembersDataItems0RoleNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, updatedAtAdded := retrieveTeamMembersDataItems0RoleUpdatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || updatedAtAdded

	return nil, retAdded
}

func retrieveTeamMembersDataItems0RoleCreatedAtFlags(depth int, m *models.TeamMembersDataItems0Role, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	createdAtFlagName := fmt.Sprintf("%v.created_at", cmdPrefix)
	if cmd.Flags().Changed(createdAtFlagName) {

		var createdAtFlagName string
		if cmdPrefix == "" {
			createdAtFlagName = "created_at"
		} else {
			createdAtFlagName = fmt.Sprintf("%v.created_at", cmdPrefix)
		}

		createdAtFlagValue, err := cmd.Flags().GetString(createdAtFlagName)
		if err != nil {
			return err, false
		}
		m.CreatedAt = createdAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTeamMembersDataItems0RoleIDFlags(depth int, m *models.TeamMembersDataItems0Role, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveTeamMembersDataItems0RoleNameFlags(depth int, m *models.TeamMembersDataItems0Role, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTeamMembersDataItems0RoleUpdatedAtFlags(depth int, m *models.TeamMembersDataItems0Role, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	updatedAtFlagName := fmt.Sprintf("%v.updated_at", cmdPrefix)
	if cmd.Flags().Changed(updatedAtFlagName) {

		var updatedAtFlagName string
		if cmdPrefix == "" {
			updatedAtFlagName = "updated_at"
		} else {
			updatedAtFlagName = fmt.Sprintf("%v.updated_at", cmdPrefix)
		}

		updatedAtFlagValue, err := cmd.Flags().GetString(updatedAtFlagName)
		if err != nil {
			return err, false
		}
		m.UpdatedAt = updatedAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}
