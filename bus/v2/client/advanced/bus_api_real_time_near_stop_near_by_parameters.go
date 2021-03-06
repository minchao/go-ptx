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

// NewBusAPIRealTimeNearStopNearByParams creates a new BusAPIRealTimeNearStopNearByParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBusAPIRealTimeNearStopNearByParams() *BusAPIRealTimeNearStopNearByParams {
	return &BusAPIRealTimeNearStopNearByParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBusAPIRealTimeNearStopNearByParamsWithTimeout creates a new BusAPIRealTimeNearStopNearByParams object
// with the ability to set a timeout on a request.
func NewBusAPIRealTimeNearStopNearByParamsWithTimeout(timeout time.Duration) *BusAPIRealTimeNearStopNearByParams {
	return &BusAPIRealTimeNearStopNearByParams{
		timeout: timeout,
	}
}

// NewBusAPIRealTimeNearStopNearByParamsWithContext creates a new BusAPIRealTimeNearStopNearByParams object
// with the ability to set a context for a request.
func NewBusAPIRealTimeNearStopNearByParamsWithContext(ctx context.Context) *BusAPIRealTimeNearStopNearByParams {
	return &BusAPIRealTimeNearStopNearByParams{
		Context: ctx,
	}
}

// NewBusAPIRealTimeNearStopNearByParamsWithHTTPClient creates a new BusAPIRealTimeNearStopNearByParams object
// with the ability to set a custom HTTPClient for a request.
func NewBusAPIRealTimeNearStopNearByParamsWithHTTPClient(client *http.Client) *BusAPIRealTimeNearStopNearByParams {
	return &BusAPIRealTimeNearStopNearByParams{
		HTTPClient: client,
	}
}

/* BusAPIRealTimeNearStopNearByParams contains all the parameters to send to the API endpoint
   for the bus Api real time near stop near by operation.

   Typically these are written to a http.Request.
*/
type BusAPIRealTimeNearStopNearByParams struct {

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

// WithDefaults hydrates default values in the bus Api real time near stop near by params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BusAPIRealTimeNearStopNearByParams) WithDefaults() *BusAPIRealTimeNearStopNearByParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bus Api real time near stop near by params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BusAPIRealTimeNearStopNearByParams) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := BusAPIRealTimeNearStopNearByParams{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) WithTimeout(timeout time.Duration) *BusAPIRealTimeNearStopNearByParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) WithContext(ctx context.Context) *BusAPIRealTimeNearStopNearByParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) WithHTTPClient(client *http.Client) *BusAPIRealTimeNearStopNearByParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) WithDollarFilter(dollarFilter *string) *BusAPIRealTimeNearStopNearByParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) WithDollarFormat(dollarFormat string) *BusAPIRealTimeNearStopNearByParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) WithDollarOrderby(dollarOrderby *string) *BusAPIRealTimeNearStopNearByParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) WithDollarSelect(dollarSelect *string) *BusAPIRealTimeNearStopNearByParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) WithDollarSkip(dollarSkip *string) *BusAPIRealTimeNearStopNearByParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarSpatialFilter adds the dollarSpatialFilter to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) WithDollarSpatialFilter(dollarSpatialFilter string) *BusAPIRealTimeNearStopNearByParams {
	o.SetDollarSpatialFilter(dollarSpatialFilter)
	return o
}

// SetDollarSpatialFilter adds the dollarSpatialFilter to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) SetDollarSpatialFilter(dollarSpatialFilter string) {
	o.DollarSpatialFilter = dollarSpatialFilter
}

// WithDollarTop adds the dollarTop to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) WithDollarTop(dollarTop *int64) *BusAPIRealTimeNearStopNearByParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithHealth adds the health to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) WithHealth(health *string) *BusAPIRealTimeNearStopNearByParams {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the bus Api real time near stop near by params
func (o *BusAPIRealTimeNearStopNearByParams) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *BusAPIRealTimeNearStopNearByParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
