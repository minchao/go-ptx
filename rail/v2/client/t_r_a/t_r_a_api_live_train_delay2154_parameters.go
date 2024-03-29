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

// NewTRAAPILiveTrainDelay2154Params creates a new TRAAPILiveTrainDelay2154Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTRAAPILiveTrainDelay2154Params() *TRAAPILiveTrainDelay2154Params {
	return &TRAAPILiveTrainDelay2154Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewTRAAPILiveTrainDelay2154ParamsWithTimeout creates a new TRAAPILiveTrainDelay2154Params object
// with the ability to set a timeout on a request.
func NewTRAAPILiveTrainDelay2154ParamsWithTimeout(timeout time.Duration) *TRAAPILiveTrainDelay2154Params {
	return &TRAAPILiveTrainDelay2154Params{
		timeout: timeout,
	}
}

// NewTRAAPILiveTrainDelay2154ParamsWithContext creates a new TRAAPILiveTrainDelay2154Params object
// with the ability to set a context for a request.
func NewTRAAPILiveTrainDelay2154ParamsWithContext(ctx context.Context) *TRAAPILiveTrainDelay2154Params {
	return &TRAAPILiveTrainDelay2154Params{
		Context: ctx,
	}
}

// NewTRAAPILiveTrainDelay2154ParamsWithHTTPClient creates a new TRAAPILiveTrainDelay2154Params object
// with the ability to set a custom HTTPClient for a request.
func NewTRAAPILiveTrainDelay2154ParamsWithHTTPClient(client *http.Client) *TRAAPILiveTrainDelay2154Params {
	return &TRAAPILiveTrainDelay2154Params{
		HTTPClient: client,
	}
}

/* TRAAPILiveTrainDelay2154Params contains all the parameters to send to the API endpoint
   for the t r a Api live train delay 2154 operation.

   Typically these are written to a http.Request.
*/
type TRAAPILiveTrainDelay2154Params struct {

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

// WithDefaults hydrates default values in the t r a Api live train delay 2154 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TRAAPILiveTrainDelay2154Params) WithDefaults() *TRAAPILiveTrainDelay2154Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the t r a Api live train delay 2154 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TRAAPILiveTrainDelay2154Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := TRAAPILiveTrainDelay2154Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) WithTimeout(timeout time.Duration) *TRAAPILiveTrainDelay2154Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) WithContext(ctx context.Context) *TRAAPILiveTrainDelay2154Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) WithHTTPClient(client *http.Client) *TRAAPILiveTrainDelay2154Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) WithDollarFilter(dollarFilter *string) *TRAAPILiveTrainDelay2154Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) WithDollarFormat(dollarFormat string) *TRAAPILiveTrainDelay2154Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) WithDollarOrderby(dollarOrderby *string) *TRAAPILiveTrainDelay2154Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) WithDollarSelect(dollarSelect *string) *TRAAPILiveTrainDelay2154Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) WithDollarSkip(dollarSkip *string) *TRAAPILiveTrainDelay2154Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) WithDollarTop(dollarTop *int64) *TRAAPILiveTrainDelay2154Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithHealth adds the health to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) WithHealth(health *string) *TRAAPILiveTrainDelay2154Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the t r a Api live train delay 2154 params
func (o *TRAAPILiveTrainDelay2154Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *TRAAPILiveTrainDelay2154Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
