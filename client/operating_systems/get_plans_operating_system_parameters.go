// Code generated by go-swagger; DO NOT EDIT.

package operating_systems

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

// NewGetPlansOperatingSystemParams creates a new GetPlansOperatingSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPlansOperatingSystemParams() *GetPlansOperatingSystemParams {
	return &GetPlansOperatingSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlansOperatingSystemParamsWithTimeout creates a new GetPlansOperatingSystemParams object
// with the ability to set a timeout on a request.
func NewGetPlansOperatingSystemParamsWithTimeout(timeout time.Duration) *GetPlansOperatingSystemParams {
	return &GetPlansOperatingSystemParams{
		timeout: timeout,
	}
}

// NewGetPlansOperatingSystemParamsWithContext creates a new GetPlansOperatingSystemParams object
// with the ability to set a context for a request.
func NewGetPlansOperatingSystemParamsWithContext(ctx context.Context) *GetPlansOperatingSystemParams {
	return &GetPlansOperatingSystemParams{
		Context: ctx,
	}
}

// NewGetPlansOperatingSystemParamsWithHTTPClient creates a new GetPlansOperatingSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPlansOperatingSystemParamsWithHTTPClient(client *http.Client) *GetPlansOperatingSystemParams {
	return &GetPlansOperatingSystemParams{
		HTTPClient: client,
	}
}

/*
GetPlansOperatingSystemParams contains all the parameters to send to the API endpoint

	for the get plans operating system operation.

	Typically these are written to a http.Request.
*/
type GetPlansOperatingSystemParams struct {

	// APIVersion.
	//
	// Default: "2023-06-01"
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get plans operating system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlansOperatingSystemParams) WithDefaults() *GetPlansOperatingSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get plans operating system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlansOperatingSystemParams) SetDefaults() {
	var (
		aPIVersionDefault = string("2023-06-01")
	)

	val := GetPlansOperatingSystemParams{
		APIVersion: &aPIVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get plans operating system params
func (o *GetPlansOperatingSystemParams) WithTimeout(timeout time.Duration) *GetPlansOperatingSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get plans operating system params
func (o *GetPlansOperatingSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get plans operating system params
func (o *GetPlansOperatingSystemParams) WithContext(ctx context.Context) *GetPlansOperatingSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get plans operating system params
func (o *GetPlansOperatingSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get plans operating system params
func (o *GetPlansOperatingSystemParams) WithHTTPClient(client *http.Client) *GetPlansOperatingSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get plans operating system params
func (o *GetPlansOperatingSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get plans operating system params
func (o *GetPlansOperatingSystemParams) WithAPIVersion(aPIVersion *string) *GetPlansOperatingSystemParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get plans operating system params
func (o *GetPlansOperatingSystemParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlansOperatingSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}