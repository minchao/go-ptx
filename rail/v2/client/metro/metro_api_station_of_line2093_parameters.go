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

// NewMetroAPIStationOfLine2093Params creates a new MetroAPIStationOfLine2093Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMetroAPIStationOfLine2093Params() *MetroAPIStationOfLine2093Params {
	return &MetroAPIStationOfLine2093Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewMetroAPIStationOfLine2093ParamsWithTimeout creates a new MetroAPIStationOfLine2093Params object
// with the ability to set a timeout on a request.
func NewMetroAPIStationOfLine2093ParamsWithTimeout(timeout time.Duration) *MetroAPIStationOfLine2093Params {
	return &MetroAPIStationOfLine2093Params{
		timeout: timeout,
	}
}

// NewMetroAPIStationOfLine2093ParamsWithContext creates a new MetroAPIStationOfLine2093Params object
// with the ability to set a context for a request.
func NewMetroAPIStationOfLine2093ParamsWithContext(ctx context.Context) *MetroAPIStationOfLine2093Params {
	return &MetroAPIStationOfLine2093Params{
		Context: ctx,
	}
}

// NewMetroAPIStationOfLine2093ParamsWithHTTPClient creates a new MetroAPIStationOfLine2093Params object
// with the ability to set a custom HTTPClient for a request.
func NewMetroAPIStationOfLine2093ParamsWithHTTPClient(client *http.Client) *MetroAPIStationOfLine2093Params {
	return &MetroAPIStationOfLine2093Params{
		HTTPClient: client,
	}
}

/* MetroAPIStationOfLine2093Params contains all the parameters to send to the API endpoint
   for the metro Api station of line 2093 operation.

   Typically these are written to a http.Request.
*/
type MetroAPIStationOfLine2093Params struct {

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

	   欲查詢軌道系統(TRTC:臺北捷運,KRTC:高雄捷運,TYMC:桃園捷運,TMRT:臺中捷運,KLRT:高雄輕軌,NTDLRT:淡海輕軌,TRTCMG:貓空纜車)
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

// WithDefaults hydrates default values in the metro Api station of line 2093 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroAPIStationOfLine2093Params) WithDefaults() *MetroAPIStationOfLine2093Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the metro Api station of line 2093 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroAPIStationOfLine2093Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := MetroAPIStationOfLine2093Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) WithTimeout(timeout time.Duration) *MetroAPIStationOfLine2093Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) WithContext(ctx context.Context) *MetroAPIStationOfLine2093Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) WithHTTPClient(client *http.Client) *MetroAPIStationOfLine2093Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) WithDollarFilter(dollarFilter *string) *MetroAPIStationOfLine2093Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) WithDollarFormat(dollarFormat string) *MetroAPIStationOfLine2093Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) WithDollarOrderby(dollarOrderby *string) *MetroAPIStationOfLine2093Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) WithDollarSelect(dollarSelect *string) *MetroAPIStationOfLine2093Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) WithDollarSkip(dollarSkip *string) *MetroAPIStationOfLine2093Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) WithDollarTop(dollarTop *int64) *MetroAPIStationOfLine2093Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithRailSystem adds the railSystem to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) WithRailSystem(railSystem string) *MetroAPIStationOfLine2093Params {
	o.SetRailSystem(railSystem)
	return o
}

// SetRailSystem adds the railSystem to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) SetRailSystem(railSystem string) {
	o.RailSystem = railSystem
}

// WithHealth adds the health to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) WithHealth(health *string) *MetroAPIStationOfLine2093Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the metro Api station of line 2093 params
func (o *MetroAPIStationOfLine2093Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *MetroAPIStationOfLine2093Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
