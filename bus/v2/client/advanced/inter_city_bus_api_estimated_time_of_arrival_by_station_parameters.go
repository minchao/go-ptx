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

// NewInterCityBusAPIEstimatedTimeOfArrivalByStationParams creates a new InterCityBusAPIEstimatedTimeOfArrivalByStationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInterCityBusAPIEstimatedTimeOfArrivalByStationParams() *InterCityBusAPIEstimatedTimeOfArrivalByStationParams {
	return &InterCityBusAPIEstimatedTimeOfArrivalByStationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInterCityBusAPIEstimatedTimeOfArrivalByStationParamsWithTimeout creates a new InterCityBusAPIEstimatedTimeOfArrivalByStationParams object
// with the ability to set a timeout on a request.
func NewInterCityBusAPIEstimatedTimeOfArrivalByStationParamsWithTimeout(timeout time.Duration) *InterCityBusAPIEstimatedTimeOfArrivalByStationParams {
	return &InterCityBusAPIEstimatedTimeOfArrivalByStationParams{
		timeout: timeout,
	}
}

// NewInterCityBusAPIEstimatedTimeOfArrivalByStationParamsWithContext creates a new InterCityBusAPIEstimatedTimeOfArrivalByStationParams object
// with the ability to set a context for a request.
func NewInterCityBusAPIEstimatedTimeOfArrivalByStationParamsWithContext(ctx context.Context) *InterCityBusAPIEstimatedTimeOfArrivalByStationParams {
	return &InterCityBusAPIEstimatedTimeOfArrivalByStationParams{
		Context: ctx,
	}
}

// NewInterCityBusAPIEstimatedTimeOfArrivalByStationParamsWithHTTPClient creates a new InterCityBusAPIEstimatedTimeOfArrivalByStationParams object
// with the ability to set a custom HTTPClient for a request.
func NewInterCityBusAPIEstimatedTimeOfArrivalByStationParamsWithHTTPClient(client *http.Client) *InterCityBusAPIEstimatedTimeOfArrivalByStationParams {
	return &InterCityBusAPIEstimatedTimeOfArrivalByStationParams{
		HTTPClient: client,
	}
}

/* InterCityBusAPIEstimatedTimeOfArrivalByStationParams contains all the parameters to send to the API endpoint
   for the inter city bus Api estimated time of arrival by station operation.

   Typically these are written to a http.Request.
*/
type InterCityBusAPIEstimatedTimeOfArrivalByStationParams struct {

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

	/* DollarTop.

	   取前幾筆

	   Default: 30
	*/
	DollarTop *int64

	/* StationID.

	   欲查詢站位代碼
	*/
	StationID string

	/* Health.

	   加入參數'?health=true'即可查詢此API服務的健康狀態
	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the inter city bus Api estimated time of arrival by station params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) WithDefaults() *InterCityBusAPIEstimatedTimeOfArrivalByStationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the inter city bus Api estimated time of arrival by station params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := InterCityBusAPIEstimatedTimeOfArrivalByStationParams{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) WithTimeout(timeout time.Duration) *InterCityBusAPIEstimatedTimeOfArrivalByStationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) WithContext(ctx context.Context) *InterCityBusAPIEstimatedTimeOfArrivalByStationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) WithHTTPClient(client *http.Client) *InterCityBusAPIEstimatedTimeOfArrivalByStationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) WithDollarFilter(dollarFilter *string) *InterCityBusAPIEstimatedTimeOfArrivalByStationParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) WithDollarFormat(dollarFormat string) *InterCityBusAPIEstimatedTimeOfArrivalByStationParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) WithDollarOrderby(dollarOrderby *string) *InterCityBusAPIEstimatedTimeOfArrivalByStationParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) WithDollarSelect(dollarSelect *string) *InterCityBusAPIEstimatedTimeOfArrivalByStationParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) WithDollarSkip(dollarSkip *string) *InterCityBusAPIEstimatedTimeOfArrivalByStationParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) WithDollarTop(dollarTop *int64) *InterCityBusAPIEstimatedTimeOfArrivalByStationParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithStationID adds the stationID to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) WithStationID(stationID string) *InterCityBusAPIEstimatedTimeOfArrivalByStationParams {
	o.SetStationID(stationID)
	return o
}

// SetStationID adds the stationId to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) SetStationID(stationID string) {
	o.StationID = stationID
}

// WithHealth adds the health to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) WithHealth(health *string) *InterCityBusAPIEstimatedTimeOfArrivalByStationParams {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the inter city bus Api estimated time of arrival by station params
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param StationID
	if err := r.SetPathParam("StationID", o.StationID); err != nil {
		return err
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
