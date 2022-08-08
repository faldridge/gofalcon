// Code generated by go-swagger; DO NOT EDIT.

package ioa_exclusions

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

// NewDeleteIOAExclusionsV1Params creates a new DeleteIOAExclusionsV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteIOAExclusionsV1Params() *DeleteIOAExclusionsV1Params {
	return &DeleteIOAExclusionsV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIOAExclusionsV1ParamsWithTimeout creates a new DeleteIOAExclusionsV1Params object
// with the ability to set a timeout on a request.
func NewDeleteIOAExclusionsV1ParamsWithTimeout(timeout time.Duration) *DeleteIOAExclusionsV1Params {
	return &DeleteIOAExclusionsV1Params{
		timeout: timeout,
	}
}

// NewDeleteIOAExclusionsV1ParamsWithContext creates a new DeleteIOAExclusionsV1Params object
// with the ability to set a context for a request.
func NewDeleteIOAExclusionsV1ParamsWithContext(ctx context.Context) *DeleteIOAExclusionsV1Params {
	return &DeleteIOAExclusionsV1Params{
		Context: ctx,
	}
}

// NewDeleteIOAExclusionsV1ParamsWithHTTPClient creates a new DeleteIOAExclusionsV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteIOAExclusionsV1ParamsWithHTTPClient(client *http.Client) *DeleteIOAExclusionsV1Params {
	return &DeleteIOAExclusionsV1Params{
		HTTPClient: client,
	}
}

/*
DeleteIOAExclusionsV1Params contains all the parameters to send to the API endpoint

	for the delete i o a exclusions v1 operation.

	Typically these are written to a http.Request.
*/
type DeleteIOAExclusionsV1Params struct {

	/* Comment.

	   Explains why this exclusions was deleted
	*/
	Comment *string

	/* Ids.

	   The ids of the exclusions to delete
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete i o a exclusions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIOAExclusionsV1Params) WithDefaults() *DeleteIOAExclusionsV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete i o a exclusions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIOAExclusionsV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete i o a exclusions v1 params
func (o *DeleteIOAExclusionsV1Params) WithTimeout(timeout time.Duration) *DeleteIOAExclusionsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete i o a exclusions v1 params
func (o *DeleteIOAExclusionsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete i o a exclusions v1 params
func (o *DeleteIOAExclusionsV1Params) WithContext(ctx context.Context) *DeleteIOAExclusionsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete i o a exclusions v1 params
func (o *DeleteIOAExclusionsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete i o a exclusions v1 params
func (o *DeleteIOAExclusionsV1Params) WithHTTPClient(client *http.Client) *DeleteIOAExclusionsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete i o a exclusions v1 params
func (o *DeleteIOAExclusionsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComment adds the comment to the delete i o a exclusions v1 params
func (o *DeleteIOAExclusionsV1Params) WithComment(comment *string) *DeleteIOAExclusionsV1Params {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the delete i o a exclusions v1 params
func (o *DeleteIOAExclusionsV1Params) SetComment(comment *string) {
	o.Comment = comment
}

// WithIds adds the ids to the delete i o a exclusions v1 params
func (o *DeleteIOAExclusionsV1Params) WithIds(ids []string) *DeleteIOAExclusionsV1Params {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the delete i o a exclusions v1 params
func (o *DeleteIOAExclusionsV1Params) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIOAExclusionsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Comment != nil {

		// query param comment
		var qrComment string

		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

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

// bindParamDeleteIOAExclusionsV1 binds the parameter ids
func (o *DeleteIOAExclusionsV1Params) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "csv"
	idsIS := swag.JoinByFormat(idsIC, "csv")

	return idsIS
}
