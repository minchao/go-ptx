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

// NewInterCityBusAPIVehicleParams creates a new InterCityBusAPIVehicleParams object
// with the default values initialized.
func NewInterCityBusAPIVehicleParams() *InterCityBusAPIVehicleParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIVehicleParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewInterCityBusAPIVehicleParamsWithTimeout creates a new InterCityBusAPIVehicleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInterCityBusAPIVehicleParamsWithTimeout(timeout time.Duration) *InterCityBusAPIVehicleParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIVehicleParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewInterCityBusAPIVehicleParamsWithContext creates a new InterCityBusAPIVehicleParams object
// with the default values initialized, and the ability to set a context for a request
func NewInterCityBusAPIVehicleParamsWithContext(ctx context.Context) *InterCityBusAPIVehicleParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIVehicleParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewInterCityBusAPIVehicleParamsWithHTTPClient creates a new InterCityBusAPIVehicleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInterCityBusAPIVehicleParamsWithHTTPClient(client *http.Client) *InterCityBusAPIVehicleParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIVehicleParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*InterCityBusAPIVehicleParams contains all the parameters to send to the API endpoint
for the inter city bus Api vehicle operation typically these are written to a http.Request
*/
type InterCityBusAPIVehicleParams struct {

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

// WithTimeout adds the timeout to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) WithTimeout(timeout time.Duration) *InterCityBusAPIVehicleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) WithContext(ctx context.Context) *InterCityBusAPIVehicleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) WithHTTPClient(client *http.Client) *InterCityBusAPIVehicleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) WithDollarFilter(dollarFilter *string) *InterCityBusAPIVehicleParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) WithDollarFormat(dollarFormat string) *InterCityBusAPIVehicleParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) WithDollarOrderby(dollarOrderby *string) *InterCityBusAPIVehicleParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) WithDollarSelect(dollarSelect *string) *InterCityBusAPIVehicleParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) WithDollarSkip(dollarSkip *string) *InterCityBusAPIVehicleParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) WithDollarTop(dollarTop *int64) *InterCityBusAPIVehicleParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithHealth adds the health to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) WithHealth(health *string) *InterCityBusAPIVehicleParams {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the inter city bus Api vehicle params
func (o *InterCityBusAPIVehicleParams) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *InterCityBusAPIVehicleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
