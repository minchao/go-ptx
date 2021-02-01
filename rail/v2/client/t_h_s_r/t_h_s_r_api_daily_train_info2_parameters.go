// Code generated by go-swagger; DO NOT EDIT.

package t_h_s_r

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

// NewTHSRAPIDailyTrainInfo2Params creates a new THSRAPIDailyTrainInfo2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTHSRAPIDailyTrainInfo2Params() *THSRAPIDailyTrainInfo2Params {
	return &THSRAPIDailyTrainInfo2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewTHSRAPIDailyTrainInfo2ParamsWithTimeout creates a new THSRAPIDailyTrainInfo2Params object
// with the ability to set a timeout on a request.
func NewTHSRAPIDailyTrainInfo2ParamsWithTimeout(timeout time.Duration) *THSRAPIDailyTrainInfo2Params {
	return &THSRAPIDailyTrainInfo2Params{
		timeout: timeout,
	}
}

// NewTHSRAPIDailyTrainInfo2ParamsWithContext creates a new THSRAPIDailyTrainInfo2Params object
// with the ability to set a context for a request.
func NewTHSRAPIDailyTrainInfo2ParamsWithContext(ctx context.Context) *THSRAPIDailyTrainInfo2Params {
	return &THSRAPIDailyTrainInfo2Params{
		Context: ctx,
	}
}

// NewTHSRAPIDailyTrainInfo2ParamsWithHTTPClient creates a new THSRAPIDailyTrainInfo2Params object
// with the ability to set a custom HTTPClient for a request.
func NewTHSRAPIDailyTrainInfo2ParamsWithHTTPClient(client *http.Client) *THSRAPIDailyTrainInfo2Params {
	return &THSRAPIDailyTrainInfo2Params{
		HTTPClient: client,
	}
}

/* THSRAPIDailyTrainInfo2Params contains all the parameters to send to the API endpoint
   for the t h s r Api daily train info 2 operation.

   Typically these are written to a http.Request.
*/
type THSRAPIDailyTrainInfo2Params struct {

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

	/* TrainDate.

	   欲查詢車次的日期(格式: yyyy-MM-dd)
	*/
	TrainDate string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the t h s r Api daily train info 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *THSRAPIDailyTrainInfo2Params) WithDefaults() *THSRAPIDailyTrainInfo2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the t h s r Api daily train info 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *THSRAPIDailyTrainInfo2Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := THSRAPIDailyTrainInfo2Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) WithTimeout(timeout time.Duration) *THSRAPIDailyTrainInfo2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) WithContext(ctx context.Context) *THSRAPIDailyTrainInfo2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) WithHTTPClient(client *http.Client) *THSRAPIDailyTrainInfo2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) WithDollarFilter(dollarFilter *string) *THSRAPIDailyTrainInfo2Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) WithDollarFormat(dollarFormat string) *THSRAPIDailyTrainInfo2Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) WithDollarOrderby(dollarOrderby *string) *THSRAPIDailyTrainInfo2Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) WithDollarSelect(dollarSelect *string) *THSRAPIDailyTrainInfo2Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) WithDollarSkip(dollarSkip *string) *THSRAPIDailyTrainInfo2Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) WithDollarTop(dollarTop *int64) *THSRAPIDailyTrainInfo2Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithTrainDate adds the trainDate to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) WithTrainDate(trainDate string) *THSRAPIDailyTrainInfo2Params {
	o.SetTrainDate(trainDate)
	return o
}

// SetTrainDate adds the trainDate to the t h s r Api daily train info 2 params
func (o *THSRAPIDailyTrainInfo2Params) SetTrainDate(trainDate string) {
	o.TrainDate = trainDate
}

// WriteToRequest writes these params to a swagger request
func (o *THSRAPIDailyTrainInfo2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param TrainDate
	if err := r.SetPathParam("TrainDate", o.TrainDate); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
