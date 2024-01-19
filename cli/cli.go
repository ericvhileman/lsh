// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/latitudesh/cli/client"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// debug flag indicating that cli should output debug logs
var debug bool

// config file location
var configFile string

// dry run flag
var dryRun bool

// name of the executable
var exeName string = filepath.Base(os.Args[0])

// logDebugf writes debug log to stdout
func logDebugf(format string, v ...interface{}) {
	if !debug {
		return
	}
	log.Printf(format, v...)
}

// depth of recursion to construct model flags
var maxDepth int = 5

// makeClient constructs a client object
func makeClient(cmd *cobra.Command, args []string) (*client.LatitudeShAPI, error) {
	hostname := viper.GetString("hostname")
	viper.SetDefault("base_path", client.DefaultBasePath)
	basePath := viper.GetString("base_path")
	scheme := viper.GetString("scheme")

	r := httptransport.New(hostname, basePath, []string{scheme})
	r.SetDebug(debug)
	// set custom producer and consumer to use the default ones

	r.Consumers["application/json"] = runtime.JSONConsumer()
	r.Consumers["application/vnd.api+json"] = runtime.JSONConsumer()

	r.Producers["application/json"] = runtime.JSONProducer()

	auth, err := makeAuthInfoWriter(cmd)
	if err != nil {
		return nil, err
	}
	r.DefaultAuthentication = auth

	appCli := client.New(r, strfmt.Default)
	logDebugf("Server url: %v://%v", scheme, hostname)
	return appCli, nil
}

// MakeRootCmd returns the root cmd
func MakeRootCmd() (*cobra.Command, error) {
	cobra.OnInitialize(initViperConfigs)

	// Use executable name as the command name
	rootCmd := &cobra.Command{
		Use: exeName,
	}

	// register basic flags
	rootCmd.PersistentFlags().String("hostname", client.DefaultHost, "hostname of the service")
	viper.BindPFlag("hostname", rootCmd.PersistentFlags().Lookup("hostname"))
	rootCmd.PersistentFlags().String("scheme", client.DefaultSchemes[0], fmt.Sprintf("Choose from: %v", client.DefaultSchemes))
	viper.BindPFlag("scheme", rootCmd.PersistentFlags().Lookup("scheme"))
	rootCmd.PersistentFlags().String("base-path", client.DefaultBasePath, fmt.Sprintf("For example: %v", client.DefaultBasePath))
	viper.BindPFlag("base_path", rootCmd.PersistentFlags().Lookup("base-path"))

	// configure debug flag
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "output debug logs")
	// configure config location
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file path")
	// configure dry run flag
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "do not send the request to server")

	// register security flags
	if err := registerAuthInoWriterFlags(rootCmd); err != nil {
		return nil, err
	}

	// add login with api -oken
	operationLoginCmd, err := makeOperationLoginCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationLoginCmd)

	// add all operation groups
	operationGroupAccountCmd, err := makeOperationGroupAccountCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupAccountCmd)

	operationGroupAPIKeysCmd, err := makeOperationGroupAPIKeysCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupAPIKeysCmd)

	operationGroupBandwidthCmd, err := makeOperationGroupBandwidthCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupBandwidthCmd)

	operationGroupBandwidthPackagesCmd, err := makeOperationGroupBandwidthPackagesCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupBandwidthPackagesCmd)

	operationGroupBillingUsageCmd, err := makeOperationGroupBillingUsageCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupBillingUsageCmd)

	operationGroupDeployConfigCmd, err := makeOperationGroupDeployConfigCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupDeployConfigCmd)

	operationGroupEventsCmd, err := makeOperationGroupEventsCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupEventsCmd)

	operationGroupIPAddressesCmd, err := makeOperationGroupIPAddressesCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupIPAddressesCmd)

	operationGroupIPmiCredentialsCmd, err := makeOperationGroupIPmiCredentialsCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupIPmiCredentialsCmd)

	operationGroupMembersCmd, err := makeOperationGroupMembersCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupMembersCmd)

	operationGroupOperatingSystemsCmd, err := makeOperationGroupOperatingSystemsCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupOperatingSystemsCmd)

	operationGroupOutOfBandCmd, err := makeOperationGroupOutOfBandCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupOutOfBandCmd)

	operationGroupPlansCmd, err := makeOperationGroupPlansCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupPlansCmd)

	operationGroupPowerActionsCmd, err := makeOperationGroupPowerActionsCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupPowerActionsCmd)

	operationGroupProjectsCmd, err := makeOperationGroupProjectsCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupProjectsCmd)

	operationGroupRegionsCmd, err := makeOperationGroupRegionsCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupRegionsCmd)

	operationGroupRescueModeCmd, err := makeOperationGroupRescueModeCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupRescueModeCmd)

	operationGroupRolesCmd, err := makeOperationGroupRolesCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupRolesCmd)

	operationGroupServerReinstallCmd, err := makeOperationGroupServerReinstallCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupServerReinstallCmd)

	operationGroupServersCmd, err := makeOperationGroupServersCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupServersCmd)

	operationGroupSSHKeysCmd, err := makeOperationGroupSSHKeysCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupSSHKeysCmd)

	operationGroupTeamsCmd, err := makeOperationGroupTeamsCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupTeamsCmd)

	operationGroupUserDataCmd, err := makeOperationGroupUserDataCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupUserDataCmd)

	operationGroupVpnSessionsCmd, err := makeOperationGroupVpnSessionsCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupVpnSessionsCmd)

	operationGroupVirtualNetworkAssignmentsCmd, err := makeOperationGroupVirtualNetworkAssignmentsCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupVirtualNetworkAssignmentsCmd)

	operationGroupVirtualNetworksCmd, err := makeOperationGroupVirtualNetworksCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupVirtualNetworksCmd)

	// add cobra completion
	rootCmd.AddCommand(makeGenCompletionCmd())

	return rootCmd, nil
}

