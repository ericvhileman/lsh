// Code generated by go-swagger; DO NOT EDIT.

package api_keys

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
	"github.com/spf13/viper"

	apierrors "github.com/latitudesh/lsh/internal/api/errors"
	"github.com/latitudesh/lsh/internal/output"
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

// PostAPIKeyReader is a Reader for the PostAPIKey structure.
type PostAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAPIKeyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := apierrors.NewBadRequest()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := apierrors.NewUnauthorized()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := apierrors.NewUnprocessableEntity()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /auth/api_keys] post-api-key", response, response.Code())
	}
}

// NewPostAPIKeyCreated creates a PostAPIKeyCreated with default headers values
func NewPostAPIKeyCreated() *PostAPIKeyCreated {
	return &PostAPIKeyCreated{}
}

/*
PostAPIKeyCreated describes a response with status code 201, with default header values.

API Key Created
*/
type PostAPIKeyCreated struct {
	Payload *PostAPIKeyCreatedBody
}

// IsSuccess returns true when this post Api key created response has a 2xx status code
func (o *PostAPIKeyCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post Api key created response has a 3xx status code
func (o *PostAPIKeyCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Api key created response has a 4xx status code
func (o *PostAPIKeyCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Api key created response has a 5xx status code
func (o *PostAPIKeyCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post Api key created response a status code equal to that given
func (o *PostAPIKeyCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post Api key created response
func (o *PostAPIKeyCreated) Code() int {
	return 201
}

func (o *PostAPIKeyCreated) Error() string {
	return fmt.Sprintf("[POST /auth/api_keys][%d] postApiKeyCreated  %+v", 201, o.Payload)
}

func (o *PostAPIKeyCreated) String() string {
	return fmt.Sprintf("[POST /auth/api_keys][%d] postApiKeyCreated  %+v", 201, o.Payload)
}

func (o *PostAPIKeyCreated) GetPayload() *PostAPIKeyCreatedBody {
	return o.Payload
}

func (o *PostAPIKeyCreated) Render() {
	formatAsJSON := viper.GetBool("json")

	if formatAsJSON {
		o.RenderJSON()
		return
	}

	formatOutputFlag := viper.GetString("output")

	switch formatOutputFlag {
	case "json":
		o.RenderJSON()
	case "table":
		o.RenderTable()
	default:
		fmt.Println("Unsupported output format")
	}
}

func (o *PostAPIKeyCreated) RenderJSON() {
	if !swag.IsZero(o) && !swag.IsZero(o.Payload) {
		JSONString, err := json.Marshal(o.Payload)
		if err != nil {
			fmt.Println("Could not decode the result as JSON.")
		}

		output.RenderJSON(JSONString)
	}
}

type CreateAPIKeyTableRow struct {
	ID string `json:"id,omitempty"`
	TokenLastSlice string `json:"token_last_slice,omitempty"`
	User string `json:"user,omitempty"`
	APIVersion string `json:"api_version,omitempty"`
	LastUsedAt string `json:"last_used_at,omitempty"`
	Name string `json:"name,omitempty"`
}

func (o *PostAPIKeyCreated) RenderTable() {
	attributes := o.Payload.Data.Attributes

	var rows []CreateAPIKeyTableRow

	row := CreateAPIKeyTableRow{
		ID:        	 					 table.String(o.Payload.Data.ID),
		TokenLastSlice:        table.String(attributes.TokenLastSlice),
		User:    							 table.String(attributes.User.Email),
		APIVersion:     			 table.String(attributes.APIVersion),
		LastUsedAt: 					 table.DateTime(attributes.LastUsedAt),
		Name: 				 				 table.String(attributes.Name),
	}

	rows = append(rows, row)

	var interfaceRows []interface{}

	for _, row := range rows {
		interfaceRows = append(interfaceRows, row)
	}

	table.Render(interfaceRows)
}

func (o *PostAPIKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostAPIKeyCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostAPIKeyCreatedBody post API key created body
swagger:model PostAPIKeyCreatedBody
*/
type PostAPIKeyCreatedBody struct {

	// data
	Data *models.APIKey `json:"data,omitempty"`
}

// Validate validates this post API key created body
func (o *PostAPIKeyCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostAPIKeyCreatedBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postApiKeyCreated" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postApiKeyCreated" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post API key created body based on the context it is used
func (o *PostAPIKeyCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostAPIKeyCreatedBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postApiKeyCreated" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postApiKeyCreated" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostAPIKeyCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAPIKeyCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostAPIKeyCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
