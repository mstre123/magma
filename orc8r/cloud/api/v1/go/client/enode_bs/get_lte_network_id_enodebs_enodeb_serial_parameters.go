// Code generated by go-swagger; DO NOT EDIT.

package enode_bs

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

// NewGetLTENetworkIDEnodebsENODEBSerialParams creates a new GetLTENetworkIDEnodebsENODEBSerialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLTENetworkIDEnodebsENODEBSerialParams() *GetLTENetworkIDEnodebsENODEBSerialParams {
	return &GetLTENetworkIDEnodebsENODEBSerialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDEnodebsENODEBSerialParamsWithTimeout creates a new GetLTENetworkIDEnodebsENODEBSerialParams object
// with the ability to set a timeout on a request.
func NewGetLTENetworkIDEnodebsENODEBSerialParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDEnodebsENODEBSerialParams {
	return &GetLTENetworkIDEnodebsENODEBSerialParams{
		timeout: timeout,
	}
}

// NewGetLTENetworkIDEnodebsENODEBSerialParamsWithContext creates a new GetLTENetworkIDEnodebsENODEBSerialParams object
// with the ability to set a context for a request.
func NewGetLTENetworkIDEnodebsENODEBSerialParamsWithContext(ctx context.Context) *GetLTENetworkIDEnodebsENODEBSerialParams {
	return &GetLTENetworkIDEnodebsENODEBSerialParams{
		Context: ctx,
	}
}

// NewGetLTENetworkIDEnodebsENODEBSerialParamsWithHTTPClient creates a new GetLTENetworkIDEnodebsENODEBSerialParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLTENetworkIDEnodebsENODEBSerialParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDEnodebsENODEBSerialParams {
	return &GetLTENetworkIDEnodebsENODEBSerialParams{
		HTTPClient: client,
	}
}

/* GetLTENetworkIDEnodebsENODEBSerialParams contains all the parameters to send to the API endpoint
   for the get LTE network ID enodebs ENODEB serial operation.

   Typically these are written to a http.Request.
*/
type GetLTENetworkIDEnodebsENODEBSerialParams struct {

	/* ENODEBSerial.

	   EnodeB serial number
	*/
	ENODEBSerial string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get LTE network ID enodebs ENODEB serial params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLTENetworkIDEnodebsENODEBSerialParams) WithDefaults() *GetLTENetworkIDEnodebsENODEBSerialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get LTE network ID enodebs ENODEB serial params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLTENetworkIDEnodebsENODEBSerialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get LTE network ID enodebs ENODEB serial params
func (o *GetLTENetworkIDEnodebsENODEBSerialParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDEnodebsENODEBSerialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID enodebs ENODEB serial params
func (o *GetLTENetworkIDEnodebsENODEBSerialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID enodebs ENODEB serial params
func (o *GetLTENetworkIDEnodebsENODEBSerialParams) WithContext(ctx context.Context) *GetLTENetworkIDEnodebsENODEBSerialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID enodebs ENODEB serial params
func (o *GetLTENetworkIDEnodebsENODEBSerialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID enodebs ENODEB serial params
func (o *GetLTENetworkIDEnodebsENODEBSerialParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDEnodebsENODEBSerialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID enodebs ENODEB serial params
func (o *GetLTENetworkIDEnodebsENODEBSerialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithENODEBSerial adds the eNODEBSerial to the get LTE network ID enodebs ENODEB serial params
func (o *GetLTENetworkIDEnodebsENODEBSerialParams) WithENODEBSerial(eNODEBSerial string) *GetLTENetworkIDEnodebsENODEBSerialParams {
	o.SetENODEBSerial(eNODEBSerial)
	return o
}

// SetENODEBSerial adds the enodebSerial to the get LTE network ID enodebs ENODEB serial params
func (o *GetLTENetworkIDEnodebsENODEBSerialParams) SetENODEBSerial(eNODEBSerial string) {
	o.ENODEBSerial = eNODEBSerial
}

// WithNetworkID adds the networkID to the get LTE network ID enodebs ENODEB serial params
func (o *GetLTENetworkIDEnodebsENODEBSerialParams) WithNetworkID(networkID string) *GetLTENetworkIDEnodebsENODEBSerialParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID enodebs ENODEB serial params
func (o *GetLTENetworkIDEnodebsENODEBSerialParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDEnodebsENODEBSerialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param enodeb_serial
	if err := r.SetPathParam("enodeb_serial", o.ENODEBSerial); err != nil {
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
