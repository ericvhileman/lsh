// Code generated by go-swagger; DO NOT EDIT.

package plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new plans API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for plans API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetBandwidthPlans(params *GetBandwidthPlansParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBandwidthPlansOK, error)

	GetPlan(params *GetPlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPlanOK, error)

	GetPlans(params *GetPlansParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPlansOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetBandwidthPlans lists all bandwidth plans

Lists all bandwidth plans.
*/
func (a *Client) GetBandwidthPlans(params *GetBandwidthPlansParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBandwidthPlansOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBandwidthPlansParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-bandwidth-plans",
		Method:             "GET",
		PathPattern:        "/plans/bandwidth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBandwidthPlansReader{formats: a.formats},
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
	success, ok := result.(*GetBandwidthPlansOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-bandwidth-plans: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPlan retrieves a plan
*/
func (a *Client) GetPlan(params *GetPlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-plan",
		Method:             "GET",
		PathPattern:        "/plans/{plan_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPlanReader{formats: a.formats},
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
	success, ok := result.(*GetPlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-plan: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPlans lists all plans

Lists all plans. Availability by region is included in `attributes.regions.locations.available[*]` node for a given plan.
*/
func (a *Client) GetPlans(params *GetPlansParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPlansOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlansParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-plans",
		Method:             "GET",
		PathPattern:        "/plans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPlansReader{formats: a.formats},
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
	success, ok := result.(*GetPlansOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-plans: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}