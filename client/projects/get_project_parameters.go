// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewGetProjectParams creates a new GetProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectParams() *GetProjectParams {
	return &GetProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectParamsWithTimeout creates a new GetProjectParams object
// with the ability to set a timeout on a request.
func NewGetProjectParamsWithTimeout(timeout time.Duration) *GetProjectParams {
	return &GetProjectParams{
		timeout: timeout,
	}
}

// NewGetProjectParamsWithContext creates a new GetProjectParams object
// with the ability to set a context for a request.
func NewGetProjectParamsWithContext(ctx context.Context) *GetProjectParams {
	return &GetProjectParams{
		Context: ctx,
	}
}

// NewGetProjectParamsWithHTTPClient creates a new GetProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectParamsWithHTTPClient(client *http.Client) *GetProjectParams {
	return &GetProjectParams{
		HTTPClient: client,
	}
}

/*
GetProjectParams contains all the parameters to send to the API endpoint

	for the get project operation.

	Typically these are written to a http.Request.
*/
type GetProjectParams struct {

	// APIVersion.
	//
	// Default: "2023-06-01"
	APIVersion *string

	/* ExtraFieldsProjects.

	   The `last_renewal_date` and `next_renewal_date` are provided as extra attributes that show previous and future billing cycle dates. To request it, just set `extra_fields[projects]=last_renewal_date,next_renewal_date` in the query string.
	*/
	ExtraFieldsProjects *string

	/* IDOrSlug.

	   The project ID or Slug
	*/
	IDOrSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectParams) WithDefaults() *GetProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectParams) SetDefaults() {
	var (
		aPIVersionDefault = string("2023-06-01")
	)

	val := GetProjectParams{
		APIVersion: &aPIVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get project params
func (o *GetProjectParams) WithTimeout(timeout time.Duration) *GetProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project params
func (o *GetProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project params
func (o *GetProjectParams) WithContext(ctx context.Context) *GetProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project params
func (o *GetProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project params
func (o *GetProjectParams) WithHTTPClient(client *http.Client) *GetProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project params
func (o *GetProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get project params
func (o *GetProjectParams) WithAPIVersion(aPIVersion *string) *GetProjectParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get project params
func (o *GetProjectParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithExtraFieldsProjects adds the extraFieldsProjects to the get project params
func (o *GetProjectParams) WithExtraFieldsProjects(extraFieldsProjects *string) *GetProjectParams {
	o.SetExtraFieldsProjects(extraFieldsProjects)
	return o
}

// SetExtraFieldsProjects adds the extraFieldsProjects to the get project params
func (o *GetProjectParams) SetExtraFieldsProjects(extraFieldsProjects *string) {
	o.ExtraFieldsProjects = extraFieldsProjects
}

// WithIDOrSlug adds the iDOrSlug to the get project params
func (o *GetProjectParams) WithIDOrSlug(iDOrSlug string) *GetProjectParams {
	o.SetIDOrSlug(iDOrSlug)
	return o
}

// SetIDOrSlug adds the idOrSlug to the get project params
func (o *GetProjectParams) SetIDOrSlug(iDOrSlug string) {
	o.IDOrSlug = iDOrSlug
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ExtraFieldsProjects != nil {

		// query param extra_fields[projects]
		var qrExtraFieldsProjects string

		if o.ExtraFieldsProjects != nil {
			qrExtraFieldsProjects = *o.ExtraFieldsProjects
		}
		qExtraFieldsProjects := qrExtraFieldsProjects
		if qExtraFieldsProjects != "" {

			if err := r.SetQueryParam("extra_fields[projects]", qExtraFieldsProjects); err != nil {
				return err
			}
		}
	}

	// path param id_or_slug
	if err := r.SetPathParam("id_or_slug", o.IDOrSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
