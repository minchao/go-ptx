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

// NewTaiwanTripBusAPIRealTimeNearStop1Params creates a new TaiwanTripBusAPIRealTimeNearStop1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTaiwanTripBusAPIRealTimeNearStop1Params() *TaiwanTripBusAPIRealTimeNearStop1Params {
	return &TaiwanTripBusAPIRealTimeNearStop1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewTaiwanTripBusAPIRealTimeNearStop1ParamsWithTimeout creates a new TaiwanTripBusAPIRealTimeNearStop1Params object
// with the ability to set a timeout on a request.
func NewTaiwanTripBusAPIRealTimeNearStop1ParamsWithTimeout(timeout time.Duration) *TaiwanTripBusAPIRealTimeNearStop1Params {
	return &TaiwanTripBusAPIRealTimeNearStop1Params{
		timeout: timeout,
	}
}

// NewTaiwanTripBusAPIRealTimeNearStop1ParamsWithContext creates a new TaiwanTripBusAPIRealTimeNearStop1Params object
// with the ability to set a context for a request.
func NewTaiwanTripBusAPIRealTimeNearStop1ParamsWithContext(ctx context.Context) *TaiwanTripBusAPIRealTimeNearStop1Params {
	return &TaiwanTripBusAPIRealTimeNearStop1Params{
		Context: ctx,
	}
}

// NewTaiwanTripBusAPIRealTimeNearStop1ParamsWithHTTPClient creates a new TaiwanTripBusAPIRealTimeNearStop1Params object
// with the ability to set a custom HTTPClient for a request.
func NewTaiwanTripBusAPIRealTimeNearStop1ParamsWithHTTPClient(client *http.Client) *TaiwanTripBusAPIRealTimeNearStop1Params {
	return &TaiwanTripBusAPIRealTimeNearStop1Params{
		HTTPClient: client,
	}
}

/* TaiwanTripBusAPIRealTimeNearStop1Params contains all the parameters to send to the API endpoint
   for the taiwan trip bus Api real time near stop 1 operation.

   Typically these are written to a http.Request.
*/
type TaiwanTripBusAPIRealTimeNearStop1Params struct {

	/* DollarFilter.

	   過濾
	*/
	DollarFilter *string

	/* DollarFormat.

	   指定來源格式
	*/
	DollarFormat string

	/* DollarOrderby.

	   排序
	*/
	DollarOrderby *string

	/* DollarSelect.

	   挑選
	*/
	DollarSelect *string

	/* DollarSkip.

	   跳過前幾筆
	*/
	DollarSkip *string

	/* DollarTop.

	   取前幾筆

	   Default: 30
	*/
	DollarTop *int64

	/* TaiwanTripName.

	   台灣好行繁體中文路線名稱，如'黃金福隆線'
	*/
	TaiwanTripName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the taiwan trip bus Api real time near stop 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) WithDefaults() *TaiwanTripBusAPIRealTimeNearStop1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the taiwan trip bus Api real time near stop 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := TaiwanTripBusAPIRealTimeNearStop1Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) WithTimeout(timeout time.Duration) *TaiwanTripBusAPIRealTimeNearStop1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) WithContext(ctx context.Context) *TaiwanTripBusAPIRealTimeNearStop1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) WithHTTPClient(client *http.Client) *TaiwanTripBusAPIRealTimeNearStop1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) WithDollarFilter(dollarFilter *string) *TaiwanTripBusAPIRealTimeNearStop1Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) WithDollarFormat(dollarFormat string) *TaiwanTripBusAPIRealTimeNearStop1Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) WithDollarOrderby(dollarOrderby *string) *TaiwanTripBusAPIRealTimeNearStop1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) WithDollarSelect(dollarSelect *string) *TaiwanTripBusAPIRealTimeNearStop1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) WithDollarSkip(dollarSkip *string) *TaiwanTripBusAPIRealTimeNearStop1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) WithDollarTop(dollarTop *int64) *TaiwanTripBusAPIRealTimeNearStop1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithTaiwanTripName adds the taiwanTripName to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) WithTaiwanTripName(taiwanTripName string) *TaiwanTripBusAPIRealTimeNearStop1Params {
	o.SetTaiwanTripName(taiwanTripName)
	return o
}

// SetTaiwanTripName adds the taiwanTripName to the taiwan trip bus Api real time near stop 1 params
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) SetTaiwanTripName(taiwanTripName string) {
	o.TaiwanTripName = taiwanTripName
}

// WriteToRequest writes these params to a swagger request
func (o *TaiwanTripBusAPIRealTimeNearStop1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
