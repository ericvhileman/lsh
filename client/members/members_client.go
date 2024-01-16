// Code generated by go-swagger; DO NOT EDIT.

package members

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new members API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for members API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DestroyTeamMember(params *DestroyTeamMemberParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DestroyTeamMemberOK, error)

	GetTeamMembers(params *GetTeamMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTeamMembersOK, error)

	PostTeamMembers(params *PostTeamMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostTeamMembersCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DestroyTeamMember removes a team member
*/
func (a *Client) DestroyTeamMember(params *DestroyTeamMemberParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DestroyTeamMemberOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDestroyTeamMemberParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "destroy-team-member",
		Method:             "DELETE",
		PathPattern:        "/team/members/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DestroyTeamMemberReader{formats: a.formats},
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
	success, ok := result.(*DestroyTeamMemberOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for destroy-team-member: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTeamMembers lists all team members
*/
func (a *Client) GetTeamMembers(params *GetTeamMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTeamMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTeamMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-team-members",
		Method:             "GET",
		PathPattern:        "/team/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTeamMembersReader{formats: a.formats},
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
	success, ok := result.(*GetTeamMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-team-members: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostTeamMembers adds a team member
*/
func (a *Client) PostTeamMembers(params *PostTeamMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostTeamMembersCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTeamMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post-team-members",
		Method:             "POST",
		PathPattern:        "/team/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostTeamMembersReader{formats: a.formats},
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
	success, ok := result.(*PostTeamMembersCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post-team-members: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
