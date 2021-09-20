// Code generated by go-swagger; DO NOT EDIT.

package air

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

// NewAirAPIMETARParams creates a new AirAPIMETARParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAirAPIMETARParams() *AirAPIMETARParams {
	return &AirAPIMETARParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAirAPIMETARParamsWithTimeout creates a new AirAPIMETARParams object
// with the ability to set a timeout on a request.
func NewAirAPIMETARParamsWithTimeout(timeout time.Duration) *AirAPIMETARParams {
	return &AirAPIMETARParams{
		timeout: timeout,
	}
}

// NewAirAPIMETARParamsWithContext creates a new AirAPIMETARParams object
// with the ability to set a context for a request.
func NewAirAPIMETARParamsWithContext(ctx context.Context) *AirAPIMETARParams {
	return &AirAPIMETARParams{
		Context: ctx,
	}
}

// NewAirAPIMETARParamsWithHTTPClient creates a new AirAPIMETARParams object
// with the ability to set a custom HTTPClient for a request.
func NewAirAPIMETARParamsWithHTTPClient(client *http.Client) *AirAPIMETARParams {
	return &AirAPIMETARParams{
		HTTPClient: client,
	}
}

/* AirAPIMETARParams contains all the parameters to send to the API endpoint
   for the air Api m e t a r operation.

   Typically these are written to a http.Request.
*/
type AirAPIMETARParams struct {

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

	   空間過濾，語法為nearby({Lat},{Lon},{DistanceInMeters})，例如nearby(25.047675, 121.517055, 100)
	*/
	DollarSpatialFilter *string

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

// WithDefaults hydrates default values in the air Api m e t a r params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AirAPIMETARParams) WithDefaults() *AirAPIMETARParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the air Api m e t a r params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AirAPIMETARParams) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := AirAPIMETARParams{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the air Api m e t a r params
func (o *AirAPIMETARParams) WithTimeout(timeout time.Duration) *AirAPIMETARParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the air Api m e t a r params
func (o *AirAPIMETARParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the air Api m e t a r params
func (o *AirAPIMETARParams) WithContext(ctx context.Context) *AirAPIMETARParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the air Api m e t a r params
func (o *AirAPIMETARParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the air Api m e t a r params
func (o *AirAPIMETARParams) WithHTTPClient(client *http.Client) *AirAPIMETARParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the air Api m e t a r params
func (o *AirAPIMETARParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the air Api m e t a r params
func (o *AirAPIMETARParams) WithDollarFilter(dollarFilter *string) *AirAPIMETARParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the air Api m e t a r params
func (o *AirAPIMETARParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the air Api m e t a r params
func (o *AirAPIMETARParams) WithDollarFormat(dollarFormat string) *AirAPIMETARParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the air Api m e t a r params
func (o *AirAPIMETARParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the air Api m e t a r params
func (o *AirAPIMETARParams) WithDollarOrderby(dollarOrderby *string) *AirAPIMETARParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the air Api m e t a r params
func (o *AirAPIMETARParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the air Api m e t a r params
func (o *AirAPIMETARParams) WithDollarSelect(dollarSelect *string) *AirAPIMETARParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the air Api m e t a r params
func (o *AirAPIMETARParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the air Api m e t a r params
func (o *AirAPIMETARParams) WithDollarSkip(dollarSkip *string) *AirAPIMETARParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the air Api m e t a r params
func (o *AirAPIMETARParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarSpatialFilter adds the dollarSpatialFilter to the air Api m e t a r params
func (o *AirAPIMETARParams) WithDollarSpatialFilter(dollarSpatialFilter *string) *AirAPIMETARParams {
	o.SetDollarSpatialFilter(dollarSpatialFilter)
	return o
}

// SetDollarSpatialFilter adds the dollarSpatialFilter to the air Api m e t a r params
func (o *AirAPIMETARParams) SetDollarSpatialFilter(dollarSpatialFilter *string) {
	o.DollarSpatialFilter = dollarSpatialFilter
}

// WithDollarTop adds the dollarTop to the air Api m e t a r params
func (o *AirAPIMETARParams) WithDollarTop(dollarTop *int64) *AirAPIMETARParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the air Api m e t a r params
func (o *AirAPIMETARParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithHealth adds the health to the air Api m e t a r params
func (o *AirAPIMETARParams) WithHealth(health *string) *AirAPIMETARParams {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the air Api m e t a r params
func (o *AirAPIMETARParams) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *AirAPIMETARParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
