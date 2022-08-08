// Code generated by go-swagger; DO NOT EDIT.

package ml_exclusions

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewCreateMLExclusionsV1Params creates a new CreateMLExclusionsV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateMLExclusionsV1Params() *CreateMLExclusionsV1Params {
	return &CreateMLExclusionsV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMLExclusionsV1ParamsWithTimeout creates a new CreateMLExclusionsV1Params object
// with the ability to set a timeout on a request.
func NewCreateMLExclusionsV1ParamsWithTimeout(timeout time.Duration) *CreateMLExclusionsV1Params {
	return &CreateMLExclusionsV1Params{
		timeout: timeout,
	}
}

// NewCreateMLExclusionsV1ParamsWithContext creates a new CreateMLExclusionsV1Params object
// with the ability to set a context for a request.
func NewCreateMLExclusionsV1ParamsWithContext(ctx context.Context) *CreateMLExclusionsV1Params {
	return &CreateMLExclusionsV1Params{
		Context: ctx,
	}
}

// NewCreateMLExclusionsV1ParamsWithHTTPClient creates a new CreateMLExclusionsV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateMLExclusionsV1ParamsWithHTTPClient(client *http.Client) *CreateMLExclusionsV1Params {
	return &CreateMLExclusionsV1Params{
		HTTPClient: client,
	}
}

/*
CreateMLExclusionsV1Params contains all the parameters to send to the API endpoint

	for the create m l exclusions v1 operation.

	Typically these are written to a http.Request.
*/
type CreateMLExclusionsV1Params struct {

	// Body.
	Body *models.RequestsMlExclusionCreateReqV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create m l exclusions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMLExclusionsV1Params) WithDefaults() *CreateMLExclusionsV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create m l exclusions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMLExclusionsV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create m l exclusions v1 params
func (o *CreateMLExclusionsV1Params) WithTimeout(timeout time.Duration) *CreateMLExclusionsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create m l exclusions v1 params
func (o *CreateMLExclusionsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create m l exclusions v1 params
func (o *CreateMLExclusionsV1Params) WithContext(ctx context.Context) *CreateMLExclusionsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create m l exclusions v1 params
func (o *CreateMLExclusionsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create m l exclusions v1 params
func (o *CreateMLExclusionsV1Params) WithHTTPClient(client *http.Client) *CreateMLExclusionsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create m l exclusions v1 params
func (o *CreateMLExclusionsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create m l exclusions v1 params
func (o *CreateMLExclusionsV1Params) WithBody(body *models.RequestsMlExclusionCreateReqV1) *CreateMLExclusionsV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create m l exclusions v1 params
func (o *CreateMLExclusionsV1Params) SetBody(body *models.RequestsMlExclusionCreateReqV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMLExclusionsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
