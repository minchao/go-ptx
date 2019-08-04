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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCityBusAPIRealTimeNearStop3Params creates a new CityBusAPIRealTimeNearStop3Params object
// with the default values initialized.
func NewCityBusAPIRealTimeNearStop3Params() *CityBusAPIRealTimeNearStop3Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &CityBusAPIRealTimeNearStop3Params{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCityBusAPIRealTimeNearStop3ParamsWithTimeout creates a new CityBusAPIRealTimeNearStop3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCityBusAPIRealTimeNearStop3ParamsWithTimeout(timeout time.Duration) *CityBusAPIRealTimeNearStop3Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &CityBusAPIRealTimeNearStop3Params{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewCityBusAPIRealTimeNearStop3ParamsWithContext creates a new CityBusAPIRealTimeNearStop3Params object
// with the default values initialized, and the ability to set a context for a request
func NewCityBusAPIRealTimeNearStop3ParamsWithContext(ctx context.Context) *CityBusAPIRealTimeNearStop3Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &CityBusAPIRealTimeNearStop3Params{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewCityBusAPIRealTimeNearStop3ParamsWithHTTPClient creates a new CityBusAPIRealTimeNearStop3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCityBusAPIRealTimeNearStop3ParamsWithHTTPClient(client *http.Client) *CityBusAPIRealTimeNearStop3Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &CityBusAPIRealTimeNearStop3Params{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*CityBusAPIRealTimeNearStop3Params contains all the parameters to send to the API endpoint
for the city bus Api real time near stop 3 operation typically these are written to a http.Request
*/
type CityBusAPIRealTimeNearStop3Params struct {

	/*DollarCount
	  查詢數量

	*/
	DollarCount *string
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
	/*City
	  欲查詢縣市

	*/
	City string
	/*RouteName
	  繁體中文路線名稱，如'307'

	*/
	RouteName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) WithTimeout(timeout time.Duration) *CityBusAPIRealTimeNearStop3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) WithContext(ctx context.Context) *CityBusAPIRealTimeNearStop3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) WithHTTPClient(client *http.Client) *CityBusAPIRealTimeNearStop3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) WithDollarCount(dollarCount *string) *CityBusAPIRealTimeNearStop3Params {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) SetDollarCount(dollarCount *string) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) WithDollarFilter(dollarFilter *string) *CityBusAPIRealTimeNearStop3Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) WithDollarFormat(dollarFormat string) *CityBusAPIRealTimeNearStop3Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) WithDollarOrderby(dollarOrderby *string) *CityBusAPIRealTimeNearStop3Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) WithDollarSelect(dollarSelect *string) *CityBusAPIRealTimeNearStop3Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) WithDollarSkip(dollarSkip *string) *CityBusAPIRealTimeNearStop3Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) WithDollarTop(dollarTop *int64) *CityBusAPIRealTimeNearStop3Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithCity adds the city to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) WithCity(city string) *CityBusAPIRealTimeNearStop3Params {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) SetCity(city string) {
	o.City = city
}

// WithRouteName adds the routeName to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) WithRouteName(routeName string) *CityBusAPIRealTimeNearStop3Params {
	o.SetRouteName(routeName)
	return o
}

// SetRouteName adds the routeName to the city bus Api real time near stop 3 params
func (o *CityBusAPIRealTimeNearStop3Params) SetRouteName(routeName string) {
	o.RouteName = routeName
}

// WriteToRequest writes these params to a swagger request
func (o *CityBusAPIRealTimeNearStop3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarCount != nil {

		// query param $count
		var qrDollarCount string
		if o.DollarCount != nil {
			qrDollarCount = *o.DollarCount
		}
		qDollarCount := qrDollarCount
		if qDollarCount != "" {
			if err := r.SetQueryParam("$count", qDollarCount); err != nil {
				return err
			}
		}

	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}