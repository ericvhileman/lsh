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

// SSHKeyData ssh key data
//
// swagger:model ssh_key_data
type SSHKeyData struct {

	// attributes
	Attributes *SSHKeyDataAttributes `json:"attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// type
	// Required: true
	// Enum: [ssh_keys]
	Type *string `json:"type"`
}

// Validate validates this ssh key data
func (m *SSHKeyData) Validate(formats strfmt.Registry) error {
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

func (m *SSHKeyData) validateAttributes(formats strfmt.Registry) error {
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

var sshKeyDataTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ssh_keys"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sshKeyDataTypeTypePropEnum = append(sshKeyDataTypeTypePropEnum, v)
	}
}

const (

	// SSHKeyDataTypeSSHKeys captures enum value "ssh_keys"
	SSHKeyDataTypeSSHKeys string = "ssh_keys"
)

// prop value enum
func (m *SSHKeyData) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sshKeyDataTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SSHKeyData) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ssh key data based on the context it is used
func (m *SSHKeyData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SSHKeyData) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SSHKeyData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSHKeyData) UnmarshalBinary(b []byte) error {
	var res SSHKeyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SSHKeyDataAttributes SSH key data attributes
//
// swagger:model SSHKeyDataAttributes
type SSHKeyDataAttributes struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// SSH Key fingerprint
	Fingerprint string `json:"fingerprint,omitempty"`

	// Name of the SSH Key
	Name string `json:"name,omitempty"`

	// SSH Public Key
	PublicKey string `json:"public_key,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// user
	User *UserInclude `json:"user,omitempty"`
}

// Validate validates this SSH key data attributes
func (m *SSHKeyDataAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SSHKeyDataAttributes) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this SSH key data attributes based on the context it is used
func (m *SSHKeyDataAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SSHKeyDataAttributes) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {

		if swag.IsZero(m.User) { // not required
			return nil
		}

		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SSHKeyDataAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSHKeyDataAttributes) UnmarshalBinary(b []byte) error {
	var res SSHKeyDataAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
