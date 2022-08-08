// Code generated by go-swagger; DO NOT EDIT.

package discover

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

// NewGetAccountsParams creates a new GetAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountsParams() *GetAccountsParams {
	return &GetAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountsParamsWithTimeout creates a new GetAccountsParams object
// with the ability to set a timeout on a request.
func NewGetAccountsParamsWithTimeout(timeout time.Duration) *GetAccountsParams {
	return &GetAccountsParams{
		timeout: timeout,
	}
}

// NewGetAccountsParamsWithContext creates a new GetAccountsParams object
// with the ability to set a context for a request.
func NewGetAccountsParamsWithContext(ctx context.Context) *GetAccountsParams {
	return &GetAccountsParams{
		Context: ctx,
	}
}

// NewGetAccountsParamsWithHTTPClient creates a new GetAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountsParamsWithHTTPClient(client *http.Client) *GetAccountsParams {
	return &GetAccountsParams{
		HTTPClient: client,
	}
}

/*
GetAccountsParams contains all the parameters to send to the API endpoint

	for the get accounts operation.

	Typically these are written to a http.Request.
*/
type GetAccountsParams struct {

	/* Ids.

	   One or more account IDs (max: 100). Find account IDs with GET `/discover/queries/accounts/v1`
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountsParams) WithDefaults() *GetAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get accounts params
func (o *GetAccountsParams) WithTimeout(timeout time.Duration) *GetAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get accounts params
func (o *GetAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get accounts params
func (o *GetAccountsParams) WithContext(ctx context.Context) *GetAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get accounts params
func (o *GetAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get accounts params
func (o *GetAccountsParams) WithHTTPClient(client *http.Client) *GetAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get accounts params
func (o *GetAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get accounts params
func (o *GetAccountsParams) WithIds(ids []string) *GetAccountsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get accounts params
func (o *GetAccountsParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetAccounts binds the parameter ids
func (o *GetAccountsParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "multi"
	idsIS := swag.JoinByFormat(idsIC, "multi")

	return idsIS
}
