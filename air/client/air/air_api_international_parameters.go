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

// NewAirAPIInternationalParams creates a new AirAPIInternationalParams object
// with the default values initialized.
func NewAirAPIInternationalParams() *AirAPIInternationalParams {
	var (
		dollarTopDefault = string("30")
	)
	return &AirAPIInternationalParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewAirAPIInternationalParamsWithTimeout creates a new AirAPIInternationalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAirAPIInternationalParamsWithTimeout(timeout time.Duration) *AirAPIInternationalParams {
	var (
		dollarTopDefault = string("30")
	)
	return &AirAPIInternationalParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewAirAPIInternationalParamsWithContext creates a new AirAPIInternationalParams object
// with the default values initialized, and the ability to set a context for a request
func NewAirAPIInternationalParamsWithContext(ctx context.Context) *AirAPIInternationalParams {
	var (
		dollarTopDefault = string("30")
	)
	return &AirAPIInternationalParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewAirAPIInternationalParamsWithHTTPClient creates a new AirAPIInternationalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAirAPIInternationalParamsWithHTTPClient(client *http.Client) *AirAPIInternationalParams {
	var (
		dollarTopDefault = string("30")
	)
	return &AirAPIInternationalParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*AirAPIInternationalParams contains all the parameters to send to the API endpoint
for the air Api international operation typically these are written to a http.Request
*/
type AirAPIInternationalParams struct {

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

// WithTimeout adds the timeout to the air Api international params
func (o *AirAPIInternationalParams) WithTimeout(timeout time.Duration) *AirAPIInternationalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the air Api international params
func (o *AirAPIInternationalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the air Api international params
func (o *AirAPIInternationalParams) WithContext(ctx context.Context) *AirAPIInternationalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the air Api international params
func (o *AirAPIInternationalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the air Api international params
func (o *AirAPIInternationalParams) WithHTTPClient(client *http.Client) *AirAPIInternationalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the air Api international params
func (o *AirAPIInternationalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the air Api international params
func (o *AirAPIInternationalParams) WithDollarFilter(dollarFilter *string) *AirAPIInternationalParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the air Api international params
func (o *AirAPIInternationalParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the air Api international params
func (o *AirAPIInternationalParams) WithDollarFormat(dollarFormat string) *AirAPIInternationalParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the air Api international params
func (o *AirAPIInternationalParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the air Api international params
func (o *AirAPIInternationalParams) WithDollarOrderby(dollarOrderby *string) *AirAPIInternationalParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the air Api international params
func (o *AirAPIInternationalParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the air Api international params
func (o *AirAPIInternationalParams) WithDollarSelect(dollarSelect *string) *AirAPIInternationalParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the air Api international params
func (o *AirAPIInternationalParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the air Api international params
func (o *AirAPIInternationalParams) WithDollarSkip(dollarSkip *string) *AirAPIInternationalParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the air Api international params
func (o *AirAPIInternationalParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the air Api international params
func (o *AirAPIInternationalParams) WithDollarTop(dollarTop *string) *AirAPIInternationalParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the air Api international params
func (o *AirAPIInternationalParams) SetDollarTop(dollarTop *string) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *AirAPIInternationalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
