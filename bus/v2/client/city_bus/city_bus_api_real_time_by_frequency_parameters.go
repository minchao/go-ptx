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

// NewCityBusAPIRealTimeByFrequencyParams creates a new CityBusAPIRealTimeByFrequencyParams object
// with the default values initialized.
func NewCityBusAPIRealTimeByFrequencyParams() *CityBusAPIRealTimeByFrequencyParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &CityBusAPIRealTimeByFrequencyParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCityBusAPIRealTimeByFrequencyParamsWithTimeout creates a new CityBusAPIRealTimeByFrequencyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCityBusAPIRealTimeByFrequencyParamsWithTimeout(timeout time.Duration) *CityBusAPIRealTimeByFrequencyParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &CityBusAPIRealTimeByFrequencyParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewCityBusAPIRealTimeByFrequencyParamsWithContext creates a new CityBusAPIRealTimeByFrequencyParams object
// with the default values initialized, and the ability to set a context for a request
func NewCityBusAPIRealTimeByFrequencyParamsWithContext(ctx context.Context) *CityBusAPIRealTimeByFrequencyParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &CityBusAPIRealTimeByFrequencyParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewCityBusAPIRealTimeByFrequencyParamsWithHTTPClient creates a new CityBusAPIRealTimeByFrequencyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCityBusAPIRealTimeByFrequencyParamsWithHTTPClient(client *http.Client) *CityBusAPIRealTimeByFrequencyParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &CityBusAPIRealTimeByFrequencyParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*CityBusAPIRealTimeByFrequencyParams contains all the parameters to send to the API endpoint
for the city bus Api real time by frequency operation typically these are written to a http.Request
*/
type CityBusAPIRealTimeByFrequencyParams struct {

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
	/*DollarSpatialFilter
	  空間過濾

	*/
	DollarSpatialFilter *string
	/*DollarTop
	  取前幾筆

	*/
	DollarTop *int64
	/*City
	  欲查詢縣市

	*/
	City string
	/*Health
	  加入參數'?health=true'即可查詢此API服務的健康狀態

	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) WithTimeout(timeout time.Duration) *CityBusAPIRealTimeByFrequencyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) WithContext(ctx context.Context) *CityBusAPIRealTimeByFrequencyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) WithHTTPClient(client *http.Client) *CityBusAPIRealTimeByFrequencyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) WithDollarFilter(dollarFilter *string) *CityBusAPIRealTimeByFrequencyParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) WithDollarFormat(dollarFormat string) *CityBusAPIRealTimeByFrequencyParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) WithDollarOrderby(dollarOrderby *string) *CityBusAPIRealTimeByFrequencyParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) WithDollarSelect(dollarSelect *string) *CityBusAPIRealTimeByFrequencyParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) WithDollarSkip(dollarSkip *string) *CityBusAPIRealTimeByFrequencyParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarSpatialFilter adds the dollarSpatialFilter to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) WithDollarSpatialFilter(dollarSpatialFilter *string) *CityBusAPIRealTimeByFrequencyParams {
	o.SetDollarSpatialFilter(dollarSpatialFilter)
	return o
}

// SetDollarSpatialFilter adds the dollarSpatialFilter to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) SetDollarSpatialFilter(dollarSpatialFilter *string) {
	o.DollarSpatialFilter = dollarSpatialFilter
}

// WithDollarTop adds the dollarTop to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) WithDollarTop(dollarTop *int64) *CityBusAPIRealTimeByFrequencyParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithCity adds the city to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) WithCity(city string) *CityBusAPIRealTimeByFrequencyParams {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) SetCity(city string) {
	o.City = city
}

// WithHealth adds the health to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) WithHealth(health *string) *CityBusAPIRealTimeByFrequencyParams {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the city bus Api real time by frequency params
func (o *CityBusAPIRealTimeByFrequencyParams) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *CityBusAPIRealTimeByFrequencyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
