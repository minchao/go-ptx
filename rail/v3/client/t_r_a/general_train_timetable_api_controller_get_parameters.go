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

// NewGeneralTrainTimetableAPIControllerGetParams creates a new GeneralTrainTimetableAPIControllerGetParams object
// with the default values initialized.
func NewGeneralTrainTimetableAPIControllerGetParams() *GeneralTrainTimetableAPIControllerGetParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &GeneralTrainTimetableAPIControllerGetParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGeneralTrainTimetableAPIControllerGetParamsWithTimeout creates a new GeneralTrainTimetableAPIControllerGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGeneralTrainTimetableAPIControllerGetParamsWithTimeout(timeout time.Duration) *GeneralTrainTimetableAPIControllerGetParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &GeneralTrainTimetableAPIControllerGetParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewGeneralTrainTimetableAPIControllerGetParamsWithContext creates a new GeneralTrainTimetableAPIControllerGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGeneralTrainTimetableAPIControllerGetParamsWithContext(ctx context.Context) *GeneralTrainTimetableAPIControllerGetParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &GeneralTrainTimetableAPIControllerGetParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewGeneralTrainTimetableAPIControllerGetParamsWithHTTPClient creates a new GeneralTrainTimetableAPIControllerGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGeneralTrainTimetableAPIControllerGetParamsWithHTTPClient(client *http.Client) *GeneralTrainTimetableAPIControllerGetParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &GeneralTrainTimetableAPIControllerGetParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*GeneralTrainTimetableAPIControllerGetParams contains all the parameters to send to the API endpoint
for the general train timetable Api controller get operation typically these are written to a http.Request
*/
type GeneralTrainTimetableAPIControllerGetParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) WithTimeout(timeout time.Duration) *GeneralTrainTimetableAPIControllerGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) WithContext(ctx context.Context) *GeneralTrainTimetableAPIControllerGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) WithHTTPClient(client *http.Client) *GeneralTrainTimetableAPIControllerGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) WithDollarCount(dollarCount *bool) *GeneralTrainTimetableAPIControllerGetParams {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) WithDollarFilter(dollarFilter *string) *GeneralTrainTimetableAPIControllerGetParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) WithDollarFormat(dollarFormat string) *GeneralTrainTimetableAPIControllerGetParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) WithDollarOrderby(dollarOrderby *string) *GeneralTrainTimetableAPIControllerGetParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) WithDollarSelect(dollarSelect *string) *GeneralTrainTimetableAPIControllerGetParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) WithDollarSkip(dollarSkip *string) *GeneralTrainTimetableAPIControllerGetParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) WithDollarTop(dollarTop *int64) *GeneralTrainTimetableAPIControllerGetParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the general train timetable Api controller get params
func (o *GeneralTrainTimetableAPIControllerGetParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *GeneralTrainTimetableAPIControllerGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
