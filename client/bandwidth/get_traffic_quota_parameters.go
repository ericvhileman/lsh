// Code generated by go-swagger; DO NOT EDIT.

package bandwidth

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

// NewGetTrafficQuotaParams creates a new GetTrafficQuotaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTrafficQuotaParams() *GetTrafficQuotaParams {
	return &GetTrafficQuotaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTrafficQuotaParamsWithTimeout creates a new GetTrafficQuotaParams object
// with the ability to set a timeout on a request.
func NewGetTrafficQuotaParamsWithTimeout(timeout time.Duration) *GetTrafficQuotaParams {
	return &GetTrafficQuotaParams{
		timeout: timeout,
	}
}

// NewGetTrafficQuotaParamsWithContext creates a new GetTrafficQuotaParams object
// with the ability to set a context for a request.
func NewGetTrafficQuotaParamsWithContext(ctx context.Context) *GetTrafficQuotaParams {
	return &GetTrafficQuotaParams{
		Context: ctx,
	}
}

// NewGetTrafficQuotaParamsWithHTTPClient creates a new GetTrafficQuotaParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTrafficQuotaParamsWithHTTPClient(client *http.Client) *GetTrafficQuotaParams {
	return &GetTrafficQuotaParams{
		HTTPClient: client,
	}
}

/*
GetTrafficQuotaParams contains all the parameters to send to the API endpoint

	for the get traffic quota operation.

	Typically these are written to a http.Request.
*/
type GetTrafficQuotaParams struct {

	// APIVersion.
	//
	// Default: "2023-06-01"
	APIVersion *string

	// FilterProject.
	FilterProject *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get traffic quota params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTrafficQuotaParams) WithDefaults() *GetTrafficQuotaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get traffic quota params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTrafficQuotaParams) SetDefaults() {
	var (
		aPIVersionDefault = string("2023-06-01")
	)

	val := GetTrafficQuotaParams{
		APIVersion: &aPIVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get traffic quota params
func (o *GetTrafficQuotaParams) WithTimeout(timeout time.Duration) *GetTrafficQuotaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get traffic quota params
func (o *GetTrafficQuotaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get traffic quota params
func (o *GetTrafficQuotaParams) WithContext(ctx context.Context) *GetTrafficQuotaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get traffic quota params
func (o *GetTrafficQuotaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get traffic quota params
func (o *GetTrafficQuotaParams) WithHTTPClient(client *http.Client) *GetTrafficQuotaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get traffic quota params
func (o *GetTrafficQuotaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get traffic quota params
func (o *GetTrafficQuotaParams) WithAPIVersion(aPIVersion *string) *GetTrafficQuotaParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get traffic quota params
func (o *GetTrafficQuotaParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithFilterProject adds the filterProject to the get traffic quota params
func (o *GetTrafficQuotaParams) WithFilterProject(filterProject *string) *GetTrafficQuotaParams {
	o.SetFilterProject(filterProject)
	return o
}

// SetFilterProject adds the filterProject to the get traffic quota params
func (o *GetTrafficQuotaParams) SetFilterProject(filterProject *string) {
	o.FilterProject = filterProject
}

// WriteToRequest writes these params to a swagger request
func (o *GetTrafficQuotaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FilterProject != nil {

		// query param filter[project]
		var qrFilterProject string

		if o.FilterProject != nil {
			qrFilterProject = *o.FilterProject
		}
		qFilterProject := qrFilterProject
		if qFilterProject != "" {

			if err := r.SetQueryParam("filter[project]", qFilterProject); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
