// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new carrier wifi gateways API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for carrier wifi gateways API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteCwfNetworkIDGatewaysGatewayID(params *DeleteCwfNetworkIDGatewaysGatewayIDParams, opts ...ClientOption) (*DeleteCwfNetworkIDGatewaysGatewayIDNoContent, error)

	GetCwfNetworkIDGateways(params *GetCwfNetworkIDGatewaysParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysOK, error)

	GetCwfNetworkIDGatewaysGatewayID(params *GetCwfNetworkIDGatewaysGatewayIDParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDOK, error)

	GetCwfNetworkIDGatewaysGatewayIDCarrierWifi(params *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDCarrierWifiOK, error)

	GetCwfNetworkIDGatewaysGatewayIDDescription(params *GetCwfNetworkIDGatewaysGatewayIDDescriptionParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDDescriptionOK, error)

	GetCwfNetworkIDGatewaysGatewayIDDevice(params *GetCwfNetworkIDGatewaysGatewayIDDeviceParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDDeviceOK, error)

	GetCwfNetworkIDGatewaysGatewayIDHealthStatus(params *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDHealthStatusOK, error)

	GetCwfNetworkIDGatewaysGatewayIDMagmad(params *GetCwfNetworkIDGatewaysGatewayIDMagmadParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDMagmadOK, error)

	GetCwfNetworkIDGatewaysGatewayIDName(params *GetCwfNetworkIDGatewaysGatewayIDNameParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDNameOK, error)

	GetCwfNetworkIDGatewaysGatewayIDStatus(params *GetCwfNetworkIDGatewaysGatewayIDStatusParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDStatusOK, error)

	GetCwfNetworkIDGatewaysGatewayIDTier(params *GetCwfNetworkIDGatewaysGatewayIDTierParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDTierOK, error)

	PostCwfNetworkIDGateways(params *PostCwfNetworkIDGatewaysParams, opts ...ClientOption) (*PostCwfNetworkIDGatewaysCreated, error)

	PutCwfNetworkIDGatewaysGatewayID(params *PutCwfNetworkIDGatewaysGatewayIDParams, opts ...ClientOption) (*PutCwfNetworkIDGatewaysGatewayIDNoContent, error)

	PutCwfNetworkIDGatewaysGatewayIDCarrierWifi(params *PutCwfNetworkIDGatewaysGatewayIDCarrierWifiParams, opts ...ClientOption) (*PutCwfNetworkIDGatewaysGatewayIDCarrierWifiNoContent, error)

	PutCwfNetworkIDGatewaysGatewayIDDescription(params *PutCwfNetworkIDGatewaysGatewayIDDescriptionParams, opts ...ClientOption) (*PutCwfNetworkIDGatewaysGatewayIDDescriptionNoContent, error)

	PutCwfNetworkIDGatewaysGatewayIDDevice(params *PutCwfNetworkIDGatewaysGatewayIDDeviceParams, opts ...ClientOption) (*PutCwfNetworkIDGatewaysGatewayIDDeviceNoContent, error)

	PutCwfNetworkIDGatewaysGatewayIDMagmad(params *PutCwfNetworkIDGatewaysGatewayIDMagmadParams, opts ...ClientOption) (*PutCwfNetworkIDGatewaysGatewayIDMagmadNoContent, error)

	PutCwfNetworkIDGatewaysGatewayIDName(params *PutCwfNetworkIDGatewaysGatewayIDNameParams, opts ...ClientOption) (*PutCwfNetworkIDGatewaysGatewayIDNameNoContent, error)

	PutCwfNetworkIDGatewaysGatewayIDTier(params *PutCwfNetworkIDGatewaysGatewayIDTierParams, opts ...ClientOption) (*PutCwfNetworkIDGatewaysGatewayIDTierNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteCwfNetworkIDGatewaysGatewayID deletes a carrier wifi gateway
*/
func (a *Client) DeleteCwfNetworkIDGatewaysGatewayID(params *DeleteCwfNetworkIDGatewaysGatewayIDParams, opts ...ClientOption) (*DeleteCwfNetworkIDGatewaysGatewayIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCwfNetworkIDGatewaysGatewayIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteCwfNetworkIDGatewaysGatewayID",
		Method:             "DELETE",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteCwfNetworkIDGatewaysGatewayIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteCwfNetworkIDGatewaysGatewayIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteCwfNetworkIDGatewaysGatewayIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCwfNetworkIDGateways lists all gateways for a carrier wifi network
*/
func (a *Client) GetCwfNetworkIDGateways(params *GetCwfNetworkIDGatewaysParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGateways",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysReader{formats: a.formats},
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
	success, ok := result.(*GetCwfNetworkIDGatewaysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCwfNetworkIDGatewaysGatewayID gets a specific carrier wifi gateway
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayID(params *GetCwfNetworkIDGatewaysGatewayIDParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayID",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDReader{formats: a.formats},
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
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCwfNetworkIDGatewaysGatewayIDCarrierWifi gets gateway carrier wifi configuration
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayIDCarrierWifi(params *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDCarrierWifiOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayIDCarrierWifi",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/carrier_wifi",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDCarrierWifiReader{formats: a.formats},
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
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDCarrierWifiOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDCarrierWifiDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCwfNetworkIDGatewaysGatewayIDDescription gets the description of a carrier wifi gateway
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayIDDescription(params *GetCwfNetworkIDGatewaysGatewayIDDescriptionParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDDescriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDDescriptionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayIDDescription",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/description",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDDescriptionReader{formats: a.formats},
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
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDDescriptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDDescriptionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCwfNetworkIDGatewaysGatewayIDDevice gets the physical device for a carrier wifi gateway
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayIDDevice(params *GetCwfNetworkIDGatewaysGatewayIDDeviceParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDDeviceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayIDDevice",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/device",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDDeviceReader{formats: a.formats},
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
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDDeviceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDDeviceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCwfNetworkIDGatewaysGatewayIDHealthStatus retrieves health status of a carrier wifi gateway
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayIDHealthStatus(params *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDHealthStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDHealthStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayIDHealthStatus",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/health_status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDHealthStatusReader{formats: a.formats},
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
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDHealthStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDHealthStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCwfNetworkIDGatewaysGatewayIDMagmad gets magmad agent configuration
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayIDMagmad(params *GetCwfNetworkIDGatewaysGatewayIDMagmadParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDMagmadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDMagmadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayIDMagmad",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/magmad",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDMagmadReader{formats: a.formats},
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
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDMagmadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDMagmadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCwfNetworkIDGatewaysGatewayIDName gets the name of a carrier wifi gateway
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayIDName(params *GetCwfNetworkIDGatewaysGatewayIDNameParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayIDName",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/name",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDNameReader{formats: a.formats},
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
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCwfNetworkIDGatewaysGatewayIDStatus gets the status of a gateway
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayIDStatus(params *GetCwfNetworkIDGatewaysGatewayIDStatusParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayIDStatus",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDStatusReader{formats: a.formats},
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
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCwfNetworkIDGatewaysGatewayIDTier gets the ID of the upgrade tier a gateway belongs to
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayIDTier(params *GetCwfNetworkIDGatewaysGatewayIDTierParams, opts ...ClientOption) (*GetCwfNetworkIDGatewaysGatewayIDTierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDTierParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayIDTier",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/tier",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDTierReader{formats: a.formats},
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
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDTierOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDTierDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostCwfNetworkIDGateways registers a new carrier wifi gateway
*/
func (a *Client) PostCwfNetworkIDGateways(params *PostCwfNetworkIDGatewaysParams, opts ...ClientOption) (*PostCwfNetworkIDGatewaysCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCwfNetworkIDGatewaysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostCwfNetworkIDGateways",
		Method:             "POST",
		PathPattern:        "/cwf/{network_id}/gateways",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostCwfNetworkIDGatewaysReader{formats: a.formats},
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
	success, ok := result.(*PostCwfNetworkIDGatewaysCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostCwfNetworkIDGatewaysDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutCwfNetworkIDGatewaysGatewayID updates an entire carrier wifi gateway record
*/
func (a *Client) PutCwfNetworkIDGatewaysGatewayID(params *PutCwfNetworkIDGatewaysGatewayIDParams, opts ...ClientOption) (*PutCwfNetworkIDGatewaysGatewayIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCwfNetworkIDGatewaysGatewayIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutCwfNetworkIDGatewaysGatewayID",
		Method:             "PUT",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCwfNetworkIDGatewaysGatewayIDReader{formats: a.formats},
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
	success, ok := result.(*PutCwfNetworkIDGatewaysGatewayIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutCwfNetworkIDGatewaysGatewayIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutCwfNetworkIDGatewaysGatewayIDCarrierWifi updates gateway carrier wifi configuration
*/
func (a *Client) PutCwfNetworkIDGatewaysGatewayIDCarrierWifi(params *PutCwfNetworkIDGatewaysGatewayIDCarrierWifiParams, opts ...ClientOption) (*PutCwfNetworkIDGatewaysGatewayIDCarrierWifiNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCwfNetworkIDGatewaysGatewayIDCarrierWifiParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutCwfNetworkIDGatewaysGatewayIDCarrierWifi",
		Method:             "PUT",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/carrier_wifi",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCwfNetworkIDGatewaysGatewayIDCarrierWifiReader{formats: a.formats},
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
	success, ok := result.(*PutCwfNetworkIDGatewaysGatewayIDCarrierWifiNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutCwfNetworkIDGatewaysGatewayIDCarrierWifiDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutCwfNetworkIDGatewaysGatewayIDDescription updates the description of a carrier wifi gateway
*/
func (a *Client) PutCwfNetworkIDGatewaysGatewayIDDescription(params *PutCwfNetworkIDGatewaysGatewayIDDescriptionParams, opts ...ClientOption) (*PutCwfNetworkIDGatewaysGatewayIDDescriptionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCwfNetworkIDGatewaysGatewayIDDescriptionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutCwfNetworkIDGatewaysGatewayIDDescription",
		Method:             "PUT",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/description",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCwfNetworkIDGatewaysGatewayIDDescriptionReader{formats: a.formats},
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
	success, ok := result.(*PutCwfNetworkIDGatewaysGatewayIDDescriptionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutCwfNetworkIDGatewaysGatewayIDDescriptionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutCwfNetworkIDGatewaysGatewayIDDevice updates the physical device for a carrier wifi gateway
*/
func (a *Client) PutCwfNetworkIDGatewaysGatewayIDDevice(params *PutCwfNetworkIDGatewaysGatewayIDDeviceParams, opts ...ClientOption) (*PutCwfNetworkIDGatewaysGatewayIDDeviceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCwfNetworkIDGatewaysGatewayIDDeviceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutCwfNetworkIDGatewaysGatewayIDDevice",
		Method:             "PUT",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/device",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCwfNetworkIDGatewaysGatewayIDDeviceReader{formats: a.formats},
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
	success, ok := result.(*PutCwfNetworkIDGatewaysGatewayIDDeviceNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutCwfNetworkIDGatewaysGatewayIDDeviceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutCwfNetworkIDGatewaysGatewayIDMagmad reconfigures magmad agent
*/
func (a *Client) PutCwfNetworkIDGatewaysGatewayIDMagmad(params *PutCwfNetworkIDGatewaysGatewayIDMagmadParams, opts ...ClientOption) (*PutCwfNetworkIDGatewaysGatewayIDMagmadNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCwfNetworkIDGatewaysGatewayIDMagmadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutCwfNetworkIDGatewaysGatewayIDMagmad",
		Method:             "PUT",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/magmad",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCwfNetworkIDGatewaysGatewayIDMagmadReader{formats: a.formats},
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
	success, ok := result.(*PutCwfNetworkIDGatewaysGatewayIDMagmadNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutCwfNetworkIDGatewaysGatewayIDMagmadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutCwfNetworkIDGatewaysGatewayIDName updates the name of a carrier wifi gateway
*/
func (a *Client) PutCwfNetworkIDGatewaysGatewayIDName(params *PutCwfNetworkIDGatewaysGatewayIDNameParams, opts ...ClientOption) (*PutCwfNetworkIDGatewaysGatewayIDNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCwfNetworkIDGatewaysGatewayIDNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutCwfNetworkIDGatewaysGatewayIDName",
		Method:             "PUT",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/name",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCwfNetworkIDGatewaysGatewayIDNameReader{formats: a.formats},
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
	success, ok := result.(*PutCwfNetworkIDGatewaysGatewayIDNameNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutCwfNetworkIDGatewaysGatewayIDNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutCwfNetworkIDGatewaysGatewayIDTier updates the ID of the upgrade tier a gateway belongs to
*/
func (a *Client) PutCwfNetworkIDGatewaysGatewayIDTier(params *PutCwfNetworkIDGatewaysGatewayIDTierParams, opts ...ClientOption) (*PutCwfNetworkIDGatewaysGatewayIDTierNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCwfNetworkIDGatewaysGatewayIDTierParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutCwfNetworkIDGatewaysGatewayIDTier",
		Method:             "PUT",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/tier",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCwfNetworkIDGatewaysGatewayIDTierReader{formats: a.formats},
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
	success, ok := result.(*PutCwfNetworkIDGatewaysGatewayIDTierNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutCwfNetworkIDGatewaysGatewayIDTierDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
