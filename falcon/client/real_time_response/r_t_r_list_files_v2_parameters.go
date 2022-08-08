// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

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

// NewRTRListFilesV2Params creates a new RTRListFilesV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRTRListFilesV2Params() *RTRListFilesV2Params {
	return &RTRListFilesV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewRTRListFilesV2ParamsWithTimeout creates a new RTRListFilesV2Params object
// with the ability to set a timeout on a request.
func NewRTRListFilesV2ParamsWithTimeout(timeout time.Duration) *RTRListFilesV2Params {
	return &RTRListFilesV2Params{
		timeout: timeout,
	}
}

// NewRTRListFilesV2ParamsWithContext creates a new RTRListFilesV2Params object
// with the ability to set a context for a request.
func NewRTRListFilesV2ParamsWithContext(ctx context.Context) *RTRListFilesV2Params {
	return &RTRListFilesV2Params{
		Context: ctx,
	}
}

// NewRTRListFilesV2ParamsWithHTTPClient creates a new RTRListFilesV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewRTRListFilesV2ParamsWithHTTPClient(client *http.Client) *RTRListFilesV2Params {
	return &RTRListFilesV2Params{
		HTTPClient: client,
	}
}

/*
RTRListFilesV2Params contains all the parameters to send to the API endpoint

	for the r t r list files v2 operation.

	Typically these are written to a http.Request.
*/
type RTRListFilesV2Params struct {

	/* SessionID.

	   RTR Session id
	*/
	SessionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the r t r list files v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRListFilesV2Params) WithDefaults() *RTRListFilesV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the r t r list files v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRListFilesV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the r t r list files v2 params
func (o *RTRListFilesV2Params) WithTimeout(timeout time.Duration) *RTRListFilesV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the r t r list files v2 params
func (o *RTRListFilesV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the r t r list files v2 params
func (o *RTRListFilesV2Params) WithContext(ctx context.Context) *RTRListFilesV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the r t r list files v2 params
func (o *RTRListFilesV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the r t r list files v2 params
func (o *RTRListFilesV2Params) WithHTTPClient(client *http.Client) *RTRListFilesV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the r t r list files v2 params
func (o *RTRListFilesV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSessionID adds the sessionID to the r t r list files v2 params
func (o *RTRListFilesV2Params) WithSessionID(sessionID string) *RTRListFilesV2Params {
	o.SetSessionID(sessionID)
	return o
}

// SetSessionID adds the sessionId to the r t r list files v2 params
func (o *RTRListFilesV2Params) SetSessionID(sessionID string) {
	o.SessionID = sessionID
}

// WriteToRequest writes these params to a swagger request
func (o *RTRListFilesV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param session_id
	qrSessionID := o.SessionID
	qSessionID := qrSessionID
	if qSessionID != "" {

		if err := r.SetQueryParam("session_id", qSessionID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
