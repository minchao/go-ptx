// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

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

// NewNewsAPIControllerGetParams creates a new NewsAPIControllerGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNewsAPIControllerGetParams() *NewsAPIControllerGetParams {
	return &NewsAPIControllerGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNewsAPIControllerGetParamsWithTimeout creates a new NewsAPIControllerGetParams object
// with the ability to set a timeout on a request.
func NewNewsAPIControllerGetParamsWithTimeout(timeout time.Duration) *NewsAPIControllerGetParams {
	return &NewsAPIControllerGetParams{
		timeout: timeout,
	}
}

// NewNewsAPIControllerGetParamsWithContext creates a new NewsAPIControllerGetParams object
// with the ability to set a context for a request.
func NewNewsAPIControllerGetParamsWithContext(ctx context.Context) *NewsAPIControllerGetParams {
	return &NewsAPIControllerGetParams{
		Context: ctx,
	}
}

// NewNewsAPIControllerGetParamsWithHTTPClient creates a new NewsAPIControllerGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNewsAPIControllerGetParamsWithHTTPClient(client *http.Client) *NewsAPIControllerGetParams {
	return &NewsAPIControllerGetParams{
		HTTPClient: client,
	}
}

/* NewsAPIControllerGetParams contains all the parameters to send to the API endpoint
   for the news Api controller get operation.

   Typically these are written to a http.Request.
*/
type NewsAPIControllerGetParams struct {

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

	/* Health.

	   加入參數'?health=true'即可查詢此API服務的健康狀態
	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the news Api controller get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NewsAPIControllerGetParams) WithDefaults() *NewsAPIControllerGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the news Api controller get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NewsAPIControllerGetParams) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := NewsAPIControllerGetParams{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the news Api controller get params
func (o *NewsAPIControllerGetParams) WithTimeout(timeout time.Duration) *NewsAPIControllerGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the news Api controller get params
func (o *NewsAPIControllerGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the news Api controller get params
func (o *NewsAPIControllerGetParams) WithContext(ctx context.Context) *NewsAPIControllerGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the news Api controller get params
func (o *NewsAPIControllerGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the news Api controller get params
func (o *NewsAPIControllerGetParams) WithHTTPClient(client *http.Client) *NewsAPIControllerGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the news Api controller get params
func (o *NewsAPIControllerGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the news Api controller get params
func (o *NewsAPIControllerGetParams) WithDollarCount(dollarCount *bool) *NewsAPIControllerGetParams {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the news Api controller get params
func (o *NewsAPIControllerGetParams) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the news Api controller get params
func (o *NewsAPIControllerGetParams) WithDollarFilter(dollarFilter *string) *NewsAPIControllerGetParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the news Api controller get params
func (o *NewsAPIControllerGetParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the news Api controller get params
func (o *NewsAPIControllerGetParams) WithDollarFormat(dollarFormat string) *NewsAPIControllerGetParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the news Api controller get params
func (o *NewsAPIControllerGetParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the news Api controller get params
func (o *NewsAPIControllerGetParams) WithDollarOrderby(dollarOrderby *string) *NewsAPIControllerGetParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the news Api controller get params
func (o *NewsAPIControllerGetParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the news Api controller get params
func (o *NewsAPIControllerGetParams) WithDollarSelect(dollarSelect *string) *NewsAPIControllerGetParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the news Api controller get params
func (o *NewsAPIControllerGetParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the news Api controller get params
func (o *NewsAPIControllerGetParams) WithDollarSkip(dollarSkip *string) *NewsAPIControllerGetParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the news Api controller get params
func (o *NewsAPIControllerGetParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the news Api controller get params
func (o *NewsAPIControllerGetParams) WithDollarTop(dollarTop *int64) *NewsAPIControllerGetParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the news Api controller get params
func (o *NewsAPIControllerGetParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithHealth adds the health to the news Api controller get params
func (o *NewsAPIControllerGetParams) WithHealth(health *string) *NewsAPIControllerGetParams {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the news Api controller get params
func (o *NewsAPIControllerGetParams) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *NewsAPIControllerGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
