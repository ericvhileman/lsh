// Code generated by go-swagger; DO NOT EDIT.

package plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/latitudesh/cli/models"
)

// GetBandwidthPlansReader is a Reader for the GetBandwidthPlans structure.
type GetBandwidthPlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBandwidthPlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBandwidthPlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /plans/bandwidth] get-bandwidth-plans", response, response.Code())
	}
}

// NewGetBandwidthPlansOK creates a GetBandwidthPlansOK with default headers values
func NewGetBandwidthPlansOK() *GetBandwidthPlansOK {
	return &GetBandwidthPlansOK{}
}

/*
GetBandwidthPlansOK describes a response with status code 200, with default header values.

Success
*/
type GetBandwidthPlansOK struct {
	Payload *models.BandwidthPlans
}

// IsSuccess returns true when this get bandwidth plans o k response has a 2xx status code
func (o *GetBandwidthPlansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get bandwidth plans o k response has a 3xx status code
func (o *GetBandwidthPlansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bandwidth plans o k response has a 4xx status code
func (o *GetBandwidthPlansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bandwidth plans o k response has a 5xx status code
func (o *GetBandwidthPlansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get bandwidth plans o k response a status code equal to that given
func (o *GetBandwidthPlansOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get bandwidth plans o k response
func (o *GetBandwidthPlansOK) Code() int {
	return 200
}

func (o *GetBandwidthPlansOK) Error() string {
	return fmt.Sprintf("[GET /plans/bandwidth][%d] getBandwidthPlansOK  %+v", 200, o.Payload)
}

func (o *GetBandwidthPlansOK) String() string {
	return fmt.Sprintf("[GET /plans/bandwidth][%d] getBandwidthPlansOK  %+v", 200, o.Payload)
}

func (o *GetBandwidthPlansOK) GetPayload() *models.BandwidthPlans {
	return o.Payload
}

func (o *GetBandwidthPlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BandwidthPlans)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
