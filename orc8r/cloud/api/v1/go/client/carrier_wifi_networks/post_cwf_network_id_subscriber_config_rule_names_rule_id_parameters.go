// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

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

// NewPostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams creates a new PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams() *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	return &PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParamsWithTimeout creates a new PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams object
// with the ability to set a timeout on a request.
func NewPostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParamsWithTimeout(timeout time.Duration) *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	return &PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams{
		timeout: timeout,
	}
}

// NewPostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParamsWithContext creates a new PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams object
// with the ability to set a context for a request.
func NewPostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParamsWithContext(ctx context.Context) *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	return &PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams{
		Context: ctx,
	}
}

// NewPostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParamsWithHTTPClient creates a new PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParamsWithHTTPClient(client *http.Client) *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	return &PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams{
		HTTPClient: client,
	}
}

/* PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams contains all the parameters to send to the API endpoint
   for the post cwf network ID subscriber config rule names rule ID operation.

   Typically these are written to a http.Request.
*/
type PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* RuleID.

	   Rule Id
	*/
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post cwf network ID subscriber config rule names rule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams) WithDefaults() *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post cwf network ID subscriber config rule names rule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post cwf network ID subscriber config rule names rule ID params
func (o *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams) WithTimeout(timeout time.Duration) *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cwf network ID subscriber config rule names rule ID params
func (o *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cwf network ID subscriber config rule names rule ID params
func (o *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams) WithContext(ctx context.Context) *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cwf network ID subscriber config rule names rule ID params
func (o *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cwf network ID subscriber config rule names rule ID params
func (o *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams) WithHTTPClient(client *http.Client) *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cwf network ID subscriber config rule names rule ID params
func (o *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the post cwf network ID subscriber config rule names rule ID params
func (o *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams) WithNetworkID(networkID string) *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post cwf network ID subscriber config rule names rule ID params
func (o *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRuleID adds the ruleID to the post cwf network ID subscriber config rule names rule ID params
func (o *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams) WithRuleID(ruleID string) *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the post cwf network ID subscriber config rule names rule ID params
func (o *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *PostCwfNetworkIDSubscriberConfigRuleNamesRuleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	// path param rule_id
	if err := r.SetPathParam("rule_id", o.RuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
