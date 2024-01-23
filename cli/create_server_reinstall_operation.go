// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/server_reinstall"
	"github.com/latitudesh/cli/internal/api"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationServerReinstallCreateServerReinstallCmd returns a cmd to handle operation createServerReinstall
func makeOperationServerReinstallCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "reinstall",
		Short: `Submit a reinstall request to a server.`,
		RunE:  runOperationServerReinstallCreateServerReinstall,
	}

	if err := registerOperationServerReinstallCreateServerReinstallParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServerReinstallCreateServerReinstall uses cmd flags to call endpoint api
func runOperationServerReinstallCreateServerReinstall(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := server_reinstall.NewCreateServerReinstallParams()
	if err, _ := retrieveOperationServerReinstallCreateServerReinstallAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServerReinstallCreateServerReinstallBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServerReinstallCreateServerReinstallServerIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	result, err := appCli.ServerReinstall.CreateServerReinstall(params, nil)
	if err != nil {
		api.RenderErrorOutput(err)
		return nil
	}

	msgStr, err := parseOperationServerReinstallCreateServerReinstallResult(result)
	if err != nil {
		return err
	}
	if !debug {

		utils.PrintOutput(msgStr)
	}
	return nil
}

// registerOperationServerReinstallCreateServerReinstallParamFlags registers all flags needed to fill params
func registerOperationServerReinstallCreateServerReinstallParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServerReinstallCreateServerReinstallAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServerReinstallCreateServerReinstallBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServerReinstallCreateServerReinstallServerIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServerReinstallCreateServerReinstallAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationServerReinstallCreateServerReinstallBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelCreateServerReinstallBodyFlags(0, "", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationServerReinstallCreateServerReinstallServerIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	serverIdDescription := `Required. `

	var serverIdFlagName string
	if cmdPrefix == "" {
		serverIdFlagName = "server_id"
	} else {
		serverIdFlagName = fmt.Sprintf("%v.server_id", cmdPrefix)
	}

	var serverIdFlagDefault string

	_ = cmd.PersistentFlags().String(serverIdFlagName, serverIdFlagDefault, serverIdDescription)
	cmd.MarkPersistentFlagRequired(serverIdFlagName)

	return nil
}

