// Code generated by go-swagger; DO NOT EDIT.

package city_bus

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

// NewCityBusAPIRealTimeNearStopUDP20471Params creates a new CityBusAPIRealTimeNearStopUDP20471Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCityBusAPIRealTimeNearStopUDP20471Params() *CityBusAPIRealTimeNearStopUDP20471Params {
	return &CityBusAPIRealTimeNearStopUDP20471Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCityBusAPIRealTimeNearStopUDP20471ParamsWithTimeout creates a new CityBusAPIRealTimeNearStopUDP20471Params object
// with the ability to set a timeout on a request.
func NewCityBusAPIRealTimeNearStopUDP20471ParamsWithTimeout(timeout time.Duration) *CityBusAPIRealTimeNearStopUDP20471Params {
	return &CityBusAPIRealTimeNearStopUDP20471Params{
		timeout: timeout,
	}
}

// NewCityBusAPIRealTimeNearStopUDP20471ParamsWithContext creates a new CityBusAPIRealTimeNearStopUDP20471Params object
// with the ability to set a context for a request.
func NewCityBusAPIRealTimeNearStopUDP20471ParamsWithContext(ctx context.Context) *CityBusAPIRealTimeNearStopUDP20471Params {
	return &CityBusAPIRealTimeNearStopUDP20471Params{
		Context: ctx,
	}
}

// NewCityBusAPIRealTimeNearStopUDP20471ParamsWithHTTPClient creates a new CityBusAPIRealTimeNearStopUDP20471Params object
// with the ability to set a custom HTTPClient for a request.
func NewCityBusAPIRealTimeNearStopUDP20471ParamsWithHTTPClient(client *http.Client) *CityBusAPIRealTimeNearStopUDP20471Params {
	return &CityBusAPIRealTimeNearStopUDP20471Params{
		HTTPClient: client,
	}
}

/* CityBusAPIRealTimeNearStopUDP20471Params contains all the parameters to send to the API endpoint
   for the city bus Api real time near stop UDP 2047 1 operation.

   Typically these are written to a http.Request.
*/
type CityBusAPIRealTimeNearStopUDP20471Params struct {

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

	/* RouteName.

	   繁體中文路線名稱，如'307'
	*/
	RouteName string

	/* Health.

	   加入參數'?health=true'即可查詢此API服務的健康狀態
	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the city bus Api real time near stop UDP 2047 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CityBusAPIRealTimeNearStopUDP20471Params) WithDefaults() *CityBusAPIRealTimeNearStopUDP20471Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the city bus Api real time near stop UDP 2047 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CityBusAPIRealTimeNearStopUDP20471Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := CityBusAPIRealTimeNearStopUDP20471Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) WithTimeout(timeout time.Duration) *CityBusAPIRealTimeNearStopUDP20471Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) WithContext(ctx context.Context) *CityBusAPIRealTimeNearStopUDP20471Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) WithHTTPClient(client *http.Client) *CityBusAPIRealTimeNearStopUDP20471Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) WithDollarFilter(dollarFilter *string) *CityBusAPIRealTimeNearStopUDP20471Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) WithDollarFormat(dollarFormat string) *CityBusAPIRealTimeNearStopUDP20471Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) WithDollarOrderby(dollarOrderby *string) *CityBusAPIRealTimeNearStopUDP20471Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) WithDollarSelect(dollarSelect *string) *CityBusAPIRealTimeNearStopUDP20471Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) WithDollarSkip(dollarSkip *string) *CityBusAPIRealTimeNearStopUDP20471Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) WithDollarTop(dollarTop *int64) *CityBusAPIRealTimeNearStopUDP20471Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithCity adds the city to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) WithCity(city string) *CityBusAPIRealTimeNearStopUDP20471Params {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) SetCity(city string) {
	o.City = city
}

// WithRouteName adds the routeName to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) WithRouteName(routeName string) *CityBusAPIRealTimeNearStopUDP20471Params {
	o.SetRouteName(routeName)
	return o
}

// SetRouteName adds the routeName to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) SetRouteName(routeName string) {
	o.RouteName = routeName
}

// WithHealth adds the health to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) WithHealth(health *string) *CityBusAPIRealTimeNearStopUDP20471Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the city bus Api real time near stop UDP 2047 1 params
func (o *CityBusAPIRealTimeNearStopUDP20471Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *CityBusAPIRealTimeNearStopUDP20471Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param RouteName
	if err := r.SetPathParam("RouteName", o.RouteName); err != nil {
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
