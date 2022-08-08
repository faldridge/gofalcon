// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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

// DeleteAWSAccountsMixin0Reader is a Reader for the DeleteAWSAccountsMixin0 structure.
type DeleteAWSAccountsMixin0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAWSAccountsMixin0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAWSAccountsMixin0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewDeleteAWSAccountsMixin0MultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAWSAccountsMixin0BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAWSAccountsMixin0Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteAWSAccountsMixin0TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAWSAccountsMixin0InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteAWSAccountsMixin0Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAWSAccountsMixin0OK creates a DeleteAWSAccountsMixin0OK with default headers values
func NewDeleteAWSAccountsMixin0OK() *DeleteAWSAccountsMixin0OK {
	return &DeleteAWSAccountsMixin0OK{}
}

/*
	DeleteAWSAccountsMixin0OK describes a response with status code 200, with default header values.

OK
*/
type DeleteAWSAccountsMixin0OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaMetaInfo
}

func (o *DeleteAWSAccountsMixin0OK) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/aws/v1][%d] deleteAWSAccountsMixin0OK  %+v", 200, o.Payload)
}
func (o *DeleteAWSAccountsMixin0OK) GetPayload() *models.MsaMetaInfo {
	return o.Payload
}

func (o *DeleteAWSAccountsMixin0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.MsaMetaInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAWSAccountsMixin0MultiStatus creates a DeleteAWSAccountsMixin0MultiStatus with default headers values
func NewDeleteAWSAccountsMixin0MultiStatus() *DeleteAWSAccountsMixin0MultiStatus {
	return &DeleteAWSAccountsMixin0MultiStatus{}
}

/*
	DeleteAWSAccountsMixin0MultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type DeleteAWSAccountsMixin0MultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaMetaInfo
}

func (o *DeleteAWSAccountsMixin0MultiStatus) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/aws/v1][%d] deleteAWSAccountsMixin0MultiStatus  %+v", 207, o.Payload)
}
func (o *DeleteAWSAccountsMixin0MultiStatus) GetPayload() *models.MsaMetaInfo {
	return o.Payload
}

func (o *DeleteAWSAccountsMixin0MultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.MsaMetaInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAWSAccountsMixin0BadRequest creates a DeleteAWSAccountsMixin0BadRequest with default headers values
func NewDeleteAWSAccountsMixin0BadRequest() *DeleteAWSAccountsMixin0BadRequest {
	return &DeleteAWSAccountsMixin0BadRequest{}
}

/*
	DeleteAWSAccountsMixin0BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteAWSAccountsMixin0BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaMetaInfo
}

func (o *DeleteAWSAccountsMixin0BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/aws/v1][%d] deleteAWSAccountsMixin0BadRequest  %+v", 400, o.Payload)
}
func (o *DeleteAWSAccountsMixin0BadRequest) GetPayload() *models.MsaMetaInfo {
	return o.Payload
}

func (o *DeleteAWSAccountsMixin0BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.MsaMetaInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAWSAccountsMixin0Forbidden creates a DeleteAWSAccountsMixin0Forbidden with default headers values
func NewDeleteAWSAccountsMixin0Forbidden() *DeleteAWSAccountsMixin0Forbidden {
	return &DeleteAWSAccountsMixin0Forbidden{}
}

/*
	DeleteAWSAccountsMixin0Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteAWSAccountsMixin0Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *DeleteAWSAccountsMixin0Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/aws/v1][%d] deleteAWSAccountsMixin0Forbidden  %+v", 403, o.Payload)
}
func (o *DeleteAWSAccountsMixin0Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteAWSAccountsMixin0Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewDeleteAWSAccountsMixin0TooManyRequests creates a DeleteAWSAccountsMixin0TooManyRequests with default headers values
func NewDeleteAWSAccountsMixin0TooManyRequests() *DeleteAWSAccountsMixin0TooManyRequests {
	return &DeleteAWSAccountsMixin0TooManyRequests{}
}

/*
	DeleteAWSAccountsMixin0TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteAWSAccountsMixin0TooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

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

func (o *DeleteAWSAccountsMixin0TooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/aws/v1][%d] deleteAWSAccountsMixin0TooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteAWSAccountsMixin0TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteAWSAccountsMixin0TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewDeleteAWSAccountsMixin0InternalServerError creates a DeleteAWSAccountsMixin0InternalServerError with default headers values
func NewDeleteAWSAccountsMixin0InternalServerError() *DeleteAWSAccountsMixin0InternalServerError {
	return &DeleteAWSAccountsMixin0InternalServerError{}
}

/*
	DeleteAWSAccountsMixin0InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteAWSAccountsMixin0InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaMetaInfo
}

func (o *DeleteAWSAccountsMixin0InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/aws/v1][%d] deleteAWSAccountsMixin0InternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteAWSAccountsMixin0InternalServerError) GetPayload() *models.MsaMetaInfo {
	return o.Payload
}

func (o *DeleteAWSAccountsMixin0InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.MsaMetaInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAWSAccountsMixin0Default creates a DeleteAWSAccountsMixin0Default with default headers values
func NewDeleteAWSAccountsMixin0Default(code int) *DeleteAWSAccountsMixin0Default {
	return &DeleteAWSAccountsMixin0Default{
		_statusCode: code,
	}
}

/*
	DeleteAWSAccountsMixin0Default describes a response with status code -1, with default header values.

OK
*/
type DeleteAWSAccountsMixin0Default struct {
	_statusCode int

	Payload *models.MsaMetaInfo
}

// Code gets the status code for the delete a w s accounts mixin0 default response
func (o *DeleteAWSAccountsMixin0Default) Code() int {
	return o._statusCode
}

func (o *DeleteAWSAccountsMixin0Default) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-protection/entities/accounts/aws/v1][%d] DeleteAWSAccountsMixin0 default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteAWSAccountsMixin0Default) GetPayload() *models.MsaMetaInfo {
	return o.Payload
}

func (o *DeleteAWSAccountsMixin0Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaMetaInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