func retrieveOperationServerReinstallCreateServerReinstallAPIVersionFlag(m *server_reinstall.CreateServerReinstallParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationServerReinstallCreateServerReinstallBodyFlag(m *server_reinstall.CreateServerReinstallParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := server_reinstall.CreateServerReinstallBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in CreateServerReinstallBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = server_reinstall.CreateServerReinstallBody{}
	}
	err, added := retrieveModelCreateServerReinstallBodyFlags(0, &bodyValueModel, "", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Body = bodyValueModel
	}
	if dryRun && debug {

		bodyValueDebugBytes, err := json.Marshal(m.Body)
		if err != nil {
			return err, false
		}
		logDebugf("Body dry-run payload: %v", string(bodyValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationServerReinstallCreateServerReinstallServerIDFlag(m *server_reinstall.CreateServerReinstallParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("server_id") {

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

	}
	return nil, retAdded
}

// parseOperationServerReinstallCreateServerReinstallResult parses request result and return the string content
func parseOperationServerReinstallCreateServerReinstallResult(resp0 *server_reinstall.CreateServerReinstallCreated) (string, error) {
	// warning: non schema response createServerReinstallCreated is not supported by go-swagger cli yet.

	return "", nil
}

// register flags to command
func registerModelCreateServerReinstallBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateServerReinstallBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateServerReinstallBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName = ""

	if err := registerModelCreateServerReinstallParamsBodyDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateServerReinstallBodyFlags(depth int, m *server_reinstall.CreateServerReinstallBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveCreateServerReinstallBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveCreateServerReinstallBodyDataFlags(depth int, m *server_reinstall.CreateServerReinstallBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%vdata", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data CreateServerReinstallParamsBodyData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &server_reinstall.CreateServerReinstallParamsBodyData{}
	}

	dataFlagName = ""
	err, dataAdded := retrieveModelCreateServerReinstallParamsBodyDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}

// register flags to command
func registerModelCreateServerReinstallParamsBodyDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateServerReinstallParamsBodyDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateServerReinstallParamsBodyDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateServerReinstallParamsBodyDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName = ""

	if err := registerModelCreateServerReinstallParamsBodyDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateServerReinstallParamsBodyDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["reinstalls"]. Required. `

	var typeFlagName = "type"

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)
	cmd.MarkPersistentFlagRequired(typeFlagName)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["reinstalls"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateServerReinstallParamsBodyDataFlags(depth int, m *server_reinstall.CreateServerReinstallParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveCreateServerReinstallParamsBodyDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, typeAdded := retrieveCreateServerReinstallParamsBodyDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveCreateServerReinstallParamsBodyDataAttributesFlags(depth int, m *server_reinstall.CreateServerReinstallParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%vattributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes CreateServerReinstallParamsBodyDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &server_reinstall.CreateServerReinstallParamsBodyDataAttributes{}
	}

	attributesFlagName = ""
	err, attributesAdded := retrieveModelCreateServerReinstallParamsBodyDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveCreateServerReinstallParamsBodyDataTypeFlags(depth int, m *server_reinstall.CreateServerReinstallParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	var typeFlagName = "type"
	if cmd.Flags().Changed(typeFlagName) {

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = &typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// register flags to command
func registerModelCreateServerReinstallParamsBodyDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateServerReinstallParamsBodyDataAttributesHostname(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateServerReinstallParamsBodyDataAttributesIpxeURL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateServerReinstallParamsBodyDataAttributesOperatingSystem(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateServerReinstallParamsBodyDataAttributesRaid(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateServerReinstallParamsBodyDataAttributesSSHKeys(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateServerReinstallParamsBodyDataAttributesUserData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateServerReinstallParamsBodyDataAttributesHostname(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	hostnameDescription := `The server hostname to set upon reinstall`

	var hostnameFlagName = "hostname"

	var hostnameFlagDefault string

	_ = cmd.PersistentFlags().String(hostnameFlagName, hostnameFlagDefault, hostnameDescription)

	return nil
}

func registerCreateServerReinstallParamsBodyDataAttributesIpxeURL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipxeUrlDescription := `URL where iPXE script is stored on, necessary for custom image reinstalls. This attribute is required when operating system iPXE is selected.`

	var ipxeUrlFlagName string
	if cmdPrefix == "" {
		ipxeUrlFlagName = "ipxe_url"
	} else {
		ipxeUrlFlagName = fmt.Sprintf("%v.ipxe_url", cmdPrefix)
	}

	var ipxeUrlFlagDefault string

	_ = cmd.PersistentFlags().String(ipxeUrlFlagName, ipxeUrlFlagDefault, ipxeUrlDescription)

	return nil
}

func registerCreateServerReinstallParamsBodyDataAttributesOperatingSystem(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	operatingSystemDescription := `Enum: ["ipxe","windows_server_2019_std_v1","ubuntu_22_04_x64_lts","debian_11","rockylinux_8","debian_10","rhel8","centos_7_4_x64","centos_8_x64","ubuntu_20_04_x64_lts","debian_12","ubuntu22_ml_in_a_box","windows2022"]. The OS selected for the reinstall process`

	var operatingSystemFlagName = "operating_system"

	var operatingSystemFlagDefault string

	_ = cmd.PersistentFlags().String(operatingSystemFlagName, operatingSystemFlagDefault, operatingSystemDescription)

	if err := cmd.RegisterFlagCompletionFunc(operatingSystemFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["ipxe","windows_server_2019_std_v1","ubuntu_22_04_x64_lts","debian_11","rockylinux_8","debian_10","rhel8","centos_7_4_x64","centos_8_x64","ubuntu_20_04_x64_lts","debian_12","ubuntu22_ml_in_a_box","windows2022"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerCreateServerReinstallParamsBodyDataAttributesRaid(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	raidDescription := `Enum: ["raid-0","raid-1"]. RAID mode for the server`

	var raidFlagName = "raid"

	var raidFlagDefault string

	_ = cmd.PersistentFlags().String(raidFlagName, raidFlagDefault, raidDescription)

	if err := cmd.RegisterFlagCompletionFunc(raidFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["raid-0","raid-1"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerCreateServerReinstallParamsBodyDataAttributesSSHKeys(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ssh_keys []string array type is not supported by go-swagger cli yet

	return nil
}

func registerCreateServerReinstallParamsBodyDataAttributesUserData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	userDataDescription := `User data to set upon reinstall`

	var userDataFlagName = "user_data"

	var userDataFlagDefault int64

	_ = cmd.PersistentFlags().Int64(userDataFlagName, userDataFlagDefault, userDataDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateServerReinstallParamsBodyDataAttributesFlags(depth int, m *server_reinstall.CreateServerReinstallParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, hostnameAdded := retrieveCreateServerReinstallParamsBodyDataAttributesHostnameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hostnameAdded

	err, ipxeUrlAdded := retrieveCreateServerReinstallParamsBodyDataAttributesIpxeURLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipxeUrlAdded

	err, operatingSystemAdded := retrieveCreateServerReinstallParamsBodyDataAttributesOperatingSystemFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || operatingSystemAdded

	err, raidAdded := retrieveCreateServerReinstallParamsBodyDataAttributesRaidFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || raidAdded

	err, sshKeysAdded := retrieveCreateServerReinstallParamsBodyDataAttributesSSHKeysFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sshKeysAdded

	err, userDataAdded := retrieveCreateServerReinstallParamsBodyDataAttributesUserDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || userDataAdded

	return nil, retAdded
}

func retrieveCreateServerReinstallParamsBodyDataAttributesHostnameFlags(depth int, m *server_reinstall.CreateServerReinstallParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	var hostnameFlagName = "hostname"
	if cmd.Flags().Changed(hostnameFlagName) {

		hostnameFlagValue, err := cmd.Flags().GetString(hostnameFlagName)
		if err != nil {
			return err, false
		}
		m.Hostname = hostnameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateServerReinstallParamsBodyDataAttributesIpxeURLFlags(depth int, m *server_reinstall.CreateServerReinstallParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	var ipxeUrlFlagName = "ipxe_url"
	if cmd.Flags().Changed(ipxeUrlFlagName) {

		ipxeUrlFlagValue, err := cmd.Flags().GetString(ipxeUrlFlagName)
		if err != nil {
			return err, false
		}
		m.IpxeURL = &ipxeUrlFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateServerReinstallParamsBodyDataAttributesOperatingSystemFlags(depth int, m *server_reinstall.CreateServerReinstallParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	operatingSystemFlagName := fmt.Sprintf("%v.operating_system", cmdPrefix)
	if cmd.Flags().Changed(operatingSystemFlagName) {

		var operatingSystemFlagName  = "operating_system"

		operatingSystemFlagValue, err := cmd.Flags().GetString(operatingSystemFlagName)
		if err != nil {
			return err, false
		}
		m.OperatingSystem = operatingSystemFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateServerReinstallParamsBodyDataAttributesRaidFlags(depth int, m *server_reinstall.CreateServerReinstallParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	var raidFlagName = "raid"
	if cmd.Flags().Changed(raidFlagName) {

		raidFlagValue, err := cmd.Flags().GetString(raidFlagName)
		if err != nil {
			return err, false
		}
		m.Raid = &raidFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateServerReinstallParamsBodyDataAttributesSSHKeysFlags(depth int, m *server_reinstall.CreateServerReinstallParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	var sshKeysFlagName = "ssh_keys"
	if cmd.Flags().Changed(sshKeysFlagName) {
		// warning: ssh_keys array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveCreateServerReinstallParamsBodyDataAttributesUserDataFlags(depth int, m *server_reinstall.CreateServerReinstallParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	userDataFlagName := fmt.Sprintf("%v.user_data", cmdPrefix)
	if cmd.Flags().Changed(userDataFlagName) {

		var userDataFlagName = "user_data"

		userDataFlagValue, err := cmd.Flags().GetInt64(userDataFlagName)
		if err != nil {
			return err, false
		}
		m.UserData = &userDataFlagValue

		retAdded = true
	}

	return nil, retAdded
}
