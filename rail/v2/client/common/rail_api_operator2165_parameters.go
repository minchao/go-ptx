// Code generated by go-swagger; DO NOT EDIT.

package common

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

// NewRailAPIOperator2165Params creates a new RailAPIOperator2165Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRailAPIOperator2165Params() *RailAPIOperator2165Params {
	return &RailAPIOperator2165Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewRailAPIOperator2165ParamsWithTimeout creates a new RailAPIOperator2165Params object
// with the ability to set a timeout on a request.
func NewRailAPIOperator2165ParamsWithTimeout(timeout time.Duration) *RailAPIOperator2165Params {
	return &RailAPIOperator2165Params{
		timeout: timeout,
	}
}

// NewRailAPIOperator2165ParamsWithContext creates a new RailAPIOperator2165Params object
// with the ability to set a context for a request.
func NewRailAPIOperator2165ParamsWithContext(ctx context.Context) *RailAPIOperator2165Params {
	return &RailAPIOperator2165Params{
		Context: ctx,
	}
}

// NewRailAPIOperator2165ParamsWithHTTPClient creates a new RailAPIOperator2165Params object
// with the ability to set a custom HTTPClient for a request.
func NewRailAPIOperator2165ParamsWithHTTPClient(client *http.Client) *RailAPIOperator2165Params {
	return &RailAPIOperator2165Params{
		HTTPClient: client,
	}
}

/* RailAPIOperator2165Params contains all the parameters to send to the API endpoint
   for the rail Api operator 2165 operation.

   Typically these are written to a http.Request.
*/
type RailAPIOperator2165Params struct {

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

// WithDefaults hydrates default values in the rail Api operator 2165 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RailAPIOperator2165Params) WithDefaults() *RailAPIOperator2165Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rail Api operator 2165 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RailAPIOperator2165Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := RailAPIOperator2165Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) WithTimeout(timeout time.Duration) *RailAPIOperator2165Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) WithContext(ctx context.Context) *RailAPIOperator2165Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) WithHTTPClient(client *http.Client) *RailAPIOperator2165Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) WithDollarFilter(dollarFilter *string) *RailAPIOperator2165Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) WithDollarFormat(dollarFormat string) *RailAPIOperator2165Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) WithDollarOrderby(dollarOrderby *string) *RailAPIOperator2165Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) WithDollarSelect(dollarSelect *string) *RailAPIOperator2165Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) WithDollarSkip(dollarSkip *string) *RailAPIOperator2165Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) WithDollarTop(dollarTop *int64) *RailAPIOperator2165Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithHealth adds the health to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) WithHealth(health *string) *RailAPIOperator2165Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the rail Api operator 2165 params
func (o *RailAPIOperator2165Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *RailAPIOperator2165Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
