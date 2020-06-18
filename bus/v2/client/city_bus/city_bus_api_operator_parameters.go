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

// NewCityBusAPIOperatorParams creates a new CityBusAPIOperatorParams object
// with the default values initialized.
func NewCityBusAPIOperatorParams() *CityBusAPIOperatorParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &CityBusAPIOperatorParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCityBusAPIOperatorParamsWithTimeout creates a new CityBusAPIOperatorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCityBusAPIOperatorParamsWithTimeout(timeout time.Duration) *CityBusAPIOperatorParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &CityBusAPIOperatorParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewCityBusAPIOperatorParamsWithContext creates a new CityBusAPIOperatorParams object
// with the default values initialized, and the ability to set a context for a request
func NewCityBusAPIOperatorParamsWithContext(ctx context.Context) *CityBusAPIOperatorParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &CityBusAPIOperatorParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewCityBusAPIOperatorParamsWithHTTPClient creates a new CityBusAPIOperatorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCityBusAPIOperatorParamsWithHTTPClient(client *http.Client) *CityBusAPIOperatorParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &CityBusAPIOperatorParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*CityBusAPIOperatorParams contains all the parameters to send to the API endpoint
for the city bus Api operator operation typically these are written to a http.Request
*/
type CityBusAPIOperatorParams struct {

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
	/*Health
	  加入參數'?health=true'即可查詢此API服務的健康狀態

	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the city bus Api operator params
func (o *CityBusAPIOperatorParams) WithTimeout(timeout time.Duration) *CityBusAPIOperatorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the city bus Api operator params
func (o *CityBusAPIOperatorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the city bus Api operator params
func (o *CityBusAPIOperatorParams) WithContext(ctx context.Context) *CityBusAPIOperatorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the city bus Api operator params
func (o *CityBusAPIOperatorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the city bus Api operator params
func (o *CityBusAPIOperatorParams) WithHTTPClient(client *http.Client) *CityBusAPIOperatorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the city bus Api operator params
func (o *CityBusAPIOperatorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the city bus Api operator params
func (o *CityBusAPIOperatorParams) WithDollarFilter(dollarFilter *string) *CityBusAPIOperatorParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the city bus Api operator params
func (o *CityBusAPIOperatorParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the city bus Api operator params
func (o *CityBusAPIOperatorParams) WithDollarFormat(dollarFormat string) *CityBusAPIOperatorParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the city bus Api operator params
func (o *CityBusAPIOperatorParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the city bus Api operator params
func (o *CityBusAPIOperatorParams) WithDollarOrderby(dollarOrderby *string) *CityBusAPIOperatorParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the city bus Api operator params
func (o *CityBusAPIOperatorParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the city bus Api operator params
func (o *CityBusAPIOperatorParams) WithDollarSelect(dollarSelect *string) *CityBusAPIOperatorParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the city bus Api operator params
func (o *CityBusAPIOperatorParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the city bus Api operator params
func (o *CityBusAPIOperatorParams) WithDollarSkip(dollarSkip *string) *CityBusAPIOperatorParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the city bus Api operator params
func (o *CityBusAPIOperatorParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the city bus Api operator params
func (o *CityBusAPIOperatorParams) WithDollarTop(dollarTop *int64) *CityBusAPIOperatorParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the city bus Api operator params
func (o *CityBusAPIOperatorParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithCity adds the city to the city bus Api operator params
func (o *CityBusAPIOperatorParams) WithCity(city string) *CityBusAPIOperatorParams {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the city bus Api operator params
func (o *CityBusAPIOperatorParams) SetCity(city string) {
	o.City = city
}

// WithHealth adds the health to the city bus Api operator params
func (o *CityBusAPIOperatorParams) WithHealth(health *string) *CityBusAPIOperatorParams {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the city bus Api operator params
func (o *CityBusAPIOperatorParams) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *CityBusAPIOperatorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
