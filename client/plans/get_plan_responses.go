// Code generated by go-swagger; DO NOT EDIT.

package plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/latitudesh/cli/internal/api"
	"github.com/latitudesh/cli/models"
)

// GetPlanReader is a Reader for the GetPlan structure.
type GetPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := api.NewNotFound()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /plans/{plan_id}] get-plan", response, response.Code())
	}
}

// NewGetPlanOK creates a GetPlanOK with default headers values
func NewGetPlanOK() *GetPlanOK {
	return &GetPlanOK{}
}

/*
GetPlanOK describes a response with status code 200, with default header values.

Success
*/
type GetPlanOK struct {
	Payload *models.Plan
}

// IsSuccess returns true when this get plan o k response has a 2xx status code
func (o *GetPlanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get plan o k response has a 3xx status code
func (o *GetPlanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get plan o k response has a 4xx status code
func (o *GetPlanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get plan o k response has a 5xx status code
func (o *GetPlanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get plan o k response a status code equal to that given
func (o *GetPlanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get plan o k response
func (o *GetPlanOK) Code() int {
	return 200
}

func (o *GetPlanOK) Error() string {
	return fmt.Sprintf("[GET /plans/{plan_id}][%d] getPlanOK  %+v", 200, o.Payload)
}

func (o *GetPlanOK) String() string {
	return fmt.Sprintf("[GET /plans/{plan_id}][%d] getPlanOK  %+v", 200, o.Payload)
}

func (o *GetPlanOK) GetPayload() *models.Plan {
	return o.Payload
}

func (o *GetPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Plan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlanNotFound creates a GetPlanNotFound with default headers values
func NewGetPlanNotFound() *GetPlanNotFound {
	return &GetPlanNotFound{}
}

/*
GetPlanNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetPlanNotFound struct {
	Payload *models.ErrorObject
}

// IsSuccess returns true when this get plan not found response has a 2xx status code
func (o *GetPlanNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get plan not found response has a 3xx status code
func (o *GetPlanNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get plan not found response has a 4xx status code
func (o *GetPlanNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get plan not found response has a 5xx status code
func (o *GetPlanNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get plan not found response a status code equal to that given
func (o *GetPlanNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get plan not found response
func (o *GetPlanNotFound) Code() int {
	return 404
}

func (o *GetPlanNotFound) Error() string {
	return fmt.Sprintf("[GET /plans/{plan_id}][%d] getPlanNotFound  %+v", 404, o.Payload)
}

func (o *GetPlanNotFound) String() string {
	return fmt.Sprintf("[GET /plans/{plan_id}][%d] getPlanNotFound  %+v", 404, o.Payload)
}

func (o *GetPlanNotFound) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *GetPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
