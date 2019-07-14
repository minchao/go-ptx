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

// NewTaiwanTripBusAPIStopOfRoute1Params creates a new TaiwanTripBusAPIStopOfRoute1Params object
// with the default values initialized.
func NewTaiwanTripBusAPIStopOfRoute1Params() *TaiwanTripBusAPIStopOfRoute1Params {
	var (
		dollarTopDefault = string("30")
	)
	return &TaiwanTripBusAPIStopOfRoute1Params{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTaiwanTripBusAPIStopOfRoute1ParamsWithTimeout creates a new TaiwanTripBusAPIStopOfRoute1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewTaiwanTripBusAPIStopOfRoute1ParamsWithTimeout(timeout time.Duration) *TaiwanTripBusAPIStopOfRoute1Params {
	var (
		dollarTopDefault = string("30")
	)
	return &TaiwanTripBusAPIStopOfRoute1Params{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTaiwanTripBusAPIStopOfRoute1ParamsWithContext creates a new TaiwanTripBusAPIStopOfRoute1Params object
// with the default values initialized, and the ability to set a context for a request
func NewTaiwanTripBusAPIStopOfRoute1ParamsWithContext(ctx context.Context) *TaiwanTripBusAPIStopOfRoute1Params {
	var (
		dollarTopDefault = string("30")
	)
	return &TaiwanTripBusAPIStopOfRoute1Params{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTaiwanTripBusAPIStopOfRoute1ParamsWithHTTPClient creates a new TaiwanTripBusAPIStopOfRoute1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTaiwanTripBusAPIStopOfRoute1ParamsWithHTTPClient(client *http.Client) *TaiwanTripBusAPIStopOfRoute1Params {
	var (
		dollarTopDefault = string("30")
	)
	return &TaiwanTripBusAPIStopOfRoute1Params{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*TaiwanTripBusAPIStopOfRoute1Params contains all the parameters to send to the API endpoint
for the taiwan trip bus Api stop of route 1 operation typically these are written to a http.Request
*/
type TaiwanTripBusAPIStopOfRoute1Params struct {

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
	/*TaiwanTripName
	  台灣好行繁體中文路線名稱，如'黃金福隆線'

	*/
	TaiwanTripName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) WithTimeout(timeout time.Duration) *TaiwanTripBusAPIStopOfRoute1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) WithContext(ctx context.Context) *TaiwanTripBusAPIStopOfRoute1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) WithHTTPClient(client *http.Client) *TaiwanTripBusAPIStopOfRoute1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) WithDollarFilter(dollarFilter *string) *TaiwanTripBusAPIStopOfRoute1Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) WithDollarFormat(dollarFormat string) *TaiwanTripBusAPIStopOfRoute1Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) WithDollarOrderby(dollarOrderby *string) *TaiwanTripBusAPIStopOfRoute1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) WithDollarSelect(dollarSelect *string) *TaiwanTripBusAPIStopOfRoute1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) WithDollarSkip(dollarSkip *string) *TaiwanTripBusAPIStopOfRoute1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) WithDollarTop(dollarTop *string) *TaiwanTripBusAPIStopOfRoute1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) SetDollarTop(dollarTop *string) {
	o.DollarTop = dollarTop
}

// WithTaiwanTripName adds the taiwanTripName to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) WithTaiwanTripName(taiwanTripName string) *TaiwanTripBusAPIStopOfRoute1Params {
	o.SetTaiwanTripName(taiwanTripName)
	return o
}

// SetTaiwanTripName adds the taiwanTripName to the taiwan trip bus Api stop of route 1 params
func (o *TaiwanTripBusAPIStopOfRoute1Params) SetTaiwanTripName(taiwanTripName string) {
	o.TaiwanTripName = taiwanTripName
}

// WriteToRequest writes these params to a swagger request
func (o *TaiwanTripBusAPIStopOfRoute1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param TaiwanTripName
	if err := r.SetPathParam("TaiwanTripName", o.TaiwanTripName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
