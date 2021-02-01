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

// NewStationLiveBoardAPIControllerGetParams creates a new StationLiveBoardAPIControllerGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStationLiveBoardAPIControllerGetParams() *StationLiveBoardAPIControllerGetParams {
	return &StationLiveBoardAPIControllerGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStationLiveBoardAPIControllerGetParamsWithTimeout creates a new StationLiveBoardAPIControllerGetParams object
// with the ability to set a timeout on a request.
func NewStationLiveBoardAPIControllerGetParamsWithTimeout(timeout time.Duration) *StationLiveBoardAPIControllerGetParams {
	return &StationLiveBoardAPIControllerGetParams{
		timeout: timeout,
	}
}

// NewStationLiveBoardAPIControllerGetParamsWithContext creates a new StationLiveBoardAPIControllerGetParams object
// with the ability to set a context for a request.
func NewStationLiveBoardAPIControllerGetParamsWithContext(ctx context.Context) *StationLiveBoardAPIControllerGetParams {
	return &StationLiveBoardAPIControllerGetParams{
		Context: ctx,
	}
}

// NewStationLiveBoardAPIControllerGetParamsWithHTTPClient creates a new StationLiveBoardAPIControllerGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewStationLiveBoardAPIControllerGetParamsWithHTTPClient(client *http.Client) *StationLiveBoardAPIControllerGetParams {
	return &StationLiveBoardAPIControllerGetParams{
		HTTPClient: client,
	}
}

/* StationLiveBoardAPIControllerGetParams contains all the parameters to send to the API endpoint
   for the station live board Api controller get operation.

   Typically these are written to a http.Request.
*/
type StationLiveBoardAPIControllerGetParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the station live board Api controller get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StationLiveBoardAPIControllerGetParams) WithDefaults() *StationLiveBoardAPIControllerGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the station live board Api controller get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StationLiveBoardAPIControllerGetParams) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := StationLiveBoardAPIControllerGetParams{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) WithTimeout(timeout time.Duration) *StationLiveBoardAPIControllerGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) WithContext(ctx context.Context) *StationLiveBoardAPIControllerGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) WithHTTPClient(client *http.Client) *StationLiveBoardAPIControllerGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) WithDollarCount(dollarCount *bool) *StationLiveBoardAPIControllerGetParams {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) WithDollarFilter(dollarFilter *string) *StationLiveBoardAPIControllerGetParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) WithDollarFormat(dollarFormat string) *StationLiveBoardAPIControllerGetParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) WithDollarOrderby(dollarOrderby *string) *StationLiveBoardAPIControllerGetParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) WithDollarSelect(dollarSelect *string) *StationLiveBoardAPIControllerGetParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) WithDollarSkip(dollarSkip *string) *StationLiveBoardAPIControllerGetParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) WithDollarTop(dollarTop *int64) *StationLiveBoardAPIControllerGetParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the station live board Api controller get params
func (o *StationLiveBoardAPIControllerGetParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *StationLiveBoardAPIControllerGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
