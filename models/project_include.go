// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectInclude project include
//
// swagger:model project_include
type ProjectInclude struct {

	// bandwidth alert
	BandwidthAlert bool `json:"bandwidth_alert,omitempty"`

	// billing
	Billing *ProjectIncludeBilling `json:"billing,omitempty"`

	// billing method
	BillingMethod *string `json:"billing_method,omitempty"`

	// billing type
	BillingType *string `json:"billing_type,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// environment
	Environment *string `json:"environment,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// stats
	Stats *ProjectIncludeStats `json:"stats,omitempty"`
}

// Validate validates this project include
func (m *ProjectInclude) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBilling(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectInclude) validateBilling(formats strfmt.Registry) error {
	if swag.IsZero(m.Billing) { // not required
		return nil
	}

	if m.Billing != nil {
		if err := m.Billing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billing")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectInclude) validateStats(formats strfmt.Registry) error {
	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this project include based on the context it is used
func (m *ProjectInclude) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBilling(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectInclude) contextValidateBilling(ctx context.Context, formats strfmt.Registry) error {

	if m.Billing != nil {

		if swag.IsZero(m.Billing) { // not required
			return nil
		}

		if err := m.Billing.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billing")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectInclude) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

	if m.Stats != nil {

		if swag.IsZero(m.Stats) { // not required
			return nil
		}

		if err := m.Stats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectInclude) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectInclude) UnmarshalBinary(b []byte) error {
	var res ProjectInclude
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ProjectIncludeBilling project include billing
//
// swagger:model ProjectIncludeBilling
type ProjectIncludeBilling struct {

	// method
	Method string `json:"method,omitempty"`

	// subscription id
	SubscriptionID string `json:"subscription_id,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this project include billing
func (m *ProjectIncludeBilling) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project include billing based on context it is used
func (m *ProjectIncludeBilling) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectIncludeBilling) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectIncludeBilling) UnmarshalBinary(b []byte) error {
	var res ProjectIncludeBilling
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ProjectIncludeStats project include stats
//
// swagger:model ProjectIncludeStats
type ProjectIncludeStats struct {

	// ip addresses
	IPAddresses int64 `json:"ip_addresses,omitempty"`

	// prefixes
	Prefixes int64 `json:"prefixes,omitempty"`

	// servers
	Servers int64 `json:"servers,omitempty"`

	// vlans
	Vlans int64 `json:"vlans,omitempty"`
}

// Validate validates this project include stats
func (m *ProjectIncludeStats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project include stats based on context it is used
func (m *ProjectIncludeStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectIncludeStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectIncludeStats) UnmarshalBinary(b []byte) error {
	var res ProjectIncludeStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
