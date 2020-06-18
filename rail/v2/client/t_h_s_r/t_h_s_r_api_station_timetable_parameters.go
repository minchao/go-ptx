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

// NewTHSRAPIStationTimetableParams creates a new THSRAPIStationTimetableParams object
// with the default values initialized.
func NewTHSRAPIStationTimetableParams() *THSRAPIStationTimetableParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &THSRAPIStationTimetableParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTHSRAPIStationTimetableParamsWithTimeout creates a new THSRAPIStationTimetableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTHSRAPIStationTimetableParamsWithTimeout(timeout time.Duration) *THSRAPIStationTimetableParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &THSRAPIStationTimetableParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTHSRAPIStationTimetableParamsWithContext creates a new THSRAPIStationTimetableParams object
// with the default values initialized, and the ability to set a context for a request
func NewTHSRAPIStationTimetableParamsWithContext(ctx context.Context) *THSRAPIStationTimetableParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &THSRAPIStationTimetableParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTHSRAPIStationTimetableParamsWithHTTPClient creates a new THSRAPIStationTimetableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTHSRAPIStationTimetableParamsWithHTTPClient(client *http.Client) *THSRAPIStationTimetableParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &THSRAPIStationTimetableParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*THSRAPIStationTimetableParams contains all the parameters to send to the API endpoint
for the t h s r Api station timetable operation typically these are written to a http.Request
*/
type THSRAPIStationTimetableParams struct {

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
	/*StationID*/
	StationID string
	/*TrainDate*/
	TrainDate string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) WithTimeout(timeout time.Duration) *THSRAPIStationTimetableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) WithContext(ctx context.Context) *THSRAPIStationTimetableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) WithHTTPClient(client *http.Client) *THSRAPIStationTimetableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) WithDollarFilter(dollarFilter *string) *THSRAPIStationTimetableParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) WithDollarFormat(dollarFormat string) *THSRAPIStationTimetableParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) WithDollarOrderby(dollarOrderby *string) *THSRAPIStationTimetableParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) WithDollarSelect(dollarSelect *string) *THSRAPIStationTimetableParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) WithDollarSkip(dollarSkip *string) *THSRAPIStationTimetableParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) WithDollarTop(dollarTop *int64) *THSRAPIStationTimetableParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithStationID adds the stationID to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) WithStationID(stationID string) *THSRAPIStationTimetableParams {
	o.SetStationID(stationID)
	return o
}

// SetStationID adds the stationId to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) SetStationID(stationID string) {
	o.StationID = stationID
}

// WithTrainDate adds the trainDate to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) WithTrainDate(trainDate string) *THSRAPIStationTimetableParams {
	o.SetTrainDate(trainDate)
	return o
}

// SetTrainDate adds the trainDate to the t h s r Api station timetable params
func (o *THSRAPIStationTimetableParams) SetTrainDate(trainDate string) {
	o.TrainDate = trainDate
}

// WriteToRequest writes these params to a swagger request
func (o *THSRAPIStationTimetableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("TrainDate", o.TrainDate); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
