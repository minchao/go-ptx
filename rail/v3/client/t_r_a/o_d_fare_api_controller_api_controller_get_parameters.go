// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

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

// NewODFareAPIControllerAPIControllerGetParams creates a new ODFareAPIControllerAPIControllerGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewODFareAPIControllerAPIControllerGetParams() *ODFareAPIControllerAPIControllerGetParams {
	return &ODFareAPIControllerAPIControllerGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewODFareAPIControllerAPIControllerGetParamsWithTimeout creates a new ODFareAPIControllerAPIControllerGetParams object
// with the ability to set a timeout on a request.
func NewODFareAPIControllerAPIControllerGetParamsWithTimeout(timeout time.Duration) *ODFareAPIControllerAPIControllerGetParams {
	return &ODFareAPIControllerAPIControllerGetParams{
		timeout: timeout,
	}
}

// NewODFareAPIControllerAPIControllerGetParamsWithContext creates a new ODFareAPIControllerAPIControllerGetParams object
// with the ability to set a context for a request.
func NewODFareAPIControllerAPIControllerGetParamsWithContext(ctx context.Context) *ODFareAPIControllerAPIControllerGetParams {
	return &ODFareAPIControllerAPIControllerGetParams{
		Context: ctx,
	}
}

// NewODFareAPIControllerAPIControllerGetParamsWithHTTPClient creates a new ODFareAPIControllerAPIControllerGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewODFareAPIControllerAPIControllerGetParamsWithHTTPClient(client *http.Client) *ODFareAPIControllerAPIControllerGetParams {
	return &ODFareAPIControllerAPIControllerGetParams{
		HTTPClient: client,
	}
}

/* ODFareAPIControllerAPIControllerGetParams contains all the parameters to send to the API endpoint
   for the o d fare Api controller Api controller get operation.

   Typically these are written to a http.Request.
*/
type ODFareAPIControllerAPIControllerGetParams struct {

	/* DollarFormat.

	   指定來源格式
	*/
	DollarFormat string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the o d fare Api controller Api controller get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ODFareAPIControllerAPIControllerGetParams) WithDefaults() *ODFareAPIControllerAPIControllerGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the o d fare Api controller Api controller get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ODFareAPIControllerAPIControllerGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the o d fare Api controller Api controller get params
func (o *ODFareAPIControllerAPIControllerGetParams) WithTimeout(timeout time.Duration) *ODFareAPIControllerAPIControllerGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the o d fare Api controller Api controller get params
func (o *ODFareAPIControllerAPIControllerGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the o d fare Api controller Api controller get params
func (o *ODFareAPIControllerAPIControllerGetParams) WithContext(ctx context.Context) *ODFareAPIControllerAPIControllerGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the o d fare Api controller Api controller get params
func (o *ODFareAPIControllerAPIControllerGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the o d fare Api controller Api controller get params
func (o *ODFareAPIControllerAPIControllerGetParams) WithHTTPClient(client *http.Client) *ODFareAPIControllerAPIControllerGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the o d fare Api controller Api controller get params
func (o *ODFareAPIControllerAPIControllerGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFormat adds the dollarFormat to the o d fare Api controller Api controller get params
func (o *ODFareAPIControllerAPIControllerGetParams) WithDollarFormat(dollarFormat string) *ODFareAPIControllerAPIControllerGetParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the o d fare Api controller Api controller get params
func (o *ODFareAPIControllerAPIControllerGetParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WriteToRequest writes these params to a swagger request
func (o *ODFareAPIControllerAPIControllerGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
