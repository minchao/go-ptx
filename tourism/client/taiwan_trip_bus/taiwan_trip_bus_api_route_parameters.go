// Code generated by go-swagger; DO NOT EDIT.

package taiwan_trip_bus

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

// NewTaiwanTripBusAPIRouteParams creates a new TaiwanTripBusAPIRouteParams object
// with the default values initialized.
func NewTaiwanTripBusAPIRouteParams() *TaiwanTripBusAPIRouteParams {
	var (
		dollarTopDefault = string("30")
	)
	return &TaiwanTripBusAPIRouteParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTaiwanTripBusAPIRouteParamsWithTimeout creates a new TaiwanTripBusAPIRouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTaiwanTripBusAPIRouteParamsWithTimeout(timeout time.Duration) *TaiwanTripBusAPIRouteParams {
	var (
		dollarTopDefault = string("30")
	)
	return &TaiwanTripBusAPIRouteParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTaiwanTripBusAPIRouteParamsWithContext creates a new TaiwanTripBusAPIRouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewTaiwanTripBusAPIRouteParamsWithContext(ctx context.Context) *TaiwanTripBusAPIRouteParams {
	var (
		dollarTopDefault = string("30")
	)
	return &TaiwanTripBusAPIRouteParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTaiwanTripBusAPIRouteParamsWithHTTPClient creates a new TaiwanTripBusAPIRouteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTaiwanTripBusAPIRouteParamsWithHTTPClient(client *http.Client) *TaiwanTripBusAPIRouteParams {
	var (
		dollarTopDefault = string("30")
	)
	return &TaiwanTripBusAPIRouteParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*TaiwanTripBusAPIRouteParams contains all the parameters to send to the API endpoint
for the taiwan trip bus Api route operation typically these are written to a http.Request
*/
type TaiwanTripBusAPIRouteParams struct {

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

// WithTimeout adds the timeout to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) WithTimeout(timeout time.Duration) *TaiwanTripBusAPIRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) WithContext(ctx context.Context) *TaiwanTripBusAPIRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) WithHTTPClient(client *http.Client) *TaiwanTripBusAPIRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) WithDollarFilter(dollarFilter *string) *TaiwanTripBusAPIRouteParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) WithDollarFormat(dollarFormat string) *TaiwanTripBusAPIRouteParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) WithDollarOrderby(dollarOrderby *string) *TaiwanTripBusAPIRouteParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) WithDollarSelect(dollarSelect *string) *TaiwanTripBusAPIRouteParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) WithDollarSkip(dollarSkip *string) *TaiwanTripBusAPIRouteParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) WithDollarTop(dollarTop *string) *TaiwanTripBusAPIRouteParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the taiwan trip bus Api route params
func (o *TaiwanTripBusAPIRouteParams) SetDollarTop(dollarTop *string) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *TaiwanTripBusAPIRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
