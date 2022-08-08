// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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

// NewAggregateRulesParams creates a new AggregateRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAggregateRulesParams() *AggregateRulesParams {
	return &AggregateRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAggregateRulesParamsWithTimeout creates a new AggregateRulesParams object
// with the ability to set a timeout on a request.
func NewAggregateRulesParamsWithTimeout(timeout time.Duration) *AggregateRulesParams {
	return &AggregateRulesParams{
		timeout: timeout,
	}
}

// NewAggregateRulesParamsWithContext creates a new AggregateRulesParams object
// with the ability to set a context for a request.
func NewAggregateRulesParamsWithContext(ctx context.Context) *AggregateRulesParams {
	return &AggregateRulesParams{
		Context: ctx,
	}
}

// NewAggregateRulesParamsWithHTTPClient creates a new AggregateRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewAggregateRulesParamsWithHTTPClient(client *http.Client) *AggregateRulesParams {
	return &AggregateRulesParams{
		HTTPClient: client,
	}
}

/*
AggregateRulesParams contains all the parameters to send to the API endpoint

	for the aggregate rules operation.

	Typically these are written to a http.Request.
*/
type AggregateRulesParams struct {

	/* Body.

	   Query criteria and settings
	*/
	Body []*models.FwmgrMsaAggregateQueryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the aggregate rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateRulesParams) WithDefaults() *AggregateRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the aggregate rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the aggregate rules params
func (o *AggregateRulesParams) WithTimeout(timeout time.Duration) *AggregateRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the aggregate rules params
func (o *AggregateRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the aggregate rules params
func (o *AggregateRulesParams) WithContext(ctx context.Context) *AggregateRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the aggregate rules params
func (o *AggregateRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the aggregate rules params
func (o *AggregateRulesParams) WithHTTPClient(client *http.Client) *AggregateRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the aggregate rules params
func (o *AggregateRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the aggregate rules params
func (o *AggregateRulesParams) WithBody(body []*models.FwmgrMsaAggregateQueryRequest) *AggregateRulesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the aggregate rules params
func (o *AggregateRulesParams) SetBody(body []*models.FwmgrMsaAggregateQueryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AggregateRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
