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

// NewCityBusAPIShapeParams creates a new CityBusAPIShapeParams object
// with the default values initialized.
func NewCityBusAPIShapeParams() *CityBusAPIShapeParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &CityBusAPIShapeParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCityBusAPIShapeParamsWithTimeout creates a new CityBusAPIShapeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCityBusAPIShapeParamsWithTimeout(timeout time.Duration) *CityBusAPIShapeParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &CityBusAPIShapeParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewCityBusAPIShapeParamsWithContext creates a new CityBusAPIShapeParams object
// with the default values initialized, and the ability to set a context for a request
func NewCityBusAPIShapeParamsWithContext(ctx context.Context) *CityBusAPIShapeParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &CityBusAPIShapeParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewCityBusAPIShapeParamsWithHTTPClient creates a new CityBusAPIShapeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCityBusAPIShapeParamsWithHTTPClient(client *http.Client) *CityBusAPIShapeParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &CityBusAPIShapeParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*CityBusAPIShapeParams contains all the parameters to send to the API endpoint
for the city bus Api shape operation typically these are written to a http.Request
*/
type CityBusAPIShapeParams struct {

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
	  縣市

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

// WithTimeout adds the timeout to the city bus Api shape params
func (o *CityBusAPIShapeParams) WithTimeout(timeout time.Duration) *CityBusAPIShapeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the city bus Api shape params
func (o *CityBusAPIShapeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the city bus Api shape params
func (o *CityBusAPIShapeParams) WithContext(ctx context.Context) *CityBusAPIShapeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the city bus Api shape params
func (o *CityBusAPIShapeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the city bus Api shape params
func (o *CityBusAPIShapeParams) WithHTTPClient(client *http.Client) *CityBusAPIShapeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the city bus Api shape params
func (o *CityBusAPIShapeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the city bus Api shape params
func (o *CityBusAPIShapeParams) WithDollarCount(dollarCount *string) *CityBusAPIShapeParams {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the city bus Api shape params
func (o *CityBusAPIShapeParams) SetDollarCount(dollarCount *string) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the city bus Api shape params
func (o *CityBusAPIShapeParams) WithDollarFilter(dollarFilter *string) *CityBusAPIShapeParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the city bus Api shape params
func (o *CityBusAPIShapeParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the city bus Api shape params
func (o *CityBusAPIShapeParams) WithDollarFormat(dollarFormat string) *CityBusAPIShapeParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the city bus Api shape params
func (o *CityBusAPIShapeParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the city bus Api shape params
func (o *CityBusAPIShapeParams) WithDollarOrderby(dollarOrderby *string) *CityBusAPIShapeParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the city bus Api shape params
func (o *CityBusAPIShapeParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the city bus Api shape params
func (o *CityBusAPIShapeParams) WithDollarSelect(dollarSelect *string) *CityBusAPIShapeParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the city bus Api shape params
func (o *CityBusAPIShapeParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the city bus Api shape params
func (o *CityBusAPIShapeParams) WithDollarSkip(dollarSkip *string) *CityBusAPIShapeParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the city bus Api shape params
func (o *CityBusAPIShapeParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the city bus Api shape params
func (o *CityBusAPIShapeParams) WithDollarTop(dollarTop *int64) *CityBusAPIShapeParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the city bus Api shape params
func (o *CityBusAPIShapeParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithCity adds the city to the city bus Api shape params
func (o *CityBusAPIShapeParams) WithCity(city string) *CityBusAPIShapeParams {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the city bus Api shape params
func (o *CityBusAPIShapeParams) SetCity(city string) {
	o.City = city
}

// WithHealth adds the health to the city bus Api shape params
func (o *CityBusAPIShapeParams) WithHealth(health *string) *CityBusAPIShapeParams {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the city bus Api shape params
func (o *CityBusAPIShapeParams) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *CityBusAPIShapeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
