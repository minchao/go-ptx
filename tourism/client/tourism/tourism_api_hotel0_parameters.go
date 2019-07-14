// Code generated by go-swagger; DO NOT EDIT.

package tourism

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

// NewTourismAPIHotel0Params creates a new TourismAPIHotel0Params object
// with the default values initialized.
func NewTourismAPIHotel0Params() *TourismAPIHotel0Params {
	var (
		dollarTopDefault = string("30")
	)
	return &TourismAPIHotel0Params{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTourismAPIHotel0ParamsWithTimeout creates a new TourismAPIHotel0Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewTourismAPIHotel0ParamsWithTimeout(timeout time.Duration) *TourismAPIHotel0Params {
	var (
		dollarTopDefault = string("30")
	)
	return &TourismAPIHotel0Params{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTourismAPIHotel0ParamsWithContext creates a new TourismAPIHotel0Params object
// with the default values initialized, and the ability to set a context for a request
func NewTourismAPIHotel0ParamsWithContext(ctx context.Context) *TourismAPIHotel0Params {
	var (
		dollarTopDefault = string("30")
	)
	return &TourismAPIHotel0Params{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTourismAPIHotel0ParamsWithHTTPClient creates a new TourismAPIHotel0Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTourismAPIHotel0ParamsWithHTTPClient(client *http.Client) *TourismAPIHotel0Params {
	var (
		dollarTopDefault = string("30")
	)
	return &TourismAPIHotel0Params{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*TourismAPIHotel0Params contains all the parameters to send to the API endpoint
for the tourism Api hotel 0 operation typically these are written to a http.Request
*/
type TourismAPIHotel0Params struct {

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
	  縣市名稱

	*/
	City string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) WithTimeout(timeout time.Duration) *TourismAPIHotel0Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) WithContext(ctx context.Context) *TourismAPIHotel0Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) WithHTTPClient(client *http.Client) *TourismAPIHotel0Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) WithDollarFilter(dollarFilter *string) *TourismAPIHotel0Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) WithDollarFormat(dollarFormat string) *TourismAPIHotel0Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) WithDollarOrderby(dollarOrderby *string) *TourismAPIHotel0Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) WithDollarSelect(dollarSelect *string) *TourismAPIHotel0Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) WithDollarSkip(dollarSkip *string) *TourismAPIHotel0Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarSpatialFilter adds the dollarSpatialFilter to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) WithDollarSpatialFilter(dollarSpatialFilter *string) *TourismAPIHotel0Params {
	o.SetDollarSpatialFilter(dollarSpatialFilter)
	return o
}

// SetDollarSpatialFilter adds the dollarSpatialFilter to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) SetDollarSpatialFilter(dollarSpatialFilter *string) {
	o.DollarSpatialFilter = dollarSpatialFilter
}

// WithDollarTop adds the dollarTop to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) WithDollarTop(dollarTop *string) *TourismAPIHotel0Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) SetDollarTop(dollarTop *string) {
	o.DollarTop = dollarTop
}

// WithCity adds the city to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) WithCity(city string) *TourismAPIHotel0Params {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the tourism Api hotel 0 params
func (o *TourismAPIHotel0Params) SetCity(city string) {
	o.City = city
}

// WriteToRequest writes these params to a swagger request
func (o *TourismAPIHotel0Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
