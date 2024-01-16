// Code generated by go-swagger; DO NOT EDIT.

package ip_m_i_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCreateIpmiSessionParams creates a new CreateIpmiSessionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateIpmiSessionParams() *CreateIpmiSessionParams {
	return &CreateIpmiSessionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateIpmiSessionParamsWithTimeout creates a new CreateIpmiSessionParams object
// with the ability to set a timeout on a request.
func NewCreateIpmiSessionParamsWithTimeout(timeout time.Duration) *CreateIpmiSessionParams {
	return &CreateIpmiSessionParams{
		timeout: timeout,
	}
}

// NewCreateIpmiSessionParamsWithContext creates a new CreateIpmiSessionParams object
// with the ability to set a context for a request.
func NewCreateIpmiSessionParamsWithContext(ctx context.Context) *CreateIpmiSessionParams {
	return &CreateIpmiSessionParams{
		Context: ctx,
	}
}

// NewCreateIpmiSessionParamsWithHTTPClient creates a new CreateIpmiSessionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateIpmiSessionParamsWithHTTPClient(client *http.Client) *CreateIpmiSessionParams {
	return &CreateIpmiSessionParams{
		HTTPClient: client,
	}
}

/*
CreateIpmiSessionParams contains all the parameters to send to the API endpoint

	for the create ipmi session operation.

	Typically these are written to a http.Request.
*/
type CreateIpmiSessionParams struct {

	// APIVersion.
	//
	// Default: "2023-06-01"
	APIVersion *string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create ipmi session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIpmiSessionParams) WithDefaults() *CreateIpmiSessionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create ipmi session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIpmiSessionParams) SetDefaults() {
	var (
		aPIVersionDefault = string("2023-06-01")
	)

	val := CreateIpmiSessionParams{
		APIVersion: &aPIVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create ipmi session params
func (o *CreateIpmiSessionParams) WithTimeout(timeout time.Duration) *CreateIpmiSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create ipmi session params
func (o *CreateIpmiSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create ipmi session params
func (o *CreateIpmiSessionParams) WithContext(ctx context.Context) *CreateIpmiSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create ipmi session params
func (o *CreateIpmiSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create ipmi session params
func (o *CreateIpmiSessionParams) WithHTTPClient(client *http.Client) *CreateIpmiSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create ipmi session params
func (o *CreateIpmiSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create ipmi session params
func (o *CreateIpmiSessionParams) WithAPIVersion(aPIVersion *string) *CreateIpmiSessionParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create ipmi session params
func (o *CreateIpmiSessionParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the create ipmi session params
func (o *CreateIpmiSessionParams) WithID(id string) *CreateIpmiSessionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create ipmi session params
func (o *CreateIpmiSessionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateIpmiSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIVersion != nil {

		// header param API-Version
		if err := r.SetHeaderParam("API-Version", *o.APIVersion); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
