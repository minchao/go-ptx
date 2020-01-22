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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewInterCityBusAPIRealTimeNearStopParams creates a new InterCityBusAPIRealTimeNearStopParams object
// with the default values initialized.
func NewInterCityBusAPIRealTimeNearStopParams() *InterCityBusAPIRealTimeNearStopParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIRealTimeNearStopParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewInterCityBusAPIRealTimeNearStopParamsWithTimeout creates a new InterCityBusAPIRealTimeNearStopParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInterCityBusAPIRealTimeNearStopParamsWithTimeout(timeout time.Duration) *InterCityBusAPIRealTimeNearStopParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIRealTimeNearStopParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewInterCityBusAPIRealTimeNearStopParamsWithContext creates a new InterCityBusAPIRealTimeNearStopParams object
// with the default values initialized, and the ability to set a context for a request
func NewInterCityBusAPIRealTimeNearStopParamsWithContext(ctx context.Context) *InterCityBusAPIRealTimeNearStopParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIRealTimeNearStopParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewInterCityBusAPIRealTimeNearStopParamsWithHTTPClient creates a new InterCityBusAPIRealTimeNearStopParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInterCityBusAPIRealTimeNearStopParamsWithHTTPClient(client *http.Client) *InterCityBusAPIRealTimeNearStopParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIRealTimeNearStopParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*InterCityBusAPIRealTimeNearStopParams contains all the parameters to send to the API endpoint
for the inter city bus Api real time near stop operation typically these are written to a http.Request
*/
type InterCityBusAPIRealTimeNearStopParams struct {

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
	/*DollarTop
	  取前幾筆

	*/
	DollarTop *int64
	/*Health
	  加入參數'?health=true'即可查詢此API服務的健康狀態

	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) WithTimeout(timeout time.Duration) *InterCityBusAPIRealTimeNearStopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) WithContext(ctx context.Context) *InterCityBusAPIRealTimeNearStopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) WithHTTPClient(client *http.Client) *InterCityBusAPIRealTimeNearStopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) WithDollarFilter(dollarFilter *string) *InterCityBusAPIRealTimeNearStopParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) WithDollarFormat(dollarFormat string) *InterCityBusAPIRealTimeNearStopParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) WithDollarOrderby(dollarOrderby *string) *InterCityBusAPIRealTimeNearStopParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) WithDollarSelect(dollarSelect *string) *InterCityBusAPIRealTimeNearStopParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) WithDollarSkip(dollarSkip *string) *InterCityBusAPIRealTimeNearStopParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) WithDollarTop(dollarTop *int64) *InterCityBusAPIRealTimeNearStopParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithHealth adds the health to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) WithHealth(health *string) *InterCityBusAPIRealTimeNearStopParams {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the inter city bus Api real time near stop params
func (o *InterCityBusAPIRealTimeNearStopParams) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *InterCityBusAPIRealTimeNearStopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
