// Code generated by go-swagger; DO NOT EDIT.

package t_h_s_r

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

// NewTHSRAPIODFareParams creates a new THSRAPIODFareParams object
// with the default values initialized.
func NewTHSRAPIODFareParams() *THSRAPIODFareParams {
	var (
		dollarTopDefault = string("30")
	)
	return &THSRAPIODFareParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTHSRAPIODFareParamsWithTimeout creates a new THSRAPIODFareParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTHSRAPIODFareParamsWithTimeout(timeout time.Duration) *THSRAPIODFareParams {
	var (
		dollarTopDefault = string("30")
	)
	return &THSRAPIODFareParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTHSRAPIODFareParamsWithContext creates a new THSRAPIODFareParams object
// with the default values initialized, and the ability to set a context for a request
func NewTHSRAPIODFareParamsWithContext(ctx context.Context) *THSRAPIODFareParams {
	var (
		dollarTopDefault = string("30")
	)
	return &THSRAPIODFareParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTHSRAPIODFareParamsWithHTTPClient creates a new THSRAPIODFareParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTHSRAPIODFareParamsWithHTTPClient(client *http.Client) *THSRAPIODFareParams {
	var (
		dollarTopDefault = string("30")
	)
	return &THSRAPIODFareParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*THSRAPIODFareParams contains all the parameters to send to the API endpoint
for the t h s r Api o d fare operation typically these are written to a http.Request
*/
type THSRAPIODFareParams struct {

	/*DollarFilter
	  過濾

	*/
	DollarFilter *string
	/*DollarFormat
	  指定來源格式

	*/
	DollarFormat string
	/*DollarOrderby
	  排序

	*/
	DollarOrderby *string
	/*DollarSelect
	  挑選

	*/
	DollarSelect *string
	/*DollarSkip
	  跳過前幾筆

	*/
	DollarSkip *string
	/*DollarTop
	  取前幾筆

	*/
	DollarTop *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) WithTimeout(timeout time.Duration) *THSRAPIODFareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) WithContext(ctx context.Context) *THSRAPIODFareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) WithHTTPClient(client *http.Client) *THSRAPIODFareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) WithDollarFilter(dollarFilter *string) *THSRAPIODFareParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) WithDollarFormat(dollarFormat string) *THSRAPIODFareParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) WithDollarOrderby(dollarOrderby *string) *THSRAPIODFareParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) WithDollarSelect(dollarSelect *string) *THSRAPIODFareParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) WithDollarSkip(dollarSkip *string) *THSRAPIODFareParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) WithDollarTop(dollarTop *string) *THSRAPIODFareParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the t h s r Api o d fare params
func (o *THSRAPIODFareParams) SetDollarTop(dollarTop *string) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *THSRAPIODFareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarFilter != nil {

		// query param $filter
		var qrDollarFilter string
		if o.DollarFilter != nil {
			qrDollarFilter = *o.DollarFilter
		}
		qDollarFilter := qrDollarFilter
		if qDollarFilter != "" {
			if err := r.SetQueryParam("$filter", qDollarFilter); err != nil {
				return err
			}
		}

	}

	// query param $format
	qrDollarFormat := o.DollarFormat
	qDollarFormat := qrDollarFormat
	if qDollarFormat != "" {
		if err := r.SetQueryParam("$format", qDollarFormat); err != nil {
			return err
		}
	}

	if o.DollarOrderby != nil {

		// query param $orderby
		var qrDollarOrderby string
		if o.DollarOrderby != nil {
			qrDollarOrderby = *o.DollarOrderby
		}
		qDollarOrderby := qrDollarOrderby
		if qDollarOrderby != "" {
			if err := r.SetQueryParam("$orderby", qDollarOrderby); err != nil {
				return err
			}
		}

	}

	if o.DollarSelect != nil {

		// query param $select
		var qrDollarSelect string
		if o.DollarSelect != nil {
			qrDollarSelect = *o.DollarSelect
		}
		qDollarSelect := qrDollarSelect
		if qDollarSelect != "" {
			if err := r.SetQueryParam("$select", qDollarSelect); err != nil {
				return err
			}
		}

	}

	if o.DollarSkip != nil {

		// query param $skip
		var qrDollarSkip string
		if o.DollarSkip != nil {
			qrDollarSkip = *o.DollarSkip
		}
		qDollarSkip := qrDollarSkip
		if qDollarSkip != "" {
			if err := r.SetQueryParam("$skip", qDollarSkip); err != nil {
				return err
			}
		}

	}

	if o.DollarTop != nil {

		// query param $top
		var qrDollarTop string
		if o.DollarTop != nil {
			qrDollarTop = *o.DollarTop
		}
		qDollarTop := qrDollarTop
		if qDollarTop != "" {
			if err := r.SetQueryParam("$top", qDollarTop); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
