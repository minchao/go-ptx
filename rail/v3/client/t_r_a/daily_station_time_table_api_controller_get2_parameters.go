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

// NewDailyStationTimeTableAPIControllerGet2Params creates a new DailyStationTimeTableAPIControllerGet2Params object
// with the default values initialized.
func NewDailyStationTimeTableAPIControllerGet2Params() *DailyStationTimeTableAPIControllerGet2Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &DailyStationTimeTableAPIControllerGet2Params{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDailyStationTimeTableAPIControllerGet2ParamsWithTimeout creates a new DailyStationTimeTableAPIControllerGet2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDailyStationTimeTableAPIControllerGet2ParamsWithTimeout(timeout time.Duration) *DailyStationTimeTableAPIControllerGet2Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &DailyStationTimeTableAPIControllerGet2Params{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewDailyStationTimeTableAPIControllerGet2ParamsWithContext creates a new DailyStationTimeTableAPIControllerGet2Params object
// with the default values initialized, and the ability to set a context for a request
func NewDailyStationTimeTableAPIControllerGet2ParamsWithContext(ctx context.Context) *DailyStationTimeTableAPIControllerGet2Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &DailyStationTimeTableAPIControllerGet2Params{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewDailyStationTimeTableAPIControllerGet2ParamsWithHTTPClient creates a new DailyStationTimeTableAPIControllerGet2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDailyStationTimeTableAPIControllerGet2ParamsWithHTTPClient(client *http.Client) *DailyStationTimeTableAPIControllerGet2Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &DailyStationTimeTableAPIControllerGet2Params{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*DailyStationTimeTableAPIControllerGet2Params contains all the parameters to send to the API endpoint
for the daily station time table Api controller get 2 operation typically these are written to a http.Request
*/
type DailyStationTimeTableAPIControllerGet2Params struct {

	/*DollarCount
	  查詢數量

	*/
	DollarCount *bool
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) WithTimeout(timeout time.Duration) *DailyStationTimeTableAPIControllerGet2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) WithContext(ctx context.Context) *DailyStationTimeTableAPIControllerGet2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) WithHTTPClient(client *http.Client) *DailyStationTimeTableAPIControllerGet2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) WithDollarCount(dollarCount *bool) *DailyStationTimeTableAPIControllerGet2Params {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) WithDollarFilter(dollarFilter *string) *DailyStationTimeTableAPIControllerGet2Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) WithDollarFormat(dollarFormat string) *DailyStationTimeTableAPIControllerGet2Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) WithDollarOrderby(dollarOrderby *string) *DailyStationTimeTableAPIControllerGet2Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) WithDollarSelect(dollarSelect *string) *DailyStationTimeTableAPIControllerGet2Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) WithDollarSkip(dollarSkip *string) *DailyStationTimeTableAPIControllerGet2Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) WithDollarTop(dollarTop *int64) *DailyStationTimeTableAPIControllerGet2Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithTrainDate adds the trainDate to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) WithTrainDate(trainDate string) *DailyStationTimeTableAPIControllerGet2Params {
	o.SetTrainDate(trainDate)
	return o
}

// SetTrainDate adds the trainDate to the daily station time table Api controller get 2 params
func (o *DailyStationTimeTableAPIControllerGet2Params) SetTrainDate(trainDate string) {
	o.TrainDate = trainDate
}

// WriteToRequest writes these params to a swagger request
func (o *DailyStationTimeTableAPIControllerGet2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param TrainDate
	if err := r.SetPathParam("TrainDate", o.TrainDate); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
