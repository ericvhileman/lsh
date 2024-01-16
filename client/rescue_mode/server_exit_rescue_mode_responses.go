// Code generated by go-swagger; DO NOT EDIT.

package rescue_mode

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/latitudesh/cli/models"
)

// ServerExitRescueModeReader is a Reader for the ServerExitRescueMode structure.
type ServerExitRescueModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServerExitRescueModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServerExitRescueModeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewServerExitRescueModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewServerExitRescueModeNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /servers/{server_id}/exit_rescue_mode] server-exit-rescue-mode", response, response.Code())
	}
}

// NewServerExitRescueModeOK creates a ServerExitRescueModeOK with default headers values
func NewServerExitRescueModeOK() *ServerExitRescueModeOK {
	return &ServerExitRescueModeOK{}
}

/*
ServerExitRescueModeOK describes a response with status code 200, with default header values.

Success
*/
type ServerExitRescueModeOK struct {
	Payload *models.ServerRescue
}

// IsSuccess returns true when this server exit rescue mode o k response has a 2xx status code
func (o *ServerExitRescueModeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this server exit rescue mode o k response has a 3xx status code
func (o *ServerExitRescueModeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this server exit rescue mode o k response has a 4xx status code
func (o *ServerExitRescueModeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this server exit rescue mode o k response has a 5xx status code
func (o *ServerExitRescueModeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this server exit rescue mode o k response a status code equal to that given
func (o *ServerExitRescueModeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the server exit rescue mode o k response
func (o *ServerExitRescueModeOK) Code() int {
	return 200
}

func (o *ServerExitRescueModeOK) Error() string {
	return fmt.Sprintf("[POST /servers/{server_id}/exit_rescue_mode][%d] serverExitRescueModeOK  %+v", 200, o.Payload)
}

func (o *ServerExitRescueModeOK) String() string {
	return fmt.Sprintf("[POST /servers/{server_id}/exit_rescue_mode][%d] serverExitRescueModeOK  %+v", 200, o.Payload)
}

func (o *ServerExitRescueModeOK) GetPayload() *models.ServerRescue {
	return o.Payload
}

func (o *ServerExitRescueModeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerRescue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServerExitRescueModeForbidden creates a ServerExitRescueModeForbidden with default headers values
func NewServerExitRescueModeForbidden() *ServerExitRescueModeForbidden {
	return &ServerExitRescueModeForbidden{}
}

/*
ServerExitRescueModeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServerExitRescueModeForbidden struct {
	Payload *models.ErrorObject
}

// IsSuccess returns true when this server exit rescue mode forbidden response has a 2xx status code
func (o *ServerExitRescueModeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this server exit rescue mode forbidden response has a 3xx status code
func (o *ServerExitRescueModeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this server exit rescue mode forbidden response has a 4xx status code
func (o *ServerExitRescueModeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this server exit rescue mode forbidden response has a 5xx status code
func (o *ServerExitRescueModeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this server exit rescue mode forbidden response a status code equal to that given
func (o *ServerExitRescueModeForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the server exit rescue mode forbidden response
func (o *ServerExitRescueModeForbidden) Code() int {
	return 403
}

func (o *ServerExitRescueModeForbidden) Error() string {
	return fmt.Sprintf("[POST /servers/{server_id}/exit_rescue_mode][%d] serverExitRescueModeForbidden  %+v", 403, o.Payload)
}

func (o *ServerExitRescueModeForbidden) String() string {
	return fmt.Sprintf("[POST /servers/{server_id}/exit_rescue_mode][%d] serverExitRescueModeForbidden  %+v", 403, o.Payload)
}

func (o *ServerExitRescueModeForbidden) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *ServerExitRescueModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServerExitRescueModeNotAcceptable creates a ServerExitRescueModeNotAcceptable with default headers values
func NewServerExitRescueModeNotAcceptable() *ServerExitRescueModeNotAcceptable {
	return &ServerExitRescueModeNotAcceptable{}
}

/*
ServerExitRescueModeNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type ServerExitRescueModeNotAcceptable struct {
	Payload *models.ErrorObject
}

// IsSuccess returns true when this server exit rescue mode not acceptable response has a 2xx status code
func (o *ServerExitRescueModeNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this server exit rescue mode not acceptable response has a 3xx status code
func (o *ServerExitRescueModeNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this server exit rescue mode not acceptable response has a 4xx status code
func (o *ServerExitRescueModeNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this server exit rescue mode not acceptable response has a 5xx status code
func (o *ServerExitRescueModeNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this server exit rescue mode not acceptable response a status code equal to that given
func (o *ServerExitRescueModeNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the server exit rescue mode not acceptable response
func (o *ServerExitRescueModeNotAcceptable) Code() int {
	return 406
}

func (o *ServerExitRescueModeNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /servers/{server_id}/exit_rescue_mode][%d] serverExitRescueModeNotAcceptable  %+v", 406, o.Payload)
}

func (o *ServerExitRescueModeNotAcceptable) String() string {
	return fmt.Sprintf("[POST /servers/{server_id}/exit_rescue_mode][%d] serverExitRescueModeNotAcceptable  %+v", 406, o.Payload)
}

func (o *ServerExitRescueModeNotAcceptable) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *ServerExitRescueModeNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}