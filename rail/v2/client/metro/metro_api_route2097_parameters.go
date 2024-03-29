// Code generated by go-swagger; DO NOT EDIT.

package metro

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

// NewMetroAPIRoute2097Params creates a new MetroAPIRoute2097Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMetroAPIRoute2097Params() *MetroAPIRoute2097Params {
	return &MetroAPIRoute2097Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewMetroAPIRoute2097ParamsWithTimeout creates a new MetroAPIRoute2097Params object
// with the ability to set a timeout on a request.
func NewMetroAPIRoute2097ParamsWithTimeout(timeout time.Duration) *MetroAPIRoute2097Params {
	return &MetroAPIRoute2097Params{
		timeout: timeout,
	}
}

// NewMetroAPIRoute2097ParamsWithContext creates a new MetroAPIRoute2097Params object
// with the ability to set a context for a request.
func NewMetroAPIRoute2097ParamsWithContext(ctx context.Context) *MetroAPIRoute2097Params {
	return &MetroAPIRoute2097Params{
		Context: ctx,
	}
}

// NewMetroAPIRoute2097ParamsWithHTTPClient creates a new MetroAPIRoute2097Params object
// with the ability to set a custom HTTPClient for a request.
func NewMetroAPIRoute2097ParamsWithHTTPClient(client *http.Client) *MetroAPIRoute2097Params {
	return &MetroAPIRoute2097Params{
		HTTPClient: client,
	}
}

/* MetroAPIRoute2097Params contains all the parameters to send to the API endpoint
   for the metro Api route 2097 operation.

   Typically these are written to a http.Request.
*/
type MetroAPIRoute2097Params struct {

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

	/* RailSystem.

	   欲查詢軌道系統(TRTC:臺北捷運,KRTC:高雄捷運,TYMC:桃園捷運,TMRT:臺中捷運,NTDLRT:淡海輕軌)
	*/
	RailSystem string

	/* Health.

	   加入參數'?health=true'即可查詢此API服務的健康狀態
	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the metro Api route 2097 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroAPIRoute2097Params) WithDefaults() *MetroAPIRoute2097Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the metro Api route 2097 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroAPIRoute2097Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := MetroAPIRoute2097Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) WithTimeout(timeout time.Duration) *MetroAPIRoute2097Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) WithContext(ctx context.Context) *MetroAPIRoute2097Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) WithHTTPClient(client *http.Client) *MetroAPIRoute2097Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) WithDollarFilter(dollarFilter *string) *MetroAPIRoute2097Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) WithDollarFormat(dollarFormat string) *MetroAPIRoute2097Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) WithDollarOrderby(dollarOrderby *string) *MetroAPIRoute2097Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) WithDollarSelect(dollarSelect *string) *MetroAPIRoute2097Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) WithDollarSkip(dollarSkip *string) *MetroAPIRoute2097Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) WithDollarTop(dollarTop *int64) *MetroAPIRoute2097Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithRailSystem adds the railSystem to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) WithRailSystem(railSystem string) *MetroAPIRoute2097Params {
	o.SetRailSystem(railSystem)
	return o
}

// SetRailSystem adds the railSystem to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) SetRailSystem(railSystem string) {
	o.RailSystem = railSystem
}

// WithHealth adds the health to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) WithHealth(health *string) *MetroAPIRoute2097Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the metro Api route 2097 params
func (o *MetroAPIRoute2097Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *MetroAPIRoute2097Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param RailSystem
	if err := r.SetPathParam("RailSystem", o.RailSystem); err != nil {
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
