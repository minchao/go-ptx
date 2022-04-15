// Code generated by go-swagger; DO NOT EDIT.

package bus_advanced_by_station

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

// NewInterCityBusAPIShapeByStation2903Params creates a new InterCityBusAPIShapeByStation2903Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInterCityBusAPIShapeByStation2903Params() *InterCityBusAPIShapeByStation2903Params {
	return &InterCityBusAPIShapeByStation2903Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewInterCityBusAPIShapeByStation2903ParamsWithTimeout creates a new InterCityBusAPIShapeByStation2903Params object
// with the ability to set a timeout on a request.
func NewInterCityBusAPIShapeByStation2903ParamsWithTimeout(timeout time.Duration) *InterCityBusAPIShapeByStation2903Params {
	return &InterCityBusAPIShapeByStation2903Params{
		timeout: timeout,
	}
}

// NewInterCityBusAPIShapeByStation2903ParamsWithContext creates a new InterCityBusAPIShapeByStation2903Params object
// with the ability to set a context for a request.
func NewInterCityBusAPIShapeByStation2903ParamsWithContext(ctx context.Context) *InterCityBusAPIShapeByStation2903Params {
	return &InterCityBusAPIShapeByStation2903Params{
		Context: ctx,
	}
}

// NewInterCityBusAPIShapeByStation2903ParamsWithHTTPClient creates a new InterCityBusAPIShapeByStation2903Params object
// with the ability to set a custom HTTPClient for a request.
func NewInterCityBusAPIShapeByStation2903ParamsWithHTTPClient(client *http.Client) *InterCityBusAPIShapeByStation2903Params {
	return &InterCityBusAPIShapeByStation2903Params{
		HTTPClient: client,
	}
}

/* InterCityBusAPIShapeByStation2903Params contains all the parameters to send to the API endpoint
   for the inter city bus Api shape by station 2903 operation.

   Typically these are written to a http.Request.
*/
type InterCityBusAPIShapeByStation2903Params struct {

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

// WithDefaults hydrates default values in the inter city bus Api shape by station 2903 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterCityBusAPIShapeByStation2903Params) WithDefaults() *InterCityBusAPIShapeByStation2903Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the inter city bus Api shape by station 2903 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterCityBusAPIShapeByStation2903Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := InterCityBusAPIShapeByStation2903Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) WithTimeout(timeout time.Duration) *InterCityBusAPIShapeByStation2903Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) WithContext(ctx context.Context) *InterCityBusAPIShapeByStation2903Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) WithHTTPClient(client *http.Client) *InterCityBusAPIShapeByStation2903Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) WithDollarFilter(dollarFilter *string) *InterCityBusAPIShapeByStation2903Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) WithDollarFormat(dollarFormat string) *InterCityBusAPIShapeByStation2903Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) WithDollarOrderby(dollarOrderby *string) *InterCityBusAPIShapeByStation2903Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) WithDollarSelect(dollarSelect *string) *InterCityBusAPIShapeByStation2903Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) WithDollarSkip(dollarSkip *string) *InterCityBusAPIShapeByStation2903Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) WithDollarTop(dollarTop *int64) *InterCityBusAPIShapeByStation2903Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithStationID adds the stationID to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) WithStationID(stationID string) *InterCityBusAPIShapeByStation2903Params {
	o.SetStationID(stationID)
	return o
}

// SetStationID adds the stationId to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) SetStationID(stationID string) {
	o.StationID = stationID
}

// WithHealth adds the health to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) WithHealth(health *string) *InterCityBusAPIShapeByStation2903Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the inter city bus Api shape by station 2903 params
func (o *InterCityBusAPIShapeByStation2903Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *InterCityBusAPIShapeByStation2903Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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