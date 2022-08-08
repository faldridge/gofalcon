// Code generated by go-swagger; DO NOT EDIT.

package real_time_response_admin

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

// RTRListScriptsReader is a Reader for the RTRListScripts structure.
type RTRListScriptsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRListScriptsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRListScriptsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRListScriptsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRListScriptsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRTRListScriptsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRListScriptsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRTRListScriptsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRTRListScriptsOK creates a RTRListScriptsOK with default headers values
func NewRTRListScriptsOK() *RTRListScriptsOK {
	return &RTRListScriptsOK{}
}

/*
	RTRListScriptsOK describes a response with status code 200, with default header values.

OK
*/
type RTRListScriptsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.BinservclientMsaPutFileResponse
}

func (o *RTRListScriptsOK) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] rTRListScriptsOK  %+v", 200, o.Payload)
}
func (o *RTRListScriptsOK) GetPayload() *models.BinservclientMsaPutFileResponse {
	return o.Payload
}

func (o *RTRListScriptsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.BinservclientMsaPutFileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListScriptsBadRequest creates a RTRListScriptsBadRequest with default headers values
func NewRTRListScriptsBadRequest() *RTRListScriptsBadRequest {
	return &RTRListScriptsBadRequest{}
}

/*
	RTRListScriptsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRListScriptsBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

func (o *RTRListScriptsBadRequest) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] rTRListScriptsBadRequest  %+v", 400, o.Payload)
}
func (o *RTRListScriptsBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListScriptsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRListScriptsForbidden creates a RTRListScriptsForbidden with default headers values
func NewRTRListScriptsForbidden() *RTRListScriptsForbidden {
	return &RTRListScriptsForbidden{}
}

/*
	RTRListScriptsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRListScriptsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *RTRListScriptsForbidden) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] rTRListScriptsForbidden  %+v", 403, o.Payload)
}
func (o *RTRListScriptsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRListScriptsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRListScriptsNotFound creates a RTRListScriptsNotFound with default headers values
func NewRTRListScriptsNotFound() *RTRListScriptsNotFound {
	return &RTRListScriptsNotFound{}
}

/*
	RTRListScriptsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RTRListScriptsNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

func (o *RTRListScriptsNotFound) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] rTRListScriptsNotFound  %+v", 404, o.Payload)
}
func (o *RTRListScriptsNotFound) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListScriptsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRListScriptsTooManyRequests creates a RTRListScriptsTooManyRequests with default headers values
func NewRTRListScriptsTooManyRequests() *RTRListScriptsTooManyRequests {
	return &RTRListScriptsTooManyRequests{}
}

/*
	RTRListScriptsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRListScriptsTooManyRequests struct {

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

func (o *RTRListScriptsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] rTRListScriptsTooManyRequests  %+v", 429, o.Payload)
}
func (o *RTRListScriptsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRListScriptsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRListScriptsDefault creates a RTRListScriptsDefault with default headers values
func NewRTRListScriptsDefault(code int) *RTRListScriptsDefault {
	return &RTRListScriptsDefault{
		_statusCode: code,
	}
}

/*
	RTRListScriptsDefault describes a response with status code -1, with default header values.

OK
*/
type RTRListScriptsDefault struct {
	_statusCode int

	Payload *models.BinservclientMsaPutFileResponse
}

// Code gets the status code for the r t r list scripts default response
func (o *RTRListScriptsDefault) Code() int {
	return o._statusCode
}

func (o *RTRListScriptsDefault) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] RTR-ListScripts default  %+v", o._statusCode, o.Payload)
}
func (o *RTRListScriptsDefault) GetPayload() *models.BinservclientMsaPutFileResponse {
	return o.Payload
}

func (o *RTRListScriptsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BinservclientMsaPutFileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
