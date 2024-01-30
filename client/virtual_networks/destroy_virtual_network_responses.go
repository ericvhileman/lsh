// Code generated by go-swagger; DO NOT EDIT.

package virtual_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	apierrors "github.com/latitudesh/lsh/internal/api/errors"
	"github.com/latitudesh/lsh/internal/output"
)

// DestroyVirtualNetworkReader is a Reader for the DestroyVirtualNetwork structure.
type DestroyVirtualNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DestroyVirtualNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDestroyVirtualNetworkNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := apierrors.NewNotFound()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := apierrors.NewNotAcceptable()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /virtual_networks/{id}] destroy-virtual-network", response, response.Code())
	}
}

// NewDestroyVirtualNetworkNoContent creates a DestroyVirtualNetworkNoContent with default headers values
func NewDestroyVirtualNetworkNoContent() *DestroyVirtualNetworkNoContent {
	return &DestroyVirtualNetworkNoContent{}
}

/*
DestroyVirtualNetworkNoContent describes a response with status code 204, with default header values.

Destroyed
*/
type DestroyVirtualNetworkNoContent struct {
}

// IsSuccess returns true when this destroy virtual network no content response has a 2xx status code
func (o *DestroyVirtualNetworkNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this destroy virtual network no content response has a 3xx status code
func (o *DestroyVirtualNetworkNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this destroy virtual network no content response has a 4xx status code
func (o *DestroyVirtualNetworkNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this destroy virtual network no content response has a 5xx status code
func (o *DestroyVirtualNetworkNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this destroy virtual network no content response a status code equal to that given
func (o *DestroyVirtualNetworkNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the destroy virtual network no content response
func (o *DestroyVirtualNetworkNoContent) Code() int {
	return 204
}

func (o *DestroyVirtualNetworkNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtual_networks/{id}][%d] destroyVirtualNetworkNoContent ", 204)
}

func (o *DestroyVirtualNetworkNoContent) String() string {
	return fmt.Sprintf("[DELETE /virtual_networks/{id}][%d] destroyVirtualNetworkNoContent ", 204)
}

func (o *DestroyVirtualNetworkNoContent) RenderOutput() {
	output.SuccessfulDeletion("Virtual Network")
}

func (o *DestroyVirtualNetworkNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
