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

// Schema cli for IpmiSession

// register flags to command
func registerModelIpmiSessionFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerIpmiSessionData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerIpmiSessionData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelIpmiSessionDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelIpmiSessionFlags(depth int, m *models.IpmiSession, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveIpmiSessionDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveIpmiSessionDataFlags(depth int, m *models.IpmiSession, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data IpmiSessionData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.IpmiSessionData{}
	}

	err, dataAdded := retrieveModelIpmiSessionDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}

// Extra schema cli for IpmiSessionData

// register flags to command
func registerModelIpmiSessionDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerIpmiSessionDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIpmiSessionDataID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIpmiSessionDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerIpmiSessionDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelIpmiSessionDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerIpmiSessionDataID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerIpmiSessionDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["ipmi_sessions"]. `

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
			if err := json.Unmarshal([]byte(`["ipmi_sessions"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelIpmiSessionDataFlags(depth int, m *models.IpmiSessionData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveIpmiSessionDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrieveIpmiSessionDataIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, typeAdded := retrieveIpmiSessionDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveIpmiSessionDataAttributesFlags(depth int, m *models.IpmiSessionData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes IpmiSessionDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &models.IpmiSessionDataAttributes{}
	}

	err, attributesAdded := retrieveModelIpmiSessionDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveIpmiSessionDataIDFlags(depth int, m *models.IpmiSessionData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveIpmiSessionDataTypeFlags(depth int, m *models.IpmiSessionData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// Extra schema cli for IpmiSessionDataAttributes

// register flags to command
func registerModelIpmiSessionDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerIpmiSessionDataAttributesIpmiAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIpmiSessionDataAttributesIpmiPassword(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIpmiSessionDataAttributesIpmiUsername(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerIpmiSessionDataAttributesIpmiAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipmiAddressDescription := `The IPMI IP Address`

	var ipmiAddressFlagName string
	if cmdPrefix == "" {
		ipmiAddressFlagName = "ipmi_address"
	} else {
		ipmiAddressFlagName = fmt.Sprintf("%v.ipmi_address", cmdPrefix)
	}

	var ipmiAddressFlagDefault string

	_ = cmd.PersistentFlags().String(ipmiAddressFlagName, ipmiAddressFlagDefault, ipmiAddressDescription)

	return nil
}

func registerIpmiSessionDataAttributesIpmiPassword(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipmiPasswordDescription := `The IPMI password`

	var ipmiPasswordFlagName string
	if cmdPrefix == "" {
		ipmiPasswordFlagName = "ipmi_password"
	} else {
		ipmiPasswordFlagName = fmt.Sprintf("%v.ipmi_password", cmdPrefix)
	}

	var ipmiPasswordFlagDefault string

	_ = cmd.PersistentFlags().String(ipmiPasswordFlagName, ipmiPasswordFlagDefault, ipmiPasswordDescription)

	return nil
}

func registerIpmiSessionDataAttributesIpmiUsername(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipmiUsernameDescription := `The IPMI username`

	var ipmiUsernameFlagName string
	if cmdPrefix == "" {
		ipmiUsernameFlagName = "ipmi_username"
	} else {
		ipmiUsernameFlagName = fmt.Sprintf("%v.ipmi_username", cmdPrefix)
	}

	var ipmiUsernameFlagDefault string

	_ = cmd.PersistentFlags().String(ipmiUsernameFlagName, ipmiUsernameFlagDefault, ipmiUsernameDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelIpmiSessionDataAttributesFlags(depth int, m *models.IpmiSessionDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, ipmiAddressAdded := retrieveIpmiSessionDataAttributesIpmiAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipmiAddressAdded

	err, ipmiPasswordAdded := retrieveIpmiSessionDataAttributesIpmiPasswordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipmiPasswordAdded

	err, ipmiUsernameAdded := retrieveIpmiSessionDataAttributesIpmiUsernameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipmiUsernameAdded

	return nil, retAdded
}

func retrieveIpmiSessionDataAttributesIpmiAddressFlags(depth int, m *models.IpmiSessionDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ipmiAddressFlagName := fmt.Sprintf("%v.ipmi_address", cmdPrefix)
	if cmd.Flags().Changed(ipmiAddressFlagName) {

		var ipmiAddressFlagName string
		if cmdPrefix == "" {
			ipmiAddressFlagName = "ipmi_address"
		} else {
			ipmiAddressFlagName = fmt.Sprintf("%v.ipmi_address", cmdPrefix)
		}

		ipmiAddressFlagValue, err := cmd.Flags().GetString(ipmiAddressFlagName)
		if err != nil {
			return err, false
		}
		m.IpmiAddress = ipmiAddressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIpmiSessionDataAttributesIpmiPasswordFlags(depth int, m *models.IpmiSessionDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ipmiPasswordFlagName := fmt.Sprintf("%v.ipmi_password", cmdPrefix)
	if cmd.Flags().Changed(ipmiPasswordFlagName) {

		var ipmiPasswordFlagName string
		if cmdPrefix == "" {
			ipmiPasswordFlagName = "ipmi_password"
		} else {
			ipmiPasswordFlagName = fmt.Sprintf("%v.ipmi_password", cmdPrefix)
		}

		ipmiPasswordFlagValue, err := cmd.Flags().GetString(ipmiPasswordFlagName)
		if err != nil {
			return err, false
		}
		m.IpmiPassword = ipmiPasswordFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIpmiSessionDataAttributesIpmiUsernameFlags(depth int, m *models.IpmiSessionDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ipmiUsernameFlagName := fmt.Sprintf("%v.ipmi_username", cmdPrefix)
	if cmd.Flags().Changed(ipmiUsernameFlagName) {

		var ipmiUsernameFlagName string
		if cmdPrefix == "" {
			ipmiUsernameFlagName = "ipmi_username"
		} else {
			ipmiUsernameFlagName = fmt.Sprintf("%v.ipmi_username", cmdPrefix)
		}

		ipmiUsernameFlagValue, err := cmd.Flags().GetString(ipmiUsernameFlagName)
		if err != nil {
			return err, false
		}
		m.IpmiUsername = ipmiUsernameFlagValue

		retAdded = true
	}

	return nil, retAdded
}