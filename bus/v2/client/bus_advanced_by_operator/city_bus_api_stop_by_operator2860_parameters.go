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

// NewCityBusAPIStopByOperator2860Params creates a new CityBusAPIStopByOperator2860Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCityBusAPIStopByOperator2860Params() *CityBusAPIStopByOperator2860Params {
	return &CityBusAPIStopByOperator2860Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCityBusAPIStopByOperator2860ParamsWithTimeout creates a new CityBusAPIStopByOperator2860Params object
// with the ability to set a timeout on a request.
func NewCityBusAPIStopByOperator2860ParamsWithTimeout(timeout time.Duration) *CityBusAPIStopByOperator2860Params {
	return &CityBusAPIStopByOperator2860Params{
		timeout: timeout,
	}
}

// NewCityBusAPIStopByOperator2860ParamsWithContext creates a new CityBusAPIStopByOperator2860Params object
// with the ability to set a context for a request.
func NewCityBusAPIStopByOperator2860ParamsWithContext(ctx context.Context) *CityBusAPIStopByOperator2860Params {
	return &CityBusAPIStopByOperator2860Params{
		Context: ctx,
	}
}

// NewCityBusAPIStopByOperator2860ParamsWithHTTPClient creates a new CityBusAPIStopByOperator2860Params object
// with the ability to set a custom HTTPClient for a request.
func NewCityBusAPIStopByOperator2860ParamsWithHTTPClient(client *http.Client) *CityBusAPIStopByOperator2860Params {
	return &CityBusAPIStopByOperator2860Params{
		HTTPClient: client,
	}
}

/* CityBusAPIStopByOperator2860Params contains all the parameters to send to the API endpoint
   for the city bus Api stop by operator 2860 operation.

   Typically these are written to a http.Request.
*/
type CityBusAPIStopByOperator2860Params struct {

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

	/* DollarSpatialFilter.

	   空間過濾，語法為nearby({Lat},{Lon},{DistanceInMeters})，例如nearby(25.047675, 121.517055, 100)
	*/
	DollarSpatialFilter *string

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

// WithDefaults hydrates default values in the city bus Api stop by operator 2860 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CityBusAPIStopByOperator2860Params) WithDefaults() *CityBusAPIStopByOperator2860Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the city bus Api stop by operator 2860 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CityBusAPIStopByOperator2860Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := CityBusAPIStopByOperator2860Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) WithTimeout(timeout time.Duration) *CityBusAPIStopByOperator2860Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) WithContext(ctx context.Context) *CityBusAPIStopByOperator2860Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) WithHTTPClient(client *http.Client) *CityBusAPIStopByOperator2860Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) WithDollarFilter(dollarFilter *string) *CityBusAPIStopByOperator2860Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) WithDollarFormat(dollarFormat string) *CityBusAPIStopByOperator2860Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) WithDollarOrderby(dollarOrderby *string) *CityBusAPIStopByOperator2860Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) WithDollarSelect(dollarSelect *string) *CityBusAPIStopByOperator2860Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) WithDollarSkip(dollarSkip *string) *CityBusAPIStopByOperator2860Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarSpatialFilter adds the dollarSpatialFilter to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) WithDollarSpatialFilter(dollarSpatialFilter *string) *CityBusAPIStopByOperator2860Params {
	o.SetDollarSpatialFilter(dollarSpatialFilter)
	return o
}

// SetDollarSpatialFilter adds the dollarSpatialFilter to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) SetDollarSpatialFilter(dollarSpatialFilter *string) {
	o.DollarSpatialFilter = dollarSpatialFilter
}

// WithDollarTop adds the dollarTop to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) WithDollarTop(dollarTop *int64) *CityBusAPIStopByOperator2860Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithCity adds the city to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) WithCity(city string) *CityBusAPIStopByOperator2860Params {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) SetCity(city string) {
	o.City = city
}

// WithOperatorNo adds the operatorNo to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) WithOperatorNo(operatorNo string) *CityBusAPIStopByOperator2860Params {
	o.SetOperatorNo(operatorNo)
	return o
}

// SetOperatorNo adds the operatorNo to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) SetOperatorNo(operatorNo string) {
	o.OperatorNo = operatorNo
}

// WithHealth adds the health to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) WithHealth(health *string) *CityBusAPIStopByOperator2860Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the city bus Api stop by operator 2860 params
func (o *CityBusAPIStopByOperator2860Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *CityBusAPIStopByOperator2860Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DollarSpatialFilter != nil {

		// query param $spatialFilter
		var qrDollarSpatialFilter string

		if o.DollarSpatialFilter != nil {
			qrDollarSpatialFilter = *o.DollarSpatialFilter
		}
		qDollarSpatialFilter := qrDollarSpatialFilter
		if qDollarSpatialFilter != "" {

			if err := r.SetQueryParam("$spatialFilter", qDollarSpatialFilter); err != nil {
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