// initViperConfigs initialize viper config using config file in '$HOME/.config/<cli name>/config.<json|yaml...>'
// currently hostname, scheme and auth tokens can be specified in this config file.
func initViperConfigs() {
	if configFile != "" {
		// use user specified config file location
		viper.SetConfigFile(configFile)
	} else {
		// look for default config
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(path.Join(home, ".config", exeName))
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		logDebugf("Error: loading config file: %v", err)
		return
	}
	logDebugf("Using config file: %v", viper.ConfigFileUsed())
}

// registerAuthInoWriterFlags registers all flags needed to perform authentication
func registerAuthInoWriterFlags(cmd *cobra.Command) error {
	/*Authorization */
	cmd.PersistentFlags().String("Authorization", "", ``)
	viper.BindPFlag("Authorization", cmd.PersistentFlags().Lookup("Authorization"))
	return nil
}

// makeAuthInfoWriter retrieves cmd flags and construct an auth info writer
func makeAuthInfoWriter(cmd *cobra.Command) (runtime.ClientAuthInfoWriter, error) {
	auths := []runtime.ClientAuthInfoWriter{}
	/*Authorization */
	if viper.IsSet("Authorization") {
		AuthorizationKey := viper.GetString("Authorization")
		ApiVersion := viper.GetString("api-version")
		auths = append(auths, httptransport.APIKeyAuth("Authorization", "header", AuthorizationKey))
		auths = append(auths, httptransport.APIKeyAuth("API-Version", "header", ApiVersion))
	}
	if len(auths) == 0 {
		logDebugf("Warning: No auth params detected.")
		return nil, nil
	}
	// compose all auths together
	return httptransport.Compose(auths...), nil
}

