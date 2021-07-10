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

// NewTRAAPIODDailyTimetableParams creates a new TRAAPIODDailyTimetableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTRAAPIODDailyTimetableParams() *TRAAPIODDailyTimetableParams {
	return &TRAAPIODDailyTimetableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTRAAPIODDailyTimetableParamsWithTimeout creates a new TRAAPIODDailyTimetableParams object
// with the ability to set a timeout on a request.
func NewTRAAPIODDailyTimetableParamsWithTimeout(timeout time.Duration) *TRAAPIODDailyTimetableParams {
	return &TRAAPIODDailyTimetableParams{
		timeout: timeout,
	}
}

// NewTRAAPIODDailyTimetableParamsWithContext creates a new TRAAPIODDailyTimetableParams object
// with the ability to set a context for a request.
func NewTRAAPIODDailyTimetableParamsWithContext(ctx context.Context) *TRAAPIODDailyTimetableParams {
	return &TRAAPIODDailyTimetableParams{
		Context: ctx,
	}
}

// NewTRAAPIODDailyTimetableParamsWithHTTPClient creates a new TRAAPIODDailyTimetableParams object
// with the ability to set a custom HTTPClient for a request.
func NewTRAAPIODDailyTimetableParamsWithHTTPClient(client *http.Client) *TRAAPIODDailyTimetableParams {
	return &TRAAPIODDailyTimetableParams{
		HTTPClient: client,
	}
}

/* TRAAPIODDailyTimetableParams contains all the parameters to send to the API endpoint
   for the t r a Api o d daily timetable operation.

   Typically these are written to a http.Request.
*/
type TRAAPIODDailyTimetableParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the t r a Api o d daily timetable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TRAAPIODDailyTimetableParams) WithDefaults() *TRAAPIODDailyTimetableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the t r a Api o d daily timetable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TRAAPIODDailyTimetableParams) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := TRAAPIODDailyTimetableParams{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) WithTimeout(timeout time.Duration) *TRAAPIODDailyTimetableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) WithContext(ctx context.Context) *TRAAPIODDailyTimetableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) WithHTTPClient(client *http.Client) *TRAAPIODDailyTimetableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) WithDollarFilter(dollarFilter *string) *TRAAPIODDailyTimetableParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) WithDollarFormat(dollarFormat string) *TRAAPIODDailyTimetableParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) WithDollarOrderby(dollarOrderby *string) *TRAAPIODDailyTimetableParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) WithDollarSelect(dollarSelect *string) *TRAAPIODDailyTimetableParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) WithDollarSkip(dollarSkip *string) *TRAAPIODDailyTimetableParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) WithDollarTop(dollarTop *int64) *TRAAPIODDailyTimetableParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithDestinationStationID adds the destinationStationID to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) WithDestinationStationID(destinationStationID string) *TRAAPIODDailyTimetableParams {
	o.SetDestinationStationID(destinationStationID)
	return o
}

// SetDestinationStationID adds the destinationStationId to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) SetDestinationStationID(destinationStationID string) {
	o.DestinationStationID = destinationStationID
}

// WithOriginStationID adds the originStationID to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) WithOriginStationID(originStationID string) *TRAAPIODDailyTimetableParams {
	o.SetOriginStationID(originStationID)
	return o
}

// SetOriginStationID adds the originStationId to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) SetOriginStationID(originStationID string) {
	o.OriginStationID = originStationID
}

// WithTrainDate adds the trainDate to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) WithTrainDate(trainDate strfmt.DateTime) *TRAAPIODDailyTimetableParams {
	o.SetTrainDate(trainDate)
	return o
}

// SetTrainDate adds the trainDate to the t r a Api o d daily timetable params
func (o *TRAAPIODDailyTimetableParams) SetTrainDate(trainDate strfmt.DateTime) {
	o.TrainDate = trainDate
}

// WriteToRequest writes these params to a swagger request
func (o *TRAAPIODDailyTimetableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
