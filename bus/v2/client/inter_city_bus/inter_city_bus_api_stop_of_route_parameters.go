// Code generated by go-swagger; DO NOT EDIT.

package inter_city_bus

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

// NewInterCityBusAPIStopOfRouteParams creates a new InterCityBusAPIStopOfRouteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInterCityBusAPIStopOfRouteParams() *InterCityBusAPIStopOfRouteParams {
	return &InterCityBusAPIStopOfRouteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInterCityBusAPIStopOfRouteParamsWithTimeout creates a new InterCityBusAPIStopOfRouteParams object
// with the ability to set a timeout on a request.
func NewInterCityBusAPIStopOfRouteParamsWithTimeout(timeout time.Duration) *InterCityBusAPIStopOfRouteParams {
	return &InterCityBusAPIStopOfRouteParams{
		timeout: timeout,
	}
}

// NewInterCityBusAPIStopOfRouteParamsWithContext creates a new InterCityBusAPIStopOfRouteParams object
// with the ability to set a context for a request.
func NewInterCityBusAPIStopOfRouteParamsWithContext(ctx context.Context) *InterCityBusAPIStopOfRouteParams {
	return &InterCityBusAPIStopOfRouteParams{
		Context: ctx,
	}
}

// NewInterCityBusAPIStopOfRouteParamsWithHTTPClient creates a new InterCityBusAPIStopOfRouteParams object
// with the ability to set a custom HTTPClient for a request.
func NewInterCityBusAPIStopOfRouteParamsWithHTTPClient(client *http.Client) *InterCityBusAPIStopOfRouteParams {
	return &InterCityBusAPIStopOfRouteParams{
		HTTPClient: client,
	}
}

/* InterCityBusAPIStopOfRouteParams contains all the parameters to send to the API endpoint
   for the inter city bus Api stop of route operation.

   Typically these are written to a http.Request.
*/
type InterCityBusAPIStopOfRouteParams struct {

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

	/* Health.

	   加入參數'?health=true'即可查詢此API服務的健康狀態
	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the inter city bus Api stop of route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterCityBusAPIStopOfRouteParams) WithDefaults() *InterCityBusAPIStopOfRouteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the inter city bus Api stop of route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterCityBusAPIStopOfRouteParams) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := InterCityBusAPIStopOfRouteParams{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) WithTimeout(timeout time.Duration) *InterCityBusAPIStopOfRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) WithContext(ctx context.Context) *InterCityBusAPIStopOfRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) WithHTTPClient(client *http.Client) *InterCityBusAPIStopOfRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) WithDollarFilter(dollarFilter *string) *InterCityBusAPIStopOfRouteParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) WithDollarFormat(dollarFormat string) *InterCityBusAPIStopOfRouteParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) WithDollarOrderby(dollarOrderby *string) *InterCityBusAPIStopOfRouteParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) WithDollarSelect(dollarSelect *string) *InterCityBusAPIStopOfRouteParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) WithDollarSkip(dollarSkip *string) *InterCityBusAPIStopOfRouteParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) WithDollarTop(dollarTop *int64) *InterCityBusAPIStopOfRouteParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithHealth adds the health to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) WithHealth(health *string) *InterCityBusAPIStopOfRouteParams {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the inter city bus Api stop of route params
func (o *InterCityBusAPIStopOfRouteParams) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *InterCityBusAPIStopOfRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
