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

// NewMetroAPIAlert2107Params creates a new MetroAPIAlert2107Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMetroAPIAlert2107Params() *MetroAPIAlert2107Params {
	return &MetroAPIAlert2107Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewMetroAPIAlert2107ParamsWithTimeout creates a new MetroAPIAlert2107Params object
// with the ability to set a timeout on a request.
func NewMetroAPIAlert2107ParamsWithTimeout(timeout time.Duration) *MetroAPIAlert2107Params {
	return &MetroAPIAlert2107Params{
		timeout: timeout,
	}
}

// NewMetroAPIAlert2107ParamsWithContext creates a new MetroAPIAlert2107Params object
// with the ability to set a context for a request.
func NewMetroAPIAlert2107ParamsWithContext(ctx context.Context) *MetroAPIAlert2107Params {
	return &MetroAPIAlert2107Params{
		Context: ctx,
	}
}

// NewMetroAPIAlert2107ParamsWithHTTPClient creates a new MetroAPIAlert2107Params object
// with the ability to set a custom HTTPClient for a request.
func NewMetroAPIAlert2107ParamsWithHTTPClient(client *http.Client) *MetroAPIAlert2107Params {
	return &MetroAPIAlert2107Params{
		HTTPClient: client,
	}
}

/* MetroAPIAlert2107Params contains all the parameters to send to the API endpoint
   for the metro Api alert 2107 operation.

   Typically these are written to a http.Request.
*/
type MetroAPIAlert2107Params struct {

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

	/* RailSystem.

	   欲查詢軌道系統(TRTC:臺北捷運,KRTC:高雄捷運,TYMC:桃園捷運,KLRT:高雄輕軌)
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

// WithDefaults hydrates default values in the metro Api alert 2107 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroAPIAlert2107Params) WithDefaults() *MetroAPIAlert2107Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the metro Api alert 2107 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroAPIAlert2107Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := MetroAPIAlert2107Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) WithTimeout(timeout time.Duration) *MetroAPIAlert2107Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) WithContext(ctx context.Context) *MetroAPIAlert2107Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) WithHTTPClient(client *http.Client) *MetroAPIAlert2107Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) WithDollarCount(dollarCount *bool) *MetroAPIAlert2107Params {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) WithDollarFilter(dollarFilter *string) *MetroAPIAlert2107Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) WithDollarFormat(dollarFormat string) *MetroAPIAlert2107Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) WithDollarOrderby(dollarOrderby *string) *MetroAPIAlert2107Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) WithDollarSelect(dollarSelect *string) *MetroAPIAlert2107Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) WithDollarSkip(dollarSkip *string) *MetroAPIAlert2107Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) WithDollarTop(dollarTop *int64) *MetroAPIAlert2107Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithRailSystem adds the railSystem to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) WithRailSystem(railSystem string) *MetroAPIAlert2107Params {
	o.SetRailSystem(railSystem)
	return o
}

// SetRailSystem adds the railSystem to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) SetRailSystem(railSystem string) {
	o.RailSystem = railSystem
}

// WithHealth adds the health to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) WithHealth(health *string) *MetroAPIAlert2107Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the metro Api alert 2107 params
func (o *MetroAPIAlert2107Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *MetroAPIAlert2107Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
