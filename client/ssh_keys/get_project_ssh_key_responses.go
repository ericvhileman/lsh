package ssh_keys

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
	tablerows "github.com/latitudesh/lsh/internal/output/table/rows"
	"github.com/latitudesh/lsh/internal/utils"
	"github.com/latitudesh/lsh/models"
)

// GetProjectSSHKeyReader is a Reader for the GetProjectSSHKey structure.
type GetProjectSSHKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectSSHKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectSSHKeyOK()
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
	default:
		return nil, runtime.NewAPIError("[GET /projects/{project_id_or_slug}/ssh_keys/{ssh_key_id}] get-project-ssh-key", response, response.Code())
	}
}

// NewGetProjectSSHKeyOK creates a GetProjectSSHKeyOK with default headers values
func NewGetProjectSSHKeyOK() *GetProjectSSHKeyOK {
	return &GetProjectSSHKeyOK{}
}

/*
GetProjectSSHKeyOK describes a response with status code 200, with default header values.

Success
*/
type GetProjectSSHKeyOK struct {
	Payload *GetProjectSSHKeyOKBody
}

// IsSuccess returns true when this get project Ssh key o k response has a 2xx status code
func (o *GetProjectSSHKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project Ssh key o k response has a 3xx status code
func (o *GetProjectSSHKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project Ssh key o k response has a 4xx status code
func (o *GetProjectSSHKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project Ssh key o k response has a 5xx status code
func (o *GetProjectSSHKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project Ssh key o k response a status code equal to that given
func (o *GetProjectSSHKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get project Ssh key o k response
func (o *GetProjectSSHKeyOK) Code() int {
	return 200
}

func (o *GetProjectSSHKeyOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id_or_slug}/ssh_keys/{ssh_key_id}][%d] getProjectSshKeyOK  %+v", 200, o.Payload)
}

func (o *GetProjectSSHKeyOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_id_or_slug}/ssh_keys/{ssh_key_id}][%d] getProjectSshKeyOK  %+v", 200, o.Payload)
}

func (o *GetProjectSSHKeyOK) GetPayload() *GetProjectSSHKeyOKBody {
	return o.Payload
}

func (o *GetProjectSSHKeyOK) Render() {
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

func (o *GetProjectSSHKeyOK) RenderJSON() {
	if !swag.IsZero(o) && !swag.IsZero(o.Payload) {
		JSONString, err := json.Marshal(o.Payload)
		if err != nil {
			fmt.Println("Could not decode the result as JSON.")
		}

		output.RenderJSON(JSONString)
	}
}

func (o *GetProjectSSHKeyOK) RenderTable() {
	rows := []table.Row{tablerows.NewSSHKeyRow(o.Payload.Data)}
	utils.RenderTableU(rows)
}

func (o *GetProjectSSHKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetProjectSSHKeyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetProjectSSHKeyOKBody get project SSH key o k body
swagger:model GetProjectSSHKeyOKBody
*/
type GetProjectSSHKeyOKBody struct {

	// data
	Data *models.SSHKeyData `json:"data,omitempty"`
}

// Validate validates this get project SSH key o k body
func (o *GetProjectSSHKeyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetProjectSSHKeyOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getProjectSshKeyOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getProjectSshKeyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get project SSH key o k body based on the context it is used
func (o *GetProjectSSHKeyOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetProjectSSHKeyOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getProjectSshKeyOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getProjectSshKeyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetProjectSSHKeyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetProjectSSHKeyOKBody) UnmarshalBinary(b []byte) error {
	var res GetProjectSSHKeyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
