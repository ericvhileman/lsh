// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/plans"
	"github.com/latitudesh/cli/internal/api"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationPlansGetPlansCmd returns a cmd to handle operation getPlans
func makeOperationPlansGetPlansCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "list",
		Short: `Lists all plans. Availability by region is included in ` + "`" + `attributes.regions.locations.available[*]` + "`" + ` node for a given plan.`,
		RunE: runOperationPlansGetPlans,
	}

	if err := registerOperationPlansGetPlansParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPlansGetPlans uses cmd flags to call endpoint api
func runOperationPlansGetPlans(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := plans.NewGetPlansParams()
	if err, _ := retrieveOperationPlansGetPlansAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPlansGetPlansFilterDiskFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPlansGetPlansFilterGpuFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPlansGetPlansFilterInStockFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPlansGetPlansFilterLocationFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPlansGetPlansFilterNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPlansGetPlansFilterRAMFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPlansGetPlansFilterSlugFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPlansGetPlansFilterStockLevelFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	result, err := appCli.Plans.GetPlans(params, nil)
	if err != nil {
		api.RenderErrorOutput(err)
		return nil
	}

	msgStr, err := parseOperationPlansGetPlansResult(result)
	if err != nil {
		return err
	}
	if !debug {
		utils.PrintOutput(msgStr)
	}
	return nil
}

