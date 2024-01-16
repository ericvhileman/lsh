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

// DestroyServerReader is a Reader for the DestroyServer structure.
type DestroyServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DestroyServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDestroyServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDestroyServerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewDestroyServerNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDestroyServerUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /servers/{server_id}] destroy-server", response, response.Code())
	}
}

// NewDestroyServerNoContent creates a DestroyServerNoContent with default headers values
func NewDestroyServerNoContent() *DestroyServerNoContent {
	return &DestroyServerNoContent{}
}

/*
DestroyServerNoContent describes a response with status code 204, with default header values.

No Content
*/
type DestroyServerNoContent struct {
}

// IsSuccess returns true when this destroy server no content response has a 2xx status code
func (o *DestroyServerNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this destroy server no content response has a 3xx status code
func (o *DestroyServerNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this destroy server no content response has a 4xx status code
func (o *DestroyServerNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this destroy server no content response has a 5xx status code
func (o *DestroyServerNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this destroy server no content response a status code equal to that given
func (o *DestroyServerNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the destroy server no content response
func (o *DestroyServerNoContent) Code() int {
	return 204
}

func (o *DestroyServerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /servers/{server_id}][%d] destroyServerNoContent ", 204)
}

func (o *DestroyServerNoContent) String() string {
	return fmt.Sprintf("[DELETE /servers/{server_id}][%d] destroyServerNoContent ", 204)
}

func (o *DestroyServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDestroyServerForbidden creates a DestroyServerForbidden with default headers values
func NewDestroyServerForbidden() *DestroyServerForbidden {
	return &DestroyServerForbidden{}
}

/*
DestroyServerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DestroyServerForbidden struct {
	Payload *models.ErrorObject
}

// IsSuccess returns true when this destroy server forbidden response has a 2xx status code
func (o *DestroyServerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this destroy server forbidden response has a 3xx status code
func (o *DestroyServerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this destroy server forbidden response has a 4xx status code
func (o *DestroyServerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this destroy server forbidden response has a 5xx status code
func (o *DestroyServerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this destroy server forbidden response a status code equal to that given
func (o *DestroyServerForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the destroy server forbidden response
func (o *DestroyServerForbidden) Code() int {
	return 403
}

func (o *DestroyServerForbidden) Error() string {
	return fmt.Sprintf("[DELETE /servers/{server_id}][%d] destroyServerForbidden  %+v", 403, o.Payload)
}

func (o *DestroyServerForbidden) String() string {
	return fmt.Sprintf("[DELETE /servers/{server_id}][%d] destroyServerForbidden  %+v", 403, o.Payload)
}

func (o *DestroyServerForbidden) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *DestroyServerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDestroyServerNotAcceptable creates a DestroyServerNotAcceptable with default headers values
func NewDestroyServerNotAcceptable() *DestroyServerNotAcceptable {
	return &DestroyServerNotAcceptable{}
}

/*
DestroyServerNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type DestroyServerNotAcceptable struct {
	Payload *models.ErrorObject
}

// IsSuccess returns true when this destroy server not acceptable response has a 2xx status code
func (o *DestroyServerNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this destroy server not acceptable response has a 3xx status code
func (o *DestroyServerNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this destroy server not acceptable response has a 4xx status code
func (o *DestroyServerNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this destroy server not acceptable response has a 5xx status code
func (o *DestroyServerNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this destroy server not acceptable response a status code equal to that given
func (o *DestroyServerNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the destroy server not acceptable response
func (o *DestroyServerNotAcceptable) Code() int {
	return 406
}

func (o *DestroyServerNotAcceptable) Error() string {
	return fmt.Sprintf("[DELETE /servers/{server_id}][%d] destroyServerNotAcceptable  %+v", 406, o.Payload)
}

func (o *DestroyServerNotAcceptable) String() string {
	return fmt.Sprintf("[DELETE /servers/{server_id}][%d] destroyServerNotAcceptable  %+v", 406, o.Payload)
}

func (o *DestroyServerNotAcceptable) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *DestroyServerNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDestroyServerUnprocessableEntity creates a DestroyServerUnprocessableEntity with default headers values
func NewDestroyServerUnprocessableEntity() *DestroyServerUnprocessableEntity {
	return &DestroyServerUnprocessableEntity{}
}

/*
DestroyServerUnprocessableEntity describes a response with status code 422, with default header values.

Unsupported operation for device
*/
type DestroyServerUnprocessableEntity struct {
	Payload *models.ErrorObject
}

// IsSuccess returns true when this destroy server unprocessable entity response has a 2xx status code
func (o *DestroyServerUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this destroy server unprocessable entity response has a 3xx status code
func (o *DestroyServerUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this destroy server unprocessable entity response has a 4xx status code
func (o *DestroyServerUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this destroy server unprocessable entity response has a 5xx status code
func (o *DestroyServerUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this destroy server unprocessable entity response a status code equal to that given
func (o *DestroyServerUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the destroy server unprocessable entity response
func (o *DestroyServerUnprocessableEntity) Code() int {
	return 422
}

func (o *DestroyServerUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /servers/{server_id}][%d] destroyServerUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DestroyServerUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /servers/{server_id}][%d] destroyServerUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DestroyServerUnprocessableEntity) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *DestroyServerUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}