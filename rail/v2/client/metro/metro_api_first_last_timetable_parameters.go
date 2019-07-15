// Code generated by go-swagger; DO NOT EDIT.

package metro

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

// NewMetroAPIFirstLastTimetableParams creates a new MetroAPIFirstLastTimetableParams object
// with the default values initialized.
func NewMetroAPIFirstLastTimetableParams() *MetroAPIFirstLastTimetableParams {
	var (
		dollarTopDefault = string("30")
	)
	return &MetroAPIFirstLastTimetableParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewMetroAPIFirstLastTimetableParamsWithTimeout creates a new MetroAPIFirstLastTimetableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMetroAPIFirstLastTimetableParamsWithTimeout(timeout time.Duration) *MetroAPIFirstLastTimetableParams {
	var (
		dollarTopDefault = string("30")
	)
	return &MetroAPIFirstLastTimetableParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewMetroAPIFirstLastTimetableParamsWithContext creates a new MetroAPIFirstLastTimetableParams object
// with the default values initialized, and the ability to set a context for a request
func NewMetroAPIFirstLastTimetableParamsWithContext(ctx context.Context) *MetroAPIFirstLastTimetableParams {
	var (
		dollarTopDefault = string("30")
	)
	return &MetroAPIFirstLastTimetableParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewMetroAPIFirstLastTimetableParamsWithHTTPClient creates a new MetroAPIFirstLastTimetableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMetroAPIFirstLastTimetableParamsWithHTTPClient(client *http.Client) *MetroAPIFirstLastTimetableParams {
	var (
		dollarTopDefault = string("30")
	)
	return &MetroAPIFirstLastTimetableParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*MetroAPIFirstLastTimetableParams contains all the parameters to send to the API endpoint
for the metro Api first last timetable operation typically these are written to a http.Request
*/
type MetroAPIFirstLastTimetableParams struct {

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
	/*Operator
	  欲查詢縣市

	*/
	Operator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) WithTimeout(timeout time.Duration) *MetroAPIFirstLastTimetableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) WithContext(ctx context.Context) *MetroAPIFirstLastTimetableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) WithHTTPClient(client *http.Client) *MetroAPIFirstLastTimetableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) WithDollarFilter(dollarFilter *string) *MetroAPIFirstLastTimetableParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) WithDollarFormat(dollarFormat string) *MetroAPIFirstLastTimetableParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) WithDollarOrderby(dollarOrderby *string) *MetroAPIFirstLastTimetableParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) WithDollarSelect(dollarSelect *string) *MetroAPIFirstLastTimetableParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) WithDollarSkip(dollarSkip *string) *MetroAPIFirstLastTimetableParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) WithDollarTop(dollarTop *string) *MetroAPIFirstLastTimetableParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) SetDollarTop(dollarTop *string) {
	o.DollarTop = dollarTop
}

// WithOperator adds the operator to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) WithOperator(operator string) *MetroAPIFirstLastTimetableParams {
	o.SetOperator(operator)
	return o
}

// SetOperator adds the operator to the metro Api first last timetable params
func (o *MetroAPIFirstLastTimetableParams) SetOperator(operator string) {
	o.Operator = operator
}

// WriteToRequest writes these params to a swagger request
func (o *MetroAPIFirstLastTimetableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param Operator
	if err := r.SetPathParam("Operator", o.Operator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
