// Code generated by go-swagger; DO NOT EDIT.

package response_policies

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
	"github.com/go-openapi/swag"
)

// NewQueryRTResponsePoliciesParams creates a new QueryRTResponsePoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryRTResponsePoliciesParams() *QueryRTResponsePoliciesParams {
	return &QueryRTResponsePoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryRTResponsePoliciesParamsWithTimeout creates a new QueryRTResponsePoliciesParams object
// with the ability to set a timeout on a request.
func NewQueryRTResponsePoliciesParamsWithTimeout(timeout time.Duration) *QueryRTResponsePoliciesParams {
	return &QueryRTResponsePoliciesParams{
		timeout: timeout,
	}
}

// NewQueryRTResponsePoliciesParamsWithContext creates a new QueryRTResponsePoliciesParams object
// with the ability to set a context for a request.
func NewQueryRTResponsePoliciesParamsWithContext(ctx context.Context) *QueryRTResponsePoliciesParams {
	return &QueryRTResponsePoliciesParams{
		Context: ctx,
	}
}

// NewQueryRTResponsePoliciesParamsWithHTTPClient creates a new QueryRTResponsePoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryRTResponsePoliciesParamsWithHTTPClient(client *http.Client) *QueryRTResponsePoliciesParams {
	return &QueryRTResponsePoliciesParams{
		HTTPClient: client,
	}
}

/*
QueryRTResponsePoliciesParams contains all the parameters to send to the API endpoint

	for the query r t response policies operation.

	Typically these are written to a http.Request.
*/
type QueryRTResponsePoliciesParams struct {

	/* Filter.

	   The filter expression that should be used to determine the results.
	*/
	Filter *string

	/* Limit.

	   The maximum number of records to return [1-5000]
	*/
	Limit *int64

	/* Offset.

	   The offset of the first record to retrieve from
	*/
	Offset *int64

	/* Sort.

	   The property to sort results by
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query r t response policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryRTResponsePoliciesParams) WithDefaults() *QueryRTResponsePoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query r t response policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryRTResponsePoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query r t response policies params
func (o *QueryRTResponsePoliciesParams) WithTimeout(timeout time.Duration) *QueryRTResponsePoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query r t response policies params
func (o *QueryRTResponsePoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query r t response policies params
func (o *QueryRTResponsePoliciesParams) WithContext(ctx context.Context) *QueryRTResponsePoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query r t response policies params
func (o *QueryRTResponsePoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query r t response policies params
func (o *QueryRTResponsePoliciesParams) WithHTTPClient(client *http.Client) *QueryRTResponsePoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query r t response policies params
func (o *QueryRTResponsePoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the query r t response policies params
func (o *QueryRTResponsePoliciesParams) WithFilter(filter *string) *QueryRTResponsePoliciesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query r t response policies params
func (o *QueryRTResponsePoliciesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query r t response policies params
func (o *QueryRTResponsePoliciesParams) WithLimit(limit *int64) *QueryRTResponsePoliciesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query r t response policies params
func (o *QueryRTResponsePoliciesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query r t response policies params
func (o *QueryRTResponsePoliciesParams) WithOffset(offset *int64) *QueryRTResponsePoliciesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query r t response policies params
func (o *QueryRTResponsePoliciesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the query r t response policies params
func (o *QueryRTResponsePoliciesParams) WithSort(sort *string) *QueryRTResponsePoliciesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query r t response policies params
func (o *QueryRTResponsePoliciesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryRTResponsePoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
