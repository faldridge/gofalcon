// Code generated by go-swagger; DO NOT EDIT.

package cloud_connect_aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cloud connect aws API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cloud connect aws API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateOrUpdateAWSSettings(params *CreateOrUpdateAWSSettingsParams, opts ...ClientOption) (*CreateOrUpdateAWSSettingsCreated, error)

	DeleteAWSAccounts(params *DeleteAWSAccountsParams, opts ...ClientOption) (*DeleteAWSAccountsOK, error)

	GetAWSAccounts(params *GetAWSAccountsParams, opts ...ClientOption) (*GetAWSAccountsOK, error)

	GetAWSSettings(params *GetAWSSettingsParams, opts ...ClientOption) (*GetAWSSettingsOK, error)

	ProvisionAWSAccounts(params *ProvisionAWSAccountsParams, opts ...ClientOption) (*ProvisionAWSAccountsCreated, error)

	QueryAWSAccounts(params *QueryAWSAccountsParams, opts ...ClientOption) (*QueryAWSAccountsOK, error)

	QueryAWSAccountsForIDs(params *QueryAWSAccountsForIDsParams, opts ...ClientOption) (*QueryAWSAccountsForIDsOK, error)

	UpdateAWSAccounts(params *UpdateAWSAccountsParams, opts ...ClientOption) (*UpdateAWSAccountsOK, error)

	VerifyAWSAccountAccess(params *VerifyAWSAccountAccessParams, opts ...ClientOption) (*VerifyAWSAccountAccessOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateOrUpdateAWSSettings creates or update global settings which are applicable to all provisioned a w s accounts
*/
func (a *Client) CreateOrUpdateAWSSettings(params *CreateOrUpdateAWSSettingsParams, opts ...ClientOption) (*CreateOrUpdateAWSSettingsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrUpdateAWSSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateOrUpdateAWSSettings",
		Method:             "POST",
		PathPattern:        "/cloud-connect-aws/entities/settings/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOrUpdateAWSSettingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateOrUpdateAWSSettingsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateOrUpdateAWSSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteAWSAccounts deletes a set of a w s accounts by specifying their i ds
*/
func (a *Client) DeleteAWSAccounts(params *DeleteAWSAccountsParams, opts ...ClientOption) (*DeleteAWSAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAWSAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAWSAccounts",
		Method:             "DELETE",
		PathPattern:        "/cloud-connect-aws/entities/accounts/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAWSAccountsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAWSAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteAWSAccountsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetAWSAccounts retrieves a set of a w s accounts by specifying their i ds
*/
func (a *Client) GetAWSAccounts(params *GetAWSAccountsParams, opts ...ClientOption) (*GetAWSAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAWSAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAWSAccounts",
		Method:             "GET",
		PathPattern:        "/cloud-connect-aws/entities/accounts/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAWSAccountsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAWSAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAWSAccountsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetAWSSettings retrieves a set of global settings which are applicable to all provisioned a w s accounts
*/
func (a *Client) GetAWSSettings(params *GetAWSSettingsParams, opts ...ClientOption) (*GetAWSSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAWSSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAWSSettings",
		Method:             "GET",
		PathPattern:        "/cloud-connect-aws/combined/settings/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAWSSettingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAWSSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAWSSettingsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProvisionAWSAccounts provisions a w s accounts by specifying details about the accounts to provision
*/
func (a *Client) ProvisionAWSAccounts(params *ProvisionAWSAccountsParams, opts ...ClientOption) (*ProvisionAWSAccountsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProvisionAWSAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ProvisionAWSAccounts",
		Method:             "POST",
		PathPattern:        "/cloud-connect-aws/entities/accounts/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProvisionAWSAccountsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProvisionAWSAccountsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProvisionAWSAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryAWSAccounts searches for provisioned a w s accounts by providing an f q l filter and paging details returns a set of a w s accounts which match the filter criteria
*/
func (a *Client) QueryAWSAccounts(params *QueryAWSAccountsParams, opts ...ClientOption) (*QueryAWSAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryAWSAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryAWSAccounts",
		Method:             "GET",
		PathPattern:        "/cloud-connect-aws/combined/accounts/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryAWSAccountsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryAWSAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryAWSAccountsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryAWSAccountsForIDs searches for provisioned a w s accounts by providing an f q l filter and paging details returns a set of a w s account i ds which match the filter criteria
*/
func (a *Client) QueryAWSAccountsForIDs(params *QueryAWSAccountsForIDsParams, opts ...ClientOption) (*QueryAWSAccountsForIDsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryAWSAccountsForIDsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryAWSAccountsForIDs",
		Method:             "GET",
		PathPattern:        "/cloud-connect-aws/queries/accounts/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryAWSAccountsForIDsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryAWSAccountsForIDsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryAWSAccountsForIDsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateAWSAccounts updates a w s accounts by specifying the ID of the account and details to update
*/
func (a *Client) UpdateAWSAccounts(params *UpdateAWSAccountsParams, opts ...ClientOption) (*UpdateAWSAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAWSAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateAWSAccounts",
		Method:             "PATCH",
		PathPattern:        "/cloud-connect-aws/entities/accounts/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAWSAccountsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateAWSAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateAWSAccountsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
VerifyAWSAccountAccess performs an access verification check on the specified a w s account i ds
*/
func (a *Client) VerifyAWSAccountAccess(params *VerifyAWSAccountAccessParams, opts ...ClientOption) (*VerifyAWSAccountAccessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVerifyAWSAccountAccessParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "VerifyAWSAccountAccess",
		Method:             "POST",
		PathPattern:        "/cloud-connect-aws/entities/verify-account-access/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VerifyAWSAccountAccessReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VerifyAWSAccountAccessOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VerifyAWSAccountAccessDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
