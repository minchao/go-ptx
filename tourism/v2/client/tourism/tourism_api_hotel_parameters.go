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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTourismAPIHotelParams creates a new TourismAPIHotelParams object
// with the default values initialized.
func NewTourismAPIHotelParams() *TourismAPIHotelParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &TourismAPIHotelParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTourismAPIHotelParamsWithTimeout creates a new TourismAPIHotelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTourismAPIHotelParamsWithTimeout(timeout time.Duration) *TourismAPIHotelParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &TourismAPIHotelParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTourismAPIHotelParamsWithContext creates a new TourismAPIHotelParams object
// with the default values initialized, and the ability to set a context for a request
func NewTourismAPIHotelParamsWithContext(ctx context.Context) *TourismAPIHotelParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &TourismAPIHotelParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTourismAPIHotelParamsWithHTTPClient creates a new TourismAPIHotelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTourismAPIHotelParamsWithHTTPClient(client *http.Client) *TourismAPIHotelParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &TourismAPIHotelParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*TourismAPIHotelParams contains all the parameters to send to the API endpoint
for the tourism Api hotel operation typically these are written to a http.Request
*/
type TourismAPIHotelParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tourism Api hotel params
func (o *TourismAPIHotelParams) WithTimeout(timeout time.Duration) *TourismAPIHotelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tourism Api hotel params
func (o *TourismAPIHotelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tourism Api hotel params
func (o *TourismAPIHotelParams) WithContext(ctx context.Context) *TourismAPIHotelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tourism Api hotel params
func (o *TourismAPIHotelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tourism Api hotel params
func (o *TourismAPIHotelParams) WithHTTPClient(client *http.Client) *TourismAPIHotelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tourism Api hotel params
func (o *TourismAPIHotelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the tourism Api hotel params
func (o *TourismAPIHotelParams) WithDollarFilter(dollarFilter *string) *TourismAPIHotelParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the tourism Api hotel params
func (o *TourismAPIHotelParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the tourism Api hotel params
func (o *TourismAPIHotelParams) WithDollarFormat(dollarFormat string) *TourismAPIHotelParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the tourism Api hotel params
func (o *TourismAPIHotelParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the tourism Api hotel params
func (o *TourismAPIHotelParams) WithDollarOrderby(dollarOrderby *string) *TourismAPIHotelParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the tourism Api hotel params
func (o *TourismAPIHotelParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the tourism Api hotel params
func (o *TourismAPIHotelParams) WithDollarSelect(dollarSelect *string) *TourismAPIHotelParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the tourism Api hotel params
func (o *TourismAPIHotelParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the tourism Api hotel params
func (o *TourismAPIHotelParams) WithDollarSkip(dollarSkip *string) *TourismAPIHotelParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the tourism Api hotel params
func (o *TourismAPIHotelParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarSpatialFilter adds the dollarSpatialFilter to the tourism Api hotel params
func (o *TourismAPIHotelParams) WithDollarSpatialFilter(dollarSpatialFilter *string) *TourismAPIHotelParams {
	o.SetDollarSpatialFilter(dollarSpatialFilter)
	return o
}

// SetDollarSpatialFilter adds the dollarSpatialFilter to the tourism Api hotel params
func (o *TourismAPIHotelParams) SetDollarSpatialFilter(dollarSpatialFilter *string) {
	o.DollarSpatialFilter = dollarSpatialFilter
}

// WithDollarTop adds the dollarTop to the tourism Api hotel params
func (o *TourismAPIHotelParams) WithDollarTop(dollarTop *int64) *TourismAPIHotelParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the tourism Api hotel params
func (o *TourismAPIHotelParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *TourismAPIHotelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
