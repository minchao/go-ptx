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

// NewTourismAPIRestaurantParams creates a new TourismAPIRestaurantParams object
// with the default values initialized.
func NewTourismAPIRestaurantParams() *TourismAPIRestaurantParams {
	var (
		dollarTopDefault = string("30")
	)
	return &TourismAPIRestaurantParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTourismAPIRestaurantParamsWithTimeout creates a new TourismAPIRestaurantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTourismAPIRestaurantParamsWithTimeout(timeout time.Duration) *TourismAPIRestaurantParams {
	var (
		dollarTopDefault = string("30")
	)
	return &TourismAPIRestaurantParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTourismAPIRestaurantParamsWithContext creates a new TourismAPIRestaurantParams object
// with the default values initialized, and the ability to set a context for a request
func NewTourismAPIRestaurantParamsWithContext(ctx context.Context) *TourismAPIRestaurantParams {
	var (
		dollarTopDefault = string("30")
	)
	return &TourismAPIRestaurantParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTourismAPIRestaurantParamsWithHTTPClient creates a new TourismAPIRestaurantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTourismAPIRestaurantParamsWithHTTPClient(client *http.Client) *TourismAPIRestaurantParams {
	var (
		dollarTopDefault = string("30")
	)
	return &TourismAPIRestaurantParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*TourismAPIRestaurantParams contains all the parameters to send to the API endpoint
for the tourism Api restaurant operation typically these are written to a http.Request
*/
type TourismAPIRestaurantParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) WithTimeout(timeout time.Duration) *TourismAPIRestaurantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) WithContext(ctx context.Context) *TourismAPIRestaurantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) WithHTTPClient(client *http.Client) *TourismAPIRestaurantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) WithDollarFilter(dollarFilter *string) *TourismAPIRestaurantParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) WithDollarFormat(dollarFormat string) *TourismAPIRestaurantParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) WithDollarOrderby(dollarOrderby *string) *TourismAPIRestaurantParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) WithDollarSelect(dollarSelect *string) *TourismAPIRestaurantParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) WithDollarSkip(dollarSkip *string) *TourismAPIRestaurantParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarSpatialFilter adds the dollarSpatialFilter to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) WithDollarSpatialFilter(dollarSpatialFilter *string) *TourismAPIRestaurantParams {
	o.SetDollarSpatialFilter(dollarSpatialFilter)
	return o
}

// SetDollarSpatialFilter adds the dollarSpatialFilter to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) SetDollarSpatialFilter(dollarSpatialFilter *string) {
	o.DollarSpatialFilter = dollarSpatialFilter
}

// WithDollarTop adds the dollarTop to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) WithDollarTop(dollarTop *string) *TourismAPIRestaurantParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the tourism Api restaurant params
func (o *TourismAPIRestaurantParams) SetDollarTop(dollarTop *string) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *TourismAPIRestaurantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
