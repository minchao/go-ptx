// Code generated by go-swagger; DO NOT EDIT.

package tourism

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

// NewTaiwanTripBusAPISchedule1Params creates a new TaiwanTripBusAPISchedule1Params object
// with the default values initialized.
func NewTaiwanTripBusAPISchedule1Params() *TaiwanTripBusAPISchedule1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &TaiwanTripBusAPISchedule1Params{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTaiwanTripBusAPISchedule1ParamsWithTimeout creates a new TaiwanTripBusAPISchedule1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewTaiwanTripBusAPISchedule1ParamsWithTimeout(timeout time.Duration) *TaiwanTripBusAPISchedule1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &TaiwanTripBusAPISchedule1Params{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTaiwanTripBusAPISchedule1ParamsWithContext creates a new TaiwanTripBusAPISchedule1Params object
// with the default values initialized, and the ability to set a context for a request
func NewTaiwanTripBusAPISchedule1ParamsWithContext(ctx context.Context) *TaiwanTripBusAPISchedule1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &TaiwanTripBusAPISchedule1Params{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTaiwanTripBusAPISchedule1ParamsWithHTTPClient creates a new TaiwanTripBusAPISchedule1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTaiwanTripBusAPISchedule1ParamsWithHTTPClient(client *http.Client) *TaiwanTripBusAPISchedule1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &TaiwanTripBusAPISchedule1Params{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*TaiwanTripBusAPISchedule1Params contains all the parameters to send to the API endpoint
for the taiwan trip bus Api schedule 1 operation typically these are written to a http.Request
*/
type TaiwanTripBusAPISchedule1Params struct {

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
	DollarTop *int64
	/*TaiwanTripName
	  台灣好行繁體中文路線名稱，如'黃金福隆線'

	*/
	TaiwanTripName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) WithTimeout(timeout time.Duration) *TaiwanTripBusAPISchedule1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) WithContext(ctx context.Context) *TaiwanTripBusAPISchedule1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) WithHTTPClient(client *http.Client) *TaiwanTripBusAPISchedule1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) WithDollarFilter(dollarFilter *string) *TaiwanTripBusAPISchedule1Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) WithDollarFormat(dollarFormat string) *TaiwanTripBusAPISchedule1Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) WithDollarOrderby(dollarOrderby *string) *TaiwanTripBusAPISchedule1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) WithDollarSelect(dollarSelect *string) *TaiwanTripBusAPISchedule1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) WithDollarSkip(dollarSkip *string) *TaiwanTripBusAPISchedule1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) WithDollarTop(dollarTop *int64) *TaiwanTripBusAPISchedule1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithTaiwanTripName adds the taiwanTripName to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) WithTaiwanTripName(taiwanTripName string) *TaiwanTripBusAPISchedule1Params {
	o.SetTaiwanTripName(taiwanTripName)
	return o
}

// SetTaiwanTripName adds the taiwanTripName to the taiwan trip bus Api schedule 1 params
func (o *TaiwanTripBusAPISchedule1Params) SetTaiwanTripName(taiwanTripName string) {
	o.TaiwanTripName = taiwanTripName
}

// WriteToRequest writes these params to a swagger request
func (o *TaiwanTripBusAPISchedule1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
		var qrDollarTop int64
		if o.DollarTop != nil {
			qrDollarTop = *o.DollarTop
		}
		qDollarTop := swag.FormatInt64(qrDollarTop)
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
