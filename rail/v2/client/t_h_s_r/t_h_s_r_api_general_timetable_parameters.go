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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTHSRAPIGeneralTimetableParams creates a new THSRAPIGeneralTimetableParams object
// with the default values initialized.
func NewTHSRAPIGeneralTimetableParams() *THSRAPIGeneralTimetableParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &THSRAPIGeneralTimetableParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTHSRAPIGeneralTimetableParamsWithTimeout creates a new THSRAPIGeneralTimetableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTHSRAPIGeneralTimetableParamsWithTimeout(timeout time.Duration) *THSRAPIGeneralTimetableParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &THSRAPIGeneralTimetableParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTHSRAPIGeneralTimetableParamsWithContext creates a new THSRAPIGeneralTimetableParams object
// with the default values initialized, and the ability to set a context for a request
func NewTHSRAPIGeneralTimetableParamsWithContext(ctx context.Context) *THSRAPIGeneralTimetableParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &THSRAPIGeneralTimetableParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTHSRAPIGeneralTimetableParamsWithHTTPClient creates a new THSRAPIGeneralTimetableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTHSRAPIGeneralTimetableParamsWithHTTPClient(client *http.Client) *THSRAPIGeneralTimetableParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &THSRAPIGeneralTimetableParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*THSRAPIGeneralTimetableParams contains all the parameters to send to the API endpoint
for the t h s r Api general timetable operation typically these are written to a http.Request
*/
type THSRAPIGeneralTimetableParams struct {

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

// WithTimeout adds the timeout to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) WithTimeout(timeout time.Duration) *THSRAPIGeneralTimetableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) WithContext(ctx context.Context) *THSRAPIGeneralTimetableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) WithHTTPClient(client *http.Client) *THSRAPIGeneralTimetableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) WithDollarFilter(dollarFilter *string) *THSRAPIGeneralTimetableParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) WithDollarFormat(dollarFormat string) *THSRAPIGeneralTimetableParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) WithDollarOrderby(dollarOrderby *string) *THSRAPIGeneralTimetableParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) WithDollarSelect(dollarSelect *string) *THSRAPIGeneralTimetableParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) WithDollarSkip(dollarSkip *string) *THSRAPIGeneralTimetableParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) WithDollarTop(dollarTop *int64) *THSRAPIGeneralTimetableParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the t h s r Api general timetable params
func (o *THSRAPIGeneralTimetableParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *THSRAPIGeneralTimetableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
