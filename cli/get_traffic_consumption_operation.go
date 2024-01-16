// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/bandwidth"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationBandwidthGetTrafficConsumptionCmd returns a cmd to handle operation getTrafficConsumption
func makeOperationBandwidthGetTrafficConsumptionCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "get-traffic-consumption",
		Short: ``,
		RunE:  runOperationBandwidthGetTrafficConsumption,
	}

	if err := registerOperationBandwidthGetTrafficConsumptionParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBandwidthGetTrafficConsumption uses cmd flags to call endpoint api
func runOperationBandwidthGetTrafficConsumption(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := bandwidth.NewGetTrafficConsumptionParams()
	if err, _ := retrieveOperationBandwidthGetTrafficConsumptionAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBandwidthGetTrafficConsumptionFilterFromDateFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBandwidthGetTrafficConsumptionFilterProjectFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBandwidthGetTrafficConsumptionFilterServerFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBandwidthGetTrafficConsumptionFilterToDateFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBandwidthGetTrafficConsumptionResult(appCli.Bandwidth.GetTrafficConsumption(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(utils.PrettifyJson(msgStr))
	}
	return nil
}

// registerOperationBandwidthGetTrafficConsumptionParamFlags registers all flags needed to fill params
func registerOperationBandwidthGetTrafficConsumptionParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBandwidthGetTrafficConsumptionAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBandwidthGetTrafficConsumptionFilterFromDateParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBandwidthGetTrafficConsumptionFilterProjectParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBandwidthGetTrafficConsumptionFilterServerParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBandwidthGetTrafficConsumptionFilterToDateParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBandwidthGetTrafficConsumptionAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationBandwidthGetTrafficConsumptionFilterFromDateParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterFromDateDescription := `The start timestamp. Must be a unix timestamp`

	var filterFromDateFlagName string
	if cmdPrefix == "" {
		filterFromDateFlagName = "filter[from_date]"
	} else {
		filterFromDateFlagName = fmt.Sprintf("%v.filter[from_date]", cmdPrefix)
	}

	var filterFromDateFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterFromDateFlagName, filterFromDateFlagDefault, filterFromDateDescription)

	return nil
}
func registerOperationBandwidthGetTrafficConsumptionFilterProjectParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterProjectDescription := ``

	var filterProjectFlagName string
	if cmdPrefix == "" {
		filterProjectFlagName = "filter[project]"
	} else {
		filterProjectFlagName = fmt.Sprintf("%v.filter[project]", cmdPrefix)
	}

	var filterProjectFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterProjectFlagName, filterProjectFlagDefault, filterProjectDescription)

	return nil
}
func registerOperationBandwidthGetTrafficConsumptionFilterServerParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterServerDescription := ``

	var filterServerFlagName string
	if cmdPrefix == "" {
		filterServerFlagName = "filter[server]"
	} else {
		filterServerFlagName = fmt.Sprintf("%v.filter[server]", cmdPrefix)
	}

	var filterServerFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterServerFlagName, filterServerFlagDefault, filterServerDescription)

	return nil
}
func registerOperationBandwidthGetTrafficConsumptionFilterToDateParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterToDateDescription := `The end timestamp. Must be a unix timestamp`

	var filterToDateFlagName string
	if cmdPrefix == "" {
		filterToDateFlagName = "filter[to_date]"
	} else {
		filterToDateFlagName = fmt.Sprintf("%v.filter[to_date]", cmdPrefix)
	}

	var filterToDateFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterToDateFlagName, filterToDateFlagDefault, filterToDateDescription)

	return nil
}

func retrieveOperationBandwidthGetTrafficConsumptionAPIVersionFlag(m *bandwidth.GetTrafficConsumptionParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBandwidthGetTrafficConsumptionFilterFromDateFlag(m *bandwidth.GetTrafficConsumptionParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[from_date]") {

		var filterFromDateFlagName string
		if cmdPrefix == "" {
			filterFromDateFlagName = "filter[from_date]"
		} else {
			filterFromDateFlagName = fmt.Sprintf("%v.filter[from_date]", cmdPrefix)
		}

		filterFromDateFlagValue, err := cmd.Flags().GetInt64(filterFromDateFlagName)
		if err != nil {
			return err, false
		}
		m.FilterFromDate = &filterFromDateFlagValue

	}
	return nil, retAdded
}
func retrieveOperationBandwidthGetTrafficConsumptionFilterProjectFlag(m *bandwidth.GetTrafficConsumptionParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[project]") {

		var filterProjectFlagName string
		if cmdPrefix == "" {
			filterProjectFlagName = "filter[project]"
		} else {
			filterProjectFlagName = fmt.Sprintf("%v.filter[project]", cmdPrefix)
		}

		filterProjectFlagValue, err := cmd.Flags().GetInt64(filterProjectFlagName)
		if err != nil {
			return err, false
		}
		m.FilterProject = &filterProjectFlagValue

	}
	return nil, retAdded
}
func retrieveOperationBandwidthGetTrafficConsumptionFilterServerFlag(m *bandwidth.GetTrafficConsumptionParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[server]") {

		var filterServerFlagName string
		if cmdPrefix == "" {
			filterServerFlagName = "filter[server]"
		} else {
			filterServerFlagName = fmt.Sprintf("%v.filter[server]", cmdPrefix)
		}

		filterServerFlagValue, err := cmd.Flags().GetInt64(filterServerFlagName)
		if err != nil {
			return err, false
		}
		m.FilterServer = &filterServerFlagValue

	}
	return nil, retAdded
}
func retrieveOperationBandwidthGetTrafficConsumptionFilterToDateFlag(m *bandwidth.GetTrafficConsumptionParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[to_date]") {

		var filterToDateFlagName string
		if cmdPrefix == "" {
			filterToDateFlagName = "filter[to_date]"
		} else {
			filterToDateFlagName = fmt.Sprintf("%v.filter[to_date]", cmdPrefix)
		}

		filterToDateFlagValue, err := cmd.Flags().GetInt64(filterToDateFlagName)
		if err != nil {
			return err, false
		}
		m.FilterToDate = &filterToDateFlagValue

	}
	return nil, retAdded
}

// parseOperationBandwidthGetTrafficConsumptionResult parses request result and return the string content
func parseOperationBandwidthGetTrafficConsumptionResult(resp0 *bandwidth.GetTrafficConsumptionOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*bandwidth.GetTrafficConsumptionOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
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