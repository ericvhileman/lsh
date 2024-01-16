// Code generated by go-swagger; DO NOT EDIT.

package api_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new api keys API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for api keys API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAPIKey(params *DeleteAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAPIKeyOK, error)

	GetAPIKeys(params *GetAPIKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPIKeysOK, error)

	PostAPIKey(params *PostAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAPIKeyCreated, error)

	UpdateAPIKey(params *UpdateAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAPIKeyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteAPIKey deletes an API key

Delete an existing API Key. Once deleted, the API Key can no longer be used to access the API.
*/
func (a *Client) DeleteAPIKey(params *DeleteAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAPIKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-api-key",
		Method:             "DELETE",
		PathPattern:        "/auth/api_keys/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIKeyReader{formats: a.formats},
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
	success, ok := result.(*DeleteAPIKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-api-key: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIKeys lists all API keys

Returns a list of all API keys from the team members
*/
func (a *Client) GetAPIKeys(params *GetAPIKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPIKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIKeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-api-keys",
		Method:             "GET",
		PathPattern:        "/auth/api_keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIKeysReader{formats: a.formats},
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
	success, ok := result.(*GetAPIKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-api-keys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAPIKey creates an API key

Create a new API Key that is tied to the current user account. The created API key is only listed ONCE upon creation. It can however be regenerated or deleted.
*/
func (a *Client) PostAPIKey(params *PostAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAPIKeyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post-api-key",
		Method:             "POST",
		PathPattern:        "/auth/api_keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIKeyReader{formats: a.formats},
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
	success, ok := result.(*PostAPIKeyCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post-api-key: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateAPIKey regenerates an API key

Regenerate an existing API Key that is tied to the current user. This overrides the previous key.
*/
func (a *Client) UpdateAPIKey(params *UpdateAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAPIKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAPIKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update-api-key",
		Method:             "PUT",
		PathPattern:        "/auth/api_keys/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAPIKeyReader{formats: a.formats},
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
	success, ok := result.(*UpdateAPIKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update-api-key: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}