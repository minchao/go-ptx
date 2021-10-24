// Code generated by go-swagger; DO NOT EDIT.

package domestic_ship

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

// NewShipDomesticAlertParams creates a new ShipDomesticAlertParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShipDomesticAlertParams() *ShipDomesticAlertParams {
	return &ShipDomesticAlertParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShipDomesticAlertParamsWithTimeout creates a new ShipDomesticAlertParams object
// with the ability to set a timeout on a request.
func NewShipDomesticAlertParamsWithTimeout(timeout time.Duration) *ShipDomesticAlertParams {
	return &ShipDomesticAlertParams{
		timeout: timeout,
	}
}

// NewShipDomesticAlertParamsWithContext creates a new ShipDomesticAlertParams object
// with the ability to set a context for a request.
func NewShipDomesticAlertParamsWithContext(ctx context.Context) *ShipDomesticAlertParams {
	return &ShipDomesticAlertParams{
		Context: ctx,
	}
}

// NewShipDomesticAlertParamsWithHTTPClient creates a new ShipDomesticAlertParams object
// with the ability to set a custom HTTPClient for a request.
func NewShipDomesticAlertParamsWithHTTPClient(client *http.Client) *ShipDomesticAlertParams {
	return &ShipDomesticAlertParams{
		HTTPClient: client,
	}
}

/* ShipDomesticAlertParams contains all the parameters to send to the API endpoint
   for the ship domestic alert operation.

   Typically these are written to a http.Request.
*/
type ShipDomesticAlertParams struct {

	/* DollarCount.

	   查詢數量
	*/
	DollarCount *bool

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

	// City.
	City string

	/* Health.

	   加入參數'?health=true'即可查詢此API服務的健康狀態
	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ship domestic alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShipDomesticAlertParams) WithDefaults() *ShipDomesticAlertParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ship domestic alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShipDomesticAlertParams) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := ShipDomesticAlertParams{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ship domestic alert params
func (o *ShipDomesticAlertParams) WithTimeout(timeout time.Duration) *ShipDomesticAlertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ship domestic alert params
func (o *ShipDomesticAlertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ship domestic alert params
func (o *ShipDomesticAlertParams) WithContext(ctx context.Context) *ShipDomesticAlertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ship domestic alert params
func (o *ShipDomesticAlertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ship domestic alert params
func (o *ShipDomesticAlertParams) WithHTTPClient(client *http.Client) *ShipDomesticAlertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ship domestic alert params
func (o *ShipDomesticAlertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the ship domestic alert params
func (o *ShipDomesticAlertParams) WithDollarCount(dollarCount *bool) *ShipDomesticAlertParams {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the ship domestic alert params
func (o *ShipDomesticAlertParams) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the ship domestic alert params
func (o *ShipDomesticAlertParams) WithDollarFilter(dollarFilter *string) *ShipDomesticAlertParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the ship domestic alert params
func (o *ShipDomesticAlertParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the ship domestic alert params
func (o *ShipDomesticAlertParams) WithDollarFormat(dollarFormat string) *ShipDomesticAlertParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the ship domestic alert params
func (o *ShipDomesticAlertParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the ship domestic alert params
func (o *ShipDomesticAlertParams) WithDollarOrderby(dollarOrderby *string) *ShipDomesticAlertParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the ship domestic alert params
func (o *ShipDomesticAlertParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the ship domestic alert params
func (o *ShipDomesticAlertParams) WithDollarSelect(dollarSelect *string) *ShipDomesticAlertParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the ship domestic alert params
func (o *ShipDomesticAlertParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the ship domestic alert params
func (o *ShipDomesticAlertParams) WithDollarSkip(dollarSkip *string) *ShipDomesticAlertParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the ship domestic alert params
func (o *ShipDomesticAlertParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the ship domestic alert params
func (o *ShipDomesticAlertParams) WithDollarTop(dollarTop *int64) *ShipDomesticAlertParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the ship domestic alert params
func (o *ShipDomesticAlertParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithCity adds the city to the ship domestic alert params
func (o *ShipDomesticAlertParams) WithCity(city string) *ShipDomesticAlertParams {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the ship domestic alert params
func (o *ShipDomesticAlertParams) SetCity(city string) {
	o.City = city
}

// WithHealth adds the health to the ship domestic alert params
func (o *ShipDomesticAlertParams) WithHealth(health *string) *ShipDomesticAlertParams {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the ship domestic alert params
func (o *ShipDomesticAlertParams) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *ShipDomesticAlertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarCount != nil {

		// query param $count
		var qrDollarCount bool

		if o.DollarCount != nil {
			qrDollarCount = *o.DollarCount
		}
		qDollarCount := swag.FormatBool(qrDollarCount)
		if qDollarCount != "" {

			if err := r.SetQueryParam("$count", qDollarCount); err != nil {
				return err
			}
		}
	}

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

	// path param City
	if err := r.SetPathParam("City", o.City); err != nil {
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
