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

// NewCityBusAPIRouteFare1Params creates a new CityBusAPIRouteFare1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCityBusAPIRouteFare1Params() *CityBusAPIRouteFare1Params {
	return &CityBusAPIRouteFare1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCityBusAPIRouteFare1ParamsWithTimeout creates a new CityBusAPIRouteFare1Params object
// with the ability to set a timeout on a request.
func NewCityBusAPIRouteFare1ParamsWithTimeout(timeout time.Duration) *CityBusAPIRouteFare1Params {
	return &CityBusAPIRouteFare1Params{
		timeout: timeout,
	}
}

// NewCityBusAPIRouteFare1ParamsWithContext creates a new CityBusAPIRouteFare1Params object
// with the ability to set a context for a request.
func NewCityBusAPIRouteFare1ParamsWithContext(ctx context.Context) *CityBusAPIRouteFare1Params {
	return &CityBusAPIRouteFare1Params{
		Context: ctx,
	}
}

// NewCityBusAPIRouteFare1ParamsWithHTTPClient creates a new CityBusAPIRouteFare1Params object
// with the ability to set a custom HTTPClient for a request.
func NewCityBusAPIRouteFare1ParamsWithHTTPClient(client *http.Client) *CityBusAPIRouteFare1Params {
	return &CityBusAPIRouteFare1Params{
		HTTPClient: client,
	}
}

/* CityBusAPIRouteFare1Params contains all the parameters to send to the API endpoint
   for the city bus Api route fare 1 operation.

   Typically these are written to a http.Request.
*/
type CityBusAPIRouteFare1Params struct {

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

// WithDefaults hydrates default values in the city bus Api route fare 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CityBusAPIRouteFare1Params) WithDefaults() *CityBusAPIRouteFare1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the city bus Api route fare 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CityBusAPIRouteFare1Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := CityBusAPIRouteFare1Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) WithTimeout(timeout time.Duration) *CityBusAPIRouteFare1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) WithContext(ctx context.Context) *CityBusAPIRouteFare1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) WithHTTPClient(client *http.Client) *CityBusAPIRouteFare1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) WithDollarFilter(dollarFilter *string) *CityBusAPIRouteFare1Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) WithDollarFormat(dollarFormat string) *CityBusAPIRouteFare1Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) WithDollarOrderby(dollarOrderby *string) *CityBusAPIRouteFare1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) WithDollarSelect(dollarSelect *string) *CityBusAPIRouteFare1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) WithDollarSkip(dollarSkip *string) *CityBusAPIRouteFare1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) WithDollarTop(dollarTop *int64) *CityBusAPIRouteFare1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithCity adds the city to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) WithCity(city string) *CityBusAPIRouteFare1Params {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) SetCity(city string) {
	o.City = city
}

// WithRouteName adds the routeName to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) WithRouteName(routeName string) *CityBusAPIRouteFare1Params {
	o.SetRouteName(routeName)
	return o
}

// SetRouteName adds the routeName to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) SetRouteName(routeName string) {
	o.RouteName = routeName
}

// WithHealth adds the health to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) WithHealth(health *string) *CityBusAPIRouteFare1Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the city bus Api route fare 1 params
func (o *CityBusAPIRouteFare1Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *CityBusAPIRouteFare1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
