// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/ip_addresses"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationIPAddressesGetIpsCmd returns a cmd to handle operation getIps
func makeOperationIPAddressesGetIpsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "get-ips",
		Short: `List all Management and Additional IP Addresses.
 • Management IPs are IPs that are used for the management IP of a device.
   This is a public IP address that a device is born and dies with. It never changes during the lifecycle of the device.
 • Additional IPs are individual IPs that can be added to a device as an additional IP that can be used.
`,
		RunE: runOperationIPAddressesGetIps,
	}

	if err := registerOperationIPAddressesGetIpsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationIPAddressesGetIps uses cmd flags to call endpoint api
func runOperationIPAddressesGetIps(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := ip_addresses.NewGetIpsParams()
	if err, _ := retrieveOperationIPAddressesGetIpsAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationIPAddressesGetIpsExtraFieldsIPAddressesFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationIPAddressesGetIpsFilterAddressFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationIPAddressesGetIpsFilterFamilyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationIPAddressesGetIpsFilterLocationFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationIPAddressesGetIpsFilterProjectFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationIPAddressesGetIpsFilterServerFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationIPAddressesGetIpsFilterTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationIPAddressesGetIpsResult(appCli.IPAddresses.GetIps(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationIPAddressesGetIpsParamFlags registers all flags needed to fill params
func registerOperationIPAddressesGetIpsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationIPAddressesGetIpsAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationIPAddressesGetIpsExtraFieldsIPAddressesParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationIPAddressesGetIpsFilterAddressParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationIPAddressesGetIpsFilterFamilyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationIPAddressesGetIpsFilterLocationParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationIPAddressesGetIpsFilterProjectParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationIPAddressesGetIpsFilterServerParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationIPAddressesGetIpsFilterTypeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationIPAddressesGetIpsAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	apiVersionDescription := ``

	var apiVersionFlagName string
	if cmdPrefix == "" {
		apiVersionFlagName = "API-Version"
	} else {
		apiVersionFlagName = fmt.Sprintf("%v.API-Version", cmdPrefix)
	}

	var apiVersionFlagDefault string = "2023-06-01"

	_ = cmd.PersistentFlags().String(apiVersionFlagName, apiVersionFlagDefault, apiVersionDescription)

	return nil
}
func registerOperationIPAddressesGetIpsExtraFieldsIPAddressesParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	extraFieldsIpAddressesDescription := `The ` + "`" + `region` + "`" + ` and ` + "`" + `server` + "`" + ` are provided as extra attributes that is lazy loaded. To request it, just set ` + "`" + `extra_fields[ip_addresses]=region,server` + "`" + ` in the query string.`

	var extraFieldsIpAddressesFlagName string
	if cmdPrefix == "" {
		extraFieldsIpAddressesFlagName = "extra_fields[ip_addresses]"
	} else {
		extraFieldsIpAddressesFlagName = fmt.Sprintf("%v.extra_fields[ip_addresses]", cmdPrefix)
	}

	var extraFieldsIpAddressesFlagDefault string

	_ = cmd.PersistentFlags().String(extraFieldsIpAddressesFlagName, extraFieldsIpAddressesFlagDefault, extraFieldsIpAddressesDescription)

	return nil
}
func registerOperationIPAddressesGetIpsFilterAddressParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterAddressDescription := `The address of IP to filter by starts_with`

	var filterAddressFlagName string
	if cmdPrefix == "" {
		filterAddressFlagName = "filter[address]"
	} else {
		filterAddressFlagName = fmt.Sprintf("%v.filter[address]", cmdPrefix)
	}

	var filterAddressFlagDefault string

	_ = cmd.PersistentFlags().String(filterAddressFlagName, filterAddressFlagDefault, filterAddressDescription)

	return nil
}
func registerOperationIPAddressesGetIpsFilterFamilyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterFamilyDescription := `Enum: ["IPv4","IPv6"]. The protocol family to filter by`

	var filterFamilyFlagName string
	if cmdPrefix == "" {
		filterFamilyFlagName = "filter[family]"
	} else {
		filterFamilyFlagName = fmt.Sprintf("%v.filter[family]", cmdPrefix)
	}

	var filterFamilyFlagDefault string

	_ = cmd.PersistentFlags().String(filterFamilyFlagName, filterFamilyFlagDefault, filterFamilyDescription)

	if err := cmd.RegisterFlagCompletionFunc(filterFamilyFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["IPv4","IPv6"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}
func registerOperationIPAddressesGetIpsFilterLocationParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterLocationDescription := `The site slug to filter by`

	var filterLocationFlagName string
	if cmdPrefix == "" {
		filterLocationFlagName = "filter[location]"
	} else {
		filterLocationFlagName = fmt.Sprintf("%v.filter[location]", cmdPrefix)
	}

	var filterLocationFlagDefault string

	_ = cmd.PersistentFlags().String(filterLocationFlagName, filterLocationFlagDefault, filterLocationDescription)

	return nil
}
func registerOperationIPAddressesGetIpsFilterProjectParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterProjectDescription := `The project ID or Slug to filter by`

	var filterProjectFlagName string
	if cmdPrefix == "" {
		filterProjectFlagName = "filter[project]"
	} else {
		filterProjectFlagName = fmt.Sprintf("%v.filter[project]", cmdPrefix)
	}

	var filterProjectFlagDefault string

	_ = cmd.PersistentFlags().String(filterProjectFlagName, filterProjectFlagDefault, filterProjectDescription)

	return nil
}
func registerOperationIPAddressesGetIpsFilterServerParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterServerDescription := `The server ID to filter by`

	var filterServerFlagName string
	if cmdPrefix == "" {
		filterServerFlagName = "filter[server]"
	} else {
		filterServerFlagName = fmt.Sprintf("%v.filter[server]", cmdPrefix)
	}

	var filterServerFlagDefault string

	_ = cmd.PersistentFlags().String(filterServerFlagName, filterServerFlagDefault, filterServerDescription)

	return nil
}
func registerOperationIPAddressesGetIpsFilterTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterTypeDescription := `Enum: ["private","public"]. The protocol type to filter by`

	var filterTypeFlagName string
	if cmdPrefix == "" {
		filterTypeFlagName = "filter[type]"
	} else {
		filterTypeFlagName = fmt.Sprintf("%v.filter[type]", cmdPrefix)
	}

	var filterTypeFlagDefault string

	_ = cmd.PersistentFlags().String(filterTypeFlagName, filterTypeFlagDefault, filterTypeDescription)

	if err := cmd.RegisterFlagCompletionFunc(filterTypeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["private","public"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func retrieveOperationIPAddressesGetIpsAPIVersionFlag(m *ip_addresses.GetIpsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("API-Version") {

		var apiVersionFlagName string
		if cmdPrefix == "" {
			apiVersionFlagName = "API-Version"
		} else {
			apiVersionFlagName = fmt.Sprintf("%v.API-Version", cmdPrefix)
		}

		apiVersionFlagValue, err := cmd.Flags().GetString(apiVersionFlagName)
		if err != nil {
			return err, false
		}
		m.APIVersion = &apiVersionFlagValue

	}
	return nil, retAdded
}
func retrieveOperationIPAddressesGetIpsExtraFieldsIPAddressesFlag(m *ip_addresses.GetIpsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("extra_fields[ip_addresses]") {

		var extraFieldsIpAddressesFlagName string
		if cmdPrefix == "" {
			extraFieldsIpAddressesFlagName = "extra_fields[ip_addresses]"
		} else {
			extraFieldsIpAddressesFlagName = fmt.Sprintf("%v.extra_fields[ip_addresses]", cmdPrefix)
		}

		extraFieldsIpAddressesFlagValue, err := cmd.Flags().GetString(extraFieldsIpAddressesFlagName)
		if err != nil {
			return err, false
		}
		m.ExtraFieldsIPAddresses = &extraFieldsIpAddressesFlagValue

	}
	return nil, retAdded
}
func retrieveOperationIPAddressesGetIpsFilterAddressFlag(m *ip_addresses.GetIpsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[address]") {

		var filterAddressFlagName string
		if cmdPrefix == "" {
			filterAddressFlagName = "filter[address]"
		} else {
			filterAddressFlagName = fmt.Sprintf("%v.filter[address]", cmdPrefix)
		}

		filterAddressFlagValue, err := cmd.Flags().GetString(filterAddressFlagName)
		if err != nil {
			return err, false
		}
		m.FilterAddress = &filterAddressFlagValue

	}
	return nil, retAdded
}
func retrieveOperationIPAddressesGetIpsFilterFamilyFlag(m *ip_addresses.GetIpsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[family]") {

		var filterFamilyFlagName string
		if cmdPrefix == "" {
			filterFamilyFlagName = "filter[family]"
		} else {
			filterFamilyFlagName = fmt.Sprintf("%v.filter[family]", cmdPrefix)
		}

		filterFamilyFlagValue, err := cmd.Flags().GetString(filterFamilyFlagName)
		if err != nil {
			return err, false
		}
		m.FilterFamily = &filterFamilyFlagValue

	}
	return nil, retAdded
}
func retrieveOperationIPAddressesGetIpsFilterLocationFlag(m *ip_addresses.GetIpsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[location]") {

		var filterLocationFlagName string
		if cmdPrefix == "" {
			filterLocationFlagName = "filter[location]"
		} else {
			filterLocationFlagName = fmt.Sprintf("%v.filter[location]", cmdPrefix)
		}

		filterLocationFlagValue, err := cmd.Flags().GetString(filterLocationFlagName)
		if err != nil {
			return err, false
		}
		m.FilterLocation = &filterLocationFlagValue

	}
	return nil, retAdded
}
func retrieveOperationIPAddressesGetIpsFilterProjectFlag(m *ip_addresses.GetIpsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[project]") {

		var filterProjectFlagName string
		if cmdPrefix == "" {
			filterProjectFlagName = "filter[project]"
		} else {
			filterProjectFlagName = fmt.Sprintf("%v.filter[project]", cmdPrefix)
		}

		filterProjectFlagValue, err := cmd.Flags().GetString(filterProjectFlagName)
		if err != nil {
			return err, false
		}
		m.FilterProject = &filterProjectFlagValue

	}
	return nil, retAdded
}
func retrieveOperationIPAddressesGetIpsFilterServerFlag(m *ip_addresses.GetIpsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[server]") {

		var filterServerFlagName string
		if cmdPrefix == "" {
			filterServerFlagName = "filter[server]"
		} else {
			filterServerFlagName = fmt.Sprintf("%v.filter[server]", cmdPrefix)
		}

		filterServerFlagValue, err := cmd.Flags().GetString(filterServerFlagName)
		if err != nil {
			return err, false
		}
		m.FilterServer = &filterServerFlagValue

	}
	return nil, retAdded
}
func retrieveOperationIPAddressesGetIpsFilterTypeFlag(m *ip_addresses.GetIpsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[type]") {

		var filterTypeFlagName string
		if cmdPrefix == "" {
			filterTypeFlagName = "filter[type]"
		} else {
			filterTypeFlagName = fmt.Sprintf("%v.filter[type]", cmdPrefix)
		}

		filterTypeFlagValue, err := cmd.Flags().GetString(filterTypeFlagName)
		if err != nil {
			return err, false
		}
		m.FilterType = &filterTypeFlagValue

	}
	return nil, retAdded
}

// parseOperationIPAddressesGetIpsResult parses request result and return the string content
func parseOperationIPAddressesGetIpsResult(resp0 *ip_addresses.GetIpsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*ip_addresses.GetIpsOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*ip_addresses.GetIpsUnprocessableEntity)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
