// Code generated by go-swagger; DO NOT EDIT.

package wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// NewPostWifiParams creates a new PostWifiParams object
// with the default values initialized.
func NewPostWifiParams() *PostWifiParams {
	var ()
	return &PostWifiParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWifiParamsWithTimeout creates a new PostWifiParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWifiParamsWithTimeout(timeout time.Duration) *PostWifiParams {
	var ()
	return &PostWifiParams{

		timeout: timeout,
	}
}

// NewPostWifiParamsWithContext creates a new PostWifiParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWifiParamsWithContext(ctx context.Context) *PostWifiParams {
	var ()
	return &PostWifiParams{

		Context: ctx,
	}
}

// NewPostWifiParamsWithHTTPClient creates a new PostWifiParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostWifiParamsWithHTTPClient(client *http.Client) *PostWifiParams {
	var ()
	return &PostWifiParams{
		HTTPClient: client,
	}
}

/*PostWifiParams contains all the parameters to send to the API endpoint
for the post wifi operation typically these are written to a http.Request
*/
type PostWifiParams struct {

	/*WifiNetwork
	  Configuration of the Wifi network to create

	*/
	WifiNetwork *models.WifiNetwork

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post wifi params
func (o *PostWifiParams) WithTimeout(timeout time.Duration) *PostWifiParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post wifi params
func (o *PostWifiParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post wifi params
func (o *PostWifiParams) WithContext(ctx context.Context) *PostWifiParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post wifi params
func (o *PostWifiParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post wifi params
func (o *PostWifiParams) WithHTTPClient(client *http.Client) *PostWifiParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post wifi params
func (o *PostWifiParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWifiNetwork adds the wifiNetwork to the post wifi params
func (o *PostWifiParams) WithWifiNetwork(wifiNetwork *models.WifiNetwork) *PostWifiParams {
	o.SetWifiNetwork(wifiNetwork)
	return o
}

// SetWifiNetwork adds the wifiNetwork to the post wifi params
func (o *PostWifiParams) SetWifiNetwork(wifiNetwork *models.WifiNetwork) {
	o.WifiNetwork = wifiNetwork
}

// WriteToRequest writes these params to a swagger request
func (o *PostWifiParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.WifiNetwork != nil {
		if err := r.SetBodyParam(o.WifiNetwork); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
