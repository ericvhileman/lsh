// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/servers"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationServersCreateServerCmd returns a cmd to handle operation createServer
func makeOperationServersCreateServerCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "create-server",
		Short: ``,
		RunE:  runOperationServersCreateServer,
	}

	if err := registerOperationServersCreateServerParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServersCreateServer uses cmd flags to call endpoint api
func runOperationServersCreateServer(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := servers.NewCreateServerParams()
	if err, _ := retrieveOperationServersCreateServerAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersCreateServerBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationServersCreateServerResult(appCli.Servers.CreateServer(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(utils.PrettifyJson(msgStr))
	}
	return nil
}

// registerOperationServersCreateServerParamFlags registers all flags needed to fill params
func registerOperationServersCreateServerParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServersCreateServerAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersCreateServerBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServersCreateServerAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationServersCreateServerBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelCreateServerBodyFlags(0, "createServerBody", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationServersCreateServerAPIVersionFlag(m *servers.CreateServerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationServersCreateServerBodyFlag(m *servers.CreateServerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := servers.CreateServerBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in CreateServerBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = servers.CreateServerBody{}
	}
	err, added := retrieveModelCreateServerBodyFlags(0, &bodyValueModel, "createServerBody", cmd)
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

// parseOperationServersCreateServerResult parses request result and return the string content
func parseOperationServersCreateServerResult(resp0 *servers.CreateServerCreated, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*servers.CreateServerCreated)
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
		resp1, ok := iResp1.(*servers.CreateServerBadRequest)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*servers.CreateServerUnprocessableEntity)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
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

// register flags to command
func registerModelCreateServerBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateServerBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateServerBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelCreateServerParamsBodyDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateServerBodyFlags(depth int, m *servers.CreateServerBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveCreateServerBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveCreateServerBodyDataFlags(depth int, m *servers.CreateServerBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data CreateServerParamsBodyData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &servers.CreateServerParamsBodyData{}
	}

	err, dataAdded := retrieveModelCreateServerParamsBodyDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
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
func registerModelCreateServerParamsBodyDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateServerParamsBodyDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateServerParamsBodyDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateServerParamsBodyDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelCreateServerParamsBodyDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateServerParamsBodyDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["servers"]. Required. `

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
			if err := json.Unmarshal([]byte(`["servers"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateServerParamsBodyDataFlags(depth int, m *servers.CreateServerParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveCreateServerParamsBodyDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, typeAdded := retrieveCreateServerParamsBodyDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveCreateServerParamsBodyDataAttributesFlags(depth int, m *servers.CreateServerParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes CreateServerParamsBodyDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &servers.CreateServerParamsBodyDataAttributes{}
	}

	err, attributesAdded := retrieveModelCreateServerParamsBodyDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveCreateServerParamsBodyDataTypeFlags(depth int, m *servers.CreateServerParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Type = &typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// register flags to command
func registerModelCreateServerParamsBodyDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateServerParamsBodyDataAttributesBilling(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateServerParamsBodyDataAttributesHostname(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateServerParamsBodyDataAttributesIpxeURL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateServerParamsBodyDataAttributesOperatingSystem(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateServerParamsBodyDataAttributesPlan(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateServerParamsBodyDataAttributesProject(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateServerParamsBodyDataAttributesRaid(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateServerParamsBodyDataAttributesSite(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateServerParamsBodyDataAttributesSSHKeys(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateServerParamsBodyDataAttributesUserData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateServerParamsBodyDataAttributesBilling(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	billingDescription := `Enum: ["hourly","monthly","yearly"]. The server billing type. Accepts ` + "`" + `hourly` + "`" + ` and ` + "`" + `monthly` + "`" + ` for on demand projects and ` + "`" + `yearly` + "`" + ` for reserved projects.`

	var billingFlagName string
	if cmdPrefix == "" {
		billingFlagName = "billing"
	} else {
		billingFlagName = fmt.Sprintf("%v.billing", cmdPrefix)
	}

	var billingFlagDefault string

	_ = cmd.PersistentFlags().String(billingFlagName, billingFlagDefault, billingDescription)

	if err := cmd.RegisterFlagCompletionFunc(billingFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["hourly","monthly","yearly"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerCreateServerParamsBodyDataAttributesHostname(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	hostnameDescription := `The server hostname`

	var hostnameFlagName string
	if cmdPrefix == "" {
		hostnameFlagName = "hostname"
	} else {
		hostnameFlagName = fmt.Sprintf("%v.hostname", cmdPrefix)
	}

	var hostnameFlagDefault string

	_ = cmd.PersistentFlags().String(hostnameFlagName, hostnameFlagDefault, hostnameDescription)

	return nil
}

func registerCreateServerParamsBodyDataAttributesIpxeURL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipxeUrlDescription := `URL where iPXE script is stored on, necessary for custom image deployments.This attribute is required when iPXE is selected as operating system.`

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

func registerCreateServerParamsBodyDataAttributesOperatingSystem(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	operatingSystemDescription := `Enum: ["ipxe","windows_server_2019_std_v1","ubuntu_22_04_x64_lts","debian_11","rockylinux_8","debian_10","rhel8","centos_7_4_x64","centos_8_x64","ubuntu_20_04_x64_lts","debian_12","ubuntu22_ml_in_a_box","windows2022"]. The operating system for the new server`

	var operatingSystemFlagName string
	if cmdPrefix == "" {
		operatingSystemFlagName = "operating_system"
	} else {
		operatingSystemFlagName = fmt.Sprintf("%v.operating_system", cmdPrefix)
	}

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

func registerCreateServerParamsBodyDataAttributesPlan(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	planDescription := `Enum: ["c2-large-x86","c2-medium-x86","c2-small-x86","c3-large-x86","c3-medium-x86","c3-small-x86","c3-xlarge-x86","g3-large-x86","g3-medium-x86","g3-small-x86","g3-xlarge-x86","m3-large-x86","s2-small-x86","s3-large-x86"]. The plan to choose server from`

	var planFlagName string
	if cmdPrefix == "" {
		planFlagName = "plan"
	} else {
		planFlagName = fmt.Sprintf("%v.plan", cmdPrefix)
	}

	var planFlagDefault string

	_ = cmd.PersistentFlags().String(planFlagName, planFlagDefault, planDescription)

	if err := cmd.RegisterFlagCompletionFunc(planFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["c2-large-x86","c2-medium-x86","c2-small-x86","c3-large-x86","c3-medium-x86","c3-small-x86","c3-xlarge-x86","g3-large-x86","g3-medium-x86","g3-small-x86","g3-xlarge-x86","m3-large-x86","s2-small-x86","s3-large-x86"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerCreateServerParamsBodyDataAttributesProject(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	projectDescription := `The project (ID or Slug) to deploy the server`

	var projectFlagName string
	if cmdPrefix == "" {
		projectFlagName = "project"
	} else {
		projectFlagName = fmt.Sprintf("%v.project", cmdPrefix)
	}

	var projectFlagDefault string

	_ = cmd.PersistentFlags().String(projectFlagName, projectFlagDefault, projectDescription)

	return nil
}

func registerCreateServerParamsBodyDataAttributesRaid(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	raidDescription := `Enum: ["raid-0","raid-1"]. RAID mode for the server`

	var raidFlagName string
	if cmdPrefix == "" {
		raidFlagName = "raid"
	} else {
		raidFlagName = fmt.Sprintf("%v.raid", cmdPrefix)
	}

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

func registerCreateServerParamsBodyDataAttributesSite(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	siteDescription := `Enum: ["ASH","BGT","BUE","CHI","DAL","FRA","LAX","LON","MEX","MEX2","MIA","MIA2","NYC","SAN","SAN2","SAO","SAO2","SYD","TYO","TYO2"]. The site to deploy the server`

	var siteFlagName string
	if cmdPrefix == "" {
		siteFlagName = "site"
	} else {
		siteFlagName = fmt.Sprintf("%v.site", cmdPrefix)
	}

	var siteFlagDefault string

	_ = cmd.PersistentFlags().String(siteFlagName, siteFlagDefault, siteDescription)

	if err := cmd.RegisterFlagCompletionFunc(siteFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["ASH","BGT","BUE","CHI","DAL","FRA","LAX","LON","MEX","MEX2","MIA","MIA2","NYC","SAN","SAN2","SAO","SAO2","SYD","TYO","TYO2"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerCreateServerParamsBodyDataAttributesSSHKeys(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ssh_keys []string array type is not supported by go-swagger cli yet

	return nil
}

func registerCreateServerParamsBodyDataAttributesUserData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	userDataDescription := `User data to set on the server`

	var userDataFlagName string
	if cmdPrefix == "" {
		userDataFlagName = "user_data"
	} else {
		userDataFlagName = fmt.Sprintf("%v.user_data", cmdPrefix)
	}

	var userDataFlagDefault int64

	_ = cmd.PersistentFlags().Int64(userDataFlagName, userDataFlagDefault, userDataDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateServerParamsBodyDataAttributesFlags(depth int, m *servers.CreateServerParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, billingAdded := retrieveCreateServerParamsBodyDataAttributesBillingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || billingAdded

	err, hostnameAdded := retrieveCreateServerParamsBodyDataAttributesHostnameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hostnameAdded

	err, ipxeUrlAdded := retrieveCreateServerParamsBodyDataAttributesIpxeURLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipxeUrlAdded

	err, operatingSystemAdded := retrieveCreateServerParamsBodyDataAttributesOperatingSystemFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || operatingSystemAdded

	err, planAdded := retrieveCreateServerParamsBodyDataAttributesPlanFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || planAdded

	err, projectAdded := retrieveCreateServerParamsBodyDataAttributesProjectFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || projectAdded

	err, raidAdded := retrieveCreateServerParamsBodyDataAttributesRaidFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || raidAdded

	err, siteAdded := retrieveCreateServerParamsBodyDataAttributesSiteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || siteAdded

	err, sshKeysAdded := retrieveCreateServerParamsBodyDataAttributesSSHKeysFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sshKeysAdded

	err, userDataAdded := retrieveCreateServerParamsBodyDataAttributesUserDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || userDataAdded

	return nil, retAdded
}

func retrieveCreateServerParamsBodyDataAttributesBillingFlags(depth int, m *servers.CreateServerParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	billingFlagName := fmt.Sprintf("%v.billing", cmdPrefix)
	if cmd.Flags().Changed(billingFlagName) {

		var billingFlagName string
		if cmdPrefix == "" {
			billingFlagName = "billing"
		} else {
			billingFlagName = fmt.Sprintf("%v.billing", cmdPrefix)
		}

		billingFlagValue, err := cmd.Flags().GetString(billingFlagName)
		if err != nil {
			return err, false
		}
		m.Billing = &billingFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateServerParamsBodyDataAttributesHostnameFlags(depth int, m *servers.CreateServerParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	hostnameFlagName := fmt.Sprintf("%v.hostname", cmdPrefix)
	if cmd.Flags().Changed(hostnameFlagName) {

		var hostnameFlagName string
		if cmdPrefix == "" {
			hostnameFlagName = "hostname"
		} else {
			hostnameFlagName = fmt.Sprintf("%v.hostname", cmdPrefix)
		}

		hostnameFlagValue, err := cmd.Flags().GetString(hostnameFlagName)
		if err != nil {
			return err, false
		}
		m.Hostname = hostnameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateServerParamsBodyDataAttributesIpxeURLFlags(depth int, m *servers.CreateServerParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ipxeUrlFlagName := fmt.Sprintf("%v.ipxe_url", cmdPrefix)
	if cmd.Flags().Changed(ipxeUrlFlagName) {

		var ipxeUrlFlagName string
		if cmdPrefix == "" {
			ipxeUrlFlagName = "ipxe_url"
		} else {
			ipxeUrlFlagName = fmt.Sprintf("%v.ipxe_url", cmdPrefix)
		}

		ipxeUrlFlagValue, err := cmd.Flags().GetString(ipxeUrlFlagName)
		if err != nil {
			return err, false
		}
		m.IpxeURL = &ipxeUrlFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateServerParamsBodyDataAttributesOperatingSystemFlags(depth int, m *servers.CreateServerParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	operatingSystemFlagName := fmt.Sprintf("%v.operating_system", cmdPrefix)
	if cmd.Flags().Changed(operatingSystemFlagName) {

		var operatingSystemFlagName string
		if cmdPrefix == "" {
			operatingSystemFlagName = "operating_system"
		} else {
			operatingSystemFlagName = fmt.Sprintf("%v.operating_system", cmdPrefix)
		}

		operatingSystemFlagValue, err := cmd.Flags().GetString(operatingSystemFlagName)
		if err != nil {
			return err, false
		}
		m.OperatingSystem = operatingSystemFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateServerParamsBodyDataAttributesPlanFlags(depth int, m *servers.CreateServerParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	planFlagName := fmt.Sprintf("%v.plan", cmdPrefix)
	if cmd.Flags().Changed(planFlagName) {

		var planFlagName string
		if cmdPrefix == "" {
			planFlagName = "plan"
		} else {
			planFlagName = fmt.Sprintf("%v.plan", cmdPrefix)
		}

		planFlagValue, err := cmd.Flags().GetString(planFlagName)
		if err != nil {
			return err, false
		}
		m.Plan = planFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateServerParamsBodyDataAttributesProjectFlags(depth int, m *servers.CreateServerParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	projectFlagName := fmt.Sprintf("%v.project", cmdPrefix)
	if cmd.Flags().Changed(projectFlagName) {

		var projectFlagName string
		if cmdPrefix == "" {
			projectFlagName = "project"
		} else {
			projectFlagName = fmt.Sprintf("%v.project", cmdPrefix)
		}

		projectFlagValue, err := cmd.Flags().GetString(projectFlagName)
		if err != nil {
			return err, false
		}
		m.Project = projectFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateServerParamsBodyDataAttributesRaidFlags(depth int, m *servers.CreateServerParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	raidFlagName := fmt.Sprintf("%v.raid", cmdPrefix)
	if cmd.Flags().Changed(raidFlagName) {

		var raidFlagName string
		if cmdPrefix == "" {
			raidFlagName = "raid"
		} else {
			raidFlagName = fmt.Sprintf("%v.raid", cmdPrefix)
		}

		raidFlagValue, err := cmd.Flags().GetString(raidFlagName)
		if err != nil {
			return err, false
		}
		m.Raid = &raidFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateServerParamsBodyDataAttributesSiteFlags(depth int, m *servers.CreateServerParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	siteFlagName := fmt.Sprintf("%v.site", cmdPrefix)
	if cmd.Flags().Changed(siteFlagName) {

		var siteFlagName string
		if cmdPrefix == "" {
			siteFlagName = "site"
		} else {
			siteFlagName = fmt.Sprintf("%v.site", cmdPrefix)
		}

		siteFlagValue, err := cmd.Flags().GetString(siteFlagName)
		if err != nil {
			return err, false
		}
		m.Site = siteFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateServerParamsBodyDataAttributesSSHKeysFlags(depth int, m *servers.CreateServerParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sshKeysFlagName := fmt.Sprintf("%v.ssh_keys", cmdPrefix)
	if cmd.Flags().Changed(sshKeysFlagName) {
		// warning: ssh_keys array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveCreateServerParamsBodyDataAttributesUserDataFlags(depth int, m *servers.CreateServerParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	userDataFlagName := fmt.Sprintf("%v.user_data", cmdPrefix)
	if cmd.Flags().Changed(userDataFlagName) {

		var userDataFlagName string
		if cmdPrefix == "" {
			userDataFlagName = "user_data"
		} else {
			userDataFlagName = fmt.Sprintf("%v.user_data", cmdPrefix)
		}

		userDataFlagValue, err := cmd.Flags().GetInt64(userDataFlagName)
		if err != nil {
			return err, false
		}
		m.UserData = &userDataFlagValue

		retAdded = true
	}

	return nil, retAdded
}