func makeOperationGroupAccountCmd() (*cobra.Command, error) {
	operationGroupAccountCmd := &cobra.Command{
		Use:  "account",
		Long: ``,
	}

	operationGetUserProfileCmd, err := makeOperationAccountGetUserProfileCmd()
	if err != nil {
		return nil, err
	}
	operationGroupAccountCmd.AddCommand(operationGetUserProfileCmd)

	operationGetUserTeamsCmd, err := makeOperationAccountGetUserTeamsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupAccountCmd.AddCommand(operationGetUserTeamsCmd)

	operationPatchUserProfileCmd, err := makeOperationAccountPatchUserProfileCmd()
	if err != nil {
		return nil, err
	}
	operationGroupAccountCmd.AddCommand(operationPatchUserProfileCmd)

	return operationGroupAccountCmd, nil
}
func makeOperationGroupAPIKeysCmd() (*cobra.Command, error) {
	operationGroupAPIKeysCmd := &cobra.Command{
		Use:  "api_keys",
		Long: ``,
	}

	operationDeleteAPIKeyCmd, err := makeOperationAPIKeysDeleteAPIKeyCmd()
	if err != nil {
		return nil, err
	}
	operationGroupAPIKeysCmd.AddCommand(operationDeleteAPIKeyCmd)

	operationGetAPIKeysCmd, err := makeOperationAPIKeysGetAPIKeysCmd()
	if err != nil {
		return nil, err
	}
	operationGroupAPIKeysCmd.AddCommand(operationGetAPIKeysCmd)

	operationPostAPIKeyCmd, err := makeOperationAPIKeysPostAPIKeyCmd()
	if err != nil {
		return nil, err
	}
	operationGroupAPIKeysCmd.AddCommand(operationPostAPIKeyCmd)

	operationUpdateAPIKeyCmd, err := makeOperationAPIKeysUpdateAPIKeyCmd()
	if err != nil {
		return nil, err
	}
	operationGroupAPIKeysCmd.AddCommand(operationUpdateAPIKeyCmd)

	return operationGroupAPIKeysCmd, nil
}
func makeOperationGroupBandwidthCmd() (*cobra.Command, error) {
	operationGroupBandwidthCmd := &cobra.Command{
		Use:  "bandwidth",
		Long: ``,
	}

	operationGetTrafficConsumptionCmd, err := makeOperationBandwidthGetTrafficConsumptionCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBandwidthCmd.AddCommand(operationGetTrafficConsumptionCmd)

	operationGetTrafficQuotaCmd, err := makeOperationBandwidthGetTrafficQuotaCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBandwidthCmd.AddCommand(operationGetTrafficQuotaCmd)

	return operationGroupBandwidthCmd, nil
}
func makeOperationGroupBandwidthPackagesCmd() (*cobra.Command, error) {
	operationGroupBandwidthPackagesCmd := &cobra.Command{
		Use:  "bandwidth_packages",
		Long: ``,
	}

	operationUpdatePlansBandwidthCmd, err := makeOperationBandwidthPackagesUpdatePlansBandwidthCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBandwidthPackagesCmd.AddCommand(operationUpdatePlansBandwidthCmd)

	return operationGroupBandwidthPackagesCmd, nil
}
func makeOperationGroupBillingUsageCmd() (*cobra.Command, error) {
	operationGroupBillingUsageCmd := &cobra.Command{
		Use:  "billing_usage",
		Long: ``,
	}

	operationGetBillingUsageCmd, err := makeOperationBillingUsageGetBillingUsageCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBillingUsageCmd.AddCommand(operationGetBillingUsageCmd)

	return operationGroupBillingUsageCmd, nil
}
func makeOperationGroupDeployConfigCmd() (*cobra.Command, error) {
	operationGroupDeployConfigCmd := &cobra.Command{
		Use:  "deploy_config",
		Long: ``,
	}

	operationGetServerDeployConfigCmd, err := makeOperationDeployConfigGetServerDeployConfigCmd()
	if err != nil {
		return nil, err
	}
	operationGroupDeployConfigCmd.AddCommand(operationGetServerDeployConfigCmd)

	operationUpdateServerDeployConfigCmd, err := makeOperationDeployConfigUpdateServerDeployConfigCmd()
	if err != nil {
		return nil, err
	}
	operationGroupDeployConfigCmd.AddCommand(operationUpdateServerDeployConfigCmd)

	return operationGroupDeployConfigCmd, nil
}
func makeOperationGroupEventsCmd() (*cobra.Command, error) {
	operationGroupEventsCmd := &cobra.Command{
		Use:  "events",
		Long: ``,
	}

	operationGetEventsCmd, err := makeOperationEventsGetEventsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupEventsCmd.AddCommand(operationGetEventsCmd)

	return operationGroupEventsCmd, nil
}
func makeOperationGroupIPAddressesCmd() (*cobra.Command, error) {
	operationGroupIPAddressesCmd := &cobra.Command{
		Use:  "ip_addresses",
		Long: ``,
	}

	operationGetIPCmd, err := makeOperationIPAddressesGetIPCmd()
	if err != nil {
		return nil, err
	}
	operationGroupIPAddressesCmd.AddCommand(operationGetIPCmd)

	operationGetIpsCmd, err := makeOperationIPAddressesGetIpsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupIPAddressesCmd.AddCommand(operationGetIpsCmd)

	return operationGroupIPAddressesCmd, nil
}
func makeOperationGroupIPmiCredentialsCmd() (*cobra.Command, error) {
	operationGroupIPmiCredentialsCmd := &cobra.Command{
		Use:  "ip_m_i_credentials",
		Long: ``,
	}

	operationCreateIpmiSessionCmd, err := makeOperationIPmiCredentialsCreateIpmiSessionCmd()
	if err != nil {
		return nil, err
	}
	operationGroupIPmiCredentialsCmd.AddCommand(operationCreateIpmiSessionCmd)

	return operationGroupIPmiCredentialsCmd, nil
}
func makeOperationGroupMembersCmd() (*cobra.Command, error) {
	operationGroupMembersCmd := &cobra.Command{
		Use:  "members",
		Long: ``,
	}

	operationDestroyTeamMemberCmd, err := makeOperationMembersDestroyTeamMemberCmd()
	if err != nil {
		return nil, err
	}
	operationGroupMembersCmd.AddCommand(operationDestroyTeamMemberCmd)

	operationGetTeamMembersCmd, err := makeOperationMembersGetTeamMembersCmd()
	if err != nil {
		return nil, err
	}
	operationGroupMembersCmd.AddCommand(operationGetTeamMembersCmd)

	operationPostTeamMembersCmd, err := makeOperationMembersPostTeamMembersCmd()
	if err != nil {
		return nil, err
	}
	operationGroupMembersCmd.AddCommand(operationPostTeamMembersCmd)

	return operationGroupMembersCmd, nil
}
func makeOperationGroupOperatingSystemsCmd() (*cobra.Command, error) {
	operationGroupOperatingSystemsCmd := &cobra.Command{
		Use:  "operating_systems",
		Long: ``,
	}

	operationGetPlansOperatingSystemCmd, err := makeOperationOperatingSystemsGetPlansOperatingSystemCmd()
	if err != nil {
		return nil, err
	}
	operationGroupOperatingSystemsCmd.AddCommand(operationGetPlansOperatingSystemCmd)

	return operationGroupOperatingSystemsCmd, nil
}
func makeOperationGroupOutOfBandCmd() (*cobra.Command, error) {
	operationGroupOutOfBandCmd := &cobra.Command{
		Use:  "out_of_band",
		Long: ``,
	}

	operationCreateServerOutOfBandCmd, err := makeOperationOutOfBandCreateServerOutOfBandCmd()
	if err != nil {
		return nil, err
	}
	operationGroupOutOfBandCmd.AddCommand(operationCreateServerOutOfBandCmd)

	operationGetServerOutOfBandCmd, err := makeOperationOutOfBandGetServerOutOfBandCmd()
	if err != nil {
		return nil, err
	}
	operationGroupOutOfBandCmd.AddCommand(operationGetServerOutOfBandCmd)

	return operationGroupOutOfBandCmd, nil
}
func makeOperationGroupPlansCmd() (*cobra.Command, error) {
	operationGroupPlansCmd := &cobra.Command{
		Use:  "plans",
		Long: ``,
	}

	operationGetBandwidthPlansCmd, err := makeOperationPlansGetBandwidthPlansCmd()
	if err != nil {
		return nil, err
	}
	operationGroupPlansCmd.AddCommand(operationGetBandwidthPlansCmd)

	operationGetPlanCmd, err := makeOperationPlansGetPlanCmd()
	if err != nil {
		return nil, err
	}
	operationGroupPlansCmd.AddCommand(operationGetPlanCmd)

	operationGetPlansCmd, err := makeOperationPlansGetPlansCmd()
	if err != nil {
		return nil, err
	}
	operationGroupPlansCmd.AddCommand(operationGetPlansCmd)

	return operationGroupPlansCmd, nil
}
func makeOperationGroupPowerActionsCmd() (*cobra.Command, error) {
	operationGroupPowerActionsCmd := &cobra.Command{
		Use:  "power_actions",
		Long: ``,
	}

	operationCreateServerActionCmd, err := makeOperationPowerActionsCreateServerActionCmd()
	if err != nil {
		return nil, err
	}
	operationGroupPowerActionsCmd.AddCommand(operationCreateServerActionCmd)

	return operationGroupPowerActionsCmd, nil
}
func makeOperationGroupProjectsCmd() (*cobra.Command, error) {
	operationGroupProjectsCmd := &cobra.Command{
		Use:  "projects",
		Long: ``,
	}

	operationCreateProjectCmd, err := makeOperationProjectsCreateProjectCmd()
	if err != nil {
		return nil, err
	}
	operationGroupProjectsCmd.AddCommand(operationCreateProjectCmd)

	operationDeleteProjectCmd, err := makeOperationProjectsDeleteProjectCmd()
	if err != nil {
		return nil, err
	}
	operationGroupProjectsCmd.AddCommand(operationDeleteProjectCmd)

	operationGetProjectCmd, err := makeOperationProjectsGetProjectCmd()
	if err != nil {
		return nil, err
	}
	operationGroupProjectsCmd.AddCommand(operationGetProjectCmd)

	operationGetProjectsCmd, err := makeOperationProjectsGetProjectsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupProjectsCmd.AddCommand(operationGetProjectsCmd)

	operationUpdateProjectCmd, err := makeOperationProjectsUpdateProjectCmd()
	if err != nil {
		return nil, err
	}
	operationGroupProjectsCmd.AddCommand(operationUpdateProjectCmd)

	return operationGroupProjectsCmd, nil
}
func makeOperationGroupRegionsCmd() (*cobra.Command, error) {
	operationGroupRegionsCmd := &cobra.Command{
		Use:  "regions",
		Long: ``,
	}

	operationGetRegionCmd, err := makeOperationRegionsGetRegionCmd()
	if err != nil {
		return nil, err
	}
	operationGroupRegionsCmd.AddCommand(operationGetRegionCmd)

	operationGetRegionsCmd, err := makeOperationRegionsGetRegionsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupRegionsCmd.AddCommand(operationGetRegionsCmd)

	return operationGroupRegionsCmd, nil
}
func makeOperationGroupRescueModeCmd() (*cobra.Command, error) {
	operationGroupRescueModeCmd := &cobra.Command{
		Use:  "rescue_mode",
		Long: ``,
	}

	operationServerExitRescueModeCmd, err := makeOperationRescueModeServerExitRescueModeCmd()
	if err != nil {
		return nil, err
	}
	operationGroupRescueModeCmd.AddCommand(operationServerExitRescueModeCmd)

	operationServerStartRescueModeCmd, err := makeOperationRescueModeServerStartRescueModeCmd()
	if err != nil {
		return nil, err
	}
	operationGroupRescueModeCmd.AddCommand(operationServerStartRescueModeCmd)

	return operationGroupRescueModeCmd, nil
}
func makeOperationGroupRolesCmd() (*cobra.Command, error) {
	operationGroupRolesCmd := &cobra.Command{
		Use:  "roles",
		Long: ``,
	}

	operationGetRoleIDCmd, err := makeOperationRolesGetRoleIDCmd()
	if err != nil {
		return nil, err
	}
	operationGroupRolesCmd.AddCommand(operationGetRoleIDCmd)

	operationGetRolesCmd, err := makeOperationRolesGetRolesCmd()
	if err != nil {
		return nil, err
	}
	operationGroupRolesCmd.AddCommand(operationGetRolesCmd)

	return operationGroupRolesCmd, nil
}
func makeOperationGroupServerReinstallCmd() (*cobra.Command, error) {
	operationGroupServerReinstallCmd := &cobra.Command{
		Use:  "server_reinstall",
		Long: ``,
	}

	operationCreateServerReinstallCmd, err := makeOperationServerReinstallCreateServerReinstallCmd()
	if err != nil {
		return nil, err
	}
	operationGroupServerReinstallCmd.AddCommand(operationCreateServerReinstallCmd)

	return operationGroupServerReinstallCmd, nil
}
func makeOperationGroupServersCmd() (*cobra.Command, error) {
	operationGroupServersCmd := &cobra.Command{
		Use:  "servers",
		Long: ``,
	}

	operationCreateServerCmd, err := makeOperationServersCreateServerCmd()
	if err != nil {
		return nil, err
	}
	operationGroupServersCmd.AddCommand(operationCreateServerCmd)

	operationDestroyServerCmd, err := makeOperationServersDestroyServerCmd()
	if err != nil {
		return nil, err
	}
	operationGroupServersCmd.AddCommand(operationDestroyServerCmd)

	operationGetServerCmd, err := makeOperationServersGetServerCmd()
	if err != nil {
		return nil, err
	}
	operationGroupServersCmd.AddCommand(operationGetServerCmd)

	operationGetServersCmd, err := makeOperationServersGetServersCmd()
	if err != nil {
		return nil, err
	}
	operationGroupServersCmd.AddCommand(operationGetServersCmd)

	operationServerScheduleDeletionCmd, err := makeOperationServersServerScheduleDeletionCmd()
	if err != nil {
		return nil, err
	}
	operationGroupServersCmd.AddCommand(operationServerScheduleDeletionCmd)

	operationServerUnscheduleDeletionCmd, err := makeOperationServersServerUnscheduleDeletionCmd()
	if err != nil {
		return nil, err
	}
	operationGroupServersCmd.AddCommand(operationServerUnscheduleDeletionCmd)

	operationUpdateServerCmd, err := makeOperationServersUpdateServerCmd()
	if err != nil {
		return nil, err
	}
	operationGroupServersCmd.AddCommand(operationUpdateServerCmd)

	return operationGroupServersCmd, nil
}
func makeOperationGroupSSHKeysCmd() (*cobra.Command, error) {
	operationGroupSSHKeysCmd := &cobra.Command{
		Use:  "ssh_keys",
		Long: ``,
	}

	operationDeleteProjectSSHKeyCmd, err := makeOperationSSHKeysDeleteProjectSSHKeyCmd()
	if err != nil {
		return nil, err
	}
	operationGroupSSHKeysCmd.AddCommand(operationDeleteProjectSSHKeyCmd)

	operationGetProjectSSHKeyCmd, err := makeOperationSSHKeysGetProjectSSHKeyCmd()
	if err != nil {
		return nil, err
	}
	operationGroupSSHKeysCmd.AddCommand(operationGetProjectSSHKeyCmd)

	operationGetProjectSSHKeysCmd, err := makeOperationSSHKeysGetProjectSSHKeysCmd()
	if err != nil {
		return nil, err
	}
	operationGroupSSHKeysCmd.AddCommand(operationGetProjectSSHKeysCmd)

	operationPostProjectSSHKeyCmd, err := makeOperationSSHKeysPostProjectSSHKeyCmd()
	if err != nil {
		return nil, err
	}
	operationGroupSSHKeysCmd.AddCommand(operationPostProjectSSHKeyCmd)

	operationPutProjectSSHKeyCmd, err := makeOperationSSHKeysPutProjectSSHKeyCmd()
	if err != nil {
		return nil, err
	}
	operationGroupSSHKeysCmd.AddCommand(operationPutProjectSSHKeyCmd)

	return operationGroupSSHKeysCmd, nil
}
func makeOperationGroupTeamsCmd() (*cobra.Command, error) {
	operationGroupTeamsCmd := &cobra.Command{
		Use:  "teams",
		Long: ``,
	}

	operationGetTeamCmd, err := makeOperationTeamsGetTeamCmd()
	if err != nil {
		return nil, err
	}
	operationGroupTeamsCmd.AddCommand(operationGetTeamCmd)

	operationPatchCurrentTeamCmd, err := makeOperationTeamsPatchCurrentTeamCmd()
	if err != nil {
		return nil, err
	}
	operationGroupTeamsCmd.AddCommand(operationPatchCurrentTeamCmd)

	operationPostTeamCmd, err := makeOperationTeamsPostTeamCmd()
	if err != nil {
		return nil, err
	}
	operationGroupTeamsCmd.AddCommand(operationPostTeamCmd)

	return operationGroupTeamsCmd, nil
}
func makeOperationGroupUserDataCmd() (*cobra.Command, error) {
	operationGroupUserDataCmd := &cobra.Command{
		Use:  "user_data",
		Long: ``,
	}

	operationDeleteProjectUserDataCmd, err := makeOperationUserDataDeleteProjectUserDataCmd()
	if err != nil {
		return nil, err
	}
	operationGroupUserDataCmd.AddCommand(operationDeleteProjectUserDataCmd)

	operationGetProjectUserDataCmd, err := makeOperationUserDataGetProjectUserDataCmd()
	if err != nil {
		return nil, err
	}
	operationGroupUserDataCmd.AddCommand(operationGetProjectUserDataCmd)

	operationGetProjectUsersDataCmd, err := makeOperationUserDataGetProjectUsersDataCmd()
	if err != nil {
		return nil, err
	}
	operationGroupUserDataCmd.AddCommand(operationGetProjectUsersDataCmd)

	operationPostProjectUserDataCmd, err := makeOperationUserDataPostProjectUserDataCmd()
	if err != nil {
		return nil, err
	}
	operationGroupUserDataCmd.AddCommand(operationPostProjectUserDataCmd)

	operationPutProjectUserDataCmd, err := makeOperationUserDataPutProjectUserDataCmd()
	if err != nil {
		return nil, err
	}
	operationGroupUserDataCmd.AddCommand(operationPutProjectUserDataCmd)

	return operationGroupUserDataCmd, nil
}
func makeOperationGroupVpnSessionsCmd() (*cobra.Command, error) {
	operationGroupVpnSessionsCmd := &cobra.Command{
		Use:  "v_p_n_sessions",
		Long: ``,
	}

	operationDeleteVpnSessionCmd, err := makeOperationVpnSessionsDeleteVpnSessionCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVpnSessionsCmd.AddCommand(operationDeleteVpnSessionCmd)

	operationGetVpnSessionsCmd, err := makeOperationVpnSessionsGetVpnSessionsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVpnSessionsCmd.AddCommand(operationGetVpnSessionsCmd)

	operationPostVpnSessionCmd, err := makeOperationVpnSessionsPostVpnSessionCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVpnSessionsCmd.AddCommand(operationPostVpnSessionCmd)

	operationPutVpnSessionCmd, err := makeOperationVpnSessionsPutVpnSessionCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVpnSessionsCmd.AddCommand(operationPutVpnSessionCmd)

	return operationGroupVpnSessionsCmd, nil
}
func makeOperationGroupVirtualNetworkAssignmentsCmd() (*cobra.Command, error) {
	operationGroupVirtualNetworkAssignmentsCmd := &cobra.Command{
		Use:  "virtual_network_assignments",
		Long: ``,
	}

	operationAssignServerVirtualNetworkCmd, err := makeOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVirtualNetworkAssignmentsCmd.AddCommand(operationAssignServerVirtualNetworkCmd)

	operationDeleteVirtualNetworksAssignmentsCmd, err := makeOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVirtualNetworkAssignmentsCmd.AddCommand(operationDeleteVirtualNetworksAssignmentsCmd)

	operationGetVirtualNetworksAssignmentsCmd, err := makeOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVirtualNetworkAssignmentsCmd.AddCommand(operationGetVirtualNetworksAssignmentsCmd)

	return operationGroupVirtualNetworkAssignmentsCmd, nil
}
func makeOperationGroupVirtualNetworksCmd() (*cobra.Command, error) {
	operationGroupVirtualNetworksCmd := &cobra.Command{
		Use:  "virtual_networks",
		Long: ``,
	}

	operationCreateVirtualNetworkCmd, err := makeOperationVirtualNetworksCreateVirtualNetworkCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVirtualNetworksCmd.AddCommand(operationCreateVirtualNetworkCmd)

	operationDestroyVirtualNetworkCmd, err := makeOperationVirtualNetworksDestroyVirtualNetworkCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVirtualNetworksCmd.AddCommand(operationDestroyVirtualNetworkCmd)

	operationGetVirtualNetworkCmd, err := makeOperationVirtualNetworksGetVirtualNetworkCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVirtualNetworksCmd.AddCommand(operationGetVirtualNetworkCmd)

	operationGetVirtualNetworksCmd, err := makeOperationVirtualNetworksGetVirtualNetworksCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVirtualNetworksCmd.AddCommand(operationGetVirtualNetworksCmd)

	operationUpdateVirtualNetworkCmd, err := makeOperationVirtualNetworksUpdateVirtualNetworkCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVirtualNetworksCmd.AddCommand(operationUpdateVirtualNetworkCmd)

	return operationGroupVirtualNetworksCmd, nil
}
