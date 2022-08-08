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

// QueryFirewallPoliciesReader is a Reader for the QueryFirewallPolicies structure.
type QueryFirewallPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryFirewallPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryFirewallPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryFirewallPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryFirewallPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryFirewallPoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryFirewallPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryFirewallPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryFirewallPoliciesOK creates a QueryFirewallPoliciesOK with default headers values
func NewQueryFirewallPoliciesOK() *QueryFirewallPoliciesOK {
	return &QueryFirewallPoliciesOK{}
}

/*
	QueryFirewallPoliciesOK describes a response with status code 200, with default header values.

OK
*/
type QueryFirewallPoliciesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryFirewallPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPoliciesOK  %+v", 200, o.Payload)
}
func (o *QueryFirewallPoliciesOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryFirewallPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryFirewallPoliciesBadRequest creates a QueryFirewallPoliciesBadRequest with default headers values
func NewQueryFirewallPoliciesBadRequest() *QueryFirewallPoliciesBadRequest {
	return &QueryFirewallPoliciesBadRequest{}
}

/*
	QueryFirewallPoliciesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryFirewallPoliciesBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryFirewallPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPoliciesBadRequest  %+v", 400, o.Payload)
}
func (o *QueryFirewallPoliciesBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryFirewallPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryFirewallPoliciesForbidden creates a QueryFirewallPoliciesForbidden with default headers values
func NewQueryFirewallPoliciesForbidden() *QueryFirewallPoliciesForbidden {
	return &QueryFirewallPoliciesForbidden{}
}

/*
	QueryFirewallPoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryFirewallPoliciesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *QueryFirewallPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPoliciesForbidden  %+v", 403, o.Payload)
}
func (o *QueryFirewallPoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryFirewallPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryFirewallPoliciesTooManyRequests creates a QueryFirewallPoliciesTooManyRequests with default headers values
func NewQueryFirewallPoliciesTooManyRequests() *QueryFirewallPoliciesTooManyRequests {
	return &QueryFirewallPoliciesTooManyRequests{}
}

/*
	QueryFirewallPoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryFirewallPoliciesTooManyRequests struct {

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

func (o *QueryFirewallPoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPoliciesTooManyRequests  %+v", 429, o.Payload)
}
func (o *QueryFirewallPoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryFirewallPoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryFirewallPoliciesInternalServerError creates a QueryFirewallPoliciesInternalServerError with default headers values
func NewQueryFirewallPoliciesInternalServerError() *QueryFirewallPoliciesInternalServerError {
	return &QueryFirewallPoliciesInternalServerError{}
}

/*
	QueryFirewallPoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryFirewallPoliciesInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryFirewallPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPoliciesInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryFirewallPoliciesInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryFirewallPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryFirewallPoliciesDefault creates a QueryFirewallPoliciesDefault with default headers values
func NewQueryFirewallPoliciesDefault(code int) *QueryFirewallPoliciesDefault {
	return &QueryFirewallPoliciesDefault{
		_statusCode: code,
	}
}

/*
	QueryFirewallPoliciesDefault describes a response with status code -1, with default header values.

OK
*/
type QueryFirewallPoliciesDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the query firewall policies default response
func (o *QueryFirewallPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *QueryFirewallPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPolicies default  %+v", o._statusCode, o.Payload)
}
func (o *QueryFirewallPoliciesDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryFirewallPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
