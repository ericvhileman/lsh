package virtual_networks

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

	apierrors "github.com/latitudesh/lsh/internal/api/errors"
	"github.com/latitudesh/lsh/internal/renderer"
)

// CreateVirtualNetworkReader is a Reader for the CreateVirtualNetwork structure.
type CreateVirtualNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVirtualNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateVirtualNetworkCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := apierrors.NewUnauthorized()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := apierrors.NewNotFound()
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
		return nil, runtime.NewAPIError("[POST /virtual_networks] create-virtual-network", response, response.Code())
	}
}

// NewCreateVirtualNetworkCreated creates a CreateVirtualNetworkCreated with default headers values
func NewCreateVirtualNetworkCreated() *CreateVirtualNetworkCreated {
	return &CreateVirtualNetworkCreated{}
}

/*
CreateVirtualNetworkCreated describes a response with status code 201, with default header values.

Created
*/
type CreateVirtualNetworkCreated struct {
	Payload *GetVirtualNetworkOKBody
}

// IsSuccess returns true when this create virtual network created response has a 2xx status code
func (o *CreateVirtualNetworkCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create virtual network created response has a 3xx status code
func (o *CreateVirtualNetworkCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create virtual network created response has a 4xx status code
func (o *CreateVirtualNetworkCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create virtual network created response has a 5xx status code
func (o *CreateVirtualNetworkCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create virtual network created response a status code equal to that given
func (o *CreateVirtualNetworkCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create virtual network created response
func (o *CreateVirtualNetworkCreated) Code() int {
	return 201
}

func (o *CreateVirtualNetworkCreated) Error() string {
	return fmt.Sprintf("[POST /virtual_networks][%d] createVirtualNetworkCreated  %+v", 201, o.Payload)
}

func (o *CreateVirtualNetworkCreated) String() string {
	return fmt.Sprintf("[POST /virtual_networks][%d] createVirtualNetworkCreated  %+v", 201, o.Payload)
}

func (o *CreateVirtualNetworkCreated) GetPayload() *GetVirtualNetworkOKBody {
	return o.Payload
}

func (o *CreateVirtualNetworkCreated) GetData() []renderer.ResponseData {
	return []renderer.ResponseData{o.Payload.Data}
}

func (o *CreateVirtualNetworkCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetVirtualNetworkOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateVirtualNetworkBody create virtual network body
swagger:model CreateVirtualNetworkBody
*/
type CreateVirtualNetworkBody struct {

	// data
	// Required: true
	Data *CreateVirtualNetworkParamsBodyData `json:"data"`
}

// Validate validates this create virtual network body
func (o *CreateVirtualNetworkBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateVirtualNetworkBody) validateData(formats strfmt.Registry) error {

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

// ContextValidate validate this create virtual network body based on the context it is used
func (o *CreateVirtualNetworkBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateVirtualNetworkBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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
func (o *CreateVirtualNetworkBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateVirtualNetworkBody) UnmarshalBinary(b []byte) error {
	var res CreateVirtualNetworkBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateVirtualNetworkParamsBodyData create virtual network params body data
swagger:model CreateVirtualNetworkParamsBodyData
*/
type CreateVirtualNetworkParamsBodyData struct {

	// attributes
	// Required: true
	Attributes *CreateVirtualNetworkParamsBodyDataAttributes `json:"attributes"`

	// type
	// Required: true
	// Enum: [virtual_network]
	Type *string `json:"type"`
}

// Validate validates this create virtual network params body data
func (o *CreateVirtualNetworkParamsBodyData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAttributes(formats); err != nil {
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

func (o *CreateVirtualNetworkParamsBodyData) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"data"+"."+"attributes", "body", o.Attributes); err != nil {
		return err
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

var createVirtualNetworkParamsBodyDataTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["virtual_network"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createVirtualNetworkParamsBodyDataTypeTypePropEnum = append(createVirtualNetworkParamsBodyDataTypeTypePropEnum, v)
	}
}

const (

	// CreateVirtualNetworkParamsBodyDataTypeVirtualNetwork captures enum value "virtual_network"
	CreateVirtualNetworkParamsBodyDataTypeVirtualNetwork string = "virtual_network"
)

// prop value enum
func (o *CreateVirtualNetworkParamsBodyData) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createVirtualNetworkParamsBodyDataTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateVirtualNetworkParamsBodyData) validateType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"data"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("body"+"."+"data"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create virtual network params body data based on the context it is used
func (o *CreateVirtualNetworkParamsBodyData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateVirtualNetworkParamsBodyData) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	if o.Attributes != nil {

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
func (o *CreateVirtualNetworkParamsBodyData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateVirtualNetworkParamsBodyData) UnmarshalBinary(b []byte) error {
	var res CreateVirtualNetworkParamsBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateVirtualNetworkParamsBodyDataAttributes create virtual network params body data attributes
swagger:model CreateVirtualNetworkParamsBodyDataAttributes
*/
type CreateVirtualNetworkParamsBodyDataAttributes struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// Project ID or slug
	// Required: true
	Project *string `json:"project"`

	// Site ID or slug
	// Enum: [ASH BGT BUE CHI DAL FRA LAX LON MEX MEX2 MIA MIA2 NYC SAN SAN2 SAO SAO2 SYD TYO TYO2]
	Site string `json:"site,omitempty"`
}

// Validate validates this create virtual network params body data attributes
func (o *CreateVirtualNetworkParamsBodyDataAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateVirtualNetworkParamsBodyDataAttributes) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"data"+"."+"attributes"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *CreateVirtualNetworkParamsBodyDataAttributes) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"data"+"."+"attributes"+"."+"project", "body", o.Project); err != nil {
		return err
	}

	return nil
}

var createVirtualNetworkParamsBodyDataAttributesTypeSitePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ASH","BGT","BUE","CHI","DAL","FRA","LAX","LON","MEX","MEX2","MIA","MIA2","NYC","SAN","SAN2","SAO","SAO2","SYD","TYO","TYO2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createVirtualNetworkParamsBodyDataAttributesTypeSitePropEnum = append(createVirtualNetworkParamsBodyDataAttributesTypeSitePropEnum, v)
	}
}

const (

	// CreateVirtualNetworkParamsBodyDataAttributesSiteASH captures enum value "ASH"
	CreateVirtualNetworkParamsBodyDataAttributesSiteASH string = "ASH"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteBGT captures enum value "BGT"
	CreateVirtualNetworkParamsBodyDataAttributesSiteBGT string = "BGT"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteBUE captures enum value "BUE"
	CreateVirtualNetworkParamsBodyDataAttributesSiteBUE string = "BUE"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteCHI captures enum value "CHI"
	CreateVirtualNetworkParamsBodyDataAttributesSiteCHI string = "CHI"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteDAL captures enum value "DAL"
	CreateVirtualNetworkParamsBodyDataAttributesSiteDAL string = "DAL"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteFRA captures enum value "FRA"
	CreateVirtualNetworkParamsBodyDataAttributesSiteFRA string = "FRA"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteLAX captures enum value "LAX"
	CreateVirtualNetworkParamsBodyDataAttributesSiteLAX string = "LAX"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteLON captures enum value "LON"
	CreateVirtualNetworkParamsBodyDataAttributesSiteLON string = "LON"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteMEX captures enum value "MEX"
	CreateVirtualNetworkParamsBodyDataAttributesSiteMEX string = "MEX"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteMEX2 captures enum value "MEX2"
	CreateVirtualNetworkParamsBodyDataAttributesSiteMEX2 string = "MEX2"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteMIA captures enum value "MIA"
	CreateVirtualNetworkParamsBodyDataAttributesSiteMIA string = "MIA"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteMIA2 captures enum value "MIA2"
	CreateVirtualNetworkParamsBodyDataAttributesSiteMIA2 string = "MIA2"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteNYC captures enum value "NYC"
	CreateVirtualNetworkParamsBodyDataAttributesSiteNYC string = "NYC"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteSAN captures enum value "SAN"
	CreateVirtualNetworkParamsBodyDataAttributesSiteSAN string = "SAN"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteSAN2 captures enum value "SAN2"
	CreateVirtualNetworkParamsBodyDataAttributesSiteSAN2 string = "SAN2"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteSAO captures enum value "SAO"
	CreateVirtualNetworkParamsBodyDataAttributesSiteSAO string = "SAO"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteSAO2 captures enum value "SAO2"
	CreateVirtualNetworkParamsBodyDataAttributesSiteSAO2 string = "SAO2"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteSYD captures enum value "SYD"
	CreateVirtualNetworkParamsBodyDataAttributesSiteSYD string = "SYD"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteTYO captures enum value "TYO"
	CreateVirtualNetworkParamsBodyDataAttributesSiteTYO string = "TYO"

	// CreateVirtualNetworkParamsBodyDataAttributesSiteTYO2 captures enum value "TYO2"
	CreateVirtualNetworkParamsBodyDataAttributesSiteTYO2 string = "TYO2"
)

// prop value enum
func (o *CreateVirtualNetworkParamsBodyDataAttributes) validateSiteEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createVirtualNetworkParamsBodyDataAttributesTypeSitePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateVirtualNetworkParamsBodyDataAttributes) validateSite(formats strfmt.Registry) error {
	if swag.IsZero(o.Site) { // not required
		return nil
	}

	// value enum
	if err := o.validateSiteEnum("body"+"."+"data"+"."+"attributes"+"."+"site", "body", o.Site); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create virtual network params body data attributes based on context it is used
func (o *CreateVirtualNetworkParamsBodyDataAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateVirtualNetworkParamsBodyDataAttributes) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateVirtualNetworkParamsBodyDataAttributes) UnmarshalBinary(b []byte) error {
	var res CreateVirtualNetworkParamsBodyDataAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
