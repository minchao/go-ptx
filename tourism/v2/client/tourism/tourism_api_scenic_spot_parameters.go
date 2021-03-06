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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewTourismAPIScenicSpotParams creates a new TourismAPIScenicSpotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTourismAPIScenicSpotParams() *TourismAPIScenicSpotParams {
	return &TourismAPIScenicSpotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTourismAPIScenicSpotParamsWithTimeout creates a new TourismAPIScenicSpotParams object
// with the ability to set a timeout on a request.
func NewTourismAPIScenicSpotParamsWithTimeout(timeout time.Duration) *TourismAPIScenicSpotParams {
	return &TourismAPIScenicSpotParams{
		timeout: timeout,
	}
}

// NewTourismAPIScenicSpotParamsWithContext creates a new TourismAPIScenicSpotParams object
// with the ability to set a context for a request.
func NewTourismAPIScenicSpotParamsWithContext(ctx context.Context) *TourismAPIScenicSpotParams {
	return &TourismAPIScenicSpotParams{
		Context: ctx,
	}
}

// NewTourismAPIScenicSpotParamsWithHTTPClient creates a new TourismAPIScenicSpotParams object
// with the ability to set a custom HTTPClient for a request.
func NewTourismAPIScenicSpotParamsWithHTTPClient(client *http.Client) *TourismAPIScenicSpotParams {
	return &TourismAPIScenicSpotParams{
		HTTPClient: client,
	}
}

/* TourismAPIScenicSpotParams contains all the parameters to send to the API endpoint
   for the tourism Api scenic spot operation.

   Typically these are written to a http.Request.
*/
type TourismAPIScenicSpotParams struct {

	/* DollarFilter.

	   過濾
	*/
	DollarFilter *string

	/* DollarFormat.

	   指定來源格式
	*/
	DollarFormat string

	/* DollarOrderby.

	   排序
	*/
	DollarOrderby *string

	/* DollarSelect.

	   挑選
	*/
	DollarSelect *string

	/* DollarSkip.

	   跳過前幾筆
	*/
	DollarSkip *string

	/* DollarSpatialFilter.

	   空間過濾，語法為nearby({Lat},{Lon},{DistanceInMeters})，例如nearby(25.047675, 121.517055, 100)
	*/
	DollarSpatialFilter *string

	/* DollarTop.

	   取前幾筆

	   Default: 30
	*/
	DollarTop *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tourism Api scenic spot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TourismAPIScenicSpotParams) WithDefaults() *TourismAPIScenicSpotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tourism Api scenic spot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TourismAPIScenicSpotParams) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := TourismAPIScenicSpotParams{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) WithTimeout(timeout time.Duration) *TourismAPIScenicSpotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) WithContext(ctx context.Context) *TourismAPIScenicSpotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) WithHTTPClient(client *http.Client) *TourismAPIScenicSpotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) WithDollarFilter(dollarFilter *string) *TourismAPIScenicSpotParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) WithDollarFormat(dollarFormat string) *TourismAPIScenicSpotParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) WithDollarOrderby(dollarOrderby *string) *TourismAPIScenicSpotParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) WithDollarSelect(dollarSelect *string) *TourismAPIScenicSpotParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) WithDollarSkip(dollarSkip *string) *TourismAPIScenicSpotParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarSpatialFilter adds the dollarSpatialFilter to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) WithDollarSpatialFilter(dollarSpatialFilter *string) *TourismAPIScenicSpotParams {
	o.SetDollarSpatialFilter(dollarSpatialFilter)
	return o
}

// SetDollarSpatialFilter adds the dollarSpatialFilter to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) SetDollarSpatialFilter(dollarSpatialFilter *string) {
	o.DollarSpatialFilter = dollarSpatialFilter
}

// WithDollarTop adds the dollarTop to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) WithDollarTop(dollarTop *int64) *TourismAPIScenicSpotParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the tourism Api scenic spot params
func (o *TourismAPIScenicSpotParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *TourismAPIScenicSpotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
