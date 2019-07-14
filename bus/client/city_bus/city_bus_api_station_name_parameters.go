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

	strfmt "github.com/go-openapi/strfmt"
)

// NewCityBusAPIStationNameParams creates a new CityBusAPIStationNameParams object
// with the default values initialized.
func NewCityBusAPIStationNameParams() *CityBusAPIStationNameParams {
	var (
		dollarTopDefault = string("30")
	)
	return &CityBusAPIStationNameParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCityBusAPIStationNameParamsWithTimeout creates a new CityBusAPIStationNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCityBusAPIStationNameParamsWithTimeout(timeout time.Duration) *CityBusAPIStationNameParams {
	var (
		dollarTopDefault = string("30")
	)
	return &CityBusAPIStationNameParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewCityBusAPIStationNameParamsWithContext creates a new CityBusAPIStationNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewCityBusAPIStationNameParamsWithContext(ctx context.Context) *CityBusAPIStationNameParams {
	var (
		dollarTopDefault = string("30")
	)
	return &CityBusAPIStationNameParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewCityBusAPIStationNameParamsWithHTTPClient creates a new CityBusAPIStationNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCityBusAPIStationNameParamsWithHTTPClient(client *http.Client) *CityBusAPIStationNameParams {
	var (
		dollarTopDefault = string("30")
	)
	return &CityBusAPIStationNameParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*CityBusAPIStationNameParams contains all the parameters to send to the API endpoint
for the city bus Api station name operation typically these are written to a http.Request
*/
type CityBusAPIStationNameParams struct {

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
	DollarTop *string
	/*City
	  欲查詢縣市

	*/
	City string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the city bus Api station name params
func (o *CityBusAPIStationNameParams) WithTimeout(timeout time.Duration) *CityBusAPIStationNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the city bus Api station name params
func (o *CityBusAPIStationNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the city bus Api station name params
func (o *CityBusAPIStationNameParams) WithContext(ctx context.Context) *CityBusAPIStationNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the city bus Api station name params
func (o *CityBusAPIStationNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the city bus Api station name params
func (o *CityBusAPIStationNameParams) WithHTTPClient(client *http.Client) *CityBusAPIStationNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the city bus Api station name params
func (o *CityBusAPIStationNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the city bus Api station name params
func (o *CityBusAPIStationNameParams) WithDollarFilter(dollarFilter *string) *CityBusAPIStationNameParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the city bus Api station name params
func (o *CityBusAPIStationNameParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the city bus Api station name params
func (o *CityBusAPIStationNameParams) WithDollarFormat(dollarFormat string) *CityBusAPIStationNameParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the city bus Api station name params
func (o *CityBusAPIStationNameParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the city bus Api station name params
func (o *CityBusAPIStationNameParams) WithDollarOrderby(dollarOrderby *string) *CityBusAPIStationNameParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the city bus Api station name params
func (o *CityBusAPIStationNameParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the city bus Api station name params
func (o *CityBusAPIStationNameParams) WithDollarSelect(dollarSelect *string) *CityBusAPIStationNameParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the city bus Api station name params
func (o *CityBusAPIStationNameParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the city bus Api station name params
func (o *CityBusAPIStationNameParams) WithDollarSkip(dollarSkip *string) *CityBusAPIStationNameParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the city bus Api station name params
func (o *CityBusAPIStationNameParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarSpatialFilter adds the dollarSpatialFilter to the city bus Api station name params
func (o *CityBusAPIStationNameParams) WithDollarSpatialFilter(dollarSpatialFilter *string) *CityBusAPIStationNameParams {
	o.SetDollarSpatialFilter(dollarSpatialFilter)
	return o
}

// SetDollarSpatialFilter adds the dollarSpatialFilter to the city bus Api station name params
func (o *CityBusAPIStationNameParams) SetDollarSpatialFilter(dollarSpatialFilter *string) {
	o.DollarSpatialFilter = dollarSpatialFilter
}

// WithDollarTop adds the dollarTop to the city bus Api station name params
func (o *CityBusAPIStationNameParams) WithDollarTop(dollarTop *string) *CityBusAPIStationNameParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the city bus Api station name params
func (o *CityBusAPIStationNameParams) SetDollarTop(dollarTop *string) {
	o.DollarTop = dollarTop
}

// WithCity adds the city to the city bus Api station name params
func (o *CityBusAPIStationNameParams) WithCity(city string) *CityBusAPIStationNameParams {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the city bus Api station name params
func (o *CityBusAPIStationNameParams) SetCity(city string) {
	o.City = city
}

// WriteToRequest writes these params to a swagger request
func (o *CityBusAPIStationNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param City
	if err := r.SetPathParam("City", o.City); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
