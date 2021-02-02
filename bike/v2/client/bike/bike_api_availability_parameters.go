// Code generated by go-swagger; DO NOT EDIT.

package bike

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

// NewBikeAPIAvailabilityParams creates a new BikeAPIAvailabilityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBikeAPIAvailabilityParams() *BikeAPIAvailabilityParams {
	return &BikeAPIAvailabilityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBikeAPIAvailabilityParamsWithTimeout creates a new BikeAPIAvailabilityParams object
// with the ability to set a timeout on a request.
func NewBikeAPIAvailabilityParamsWithTimeout(timeout time.Duration) *BikeAPIAvailabilityParams {
	return &BikeAPIAvailabilityParams{
		timeout: timeout,
	}
}

// NewBikeAPIAvailabilityParamsWithContext creates a new BikeAPIAvailabilityParams object
// with the ability to set a context for a request.
func NewBikeAPIAvailabilityParamsWithContext(ctx context.Context) *BikeAPIAvailabilityParams {
	return &BikeAPIAvailabilityParams{
		Context: ctx,
	}
}

// NewBikeAPIAvailabilityParamsWithHTTPClient creates a new BikeAPIAvailabilityParams object
// with the ability to set a custom HTTPClient for a request.
func NewBikeAPIAvailabilityParamsWithHTTPClient(client *http.Client) *BikeAPIAvailabilityParams {
	return &BikeAPIAvailabilityParams{
		HTTPClient: client,
	}
}

/* BikeAPIAvailabilityParams contains all the parameters to send to the API endpoint
   for the bike Api availability operation.

   Typically these are written to a http.Request.
*/
type BikeAPIAvailabilityParams struct {

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

	/* City.

	   欲查詢縣市
	*/
	City string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bike Api availability params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BikeAPIAvailabilityParams) WithDefaults() *BikeAPIAvailabilityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bike Api availability params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BikeAPIAvailabilityParams) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := BikeAPIAvailabilityParams{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the bike Api availability params
func (o *BikeAPIAvailabilityParams) WithTimeout(timeout time.Duration) *BikeAPIAvailabilityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bike Api availability params
func (o *BikeAPIAvailabilityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bike Api availability params
func (o *BikeAPIAvailabilityParams) WithContext(ctx context.Context) *BikeAPIAvailabilityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bike Api availability params
func (o *BikeAPIAvailabilityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bike Api availability params
func (o *BikeAPIAvailabilityParams) WithHTTPClient(client *http.Client) *BikeAPIAvailabilityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bike Api availability params
func (o *BikeAPIAvailabilityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the bike Api availability params
func (o *BikeAPIAvailabilityParams) WithDollarFilter(dollarFilter *string) *BikeAPIAvailabilityParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the bike Api availability params
func (o *BikeAPIAvailabilityParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the bike Api availability params
func (o *BikeAPIAvailabilityParams) WithDollarFormat(dollarFormat string) *BikeAPIAvailabilityParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the bike Api availability params
func (o *BikeAPIAvailabilityParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the bike Api availability params
func (o *BikeAPIAvailabilityParams) WithDollarOrderby(dollarOrderby *string) *BikeAPIAvailabilityParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the bike Api availability params
func (o *BikeAPIAvailabilityParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the bike Api availability params
func (o *BikeAPIAvailabilityParams) WithDollarSelect(dollarSelect *string) *BikeAPIAvailabilityParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the bike Api availability params
func (o *BikeAPIAvailabilityParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the bike Api availability params
func (o *BikeAPIAvailabilityParams) WithDollarSkip(dollarSkip *string) *BikeAPIAvailabilityParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the bike Api availability params
func (o *BikeAPIAvailabilityParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the bike Api availability params
func (o *BikeAPIAvailabilityParams) WithDollarTop(dollarTop *int64) *BikeAPIAvailabilityParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the bike Api availability params
func (o *BikeAPIAvailabilityParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithCity adds the city to the bike Api availability params
func (o *BikeAPIAvailabilityParams) WithCity(city string) *BikeAPIAvailabilityParams {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the bike Api availability params
func (o *BikeAPIAvailabilityParams) SetCity(city string) {
	o.City = city
}

// WriteToRequest writes these params to a swagger request
func (o *BikeAPIAvailabilityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param City
	if err := r.SetPathParam("City", o.City); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
