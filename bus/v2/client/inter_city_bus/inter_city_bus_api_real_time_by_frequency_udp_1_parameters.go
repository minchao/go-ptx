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

// NewInterCityBusAPIRealTimeByFrequencyUDP1Params creates a new InterCityBusAPIRealTimeByFrequencyUDP1Params object
// with the default values initialized.
func NewInterCityBusAPIRealTimeByFrequencyUDP1Params() *InterCityBusAPIRealTimeByFrequencyUDP1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIRealTimeByFrequencyUDP1Params{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewInterCityBusAPIRealTimeByFrequencyUDP1ParamsWithTimeout creates a new InterCityBusAPIRealTimeByFrequencyUDP1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewInterCityBusAPIRealTimeByFrequencyUDP1ParamsWithTimeout(timeout time.Duration) *InterCityBusAPIRealTimeByFrequencyUDP1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIRealTimeByFrequencyUDP1Params{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewInterCityBusAPIRealTimeByFrequencyUDP1ParamsWithContext creates a new InterCityBusAPIRealTimeByFrequencyUDP1Params object
// with the default values initialized, and the ability to set a context for a request
func NewInterCityBusAPIRealTimeByFrequencyUDP1ParamsWithContext(ctx context.Context) *InterCityBusAPIRealTimeByFrequencyUDP1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIRealTimeByFrequencyUDP1Params{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewInterCityBusAPIRealTimeByFrequencyUDP1ParamsWithHTTPClient creates a new InterCityBusAPIRealTimeByFrequencyUDP1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInterCityBusAPIRealTimeByFrequencyUDP1ParamsWithHTTPClient(client *http.Client) *InterCityBusAPIRealTimeByFrequencyUDP1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIRealTimeByFrequencyUDP1Params{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*InterCityBusAPIRealTimeByFrequencyUDP1Params contains all the parameters to send to the API endpoint
for the inter city bus Api real time by frequency UDP 1 operation typically these are written to a http.Request
*/
type InterCityBusAPIRealTimeByFrequencyUDP1Params struct {

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
	/*DollarSpatialFilter
	  空間過濾

	*/
	DollarSpatialFilter *string
	/*DollarTop
	  取前幾筆

	*/
	DollarTop *int64
	/*RouteName
	  繁體中文路線名稱，如'9102'

	*/
	RouteName string
	/*Health
	  加入參數'?health=true'即可查詢此API服務的健康狀態

	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) WithTimeout(timeout time.Duration) *InterCityBusAPIRealTimeByFrequencyUDP1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) WithContext(ctx context.Context) *InterCityBusAPIRealTimeByFrequencyUDP1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) WithHTTPClient(client *http.Client) *InterCityBusAPIRealTimeByFrequencyUDP1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) WithDollarFilter(dollarFilter *string) *InterCityBusAPIRealTimeByFrequencyUDP1Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) WithDollarFormat(dollarFormat string) *InterCityBusAPIRealTimeByFrequencyUDP1Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) WithDollarOrderby(dollarOrderby *string) *InterCityBusAPIRealTimeByFrequencyUDP1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) WithDollarSelect(dollarSelect *string) *InterCityBusAPIRealTimeByFrequencyUDP1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) WithDollarSkip(dollarSkip *string) *InterCityBusAPIRealTimeByFrequencyUDP1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarSpatialFilter adds the dollarSpatialFilter to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) WithDollarSpatialFilter(dollarSpatialFilter *string) *InterCityBusAPIRealTimeByFrequencyUDP1Params {
	o.SetDollarSpatialFilter(dollarSpatialFilter)
	return o
}

// SetDollarSpatialFilter adds the dollarSpatialFilter to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) SetDollarSpatialFilter(dollarSpatialFilter *string) {
	o.DollarSpatialFilter = dollarSpatialFilter
}

// WithDollarTop adds the dollarTop to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) WithDollarTop(dollarTop *int64) *InterCityBusAPIRealTimeByFrequencyUDP1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithRouteName adds the routeName to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) WithRouteName(routeName string) *InterCityBusAPIRealTimeByFrequencyUDP1Params {
	o.SetRouteName(routeName)
	return o
}

// SetRouteName adds the routeName to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) SetRouteName(routeName string) {
	o.RouteName = routeName
}

// WithHealth adds the health to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) WithHealth(health *string) *InterCityBusAPIRealTimeByFrequencyUDP1Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the inter city bus Api real time by frequency UDP 1 params
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param RouteName
	if err := r.SetPathParam("RouteName", o.RouteName); err != nil {
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