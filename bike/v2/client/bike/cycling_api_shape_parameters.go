// Code generated by go-swagger; DO NOT EDIT.

package bike

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

// NewCyclingAPIShapeParams creates a new CyclingAPIShapeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCyclingAPIShapeParams() *CyclingAPIShapeParams {
	return &CyclingAPIShapeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCyclingAPIShapeParamsWithTimeout creates a new CyclingAPIShapeParams object
// with the ability to set a timeout on a request.
func NewCyclingAPIShapeParamsWithTimeout(timeout time.Duration) *CyclingAPIShapeParams {
	return &CyclingAPIShapeParams{
		timeout: timeout,
	}
}

// NewCyclingAPIShapeParamsWithContext creates a new CyclingAPIShapeParams object
// with the ability to set a context for a request.
func NewCyclingAPIShapeParamsWithContext(ctx context.Context) *CyclingAPIShapeParams {
	return &CyclingAPIShapeParams{
		Context: ctx,
	}
}

// NewCyclingAPIShapeParamsWithHTTPClient creates a new CyclingAPIShapeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCyclingAPIShapeParamsWithHTTPClient(client *http.Client) *CyclingAPIShapeParams {
	return &CyclingAPIShapeParams{
		HTTPClient: client,
	}
}

/* CyclingAPIShapeParams contains all the parameters to send to the API endpoint
   for the cycling Api shape operation.

   Typically these are written to a http.Request.
*/
type CyclingAPIShapeParams struct {

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

	/* City.

	   欲查詢縣市
	*/
	City string

	/* Health.

	   加入參數'?health=true'即可查詢此API服務的健康狀態
	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cycling Api shape params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CyclingAPIShapeParams) WithDefaults() *CyclingAPIShapeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cycling Api shape params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CyclingAPIShapeParams) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := CyclingAPIShapeParams{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cycling Api shape params
func (o *CyclingAPIShapeParams) WithTimeout(timeout time.Duration) *CyclingAPIShapeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cycling Api shape params
func (o *CyclingAPIShapeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cycling Api shape params
func (o *CyclingAPIShapeParams) WithContext(ctx context.Context) *CyclingAPIShapeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cycling Api shape params
func (o *CyclingAPIShapeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cycling Api shape params
func (o *CyclingAPIShapeParams) WithHTTPClient(client *http.Client) *CyclingAPIShapeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cycling Api shape params
func (o *CyclingAPIShapeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the cycling Api shape params
func (o *CyclingAPIShapeParams) WithDollarFilter(dollarFilter *string) *CyclingAPIShapeParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the cycling Api shape params
func (o *CyclingAPIShapeParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the cycling Api shape params
func (o *CyclingAPIShapeParams) WithDollarFormat(dollarFormat string) *CyclingAPIShapeParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the cycling Api shape params
func (o *CyclingAPIShapeParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the cycling Api shape params
func (o *CyclingAPIShapeParams) WithDollarOrderby(dollarOrderby *string) *CyclingAPIShapeParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the cycling Api shape params
func (o *CyclingAPIShapeParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the cycling Api shape params
func (o *CyclingAPIShapeParams) WithDollarSelect(dollarSelect *string) *CyclingAPIShapeParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the cycling Api shape params
func (o *CyclingAPIShapeParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the cycling Api shape params
func (o *CyclingAPIShapeParams) WithDollarSkip(dollarSkip *string) *CyclingAPIShapeParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the cycling Api shape params
func (o *CyclingAPIShapeParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the cycling Api shape params
func (o *CyclingAPIShapeParams) WithDollarTop(dollarTop *int64) *CyclingAPIShapeParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the cycling Api shape params
func (o *CyclingAPIShapeParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithCity adds the city to the cycling Api shape params
func (o *CyclingAPIShapeParams) WithCity(city string) *CyclingAPIShapeParams {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the cycling Api shape params
func (o *CyclingAPIShapeParams) SetCity(city string) {
	o.City = city
}

// WithHealth adds the health to the cycling Api shape params
func (o *CyclingAPIShapeParams) WithHealth(health *string) *CyclingAPIShapeParams {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the cycling Api shape params
func (o *CyclingAPIShapeParams) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *CyclingAPIShapeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
