// Code generated by go-swagger; DO NOT EDIT.

package sensor_visibility_exclusions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new sensor visibility exclusions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sensor visibility exclusions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSVExclusionsV1(params *CreateSVExclusionsV1Params) (*CreateSVExclusionsV1OK, error)

	DeleteSensorVisibilityExclusionsV1(params *DeleteSensorVisibilityExclusionsV1Params) (*DeleteSensorVisibilityExclusionsV1OK, error)

	GetSensorVisibilityExclusionsV1(params *GetSensorVisibilityExclusionsV1Params) (*GetSensorVisibilityExclusionsV1OK, error)

	QuerySensorVisibilityExclusionsV1(params *QuerySensorVisibilityExclusionsV1Params) (*QuerySensorVisibilityExclusionsV1OK, error)

	UpdateSensorVisibilityExclusionsV1(params *UpdateSensorVisibilityExclusionsV1Params) (*UpdateSensorVisibilityExclusionsV1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateSVExclusionsV1 creates the sensor visibility exclusions
*/
func (a *Client) CreateSVExclusionsV1(params *CreateSVExclusionsV1Params) (*CreateSVExclusionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSVExclusionsV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSVExclusionsV1",
		Method:             "POST",
		PathPattern:        "/policy/entities/sv-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSVExclusionsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateSVExclusionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateSVExclusionsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteSensorVisibilityExclusionsV1 deletes the sensor visibility exclusions by id
*/
func (a *Client) DeleteSensorVisibilityExclusionsV1(params *DeleteSensorVisibilityExclusionsV1Params) (*DeleteSensorVisibilityExclusionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSensorVisibilityExclusionsV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSensorVisibilityExclusionsV1",
		Method:             "DELETE",
		PathPattern:        "/policy/entities/sv-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSensorVisibilityExclusionsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteSensorVisibilityExclusionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSensorVisibilityExclusionsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSensorVisibilityExclusionsV1 gets a set of sensor visibility exclusions by specifying their i ds
*/
func (a *Client) GetSensorVisibilityExclusionsV1(params *GetSensorVisibilityExclusionsV1Params) (*GetSensorVisibilityExclusionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSensorVisibilityExclusionsV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSensorVisibilityExclusionsV1",
		Method:             "GET",
		PathPattern:        "/policy/entities/sv-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSensorVisibilityExclusionsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSensorVisibilityExclusionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSensorVisibilityExclusionsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  QuerySensorVisibilityExclusionsV1 searches for sensor visibility exclusions
*/
func (a *Client) QuerySensorVisibilityExclusionsV1(params *QuerySensorVisibilityExclusionsV1Params) (*QuerySensorVisibilityExclusionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuerySensorVisibilityExclusionsV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "querySensorVisibilityExclusionsV1",
		Method:             "GET",
		PathPattern:        "/policy/queries/sv-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuerySensorVisibilityExclusionsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QuerySensorVisibilityExclusionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QuerySensorVisibilityExclusionsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateSensorVisibilityExclusionsV1 updates the sensor visibility exclusions
*/
func (a *Client) UpdateSensorVisibilityExclusionsV1(params *UpdateSensorVisibilityExclusionsV1Params) (*UpdateSensorVisibilityExclusionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSensorVisibilityExclusionsV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSensorVisibilityExclusionsV1",
		Method:             "PATCH",
		PathPattern:        "/policy/entities/sv-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSensorVisibilityExclusionsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateSensorVisibilityExclusionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateSensorVisibilityExclusionsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