// registerOperationPlansGetPlansParamFlags registers all flags needed to fill params
func registerOperationPlansGetPlansParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPlansGetPlansAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPlansGetPlansFilterDiskParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPlansGetPlansFilterGpuParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPlansGetPlansFilterInStockParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPlansGetPlansFilterLocationParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPlansGetPlansFilterNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPlansGetPlansFilterRAMParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPlansGetPlansFilterSlugParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPlansGetPlansFilterStockLevelParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPlansGetPlansAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationPlansGetPlansFilterDiskParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterDiskDescription := `The disk size in Gigabytes to filter by, should be used with the following options:
                              [eql] to filter for values equal to the provided value.
                              [gte] to filter for values greater or equal to the provided value.
                              [lte] to filter by values lower or equal to the provided value.`

	var filterDiskFlagName string
	if cmdPrefix == "" {
		filterDiskFlagName = "filter[disk]"
	} else {
		filterDiskFlagName = fmt.Sprintf("%v.filter[disk]", cmdPrefix)
	}

	var filterDiskFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterDiskFlagName, filterDiskFlagDefault, filterDiskDescription)

	return nil
}
func registerOperationPlansGetPlansFilterGpuParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterGpuDescription := `Filter by the existence of an associated GPU`

	var filterGpuFlagName string
	if cmdPrefix == "" {
		filterGpuFlagName = "filter[gpu]"
	} else {
		filterGpuFlagName = fmt.Sprintf("%v.filter[gpu]", cmdPrefix)
	}

	var filterGpuFlagDefault bool

	_ = cmd.PersistentFlags().Bool(filterGpuFlagName, filterGpuFlagDefault, filterGpuDescription)

	return nil
}
func registerOperationPlansGetPlansFilterInStockParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterInStockDescription := `The stock available at the site to filter by`

	var filterInStockFlagName string
	if cmdPrefix == "" {
		filterInStockFlagName = "filter[in_stock]"
	} else {
		filterInStockFlagName = fmt.Sprintf("%v.filter[in_stock]", cmdPrefix)
	}

	var filterInStockFlagDefault bool

	_ = cmd.PersistentFlags().Bool(filterInStockFlagName, filterInStockFlagDefault, filterInStockDescription)

	return nil
}
func registerOperationPlansGetPlansFilterLocationParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterLocationDescription := `The location of the site to filter by`

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
func registerOperationPlansGetPlansFilterNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterNameDescription := `The plan name to filter by`

	var filterNameFlagName string
	if cmdPrefix == "" {
		filterNameFlagName = "filter[name]"
	} else {
		filterNameFlagName = fmt.Sprintf("%v.filter[name]", cmdPrefix)
	}

	var filterNameFlagDefault string

	_ = cmd.PersistentFlags().String(filterNameFlagName, filterNameFlagDefault, filterNameDescription)

	return nil
}
func registerOperationPlansGetPlansFilterRAMParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterRamDescription := `The ram size in Gigabytes to filter by, should be used with the following options:
                              [eql] to filter for values equal to the provided value.
                              [gte] to filter for values greater or equal to the provided value.
                              [lte] to filter by values lower or equal to the provided value.`

	var filterRamFlagName string
	if cmdPrefix == "" {
		filterRamFlagName = "filter[ram]"
	} else {
		filterRamFlagName = fmt.Sprintf("%v.filter[ram]", cmdPrefix)
	}

	var filterRamFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterRamFlagName, filterRamFlagDefault, filterRamDescription)

	return nil
}
func registerOperationPlansGetPlansFilterSlugParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterSlugDescription := `The plan slug to filter by`

	var filterSlugFlagName string
	if cmdPrefix == "" {
		filterSlugFlagName = "filter[slug]"
	} else {
		filterSlugFlagName = fmt.Sprintf("%v.filter[slug]", cmdPrefix)
	}

	var filterSlugFlagDefault string

	_ = cmd.PersistentFlags().String(filterSlugFlagName, filterSlugFlagDefault, filterSlugDescription)

	return nil
}
func registerOperationPlansGetPlansFilterStockLevelParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterStockLevelDescription := `Enum: ["Unavailable","Low","Medium","High","Unique"]. The stock level at the site to filter by`

	var filterStockLevelFlagName string
	if cmdPrefix == "" {
		filterStockLevelFlagName = "filter[stock_level]"
	} else {
		filterStockLevelFlagName = fmt.Sprintf("%v.filter[stock_level]", cmdPrefix)
	}

	var filterStockLevelFlagDefault string

	_ = cmd.PersistentFlags().String(filterStockLevelFlagName, filterStockLevelFlagDefault, filterStockLevelDescription)

	if err := cmd.RegisterFlagCompletionFunc(filterStockLevelFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["Unavailable","Low","Medium","High","Unique"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func retrieveOperationPlansGetPlansAPIVersionFlag(m *plans.GetPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationPlansGetPlansFilterDiskFlag(m *plans.GetPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[disk]") {

		var filterDiskFlagName string
		if cmdPrefix == "" {
			filterDiskFlagName = "filter[disk]"
		} else {
			filterDiskFlagName = fmt.Sprintf("%v.filter[disk]", cmdPrefix)
		}

		filterDiskFlagValue, err := cmd.Flags().GetInt64(filterDiskFlagName)
		if err != nil {
			return err, false
		}
		m.FilterDisk = &filterDiskFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPlansGetPlansFilterGpuFlag(m *plans.GetPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[gpu]") {

		var filterGpuFlagName string
		if cmdPrefix == "" {
			filterGpuFlagName = "filter[gpu]"
		} else {
			filterGpuFlagName = fmt.Sprintf("%v.filter[gpu]", cmdPrefix)
		}

		filterGpuFlagValue, err := cmd.Flags().GetBool(filterGpuFlagName)
		if err != nil {
			return err, false
		}
		m.FilterGpu = &filterGpuFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPlansGetPlansFilterInStockFlag(m *plans.GetPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[in_stock]") {

		var filterInStockFlagName string
		if cmdPrefix == "" {
			filterInStockFlagName = "filter[in_stock]"
		} else {
			filterInStockFlagName = fmt.Sprintf("%v.filter[in_stock]", cmdPrefix)
		}

		filterInStockFlagValue, err := cmd.Flags().GetBool(filterInStockFlagName)
		if err != nil {
			return err, false
		}
		m.FilterInStock = &filterInStockFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPlansGetPlansFilterLocationFlag(m *plans.GetPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationPlansGetPlansFilterNameFlag(m *plans.GetPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[name]") {

		var filterNameFlagName string
		if cmdPrefix == "" {
			filterNameFlagName = "filter[name]"
		} else {
			filterNameFlagName = fmt.Sprintf("%v.filter[name]", cmdPrefix)
		}

		filterNameFlagValue, err := cmd.Flags().GetString(filterNameFlagName)
		if err != nil {
			return err, false
		}
		m.FilterName = &filterNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPlansGetPlansFilterRAMFlag(m *plans.GetPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[ram]") {

		var filterRamFlagName string
		if cmdPrefix == "" {
			filterRamFlagName = "filter[ram]"
		} else {
			filterRamFlagName = fmt.Sprintf("%v.filter[ram]", cmdPrefix)
		}

		filterRamFlagValue, err := cmd.Flags().GetInt64(filterRamFlagName)
		if err != nil {
			return err, false
		}
		m.FilterRAM = &filterRamFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPlansGetPlansFilterSlugFlag(m *plans.GetPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[slug]") {

		var filterSlugFlagName string
		if cmdPrefix == "" {
			filterSlugFlagName = "filter[slug]"
		} else {
			filterSlugFlagName = fmt.Sprintf("%v.filter[slug]", cmdPrefix)
		}

		filterSlugFlagValue, err := cmd.Flags().GetString(filterSlugFlagName)
		if err != nil {
			return err, false
		}
		m.FilterSlug = &filterSlugFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPlansGetPlansFilterStockLevelFlag(m *plans.GetPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter[stock_level]") {

		var filterStockLevelFlagName string
		if cmdPrefix == "" {
			filterStockLevelFlagName = "filter[stock_level]"
		} else {
			filterStockLevelFlagName = fmt.Sprintf("%v.filter[stock_level]", cmdPrefix)
		}

		filterStockLevelFlagValue, err := cmd.Flags().GetString(filterStockLevelFlagName)
		if err != nil {
			return err, false
		}
		m.FilterStockLevel = &filterStockLevelFlagValue

	}
	return nil, retAdded
}

// parseOperationPlansGetPlansResult parses request result and return the string content
func parseOperationPlansGetPlansResult(resp0 *plans.GetPlansOK) (string, error) {
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
func registerModelGetPlansOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetPlansOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetPlansOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data []*models.PlanData array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetPlansOKBodyFlags(depth int, m *plans.GetPlansOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveGetPlansOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveGetPlansOKBodyDataFlags(depth int, m *plans.GetPlansOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// warning: data array type []*models.PlanData is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
