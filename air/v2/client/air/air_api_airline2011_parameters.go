// Code generated by go-swagger; DO NOT EDIT.

package air

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

// NewAirAPIAirline2011Params creates a new AirAPIAirline2011Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAirAPIAirline2011Params() *AirAPIAirline2011Params {
	return &AirAPIAirline2011Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewAirAPIAirline2011ParamsWithTimeout creates a new AirAPIAirline2011Params object
// with the ability to set a timeout on a request.
func NewAirAPIAirline2011ParamsWithTimeout(timeout time.Duration) *AirAPIAirline2011Params {
	return &AirAPIAirline2011Params{
		timeout: timeout,
	}
}

// NewAirAPIAirline2011ParamsWithContext creates a new AirAPIAirline2011Params object
// with the ability to set a context for a request.
func NewAirAPIAirline2011ParamsWithContext(ctx context.Context) *AirAPIAirline2011Params {
	return &AirAPIAirline2011Params{
		Context: ctx,
	}
}

// NewAirAPIAirline2011ParamsWithHTTPClient creates a new AirAPIAirline2011Params object
// with the ability to set a custom HTTPClient for a request.
func NewAirAPIAirline2011ParamsWithHTTPClient(client *http.Client) *AirAPIAirline2011Params {
	return &AirAPIAirline2011Params{
		HTTPClient: client,
	}
}

/* AirAPIAirline2011Params contains all the parameters to send to the API endpoint
   for the air Api airline 2011 operation.

   Typically these are written to a http.Request.
*/
type AirAPIAirline2011Params struct {

	/* DollarFormat.

	   指定來源格式
	*/
	DollarFormat string

	/* IATA.

	   航空公司代碼
	*/
	IATA string

	/* Health.

	   加入參數'?health=true'即可查詢此API服務的健康狀態
	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the air Api airline 2011 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AirAPIAirline2011Params) WithDefaults() *AirAPIAirline2011Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the air Api airline 2011 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AirAPIAirline2011Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the air Api airline 2011 params
func (o *AirAPIAirline2011Params) WithTimeout(timeout time.Duration) *AirAPIAirline2011Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the air Api airline 2011 params
func (o *AirAPIAirline2011Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the air Api airline 2011 params
func (o *AirAPIAirline2011Params) WithContext(ctx context.Context) *AirAPIAirline2011Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the air Api airline 2011 params
func (o *AirAPIAirline2011Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the air Api airline 2011 params
func (o *AirAPIAirline2011Params) WithHTTPClient(client *http.Client) *AirAPIAirline2011Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the air Api airline 2011 params
func (o *AirAPIAirline2011Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFormat adds the dollarFormat to the air Api airline 2011 params
func (o *AirAPIAirline2011Params) WithDollarFormat(dollarFormat string) *AirAPIAirline2011Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the air Api airline 2011 params
func (o *AirAPIAirline2011Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithIATA adds the iATA to the air Api airline 2011 params
func (o *AirAPIAirline2011Params) WithIATA(iATA string) *AirAPIAirline2011Params {
	o.SetIATA(iATA)
	return o
}

// SetIATA adds the iATA to the air Api airline 2011 params
func (o *AirAPIAirline2011Params) SetIATA(iATA string) {
	o.IATA = iATA
}

// WithHealth adds the health to the air Api airline 2011 params
func (o *AirAPIAirline2011Params) WithHealth(health *string) *AirAPIAirline2011Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the air Api airline 2011 params
func (o *AirAPIAirline2011Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *AirAPIAirline2011Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param $format
	qrDollarFormat := o.DollarFormat
	qDollarFormat := qrDollarFormat
	if qDollarFormat != "" {

		if err := r.SetQueryParam("$format", qDollarFormat); err != nil {
			return err
		}
	}

	// path param IATA
	if err := r.SetPathParam("IATA", o.IATA); err != nil {
		return err
	}

	if o.Health != nil {

		// query param health
		var qrHealth string

		if o.Health != nil {
			qrHealth = *o.Health
		}
		qHealth := qrHealth
		if qHealth != "" {

			if err := r.SetQueryParam("health", qHealth); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
