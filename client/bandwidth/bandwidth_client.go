// Code generated by go-swagger; DO NOT EDIT.

package bandwidth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new bandwidth API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bandwidth API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetTrafficConsumption(params *GetTrafficConsumptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTrafficConsumptionOK, error)

	GetTrafficQuota(params *GetTrafficQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTrafficQuotaOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetTrafficConsumption retrieves traffic consumption
*/
func (a *Client) GetTrafficConsumption(params *GetTrafficConsumptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTrafficConsumptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTrafficConsumptionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-traffic-consumption",
		Method:             "GET",
		PathPattern:        "/traffic",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTrafficConsumptionReader{formats: a.formats},
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
	success, ok := result.(*GetTrafficConsumptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-traffic-consumption: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTrafficQuota retrieves traffic quota
*/
func (a *Client) GetTrafficQuota(params *GetTrafficQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTrafficQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTrafficQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-traffic-quota",
		Method:             "GET",
		PathPattern:        "/traffic/quota",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTrafficQuotaReader{formats: a.formats},
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
	success, ok := result.(*GetTrafficQuotaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-traffic-quota: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}