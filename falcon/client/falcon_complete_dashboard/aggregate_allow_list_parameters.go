// Code generated by go-swagger; DO NOT EDIT.

package falcon_complete_dashboard

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

// NewAggregateAllowListParams creates a new AggregateAllowListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAggregateAllowListParams() *AggregateAllowListParams {
	return &AggregateAllowListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAggregateAllowListParamsWithTimeout creates a new AggregateAllowListParams object
// with the ability to set a timeout on a request.
func NewAggregateAllowListParamsWithTimeout(timeout time.Duration) *AggregateAllowListParams {
	return &AggregateAllowListParams{
		timeout: timeout,
	}
}

// NewAggregateAllowListParamsWithContext creates a new AggregateAllowListParams object
// with the ability to set a context for a request.
func NewAggregateAllowListParamsWithContext(ctx context.Context) *AggregateAllowListParams {
	return &AggregateAllowListParams{
		Context: ctx,
	}
}

// NewAggregateAllowListParamsWithHTTPClient creates a new AggregateAllowListParams object
// with the ability to set a custom HTTPClient for a request.
func NewAggregateAllowListParamsWithHTTPClient(client *http.Client) *AggregateAllowListParams {
	return &AggregateAllowListParams{
		HTTPClient: client,
	}
}

/*
AggregateAllowListParams contains all the parameters to send to the API endpoint

	for the aggregate allow list operation.

	Typically these are written to a http.Request.
*/
type AggregateAllowListParams struct {

	// Body.
	Body []*models.MsaAggregateQueryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the aggregate allow list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateAllowListParams) WithDefaults() *AggregateAllowListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the aggregate allow list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateAllowListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the aggregate allow list params
func (o *AggregateAllowListParams) WithTimeout(timeout time.Duration) *AggregateAllowListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the aggregate allow list params
func (o *AggregateAllowListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the aggregate allow list params
func (o *AggregateAllowListParams) WithContext(ctx context.Context) *AggregateAllowListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the aggregate allow list params
func (o *AggregateAllowListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the aggregate allow list params
func (o *AggregateAllowListParams) WithHTTPClient(client *http.Client) *AggregateAllowListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the aggregate allow list params
func (o *AggregateAllowListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the aggregate allow list params
func (o *AggregateAllowListParams) WithBody(body []*models.MsaAggregateQueryRequest) *AggregateAllowListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the aggregate allow list params
func (o *AggregateAllowListParams) SetBody(body []*models.MsaAggregateQueryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AggregateAllowListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
