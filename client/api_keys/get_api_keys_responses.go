package api_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/spf13/viper"

	apierrors "github.com/latitudesh/lsh/internal/api/errors"
	"github.com/latitudesh/lsh/internal/output"
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

// GetAPIKeysReader is a Reader for the GetAPIKeys structure.
type GetAPIKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIKeysOK()
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
	default:
		return nil, runtime.NewAPIError("[GET /auth/api_keys] get-api-keys", response, response.Code())
	}
}

// NewGetAPIKeysOK creates a GetAPIKeysOK with default headers values
func NewGetAPIKeysOK() *GetAPIKeysOK {
	return &GetAPIKeysOK{}
}

/*
GetAPIKeysOK describes a response with status code 200, with default header values.

Success
*/
type GetAPIKeysOK struct {
	Payload *models.APIKeys
}

// IsSuccess returns true when this get Api keys o k response has a 2xx status code
func (o *GetAPIKeysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api keys o k response has a 3xx status code
func (o *GetAPIKeysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api keys o k response has a 4xx status code
func (o *GetAPIKeysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api keys o k response has a 5xx status code
func (o *GetAPIKeysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api keys o k response a status code equal to that given
func (o *GetAPIKeysOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api keys o k response
func (o *GetAPIKeysOK) Code() int {
	return 200
}

func (o *GetAPIKeysOK) Error() string {
	return fmt.Sprintf("[GET /auth/api_keys][%d] getApiKeysOK  %+v", 200, o.Payload)
}

func (o *GetAPIKeysOK) String() string {
	return fmt.Sprintf("[GET /auth/api_keys][%d] getApiKeysOK  %+v", 200, o.Payload)
}

func (o *GetAPIKeysOK) GetPayload() *models.APIKeys {
	return o.Payload
}

type APIKeyTableRow struct {
	ID             string `json:"id,omitempty"`
	TokenLastSlice string `json:"token_last_slice,omitempty"`
	User           string `json:"user,omitempty"`
	APIVersion     string `json:"api_version,omitempty"`
	LastUsedAt     string `json:"last_used_at,omitempty"`
	Name           string `json:"name,omitempty"`
}

func (o *GetAPIKeysOK) Render() {
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

func (o *GetAPIKeysOK) RenderJSON() {
	if !swag.IsZero(o) && !swag.IsZero(o.Payload) {
		JSONString, err := json.Marshal(o.Payload)
		if err != nil {
			fmt.Println("Could not decode the result as JSON.")
		}

		output.RenderJSON(JSONString)
	}
}

func (o *GetAPIKeysOK) RenderTable() {
	data := o.Payload.Data

	var rows []APIKeyTableRow

	for _, api_key := range data {
		attributes := api_key.Attributes

		row := APIKeyTableRow{
			ID:             table.String(api_key.ID),
			TokenLastSlice: table.String(attributes.TokenLastSlice),
			User:           table.String(attributes.User.Email),
			APIVersion:     table.String(attributes.APIVersion),
			LastUsedAt:     table.DateTime(attributes.LastUsedAt),
			Name:           table.String(attributes.Name),
		}

		rows = append(rows, row)
	}

	var interfaceRows []interface{}

	for _, row := range rows {
		interfaceRows = append(interfaceRows, row)
	}

	table.Render(interfaceRows)
}

func (o *GetAPIKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.APIKeys)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
