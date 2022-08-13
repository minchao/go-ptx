// Code generated by go-swagger; DO NOT EDIT.

package bus_advanced_by_operator

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

// NewCityBusAPIRealTimeByFrequencyByOperator2856Params creates a new CityBusAPIRealTimeByFrequencyByOperator2856Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCityBusAPIRealTimeByFrequencyByOperator2856Params() *CityBusAPIRealTimeByFrequencyByOperator2856Params {
	return &CityBusAPIRealTimeByFrequencyByOperator2856Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCityBusAPIRealTimeByFrequencyByOperator2856ParamsWithTimeout creates a new CityBusAPIRealTimeByFrequencyByOperator2856Params object
// with the ability to set a timeout on a request.
func NewCityBusAPIRealTimeByFrequencyByOperator2856ParamsWithTimeout(timeout time.Duration) *CityBusAPIRealTimeByFrequencyByOperator2856Params {
	return &CityBusAPIRealTimeByFrequencyByOperator2856Params{
		timeout: timeout,
	}
}

// NewCityBusAPIRealTimeByFrequencyByOperator2856ParamsWithContext creates a new CityBusAPIRealTimeByFrequencyByOperator2856Params object
// with the ability to set a context for a request.
func NewCityBusAPIRealTimeByFrequencyByOperator2856ParamsWithContext(ctx context.Context) *CityBusAPIRealTimeByFrequencyByOperator2856Params {
	return &CityBusAPIRealTimeByFrequencyByOperator2856Params{
		Context: ctx,
	}
}

// NewCityBusAPIRealTimeByFrequencyByOperator2856ParamsWithHTTPClient creates a new CityBusAPIRealTimeByFrequencyByOperator2856Params object
// with the ability to set a custom HTTPClient for a request.
func NewCityBusAPIRealTimeByFrequencyByOperator2856ParamsWithHTTPClient(client *http.Client) *CityBusAPIRealTimeByFrequencyByOperator2856Params {
	return &CityBusAPIRealTimeByFrequencyByOperator2856Params{
		HTTPClient: client,
	}
}

/* CityBusAPIRealTimeByFrequencyByOperator2856Params contains all the parameters to send to the API endpoint
   for the city bus Api real time by frequency by operator 2856 operation.

   Typically these are written to a http.Request.
*/
type CityBusAPIRealTimeByFrequencyByOperator2856Params struct {

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

	/* City.

	   欲查詢縣市
	*/
	City string

	/* OperatorNo.

	   欲查詢營運業者編號
	*/
	OperatorNo string

	/* Health.

	   加入參數'?health=true'即可查詢此API服務的健康狀態
	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the city bus Api real time by frequency by operator 2856 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) WithDefaults() *CityBusAPIRealTimeByFrequencyByOperator2856Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the city bus Api real time by frequency by operator 2856 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := CityBusAPIRealTimeByFrequencyByOperator2856Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) WithTimeout(timeout time.Duration) *CityBusAPIRealTimeByFrequencyByOperator2856Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) WithContext(ctx context.Context) *CityBusAPIRealTimeByFrequencyByOperator2856Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) WithHTTPClient(client *http.Client) *CityBusAPIRealTimeByFrequencyByOperator2856Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) WithDollarFilter(dollarFilter *string) *CityBusAPIRealTimeByFrequencyByOperator2856Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) WithDollarFormat(dollarFormat string) *CityBusAPIRealTimeByFrequencyByOperator2856Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) WithDollarOrderby(dollarOrderby *string) *CityBusAPIRealTimeByFrequencyByOperator2856Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) WithDollarSelect(dollarSelect *string) *CityBusAPIRealTimeByFrequencyByOperator2856Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) WithDollarSkip(dollarSkip *string) *CityBusAPIRealTimeByFrequencyByOperator2856Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) WithDollarTop(dollarTop *int64) *CityBusAPIRealTimeByFrequencyByOperator2856Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithCity adds the city to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) WithCity(city string) *CityBusAPIRealTimeByFrequencyByOperator2856Params {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) SetCity(city string) {
	o.City = city
}

// WithOperatorNo adds the operatorNo to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) WithOperatorNo(operatorNo string) *CityBusAPIRealTimeByFrequencyByOperator2856Params {
	o.SetOperatorNo(operatorNo)
	return o
}

// SetOperatorNo adds the operatorNo to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) SetOperatorNo(operatorNo string) {
	o.OperatorNo = operatorNo
}

// WithHealth adds the health to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) WithHealth(health *string) *CityBusAPIRealTimeByFrequencyByOperator2856Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the city bus Api real time by frequency by operator 2856 params
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *CityBusAPIRealTimeByFrequencyByOperator2856Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param City
	if err := r.SetPathParam("City", o.City); err != nil {
		return err
	}

	// path param OperatorNo
	if err := r.SetPathParam("OperatorNo", o.OperatorNo); err != nil {
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
