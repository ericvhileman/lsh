package virtual_network_assignments

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
	tablerows "github.com/latitudesh/lsh/internal/output/table/rows"
	"github.com/latitudesh/lsh/models"
)

// GetVirtualNetworksAssignmentsReader is a Reader for the GetVirtualNetworksAssignments structure.
type GetVirtualNetworksAssignmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVirtualNetworksAssignmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVirtualNetworksAssignmentsOK()
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
		return nil, runtime.NewAPIError("[GET /virtual_networks/assignments] get-virtual-networks-assignments", response, response.Code())
	}
}

// NewGetVirtualNetworksAssignmentsOK creates a GetVirtualNetworksAssignmentsOK with default headers values
func NewGetVirtualNetworksAssignmentsOK() *GetVirtualNetworksAssignmentsOK {
	return &GetVirtualNetworksAssignmentsOK{}
}

/*
GetVirtualNetworksAssignmentsOK describes a response with status code 200, with default header values.

Success
*/
type GetVirtualNetworksAssignmentsOK struct {
	Payload *models.VirtualNetworkAssignments
}

// IsSuccess returns true when this get virtual networks assignments o k response has a 2xx status code
func (o *GetVirtualNetworksAssignmentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get virtual networks assignments o k response has a 3xx status code
func (o *GetVirtualNetworksAssignmentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get virtual networks assignments o k response has a 4xx status code
func (o *GetVirtualNetworksAssignmentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get virtual networks assignments o k response has a 5xx status code
func (o *GetVirtualNetworksAssignmentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get virtual networks assignments o k response a status code equal to that given
func (o *GetVirtualNetworksAssignmentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get virtual networks assignments o k response
func (o *GetVirtualNetworksAssignmentsOK) Code() int {
	return 200
}

func (o *GetVirtualNetworksAssignmentsOK) Error() string {
	return fmt.Sprintf("[GET /virtual_networks/assignments][%d] getVirtualNetworksAssignmentsOK  %+v", 200, o.Payload)
}

func (o *GetVirtualNetworksAssignmentsOK) String() string {
	return fmt.Sprintf("[GET /virtual_networks/assignments][%d] getVirtualNetworksAssignmentsOK  %+v", 200, o.Payload)
}

func (o *GetVirtualNetworksAssignmentsOK) GetPayload() *models.VirtualNetworkAssignments {
	return o.Payload
}

func (o *GetVirtualNetworksAssignmentsOK) Render() {
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

func (o *GetVirtualNetworksAssignmentsOK) RenderJSON() {
	if !swag.IsZero(o) && !swag.IsZero(o.Payload) {
		JSONString, err := json.Marshal(o.Payload)
		if err != nil {
			fmt.Println("Could not decode the result as JSON.")
		}

		output.RenderJSON(JSONString)
	}
}

func (o *GetVirtualNetworksAssignmentsOK) RenderTable() {
	var rows []table.Row

	for _, assignment := range o.Payload.Data {
		rows = append(rows, tablerows.NewVirtualNetworkAssignmentRow(assignment))
	}

	table.Render(rows)
}

func (o *GetVirtualNetworksAssignmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualNetworkAssignments)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
