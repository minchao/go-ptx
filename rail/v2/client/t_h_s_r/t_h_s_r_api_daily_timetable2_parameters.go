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

	strfmt "github.com/go-openapi/strfmt"
)

// NewTHSRAPIDailyTimetable2Params creates a new THSRAPIDailyTimetable2Params object
// with the default values initialized.
func NewTHSRAPIDailyTimetable2Params() *THSRAPIDailyTimetable2Params {
	var (
		dollarTopDefault = string("30")
	)
	return &THSRAPIDailyTimetable2Params{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTHSRAPIDailyTimetable2ParamsWithTimeout creates a new THSRAPIDailyTimetable2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewTHSRAPIDailyTimetable2ParamsWithTimeout(timeout time.Duration) *THSRAPIDailyTimetable2Params {
	var (
		dollarTopDefault = string("30")
	)
	return &THSRAPIDailyTimetable2Params{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTHSRAPIDailyTimetable2ParamsWithContext creates a new THSRAPIDailyTimetable2Params object
// with the default values initialized, and the ability to set a context for a request
func NewTHSRAPIDailyTimetable2ParamsWithContext(ctx context.Context) *THSRAPIDailyTimetable2Params {
	var (
		dollarTopDefault = string("30")
	)
	return &THSRAPIDailyTimetable2Params{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTHSRAPIDailyTimetable2ParamsWithHTTPClient creates a new THSRAPIDailyTimetable2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTHSRAPIDailyTimetable2ParamsWithHTTPClient(client *http.Client) *THSRAPIDailyTimetable2Params {
	var (
		dollarTopDefault = string("30")
	)
	return &THSRAPIDailyTimetable2Params{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*THSRAPIDailyTimetable2Params contains all the parameters to send to the API endpoint
for the t h s r Api daily timetable 2 operation typically these are written to a http.Request
*/
type THSRAPIDailyTimetable2Params struct {

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
	DollarTop *string
	/*TrainDate
	  欲查詢的日期(格式: yyyy-MM-dd)

	*/
	TrainDate string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) WithTimeout(timeout time.Duration) *THSRAPIDailyTimetable2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) WithContext(ctx context.Context) *THSRAPIDailyTimetable2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) WithHTTPClient(client *http.Client) *THSRAPIDailyTimetable2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) WithDollarFilter(dollarFilter *string) *THSRAPIDailyTimetable2Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) WithDollarFormat(dollarFormat string) *THSRAPIDailyTimetable2Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) WithDollarOrderby(dollarOrderby *string) *THSRAPIDailyTimetable2Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) WithDollarSelect(dollarSelect *string) *THSRAPIDailyTimetable2Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) WithDollarSkip(dollarSkip *string) *THSRAPIDailyTimetable2Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) WithDollarTop(dollarTop *string) *THSRAPIDailyTimetable2Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) SetDollarTop(dollarTop *string) {
	o.DollarTop = dollarTop
}

// WithTrainDate adds the trainDate to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) WithTrainDate(trainDate string) *THSRAPIDailyTimetable2Params {
	o.SetTrainDate(trainDate)
	return o
}

// SetTrainDate adds the trainDate to the t h s r Api daily timetable 2 params
func (o *THSRAPIDailyTimetable2Params) SetTrainDate(trainDate string) {
	o.TrainDate = trainDate
}

// WriteToRequest writes these params to a swagger request
func (o *THSRAPIDailyTimetable2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
		var qrDollarTop string
		if o.DollarTop != nil {
			qrDollarTop = *o.DollarTop
		}
		qDollarTop := qrDollarTop
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
