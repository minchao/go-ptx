// Code generated by go-swagger; DO NOT EDIT.

package tourism

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

// NewTourismAPIScenicSpot22401Params creates a new TourismAPIScenicSpot22401Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTourismAPIScenicSpot22401Params() *TourismAPIScenicSpot22401Params {
	return &TourismAPIScenicSpot22401Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewTourismAPIScenicSpot22401ParamsWithTimeout creates a new TourismAPIScenicSpot22401Params object
// with the ability to set a timeout on a request.
func NewTourismAPIScenicSpot22401ParamsWithTimeout(timeout time.Duration) *TourismAPIScenicSpot22401Params {
	return &TourismAPIScenicSpot22401Params{
		timeout: timeout,
	}
}

// NewTourismAPIScenicSpot22401ParamsWithContext creates a new TourismAPIScenicSpot22401Params object
// with the ability to set a context for a request.
func NewTourismAPIScenicSpot22401ParamsWithContext(ctx context.Context) *TourismAPIScenicSpot22401Params {
	return &TourismAPIScenicSpot22401Params{
		Context: ctx,
	}
}

// NewTourismAPIScenicSpot22401ParamsWithHTTPClient creates a new TourismAPIScenicSpot22401Params object
// with the ability to set a custom HTTPClient for a request.
func NewTourismAPIScenicSpot22401ParamsWithHTTPClient(client *http.Client) *TourismAPIScenicSpot22401Params {
	return &TourismAPIScenicSpot22401Params{
		HTTPClient: client,
	}
}

/* TourismAPIScenicSpot22401Params contains all the parameters to send to the API endpoint
   for the tourism Api scenic spot 2240 1 operation.

   Typically these are written to a http.Request.
*/
type TourismAPIScenicSpot22401Params struct {

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

	   縣市名稱
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

// WithDefaults hydrates default values in the tourism Api scenic spot 2240 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TourismAPIScenicSpot22401Params) WithDefaults() *TourismAPIScenicSpot22401Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tourism Api scenic spot 2240 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TourismAPIScenicSpot22401Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := TourismAPIScenicSpot22401Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) WithTimeout(timeout time.Duration) *TourismAPIScenicSpot22401Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) WithContext(ctx context.Context) *TourismAPIScenicSpot22401Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) WithHTTPClient(client *http.Client) *TourismAPIScenicSpot22401Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) WithDollarFilter(dollarFilter *string) *TourismAPIScenicSpot22401Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) WithDollarFormat(dollarFormat string) *TourismAPIScenicSpot22401Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) WithDollarOrderby(dollarOrderby *string) *TourismAPIScenicSpot22401Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) WithDollarSelect(dollarSelect *string) *TourismAPIScenicSpot22401Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) WithDollarSkip(dollarSkip *string) *TourismAPIScenicSpot22401Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) WithDollarTop(dollarTop *int64) *TourismAPIScenicSpot22401Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithCity adds the city to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) WithCity(city string) *TourismAPIScenicSpot22401Params {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) SetCity(city string) {
	o.City = city
}

// WithHealth adds the health to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) WithHealth(health *string) *TourismAPIScenicSpot22401Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the tourism Api scenic spot 2240 1 params
func (o *TourismAPIScenicSpot22401Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *TourismAPIScenicSpot22401Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
