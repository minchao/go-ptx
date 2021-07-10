// Code generated by go-swagger; DO NOT EDIT.

package a_f_r

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

// NewLiteTrainODFare1Params creates a new LiteTrainODFare1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLiteTrainODFare1Params() *LiteTrainODFare1Params {
	return &LiteTrainODFare1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewLiteTrainODFare1ParamsWithTimeout creates a new LiteTrainODFare1Params object
// with the ability to set a timeout on a request.
func NewLiteTrainODFare1ParamsWithTimeout(timeout time.Duration) *LiteTrainODFare1Params {
	return &LiteTrainODFare1Params{
		timeout: timeout,
	}
}

// NewLiteTrainODFare1ParamsWithContext creates a new LiteTrainODFare1Params object
// with the ability to set a context for a request.
func NewLiteTrainODFare1ParamsWithContext(ctx context.Context) *LiteTrainODFare1Params {
	return &LiteTrainODFare1Params{
		Context: ctx,
	}
}

// NewLiteTrainODFare1ParamsWithHTTPClient creates a new LiteTrainODFare1Params object
// with the ability to set a custom HTTPClient for a request.
func NewLiteTrainODFare1ParamsWithHTTPClient(client *http.Client) *LiteTrainODFare1Params {
	return &LiteTrainODFare1Params{
		HTTPClient: client,
	}
}

/* LiteTrainODFare1Params contains all the parameters to send to the API endpoint
   for the lite train o d fare 1 operation.

   Typically these are written to a http.Request.
*/
type LiteTrainODFare1Params struct {

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

	/* DestinationStationID.

	   迄點車站代碼
	*/
	DestinationStationID string

	/* OriginStationID.

	   起點車站代碼
	*/
	OriginStationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lite train o d fare 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LiteTrainODFare1Params) WithDefaults() *LiteTrainODFare1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lite train o d fare 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LiteTrainODFare1Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := LiteTrainODFare1Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) WithTimeout(timeout time.Duration) *LiteTrainODFare1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) WithContext(ctx context.Context) *LiteTrainODFare1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) WithHTTPClient(client *http.Client) *LiteTrainODFare1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) WithDollarCount(dollarCount *bool) *LiteTrainODFare1Params {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) WithDollarFilter(dollarFilter *string) *LiteTrainODFare1Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) WithDollarFormat(dollarFormat string) *LiteTrainODFare1Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) WithDollarOrderby(dollarOrderby *string) *LiteTrainODFare1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) WithDollarSelect(dollarSelect *string) *LiteTrainODFare1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) WithDollarSkip(dollarSkip *string) *LiteTrainODFare1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) WithDollarTop(dollarTop *int64) *LiteTrainODFare1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithDestinationStationID adds the destinationStationID to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) WithDestinationStationID(destinationStationID string) *LiteTrainODFare1Params {
	o.SetDestinationStationID(destinationStationID)
	return o
}

// SetDestinationStationID adds the destinationStationId to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) SetDestinationStationID(destinationStationID string) {
	o.DestinationStationID = destinationStationID
}

// WithOriginStationID adds the originStationID to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) WithOriginStationID(originStationID string) *LiteTrainODFare1Params {
	o.SetOriginStationID(originStationID)
	return o
}

// SetOriginStationID adds the originStationId to the lite train o d fare 1 params
func (o *LiteTrainODFare1Params) SetOriginStationID(originStationID string) {
	o.OriginStationID = originStationID
}

// WriteToRequest writes these params to a swagger request
func (o *LiteTrainODFare1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param DestinationStationID
	if err := r.SetPathParam("DestinationStationID", o.DestinationStationID); err != nil {
		return err
	}

	// path param OriginStationID
	if err := r.SetPathParam("OriginStationID", o.OriginStationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
