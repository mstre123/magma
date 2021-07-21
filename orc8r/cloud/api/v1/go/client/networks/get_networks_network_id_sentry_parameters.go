// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewGetNetworksNetworkIDSentryParams creates a new GetNetworksNetworkIDSentryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworksNetworkIDSentryParams() *GetNetworksNetworkIDSentryParams {
	return &GetNetworksNetworkIDSentryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksNetworkIDSentryParamsWithTimeout creates a new GetNetworksNetworkIDSentryParams object
// with the ability to set a timeout on a request.
func NewGetNetworksNetworkIDSentryParamsWithTimeout(timeout time.Duration) *GetNetworksNetworkIDSentryParams {
	return &GetNetworksNetworkIDSentryParams{
		timeout: timeout,
	}
}

// NewGetNetworksNetworkIDSentryParamsWithContext creates a new GetNetworksNetworkIDSentryParams object
// with the ability to set a context for a request.
func NewGetNetworksNetworkIDSentryParamsWithContext(ctx context.Context) *GetNetworksNetworkIDSentryParams {
	return &GetNetworksNetworkIDSentryParams{
		Context: ctx,
	}
}

// NewGetNetworksNetworkIDSentryParamsWithHTTPClient creates a new GetNetworksNetworkIDSentryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworksNetworkIDSentryParamsWithHTTPClient(client *http.Client) *GetNetworksNetworkIDSentryParams {
	return &GetNetworksNetworkIDSentryParams{
		HTTPClient: client,
	}
}

/* GetNetworksNetworkIDSentryParams contains all the parameters to send to the API endpoint
   for the get networks network ID sentry operation.

   Typically these are written to a http.Request.
*/
type GetNetworksNetworkIDSentryParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get networks network ID sentry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworksNetworkIDSentryParams) WithDefaults() *GetNetworksNetworkIDSentryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get networks network ID sentry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworksNetworkIDSentryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get networks network ID sentry params
func (o *GetNetworksNetworkIDSentryParams) WithTimeout(timeout time.Duration) *GetNetworksNetworkIDSentryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks network ID sentry params
func (o *GetNetworksNetworkIDSentryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks network ID sentry params
func (o *GetNetworksNetworkIDSentryParams) WithContext(ctx context.Context) *GetNetworksNetworkIDSentryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks network ID sentry params
func (o *GetNetworksNetworkIDSentryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks network ID sentry params
func (o *GetNetworksNetworkIDSentryParams) WithHTTPClient(client *http.Client) *GetNetworksNetworkIDSentryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks network ID sentry params
func (o *GetNetworksNetworkIDSentryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get networks network ID sentry params
func (o *GetNetworksNetworkIDSentryParams) WithNetworkID(networkID string) *GetNetworksNetworkIDSentryParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get networks network ID sentry params
func (o *GetNetworksNetworkIDSentryParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksNetworkIDSentryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
