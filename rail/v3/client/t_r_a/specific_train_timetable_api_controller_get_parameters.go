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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSpecificTrainTimetableAPIControllerGetParams creates a new SpecificTrainTimetableAPIControllerGetParams object
// with the default values initialized.
func NewSpecificTrainTimetableAPIControllerGetParams() *SpecificTrainTimetableAPIControllerGetParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &SpecificTrainTimetableAPIControllerGetParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSpecificTrainTimetableAPIControllerGetParamsWithTimeout creates a new SpecificTrainTimetableAPIControllerGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSpecificTrainTimetableAPIControllerGetParamsWithTimeout(timeout time.Duration) *SpecificTrainTimetableAPIControllerGetParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &SpecificTrainTimetableAPIControllerGetParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewSpecificTrainTimetableAPIControllerGetParamsWithContext creates a new SpecificTrainTimetableAPIControllerGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSpecificTrainTimetableAPIControllerGetParamsWithContext(ctx context.Context) *SpecificTrainTimetableAPIControllerGetParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &SpecificTrainTimetableAPIControllerGetParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewSpecificTrainTimetableAPIControllerGetParamsWithHTTPClient creates a new SpecificTrainTimetableAPIControllerGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSpecificTrainTimetableAPIControllerGetParamsWithHTTPClient(client *http.Client) *SpecificTrainTimetableAPIControllerGetParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &SpecificTrainTimetableAPIControllerGetParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*SpecificTrainTimetableAPIControllerGetParams contains all the parameters to send to the API endpoint
for the specific train timetable Api controller get operation typically these are written to a http.Request
*/
type SpecificTrainTimetableAPIControllerGetParams struct {

	/*DollarCount
	  查詢數量

	*/
	DollarCount *string
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) WithTimeout(timeout time.Duration) *SpecificTrainTimetableAPIControllerGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) WithContext(ctx context.Context) *SpecificTrainTimetableAPIControllerGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) WithHTTPClient(client *http.Client) *SpecificTrainTimetableAPIControllerGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) WithDollarCount(dollarCount *string) *SpecificTrainTimetableAPIControllerGetParams {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) SetDollarCount(dollarCount *string) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) WithDollarFilter(dollarFilter *string) *SpecificTrainTimetableAPIControllerGetParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) WithDollarFormat(dollarFormat string) *SpecificTrainTimetableAPIControllerGetParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) WithDollarOrderby(dollarOrderby *string) *SpecificTrainTimetableAPIControllerGetParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) WithDollarSelect(dollarSelect *string) *SpecificTrainTimetableAPIControllerGetParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) WithDollarSkip(dollarSkip *string) *SpecificTrainTimetableAPIControllerGetParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) WithDollarTop(dollarTop *int64) *SpecificTrainTimetableAPIControllerGetParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the specific train timetable Api controller get params
func (o *SpecificTrainTimetableAPIControllerGetParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *SpecificTrainTimetableAPIControllerGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarCount != nil {

		// query param $count
		var qrDollarCount string
		if o.DollarCount != nil {
			qrDollarCount = *o.DollarCount
		}
		qDollarCount := qrDollarCount
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
