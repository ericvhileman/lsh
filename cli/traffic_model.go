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

// Schema cli for Traffic

// register flags to command
func registerModelTrafficFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTrafficData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTrafficData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelTrafficDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTrafficFlags(depth int, m *models.Traffic, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveTrafficDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveTrafficDataFlags(depth int, m *models.Traffic, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data TrafficData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.TrafficData{}
	}

	err, dataAdded := retrieveModelTrafficDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}

// Extra schema cli for TrafficData

// register flags to command
func registerModelTrafficDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTrafficDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficDataID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTrafficDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelTrafficDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerTrafficDataID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerTrafficDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["traffic"]. `

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
			if err := json.Unmarshal([]byte(`["traffic"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTrafficDataFlags(depth int, m *models.TrafficData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveTrafficDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrieveTrafficDataIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, typeAdded := retrieveTrafficDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveTrafficDataAttributesFlags(depth int, m *models.TrafficData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes TrafficDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &models.TrafficDataAttributes{}
	}

	err, attributesAdded := retrieveModelTrafficDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveTrafficDataIDFlags(depth int, m *models.TrafficData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveTrafficDataTypeFlags(depth int, m *models.TrafficData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// Extra schema cli for TrafficDataAttributes

// register flags to command
func registerModelTrafficDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTrafficDataAttributesFromDate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficDataAttributesRegions(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficDataAttributesToDate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficDataAttributesTotalInbound95thPercentileMbps(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficDataAttributesTotalInboundGb(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficDataAttributesTotalOutbound95thPercentileMbps(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficDataAttributesTotalOutboundGb(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTrafficDataAttributesFromDate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	fromDateDescription := `The start timestamp. Must be a unix timestamp`

	var fromDateFlagName string
	if cmdPrefix == "" {
		fromDateFlagName = "from_date"
	} else {
		fromDateFlagName = fmt.Sprintf("%v.from_date", cmdPrefix)
	}

	var fromDateFlagDefault int64

	_ = cmd.PersistentFlags().Int64(fromDateFlagName, fromDateFlagDefault, fromDateDescription)

	return nil
}

func registerTrafficDataAttributesRegions(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: regions []*TrafficDataAttributesRegionsItems0 array type is not supported by go-swagger cli yet

	return nil
}

func registerTrafficDataAttributesToDate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	toDateDescription := `The end timestamp. Must be a unix timestamp`

	var toDateFlagName string
	if cmdPrefix == "" {
		toDateFlagName = "to_date"
	} else {
		toDateFlagName = fmt.Sprintf("%v.to_date", cmdPrefix)
	}

	var toDateFlagDefault int64

	_ = cmd.PersistentFlags().Int64(toDateFlagName, toDateFlagDefault, toDateDescription)

	return nil
}

func registerTrafficDataAttributesTotalInbound95thPercentileMbps(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalInbound95thPercentileMbpsDescription := `Value in MBps`

	var totalInbound95thPercentileMbpsFlagName string
	if cmdPrefix == "" {
		totalInbound95thPercentileMbpsFlagName = "total_inbound_95th_percentile_mbps"
	} else {
		totalInbound95thPercentileMbpsFlagName = fmt.Sprintf("%v.total_inbound_95th_percentile_mbps", cmdPrefix)
	}

	var totalInbound95thPercentileMbpsFlagDefault float64

	_ = cmd.PersistentFlags().Float64(totalInbound95thPercentileMbpsFlagName, totalInbound95thPercentileMbpsFlagDefault, totalInbound95thPercentileMbpsDescription)

	return nil
}

func registerTrafficDataAttributesTotalInboundGb(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalInboundGbDescription := `Value in GB`

	var totalInboundGbFlagName string
	if cmdPrefix == "" {
		totalInboundGbFlagName = "total_inbound_gb"
	} else {
		totalInboundGbFlagName = fmt.Sprintf("%v.total_inbound_gb", cmdPrefix)
	}

	var totalInboundGbFlagDefault int64

	_ = cmd.PersistentFlags().Int64(totalInboundGbFlagName, totalInboundGbFlagDefault, totalInboundGbDescription)

	return nil
}

func registerTrafficDataAttributesTotalOutbound95thPercentileMbps(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalOutbound95thPercentileMbpsDescription := `Value in MBps`

	var totalOutbound95thPercentileMbpsFlagName string
	if cmdPrefix == "" {
		totalOutbound95thPercentileMbpsFlagName = "total_outbound_95th_percentile_mbps"
	} else {
		totalOutbound95thPercentileMbpsFlagName = fmt.Sprintf("%v.total_outbound_95th_percentile_mbps", cmdPrefix)
	}

	var totalOutbound95thPercentileMbpsFlagDefault float64

	_ = cmd.PersistentFlags().Float64(totalOutbound95thPercentileMbpsFlagName, totalOutbound95thPercentileMbpsFlagDefault, totalOutbound95thPercentileMbpsDescription)

	return nil
}

func registerTrafficDataAttributesTotalOutboundGb(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalOutboundGbDescription := `Value in GB`

	var totalOutboundGbFlagName string
	if cmdPrefix == "" {
		totalOutboundGbFlagName = "total_outbound_gb"
	} else {
		totalOutboundGbFlagName = fmt.Sprintf("%v.total_outbound_gb", cmdPrefix)
	}

	var totalOutboundGbFlagDefault int64

	_ = cmd.PersistentFlags().Int64(totalOutboundGbFlagName, totalOutboundGbFlagDefault, totalOutboundGbDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTrafficDataAttributesFlags(depth int, m *models.TrafficDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, fromDateAdded := retrieveTrafficDataAttributesFromDateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fromDateAdded

	err, regionsAdded := retrieveTrafficDataAttributesRegionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || regionsAdded

	err, toDateAdded := retrieveTrafficDataAttributesToDateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || toDateAdded

	err, totalInbound95thPercentileMbpsAdded := retrieveTrafficDataAttributesTotalInbound95thPercentileMbpsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalInbound95thPercentileMbpsAdded

	err, totalInboundGbAdded := retrieveTrafficDataAttributesTotalInboundGbFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalInboundGbAdded

	err, totalOutbound95thPercentileMbpsAdded := retrieveTrafficDataAttributesTotalOutbound95thPercentileMbpsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalOutbound95thPercentileMbpsAdded

	err, totalOutboundGbAdded := retrieveTrafficDataAttributesTotalOutboundGbFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalOutboundGbAdded

	return nil, retAdded
}

func retrieveTrafficDataAttributesFromDateFlags(depth int, m *models.TrafficDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	fromDateFlagName := fmt.Sprintf("%v.from_date", cmdPrefix)
	if cmd.Flags().Changed(fromDateFlagName) {

		var fromDateFlagName string
		if cmdPrefix == "" {
			fromDateFlagName = "from_date"
		} else {
			fromDateFlagName = fmt.Sprintf("%v.from_date", cmdPrefix)
		}

		fromDateFlagValue, err := cmd.Flags().GetInt64(fromDateFlagName)
		if err != nil {
			return err, false
		}
		m.FromDate = fromDateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficDataAttributesRegionsFlags(depth int, m *models.TrafficDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	regionsFlagName := fmt.Sprintf("%v.regions", cmdPrefix)
	if cmd.Flags().Changed(regionsFlagName) {
		// warning: regions array type []*TrafficDataAttributesRegionsItems0 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveTrafficDataAttributesToDateFlags(depth int, m *models.TrafficDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	toDateFlagName := fmt.Sprintf("%v.to_date", cmdPrefix)
	if cmd.Flags().Changed(toDateFlagName) {

		var toDateFlagName string
		if cmdPrefix == "" {
			toDateFlagName = "to_date"
		} else {
			toDateFlagName = fmt.Sprintf("%v.to_date", cmdPrefix)
		}

		toDateFlagValue, err := cmd.Flags().GetInt64(toDateFlagName)
		if err != nil {
			return err, false
		}
		m.ToDate = toDateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficDataAttributesTotalInbound95thPercentileMbpsFlags(depth int, m *models.TrafficDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalInbound95thPercentileMbpsFlagName := fmt.Sprintf("%v.total_inbound_95th_percentile_mbps", cmdPrefix)
	if cmd.Flags().Changed(totalInbound95thPercentileMbpsFlagName) {

		var totalInbound95thPercentileMbpsFlagName string
		if cmdPrefix == "" {
			totalInbound95thPercentileMbpsFlagName = "total_inbound_95th_percentile_mbps"
		} else {
			totalInbound95thPercentileMbpsFlagName = fmt.Sprintf("%v.total_inbound_95th_percentile_mbps", cmdPrefix)
		}

		totalInbound95thPercentileMbpsFlagValue, err := cmd.Flags().GetFloat64(totalInbound95thPercentileMbpsFlagName)
		if err != nil {
			return err, false
		}
		m.TotalInbound95thPercentileMbps = totalInbound95thPercentileMbpsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficDataAttributesTotalInboundGbFlags(depth int, m *models.TrafficDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalInboundGbFlagName := fmt.Sprintf("%v.total_inbound_gb", cmdPrefix)
	if cmd.Flags().Changed(totalInboundGbFlagName) {

		var totalInboundGbFlagName string
		if cmdPrefix == "" {
			totalInboundGbFlagName = "total_inbound_gb"
		} else {
			totalInboundGbFlagName = fmt.Sprintf("%v.total_inbound_gb", cmdPrefix)
		}

		totalInboundGbFlagValue, err := cmd.Flags().GetInt64(totalInboundGbFlagName)
		if err != nil {
			return err, false
		}
		m.TotalInboundGb = totalInboundGbFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficDataAttributesTotalOutbound95thPercentileMbpsFlags(depth int, m *models.TrafficDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalOutbound95thPercentileMbpsFlagName := fmt.Sprintf("%v.total_outbound_95th_percentile_mbps", cmdPrefix)
	if cmd.Flags().Changed(totalOutbound95thPercentileMbpsFlagName) {

		var totalOutbound95thPercentileMbpsFlagName string
		if cmdPrefix == "" {
			totalOutbound95thPercentileMbpsFlagName = "total_outbound_95th_percentile_mbps"
		} else {
			totalOutbound95thPercentileMbpsFlagName = fmt.Sprintf("%v.total_outbound_95th_percentile_mbps", cmdPrefix)
		}

		totalOutbound95thPercentileMbpsFlagValue, err := cmd.Flags().GetFloat64(totalOutbound95thPercentileMbpsFlagName)
		if err != nil {
			return err, false
		}
		m.TotalOutbound95thPercentileMbps = totalOutbound95thPercentileMbpsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficDataAttributesTotalOutboundGbFlags(depth int, m *models.TrafficDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalOutboundGbFlagName := fmt.Sprintf("%v.total_outbound_gb", cmdPrefix)
	if cmd.Flags().Changed(totalOutboundGbFlagName) {

		var totalOutboundGbFlagName string
		if cmdPrefix == "" {
			totalOutboundGbFlagName = "total_outbound_gb"
		} else {
			totalOutboundGbFlagName = fmt.Sprintf("%v.total_outbound_gb", cmdPrefix)
		}

		totalOutboundGbFlagValue, err := cmd.Flags().GetInt64(totalOutboundGbFlagName)
		if err != nil {
			return err, false
		}
		m.TotalOutboundGb = totalOutboundGbFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for TrafficDataAttributesRegionsItems0

// register flags to command
func registerModelTrafficDataAttributesRegionsItems0Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTrafficDataAttributesRegionsItems0Data(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficDataAttributesRegionsItems0RegionSlug(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficDataAttributesRegionsItems0TotalInbound95thPercentileMbps(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficDataAttributesRegionsItems0TotalInboundGb(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficDataAttributesRegionsItems0TotalOutbound95thPercentileMbps(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficDataAttributesRegionsItems0TotalOutboundGb(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTrafficDataAttributesRegionsItems0Data(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data []*TrafficDataAttributesRegionsItems0DataItems0 array type is not supported by go-swagger cli yet

	return nil
}

func registerTrafficDataAttributesRegionsItems0RegionSlug(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	regionSlugDescription := ``

	var regionSlugFlagName string
	if cmdPrefix == "" {
		regionSlugFlagName = "region_slug"
	} else {
		regionSlugFlagName = fmt.Sprintf("%v.region_slug", cmdPrefix)
	}

	var regionSlugFlagDefault string

	_ = cmd.PersistentFlags().String(regionSlugFlagName, regionSlugFlagDefault, regionSlugDescription)

	return nil
}

func registerTrafficDataAttributesRegionsItems0TotalInbound95thPercentileMbps(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalInbound95thPercentileMbpsDescription := `Value in MBps`

	var totalInbound95thPercentileMbpsFlagName string
	if cmdPrefix == "" {
		totalInbound95thPercentileMbpsFlagName = "total_inbound_95th_percentile_mbps"
	} else {
		totalInbound95thPercentileMbpsFlagName = fmt.Sprintf("%v.total_inbound_95th_percentile_mbps", cmdPrefix)
	}

	var totalInbound95thPercentileMbpsFlagDefault float64

	_ = cmd.PersistentFlags().Float64(totalInbound95thPercentileMbpsFlagName, totalInbound95thPercentileMbpsFlagDefault, totalInbound95thPercentileMbpsDescription)

	return nil
}

func registerTrafficDataAttributesRegionsItems0TotalInboundGb(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalInboundGbDescription := `Value in GB`

	var totalInboundGbFlagName string
	if cmdPrefix == "" {
		totalInboundGbFlagName = "total_inbound_gb"
	} else {
		totalInboundGbFlagName = fmt.Sprintf("%v.total_inbound_gb", cmdPrefix)
	}

	var totalInboundGbFlagDefault int64

	_ = cmd.PersistentFlags().Int64(totalInboundGbFlagName, totalInboundGbFlagDefault, totalInboundGbDescription)

	return nil
}

func registerTrafficDataAttributesRegionsItems0TotalOutbound95thPercentileMbps(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalOutbound95thPercentileMbpsDescription := `Value in MBps`

	var totalOutbound95thPercentileMbpsFlagName string
	if cmdPrefix == "" {
		totalOutbound95thPercentileMbpsFlagName = "total_outbound_95th_percentile_mbps"
	} else {
		totalOutbound95thPercentileMbpsFlagName = fmt.Sprintf("%v.total_outbound_95th_percentile_mbps", cmdPrefix)
	}

	var totalOutbound95thPercentileMbpsFlagDefault float64

	_ = cmd.PersistentFlags().Float64(totalOutbound95thPercentileMbpsFlagName, totalOutbound95thPercentileMbpsFlagDefault, totalOutbound95thPercentileMbpsDescription)

	return nil
}

func registerTrafficDataAttributesRegionsItems0TotalOutboundGb(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalOutboundGbDescription := `Value in GB`

	var totalOutboundGbFlagName string
	if cmdPrefix == "" {
		totalOutboundGbFlagName = "total_outbound_gb"
	} else {
		totalOutboundGbFlagName = fmt.Sprintf("%v.total_outbound_gb", cmdPrefix)
	}

	var totalOutboundGbFlagDefault int64

	_ = cmd.PersistentFlags().Int64(totalOutboundGbFlagName, totalOutboundGbFlagDefault, totalOutboundGbDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTrafficDataAttributesRegionsItems0Flags(depth int, m *models.TrafficDataAttributesRegionsItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveTrafficDataAttributesRegionsItems0DataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	err, regionSlugAdded := retrieveTrafficDataAttributesRegionsItems0RegionSlugFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || regionSlugAdded

	err, totalInbound95thPercentileMbpsAdded := retrieveTrafficDataAttributesRegionsItems0TotalInbound95thPercentileMbpsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalInbound95thPercentileMbpsAdded

	err, totalInboundGbAdded := retrieveTrafficDataAttributesRegionsItems0TotalInboundGbFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalInboundGbAdded

	err, totalOutbound95thPercentileMbpsAdded := retrieveTrafficDataAttributesRegionsItems0TotalOutbound95thPercentileMbpsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalOutbound95thPercentileMbpsAdded

	err, totalOutboundGbAdded := retrieveTrafficDataAttributesRegionsItems0TotalOutboundGbFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalOutboundGbAdded

	return nil, retAdded
}

func retrieveTrafficDataAttributesRegionsItems0DataFlags(depth int, m *models.TrafficDataAttributesRegionsItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// warning: data array type []*TrafficDataAttributesRegionsItems0DataItems0 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveTrafficDataAttributesRegionsItems0RegionSlugFlags(depth int, m *models.TrafficDataAttributesRegionsItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	regionSlugFlagName := fmt.Sprintf("%v.region_slug", cmdPrefix)
	if cmd.Flags().Changed(regionSlugFlagName) {

		var regionSlugFlagName string
		if cmdPrefix == "" {
			regionSlugFlagName = "region_slug"
		} else {
			regionSlugFlagName = fmt.Sprintf("%v.region_slug", cmdPrefix)
		}

		regionSlugFlagValue, err := cmd.Flags().GetString(regionSlugFlagName)
		if err != nil {
			return err, false
		}
		m.RegionSlug = regionSlugFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficDataAttributesRegionsItems0TotalInbound95thPercentileMbpsFlags(depth int, m *models.TrafficDataAttributesRegionsItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalInbound95thPercentileMbpsFlagName := fmt.Sprintf("%v.total_inbound_95th_percentile_mbps", cmdPrefix)
	if cmd.Flags().Changed(totalInbound95thPercentileMbpsFlagName) {

		var totalInbound95thPercentileMbpsFlagName string
		if cmdPrefix == "" {
			totalInbound95thPercentileMbpsFlagName = "total_inbound_95th_percentile_mbps"
		} else {
			totalInbound95thPercentileMbpsFlagName = fmt.Sprintf("%v.total_inbound_95th_percentile_mbps", cmdPrefix)
		}

		totalInbound95thPercentileMbpsFlagValue, err := cmd.Flags().GetFloat64(totalInbound95thPercentileMbpsFlagName)
		if err != nil {
			return err, false
		}
		m.TotalInbound95thPercentileMbps = totalInbound95thPercentileMbpsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficDataAttributesRegionsItems0TotalInboundGbFlags(depth int, m *models.TrafficDataAttributesRegionsItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalInboundGbFlagName := fmt.Sprintf("%v.total_inbound_gb", cmdPrefix)
	if cmd.Flags().Changed(totalInboundGbFlagName) {

		var totalInboundGbFlagName string
		if cmdPrefix == "" {
			totalInboundGbFlagName = "total_inbound_gb"
		} else {
			totalInboundGbFlagName = fmt.Sprintf("%v.total_inbound_gb", cmdPrefix)
		}

		totalInboundGbFlagValue, err := cmd.Flags().GetInt64(totalInboundGbFlagName)
		if err != nil {
			return err, false
		}
		m.TotalInboundGb = totalInboundGbFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficDataAttributesRegionsItems0TotalOutbound95thPercentileMbpsFlags(depth int, m *models.TrafficDataAttributesRegionsItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalOutbound95thPercentileMbpsFlagName := fmt.Sprintf("%v.total_outbound_95th_percentile_mbps", cmdPrefix)
	if cmd.Flags().Changed(totalOutbound95thPercentileMbpsFlagName) {

		var totalOutbound95thPercentileMbpsFlagName string
		if cmdPrefix == "" {
			totalOutbound95thPercentileMbpsFlagName = "total_outbound_95th_percentile_mbps"
		} else {
			totalOutbound95thPercentileMbpsFlagName = fmt.Sprintf("%v.total_outbound_95th_percentile_mbps", cmdPrefix)
		}

		totalOutbound95thPercentileMbpsFlagValue, err := cmd.Flags().GetFloat64(totalOutbound95thPercentileMbpsFlagName)
		if err != nil {
			return err, false
		}
		m.TotalOutbound95thPercentileMbps = totalOutbound95thPercentileMbpsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficDataAttributesRegionsItems0TotalOutboundGbFlags(depth int, m *models.TrafficDataAttributesRegionsItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalOutboundGbFlagName := fmt.Sprintf("%v.total_outbound_gb", cmdPrefix)
	if cmd.Flags().Changed(totalOutboundGbFlagName) {

		var totalOutboundGbFlagName string
		if cmdPrefix == "" {
			totalOutboundGbFlagName = "total_outbound_gb"
		} else {
			totalOutboundGbFlagName = fmt.Sprintf("%v.total_outbound_gb", cmdPrefix)
		}

		totalOutboundGbFlagValue, err := cmd.Flags().GetInt64(totalOutboundGbFlagName)
		if err != nil {
			return err, false
		}
		m.TotalOutboundGb = totalOutboundGbFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for TrafficDataAttributesRegionsItems0DataItems0

// register flags to command
func registerModelTrafficDataAttributesRegionsItems0DataItems0Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTrafficDataAttributesRegionsItems0DataItems0AvgInboundSpeedMbps(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficDataAttributesRegionsItems0DataItems0AvgOutboundSpeedMbps(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficDataAttributesRegionsItems0DataItems0Date(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficDataAttributesRegionsItems0DataItems0InboundGb(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficDataAttributesRegionsItems0DataItems0OutboundGb(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTrafficDataAttributesRegionsItems0DataItems0AvgInboundSpeedMbps(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	avgInboundSpeedMbpsDescription := `Value in MBps`

	var avgInboundSpeedMbpsFlagName string
	if cmdPrefix == "" {
		avgInboundSpeedMbpsFlagName = "avg_inbound_speed_mbps"
	} else {
		avgInboundSpeedMbpsFlagName = fmt.Sprintf("%v.avg_inbound_speed_mbps", cmdPrefix)
	}

	var avgInboundSpeedMbpsFlagDefault float64

	_ = cmd.PersistentFlags().Float64(avgInboundSpeedMbpsFlagName, avgInboundSpeedMbpsFlagDefault, avgInboundSpeedMbpsDescription)

	return nil
}

func registerTrafficDataAttributesRegionsItems0DataItems0AvgOutboundSpeedMbps(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	avgOutboundSpeedMbpsDescription := `Value in MBps`

	var avgOutboundSpeedMbpsFlagName string
	if cmdPrefix == "" {
		avgOutboundSpeedMbpsFlagName = "avg_outbound_speed_mbps"
	} else {
		avgOutboundSpeedMbpsFlagName = fmt.Sprintf("%v.avg_outbound_speed_mbps", cmdPrefix)
	}

	var avgOutboundSpeedMbpsFlagDefault float64

	_ = cmd.PersistentFlags().Float64(avgOutboundSpeedMbpsFlagName, avgOutboundSpeedMbpsFlagDefault, avgOutboundSpeedMbpsDescription)

	return nil
}

func registerTrafficDataAttributesRegionsItems0DataItems0Date(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	dateDescription := `The datetime of the day`

	var dateFlagName string
	if cmdPrefix == "" {
		dateFlagName = "date"
	} else {
		dateFlagName = fmt.Sprintf("%v.date", cmdPrefix)
	}

	var dateFlagDefault string

	_ = cmd.PersistentFlags().String(dateFlagName, dateFlagDefault, dateDescription)

	return nil
}

func registerTrafficDataAttributesRegionsItems0DataItems0InboundGb(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	inboundGbDescription := `Value in GB`

	var inboundGbFlagName string
	if cmdPrefix == "" {
		inboundGbFlagName = "inbound_gb"
	} else {
		inboundGbFlagName = fmt.Sprintf("%v.inbound_gb", cmdPrefix)
	}

	var inboundGbFlagDefault int64

	_ = cmd.PersistentFlags().Int64(inboundGbFlagName, inboundGbFlagDefault, inboundGbDescription)

	return nil
}

func registerTrafficDataAttributesRegionsItems0DataItems0OutboundGb(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	outboundGbDescription := `Value in GB`

	var outboundGbFlagName string
	if cmdPrefix == "" {
		outboundGbFlagName = "outbound_gb"
	} else {
		outboundGbFlagName = fmt.Sprintf("%v.outbound_gb", cmdPrefix)
	}

	var outboundGbFlagDefault int64

	_ = cmd.PersistentFlags().Int64(outboundGbFlagName, outboundGbFlagDefault, outboundGbDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTrafficDataAttributesRegionsItems0DataItems0Flags(depth int, m *models.TrafficDataAttributesRegionsItems0DataItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, avgInboundSpeedMbpsAdded := retrieveTrafficDataAttributesRegionsItems0DataItems0AvgInboundSpeedMbpsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || avgInboundSpeedMbpsAdded

	err, avgOutboundSpeedMbpsAdded := retrieveTrafficDataAttributesRegionsItems0DataItems0AvgOutboundSpeedMbpsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || avgOutboundSpeedMbpsAdded

	err, dateAdded := retrieveTrafficDataAttributesRegionsItems0DataItems0DateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dateAdded

	err, inboundGbAdded := retrieveTrafficDataAttributesRegionsItems0DataItems0InboundGbFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || inboundGbAdded

	err, outboundGbAdded := retrieveTrafficDataAttributesRegionsItems0DataItems0OutboundGbFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || outboundGbAdded

	return nil, retAdded
}

func retrieveTrafficDataAttributesRegionsItems0DataItems0AvgInboundSpeedMbpsFlags(depth int, m *models.TrafficDataAttributesRegionsItems0DataItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	avgInboundSpeedMbpsFlagName := fmt.Sprintf("%v.avg_inbound_speed_mbps", cmdPrefix)
	if cmd.Flags().Changed(avgInboundSpeedMbpsFlagName) {

		var avgInboundSpeedMbpsFlagName string
		if cmdPrefix == "" {
			avgInboundSpeedMbpsFlagName = "avg_inbound_speed_mbps"
		} else {
			avgInboundSpeedMbpsFlagName = fmt.Sprintf("%v.avg_inbound_speed_mbps", cmdPrefix)
		}

		avgInboundSpeedMbpsFlagValue, err := cmd.Flags().GetFloat64(avgInboundSpeedMbpsFlagName)
		if err != nil {
			return err, false
		}
		m.AvgInboundSpeedMbps = avgInboundSpeedMbpsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficDataAttributesRegionsItems0DataItems0AvgOutboundSpeedMbpsFlags(depth int, m *models.TrafficDataAttributesRegionsItems0DataItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	avgOutboundSpeedMbpsFlagName := fmt.Sprintf("%v.avg_outbound_speed_mbps", cmdPrefix)
	if cmd.Flags().Changed(avgOutboundSpeedMbpsFlagName) {

		var avgOutboundSpeedMbpsFlagName string
		if cmdPrefix == "" {
			avgOutboundSpeedMbpsFlagName = "avg_outbound_speed_mbps"
		} else {
			avgOutboundSpeedMbpsFlagName = fmt.Sprintf("%v.avg_outbound_speed_mbps", cmdPrefix)
		}

		avgOutboundSpeedMbpsFlagValue, err := cmd.Flags().GetFloat64(avgOutboundSpeedMbpsFlagName)
		if err != nil {
			return err, false
		}
		m.AvgOutboundSpeedMbps = avgOutboundSpeedMbpsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficDataAttributesRegionsItems0DataItems0DateFlags(depth int, m *models.TrafficDataAttributesRegionsItems0DataItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dateFlagName := fmt.Sprintf("%v.date", cmdPrefix)
	if cmd.Flags().Changed(dateFlagName) {

		var dateFlagName string
		if cmdPrefix == "" {
			dateFlagName = "date"
		} else {
			dateFlagName = fmt.Sprintf("%v.date", cmdPrefix)
		}

		dateFlagValue, err := cmd.Flags().GetString(dateFlagName)
		if err != nil {
			return err, false
		}
		m.Date = dateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficDataAttributesRegionsItems0DataItems0InboundGbFlags(depth int, m *models.TrafficDataAttributesRegionsItems0DataItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	inboundGbFlagName := fmt.Sprintf("%v.inbound_gb", cmdPrefix)
	if cmd.Flags().Changed(inboundGbFlagName) {

		var inboundGbFlagName string
		if cmdPrefix == "" {
			inboundGbFlagName = "inbound_gb"
		} else {
			inboundGbFlagName = fmt.Sprintf("%v.inbound_gb", cmdPrefix)
		}

		inboundGbFlagValue, err := cmd.Flags().GetInt64(inboundGbFlagName)
		if err != nil {
			return err, false
		}
		m.InboundGb = inboundGbFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficDataAttributesRegionsItems0DataItems0OutboundGbFlags(depth int, m *models.TrafficDataAttributesRegionsItems0DataItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	outboundGbFlagName := fmt.Sprintf("%v.outbound_gb", cmdPrefix)
	if cmd.Flags().Changed(outboundGbFlagName) {

		var outboundGbFlagName string
		if cmdPrefix == "" {
			outboundGbFlagName = "outbound_gb"
		} else {
			outboundGbFlagName = fmt.Sprintf("%v.outbound_gb", cmdPrefix)
		}

		outboundGbFlagValue, err := cmd.Flags().GetInt64(outboundGbFlagName)
		if err != nil {
			return err, false
		}
		m.OutboundGb = outboundGbFlagValue

		retAdded = true
	}

	return nil, retAdded
}
