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

// NewAirAPIFlightParams creates a new AirAPIFlightParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAirAPIFlightParams() *AirAPIFlightParams {
	return &AirAPIFlightParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAirAPIFlightParamsWithTimeout creates a new AirAPIFlightParams object
// with the ability to set a timeout on a request.
func NewAirAPIFlightParamsWithTimeout(timeout time.Duration) *AirAPIFlightParams {
	return &AirAPIFlightParams{
		timeout: timeout,
	}
}

// NewAirAPIFlightParamsWithContext creates a new AirAPIFlightParams object
// with the ability to set a context for a request.
func NewAirAPIFlightParamsWithContext(ctx context.Context) *AirAPIFlightParams {
	return &AirAPIFlightParams{
		Context: ctx,
	}
}

// NewAirAPIFlightParamsWithHTTPClient creates a new AirAPIFlightParams object
// with the ability to set a custom HTTPClient for a request.
func NewAirAPIFlightParamsWithHTTPClient(client *http.Client) *AirAPIFlightParams {
	return &AirAPIFlightParams{
		HTTPClient: client,
	}
}

/* AirAPIFlightParams contains all the parameters to send to the API endpoint
   for the air Api flight operation.

   Typically these are written to a http.Request.
*/
type AirAPIFlightParams struct {

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

// WithDefaults hydrates default values in the air Api flight params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AirAPIFlightParams) WithDefaults() *AirAPIFlightParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the air Api flight params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AirAPIFlightParams) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := AirAPIFlightParams{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the air Api flight params
func (o *AirAPIFlightParams) WithTimeout(timeout time.Duration) *AirAPIFlightParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the air Api flight params
func (o *AirAPIFlightParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the air Api flight params
func (o *AirAPIFlightParams) WithContext(ctx context.Context) *AirAPIFlightParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the air Api flight params
func (o *AirAPIFlightParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the air Api flight params
func (o *AirAPIFlightParams) WithHTTPClient(client *http.Client) *AirAPIFlightParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the air Api flight params
func (o *AirAPIFlightParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the air Api flight params
func (o *AirAPIFlightParams) WithDollarFilter(dollarFilter *string) *AirAPIFlightParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the air Api flight params
func (o *AirAPIFlightParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the air Api flight params
func (o *AirAPIFlightParams) WithDollarFormat(dollarFormat string) *AirAPIFlightParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the air Api flight params
func (o *AirAPIFlightParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the air Api flight params
func (o *AirAPIFlightParams) WithDollarOrderby(dollarOrderby *string) *AirAPIFlightParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the air Api flight params
func (o *AirAPIFlightParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the air Api flight params
func (o *AirAPIFlightParams) WithDollarSelect(dollarSelect *string) *AirAPIFlightParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the air Api flight params
func (o *AirAPIFlightParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the air Api flight params
func (o *AirAPIFlightParams) WithDollarSkip(dollarSkip *string) *AirAPIFlightParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the air Api flight params
func (o *AirAPIFlightParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the air Api flight params
func (o *AirAPIFlightParams) WithDollarTop(dollarTop *int64) *AirAPIFlightParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the air Api flight params
func (o *AirAPIFlightParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithIsCargo adds the isCargo to the air Api flight params
func (o *AirAPIFlightParams) WithIsCargo(isCargo *bool) *AirAPIFlightParams {
	o.SetIsCargo(isCargo)
	return o
}

// SetIsCargo adds the isCargo to the air Api flight params
func (o *AirAPIFlightParams) SetIsCargo(isCargo *bool) {
	o.IsCargo = isCargo
}

// WriteToRequest writes these params to a swagger request
func (o *AirAPIFlightParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
