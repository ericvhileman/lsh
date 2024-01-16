// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BillingUsage billing usage
//
// swagger:model billing_usage
type BillingUsage struct {

	// data
	Data *BillingUsageData `json:"data,omitempty"`
}

// Validate validates this billing usage
func (m *BillingUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingUsage) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this billing usage based on the context it is used
func (m *BillingUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingUsage) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {

		if swag.IsZero(m.Data) { // not required
			return nil
		}

		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingUsage) UnmarshalBinary(b []byte) error {
	var res BillingUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BillingUsageData billing usage data
//
// swagger:model BillingUsageData
type BillingUsageData struct {

	// attributes
	Attributes *BillingUsageDataAttributes `json:"attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this billing usage data
func (m *BillingUsageData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingUsageData) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "attributes")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this billing usage data based on the context it is used
func (m *BillingUsageData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingUsageData) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	if m.Attributes != nil {

		if swag.IsZero(m.Attributes) { // not required
			return nil
		}

		if err := m.Attributes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "attributes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingUsageData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingUsageData) UnmarshalBinary(b []byte) error {
	var res BillingUsageData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BillingUsageDataAttributes billing usage data attributes
//
// swagger:model BillingUsageDataAttributes
type BillingUsageDataAttributes struct {

	// period
	Period *BillingUsageDataAttributesPeriod `json:"period,omitempty"`

	// The total usage price in cents
	Price float64 `json:"price,omitempty"`

	// products
	Products []*BillingUsageDataAttributesProductsItems0 `json:"products"`

	// project
	Project *BillingUsageDataAttributesProject `json:"project,omitempty"`
}

// Validate validates this billing usage data attributes
func (m *BillingUsageDataAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProducts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingUsageDataAttributes) validatePeriod(formats strfmt.Registry) error {
	if swag.IsZero(m.Period) { // not required
		return nil
	}

	if m.Period != nil {
		if err := m.Period.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "attributes" + "." + "period")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "attributes" + "." + "period")
			}
			return err
		}
	}

	return nil
}

func (m *BillingUsageDataAttributes) validateProducts(formats strfmt.Registry) error {
	if swag.IsZero(m.Products) { // not required
		return nil
	}

	for i := 0; i < len(m.Products); i++ {
		if swag.IsZero(m.Products[i]) { // not required
			continue
		}

		if m.Products[i] != nil {
			if err := m.Products[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "attributes" + "." + "products" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "attributes" + "." + "products" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BillingUsageDataAttributes) validateProject(formats strfmt.Registry) error {
	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "attributes" + "." + "project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "attributes" + "." + "project")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this billing usage data attributes based on the context it is used
func (m *BillingUsageDataAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePeriod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProducts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingUsageDataAttributes) contextValidatePeriod(ctx context.Context, formats strfmt.Registry) error {

	if m.Period != nil {

		if swag.IsZero(m.Period) { // not required
			return nil
		}

		if err := m.Period.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "attributes" + "." + "period")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "attributes" + "." + "period")
			}
			return err
		}
	}

	return nil
}

func (m *BillingUsageDataAttributes) contextValidateProducts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Products); i++ {

		if m.Products[i] != nil {

			if swag.IsZero(m.Products[i]) { // not required
				return nil
			}

			if err := m.Products[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "attributes" + "." + "products" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "attributes" + "." + "products" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BillingUsageDataAttributes) contextValidateProject(ctx context.Context, formats strfmt.Registry) error {

	if m.Project != nil {

		if swag.IsZero(m.Project) { // not required
			return nil
		}

		if err := m.Project.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "attributes" + "." + "project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "attributes" + "." + "project")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingUsageDataAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingUsageDataAttributes) UnmarshalBinary(b []byte) error {
	var res BillingUsageDataAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BillingUsageDataAttributesPeriod The period from the returned billing cycle
//
// swagger:model BillingUsageDataAttributesPeriod
type BillingUsageDataAttributesPeriod struct {

	// end
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`

	// start
	// Format: date-time
	Start strfmt.DateTime `json:"start,omitempty"`
}

// Validate validates this billing usage data attributes period
func (m *BillingUsageDataAttributesPeriod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingUsageDataAttributesPeriod) validateEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("data"+"."+"attributes"+"."+"period"+"."+"end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BillingUsageDataAttributesPeriod) validateStart(formats strfmt.Registry) error {
	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("data"+"."+"attributes"+"."+"period"+"."+"start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this billing usage data attributes period based on context it is used
func (m *BillingUsageDataAttributesPeriod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BillingUsageDataAttributesPeriod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingUsageDataAttributesPeriod) UnmarshalBinary(b []byte) error {
	var res BillingUsageDataAttributesPeriod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BillingUsageDataAttributesProductsItems0 billing usage data attributes products items0
//
// swagger:model BillingUsageDataAttributesProductsItems0
type BillingUsageDataAttributesProductsItems0 struct {

	// end
	// Format: date-time
	End *strfmt.DateTime `json:"end,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// metadata
	Metadata *BillingUsageDataAttributesProductsItems0Metadata `json:"metadata,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The total usage price of the product in cents
	Price float64 `json:"price,omitempty"`

	// quantity
	Quantity float64 `json:"quantity,omitempty"`

	// resource
	// Enum: [servers networking others]
	Resource string `json:"resource,omitempty"`

	// start
	// Format: date-time
	Start strfmt.DateTime `json:"start,omitempty"`

	// The unit price of the product in cents
	UnitPrice float64 `json:"unit_price,omitempty"`

	// usage type
	// Enum: [licensed metered]
	UsageType string `json:"usage_type,omitempty"`
}

// Validate validates this billing usage data attributes products items0
func (m *BillingUsageDataAttributesProductsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsageType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingUsageDataAttributesProductsItems0) validateEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BillingUsageDataAttributesProductsItems0) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

var billingUsageDataAttributesProductsItems0TypeResourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["servers","networking","others"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billingUsageDataAttributesProductsItems0TypeResourcePropEnum = append(billingUsageDataAttributesProductsItems0TypeResourcePropEnum, v)
	}
}

const (

	// BillingUsageDataAttributesProductsItems0ResourceServers captures enum value "servers"
	BillingUsageDataAttributesProductsItems0ResourceServers string = "servers"

	// BillingUsageDataAttributesProductsItems0ResourceNetworking captures enum value "networking"
	BillingUsageDataAttributesProductsItems0ResourceNetworking string = "networking"

	// BillingUsageDataAttributesProductsItems0ResourceOthers captures enum value "others"
	BillingUsageDataAttributesProductsItems0ResourceOthers string = "others"
)

// prop value enum
func (m *BillingUsageDataAttributesProductsItems0) validateResourceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, billingUsageDataAttributesProductsItems0TypeResourcePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BillingUsageDataAttributesProductsItems0) validateResource(formats strfmt.Registry) error {
	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	// value enum
	if err := m.validateResourceEnum("resource", "body", m.Resource); err != nil {
		return err
	}

	return nil
}

func (m *BillingUsageDataAttributesProductsItems0) validateStart(formats strfmt.Registry) error {
	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

var billingUsageDataAttributesProductsItems0TypeUsageTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["licensed","metered"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billingUsageDataAttributesProductsItems0TypeUsageTypePropEnum = append(billingUsageDataAttributesProductsItems0TypeUsageTypePropEnum, v)
	}
}

const (

	// BillingUsageDataAttributesProductsItems0UsageTypeLicensed captures enum value "licensed"
	BillingUsageDataAttributesProductsItems0UsageTypeLicensed string = "licensed"

	// BillingUsageDataAttributesProductsItems0UsageTypeMetered captures enum value "metered"
	BillingUsageDataAttributesProductsItems0UsageTypeMetered string = "metered"
)

// prop value enum
func (m *BillingUsageDataAttributesProductsItems0) validateUsageTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, billingUsageDataAttributesProductsItems0TypeUsageTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BillingUsageDataAttributesProductsItems0) validateUsageType(formats strfmt.Registry) error {
	if swag.IsZero(m.UsageType) { // not required
		return nil
	}

	// value enum
	if err := m.validateUsageTypeEnum("usage_type", "body", m.UsageType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this billing usage data attributes products items0 based on the context it is used
func (m *BillingUsageDataAttributesProductsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingUsageDataAttributesProductsItems0) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {

		if swag.IsZero(m.Metadata) { // not required
			return nil
		}

		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingUsageDataAttributesProductsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingUsageDataAttributesProductsItems0) UnmarshalBinary(b []byte) error {
	var res BillingUsageDataAttributesProductsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BillingUsageDataAttributesProductsItems0Metadata billing usage data attributes products items0 metadata
//
// swagger:model BillingUsageDataAttributesProductsItems0Metadata
type BillingUsageDataAttributesProductsItems0Metadata struct {

	// hostname
	Hostname *string `json:"hostname,omitempty"`

	// id
	ID *string `json:"id,omitempty"`

	// plan
	Plan *string `json:"plan,omitempty"`
}

// Validate validates this billing usage data attributes products items0 metadata
func (m *BillingUsageDataAttributesProductsItems0Metadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this billing usage data attributes products items0 metadata based on context it is used
func (m *BillingUsageDataAttributesProductsItems0Metadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BillingUsageDataAttributesProductsItems0Metadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingUsageDataAttributesProductsItems0Metadata) UnmarshalBinary(b []byte) error {
	var res BillingUsageDataAttributesProductsItems0Metadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BillingUsageDataAttributesProject The project in which the returned usage belongs to
//
// swagger:model BillingUsageDataAttributesProject
type BillingUsageDataAttributesProject struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`
}

// Validate validates this billing usage data attributes project
func (m *BillingUsageDataAttributesProject) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this billing usage data attributes project based on context it is used
func (m *BillingUsageDataAttributesProject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BillingUsageDataAttributesProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingUsageDataAttributesProject) UnmarshalBinary(b []byte) error {
	var res BillingUsageDataAttributesProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
