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

// NewDailyTrainTimeTableAPIControllerGet4Params creates a new DailyTrainTimeTableAPIControllerGet4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDailyTrainTimeTableAPIControllerGet4Params() *DailyTrainTimeTableAPIControllerGet4Params {
	return &DailyTrainTimeTableAPIControllerGet4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDailyTrainTimeTableAPIControllerGet4ParamsWithTimeout creates a new DailyTrainTimeTableAPIControllerGet4Params object
// with the ability to set a timeout on a request.
func NewDailyTrainTimeTableAPIControllerGet4ParamsWithTimeout(timeout time.Duration) *DailyTrainTimeTableAPIControllerGet4Params {
	return &DailyTrainTimeTableAPIControllerGet4Params{
		timeout: timeout,
	}
}

// NewDailyTrainTimeTableAPIControllerGet4ParamsWithContext creates a new DailyTrainTimeTableAPIControllerGet4Params object
// with the ability to set a context for a request.
func NewDailyTrainTimeTableAPIControllerGet4ParamsWithContext(ctx context.Context) *DailyTrainTimeTableAPIControllerGet4Params {
	return &DailyTrainTimeTableAPIControllerGet4Params{
		Context: ctx,
	}
}

// NewDailyTrainTimeTableAPIControllerGet4ParamsWithHTTPClient creates a new DailyTrainTimeTableAPIControllerGet4Params object
// with the ability to set a custom HTTPClient for a request.
func NewDailyTrainTimeTableAPIControllerGet4ParamsWithHTTPClient(client *http.Client) *DailyTrainTimeTableAPIControllerGet4Params {
	return &DailyTrainTimeTableAPIControllerGet4Params{
		HTTPClient: client,
	}
}

/* DailyTrainTimeTableAPIControllerGet4Params contains all the parameters to send to the API endpoint
   for the daily train time table Api controller get 4 operation.

   Typically these are written to a http.Request.
*/
type DailyTrainTimeTableAPIControllerGet4Params struct {

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
	*/
	TrainDate string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the daily train time table Api controller get 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DailyTrainTimeTableAPIControllerGet4Params) WithDefaults() *DailyTrainTimeTableAPIControllerGet4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the daily train time table Api controller get 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DailyTrainTimeTableAPIControllerGet4Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := DailyTrainTimeTableAPIControllerGet4Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) WithTimeout(timeout time.Duration) *DailyTrainTimeTableAPIControllerGet4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) WithContext(ctx context.Context) *DailyTrainTimeTableAPIControllerGet4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) WithHTTPClient(client *http.Client) *DailyTrainTimeTableAPIControllerGet4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) WithDollarCount(dollarCount *bool) *DailyTrainTimeTableAPIControllerGet4Params {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) WithDollarFilter(dollarFilter *string) *DailyTrainTimeTableAPIControllerGet4Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) WithDollarFormat(dollarFormat string) *DailyTrainTimeTableAPIControllerGet4Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) WithDollarOrderby(dollarOrderby *string) *DailyTrainTimeTableAPIControllerGet4Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) WithDollarSelect(dollarSelect *string) *DailyTrainTimeTableAPIControllerGet4Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) WithDollarSkip(dollarSkip *string) *DailyTrainTimeTableAPIControllerGet4Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) WithDollarTop(dollarTop *int64) *DailyTrainTimeTableAPIControllerGet4Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithDestinationStationID adds the destinationStationID to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) WithDestinationStationID(destinationStationID string) *DailyTrainTimeTableAPIControllerGet4Params {
	o.SetDestinationStationID(destinationStationID)
	return o
}

// SetDestinationStationID adds the destinationStationId to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) SetDestinationStationID(destinationStationID string) {
	o.DestinationStationID = destinationStationID
}

// WithOriginStationID adds the originStationID to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) WithOriginStationID(originStationID string) *DailyTrainTimeTableAPIControllerGet4Params {
	o.SetOriginStationID(originStationID)
	return o
}

// SetOriginStationID adds the originStationId to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) SetOriginStationID(originStationID string) {
	o.OriginStationID = originStationID
}

// WithTrainDate adds the trainDate to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) WithTrainDate(trainDate string) *DailyTrainTimeTableAPIControllerGet4Params {
	o.SetTrainDate(trainDate)
	return o
}

// SetTrainDate adds the trainDate to the daily train time table Api controller get 4 params
func (o *DailyTrainTimeTableAPIControllerGet4Params) SetTrainDate(trainDate string) {
	o.TrainDate = trainDate
}

// WriteToRequest writes these params to a swagger request
func (o *DailyTrainTimeTableAPIControllerGet4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("TrainDate", o.TrainDate); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
