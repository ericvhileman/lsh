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

// Schema cli for Regions

// register flags to command
func registerModelRegionsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRegionsData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRegionsData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data []*RegionsDataItems0 array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRegionsFlags(depth int, m *models.Regions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveRegionsDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveRegionsDataFlags(depth int, m *models.Regions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// warning: data array type []*RegionsDataItems0 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

// Extra schema cli for RegionsDataItems0

// register flags to command
func registerModelRegionsDataItems0Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRegionsDataItems0Attributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRegionsDataItems0ID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRegionsDataItems0Attributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelRegionsDataItems0AttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerRegionsDataItems0ID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRegionsDataItems0Flags(depth int, m *models.RegionsDataItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveRegionsDataItems0AttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrieveRegionsDataItems0IDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	return nil, retAdded
}

func retrieveRegionsDataItems0AttributesFlags(depth int, m *models.RegionsDataItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes RegionsDataItems0Attributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &models.RegionsDataItems0Attributes{}
	}

	err, attributesAdded := retrieveModelRegionsDataItems0AttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveRegionsDataItems0IDFlags(depth int, m *models.RegionsDataItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// Extra schema cli for RegionsDataItems0Attributes

// register flags to command
func registerModelRegionsDataItems0AttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRegionsDataItems0AttributesCountry(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRegionsDataItems0AttributesName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRegionsDataItems0AttributesSlug(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRegionsDataItems0AttributesCountry(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var countryFlagName string
	if cmdPrefix == "" {
		countryFlagName = "country"
	} else {
		countryFlagName = fmt.Sprintf("%v.country", cmdPrefix)
	}

	if err := registerModelRegionsDataItems0AttributesCountryFlags(depth+1, countryFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerRegionsDataItems0AttributesName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerRegionsDataItems0AttributesSlug(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	slugDescription := ``

	var slugFlagName string
	if cmdPrefix == "" {
		slugFlagName = "slug"
	} else {
		slugFlagName = fmt.Sprintf("%v.slug", cmdPrefix)
	}

	var slugFlagDefault string

	_ = cmd.PersistentFlags().String(slugFlagName, slugFlagDefault, slugDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRegionsDataItems0AttributesFlags(depth int, m *models.RegionsDataItems0Attributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, countryAdded := retrieveRegionsDataItems0AttributesCountryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || countryAdded

	err, nameAdded := retrieveRegionsDataItems0AttributesNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, slugAdded := retrieveRegionsDataItems0AttributesSlugFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || slugAdded

	return nil, retAdded
}

func retrieveRegionsDataItems0AttributesCountryFlags(depth int, m *models.RegionsDataItems0Attributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	countryFlagName := fmt.Sprintf("%v.country", cmdPrefix)
	if cmd.Flags().Changed(countryFlagName) {
		// info: complex object country RegionsDataItems0AttributesCountry is retrieved outside this Changed() block
	}
	countryFlagValue := m.Country
	if swag.IsZero(countryFlagValue) {
		countryFlagValue = &models.RegionsDataItems0AttributesCountry{}
	}

	err, countryAdded := retrieveModelRegionsDataItems0AttributesCountryFlags(depth+1, countryFlagValue, countryFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || countryAdded
	if countryAdded {
		m.Country = countryFlagValue
	}

	return nil, retAdded
}

func retrieveRegionsDataItems0AttributesNameFlags(depth int, m *models.RegionsDataItems0Attributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveRegionsDataItems0AttributesSlugFlags(depth int, m *models.RegionsDataItems0Attributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	slugFlagName := fmt.Sprintf("%v.slug", cmdPrefix)
	if cmd.Flags().Changed(slugFlagName) {

		var slugFlagName string
		if cmdPrefix == "" {
			slugFlagName = "slug"
		} else {
			slugFlagName = fmt.Sprintf("%v.slug", cmdPrefix)
		}

		slugFlagValue, err := cmd.Flags().GetString(slugFlagName)
		if err != nil {
			return err, false
		}
		m.Slug = slugFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for RegionsDataItems0AttributesCountry

// register flags to command
func registerModelRegionsDataItems0AttributesCountryFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRegionsDataItems0AttributesCountryName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRegionsDataItems0AttributesCountrySlug(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRegionsDataItems0AttributesCountryName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerRegionsDataItems0AttributesCountrySlug(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	slugDescription := ``

	var slugFlagName string
	if cmdPrefix == "" {
		slugFlagName = "slug"
	} else {
		slugFlagName = fmt.Sprintf("%v.slug", cmdPrefix)
	}

	var slugFlagDefault string

	_ = cmd.PersistentFlags().String(slugFlagName, slugFlagDefault, slugDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRegionsDataItems0AttributesCountryFlags(depth int, m *models.RegionsDataItems0AttributesCountry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, nameAdded := retrieveRegionsDataItems0AttributesCountryNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, slugAdded := retrieveRegionsDataItems0AttributesCountrySlugFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || slugAdded

	return nil, retAdded
}

func retrieveRegionsDataItems0AttributesCountryNameFlags(depth int, m *models.RegionsDataItems0AttributesCountry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveRegionsDataItems0AttributesCountrySlugFlags(depth int, m *models.RegionsDataItems0AttributesCountry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	slugFlagName := fmt.Sprintf("%v.slug", cmdPrefix)
	if cmd.Flags().Changed(slugFlagName) {

		var slugFlagName string
		if cmdPrefix == "" {
			slugFlagName = "slug"
		} else {
			slugFlagName = fmt.Sprintf("%v.slug", cmdPrefix)
		}

		slugFlagValue, err := cmd.Flags().GetString(slugFlagName)
		if err != nil {
			return err, false
		}
		m.Slug = slugFlagValue

		retAdded = true
	}

	return nil, retAdded
}