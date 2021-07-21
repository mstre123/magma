// Code generated by go-swagger; DO NOT EDIT.

package federated_lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new federated lte networks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for federated lte networks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteFegLTENetworkID(params *DeleteFegLTENetworkIDParams, opts ...ClientOption) (*DeleteFegLTENetworkIDNoContent, error)

	DeleteFegLTENetworkIDFederation(params *DeleteFegLTENetworkIDFederationParams, opts ...ClientOption) (*DeleteFegLTENetworkIDFederationNoContent, error)

	DeleteFegLTENetworkIDSubscriberConfigBaseNamesBaseName(params *DeleteFegLTENetworkIDSubscriberConfigBaseNamesBaseNameParams, opts ...ClientOption) (*DeleteFegLTENetworkIDSubscriberConfigBaseNamesBaseNameNoContent, error)

	DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleID(params *DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDParams, opts ...ClientOption) (*DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDNoContent, error)

	GetFegLTE(params *GetFegLTEParams, opts ...ClientOption) (*GetFegLTEOK, error)

	GetFegLTENetworkID(params *GetFegLTENetworkIDParams, opts ...ClientOption) (*GetFegLTENetworkIDOK, error)

	GetFegLTENetworkIDFederation(params *GetFegLTENetworkIDFederationParams, opts ...ClientOption) (*GetFegLTENetworkIDFederationOK, error)

	GetFegLTENetworkIDSubscriberConfig(params *GetFegLTENetworkIDSubscriberConfigParams, opts ...ClientOption) (*GetFegLTENetworkIDSubscriberConfigOK, error)

	GetFegLTENetworkIDSubscriberConfigBaseNames(params *GetFegLTENetworkIDSubscriberConfigBaseNamesParams, opts ...ClientOption) (*GetFegLTENetworkIDSubscriberConfigBaseNamesOK, error)

	GetFegLTENetworkIDSubscriberConfigRuleNames(params *GetFegLTENetworkIDSubscriberConfigRuleNamesParams, opts ...ClientOption) (*GetFegLTENetworkIDSubscriberConfigRuleNamesOK, error)

	PostFegLTE(params *PostFegLTEParams, opts ...ClientOption) (*PostFegLTECreated, error)

	PostFegLTENetworkIDSubscriberConfigBaseNamesBaseName(params *PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameParams, opts ...ClientOption) (*PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated, error)

	PostFegLTENetworkIDSubscriberConfigRuleNamesRuleID(params *PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDParams, opts ...ClientOption) (*PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDCreated, error)

	PutFegLTENetworkID(params *PutFegLTENetworkIDParams, opts ...ClientOption) (*PutFegLTENetworkIDNoContent, error)

	PutFegLTENetworkIDFederation(params *PutFegLTENetworkIDFederationParams, opts ...ClientOption) (*PutFegLTENetworkIDFederationOK, error)

	PutFegLTENetworkIDSubscriberConfig(params *PutFegLTENetworkIDSubscriberConfigParams, opts ...ClientOption) (*PutFegLTENetworkIDSubscriberConfigNoContent, error)

	PutFegLTENetworkIDSubscriberConfigBaseNames(params *PutFegLTENetworkIDSubscriberConfigBaseNamesParams, opts ...ClientOption) (*PutFegLTENetworkIDSubscriberConfigBaseNamesNoContent, error)

	PutFegLTENetworkIDSubscriberConfigRuleNames(params *PutFegLTENetworkIDSubscriberConfigRuleNamesParams, opts ...ClientOption) (*PutFegLTENetworkIDSubscriberConfigRuleNamesNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteFegLTENetworkID deletes a federated LTE network
*/
func (a *Client) DeleteFegLTENetworkID(params *DeleteFegLTENetworkIDParams, opts ...ClientOption) (*DeleteFegLTENetworkIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFegLTENetworkIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteFegLTENetworkID",
		Method:             "DELETE",
		PathPattern:        "/feg_lte/{network_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteFegLTENetworkIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteFegLTENetworkIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteFegLTENetworkIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteFegLTENetworkIDFederation deletes network fe g configs
*/
func (a *Client) DeleteFegLTENetworkIDFederation(params *DeleteFegLTENetworkIDFederationParams, opts ...ClientOption) (*DeleteFegLTENetworkIDFederationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFegLTENetworkIDFederationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteFegLTENetworkIDFederation",
		Method:             "DELETE",
		PathPattern:        "/feg_lte/{network_id}/federation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteFegLTENetworkIDFederationReader{formats: a.formats},
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
	success, ok := result.(*DeleteFegLTENetworkIDFederationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteFegLTENetworkIDFederationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteFegLTENetworkIDSubscriberConfigBaseNamesBaseName adds a network wide base name
*/
func (a *Client) DeleteFegLTENetworkIDSubscriberConfigBaseNamesBaseName(params *DeleteFegLTENetworkIDSubscriberConfigBaseNamesBaseNameParams, opts ...ClientOption) (*DeleteFegLTENetworkIDSubscriberConfigBaseNamesBaseNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFegLTENetworkIDSubscriberConfigBaseNamesBaseNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteFegLTENetworkIDSubscriberConfigBaseNamesBaseName",
		Method:             "DELETE",
		PathPattern:        "/feg_lte/{network_id}/subscriber_config/base_names/{base_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteFegLTENetworkIDSubscriberConfigBaseNamesBaseNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteFegLTENetworkIDSubscriberConfigBaseNamesBaseNameNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteFegLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleID adds a network wide rule name
*/
func (a *Client) DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleID(params *DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDParams, opts ...ClientOption) (*DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleID",
		Method:             "DELETE",
		PathPattern:        "/feg_lte/{network_id}/subscriber_config/rule_names/{rule_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetFegLTE lists all federated LTE network i ds
*/
func (a *Client) GetFegLTE(params *GetFegLTEParams, opts ...ClientOption) (*GetFegLTEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFegLTEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetFegLTE",
		Method:             "GET",
		PathPattern:        "/feg_lte",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFegLTEReader{formats: a.formats},
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
	success, ok := result.(*GetFegLTEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFegLTEDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetFegLTENetworkID describes a federated LTE network
*/
func (a *Client) GetFegLTENetworkID(params *GetFegLTENetworkIDParams, opts ...ClientOption) (*GetFegLTENetworkIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFegLTENetworkIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetFegLTENetworkID",
		Method:             "GET",
		PathPattern:        "/feg_lte/{network_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFegLTENetworkIDReader{formats: a.formats},
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
	success, ok := result.(*GetFegLTENetworkIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFegLTENetworkIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetFegLTENetworkIDFederation retrieves network fe g configs
*/
func (a *Client) GetFegLTENetworkIDFederation(params *GetFegLTENetworkIDFederationParams, opts ...ClientOption) (*GetFegLTENetworkIDFederationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFegLTENetworkIDFederationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetFegLTENetworkIDFederation",
		Method:             "GET",
		PathPattern:        "/feg_lte/{network_id}/federation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFegLTENetworkIDFederationReader{formats: a.formats},
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
	success, ok := result.(*GetFegLTENetworkIDFederationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFegLTENetworkIDFederationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetFegLTENetworkIDSubscriberConfig gets a network wide subscriber config
*/
func (a *Client) GetFegLTENetworkIDSubscriberConfig(params *GetFegLTENetworkIDSubscriberConfigParams, opts ...ClientOption) (*GetFegLTENetworkIDSubscriberConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFegLTENetworkIDSubscriberConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetFegLTENetworkIDSubscriberConfig",
		Method:             "GET",
		PathPattern:        "/feg_lte/{network_id}/subscriber_config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFegLTENetworkIDSubscriberConfigReader{formats: a.formats},
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
	success, ok := result.(*GetFegLTENetworkIDSubscriberConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFegLTENetworkIDSubscriberConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetFegLTENetworkIDSubscriberConfigBaseNames gets network wide base names
*/
func (a *Client) GetFegLTENetworkIDSubscriberConfigBaseNames(params *GetFegLTENetworkIDSubscriberConfigBaseNamesParams, opts ...ClientOption) (*GetFegLTENetworkIDSubscriberConfigBaseNamesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFegLTENetworkIDSubscriberConfigBaseNamesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetFegLTENetworkIDSubscriberConfigBaseNames",
		Method:             "GET",
		PathPattern:        "/feg_lte/{network_id}/subscriber_config/base_names",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFegLTENetworkIDSubscriberConfigBaseNamesReader{formats: a.formats},
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
	success, ok := result.(*GetFegLTENetworkIDSubscriberConfigBaseNamesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFegLTENetworkIDSubscriberConfigBaseNamesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetFegLTENetworkIDSubscriberConfigRuleNames gets network wide rule names
*/
func (a *Client) GetFegLTENetworkIDSubscriberConfigRuleNames(params *GetFegLTENetworkIDSubscriberConfigRuleNamesParams, opts ...ClientOption) (*GetFegLTENetworkIDSubscriberConfigRuleNamesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFegLTENetworkIDSubscriberConfigRuleNamesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetFegLTENetworkIDSubscriberConfigRuleNames",
		Method:             "GET",
		PathPattern:        "/feg_lte/{network_id}/subscriber_config/rule_names",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFegLTENetworkIDSubscriberConfigRuleNamesReader{formats: a.formats},
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
	success, ok := result.(*GetFegLTENetworkIDSubscriberConfigRuleNamesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFegLTENetworkIDSubscriberConfigRuleNamesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostFegLTE creates a new federated LTE network
*/
func (a *Client) PostFegLTE(params *PostFegLTEParams, opts ...ClientOption) (*PostFegLTECreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFegLTEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostFegLTE",
		Method:             "POST",
		PathPattern:        "/feg_lte",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostFegLTEReader{formats: a.formats},
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
	success, ok := result.(*PostFegLTECreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostFegLTEDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostFegLTENetworkIDSubscriberConfigBaseNamesBaseName adds a network wide base name
*/
func (a *Client) PostFegLTENetworkIDSubscriberConfigBaseNamesBaseName(params *PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameParams, opts ...ClientOption) (*PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostFegLTENetworkIDSubscriberConfigBaseNamesBaseName",
		Method:             "POST",
		PathPattern:        "/feg_lte/{network_id}/subscriber_config/base_names/{base_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameReader{formats: a.formats},
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
	success, ok := result.(*PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostFegLTENetworkIDSubscriberConfigRuleNamesRuleID adds a network wide rule name
*/
func (a *Client) PostFegLTENetworkIDSubscriberConfigRuleNamesRuleID(params *PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDParams, opts ...ClientOption) (*PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostFegLTENetworkIDSubscriberConfigRuleNamesRuleID",
		Method:             "POST",
		PathPattern:        "/feg_lte/{network_id}/subscriber_config/rule_names/{rule_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDReader{formats: a.formats},
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
	success, ok := result.(*PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutFegLTENetworkID updates an entire federated LTE network
*/
func (a *Client) PutFegLTENetworkID(params *PutFegLTENetworkIDParams, opts ...ClientOption) (*PutFegLTENetworkIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutFegLTENetworkIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutFegLTENetworkID",
		Method:             "PUT",
		PathPattern:        "/feg_lte/{network_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutFegLTENetworkIDReader{formats: a.formats},
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
	success, ok := result.(*PutFegLTENetworkIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutFegLTENetworkIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutFegLTENetworkIDFederation creates or modify network fe g configs
*/
func (a *Client) PutFegLTENetworkIDFederation(params *PutFegLTENetworkIDFederationParams, opts ...ClientOption) (*PutFegLTENetworkIDFederationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutFegLTENetworkIDFederationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutFegLTENetworkIDFederation",
		Method:             "PUT",
		PathPattern:        "/feg_lte/{network_id}/federation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutFegLTENetworkIDFederationReader{formats: a.formats},
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
	success, ok := result.(*PutFegLTENetworkIDFederationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutFegLTENetworkIDFederationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutFegLTENetworkIDSubscriberConfig updates a network wide subscriber config
*/
func (a *Client) PutFegLTENetworkIDSubscriberConfig(params *PutFegLTENetworkIDSubscriberConfigParams, opts ...ClientOption) (*PutFegLTENetworkIDSubscriberConfigNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutFegLTENetworkIDSubscriberConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutFegLTENetworkIDSubscriberConfig",
		Method:             "PUT",
		PathPattern:        "/feg_lte/{network_id}/subscriber_config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutFegLTENetworkIDSubscriberConfigReader{formats: a.formats},
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
	success, ok := result.(*PutFegLTENetworkIDSubscriberConfigNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutFegLTENetworkIDSubscriberConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutFegLTENetworkIDSubscriberConfigBaseNames updates network wide base names
*/
func (a *Client) PutFegLTENetworkIDSubscriberConfigBaseNames(params *PutFegLTENetworkIDSubscriberConfigBaseNamesParams, opts ...ClientOption) (*PutFegLTENetworkIDSubscriberConfigBaseNamesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutFegLTENetworkIDSubscriberConfigBaseNamesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutFegLTENetworkIDSubscriberConfigBaseNames",
		Method:             "PUT",
		PathPattern:        "/feg_lte/{network_id}/subscriber_config/base_names",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutFegLTENetworkIDSubscriberConfigBaseNamesReader{formats: a.formats},
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
	success, ok := result.(*PutFegLTENetworkIDSubscriberConfigBaseNamesNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutFegLTENetworkIDSubscriberConfigBaseNamesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutFegLTENetworkIDSubscriberConfigRuleNames updates network wide rule names
*/
func (a *Client) PutFegLTENetworkIDSubscriberConfigRuleNames(params *PutFegLTENetworkIDSubscriberConfigRuleNamesParams, opts ...ClientOption) (*PutFegLTENetworkIDSubscriberConfigRuleNamesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutFegLTENetworkIDSubscriberConfigRuleNamesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutFegLTENetworkIDSubscriberConfigRuleNames",
		Method:             "PUT",
		PathPattern:        "/feg_lte/{network_id}/subscriber_config/rule_names",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutFegLTENetworkIDSubscriberConfigRuleNamesReader{formats: a.formats},
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
	success, ok := result.(*PutFegLTENetworkIDSubscriberConfigRuleNamesNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutFegLTENetworkIDSubscriberConfigRuleNamesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
