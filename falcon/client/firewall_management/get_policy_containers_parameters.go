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
	"github.com/go-openapi/swag"
)

// NewGetPolicyContainersParams creates a new GetPolicyContainersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPolicyContainersParams() *GetPolicyContainersParams {
	return &GetPolicyContainersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPolicyContainersParamsWithTimeout creates a new GetPolicyContainersParams object
// with the ability to set a timeout on a request.
func NewGetPolicyContainersParamsWithTimeout(timeout time.Duration) *GetPolicyContainersParams {
	return &GetPolicyContainersParams{
		timeout: timeout,
	}
}

// NewGetPolicyContainersParamsWithContext creates a new GetPolicyContainersParams object
// with the ability to set a context for a request.
func NewGetPolicyContainersParamsWithContext(ctx context.Context) *GetPolicyContainersParams {
	return &GetPolicyContainersParams{
		Context: ctx,
	}
}

// NewGetPolicyContainersParamsWithHTTPClient creates a new GetPolicyContainersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPolicyContainersParamsWithHTTPClient(client *http.Client) *GetPolicyContainersParams {
	return &GetPolicyContainersParams{
		HTTPClient: client,
	}
}

/*
GetPolicyContainersParams contains all the parameters to send to the API endpoint

	for the get policy containers operation.

	Typically these are written to a http.Request.
*/
type GetPolicyContainersParams struct {

	/* Ids.

	   The policy container(s) to retrieve, identified by policy ID
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get policy containers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPolicyContainersParams) WithDefaults() *GetPolicyContainersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get policy containers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPolicyContainersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get policy containers params
func (o *GetPolicyContainersParams) WithTimeout(timeout time.Duration) *GetPolicyContainersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policy containers params
func (o *GetPolicyContainersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policy containers params
func (o *GetPolicyContainersParams) WithContext(ctx context.Context) *GetPolicyContainersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policy containers params
func (o *GetPolicyContainersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policy containers params
func (o *GetPolicyContainersParams) WithHTTPClient(client *http.Client) *GetPolicyContainersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policy containers params
func (o *GetPolicyContainersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get policy containers params
func (o *GetPolicyContainersParams) WithIds(ids []string) *GetPolicyContainersParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get policy containers params
func (o *GetPolicyContainersParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetPolicyContainersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamGetPolicyContainers binds the parameter ids
func (o *GetPolicyContainersParams) bindParamIds(formats strfmt.Registry) []string {
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
