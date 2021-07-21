// Code generated by go-swagger; DO NOT EDIT.

package commands

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams creates a new PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams() *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams {
	return &PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostNetworksNetworkIDGatewaysGatewayIDCommandRebootParamsWithTimeout creates a new PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams object
// with the ability to set a timeout on a request.
func NewPostNetworksNetworkIDGatewaysGatewayIDCommandRebootParamsWithTimeout(timeout time.Duration) *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams {
	return &PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams{
		timeout: timeout,
	}
}

// NewPostNetworksNetworkIDGatewaysGatewayIDCommandRebootParamsWithContext creates a new PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams object
// with the ability to set a context for a request.
func NewPostNetworksNetworkIDGatewaysGatewayIDCommandRebootParamsWithContext(ctx context.Context) *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams {
	return &PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams{
		Context: ctx,
	}
}

// NewPostNetworksNetworkIDGatewaysGatewayIDCommandRebootParamsWithHTTPClient creates a new PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostNetworksNetworkIDGatewaysGatewayIDCommandRebootParamsWithHTTPClient(client *http.Client) *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams {
	return &PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams{
		HTTPClient: client,
	}
}

/* PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams contains all the parameters to send to the API endpoint
   for the post networks network ID gateways gateway ID command reboot operation.

   Typically these are written to a http.Request.
*/
type PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams struct {

	/* GatewayID.

	   Gateway ID
	*/
	GatewayID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post networks network ID gateways gateway ID command reboot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams) WithDefaults() *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post networks network ID gateways gateway ID command reboot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post networks network ID gateways gateway ID command reboot params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams) WithTimeout(timeout time.Duration) *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post networks network ID gateways gateway ID command reboot params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post networks network ID gateways gateway ID command reboot params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams) WithContext(ctx context.Context) *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post networks network ID gateways gateway ID command reboot params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post networks network ID gateways gateway ID command reboot params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams) WithHTTPClient(client *http.Client) *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post networks network ID gateways gateway ID command reboot params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the post networks network ID gateways gateway ID command reboot params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams) WithGatewayID(gatewayID string) *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the post networks network ID gateways gateway ID command reboot params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the post networks network ID gateways gateway ID command reboot params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams) WithNetworkID(networkID string) *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post networks network ID gateways gateway ID command reboot params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandRebootParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
