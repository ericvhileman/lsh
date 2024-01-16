// Code generated by go-swagger; DO NOT EDIT.

package ssh_keys

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

// NewGetProjectSSHKeyParams creates a new GetProjectSSHKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectSSHKeyParams() *GetProjectSSHKeyParams {
	return &GetProjectSSHKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectSSHKeyParamsWithTimeout creates a new GetProjectSSHKeyParams object
// with the ability to set a timeout on a request.
func NewGetProjectSSHKeyParamsWithTimeout(timeout time.Duration) *GetProjectSSHKeyParams {
	return &GetProjectSSHKeyParams{
		timeout: timeout,
	}
}

// NewGetProjectSSHKeyParamsWithContext creates a new GetProjectSSHKeyParams object
// with the ability to set a context for a request.
func NewGetProjectSSHKeyParamsWithContext(ctx context.Context) *GetProjectSSHKeyParams {
	return &GetProjectSSHKeyParams{
		Context: ctx,
	}
}

// NewGetProjectSSHKeyParamsWithHTTPClient creates a new GetProjectSSHKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectSSHKeyParamsWithHTTPClient(client *http.Client) *GetProjectSSHKeyParams {
	return &GetProjectSSHKeyParams{
		HTTPClient: client,
	}
}

/*
GetProjectSSHKeyParams contains all the parameters to send to the API endpoint

	for the get project ssh key operation.

	Typically these are written to a http.Request.
*/
type GetProjectSSHKeyParams struct {

	// APIVersion.
	//
	// Default: "2023-06-01"
	APIVersion *string

	// ProjectIDOrSlug.
	ProjectIDOrSlug string

	// SSHKeyID.
	SSHKeyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project ssh key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectSSHKeyParams) WithDefaults() *GetProjectSSHKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project ssh key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectSSHKeyParams) SetDefaults() {
	var (
		aPIVersionDefault = string("2023-06-01")
	)

	val := GetProjectSSHKeyParams{
		APIVersion: &aPIVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get project ssh key params
func (o *GetProjectSSHKeyParams) WithTimeout(timeout time.Duration) *GetProjectSSHKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project ssh key params
func (o *GetProjectSSHKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project ssh key params
func (o *GetProjectSSHKeyParams) WithContext(ctx context.Context) *GetProjectSSHKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project ssh key params
func (o *GetProjectSSHKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project ssh key params
func (o *GetProjectSSHKeyParams) WithHTTPClient(client *http.Client) *GetProjectSSHKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project ssh key params
func (o *GetProjectSSHKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get project ssh key params
func (o *GetProjectSSHKeyParams) WithAPIVersion(aPIVersion *string) *GetProjectSSHKeyParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get project ssh key params
func (o *GetProjectSSHKeyParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithProjectIDOrSlug adds the projectIDOrSlug to the get project ssh key params
func (o *GetProjectSSHKeyParams) WithProjectIDOrSlug(projectIDOrSlug string) *GetProjectSSHKeyParams {
	o.SetProjectIDOrSlug(projectIDOrSlug)
	return o
}

// SetProjectIDOrSlug adds the projectIdOrSlug to the get project ssh key params
func (o *GetProjectSSHKeyParams) SetProjectIDOrSlug(projectIDOrSlug string) {
	o.ProjectIDOrSlug = projectIDOrSlug
}

// WithSSHKeyID adds the sSHKeyID to the get project ssh key params
func (o *GetProjectSSHKeyParams) WithSSHKeyID(sSHKeyID string) *GetProjectSSHKeyParams {
	o.SetSSHKeyID(sSHKeyID)
	return o
}

// SetSSHKeyID adds the sshKeyId to the get project ssh key params
func (o *GetProjectSSHKeyParams) SetSSHKeyID(sSHKeyID string) {
	o.SSHKeyID = sSHKeyID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectSSHKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project_id_or_slug
	if err := r.SetPathParam("project_id_or_slug", o.ProjectIDOrSlug); err != nil {
		return err
	}

	// path param ssh_key_id
	if err := r.SetPathParam("ssh_key_id", o.SSHKeyID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}