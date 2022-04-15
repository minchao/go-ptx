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

// NewTaiwanTripBusAPIEstimatedTimeOfArrival22621Params creates a new TaiwanTripBusAPIEstimatedTimeOfArrival22621Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTaiwanTripBusAPIEstimatedTimeOfArrival22621Params() *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params {
	return &TaiwanTripBusAPIEstimatedTimeOfArrival22621Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewTaiwanTripBusAPIEstimatedTimeOfArrival22621ParamsWithTimeout creates a new TaiwanTripBusAPIEstimatedTimeOfArrival22621Params object
// with the ability to set a timeout on a request.
func NewTaiwanTripBusAPIEstimatedTimeOfArrival22621ParamsWithTimeout(timeout time.Duration) *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params {
	return &TaiwanTripBusAPIEstimatedTimeOfArrival22621Params{
		timeout: timeout,
	}
}

// NewTaiwanTripBusAPIEstimatedTimeOfArrival22621ParamsWithContext creates a new TaiwanTripBusAPIEstimatedTimeOfArrival22621Params object
// with the ability to set a context for a request.
func NewTaiwanTripBusAPIEstimatedTimeOfArrival22621ParamsWithContext(ctx context.Context) *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params {
	return &TaiwanTripBusAPIEstimatedTimeOfArrival22621Params{
		Context: ctx,
	}
}

// NewTaiwanTripBusAPIEstimatedTimeOfArrival22621ParamsWithHTTPClient creates a new TaiwanTripBusAPIEstimatedTimeOfArrival22621Params object
// with the ability to set a custom HTTPClient for a request.
func NewTaiwanTripBusAPIEstimatedTimeOfArrival22621ParamsWithHTTPClient(client *http.Client) *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params {
	return &TaiwanTripBusAPIEstimatedTimeOfArrival22621Params{
		HTTPClient: client,
	}
}

/* TaiwanTripBusAPIEstimatedTimeOfArrival22621Params contains all the parameters to send to the API endpoint
   for the taiwan trip bus Api estimated time of arrival 2262 1 operation.

   Typically these are written to a http.Request.
*/
type TaiwanTripBusAPIEstimatedTimeOfArrival22621Params struct {

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

	/* Health.

	   加入參數'?health=true'即可查詢此API服務的健康狀態
	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the taiwan trip bus Api estimated time of arrival 2262 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) WithDefaults() *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the taiwan trip bus Api estimated time of arrival 2262 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := TaiwanTripBusAPIEstimatedTimeOfArrival22621Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) WithTimeout(timeout time.Duration) *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) WithContext(ctx context.Context) *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) WithHTTPClient(client *http.Client) *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) WithDollarFilter(dollarFilter *string) *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) WithDollarFormat(dollarFormat string) *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) WithDollarOrderby(dollarOrderby *string) *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) WithDollarSelect(dollarSelect *string) *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) WithDollarSkip(dollarSkip *string) *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) WithDollarTop(dollarTop *int64) *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithTaiwanTripName adds the taiwanTripName to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) WithTaiwanTripName(taiwanTripName string) *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params {
	o.SetTaiwanTripName(taiwanTripName)
	return o
}

// SetTaiwanTripName adds the taiwanTripName to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) SetTaiwanTripName(taiwanTripName string) {
	o.TaiwanTripName = taiwanTripName
}

// WithHealth adds the health to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) WithHealth(health *string) *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the taiwan trip bus Api estimated time of arrival 2262 1 params
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival22621Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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