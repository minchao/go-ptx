// Code generated by go-swagger; DO NOT EDIT.

package bus_advanced_by_operator

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

// NewBusAPIStationByOperator2873Params creates a new BusAPIStationByOperator2873Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBusAPIStationByOperator2873Params() *BusAPIStationByOperator2873Params {
	return &BusAPIStationByOperator2873Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewBusAPIStationByOperator2873ParamsWithTimeout creates a new BusAPIStationByOperator2873Params object
// with the ability to set a timeout on a request.
func NewBusAPIStationByOperator2873ParamsWithTimeout(timeout time.Duration) *BusAPIStationByOperator2873Params {
	return &BusAPIStationByOperator2873Params{
		timeout: timeout,
	}
}

// NewBusAPIStationByOperator2873ParamsWithContext creates a new BusAPIStationByOperator2873Params object
// with the ability to set a context for a request.
func NewBusAPIStationByOperator2873ParamsWithContext(ctx context.Context) *BusAPIStationByOperator2873Params {
	return &BusAPIStationByOperator2873Params{
		Context: ctx,
	}
}

// NewBusAPIStationByOperator2873ParamsWithHTTPClient creates a new BusAPIStationByOperator2873Params object
// with the ability to set a custom HTTPClient for a request.
func NewBusAPIStationByOperator2873ParamsWithHTTPClient(client *http.Client) *BusAPIStationByOperator2873Params {
	return &BusAPIStationByOperator2873Params{
		HTTPClient: client,
	}
}

/* BusAPIStationByOperator2873Params contains all the parameters to send to the API endpoint
   for the bus Api station by operator 2873 operation.

   Typically these are written to a http.Request.
*/
type BusAPIStationByOperator2873Params struct {

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

	/* OperatorNo.

	   欲查詢營運業者編號
	*/
	OperatorNo string

	/* Health.

	   加入參數'?health=true'即可查詢此API服務的健康狀態
	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bus Api station by operator 2873 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BusAPIStationByOperator2873Params) WithDefaults() *BusAPIStationByOperator2873Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bus Api station by operator 2873 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BusAPIStationByOperator2873Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := BusAPIStationByOperator2873Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) WithTimeout(timeout time.Duration) *BusAPIStationByOperator2873Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) WithContext(ctx context.Context) *BusAPIStationByOperator2873Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) WithHTTPClient(client *http.Client) *BusAPIStationByOperator2873Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) WithDollarFilter(dollarFilter *string) *BusAPIStationByOperator2873Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) WithDollarFormat(dollarFormat string) *BusAPIStationByOperator2873Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) WithDollarOrderby(dollarOrderby *string) *BusAPIStationByOperator2873Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) WithDollarSelect(dollarSelect *string) *BusAPIStationByOperator2873Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) WithDollarSkip(dollarSkip *string) *BusAPIStationByOperator2873Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) WithDollarTop(dollarTop *int64) *BusAPIStationByOperator2873Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithOperatorNo adds the operatorNo to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) WithOperatorNo(operatorNo string) *BusAPIStationByOperator2873Params {
	o.SetOperatorNo(operatorNo)
	return o
}

// SetOperatorNo adds the operatorNo to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) SetOperatorNo(operatorNo string) {
	o.OperatorNo = operatorNo
}

// WithHealth adds the health to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) WithHealth(health *string) *BusAPIStationByOperator2873Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the bus Api station by operator 2873 params
func (o *BusAPIStationByOperator2873Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *BusAPIStationByOperator2873Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param OperatorNo
	if err := r.SetPathParam("OperatorNo", o.OperatorNo); err != nil {
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
