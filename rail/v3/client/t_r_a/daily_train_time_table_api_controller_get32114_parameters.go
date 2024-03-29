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

// NewDailyTrainTimeTableAPIControllerGet32114Params creates a new DailyTrainTimeTableAPIControllerGet32114Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDailyTrainTimeTableAPIControllerGet32114Params() *DailyTrainTimeTableAPIControllerGet32114Params {
	return &DailyTrainTimeTableAPIControllerGet32114Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDailyTrainTimeTableAPIControllerGet32114ParamsWithTimeout creates a new DailyTrainTimeTableAPIControllerGet32114Params object
// with the ability to set a timeout on a request.
func NewDailyTrainTimeTableAPIControllerGet32114ParamsWithTimeout(timeout time.Duration) *DailyTrainTimeTableAPIControllerGet32114Params {
	return &DailyTrainTimeTableAPIControllerGet32114Params{
		timeout: timeout,
	}
}

// NewDailyTrainTimeTableAPIControllerGet32114ParamsWithContext creates a new DailyTrainTimeTableAPIControllerGet32114Params object
// with the ability to set a context for a request.
func NewDailyTrainTimeTableAPIControllerGet32114ParamsWithContext(ctx context.Context) *DailyTrainTimeTableAPIControllerGet32114Params {
	return &DailyTrainTimeTableAPIControllerGet32114Params{
		Context: ctx,
	}
}

// NewDailyTrainTimeTableAPIControllerGet32114ParamsWithHTTPClient creates a new DailyTrainTimeTableAPIControllerGet32114Params object
// with the ability to set a custom HTTPClient for a request.
func NewDailyTrainTimeTableAPIControllerGet32114ParamsWithHTTPClient(client *http.Client) *DailyTrainTimeTableAPIControllerGet32114Params {
	return &DailyTrainTimeTableAPIControllerGet32114Params{
		HTTPClient: client,
	}
}

/* DailyTrainTimeTableAPIControllerGet32114Params contains all the parameters to send to the API endpoint
   for the daily train time table Api controller get 3211 4 operation.

   Typically these are written to a http.Request.
*/
type DailyTrainTimeTableAPIControllerGet32114Params struct {

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

	/* TrainDate.

	   欲查詢的日期(格式: yyyy-MM-dd)

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

// WithDefaults hydrates default values in the daily train time table Api controller get 3211 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DailyTrainTimeTableAPIControllerGet32114Params) WithDefaults() *DailyTrainTimeTableAPIControllerGet32114Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the daily train time table Api controller get 3211 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DailyTrainTimeTableAPIControllerGet32114Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := DailyTrainTimeTableAPIControllerGet32114Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) WithTimeout(timeout time.Duration) *DailyTrainTimeTableAPIControllerGet32114Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) WithContext(ctx context.Context) *DailyTrainTimeTableAPIControllerGet32114Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) WithHTTPClient(client *http.Client) *DailyTrainTimeTableAPIControllerGet32114Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) WithDollarCount(dollarCount *bool) *DailyTrainTimeTableAPIControllerGet32114Params {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) WithDollarFilter(dollarFilter *string) *DailyTrainTimeTableAPIControllerGet32114Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) WithDollarFormat(dollarFormat string) *DailyTrainTimeTableAPIControllerGet32114Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) WithDollarOrderby(dollarOrderby *string) *DailyTrainTimeTableAPIControllerGet32114Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) WithDollarSelect(dollarSelect *string) *DailyTrainTimeTableAPIControllerGet32114Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) WithDollarSkip(dollarSkip *string) *DailyTrainTimeTableAPIControllerGet32114Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) WithDollarTop(dollarTop *int64) *DailyTrainTimeTableAPIControllerGet32114Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithDestinationStationID adds the destinationStationID to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) WithDestinationStationID(destinationStationID string) *DailyTrainTimeTableAPIControllerGet32114Params {
	o.SetDestinationStationID(destinationStationID)
	return o
}

// SetDestinationStationID adds the destinationStationId to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) SetDestinationStationID(destinationStationID string) {
	o.DestinationStationID = destinationStationID
}

// WithOriginStationID adds the originStationID to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) WithOriginStationID(originStationID string) *DailyTrainTimeTableAPIControllerGet32114Params {
	o.SetOriginStationID(originStationID)
	return o
}

// SetOriginStationID adds the originStationId to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) SetOriginStationID(originStationID string) {
	o.OriginStationID = originStationID
}

// WithTrainDate adds the trainDate to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) WithTrainDate(trainDate strfmt.DateTime) *DailyTrainTimeTableAPIControllerGet32114Params {
	o.SetTrainDate(trainDate)
	return o
}

// SetTrainDate adds the trainDate to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) SetTrainDate(trainDate strfmt.DateTime) {
	o.TrainDate = trainDate
}

// WithHealth adds the health to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) WithHealth(health *string) *DailyTrainTimeTableAPIControllerGet32114Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the daily train time table Api controller get 3211 4 params
func (o *DailyTrainTimeTableAPIControllerGet32114Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *DailyTrainTimeTableAPIControllerGet32114Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
