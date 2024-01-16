// Code generated by go-swagger; DO NOT EDIT.

package ip_addresses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/latitudesh/cli/models"
)

// GetIPReader is a Reader for the GetIP structure.
type GetIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetIPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /ips/{id}] get-ip", response, response.Code())
	}
}

// NewGetIPOK creates a GetIPOK with default headers values
func NewGetIPOK() *GetIPOK {
	return &GetIPOK{}
}

/*
GetIPOK describes a response with status code 200, with default header values.

Success
*/
type GetIPOK struct {
	Payload *models.IPAddress
}

// IsSuccess returns true when this get Ip o k response has a 2xx status code
func (o *GetIPOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Ip o k response has a 3xx status code
func (o *GetIPOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Ip o k response has a 4xx status code
func (o *GetIPOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Ip o k response has a 5xx status code
func (o *GetIPOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Ip o k response a status code equal to that given
func (o *GetIPOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Ip o k response
func (o *GetIPOK) Code() int {
	return 200
}

func (o *GetIPOK) Error() string {
	return fmt.Sprintf("[GET /ips/{id}][%d] getIpOK  %+v", 200, o.Payload)
}

func (o *GetIPOK) String() string {
	return fmt.Sprintf("[GET /ips/{id}][%d] getIpOK  %+v", 200, o.Payload)
}

func (o *GetIPOK) GetPayload() *models.IPAddress {
	return o.Payload
}

func (o *GetIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPNotFound creates a GetIPNotFound with default headers values
func NewGetIPNotFound() *GetIPNotFound {
	return &GetIPNotFound{}
}

/*
GetIPNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetIPNotFound struct {
	Payload *models.ErrorObject
}

// IsSuccess returns true when this get Ip not found response has a 2xx status code
func (o *GetIPNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Ip not found response has a 3xx status code
func (o *GetIPNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Ip not found response has a 4xx status code
func (o *GetIPNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Ip not found response has a 5xx status code
func (o *GetIPNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Ip not found response a status code equal to that given
func (o *GetIPNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get Ip not found response
func (o *GetIPNotFound) Code() int {
	return 404
}

func (o *GetIPNotFound) Error() string {
	return fmt.Sprintf("[GET /ips/{id}][%d] getIpNotFound  %+v", 404, o.Payload)
}

func (o *GetIPNotFound) String() string {
	return fmt.Sprintf("[GET /ips/{id}][%d] getIpNotFound  %+v", 404, o.Payload)
}

func (o *GetIPNotFound) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *GetIPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
