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

// NewMetroAPIStationOfRoute2098Params creates a new MetroAPIStationOfRoute2098Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMetroAPIStationOfRoute2098Params() *MetroAPIStationOfRoute2098Params {
	return &MetroAPIStationOfRoute2098Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewMetroAPIStationOfRoute2098ParamsWithTimeout creates a new MetroAPIStationOfRoute2098Params object
// with the ability to set a timeout on a request.
func NewMetroAPIStationOfRoute2098ParamsWithTimeout(timeout time.Duration) *MetroAPIStationOfRoute2098Params {
	return &MetroAPIStationOfRoute2098Params{
		timeout: timeout,
	}
}

// NewMetroAPIStationOfRoute2098ParamsWithContext creates a new MetroAPIStationOfRoute2098Params object
// with the ability to set a context for a request.
func NewMetroAPIStationOfRoute2098ParamsWithContext(ctx context.Context) *MetroAPIStationOfRoute2098Params {
	return &MetroAPIStationOfRoute2098Params{
		Context: ctx,
	}
}

// NewMetroAPIStationOfRoute2098ParamsWithHTTPClient creates a new MetroAPIStationOfRoute2098Params object
// with the ability to set a custom HTTPClient for a request.
func NewMetroAPIStationOfRoute2098ParamsWithHTTPClient(client *http.Client) *MetroAPIStationOfRoute2098Params {
	return &MetroAPIStationOfRoute2098Params{
		HTTPClient: client,
	}
}

/* MetroAPIStationOfRoute2098Params contains all the parameters to send to the API endpoint
   for the metro Api station of route 2098 operation.

   Typically these are written to a http.Request.
*/
type MetroAPIStationOfRoute2098Params struct {

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

// WithDefaults hydrates default values in the metro Api station of route 2098 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroAPIStationOfRoute2098Params) WithDefaults() *MetroAPIStationOfRoute2098Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the metro Api station of route 2098 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroAPIStationOfRoute2098Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := MetroAPIStationOfRoute2098Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) WithTimeout(timeout time.Duration) *MetroAPIStationOfRoute2098Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) WithContext(ctx context.Context) *MetroAPIStationOfRoute2098Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) WithHTTPClient(client *http.Client) *MetroAPIStationOfRoute2098Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) WithDollarFilter(dollarFilter *string) *MetroAPIStationOfRoute2098Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) WithDollarFormat(dollarFormat string) *MetroAPIStationOfRoute2098Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) WithDollarOrderby(dollarOrderby *string) *MetroAPIStationOfRoute2098Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) WithDollarSelect(dollarSelect *string) *MetroAPIStationOfRoute2098Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) WithDollarSkip(dollarSkip *string) *MetroAPIStationOfRoute2098Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) WithDollarTop(dollarTop *int64) *MetroAPIStationOfRoute2098Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithRailSystem adds the railSystem to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) WithRailSystem(railSystem string) *MetroAPIStationOfRoute2098Params {
	o.SetRailSystem(railSystem)
	return o
}

// SetRailSystem adds the railSystem to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) SetRailSystem(railSystem string) {
	o.RailSystem = railSystem
}

// WithHealth adds the health to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) WithHealth(health *string) *MetroAPIStationOfRoute2098Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the metro Api station of route 2098 params
func (o *MetroAPIStationOfRoute2098Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *MetroAPIStationOfRoute2098Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
