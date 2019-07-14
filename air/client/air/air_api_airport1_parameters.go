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

	strfmt "github.com/go-openapi/strfmt"
)

// NewAirAPIAirport1Params creates a new AirAPIAirport1Params object
// with the default values initialized.
func NewAirAPIAirport1Params() *AirAPIAirport1Params {
	var ()
	return &AirAPIAirport1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAirAPIAirport1ParamsWithTimeout creates a new AirAPIAirport1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAirAPIAirport1ParamsWithTimeout(timeout time.Duration) *AirAPIAirport1Params {
	var ()
	return &AirAPIAirport1Params{

		timeout: timeout,
	}
}

// NewAirAPIAirport1ParamsWithContext creates a new AirAPIAirport1Params object
// with the default values initialized, and the ability to set a context for a request
func NewAirAPIAirport1ParamsWithContext(ctx context.Context) *AirAPIAirport1Params {
	var ()
	return &AirAPIAirport1Params{

		Context: ctx,
	}
}

// NewAirAPIAirport1ParamsWithHTTPClient creates a new AirAPIAirport1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAirAPIAirport1ParamsWithHTTPClient(client *http.Client) *AirAPIAirport1Params {
	var ()
	return &AirAPIAirport1Params{
		HTTPClient: client,
	}
}

/*AirAPIAirport1Params contains all the parameters to send to the API endpoint
for the air Api airport 1 operation typically these are written to a http.Request
*/
type AirAPIAirport1Params struct {

	/*DollarFormat
	  指定來源格式

	*/
	DollarFormat string
	/*IATA
	  機場代碼

	*/
	IATA string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the air Api airport 1 params
func (o *AirAPIAirport1Params) WithTimeout(timeout time.Duration) *AirAPIAirport1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the air Api airport 1 params
func (o *AirAPIAirport1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the air Api airport 1 params
func (o *AirAPIAirport1Params) WithContext(ctx context.Context) *AirAPIAirport1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the air Api airport 1 params
func (o *AirAPIAirport1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the air Api airport 1 params
func (o *AirAPIAirport1Params) WithHTTPClient(client *http.Client) *AirAPIAirport1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the air Api airport 1 params
func (o *AirAPIAirport1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFormat adds the dollarFormat to the air Api airport 1 params
func (o *AirAPIAirport1Params) WithDollarFormat(dollarFormat string) *AirAPIAirport1Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the air Api airport 1 params
func (o *AirAPIAirport1Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithIATA adds the iATA to the air Api airport 1 params
func (o *AirAPIAirport1Params) WithIATA(iATA string) *AirAPIAirport1Params {
	o.SetIATA(iATA)
	return o
}

// SetIATA adds the iATA to the air Api airport 1 params
func (o *AirAPIAirport1Params) SetIATA(iATA string) {
	o.IATA = iATA
}

// WriteToRequest writes these params to a swagger request
func (o *AirAPIAirport1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
