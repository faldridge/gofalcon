// Code generated by go-swagger; DO NOT EDIT.

package host_group

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewPerformGroupActionParams creates a new PerformGroupActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformGroupActionParams() *PerformGroupActionParams {
	return &PerformGroupActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformGroupActionParamsWithTimeout creates a new PerformGroupActionParams object
// with the ability to set a timeout on a request.
func NewPerformGroupActionParamsWithTimeout(timeout time.Duration) *PerformGroupActionParams {
	return &PerformGroupActionParams{
		timeout: timeout,
	}
}

// NewPerformGroupActionParamsWithContext creates a new PerformGroupActionParams object
// with the ability to set a context for a request.
func NewPerformGroupActionParamsWithContext(ctx context.Context) *PerformGroupActionParams {
	return &PerformGroupActionParams{
		Context: ctx,
	}
}

// NewPerformGroupActionParamsWithHTTPClient creates a new PerformGroupActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformGroupActionParamsWithHTTPClient(client *http.Client) *PerformGroupActionParams {
	return &PerformGroupActionParams{
		HTTPClient: client,
	}
}

/*
PerformGroupActionParams contains all the parameters to send to the API endpoint

	for the perform group action operation.

	Typically these are written to a http.Request.
*/
type PerformGroupActionParams struct {

	/* ActionName.

	   The action to perform
	*/
	ActionName string

	// Body.
	Body *models.MsaEntityActionRequestV2

	/* DisableHostnameCheck.

	   Bool to disable hostname check on add-member
	*/
	DisableHostnameCheck *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the perform group action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformGroupActionParams) WithDefaults() *PerformGroupActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the perform group action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformGroupActionParams) SetDefaults() {
	var (
		disableHostnameCheckDefault = bool(false)
	)

	val := PerformGroupActionParams{
		DisableHostnameCheck: &disableHostnameCheckDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the perform group action params
func (o *PerformGroupActionParams) WithTimeout(timeout time.Duration) *PerformGroupActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the perform group action params
func (o *PerformGroupActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the perform group action params
func (o *PerformGroupActionParams) WithContext(ctx context.Context) *PerformGroupActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the perform group action params
func (o *PerformGroupActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the perform group action params
func (o *PerformGroupActionParams) WithHTTPClient(client *http.Client) *PerformGroupActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the perform group action params
func (o *PerformGroupActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionName adds the actionName to the perform group action params
func (o *PerformGroupActionParams) WithActionName(actionName string) *PerformGroupActionParams {
	o.SetActionName(actionName)
	return o
}

// SetActionName adds the actionName to the perform group action params
func (o *PerformGroupActionParams) SetActionName(actionName string) {
	o.ActionName = actionName
}

// WithBody adds the body to the perform group action params
func (o *PerformGroupActionParams) WithBody(body *models.MsaEntityActionRequestV2) *PerformGroupActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the perform group action params
func (o *PerformGroupActionParams) SetBody(body *models.MsaEntityActionRequestV2) {
	o.Body = body
}

// WithDisableHostnameCheck adds the disableHostnameCheck to the perform group action params
func (o *PerformGroupActionParams) WithDisableHostnameCheck(disableHostnameCheck *bool) *PerformGroupActionParams {
	o.SetDisableHostnameCheck(disableHostnameCheck)
	return o
}

// SetDisableHostnameCheck adds the disableHostnameCheck to the perform group action params
func (o *PerformGroupActionParams) SetDisableHostnameCheck(disableHostnameCheck *bool) {
	o.DisableHostnameCheck = disableHostnameCheck
}

// WriteToRequest writes these params to a swagger request
func (o *PerformGroupActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param action_name
	qrActionName := o.ActionName
	qActionName := qrActionName
	if qActionName != "" {

		if err := r.SetQueryParam("action_name", qActionName); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.DisableHostnameCheck != nil {

		// query param disable_hostname_check
		var qrDisableHostnameCheck bool

		if o.DisableHostnameCheck != nil {
			qrDisableHostnameCheck = *o.DisableHostnameCheck
		}
		qDisableHostnameCheck := swag.FormatBool(qrDisableHostnameCheck)
		if qDisableHostnameCheck != "" {

			if err := r.SetQueryParam("disable_hostname_check", qDisableHostnameCheck); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
