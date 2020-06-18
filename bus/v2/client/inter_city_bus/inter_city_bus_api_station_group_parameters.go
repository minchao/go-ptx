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

// NewInterCityBusAPIStationGroupParams creates a new InterCityBusAPIStationGroupParams object
// with the default values initialized.
func NewInterCityBusAPIStationGroupParams() *InterCityBusAPIStationGroupParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIStationGroupParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewInterCityBusAPIStationGroupParamsWithTimeout creates a new InterCityBusAPIStationGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInterCityBusAPIStationGroupParamsWithTimeout(timeout time.Duration) *InterCityBusAPIStationGroupParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIStationGroupParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewInterCityBusAPIStationGroupParamsWithContext creates a new InterCityBusAPIStationGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewInterCityBusAPIStationGroupParamsWithContext(ctx context.Context) *InterCityBusAPIStationGroupParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIStationGroupParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewInterCityBusAPIStationGroupParamsWithHTTPClient creates a new InterCityBusAPIStationGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInterCityBusAPIStationGroupParamsWithHTTPClient(client *http.Client) *InterCityBusAPIStationGroupParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIStationGroupParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*InterCityBusAPIStationGroupParams contains all the parameters to send to the API endpoint
for the inter city bus Api station group operation typically these are written to a http.Request
*/
type InterCityBusAPIStationGroupParams struct {

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
	/*Health
	  加入參數'?health=true'即可查詢此API服務的健康狀態

	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) WithTimeout(timeout time.Duration) *InterCityBusAPIStationGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) WithContext(ctx context.Context) *InterCityBusAPIStationGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) WithHTTPClient(client *http.Client) *InterCityBusAPIStationGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) WithDollarFilter(dollarFilter *string) *InterCityBusAPIStationGroupParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) WithDollarFormat(dollarFormat string) *InterCityBusAPIStationGroupParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) WithDollarOrderby(dollarOrderby *string) *InterCityBusAPIStationGroupParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) WithDollarSelect(dollarSelect *string) *InterCityBusAPIStationGroupParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) WithDollarSkip(dollarSkip *string) *InterCityBusAPIStationGroupParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarSpatialFilter adds the dollarSpatialFilter to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) WithDollarSpatialFilter(dollarSpatialFilter *string) *InterCityBusAPIStationGroupParams {
	o.SetDollarSpatialFilter(dollarSpatialFilter)
	return o
}

// SetDollarSpatialFilter adds the dollarSpatialFilter to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) SetDollarSpatialFilter(dollarSpatialFilter *string) {
	o.DollarSpatialFilter = dollarSpatialFilter
}

// WithDollarTop adds the dollarTop to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) WithDollarTop(dollarTop *int64) *InterCityBusAPIStationGroupParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithHealth adds the health to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) WithHealth(health *string) *InterCityBusAPIStationGroupParams {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the inter city bus Api station group params
func (o *InterCityBusAPIStationGroupParams) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *InterCityBusAPIStationGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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