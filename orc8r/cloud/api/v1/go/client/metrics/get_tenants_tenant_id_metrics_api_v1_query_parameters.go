// Code generated by go-swagger; DO NOT EDIT.

package metrics

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
	"github.com/go-openapi/swag"
)

// NewGetTenantsTenantIDMetricsAPIV1QueryParams creates a new GetTenantsTenantIDMetricsAPIV1QueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTenantsTenantIDMetricsAPIV1QueryParams() *GetTenantsTenantIDMetricsAPIV1QueryParams {
	return &GetTenantsTenantIDMetricsAPIV1QueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTenantsTenantIDMetricsAPIV1QueryParamsWithTimeout creates a new GetTenantsTenantIDMetricsAPIV1QueryParams object
// with the ability to set a timeout on a request.
func NewGetTenantsTenantIDMetricsAPIV1QueryParamsWithTimeout(timeout time.Duration) *GetTenantsTenantIDMetricsAPIV1QueryParams {
	return &GetTenantsTenantIDMetricsAPIV1QueryParams{
		timeout: timeout,
	}
}

// NewGetTenantsTenantIDMetricsAPIV1QueryParamsWithContext creates a new GetTenantsTenantIDMetricsAPIV1QueryParams object
// with the ability to set a context for a request.
func NewGetTenantsTenantIDMetricsAPIV1QueryParamsWithContext(ctx context.Context) *GetTenantsTenantIDMetricsAPIV1QueryParams {
	return &GetTenantsTenantIDMetricsAPIV1QueryParams{
		Context: ctx,
	}
}

// NewGetTenantsTenantIDMetricsAPIV1QueryParamsWithHTTPClient creates a new GetTenantsTenantIDMetricsAPIV1QueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTenantsTenantIDMetricsAPIV1QueryParamsWithHTTPClient(client *http.Client) *GetTenantsTenantIDMetricsAPIV1QueryParams {
	return &GetTenantsTenantIDMetricsAPIV1QueryParams{
		HTTPClient: client,
	}
}

/* GetTenantsTenantIDMetricsAPIV1QueryParams contains all the parameters to send to the API endpoint
   for the get tenants tenant ID metrics API v1 query operation.

   Typically these are written to a http.Request.
*/
type GetTenantsTenantIDMetricsAPIV1QueryParams struct {

	/* Query.

	   PromQL query to proxy to prometheus
	*/
	Query string

	/* TenantID.

	   Tenant ID
	*/
	TenantID int64

	/* Time.

	   time for query (UnixTime or RFC3339)
	*/
	Time *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get tenants tenant ID metrics API v1 query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTenantsTenantIDMetricsAPIV1QueryParams) WithDefaults() *GetTenantsTenantIDMetricsAPIV1QueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tenants tenant ID metrics API v1 query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTenantsTenantIDMetricsAPIV1QueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get tenants tenant ID metrics API v1 query params
func (o *GetTenantsTenantIDMetricsAPIV1QueryParams) WithTimeout(timeout time.Duration) *GetTenantsTenantIDMetricsAPIV1QueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tenants tenant ID metrics API v1 query params
func (o *GetTenantsTenantIDMetricsAPIV1QueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tenants tenant ID metrics API v1 query params
func (o *GetTenantsTenantIDMetricsAPIV1QueryParams) WithContext(ctx context.Context) *GetTenantsTenantIDMetricsAPIV1QueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tenants tenant ID metrics API v1 query params
func (o *GetTenantsTenantIDMetricsAPIV1QueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tenants tenant ID metrics API v1 query params
func (o *GetTenantsTenantIDMetricsAPIV1QueryParams) WithHTTPClient(client *http.Client) *GetTenantsTenantIDMetricsAPIV1QueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tenants tenant ID metrics API v1 query params
func (o *GetTenantsTenantIDMetricsAPIV1QueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the get tenants tenant ID metrics API v1 query params
func (o *GetTenantsTenantIDMetricsAPIV1QueryParams) WithQuery(query string) *GetTenantsTenantIDMetricsAPIV1QueryParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get tenants tenant ID metrics API v1 query params
func (o *GetTenantsTenantIDMetricsAPIV1QueryParams) SetQuery(query string) {
	o.Query = query
}

// WithTenantID adds the tenantID to the get tenants tenant ID metrics API v1 query params
func (o *GetTenantsTenantIDMetricsAPIV1QueryParams) WithTenantID(tenantID int64) *GetTenantsTenantIDMetricsAPIV1QueryParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get tenants tenant ID metrics API v1 query params
func (o *GetTenantsTenantIDMetricsAPIV1QueryParams) SetTenantID(tenantID int64) {
	o.TenantID = tenantID
}

// WithTime adds the time to the get tenants tenant ID metrics API v1 query params
func (o *GetTenantsTenantIDMetricsAPIV1QueryParams) WithTime(time *string) *GetTenantsTenantIDMetricsAPIV1QueryParams {
	o.SetTime(time)
	return o
}

// SetTime adds the time to the get tenants tenant ID metrics API v1 query params
func (o *GetTenantsTenantIDMetricsAPIV1QueryParams) SetTime(time *string) {
	o.Time = time
}

// WriteToRequest writes these params to a swagger request
func (o *GetTenantsTenantIDMetricsAPIV1QueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param query
	qrQuery := o.Query
	qQuery := qrQuery
	if qQuery != "" {

		if err := r.SetQueryParam("query", qQuery); err != nil {
			return err
		}
	}

	// path param tenant_id
	if err := r.SetPathParam("tenant_id", swag.FormatInt64(o.TenantID)); err != nil {
		return err
	}

	if o.Time != nil {

		// query param time
		var qrTime string

		if o.Time != nil {
			qrTime = *o.Time
		}
		qTime := qrTime
		if qTime != "" {

			if err := r.SetQueryParam("time", qTime); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
