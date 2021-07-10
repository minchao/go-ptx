// Code generated by go-swagger; DO NOT EDIT.

package advanced

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

// NewBikeAPIAvailabilityAllCityNearByParams creates a new BikeAPIAvailabilityAllCityNearByParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBikeAPIAvailabilityAllCityNearByParams() *BikeAPIAvailabilityAllCityNearByParams {
	return &BikeAPIAvailabilityAllCityNearByParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBikeAPIAvailabilityAllCityNearByParamsWithTimeout creates a new BikeAPIAvailabilityAllCityNearByParams object
// with the ability to set a timeout on a request.
func NewBikeAPIAvailabilityAllCityNearByParamsWithTimeout(timeout time.Duration) *BikeAPIAvailabilityAllCityNearByParams {
	return &BikeAPIAvailabilityAllCityNearByParams{
		timeout: timeout,
	}
}

// NewBikeAPIAvailabilityAllCityNearByParamsWithContext creates a new BikeAPIAvailabilityAllCityNearByParams object
// with the ability to set a context for a request.
func NewBikeAPIAvailabilityAllCityNearByParamsWithContext(ctx context.Context) *BikeAPIAvailabilityAllCityNearByParams {
	return &BikeAPIAvailabilityAllCityNearByParams{
		Context: ctx,
	}
}

// NewBikeAPIAvailabilityAllCityNearByParamsWithHTTPClient creates a new BikeAPIAvailabilityAllCityNearByParams object
// with the ability to set a custom HTTPClient for a request.
func NewBikeAPIAvailabilityAllCityNearByParamsWithHTTPClient(client *http.Client) *BikeAPIAvailabilityAllCityNearByParams {
	return &BikeAPIAvailabilityAllCityNearByParams{
		HTTPClient: client,
	}
}

/* BikeAPIAvailabilityAllCityNearByParams contains all the parameters to send to the API endpoint
   for the bike Api availability all city near by operation.

   Typically these are written to a http.Request.
*/
type BikeAPIAvailabilityAllCityNearByParams struct {

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

	   空間過濾(最大搜尋半徑為1000公尺)，語法為nearby({Lat},{Lon},{DistanceInMeters})，例如nearby(25.047675, 121.517055, 100)
	*/
	DollarSpatialFilter string

	/* DollarTop.

	   取前幾筆

	   Default: 30
	*/
	DollarTop *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bike Api availability all city near by params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BikeAPIAvailabilityAllCityNearByParams) WithDefaults() *BikeAPIAvailabilityAllCityNearByParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bike Api availability all city near by params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BikeAPIAvailabilityAllCityNearByParams) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := BikeAPIAvailabilityAllCityNearByParams{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) WithTimeout(timeout time.Duration) *BikeAPIAvailabilityAllCityNearByParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) WithContext(ctx context.Context) *BikeAPIAvailabilityAllCityNearByParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) WithHTTPClient(client *http.Client) *BikeAPIAvailabilityAllCityNearByParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) WithDollarFilter(dollarFilter *string) *BikeAPIAvailabilityAllCityNearByParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) WithDollarFormat(dollarFormat string) *BikeAPIAvailabilityAllCityNearByParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) WithDollarOrderby(dollarOrderby *string) *BikeAPIAvailabilityAllCityNearByParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) WithDollarSelect(dollarSelect *string) *BikeAPIAvailabilityAllCityNearByParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) WithDollarSkip(dollarSkip *string) *BikeAPIAvailabilityAllCityNearByParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarSpatialFilter adds the dollarSpatialFilter to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) WithDollarSpatialFilter(dollarSpatialFilter string) *BikeAPIAvailabilityAllCityNearByParams {
	o.SetDollarSpatialFilter(dollarSpatialFilter)
	return o
}

// SetDollarSpatialFilter adds the dollarSpatialFilter to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) SetDollarSpatialFilter(dollarSpatialFilter string) {
	o.DollarSpatialFilter = dollarSpatialFilter
}

// WithDollarTop adds the dollarTop to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) WithDollarTop(dollarTop *int64) *BikeAPIAvailabilityAllCityNearByParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the bike Api availability all city near by params
func (o *BikeAPIAvailabilityAllCityNearByParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *BikeAPIAvailabilityAllCityNearByParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param $spatialFilter
	qrDollarSpatialFilter := o.DollarSpatialFilter
	qDollarSpatialFilter := qrDollarSpatialFilter
	if qDollarSpatialFilter != "" {

		if err := r.SetQueryParam("$spatialFilter", qDollarSpatialFilter); err != nil {
			return err
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
