// Code generated by go-swagger; DO NOT EDIT.

package air

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

// NewAirAPIArrivalParams creates a new AirAPIArrivalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAirAPIArrivalParams() *AirAPIArrivalParams {
	return &AirAPIArrivalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAirAPIArrivalParamsWithTimeout creates a new AirAPIArrivalParams object
// with the ability to set a timeout on a request.
func NewAirAPIArrivalParamsWithTimeout(timeout time.Duration) *AirAPIArrivalParams {
	return &AirAPIArrivalParams{
		timeout: timeout,
	}
}

// NewAirAPIArrivalParamsWithContext creates a new AirAPIArrivalParams object
// with the ability to set a context for a request.
func NewAirAPIArrivalParamsWithContext(ctx context.Context) *AirAPIArrivalParams {
	return &AirAPIArrivalParams{
		Context: ctx,
	}
}

// NewAirAPIArrivalParamsWithHTTPClient creates a new AirAPIArrivalParams object
// with the ability to set a custom HTTPClient for a request.
func NewAirAPIArrivalParamsWithHTTPClient(client *http.Client) *AirAPIArrivalParams {
	return &AirAPIArrivalParams{
		HTTPClient: client,
	}
}

/* AirAPIArrivalParams contains all the parameters to send to the API endpoint
   for the air Api arrival operation.

   Typically these are written to a http.Request.
*/
type AirAPIArrivalParams struct {

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

	/* IsCargo.

	   是否為貨機
	*/
	IsCargo *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the air Api arrival params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AirAPIArrivalParams) WithDefaults() *AirAPIArrivalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the air Api arrival params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AirAPIArrivalParams) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := AirAPIArrivalParams{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the air Api arrival params
func (o *AirAPIArrivalParams) WithTimeout(timeout time.Duration) *AirAPIArrivalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the air Api arrival params
func (o *AirAPIArrivalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the air Api arrival params
func (o *AirAPIArrivalParams) WithContext(ctx context.Context) *AirAPIArrivalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the air Api arrival params
func (o *AirAPIArrivalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the air Api arrival params
func (o *AirAPIArrivalParams) WithHTTPClient(client *http.Client) *AirAPIArrivalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the air Api arrival params
func (o *AirAPIArrivalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the air Api arrival params
func (o *AirAPIArrivalParams) WithDollarFilter(dollarFilter *string) *AirAPIArrivalParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the air Api arrival params
func (o *AirAPIArrivalParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the air Api arrival params
func (o *AirAPIArrivalParams) WithDollarFormat(dollarFormat string) *AirAPIArrivalParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the air Api arrival params
func (o *AirAPIArrivalParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the air Api arrival params
func (o *AirAPIArrivalParams) WithDollarOrderby(dollarOrderby *string) *AirAPIArrivalParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the air Api arrival params
func (o *AirAPIArrivalParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the air Api arrival params
func (o *AirAPIArrivalParams) WithDollarSelect(dollarSelect *string) *AirAPIArrivalParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the air Api arrival params
func (o *AirAPIArrivalParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the air Api arrival params
func (o *AirAPIArrivalParams) WithDollarSkip(dollarSkip *string) *AirAPIArrivalParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the air Api arrival params
func (o *AirAPIArrivalParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the air Api arrival params
func (o *AirAPIArrivalParams) WithDollarTop(dollarTop *int64) *AirAPIArrivalParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the air Api arrival params
func (o *AirAPIArrivalParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithIsCargo adds the isCargo to the air Api arrival params
func (o *AirAPIArrivalParams) WithIsCargo(isCargo *bool) *AirAPIArrivalParams {
	o.SetIsCargo(isCargo)
	return o
}

// SetIsCargo adds the isCargo to the air Api arrival params
func (o *AirAPIArrivalParams) SetIsCargo(isCargo *bool) {
	o.IsCargo = isCargo
}

// WriteToRequest writes these params to a swagger request
func (o *AirAPIArrivalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IsCargo != nil {

		// query param IsCargo
		var qrIsCargo bool

		if o.IsCargo != nil {
			qrIsCargo = *o.IsCargo
		}
		qIsCargo := swag.FormatBool(qrIsCargo)
		if qIsCargo != "" {

			if err := r.SetQueryParam("IsCargo", qIsCargo); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
