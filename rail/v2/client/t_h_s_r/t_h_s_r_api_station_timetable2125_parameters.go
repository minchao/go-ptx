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

// NewTHSRAPIStationTimetable2125Params creates a new THSRAPIStationTimetable2125Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTHSRAPIStationTimetable2125Params() *THSRAPIStationTimetable2125Params {
	return &THSRAPIStationTimetable2125Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewTHSRAPIStationTimetable2125ParamsWithTimeout creates a new THSRAPIStationTimetable2125Params object
// with the ability to set a timeout on a request.
func NewTHSRAPIStationTimetable2125ParamsWithTimeout(timeout time.Duration) *THSRAPIStationTimetable2125Params {
	return &THSRAPIStationTimetable2125Params{
		timeout: timeout,
	}
}

// NewTHSRAPIStationTimetable2125ParamsWithContext creates a new THSRAPIStationTimetable2125Params object
// with the ability to set a context for a request.
func NewTHSRAPIStationTimetable2125ParamsWithContext(ctx context.Context) *THSRAPIStationTimetable2125Params {
	return &THSRAPIStationTimetable2125Params{
		Context: ctx,
	}
}

// NewTHSRAPIStationTimetable2125ParamsWithHTTPClient creates a new THSRAPIStationTimetable2125Params object
// with the ability to set a custom HTTPClient for a request.
func NewTHSRAPIStationTimetable2125ParamsWithHTTPClient(client *http.Client) *THSRAPIStationTimetable2125Params {
	return &THSRAPIStationTimetable2125Params{
		HTTPClient: client,
	}
}

/* THSRAPIStationTimetable2125Params contains all the parameters to send to the API endpoint
   for the t h s r Api station timetable 2125 operation.

   Typically these are written to a http.Request.
*/
type THSRAPIStationTimetable2125Params struct {

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

	/* StationID.

	   車站代號
	*/
	StationID string

	/* TrainDate.

	   時刻表日期(格式: yyyy-MM-dd)

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

// WithDefaults hydrates default values in the t h s r Api station timetable 2125 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *THSRAPIStationTimetable2125Params) WithDefaults() *THSRAPIStationTimetable2125Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the t h s r Api station timetable 2125 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *THSRAPIStationTimetable2125Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := THSRAPIStationTimetable2125Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) WithTimeout(timeout time.Duration) *THSRAPIStationTimetable2125Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) WithContext(ctx context.Context) *THSRAPIStationTimetable2125Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) WithHTTPClient(client *http.Client) *THSRAPIStationTimetable2125Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) WithDollarFilter(dollarFilter *string) *THSRAPIStationTimetable2125Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) WithDollarFormat(dollarFormat string) *THSRAPIStationTimetable2125Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) WithDollarOrderby(dollarOrderby *string) *THSRAPIStationTimetable2125Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) WithDollarSelect(dollarSelect *string) *THSRAPIStationTimetable2125Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) WithDollarSkip(dollarSkip *string) *THSRAPIStationTimetable2125Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) WithDollarTop(dollarTop *int64) *THSRAPIStationTimetable2125Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithStationID adds the stationID to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) WithStationID(stationID string) *THSRAPIStationTimetable2125Params {
	o.SetStationID(stationID)
	return o
}

// SetStationID adds the stationId to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) SetStationID(stationID string) {
	o.StationID = stationID
}

// WithTrainDate adds the trainDate to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) WithTrainDate(trainDate strfmt.DateTime) *THSRAPIStationTimetable2125Params {
	o.SetTrainDate(trainDate)
	return o
}

// SetTrainDate adds the trainDate to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) SetTrainDate(trainDate strfmt.DateTime) {
	o.TrainDate = trainDate
}

// WithHealth adds the health to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) WithHealth(health *string) *THSRAPIStationTimetable2125Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the t h s r Api station timetable 2125 params
func (o *THSRAPIStationTimetable2125Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *THSRAPIStationTimetable2125Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param StationID
	if err := r.SetPathParam("StationID", o.StationID); err != nil {
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
