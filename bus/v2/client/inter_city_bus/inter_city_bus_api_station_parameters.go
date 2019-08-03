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

// NewInterCityBusAPIStationParams creates a new InterCityBusAPIStationParams object
// with the default values initialized.
func NewInterCityBusAPIStationParams() *InterCityBusAPIStationParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIStationParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewInterCityBusAPIStationParamsWithTimeout creates a new InterCityBusAPIStationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInterCityBusAPIStationParamsWithTimeout(timeout time.Duration) *InterCityBusAPIStationParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIStationParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewInterCityBusAPIStationParamsWithContext creates a new InterCityBusAPIStationParams object
// with the default values initialized, and the ability to set a context for a request
func NewInterCityBusAPIStationParamsWithContext(ctx context.Context) *InterCityBusAPIStationParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIStationParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewInterCityBusAPIStationParamsWithHTTPClient creates a new InterCityBusAPIStationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInterCityBusAPIStationParamsWithHTTPClient(client *http.Client) *InterCityBusAPIStationParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &InterCityBusAPIStationParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*InterCityBusAPIStationParams contains all the parameters to send to the API endpoint
for the inter city bus Api station operation typically these are written to a http.Request
*/
type InterCityBusAPIStationParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) WithTimeout(timeout time.Duration) *InterCityBusAPIStationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) WithContext(ctx context.Context) *InterCityBusAPIStationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) WithHTTPClient(client *http.Client) *InterCityBusAPIStationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) WithDollarFilter(dollarFilter *string) *InterCityBusAPIStationParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) WithDollarFormat(dollarFormat string) *InterCityBusAPIStationParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) WithDollarOrderby(dollarOrderby *string) *InterCityBusAPIStationParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) WithDollarSelect(dollarSelect *string) *InterCityBusAPIStationParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) WithDollarSkip(dollarSkip *string) *InterCityBusAPIStationParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarSpatialFilter adds the dollarSpatialFilter to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) WithDollarSpatialFilter(dollarSpatialFilter *string) *InterCityBusAPIStationParams {
	o.SetDollarSpatialFilter(dollarSpatialFilter)
	return o
}

// SetDollarSpatialFilter adds the dollarSpatialFilter to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) SetDollarSpatialFilter(dollarSpatialFilter *string) {
	o.DollarSpatialFilter = dollarSpatialFilter
}

// WithDollarTop adds the dollarTop to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) WithDollarTop(dollarTop *int64) *InterCityBusAPIStationParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the inter city bus Api station params
func (o *InterCityBusAPIStationParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *InterCityBusAPIStationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
