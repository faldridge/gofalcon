// Code generated by go-swagger; DO NOT EDIT.

package oauth2

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

// Oauth2AccessTokenReader is a Reader for the Oauth2AccessToken structure.
type Oauth2AccessTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Oauth2AccessTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewOauth2AccessTokenCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOauth2AccessTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOauth2AccessTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOauth2AccessTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOauth2AccessTokenCreated creates a Oauth2AccessTokenCreated with default headers values
func NewOauth2AccessTokenCreated() *Oauth2AccessTokenCreated {
	return &Oauth2AccessTokenCreated{}
}

/*
	Oauth2AccessTokenCreated describes a response with status code 201, with default header values.

Successfully issued token
*/
type Oauth2AccessTokenCreated struct {
	XCSRegion string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAccessTokenResponseV1
}

func (o *Oauth2AccessTokenCreated) Error() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauth2AccessTokenCreated  %+v", 201, o.Payload)
}
func (o *Oauth2AccessTokenCreated) GetPayload() *models.DomainAccessTokenResponseV1 {
	return o.Payload
}

func (o *Oauth2AccessTokenCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-Region
	hdrXCSRegion := response.GetHeader("X-CS-Region")

	if hdrXCSRegion != "" {
		o.XCSRegion = hdrXCSRegion
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

	o.Payload = new(models.DomainAccessTokenResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOauth2AccessTokenBadRequest creates a Oauth2AccessTokenBadRequest with default headers values
func NewOauth2AccessTokenBadRequest() *Oauth2AccessTokenBadRequest {
	return &Oauth2AccessTokenBadRequest{}
}

/*
	Oauth2AccessTokenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type Oauth2AccessTokenBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *Oauth2AccessTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauth2AccessTokenBadRequest  %+v", 400, o.Payload)
}
func (o *Oauth2AccessTokenBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *Oauth2AccessTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewOauth2AccessTokenForbidden creates a Oauth2AccessTokenForbidden with default headers values
func NewOauth2AccessTokenForbidden() *Oauth2AccessTokenForbidden {
	return &Oauth2AccessTokenForbidden{}
}

/*
	Oauth2AccessTokenForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type Oauth2AccessTokenForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *Oauth2AccessTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauth2AccessTokenForbidden  %+v", 403, o.Payload)
}
func (o *Oauth2AccessTokenForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *Oauth2AccessTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewOauth2AccessTokenInternalServerError creates a Oauth2AccessTokenInternalServerError with default headers values
func NewOauth2AccessTokenInternalServerError() *Oauth2AccessTokenInternalServerError {
	return &Oauth2AccessTokenInternalServerError{}
}

/*
	Oauth2AccessTokenInternalServerError describes a response with status code 500, with default header values.

Failed to issue token
*/
type Oauth2AccessTokenInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *Oauth2AccessTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauth2AccessTokenInternalServerError  %+v", 500, o.Payload)
}
func (o *Oauth2AccessTokenInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *Oauth2AccessTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
