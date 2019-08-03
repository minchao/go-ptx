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

// NewTourismAPIScenicSpot0Params creates a new TourismAPIScenicSpot0Params object
// with the default values initialized.
func NewTourismAPIScenicSpot0Params() *TourismAPIScenicSpot0Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &TourismAPIScenicSpot0Params{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTourismAPIScenicSpot0ParamsWithTimeout creates a new TourismAPIScenicSpot0Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewTourismAPIScenicSpot0ParamsWithTimeout(timeout time.Duration) *TourismAPIScenicSpot0Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &TourismAPIScenicSpot0Params{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTourismAPIScenicSpot0ParamsWithContext creates a new TourismAPIScenicSpot0Params object
// with the default values initialized, and the ability to set a context for a request
func NewTourismAPIScenicSpot0ParamsWithContext(ctx context.Context) *TourismAPIScenicSpot0Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &TourismAPIScenicSpot0Params{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTourismAPIScenicSpot0ParamsWithHTTPClient creates a new TourismAPIScenicSpot0Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTourismAPIScenicSpot0ParamsWithHTTPClient(client *http.Client) *TourismAPIScenicSpot0Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &TourismAPIScenicSpot0Params{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*TourismAPIScenicSpot0Params contains all the parameters to send to the API endpoint
for the tourism Api scenic spot 0 operation typically these are written to a http.Request
*/
type TourismAPIScenicSpot0Params struct {

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
	  縣市名稱

	*/
	City string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) WithTimeout(timeout time.Duration) *TourismAPIScenicSpot0Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) WithContext(ctx context.Context) *TourismAPIScenicSpot0Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) WithHTTPClient(client *http.Client) *TourismAPIScenicSpot0Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) WithDollarFilter(dollarFilter *string) *TourismAPIScenicSpot0Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) WithDollarFormat(dollarFormat string) *TourismAPIScenicSpot0Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) WithDollarOrderby(dollarOrderby *string) *TourismAPIScenicSpot0Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) WithDollarSelect(dollarSelect *string) *TourismAPIScenicSpot0Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) WithDollarSkip(dollarSkip *string) *TourismAPIScenicSpot0Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarSpatialFilter adds the dollarSpatialFilter to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) WithDollarSpatialFilter(dollarSpatialFilter *string) *TourismAPIScenicSpot0Params {
	o.SetDollarSpatialFilter(dollarSpatialFilter)
	return o
}

// SetDollarSpatialFilter adds the dollarSpatialFilter to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) SetDollarSpatialFilter(dollarSpatialFilter *string) {
	o.DollarSpatialFilter = dollarSpatialFilter
}

// WithDollarTop adds the dollarTop to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) WithDollarTop(dollarTop *int64) *TourismAPIScenicSpot0Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithCity adds the city to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) WithCity(city string) *TourismAPIScenicSpot0Params {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the tourism Api scenic spot 0 params
func (o *TourismAPIScenicSpot0Params) SetCity(city string) {
	o.City = city
}

// WriteToRequest writes these params to a swagger request
func (o *TourismAPIScenicSpot0Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
