// Code generated by go-swagger; DO NOT EDIT.

package sensor_update_policies

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

// NewUpdateSensorUpdatePoliciesParams creates a new UpdateSensorUpdatePoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateSensorUpdatePoliciesParams() *UpdateSensorUpdatePoliciesParams {
	return &UpdateSensorUpdatePoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSensorUpdatePoliciesParamsWithTimeout creates a new UpdateSensorUpdatePoliciesParams object
// with the ability to set a timeout on a request.
func NewUpdateSensorUpdatePoliciesParamsWithTimeout(timeout time.Duration) *UpdateSensorUpdatePoliciesParams {
	return &UpdateSensorUpdatePoliciesParams{
		timeout: timeout,
	}
}

// NewUpdateSensorUpdatePoliciesParamsWithContext creates a new UpdateSensorUpdatePoliciesParams object
// with the ability to set a context for a request.
func NewUpdateSensorUpdatePoliciesParamsWithContext(ctx context.Context) *UpdateSensorUpdatePoliciesParams {
	return &UpdateSensorUpdatePoliciesParams{
		Context: ctx,
	}
}

// NewUpdateSensorUpdatePoliciesParamsWithHTTPClient creates a new UpdateSensorUpdatePoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateSensorUpdatePoliciesParamsWithHTTPClient(client *http.Client) *UpdateSensorUpdatePoliciesParams {
	return &UpdateSensorUpdatePoliciesParams{
		HTTPClient: client,
	}
}

/*
UpdateSensorUpdatePoliciesParams contains all the parameters to send to the API endpoint

	for the update sensor update policies operation.

	Typically these are written to a http.Request.
*/
type UpdateSensorUpdatePoliciesParams struct {

	// Body.
	Body *models.RequestsUpdateSensorUpdatePoliciesV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update sensor update policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSensorUpdatePoliciesParams) WithDefaults() *UpdateSensorUpdatePoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update sensor update policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSensorUpdatePoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update sensor update policies params
func (o *UpdateSensorUpdatePoliciesParams) WithTimeout(timeout time.Duration) *UpdateSensorUpdatePoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update sensor update policies params
func (o *UpdateSensorUpdatePoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update sensor update policies params
func (o *UpdateSensorUpdatePoliciesParams) WithContext(ctx context.Context) *UpdateSensorUpdatePoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update sensor update policies params
func (o *UpdateSensorUpdatePoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update sensor update policies params
func (o *UpdateSensorUpdatePoliciesParams) WithHTTPClient(client *http.Client) *UpdateSensorUpdatePoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update sensor update policies params
func (o *UpdateSensorUpdatePoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update sensor update policies params
func (o *UpdateSensorUpdatePoliciesParams) WithBody(body *models.RequestsUpdateSensorUpdatePoliciesV1) *UpdateSensorUpdatePoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update sensor update policies params
func (o *UpdateSensorUpdatePoliciesParams) SetBody(body *models.RequestsUpdateSensorUpdatePoliciesV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSensorUpdatePoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
