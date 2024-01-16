// Code generated by go-swagger; DO NOT EDIT.

package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/latitudesh/cli/models"
)

// GetServersReader is a Reader for the GetServers structure.
type GetServersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /servers] get-servers", response, response.Code())
	}
}

// NewGetServersOK creates a GetServersOK with default headers values
func NewGetServersOK() *GetServersOK {
	return &GetServersOK{}
}

/*
GetServersOK describes a response with status code 200, with default header values.

Success
*/
type GetServersOK struct {
	Payload *models.Servers
}

// IsSuccess returns true when this get servers o k response has a 2xx status code
func (o *GetServersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get servers o k response has a 3xx status code
func (o *GetServersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get servers o k response has a 4xx status code
func (o *GetServersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get servers o k response has a 5xx status code
func (o *GetServersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get servers o k response a status code equal to that given
func (o *GetServersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get servers o k response
func (o *GetServersOK) Code() int {
	return 200
}

func (o *GetServersOK) Error() string {
	return fmt.Sprintf("[GET /servers][%d] getServersOK  %+v", 200, o.Payload)
}

func (o *GetServersOK) String() string {
	return fmt.Sprintf("[GET /servers][%d] getServersOK  %+v", 200, o.Payload)
}

func (o *GetServersOK) GetPayload() *models.Servers {
	return o.Payload
}

func (o *GetServersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Servers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}