// Code generated by go-swagger; DO NOT EDIT.

package operating_systems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operating systems API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operating systems API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetPlansOperatingSystem(params *GetPlansOperatingSystemParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPlansOperatingSystemOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetPlansOperatingSystem lists all operating systems available

Lists all operating systems available to deploy and reinstall.
*/
func (a *Client) GetPlansOperatingSystem(params *GetPlansOperatingSystemParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPlansOperatingSystemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlansOperatingSystemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-plans-operating-system",
		Method:             "GET",
		PathPattern:        "/plans/operating_systems",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPlansOperatingSystemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPlansOperatingSystemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-plans-operating-system: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}