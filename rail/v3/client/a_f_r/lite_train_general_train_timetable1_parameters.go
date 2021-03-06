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

// NewLiteTrainGeneralTrainTimetable1Params creates a new LiteTrainGeneralTrainTimetable1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLiteTrainGeneralTrainTimetable1Params() *LiteTrainGeneralTrainTimetable1Params {
	return &LiteTrainGeneralTrainTimetable1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewLiteTrainGeneralTrainTimetable1ParamsWithTimeout creates a new LiteTrainGeneralTrainTimetable1Params object
// with the ability to set a timeout on a request.
func NewLiteTrainGeneralTrainTimetable1ParamsWithTimeout(timeout time.Duration) *LiteTrainGeneralTrainTimetable1Params {
	return &LiteTrainGeneralTrainTimetable1Params{
		timeout: timeout,
	}
}

// NewLiteTrainGeneralTrainTimetable1ParamsWithContext creates a new LiteTrainGeneralTrainTimetable1Params object
// with the ability to set a context for a request.
func NewLiteTrainGeneralTrainTimetable1ParamsWithContext(ctx context.Context) *LiteTrainGeneralTrainTimetable1Params {
	return &LiteTrainGeneralTrainTimetable1Params{
		Context: ctx,
	}
}

// NewLiteTrainGeneralTrainTimetable1ParamsWithHTTPClient creates a new LiteTrainGeneralTrainTimetable1Params object
// with the ability to set a custom HTTPClient for a request.
func NewLiteTrainGeneralTrainTimetable1ParamsWithHTTPClient(client *http.Client) *LiteTrainGeneralTrainTimetable1Params {
	return &LiteTrainGeneralTrainTimetable1Params{
		HTTPClient: client,
	}
}

/* LiteTrainGeneralTrainTimetable1Params contains all the parameters to send to the API endpoint
   for the lite train general train timetable 1 operation.

   Typically these are written to a http.Request.
*/
type LiteTrainGeneralTrainTimetable1Params struct {

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

	/* TrainNo.

	   欲查詢車次的代碼
	*/
	TrainNo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lite train general train timetable 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LiteTrainGeneralTrainTimetable1Params) WithDefaults() *LiteTrainGeneralTrainTimetable1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lite train general train timetable 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LiteTrainGeneralTrainTimetable1Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := LiteTrainGeneralTrainTimetable1Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) WithTimeout(timeout time.Duration) *LiteTrainGeneralTrainTimetable1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) WithContext(ctx context.Context) *LiteTrainGeneralTrainTimetable1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) WithHTTPClient(client *http.Client) *LiteTrainGeneralTrainTimetable1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) WithDollarCount(dollarCount *bool) *LiteTrainGeneralTrainTimetable1Params {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) WithDollarFilter(dollarFilter *string) *LiteTrainGeneralTrainTimetable1Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) WithDollarFormat(dollarFormat string) *LiteTrainGeneralTrainTimetable1Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) WithDollarOrderby(dollarOrderby *string) *LiteTrainGeneralTrainTimetable1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) WithDollarSelect(dollarSelect *string) *LiteTrainGeneralTrainTimetable1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) WithDollarSkip(dollarSkip *string) *LiteTrainGeneralTrainTimetable1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) WithDollarTop(dollarTop *int64) *LiteTrainGeneralTrainTimetable1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithTrainNo adds the trainNo to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) WithTrainNo(trainNo string) *LiteTrainGeneralTrainTimetable1Params {
	o.SetTrainNo(trainNo)
	return o
}

// SetTrainNo adds the trainNo to the lite train general train timetable 1 params
func (o *LiteTrainGeneralTrainTimetable1Params) SetTrainNo(trainNo string) {
	o.TrainNo = trainNo
}

// WriteToRequest writes these params to a swagger request
func (o *LiteTrainGeneralTrainTimetable1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param TrainNo
	if err := r.SetPathParam("TrainNo", o.TrainNo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
