// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

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

// RTRListQueuedSessionsReader is a Reader for the RTRListQueuedSessions structure.
type RTRListQueuedSessionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRListQueuedSessionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRListQueuedSessionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRListQueuedSessionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRTRListQueuedSessionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRListQueuedSessionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRTRListQueuedSessionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRListQueuedSessionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRTRListQueuedSessionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRTRListQueuedSessionsOK creates a RTRListQueuedSessionsOK with default headers values
func NewRTRListQueuedSessionsOK() *RTRListQueuedSessionsOK {
	return &RTRListQueuedSessionsOK{}
}

/*
	RTRListQueuedSessionsOK describes a response with status code 200, with default header values.

success
*/
type RTRListQueuedSessionsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainQueuedSessionResponseWrapper
}

func (o *RTRListQueuedSessionsOK) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsOK  %+v", 200, o.Payload)
}
func (o *RTRListQueuedSessionsOK) GetPayload() *models.DomainQueuedSessionResponseWrapper {
	return o.Payload
}

func (o *RTRListQueuedSessionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainQueuedSessionResponseWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListQueuedSessionsBadRequest creates a RTRListQueuedSessionsBadRequest with default headers values
func NewRTRListQueuedSessionsBadRequest() *RTRListQueuedSessionsBadRequest {
	return &RTRListQueuedSessionsBadRequest{}
}

/*
	RTRListQueuedSessionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRListQueuedSessionsBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

func (o *RTRListQueuedSessionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsBadRequest  %+v", 400, o.Payload)
}
func (o *RTRListQueuedSessionsBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListQueuedSessionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListQueuedSessionsUnauthorized creates a RTRListQueuedSessionsUnauthorized with default headers values
func NewRTRListQueuedSessionsUnauthorized() *RTRListQueuedSessionsUnauthorized {
	return &RTRListQueuedSessionsUnauthorized{}
}

/*
	RTRListQueuedSessionsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RTRListQueuedSessionsUnauthorized struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

func (o *RTRListQueuedSessionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsUnauthorized  %+v", 401, o.Payload)
}
func (o *RTRListQueuedSessionsUnauthorized) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListQueuedSessionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListQueuedSessionsForbidden creates a RTRListQueuedSessionsForbidden with default headers values
func NewRTRListQueuedSessionsForbidden() *RTRListQueuedSessionsForbidden {
	return &RTRListQueuedSessionsForbidden{}
}

/*
	RTRListQueuedSessionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRListQueuedSessionsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *RTRListQueuedSessionsForbidden) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsForbidden  %+v", 403, o.Payload)
}
func (o *RTRListQueuedSessionsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRListQueuedSessionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListQueuedSessionsNotFound creates a RTRListQueuedSessionsNotFound with default headers values
func NewRTRListQueuedSessionsNotFound() *RTRListQueuedSessionsNotFound {
	return &RTRListQueuedSessionsNotFound{}
}

/*
	RTRListQueuedSessionsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RTRListQueuedSessionsNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

func (o *RTRListQueuedSessionsNotFound) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsNotFound  %+v", 404, o.Payload)
}
func (o *RTRListQueuedSessionsNotFound) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListQueuedSessionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListQueuedSessionsTooManyRequests creates a RTRListQueuedSessionsTooManyRequests with default headers values
func NewRTRListQueuedSessionsTooManyRequests() *RTRListQueuedSessionsTooManyRequests {
	return &RTRListQueuedSessionsTooManyRequests{}
}

/*
	RTRListQueuedSessionsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRListQueuedSessionsTooManyRequests struct {

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

func (o *RTRListQueuedSessionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsTooManyRequests  %+v", 429, o.Payload)
}
func (o *RTRListQueuedSessionsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRListQueuedSessionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRListQueuedSessionsDefault creates a RTRListQueuedSessionsDefault with default headers values
func NewRTRListQueuedSessionsDefault(code int) *RTRListQueuedSessionsDefault {
	return &RTRListQueuedSessionsDefault{
		_statusCode: code,
	}
}

/*
	RTRListQueuedSessionsDefault describes a response with status code -1, with default header values.

success
*/
type RTRListQueuedSessionsDefault struct {
	_statusCode int

	Payload *models.DomainQueuedSessionResponseWrapper
}

// Code gets the status code for the r t r list queued sessions default response
func (o *RTRListQueuedSessionsDefault) Code() int {
	return o._statusCode
}

func (o *RTRListQueuedSessionsDefault) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] RTR-ListQueuedSessions default  %+v", o._statusCode, o.Payload)
}
func (o *RTRListQueuedSessionsDefault) GetPayload() *models.DomainQueuedSessionResponseWrapper {
	return o.Payload
}

func (o *RTRListQueuedSessionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainQueuedSessionResponseWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
