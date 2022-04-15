// Code generated by go-swagger; DO NOT EDIT.

package bus_advanced_near_by

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

// NewBusAPIRouteNearBy2852Params creates a new BusAPIRouteNearBy2852Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBusAPIRouteNearBy2852Params() *BusAPIRouteNearBy2852Params {
	return &BusAPIRouteNearBy2852Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewBusAPIRouteNearBy2852ParamsWithTimeout creates a new BusAPIRouteNearBy2852Params object
// with the ability to set a timeout on a request.
func NewBusAPIRouteNearBy2852ParamsWithTimeout(timeout time.Duration) *BusAPIRouteNearBy2852Params {
	return &BusAPIRouteNearBy2852Params{
		timeout: timeout,
	}
}

// NewBusAPIRouteNearBy2852ParamsWithContext creates a new BusAPIRouteNearBy2852Params object
// with the ability to set a context for a request.
func NewBusAPIRouteNearBy2852ParamsWithContext(ctx context.Context) *BusAPIRouteNearBy2852Params {
	return &BusAPIRouteNearBy2852Params{
		Context: ctx,
	}
}

// NewBusAPIRouteNearBy2852ParamsWithHTTPClient creates a new BusAPIRouteNearBy2852Params object
// with the ability to set a custom HTTPClient for a request.
func NewBusAPIRouteNearBy2852ParamsWithHTTPClient(client *http.Client) *BusAPIRouteNearBy2852Params {
	return &BusAPIRouteNearBy2852Params{
		HTTPClient: client,
	}
}

/* BusAPIRouteNearBy2852Params contains all the parameters to send to the API endpoint
   for the bus Api route near by 2852 operation.

   Typically these are written to a http.Request.
*/
type BusAPIRouteNearBy2852Params struct {

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

	/* Health.

	   加入參數'?health=true'即可查詢此API服務的健康狀態
	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bus Api route near by 2852 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BusAPIRouteNearBy2852Params) WithDefaults() *BusAPIRouteNearBy2852Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bus Api route near by 2852 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BusAPIRouteNearBy2852Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := BusAPIRouteNearBy2852Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) WithTimeout(timeout time.Duration) *BusAPIRouteNearBy2852Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) WithContext(ctx context.Context) *BusAPIRouteNearBy2852Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) WithHTTPClient(client *http.Client) *BusAPIRouteNearBy2852Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) WithDollarFilter(dollarFilter *string) *BusAPIRouteNearBy2852Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) WithDollarFormat(dollarFormat string) *BusAPIRouteNearBy2852Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) WithDollarOrderby(dollarOrderby *string) *BusAPIRouteNearBy2852Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) WithDollarSelect(dollarSelect *string) *BusAPIRouteNearBy2852Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) WithDollarSkip(dollarSkip *string) *BusAPIRouteNearBy2852Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarSpatialFilter adds the dollarSpatialFilter to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) WithDollarSpatialFilter(dollarSpatialFilter string) *BusAPIRouteNearBy2852Params {
	o.SetDollarSpatialFilter(dollarSpatialFilter)
	return o
}

// SetDollarSpatialFilter adds the dollarSpatialFilter to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) SetDollarSpatialFilter(dollarSpatialFilter string) {
	o.DollarSpatialFilter = dollarSpatialFilter
}

// WithDollarTop adds the dollarTop to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) WithDollarTop(dollarTop *int64) *BusAPIRouteNearBy2852Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithHealth adds the health to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) WithHealth(health *string) *BusAPIRouteNearBy2852Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the bus Api route near by 2852 params
func (o *BusAPIRouteNearBy2852Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *BusAPIRouteNearBy2852Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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