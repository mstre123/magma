// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

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

// NewGetLTENetworkIDParams creates a new GetLTENetworkIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLTENetworkIDParams() *GetLTENetworkIDParams {
	return &GetLTENetworkIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDParamsWithTimeout creates a new GetLTENetworkIDParams object
// with the ability to set a timeout on a request.
func NewGetLTENetworkIDParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDParams {
	return &GetLTENetworkIDParams{
		timeout: timeout,
	}
}

// NewGetLTENetworkIDParamsWithContext creates a new GetLTENetworkIDParams object
// with the ability to set a context for a request.
func NewGetLTENetworkIDParamsWithContext(ctx context.Context) *GetLTENetworkIDParams {
	return &GetLTENetworkIDParams{
		Context: ctx,
	}
}

// NewGetLTENetworkIDParamsWithHTTPClient creates a new GetLTENetworkIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLTENetworkIDParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDParams {
	return &GetLTENetworkIDParams{
		HTTPClient: client,
	}
}

/* GetLTENetworkIDParams contains all the parameters to send to the API endpoint
   for the get LTE network ID operation.

   Typically these are written to a http.Request.
*/
type GetLTENetworkIDParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get LTE network ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLTENetworkIDParams) WithDefaults() *GetLTENetworkIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get LTE network ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLTENetworkIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get LTE network ID params
func (o *GetLTENetworkIDParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID params
func (o *GetLTENetworkIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID params
func (o *GetLTENetworkIDParams) WithContext(ctx context.Context) *GetLTENetworkIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID params
func (o *GetLTENetworkIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID params
func (o *GetLTENetworkIDParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID params
func (o *GetLTENetworkIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get LTE network ID params
func (o *GetLTENetworkIDParams) WithNetworkID(networkID string) *GetLTENetworkIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID params
func (o *GetLTENetworkIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
