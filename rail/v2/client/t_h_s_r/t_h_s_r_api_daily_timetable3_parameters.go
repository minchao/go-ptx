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

// NewTHSRAPIDailyTimetable3Params creates a new THSRAPIDailyTimetable3Params object
// with the default values initialized.
func NewTHSRAPIDailyTimetable3Params() *THSRAPIDailyTimetable3Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &THSRAPIDailyTimetable3Params{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTHSRAPIDailyTimetable3ParamsWithTimeout creates a new THSRAPIDailyTimetable3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewTHSRAPIDailyTimetable3ParamsWithTimeout(timeout time.Duration) *THSRAPIDailyTimetable3Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &THSRAPIDailyTimetable3Params{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTHSRAPIDailyTimetable3ParamsWithContext creates a new THSRAPIDailyTimetable3Params object
// with the default values initialized, and the ability to set a context for a request
func NewTHSRAPIDailyTimetable3ParamsWithContext(ctx context.Context) *THSRAPIDailyTimetable3Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &THSRAPIDailyTimetable3Params{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTHSRAPIDailyTimetable3ParamsWithHTTPClient creates a new THSRAPIDailyTimetable3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTHSRAPIDailyTimetable3ParamsWithHTTPClient(client *http.Client) *THSRAPIDailyTimetable3Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &THSRAPIDailyTimetable3Params{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*THSRAPIDailyTimetable3Params contains all the parameters to send to the API endpoint
for the t h s r Api daily timetable 3 operation typically these are written to a http.Request
*/
type THSRAPIDailyTimetable3Params struct {

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
	/*TrainDate
	  欲查詢的日期(格式: yyyy-MM-dd)

	*/
	TrainDate string
	/*TrainNo
	  欲查詢車次的代碼

	*/
	TrainNo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) WithTimeout(timeout time.Duration) *THSRAPIDailyTimetable3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) WithContext(ctx context.Context) *THSRAPIDailyTimetable3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) WithHTTPClient(client *http.Client) *THSRAPIDailyTimetable3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) WithDollarFilter(dollarFilter *string) *THSRAPIDailyTimetable3Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) WithDollarFormat(dollarFormat string) *THSRAPIDailyTimetable3Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) WithDollarOrderby(dollarOrderby *string) *THSRAPIDailyTimetable3Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) WithDollarSelect(dollarSelect *string) *THSRAPIDailyTimetable3Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) WithDollarSkip(dollarSkip *string) *THSRAPIDailyTimetable3Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) WithDollarTop(dollarTop *int64) *THSRAPIDailyTimetable3Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithTrainDate adds the trainDate to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) WithTrainDate(trainDate string) *THSRAPIDailyTimetable3Params {
	o.SetTrainDate(trainDate)
	return o
}

// SetTrainDate adds the trainDate to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) SetTrainDate(trainDate string) {
	o.TrainDate = trainDate
}

// WithTrainNo adds the trainNo to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) WithTrainNo(trainNo string) *THSRAPIDailyTimetable3Params {
	o.SetTrainNo(trainNo)
	return o
}

// SetTrainNo adds the trainNo to the t h s r Api daily timetable 3 params
func (o *THSRAPIDailyTimetable3Params) SetTrainNo(trainNo string) {
	o.TrainNo = trainNo
}

// WriteToRequest writes these params to a swagger request
func (o *THSRAPIDailyTimetable3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param TrainNo
	if err := r.SetPathParam("TrainNo", o.TrainNo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
