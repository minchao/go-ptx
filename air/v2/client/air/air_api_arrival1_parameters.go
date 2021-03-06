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

// NewAirAPIArrival1Params creates a new AirAPIArrival1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAirAPIArrival1Params() *AirAPIArrival1Params {
	return &AirAPIArrival1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewAirAPIArrival1ParamsWithTimeout creates a new AirAPIArrival1Params object
// with the ability to set a timeout on a request.
func NewAirAPIArrival1ParamsWithTimeout(timeout time.Duration) *AirAPIArrival1Params {
	return &AirAPIArrival1Params{
		timeout: timeout,
	}
}

// NewAirAPIArrival1ParamsWithContext creates a new AirAPIArrival1Params object
// with the ability to set a context for a request.
func NewAirAPIArrival1ParamsWithContext(ctx context.Context) *AirAPIArrival1Params {
	return &AirAPIArrival1Params{
		Context: ctx,
	}
}

// NewAirAPIArrival1ParamsWithHTTPClient creates a new AirAPIArrival1Params object
// with the ability to set a custom HTTPClient for a request.
func NewAirAPIArrival1ParamsWithHTTPClient(client *http.Client) *AirAPIArrival1Params {
	return &AirAPIArrival1Params{
		HTTPClient: client,
	}
}

/* AirAPIArrival1Params contains all the parameters to send to the API endpoint
   for the air Api arrival 1 operation.

   Typically these are written to a http.Request.
*/
type AirAPIArrival1Params struct {

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

	/* IATA.

	   機場代碼
	*/
	IATA string

	/* IsCargo.

	   是否為貨機
	*/
	IsCargo *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the air Api arrival 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AirAPIArrival1Params) WithDefaults() *AirAPIArrival1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the air Api arrival 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AirAPIArrival1Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := AirAPIArrival1Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the air Api arrival 1 params
func (o *AirAPIArrival1Params) WithTimeout(timeout time.Duration) *AirAPIArrival1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the air Api arrival 1 params
func (o *AirAPIArrival1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the air Api arrival 1 params
func (o *AirAPIArrival1Params) WithContext(ctx context.Context) *AirAPIArrival1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the air Api arrival 1 params
func (o *AirAPIArrival1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the air Api arrival 1 params
func (o *AirAPIArrival1Params) WithHTTPClient(client *http.Client) *AirAPIArrival1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the air Api arrival 1 params
func (o *AirAPIArrival1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the air Api arrival 1 params
func (o *AirAPIArrival1Params) WithDollarFilter(dollarFilter *string) *AirAPIArrival1Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the air Api arrival 1 params
func (o *AirAPIArrival1Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the air Api arrival 1 params
func (o *AirAPIArrival1Params) WithDollarFormat(dollarFormat string) *AirAPIArrival1Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the air Api arrival 1 params
func (o *AirAPIArrival1Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the air Api arrival 1 params
func (o *AirAPIArrival1Params) WithDollarOrderby(dollarOrderby *string) *AirAPIArrival1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the air Api arrival 1 params
func (o *AirAPIArrival1Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the air Api arrival 1 params
func (o *AirAPIArrival1Params) WithDollarSelect(dollarSelect *string) *AirAPIArrival1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the air Api arrival 1 params
func (o *AirAPIArrival1Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the air Api arrival 1 params
func (o *AirAPIArrival1Params) WithDollarSkip(dollarSkip *string) *AirAPIArrival1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the air Api arrival 1 params
func (o *AirAPIArrival1Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the air Api arrival 1 params
func (o *AirAPIArrival1Params) WithDollarTop(dollarTop *int64) *AirAPIArrival1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the air Api arrival 1 params
func (o *AirAPIArrival1Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithIATA adds the iATA to the air Api arrival 1 params
func (o *AirAPIArrival1Params) WithIATA(iATA string) *AirAPIArrival1Params {
	o.SetIATA(iATA)
	return o
}

// SetIATA adds the iATA to the air Api arrival 1 params
func (o *AirAPIArrival1Params) SetIATA(iATA string) {
	o.IATA = iATA
}

// WithIsCargo adds the isCargo to the air Api arrival 1 params
func (o *AirAPIArrival1Params) WithIsCargo(isCargo *bool) *AirAPIArrival1Params {
	o.SetIsCargo(isCargo)
	return o
}

// SetIsCargo adds the isCargo to the air Api arrival 1 params
func (o *AirAPIArrival1Params) SetIsCargo(isCargo *bool) {
	o.IsCargo = isCargo
}

// WriteToRequest writes these params to a swagger request
func (o *AirAPIArrival1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param IATA
	if err := r.SetPathParam("IATA", o.IATA); err != nil {
		return err
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
