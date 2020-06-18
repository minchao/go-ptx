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

// NewGeneralTrainTimetableAPIControllerGet1Params creates a new GeneralTrainTimetableAPIControllerGet1Params object
// with the default values initialized.
func NewGeneralTrainTimetableAPIControllerGet1Params() *GeneralTrainTimetableAPIControllerGet1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &GeneralTrainTimetableAPIControllerGet1Params{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGeneralTrainTimetableAPIControllerGet1ParamsWithTimeout creates a new GeneralTrainTimetableAPIControllerGet1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGeneralTrainTimetableAPIControllerGet1ParamsWithTimeout(timeout time.Duration) *GeneralTrainTimetableAPIControllerGet1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &GeneralTrainTimetableAPIControllerGet1Params{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewGeneralTrainTimetableAPIControllerGet1ParamsWithContext creates a new GeneralTrainTimetableAPIControllerGet1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGeneralTrainTimetableAPIControllerGet1ParamsWithContext(ctx context.Context) *GeneralTrainTimetableAPIControllerGet1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &GeneralTrainTimetableAPIControllerGet1Params{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewGeneralTrainTimetableAPIControllerGet1ParamsWithHTTPClient creates a new GeneralTrainTimetableAPIControllerGet1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGeneralTrainTimetableAPIControllerGet1ParamsWithHTTPClient(client *http.Client) *GeneralTrainTimetableAPIControllerGet1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &GeneralTrainTimetableAPIControllerGet1Params{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*GeneralTrainTimetableAPIControllerGet1Params contains all the parameters to send to the API endpoint
for the general train timetable Api controller get 1 operation typically these are written to a http.Request
*/
type GeneralTrainTimetableAPIControllerGet1Params struct {

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
	/*TrainNo
	  欲查詢車次的代碼

	*/
	TrainNo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) WithTimeout(timeout time.Duration) *GeneralTrainTimetableAPIControllerGet1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) WithContext(ctx context.Context) *GeneralTrainTimetableAPIControllerGet1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) WithHTTPClient(client *http.Client) *GeneralTrainTimetableAPIControllerGet1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) WithDollarCount(dollarCount *string) *GeneralTrainTimetableAPIControllerGet1Params {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) SetDollarCount(dollarCount *string) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) WithDollarFilter(dollarFilter *string) *GeneralTrainTimetableAPIControllerGet1Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) WithDollarFormat(dollarFormat string) *GeneralTrainTimetableAPIControllerGet1Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) WithDollarOrderby(dollarOrderby *string) *GeneralTrainTimetableAPIControllerGet1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) WithDollarSelect(dollarSelect *string) *GeneralTrainTimetableAPIControllerGet1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) WithDollarSkip(dollarSkip *string) *GeneralTrainTimetableAPIControllerGet1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) WithDollarTop(dollarTop *int64) *GeneralTrainTimetableAPIControllerGet1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithTrainNo adds the trainNo to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) WithTrainNo(trainNo string) *GeneralTrainTimetableAPIControllerGet1Params {
	o.SetTrainNo(trainNo)
	return o
}

// SetTrainNo adds the trainNo to the general train timetable Api controller get 1 params
func (o *GeneralTrainTimetableAPIControllerGet1Params) SetTrainNo(trainNo string) {
	o.TrainNo = trainNo
}

// WriteToRequest writes these params to a swagger request
func (o *GeneralTrainTimetableAPIControllerGet1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param TrainNo
	if err := r.SetPathParam("TrainNo", o.TrainNo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
