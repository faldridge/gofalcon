// Code generated by go-swagger; DO NOT EDIT.

package d4c_registration

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

// NewGetCSPMGCPUserScriptsAttachmentParams creates a new GetCSPMGCPUserScriptsAttachmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCSPMGCPUserScriptsAttachmentParams() *GetCSPMGCPUserScriptsAttachmentParams {
	return &GetCSPMGCPUserScriptsAttachmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCSPMGCPUserScriptsAttachmentParamsWithTimeout creates a new GetCSPMGCPUserScriptsAttachmentParams object
// with the ability to set a timeout on a request.
func NewGetCSPMGCPUserScriptsAttachmentParamsWithTimeout(timeout time.Duration) *GetCSPMGCPUserScriptsAttachmentParams {
	return &GetCSPMGCPUserScriptsAttachmentParams{
		timeout: timeout,
	}
}

// NewGetCSPMGCPUserScriptsAttachmentParamsWithContext creates a new GetCSPMGCPUserScriptsAttachmentParams object
// with the ability to set a context for a request.
func NewGetCSPMGCPUserScriptsAttachmentParamsWithContext(ctx context.Context) *GetCSPMGCPUserScriptsAttachmentParams {
	return &GetCSPMGCPUserScriptsAttachmentParams{
		Context: ctx,
	}
}

// NewGetCSPMGCPUserScriptsAttachmentParamsWithHTTPClient creates a new GetCSPMGCPUserScriptsAttachmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCSPMGCPUserScriptsAttachmentParamsWithHTTPClient(client *http.Client) *GetCSPMGCPUserScriptsAttachmentParams {
	return &GetCSPMGCPUserScriptsAttachmentParams{
		HTTPClient: client,
	}
}

/*
GetCSPMGCPUserScriptsAttachmentParams contains all the parameters to send to the API endpoint

	for the get c s p m g c p user scripts attachment operation.

	Typically these are written to a http.Request.
*/
type GetCSPMGCPUserScriptsAttachmentParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get c s p m g c p user scripts attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCSPMGCPUserScriptsAttachmentParams) WithDefaults() *GetCSPMGCPUserScriptsAttachmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get c s p m g c p user scripts attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCSPMGCPUserScriptsAttachmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get c s p m g c p user scripts attachment params
func (o *GetCSPMGCPUserScriptsAttachmentParams) WithTimeout(timeout time.Duration) *GetCSPMGCPUserScriptsAttachmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get c s p m g c p user scripts attachment params
func (o *GetCSPMGCPUserScriptsAttachmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get c s p m g c p user scripts attachment params
func (o *GetCSPMGCPUserScriptsAttachmentParams) WithContext(ctx context.Context) *GetCSPMGCPUserScriptsAttachmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get c s p m g c p user scripts attachment params
func (o *GetCSPMGCPUserScriptsAttachmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get c s p m g c p user scripts attachment params
func (o *GetCSPMGCPUserScriptsAttachmentParams) WithHTTPClient(client *http.Client) *GetCSPMGCPUserScriptsAttachmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get c s p m g c p user scripts attachment params
func (o *GetCSPMGCPUserScriptsAttachmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCSPMGCPUserScriptsAttachmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
