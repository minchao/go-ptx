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

// NewDailyTrainTimeTableAPIControllerGet32113Params creates a new DailyTrainTimeTableAPIControllerGet32113Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDailyTrainTimeTableAPIControllerGet32113Params() *DailyTrainTimeTableAPIControllerGet32113Params {
	return &DailyTrainTimeTableAPIControllerGet32113Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDailyTrainTimeTableAPIControllerGet32113ParamsWithTimeout creates a new DailyTrainTimeTableAPIControllerGet32113Params object
// with the ability to set a timeout on a request.
func NewDailyTrainTimeTableAPIControllerGet32113ParamsWithTimeout(timeout time.Duration) *DailyTrainTimeTableAPIControllerGet32113Params {
	return &DailyTrainTimeTableAPIControllerGet32113Params{
		timeout: timeout,
	}
}

// NewDailyTrainTimeTableAPIControllerGet32113ParamsWithContext creates a new DailyTrainTimeTableAPIControllerGet32113Params object
// with the ability to set a context for a request.
func NewDailyTrainTimeTableAPIControllerGet32113ParamsWithContext(ctx context.Context) *DailyTrainTimeTableAPIControllerGet32113Params {
	return &DailyTrainTimeTableAPIControllerGet32113Params{
		Context: ctx,
	}
}

// NewDailyTrainTimeTableAPIControllerGet32113ParamsWithHTTPClient creates a new DailyTrainTimeTableAPIControllerGet32113Params object
// with the ability to set a custom HTTPClient for a request.
func NewDailyTrainTimeTableAPIControllerGet32113ParamsWithHTTPClient(client *http.Client) *DailyTrainTimeTableAPIControllerGet32113Params {
	return &DailyTrainTimeTableAPIControllerGet32113Params{
		HTTPClient: client,
	}
}

/* DailyTrainTimeTableAPIControllerGet32113Params contains all the parameters to send to the API endpoint
   for the daily train time table Api controller get 3211 3 operation.

   Typically these are written to a http.Request.
*/
type DailyTrainTimeTableAPIControllerGet32113Params struct {

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

	/* TrainDate.

	   欲查詢車次的日期(格式: yyyy-MM-dd)

	   Format: date-time
	*/
	TrainDate strfmt.DateTime

	/* Health.

	   加入參數'?health=true'即可查詢此API服務的健康狀態
	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the daily train time table Api controller get 3211 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DailyTrainTimeTableAPIControllerGet32113Params) WithDefaults() *DailyTrainTimeTableAPIControllerGet32113Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the daily train time table Api controller get 3211 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DailyTrainTimeTableAPIControllerGet32113Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := DailyTrainTimeTableAPIControllerGet32113Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) WithTimeout(timeout time.Duration) *DailyTrainTimeTableAPIControllerGet32113Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) WithContext(ctx context.Context) *DailyTrainTimeTableAPIControllerGet32113Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) WithHTTPClient(client *http.Client) *DailyTrainTimeTableAPIControllerGet32113Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) WithDollarCount(dollarCount *bool) *DailyTrainTimeTableAPIControllerGet32113Params {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) WithDollarFilter(dollarFilter *string) *DailyTrainTimeTableAPIControllerGet32113Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) WithDollarFormat(dollarFormat string) *DailyTrainTimeTableAPIControllerGet32113Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) WithDollarOrderby(dollarOrderby *string) *DailyTrainTimeTableAPIControllerGet32113Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) WithDollarSelect(dollarSelect *string) *DailyTrainTimeTableAPIControllerGet32113Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) WithDollarSkip(dollarSkip *string) *DailyTrainTimeTableAPIControllerGet32113Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) WithDollarTop(dollarTop *int64) *DailyTrainTimeTableAPIControllerGet32113Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithTrainDate adds the trainDate to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) WithTrainDate(trainDate strfmt.DateTime) *DailyTrainTimeTableAPIControllerGet32113Params {
	o.SetTrainDate(trainDate)
	return o
}

// SetTrainDate adds the trainDate to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) SetTrainDate(trainDate strfmt.DateTime) {
	o.TrainDate = trainDate
}

// WithHealth adds the health to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) WithHealth(health *string) *DailyTrainTimeTableAPIControllerGet32113Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the daily train time table Api controller get 3211 3 params
func (o *DailyTrainTimeTableAPIControllerGet32113Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *DailyTrainTimeTableAPIControllerGet32113Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("TrainDate", o.TrainDate.String()); err != nil {
		return err
	}

	if o.Health != nil {

		// query param health
		var qrHealth string

		if o.Health != nil {
			qrHealth = *o.Health
		}
		qHealth := qrHealth
		if qHealth != "" {

			if err := r.SetQueryParam("health", qHealth); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}