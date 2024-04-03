package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/lsh/client/plans"
	"github.com/latitudesh/lsh/cmd/lsh"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

// makeOperationPlansGetPlansCmd returns a cmd to handle operation getPlans
func makeOperationPlansGetPlansCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List available plans",
		RunE:  runOperationPlansGetPlans,
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
	if err, _ := retrieveOperationPlansGetPlansFilterDiskEqlFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPlansGetPlansFilterDiskLteFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPlansGetPlansFilterDiskGteFlag(params, "", cmd); err != nil {
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
	if err, _ := retrieveOperationPlansGetPlansFilterRAMEqlFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPlansGetPlansFilterRAMLteFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPlansGetPlansFilterRAMGteFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPlansGetPlansFilterSlugFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPlansGetPlansFilterStockLevelFlag(params, "", cmd); err != nil {
		return err
	}
	if lsh.DryRun {

		lsh.LogDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Plans.GetPlans(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !lsh.Debug {
		utils.Render(response.GetData())
	}
	return nil
}

// registerOperationPlansGetPlansParamFlags registers all flags needed to fill params
func registerOperationPlansGetPlansParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPlansGetPlansFilterDiskEqlParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPlansGetPlansFilterDiskLteParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPlansGetPlansFilterDiskGteParamFlags("", cmd); err != nil {
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
	if err := registerOperationPlansGetPlansFilterRAMEqlParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPlansGetPlansFilterRAMLteParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPlansGetPlansFilterRAMGteParamFlags("", cmd); err != nil {
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

func registerOperationPlansGetPlansFilterDiskEqlParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterDiskDescription := "Filter plans with disk size in Gigabytes equals the provided value."

	var filterDiskFlagName string
	if cmdPrefix == "" {
		filterDiskFlagName = "disk_eql"
	} else {
		filterDiskFlagName = fmt.Sprintf("%v.disk_eql", cmdPrefix)
	}

	var filterDiskFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterDiskFlagName, filterDiskFlagDefault, filterDiskDescription)

	return nil
}
func registerOperationPlansGetPlansFilterDiskLteParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterDiskDescription := "Filter plans with disk size in Gigabytes less than or equal the provided value."

	var filterDiskFlagName string
	if cmdPrefix == "" {
		filterDiskFlagName = "disk_lte"
	} else {
		filterDiskFlagName = fmt.Sprintf("%v.disk_lte", cmdPrefix)
	}

	var filterDiskFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterDiskFlagName, filterDiskFlagDefault, filterDiskDescription)

	return nil
}
func registerOperationPlansGetPlansFilterDiskGteParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterDiskDescription := "Filter plans with disk size in Gigabytes greater than or equal the provided value."

	var filterDiskFlagName string
	if cmdPrefix == "" {
		filterDiskFlagName = "disk_gte"
	} else {
		filterDiskFlagName = fmt.Sprintf("%v.disk_gte", cmdPrefix)
	}

	var filterDiskFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterDiskFlagName, filterDiskFlagDefault, filterDiskDescription)

	return nil
}
func registerOperationPlansGetPlansFilterGpuParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterGpuDescription := `Filter by the existence of an associated GPU`

	var filterGpuFlagName string
	if cmdPrefix == "" {
		filterGpuFlagName = "gpu"
	} else {
		filterGpuFlagName = fmt.Sprintf("%v.gpu", cmdPrefix)
	}

	var filterGpuFlagDefault bool

	_ = cmd.PersistentFlags().Bool(filterGpuFlagName, filterGpuFlagDefault, filterGpuDescription)

	return nil
}
func registerOperationPlansGetPlansFilterInStockParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterInStockDescription := `The stock available at the site to filter by`

	var filterInStockFlagName string
	if cmdPrefix == "" {
		filterInStockFlagName = "in_stock"
	} else {
		filterInStockFlagName = fmt.Sprintf("%v.in_stock", cmdPrefix)
	}

	var filterInStockFlagDefault bool

	_ = cmd.PersistentFlags().Bool(filterInStockFlagName, filterInStockFlagDefault, filterInStockDescription)

	return nil
}
func registerOperationPlansGetPlansFilterLocationParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterLocationDescription := `The location of the site to filter by`

	var filterLocationFlagName string
	if cmdPrefix == "" {
		filterLocationFlagName = "location"
	} else {
		filterLocationFlagName = fmt.Sprintf("%v.location", cmdPrefix)
	}

	var filterLocationFlagDefault string

	_ = cmd.PersistentFlags().String(filterLocationFlagName, filterLocationFlagDefault, filterLocationDescription)

	return nil
}
func registerOperationPlansGetPlansFilterNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterNameDescription := `The plan name to filter by`

	var filterNameFlagName string
	if cmdPrefix == "" {
		filterNameFlagName = "name"
	} else {
		filterNameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var filterNameFlagDefault string

	_ = cmd.PersistentFlags().String(filterNameFlagName, filterNameFlagDefault, filterNameDescription)

	return nil
}
func registerOperationPlansGetPlansFilterRAMEqlParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterRamEqlDescription := `Filter plans with RAM size (in GB) equals the provided value.`

	var filterRamEqlFlagName string
	if cmdPrefix == "" {
		filterRamEqlFlagName = "ram_eql"
	} else {
		filterRamEqlFlagName = fmt.Sprintf("%v.ram_eql", cmdPrefix)
	}

	var filterRamEqlFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterRamEqlFlagName, filterRamEqlFlagDefault, filterRamEqlDescription)

	return nil
}
func registerOperationPlansGetPlansFilterRAMGteParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterRamGteDescription := `Filter plans with RAM size (in GB) greater than or equal the provided value.`

	var filterRamGteFlagName string
	if cmdPrefix == "" {
		filterRamGteFlagName = "ram_gte"
	} else {
		filterRamGteFlagName = fmt.Sprintf("%v.ram_gte", cmdPrefix)
	}

	var filterRamGteFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterRamGteFlagName, filterRamGteFlagDefault, filterRamGteDescription)

	return nil
}
func registerOperationPlansGetPlansFilterRAMLteParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterRamLteDescription := `Filter plans with RAM size (in GB) less than or equal the provided value.`

	var filterRamLteFlagName string
	if cmdPrefix == "" {
		filterRamLteFlagName = "ram_lte"
	} else {
		filterRamLteFlagName = fmt.Sprintf("%v.ram_lte", cmdPrefix)
	}

	var filterRamLteFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterRamLteFlagName, filterRamLteFlagDefault, filterRamLteDescription)

	return nil
}
func registerOperationPlansGetPlansFilterSlugParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterSlugDescription := `The plan slug to filter by`

	var filterSlugFlagName string
	if cmdPrefix == "" {
		filterSlugFlagName = "slug"
	} else {
		filterSlugFlagName = fmt.Sprintf("%v.slug", cmdPrefix)
	}

	var filterSlugFlagDefault string

	_ = cmd.PersistentFlags().String(filterSlugFlagName, filterSlugFlagDefault, filterSlugDescription)

	return nil
}
func registerOperationPlansGetPlansFilterStockLevelParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterStockLevelDescription := `Enum: ["Unavailable","Low","Medium","High","Unique"]. The stock level at the site to filter by`

	var filterStockLevelFlagName string
	if cmdPrefix == "" {
		filterStockLevelFlagName = "stock_level"
	} else {
		filterStockLevelFlagName = fmt.Sprintf("%v.stock_level", cmdPrefix)
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

func retrieveOperationPlansGetPlansFilterDiskEqlFlag(m *plans.GetPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("disk_eql") {

		var filterDiskFlagName string
		if cmdPrefix == "" {
			filterDiskFlagName = "disk_eql"
		} else {
			filterDiskFlagName = fmt.Sprintf("%v.disk_eql", cmdPrefix)
		}

		filterDiskFlagValue, err := cmd.Flags().GetInt64(filterDiskFlagName)
		if err != nil {
			return err, false
		}
		m.FilterDiskEql = &filterDiskFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPlansGetPlansFilterDiskLteFlag(m *plans.GetPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("disk_lte") {

		var filterDiskFlagName string
		if cmdPrefix == "" {
			filterDiskFlagName = "disk_lte"
		} else {
			filterDiskFlagName = fmt.Sprintf("%v.disk_lte", cmdPrefix)
		}

		filterDiskFlagValue, err := cmd.Flags().GetInt64(filterDiskFlagName)
		if err != nil {
			return err, false
		}
		m.FilterDiskLte = &filterDiskFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPlansGetPlansFilterDiskGteFlag(m *plans.GetPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("disk_gte") {

		var filterDiskFlagName string
		if cmdPrefix == "" {
			filterDiskFlagName = "disk_gte"
		} else {
			filterDiskFlagName = fmt.Sprintf("%v.disk_gte", cmdPrefix)
		}

		filterDiskFlagValue, err := cmd.Flags().GetInt64(filterDiskFlagName)
		if err != nil {
			return err, false
		}
		m.FilterDiskGte = &filterDiskFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPlansGetPlansFilterGpuFlag(m *plans.GetPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("gpu") {

		var filterGpuFlagName string
		if cmdPrefix == "" {
			filterGpuFlagName = "gpu"
		} else {
			filterGpuFlagName = fmt.Sprintf("%v.gpu", cmdPrefix)
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
	if cmd.Flags().Changed("in_stock") {

		var filterInStockFlagName string
		if cmdPrefix == "" {
			filterInStockFlagName = "in_stock"
		} else {
			filterInStockFlagName = fmt.Sprintf("%v.in_stock", cmdPrefix)
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
	if cmd.Flags().Changed("location") {

		var filterLocationFlagName string
		if cmdPrefix == "" {
			filterLocationFlagName = "location"
		} else {
			filterLocationFlagName = fmt.Sprintf("%v.location", cmdPrefix)
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
	if cmd.Flags().Changed("name") {

		var filterNameFlagName string
		if cmdPrefix == "" {
			filterNameFlagName = "name"
		} else {
			filterNameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		filterNameFlagValue, err := cmd.Flags().GetString(filterNameFlagName)
		if err != nil {
			return err, false
		}
		m.FilterName = &filterNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPlansGetPlansFilterRAMEqlFlag(m *plans.GetPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("ram_eql") {

		var filterRamFlagName string
		if cmdPrefix == "" {
			filterRamFlagName = "ram_eql"
		} else {
			filterRamFlagName = fmt.Sprintf("%v.ram_eql", cmdPrefix)
		}

		filterRamFlagValue, err := cmd.Flags().GetInt64(filterRamFlagName)
		if err != nil {
			return err, false
		}
		m.FilterRAMEql = &filterRamFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPlansGetPlansFilterRAMLteFlag(m *plans.GetPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("ram_lte") {

		var filterRamFlagName string
		if cmdPrefix == "" {
			filterRamFlagName = "ram_lte"
		} else {
			filterRamFlagName = fmt.Sprintf("%v.ram_lte", cmdPrefix)
		}

		filterRamFlagValue, err := cmd.Flags().GetInt64(filterRamFlagName)
		if err != nil {
			return err, false
		}
		m.FilterRAMLte = &filterRamFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPlansGetPlansFilterRAMGteFlag(m *plans.GetPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("ram_gte") {

		var filterRamFlagName string
		if cmdPrefix == "" {
			filterRamFlagName = "ram_gte"
		} else {
			filterRamFlagName = fmt.Sprintf("%v.ram_gte", cmdPrefix)
		}

		filterRamFlagValue, err := cmd.Flags().GetInt64(filterRamFlagName)
		if err != nil {
			return err, false
		}
		m.FilterRAMGte = &filterRamFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPlansGetPlansFilterSlugFlag(m *plans.GetPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("slug") {

		var filterSlugFlagName string
		if cmdPrefix == "" {
			filterSlugFlagName = "slug"
		} else {
			filterSlugFlagName = fmt.Sprintf("%v.slug", cmdPrefix)
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
	if cmd.Flags().Changed("stock_level") {

		var filterStockLevelFlagName string
		if cmdPrefix == "" {
			filterStockLevelFlagName = "stock_level"
		} else {
			filterStockLevelFlagName = fmt.Sprintf("%v.stock_level", cmdPrefix)
		}

		filterStockLevelFlagValue, err := cmd.Flags().GetString(filterStockLevelFlagName)
		if err != nil {
			return err, false
		}
		m.FilterStockLevel = &filterStockLevelFlagValue

	}
	return nil, retAdded
}
