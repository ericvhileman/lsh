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

// NewGetProjectsParams creates a new GetProjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectsParams() *GetProjectsParams {
	return &GetProjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectsParamsWithTimeout creates a new GetProjectsParams object
// with the ability to set a timeout on a request.
func NewGetProjectsParamsWithTimeout(timeout time.Duration) *GetProjectsParams {
	return &GetProjectsParams{
		timeout: timeout,
	}
}

// NewGetProjectsParamsWithContext creates a new GetProjectsParams object
// with the ability to set a context for a request.
func NewGetProjectsParamsWithContext(ctx context.Context) *GetProjectsParams {
	return &GetProjectsParams{
		Context: ctx,
	}
}

// NewGetProjectsParamsWithHTTPClient creates a new GetProjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectsParamsWithHTTPClient(client *http.Client) *GetProjectsParams {
	return &GetProjectsParams{
		HTTPClient: client,
	}
}

/*
GetProjectsParams contains all the parameters to send to the API endpoint

	for the get projects operation.

	Typically these are written to a http.Request.
*/
type GetProjectsParams struct {

	/* ExtraFieldsProjects.

	   The `last_renewal_date` and `next_renewal_date` are provided as extra attributes that show previous and future billing cycle dates. To request it, just set `extra_fields[projects]=last_renewal_date,next_renewal_date` in the query string.
	*/
	ExtraFieldsProjects *string

	/* FilterBillingType.

	   The billing type to filter by
	*/
	FilterBillingType *string

	/* FilterDescription.

	   The project description to filter by
	*/
	FilterDescription *string

	/* FilterEnvironment.

	   The environment to filter by
	*/
	FilterEnvironment *string

	/* FilterName.

	   The project name to filter by
	*/
	FilterName *string

	/* FilterSlug.

	   The project slug to filter by
	*/
	FilterSlug *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectsParams) WithDefaults() *GetProjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectsParams) SetDefaults() {
	val := GetProjectsParams{}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get projects params
func (o *GetProjectsParams) WithTimeout(timeout time.Duration) *GetProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get projects params
func (o *GetProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get projects params
func (o *GetProjectsParams) WithContext(ctx context.Context) *GetProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get projects params
func (o *GetProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get projects params
func (o *GetProjectsParams) WithHTTPClient(client *http.Client) *GetProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get projects params
func (o *GetProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtraFieldsProjects adds the extraFieldsProjects to the get projects params
func (o *GetProjectsParams) WithExtraFieldsProjects(extraFieldsProjects *string) *GetProjectsParams {
	o.SetExtraFieldsProjects(extraFieldsProjects)
	return o
}

// SetExtraFieldsProjects adds the extraFieldsProjects to the get projects params
func (o *GetProjectsParams) SetExtraFieldsProjects(extraFieldsProjects *string) {
	o.ExtraFieldsProjects = extraFieldsProjects
}

// WithFilterBillingType adds the filterBillingType to the get projects params
func (o *GetProjectsParams) WithFilterBillingType(filterBillingType *string) *GetProjectsParams {
	o.SetFilterBillingType(filterBillingType)
	return o
}

// SetFilterBillingType adds the filterBillingType to the get projects params
func (o *GetProjectsParams) SetFilterBillingType(filterBillingType *string) {
	o.FilterBillingType = filterBillingType
}

// WithFilterDescription adds the filterDescription to the get projects params
func (o *GetProjectsParams) WithFilterDescription(filterDescription *string) *GetProjectsParams {
	o.SetFilterDescription(filterDescription)
	return o
}

// SetFilterDescription adds the filterDescription to the get projects params
func (o *GetProjectsParams) SetFilterDescription(filterDescription *string) {
	o.FilterDescription = filterDescription
}

// WithFilterEnvironment adds the filterEnvironment to the get projects params
func (o *GetProjectsParams) WithFilterEnvironment(filterEnvironment *string) *GetProjectsParams {
	o.SetFilterEnvironment(filterEnvironment)
	return o
}

// SetFilterEnvironment adds the filterEnvironment to the get projects params
func (o *GetProjectsParams) SetFilterEnvironment(filterEnvironment *string) {
	o.FilterEnvironment = filterEnvironment
}

// WithFilterName adds the filterName to the get projects params
func (o *GetProjectsParams) WithFilterName(filterName *string) *GetProjectsParams {
	o.SetFilterName(filterName)
	return o
}

// SetFilterName adds the filterName to the get projects params
func (o *GetProjectsParams) SetFilterName(filterName *string) {
	o.FilterName = filterName
}

// WithFilterSlug adds the filterSlug to the get projects params
func (o *GetProjectsParams) WithFilterSlug(filterSlug *string) *GetProjectsParams {
	o.SetFilterSlug(filterSlug)
	return o
}

// SetFilterSlug adds the filterSlug to the get projects params
func (o *GetProjectsParams) SetFilterSlug(filterSlug *string) {
	o.FilterSlug = filterSlug
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetQueryParam("page[size]", "100"); err != nil {
		return err
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

	if o.FilterBillingType != nil {

		// query param filter[billing_type]
		var qrFilterBillingType string

		if o.FilterBillingType != nil {
			qrFilterBillingType = *o.FilterBillingType
		}
		qFilterBillingType := qrFilterBillingType
		if qFilterBillingType != "" {

			if err := r.SetQueryParam("filter[billing_type]", qFilterBillingType); err != nil {
				return err
			}
		}
	}

	if o.FilterDescription != nil {

		// query param filter[description]
		var qrFilterDescription string

		if o.FilterDescription != nil {
			qrFilterDescription = *o.FilterDescription
		}
		qFilterDescription := qrFilterDescription
		if qFilterDescription != "" {

			if err := r.SetQueryParam("filter[description]", qFilterDescription); err != nil {
				return err
			}
		}
	}

	if o.FilterEnvironment != nil {

		// query param filter[environment]
		var qrFilterEnvironment string

		if o.FilterEnvironment != nil {
			qrFilterEnvironment = *o.FilterEnvironment
		}
		qFilterEnvironment := qrFilterEnvironment
		if qFilterEnvironment != "" {

			if err := r.SetQueryParam("filter[environment]", qFilterEnvironment); err != nil {
				return err
			}
		}
	}

	if o.FilterName != nil {

		// query param filter[name]
		var qrFilterName string

		if o.FilterName != nil {
			qrFilterName = *o.FilterName
		}
		qFilterName := qrFilterName
		if qFilterName != "" {

			if err := r.SetQueryParam("filter[name]", qFilterName); err != nil {
				return err
			}
		}
	}

	if o.FilterSlug != nil {

		// query param filter[slug]
		var qrFilterSlug string

		if o.FilterSlug != nil {
			qrFilterSlug = *o.FilterSlug
		}
		qFilterSlug := qrFilterSlug
		if qFilterSlug != "" {

			if err := r.SetQueryParam("filter[slug]", qFilterSlug); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
