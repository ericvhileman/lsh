// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/latitudesh/cli/internal/api"
	"github.com/latitudesh/cli/models"
)

// PatchCurrentTeamReader is a Reader for the PatchCurrentTeam structure.
type PatchCurrentTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCurrentTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchCurrentTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := api.NewForbidden()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := api.NewNotFound()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /team/{team_id}] patch-current-team", response, response.Code())
	}
}

// NewPatchCurrentTeamOK creates a PatchCurrentTeamOK with default headers values
func NewPatchCurrentTeamOK() *PatchCurrentTeamOK {
	return &PatchCurrentTeamOK{}
}

/*
PatchCurrentTeamOK describes a response with status code 200, with default header values.

Success
*/
type PatchCurrentTeamOK struct {
	Payload *PatchCurrentTeamOKBody
}

// IsSuccess returns true when this patch current team o k response has a 2xx status code
func (o *PatchCurrentTeamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch current team o k response has a 3xx status code
func (o *PatchCurrentTeamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch current team o k response has a 4xx status code
func (o *PatchCurrentTeamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch current team o k response has a 5xx status code
func (o *PatchCurrentTeamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch current team o k response a status code equal to that given
func (o *PatchCurrentTeamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch current team o k response
func (o *PatchCurrentTeamOK) Code() int {
	return 200
}

func (o *PatchCurrentTeamOK) Error() string {
	return fmt.Sprintf("[PATCH /team/{team_id}][%d] patchCurrentTeamOK  %+v", 200, o.Payload)
}

func (o *PatchCurrentTeamOK) String() string {
	return fmt.Sprintf("[PATCH /team/{team_id}][%d] patchCurrentTeamOK  %+v", 200, o.Payload)
}

func (o *PatchCurrentTeamOK) GetPayload() *PatchCurrentTeamOKBody {
	return o.Payload
}

func (o *PatchCurrentTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PatchCurrentTeamOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCurrentTeamForbidden creates a PatchCurrentTeamForbidden with default headers values
func NewPatchCurrentTeamForbidden() *PatchCurrentTeamForbidden {
	return &PatchCurrentTeamForbidden{}
}

/*
PatchCurrentTeamForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchCurrentTeamForbidden struct {
	Payload *models.ErrorObject
}

// IsSuccess returns true when this patch current team forbidden response has a 2xx status code
func (o *PatchCurrentTeamForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch current team forbidden response has a 3xx status code
func (o *PatchCurrentTeamForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch current team forbidden response has a 4xx status code
func (o *PatchCurrentTeamForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch current team forbidden response has a 5xx status code
func (o *PatchCurrentTeamForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch current team forbidden response a status code equal to that given
func (o *PatchCurrentTeamForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch current team forbidden response
func (o *PatchCurrentTeamForbidden) Code() int {
	return 403
}

func (o *PatchCurrentTeamForbidden) Error() string {
	return fmt.Sprintf("[PATCH /team/{team_id}][%d] patchCurrentTeamForbidden  %+v", 403, o.Payload)
}

func (o *PatchCurrentTeamForbidden) String() string {
	return fmt.Sprintf("[PATCH /team/{team_id}][%d] patchCurrentTeamForbidden  %+v", 403, o.Payload)
}

func (o *PatchCurrentTeamForbidden) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *PatchCurrentTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCurrentTeamNotFound creates a PatchCurrentTeamNotFound with default headers values
func NewPatchCurrentTeamNotFound() *PatchCurrentTeamNotFound {
	return &PatchCurrentTeamNotFound{}
}

/*
PatchCurrentTeamNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PatchCurrentTeamNotFound struct {
	Payload *models.ErrorObject
}

// IsSuccess returns true when this patch current team not found response has a 2xx status code
func (o *PatchCurrentTeamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch current team not found response has a 3xx status code
func (o *PatchCurrentTeamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch current team not found response has a 4xx status code
func (o *PatchCurrentTeamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch current team not found response has a 5xx status code
func (o *PatchCurrentTeamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch current team not found response a status code equal to that given
func (o *PatchCurrentTeamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the patch current team not found response
func (o *PatchCurrentTeamNotFound) Code() int {
	return 404
}

func (o *PatchCurrentTeamNotFound) Error() string {
	return fmt.Sprintf("[PATCH /team/{team_id}][%d] patchCurrentTeamNotFound  %+v", 404, o.Payload)
}

func (o *PatchCurrentTeamNotFound) String() string {
	return fmt.Sprintf("[PATCH /team/{team_id}][%d] patchCurrentTeamNotFound  %+v", 404, o.Payload)
}

func (o *PatchCurrentTeamNotFound) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *PatchCurrentTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PatchCurrentTeamBody patch current team body
swagger:model PatchCurrentTeamBody
*/
type PatchCurrentTeamBody struct {

	// data
	// Required: true
	Data *PatchCurrentTeamParamsBodyData `json:"data"`
}

// Validate validates this patch current team body
func (o *PatchCurrentTeamBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchCurrentTeamBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this patch current team body based on the context it is used
func (o *PatchCurrentTeamBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchCurrentTeamBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchCurrentTeamBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchCurrentTeamBody) UnmarshalBinary(b []byte) error {
	var res PatchCurrentTeamBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PatchCurrentTeamOKBody patch current team o k body
swagger:model PatchCurrentTeamOKBody
*/
type PatchCurrentTeamOKBody struct {

	// data
	Data *models.Team `json:"data,omitempty"`
}

// Validate validates this patch current team o k body
func (o *PatchCurrentTeamOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchCurrentTeamOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patchCurrentTeamOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("patchCurrentTeamOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this patch current team o k body based on the context it is used
func (o *PatchCurrentTeamOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchCurrentTeamOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patchCurrentTeamOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("patchCurrentTeamOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchCurrentTeamOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchCurrentTeamOKBody) UnmarshalBinary(b []byte) error {
	var res PatchCurrentTeamOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PatchCurrentTeamParamsBodyData patch current team params body data
swagger:model PatchCurrentTeamParamsBodyData
*/
type PatchCurrentTeamParamsBodyData struct {

	// attributes
	Attributes *PatchCurrentTeamParamsBodyDataAttributes `json:"attributes,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// type
	// Required: true
	// Enum: [teams]
	Type *string `json:"type"`
}

// Validate validates this patch current team params body data
func (o *PatchCurrentTeamParamsBodyData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchCurrentTeamParamsBodyData) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(o.Attributes) { // not required
		return nil
	}

	if o.Attributes != nil {
		if err := o.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data" + "." + "attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data" + "." + "attributes")
			}
			return err
		}
	}

	return nil
}

func (o *PatchCurrentTeamParamsBodyData) validateID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"data"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

var patchCurrentTeamParamsBodyDataTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["teams"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchCurrentTeamParamsBodyDataTypeTypePropEnum = append(patchCurrentTeamParamsBodyDataTypeTypePropEnum, v)
	}
}

const (

	// PatchCurrentTeamParamsBodyDataTypeTeams captures enum value "teams"
	PatchCurrentTeamParamsBodyDataTypeTeams string = "teams"
)

// prop value enum
func (o *PatchCurrentTeamParamsBodyData) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchCurrentTeamParamsBodyDataTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PatchCurrentTeamParamsBodyData) validateType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"data"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("body"+"."+"data"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this patch current team params body data based on the context it is used
func (o *PatchCurrentTeamParamsBodyData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchCurrentTeamParamsBodyData) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	if o.Attributes != nil {

		if swag.IsZero(o.Attributes) { // not required
			return nil
		}

		if err := o.Attributes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data" + "." + "attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data" + "." + "attributes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchCurrentTeamParamsBodyData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchCurrentTeamParamsBodyData) UnmarshalBinary(b []byte) error {
	var res PatchCurrentTeamParamsBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PatchCurrentTeamParamsBodyDataAttributes patch current team params body data attributes
swagger:model PatchCurrentTeamParamsBodyDataAttributes
*/
type PatchCurrentTeamParamsBodyDataAttributes struct {

	// address
	Address string `json:"address,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// referred code
	ReferredCode string `json:"referred_code,omitempty"`
}

// Validate validates this patch current team params body data attributes
func (o *PatchCurrentTeamParamsBodyDataAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch current team params body data attributes based on context it is used
func (o *PatchCurrentTeamParamsBodyDataAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PatchCurrentTeamParamsBodyDataAttributes) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchCurrentTeamParamsBodyDataAttributes) UnmarshalBinary(b []byte) error {
	var res PatchCurrentTeamParamsBodyDataAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
