// Code generated by go-swagger; DO NOT EDIT.

package firewall_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// QueryCombinedFirewallPoliciesReader is a Reader for the QueryCombinedFirewallPolicies structure.
type QueryCombinedFirewallPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryCombinedFirewallPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryCombinedFirewallPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryCombinedFirewallPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryCombinedFirewallPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryCombinedFirewallPoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryCombinedFirewallPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryCombinedFirewallPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryCombinedFirewallPoliciesOK creates a QueryCombinedFirewallPoliciesOK with default headers values
func NewQueryCombinedFirewallPoliciesOK() *QueryCombinedFirewallPoliciesOK {
	return &QueryCombinedFirewallPoliciesOK{}
}

/*
	QueryCombinedFirewallPoliciesOK describes a response with status code 200, with default header values.

OK
*/
type QueryCombinedFirewallPoliciesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

func (o *QueryCombinedFirewallPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall/v1][%d] queryCombinedFirewallPoliciesOK  %+v", 200, o.Payload)
}
func (o *QueryCombinedFirewallPoliciesOK) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *QueryCombinedFirewallPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedFirewallPoliciesBadRequest creates a QueryCombinedFirewallPoliciesBadRequest with default headers values
func NewQueryCombinedFirewallPoliciesBadRequest() *QueryCombinedFirewallPoliciesBadRequest {
	return &QueryCombinedFirewallPoliciesBadRequest{}
}

/*
	QueryCombinedFirewallPoliciesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryCombinedFirewallPoliciesBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

func (o *QueryCombinedFirewallPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall/v1][%d] queryCombinedFirewallPoliciesBadRequest  %+v", 400, o.Payload)
}
func (o *QueryCombinedFirewallPoliciesBadRequest) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *QueryCombinedFirewallPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedFirewallPoliciesForbidden creates a QueryCombinedFirewallPoliciesForbidden with default headers values
func NewQueryCombinedFirewallPoliciesForbidden() *QueryCombinedFirewallPoliciesForbidden {
	return &QueryCombinedFirewallPoliciesForbidden{}
}

/*
	QueryCombinedFirewallPoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryCombinedFirewallPoliciesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *QueryCombinedFirewallPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall/v1][%d] queryCombinedFirewallPoliciesForbidden  %+v", 403, o.Payload)
}
func (o *QueryCombinedFirewallPoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryCombinedFirewallPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedFirewallPoliciesTooManyRequests creates a QueryCombinedFirewallPoliciesTooManyRequests with default headers values
func NewQueryCombinedFirewallPoliciesTooManyRequests() *QueryCombinedFirewallPoliciesTooManyRequests {
	return &QueryCombinedFirewallPoliciesTooManyRequests{}
}

/*
	QueryCombinedFirewallPoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryCombinedFirewallPoliciesTooManyRequests struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

func (o *QueryCombinedFirewallPoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall/v1][%d] queryCombinedFirewallPoliciesTooManyRequests  %+v", 429, o.Payload)
}
func (o *QueryCombinedFirewallPoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryCombinedFirewallPoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedFirewallPoliciesInternalServerError creates a QueryCombinedFirewallPoliciesInternalServerError with default headers values
func NewQueryCombinedFirewallPoliciesInternalServerError() *QueryCombinedFirewallPoliciesInternalServerError {
	return &QueryCombinedFirewallPoliciesInternalServerError{}
}

/*
	QueryCombinedFirewallPoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryCombinedFirewallPoliciesInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

func (o *QueryCombinedFirewallPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall/v1][%d] queryCombinedFirewallPoliciesInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryCombinedFirewallPoliciesInternalServerError) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *QueryCombinedFirewallPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedFirewallPoliciesDefault creates a QueryCombinedFirewallPoliciesDefault with default headers values
func NewQueryCombinedFirewallPoliciesDefault(code int) *QueryCombinedFirewallPoliciesDefault {
	return &QueryCombinedFirewallPoliciesDefault{
		_statusCode: code,
	}
}

/*
	QueryCombinedFirewallPoliciesDefault describes a response with status code -1, with default header values.

OK
*/
type QueryCombinedFirewallPoliciesDefault struct {
	_statusCode int

	Payload *models.ResponsesFirewallPoliciesV1
}

// Code gets the status code for the query combined firewall policies default response
func (o *QueryCombinedFirewallPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *QueryCombinedFirewallPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall/v1][%d] queryCombinedFirewallPolicies default  %+v", o._statusCode, o.Payload)
}
func (o *QueryCombinedFirewallPoliciesDefault) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *QueryCombinedFirewallPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
