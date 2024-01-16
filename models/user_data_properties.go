// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserDataProperties user data properties
//
// swagger:model user_data_properties
type UserDataProperties struct {

	// attributes
	Attributes *UserDataPropertiesAttributes `json:"attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// type
	// Required: true
	// Enum: [user_data]
	Type *string `json:"type"`
}

// Validate validates this user data properties
func (m *UserDataProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserDataProperties) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

var userDataPropertiesTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user_data"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userDataPropertiesTypeTypePropEnum = append(userDataPropertiesTypeTypePropEnum, v)
	}
}

const (

	// UserDataPropertiesTypeUserData captures enum value "user_data"
	UserDataPropertiesTypeUserData string = "user_data"
)

// prop value enum
func (m *UserDataProperties) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userDataPropertiesTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserDataProperties) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user data properties based on the context it is used
func (m *UserDataProperties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserDataProperties) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	if m.Attributes != nil {

		if swag.IsZero(m.Attributes) { // not required
			return nil
		}

		if err := m.Attributes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserDataProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserDataProperties) UnmarshalBinary(b []byte) error {
	var res UserDataProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UserDataPropertiesAttributes user data properties attributes
//
// swagger:model UserDataPropertiesAttributes
type UserDataPropertiesAttributes struct {

	// content of the User Data
	Content string `json:"content,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// description of the User Data
	Description string `json:"description,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this user data properties attributes
func (m *UserDataPropertiesAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user data properties attributes based on context it is used
func (m *UserDataPropertiesAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserDataPropertiesAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserDataPropertiesAttributes) UnmarshalBinary(b []byte) error {
	var res UserDataPropertiesAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}