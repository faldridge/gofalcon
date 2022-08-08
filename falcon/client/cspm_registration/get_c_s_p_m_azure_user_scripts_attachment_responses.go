// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// GetCSPMAzureUserScriptsAttachmentReader is a Reader for the GetCSPMAzureUserScriptsAttachment structure.
type GetCSPMAzureUserScriptsAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCSPMAzureUserScriptsAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCSPMAzureUserScriptsAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCSPMAzureUserScriptsAttachmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCSPMAzureUserScriptsAttachmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCSPMAzureUserScriptsAttachmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCSPMAzureUserScriptsAttachmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCSPMAzureUserScriptsAttachmentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCSPMAzureUserScriptsAttachmentOK creates a GetCSPMAzureUserScriptsAttachmentOK with default headers values
func NewGetCSPMAzureUserScriptsAttachmentOK() *GetCSPMAzureUserScriptsAttachmentOK {
	return &GetCSPMAzureUserScriptsAttachmentOK{}
}

/*
	GetCSPMAzureUserScriptsAttachmentOK describes a response with status code 200, with default header values.

OK
*/
type GetCSPMAzureUserScriptsAttachmentOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureProvisionGetUserScriptResponseV1
}

func (o *GetCSPMAzureUserScriptsAttachmentOK) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-azure/entities/user-scripts-download/v1][%d] getCSPMAzureUserScriptsAttachmentOK  %+v", 200, o.Payload)
}
func (o *GetCSPMAzureUserScriptsAttachmentOK) GetPayload() *models.RegistrationAzureProvisionGetUserScriptResponseV1 {
	return o.Payload
}

func (o *GetCSPMAzureUserScriptsAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAzureProvisionGetUserScriptResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMAzureUserScriptsAttachmentBadRequest creates a GetCSPMAzureUserScriptsAttachmentBadRequest with default headers values
func NewGetCSPMAzureUserScriptsAttachmentBadRequest() *GetCSPMAzureUserScriptsAttachmentBadRequest {
	return &GetCSPMAzureUserScriptsAttachmentBadRequest{}
}

/*
	GetCSPMAzureUserScriptsAttachmentBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCSPMAzureUserScriptsAttachmentBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureProvisionGetUserScriptResponseV1
}

func (o *GetCSPMAzureUserScriptsAttachmentBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-azure/entities/user-scripts-download/v1][%d] getCSPMAzureUserScriptsAttachmentBadRequest  %+v", 400, o.Payload)
}
func (o *GetCSPMAzureUserScriptsAttachmentBadRequest) GetPayload() *models.RegistrationAzureProvisionGetUserScriptResponseV1 {
	return o.Payload
}

func (o *GetCSPMAzureUserScriptsAttachmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAzureProvisionGetUserScriptResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMAzureUserScriptsAttachmentForbidden creates a GetCSPMAzureUserScriptsAttachmentForbidden with default headers values
func NewGetCSPMAzureUserScriptsAttachmentForbidden() *GetCSPMAzureUserScriptsAttachmentForbidden {
	return &GetCSPMAzureUserScriptsAttachmentForbidden{}
}

/*
	GetCSPMAzureUserScriptsAttachmentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCSPMAzureUserScriptsAttachmentForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *GetCSPMAzureUserScriptsAttachmentForbidden) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-azure/entities/user-scripts-download/v1][%d] getCSPMAzureUserScriptsAttachmentForbidden  %+v", 403, o.Payload)
}
func (o *GetCSPMAzureUserScriptsAttachmentForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMAzureUserScriptsAttachmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMAzureUserScriptsAttachmentTooManyRequests creates a GetCSPMAzureUserScriptsAttachmentTooManyRequests with default headers values
func NewGetCSPMAzureUserScriptsAttachmentTooManyRequests() *GetCSPMAzureUserScriptsAttachmentTooManyRequests {
	return &GetCSPMAzureUserScriptsAttachmentTooManyRequests{}
}

/*
	GetCSPMAzureUserScriptsAttachmentTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCSPMAzureUserScriptsAttachmentTooManyRequests struct {

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

func (o *GetCSPMAzureUserScriptsAttachmentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-azure/entities/user-scripts-download/v1][%d] getCSPMAzureUserScriptsAttachmentTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetCSPMAzureUserScriptsAttachmentTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMAzureUserScriptsAttachmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMAzureUserScriptsAttachmentInternalServerError creates a GetCSPMAzureUserScriptsAttachmentInternalServerError with default headers values
func NewGetCSPMAzureUserScriptsAttachmentInternalServerError() *GetCSPMAzureUserScriptsAttachmentInternalServerError {
	return &GetCSPMAzureUserScriptsAttachmentInternalServerError{}
}

/*
	GetCSPMAzureUserScriptsAttachmentInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCSPMAzureUserScriptsAttachmentInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureProvisionGetUserScriptResponseV1
}

func (o *GetCSPMAzureUserScriptsAttachmentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-azure/entities/user-scripts-download/v1][%d] getCSPMAzureUserScriptsAttachmentInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCSPMAzureUserScriptsAttachmentInternalServerError) GetPayload() *models.RegistrationAzureProvisionGetUserScriptResponseV1 {
	return o.Payload
}

func (o *GetCSPMAzureUserScriptsAttachmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAzureProvisionGetUserScriptResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMAzureUserScriptsAttachmentDefault creates a GetCSPMAzureUserScriptsAttachmentDefault with default headers values
func NewGetCSPMAzureUserScriptsAttachmentDefault(code int) *GetCSPMAzureUserScriptsAttachmentDefault {
	return &GetCSPMAzureUserScriptsAttachmentDefault{
		_statusCode: code,
	}
}

/*
	GetCSPMAzureUserScriptsAttachmentDefault describes a response with status code -1, with default header values.

OK
*/
type GetCSPMAzureUserScriptsAttachmentDefault struct {
	_statusCode int

	Payload *models.RegistrationAzureProvisionGetUserScriptResponseV1
}

// Code gets the status code for the get c s p m azure user scripts attachment default response
func (o *GetCSPMAzureUserScriptsAttachmentDefault) Code() int {
	return o._statusCode
}

func (o *GetCSPMAzureUserScriptsAttachmentDefault) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-azure/entities/user-scripts-download/v1][%d] GetCSPMAzureUserScriptsAttachment default  %+v", o._statusCode, o.Payload)
}
func (o *GetCSPMAzureUserScriptsAttachmentDefault) GetPayload() *models.RegistrationAzureProvisionGetUserScriptResponseV1 {
	return o.Payload
}

func (o *GetCSPMAzureUserScriptsAttachmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistrationAzureProvisionGetUserScriptResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
