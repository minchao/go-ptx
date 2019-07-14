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

// NewCityBusAPINewsParams creates a new CityBusAPINewsParams object
// with the default values initialized.
func NewCityBusAPINewsParams() *CityBusAPINewsParams {
	var (
		dollarTopDefault = string("30")
	)
	return &CityBusAPINewsParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCityBusAPINewsParamsWithTimeout creates a new CityBusAPINewsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCityBusAPINewsParamsWithTimeout(timeout time.Duration) *CityBusAPINewsParams {
	var (
		dollarTopDefault = string("30")
	)
	return &CityBusAPINewsParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewCityBusAPINewsParamsWithContext creates a new CityBusAPINewsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCityBusAPINewsParamsWithContext(ctx context.Context) *CityBusAPINewsParams {
	var (
		dollarTopDefault = string("30")
	)
	return &CityBusAPINewsParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewCityBusAPINewsParamsWithHTTPClient creates a new CityBusAPINewsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCityBusAPINewsParamsWithHTTPClient(client *http.Client) *CityBusAPINewsParams {
	var (
		dollarTopDefault = string("30")
	)
	return &CityBusAPINewsParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*CityBusAPINewsParams contains all the parameters to send to the API endpoint
for the city bus Api news operation typically these are written to a http.Request
*/
type CityBusAPINewsParams struct {

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
	DollarTop *string
	/*City
	  欲查詢縣市

	*/
	City string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the city bus Api news params
func (o *CityBusAPINewsParams) WithTimeout(timeout time.Duration) *CityBusAPINewsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the city bus Api news params
func (o *CityBusAPINewsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the city bus Api news params
func (o *CityBusAPINewsParams) WithContext(ctx context.Context) *CityBusAPINewsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the city bus Api news params
func (o *CityBusAPINewsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the city bus Api news params
func (o *CityBusAPINewsParams) WithHTTPClient(client *http.Client) *CityBusAPINewsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the city bus Api news params
func (o *CityBusAPINewsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the city bus Api news params
func (o *CityBusAPINewsParams) WithDollarFilter(dollarFilter *string) *CityBusAPINewsParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the city bus Api news params
func (o *CityBusAPINewsParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the city bus Api news params
func (o *CityBusAPINewsParams) WithDollarFormat(dollarFormat string) *CityBusAPINewsParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the city bus Api news params
func (o *CityBusAPINewsParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the city bus Api news params
func (o *CityBusAPINewsParams) WithDollarOrderby(dollarOrderby *string) *CityBusAPINewsParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the city bus Api news params
func (o *CityBusAPINewsParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the city bus Api news params
func (o *CityBusAPINewsParams) WithDollarSelect(dollarSelect *string) *CityBusAPINewsParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the city bus Api news params
func (o *CityBusAPINewsParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the city bus Api news params
func (o *CityBusAPINewsParams) WithDollarSkip(dollarSkip *string) *CityBusAPINewsParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the city bus Api news params
func (o *CityBusAPINewsParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the city bus Api news params
func (o *CityBusAPINewsParams) WithDollarTop(dollarTop *string) *CityBusAPINewsParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the city bus Api news params
func (o *CityBusAPINewsParams) SetDollarTop(dollarTop *string) {
	o.DollarTop = dollarTop
}

// WithCity adds the city to the city bus Api news params
func (o *CityBusAPINewsParams) WithCity(city string) *CityBusAPINewsParams {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the city bus Api news params
func (o *CityBusAPINewsParams) SetCity(city string) {
	o.City = city
}

// WriteToRequest writes these params to a swagger request
func (o *CityBusAPINewsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
