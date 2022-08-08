// Code generated by go-swagger; DO NOT EDIT.

package recon

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

// NewUpdateActionV1Params creates a new UpdateActionV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateActionV1Params() *UpdateActionV1Params {
	return &UpdateActionV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateActionV1ParamsWithTimeout creates a new UpdateActionV1Params object
// with the ability to set a timeout on a request.
func NewUpdateActionV1ParamsWithTimeout(timeout time.Duration) *UpdateActionV1Params {
	return &UpdateActionV1Params{
		timeout: timeout,
	}
}

// NewUpdateActionV1ParamsWithContext creates a new UpdateActionV1Params object
// with the ability to set a context for a request.
func NewUpdateActionV1ParamsWithContext(ctx context.Context) *UpdateActionV1Params {
	return &UpdateActionV1Params{
		Context: ctx,
	}
}

// NewUpdateActionV1ParamsWithHTTPClient creates a new UpdateActionV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateActionV1ParamsWithHTTPClient(client *http.Client) *UpdateActionV1Params {
	return &UpdateActionV1Params{
		HTTPClient: client,
	}
}

/*
UpdateActionV1Params contains all the parameters to send to the API endpoint

	for the update action v1 operation.

	Typically these are written to a http.Request.
*/
type UpdateActionV1Params struct {

	// Body.
	Body *models.DomainUpdateActionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update action v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateActionV1Params) WithDefaults() *UpdateActionV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update action v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateActionV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update action v1 params
func (o *UpdateActionV1Params) WithTimeout(timeout time.Duration) *UpdateActionV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update action v1 params
func (o *UpdateActionV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update action v1 params
func (o *UpdateActionV1Params) WithContext(ctx context.Context) *UpdateActionV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update action v1 params
func (o *UpdateActionV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update action v1 params
func (o *UpdateActionV1Params) WithHTTPClient(client *http.Client) *UpdateActionV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update action v1 params
func (o *UpdateActionV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update action v1 params
func (o *UpdateActionV1Params) WithBody(body *models.DomainUpdateActionRequest) *UpdateActionV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update action v1 params
func (o *UpdateActionV1Params) SetBody(body *models.DomainUpdateActionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateActionV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
