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

// NewTHSRAPIDailyTrainInfo1Params creates a new THSRAPIDailyTrainInfo1Params object
// with the default values initialized.
func NewTHSRAPIDailyTrainInfo1Params() *THSRAPIDailyTrainInfo1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &THSRAPIDailyTrainInfo1Params{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTHSRAPIDailyTrainInfo1ParamsWithTimeout creates a new THSRAPIDailyTrainInfo1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewTHSRAPIDailyTrainInfo1ParamsWithTimeout(timeout time.Duration) *THSRAPIDailyTrainInfo1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &THSRAPIDailyTrainInfo1Params{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTHSRAPIDailyTrainInfo1ParamsWithContext creates a new THSRAPIDailyTrainInfo1Params object
// with the default values initialized, and the ability to set a context for a request
func NewTHSRAPIDailyTrainInfo1ParamsWithContext(ctx context.Context) *THSRAPIDailyTrainInfo1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &THSRAPIDailyTrainInfo1Params{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTHSRAPIDailyTrainInfo1ParamsWithHTTPClient creates a new THSRAPIDailyTrainInfo1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTHSRAPIDailyTrainInfo1ParamsWithHTTPClient(client *http.Client) *THSRAPIDailyTrainInfo1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &THSRAPIDailyTrainInfo1Params{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*THSRAPIDailyTrainInfo1Params contains all the parameters to send to the API endpoint
for the t h s r Api daily train info 1 operation typically these are written to a http.Request
*/
type THSRAPIDailyTrainInfo1Params struct {

	/*DollarFilter
	  過濾

	*/
	DollarFilter *string
	/*DollarFormat
	  指定來源格式

	*/
	DollarFormat string
	/*DollarOrderby
	  排序

	*/
	DollarOrderby *string
	/*DollarSelect
	  挑選

	*/
	DollarSelect *string
	/*DollarSkip
	  跳過前幾筆

	*/
	DollarSkip *string
	/*DollarTop
	  取前幾筆

	*/
	DollarTop *int64
	/*TrainNo
	  欲查詢車次的代碼

	*/
	TrainNo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) WithTimeout(timeout time.Duration) *THSRAPIDailyTrainInfo1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) WithContext(ctx context.Context) *THSRAPIDailyTrainInfo1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) WithHTTPClient(client *http.Client) *THSRAPIDailyTrainInfo1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) WithDollarFilter(dollarFilter *string) *THSRAPIDailyTrainInfo1Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) WithDollarFormat(dollarFormat string) *THSRAPIDailyTrainInfo1Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) WithDollarOrderby(dollarOrderby *string) *THSRAPIDailyTrainInfo1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) WithDollarSelect(dollarSelect *string) *THSRAPIDailyTrainInfo1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) WithDollarSkip(dollarSkip *string) *THSRAPIDailyTrainInfo1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) WithDollarTop(dollarTop *int64) *THSRAPIDailyTrainInfo1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithTrainNo adds the trainNo to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) WithTrainNo(trainNo string) *THSRAPIDailyTrainInfo1Params {
	o.SetTrainNo(trainNo)
	return o
}

// SetTrainNo adds the trainNo to the t h s r Api daily train info 1 params
func (o *THSRAPIDailyTrainInfo1Params) SetTrainNo(trainNo string) {
	o.TrainNo = trainNo
}

// WriteToRequest writes these params to a swagger request
func (o *THSRAPIDailyTrainInfo1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param TrainNo
	if err := r.SetPathParam("TrainNo", o.TrainNo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
