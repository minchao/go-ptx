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

// NewCityBusAPIRouteFareParams creates a new CityBusAPIRouteFareParams object
// with the default values initialized.
func NewCityBusAPIRouteFareParams() *CityBusAPIRouteFareParams {
	var (
		dollarTopDefault = string("30")
	)
	return &CityBusAPIRouteFareParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCityBusAPIRouteFareParamsWithTimeout creates a new CityBusAPIRouteFareParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCityBusAPIRouteFareParamsWithTimeout(timeout time.Duration) *CityBusAPIRouteFareParams {
	var (
		dollarTopDefault = string("30")
	)
	return &CityBusAPIRouteFareParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewCityBusAPIRouteFareParamsWithContext creates a new CityBusAPIRouteFareParams object
// with the default values initialized, and the ability to set a context for a request
func NewCityBusAPIRouteFareParamsWithContext(ctx context.Context) *CityBusAPIRouteFareParams {
	var (
		dollarTopDefault = string("30")
	)
	return &CityBusAPIRouteFareParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewCityBusAPIRouteFareParamsWithHTTPClient creates a new CityBusAPIRouteFareParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCityBusAPIRouteFareParamsWithHTTPClient(client *http.Client) *CityBusAPIRouteFareParams {
	var (
		dollarTopDefault = string("30")
	)
	return &CityBusAPIRouteFareParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*CityBusAPIRouteFareParams contains all the parameters to send to the API endpoint
for the city bus Api route fare operation typically these are written to a http.Request
*/
type CityBusAPIRouteFareParams struct {

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

// WithTimeout adds the timeout to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) WithTimeout(timeout time.Duration) *CityBusAPIRouteFareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) WithContext(ctx context.Context) *CityBusAPIRouteFareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) WithHTTPClient(client *http.Client) *CityBusAPIRouteFareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) WithDollarFilter(dollarFilter *string) *CityBusAPIRouteFareParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) WithDollarFormat(dollarFormat string) *CityBusAPIRouteFareParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) WithDollarOrderby(dollarOrderby *string) *CityBusAPIRouteFareParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) WithDollarSelect(dollarSelect *string) *CityBusAPIRouteFareParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) WithDollarSkip(dollarSkip *string) *CityBusAPIRouteFareParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) WithDollarTop(dollarTop *string) *CityBusAPIRouteFareParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) SetDollarTop(dollarTop *string) {
	o.DollarTop = dollarTop
}

// WithCity adds the city to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) WithCity(city string) *CityBusAPIRouteFareParams {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the city bus Api route fare params
func (o *CityBusAPIRouteFareParams) SetCity(city string) {
	o.City = city
}

// WriteToRequest writes these params to a swagger request
func (o *CityBusAPIRouteFareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
