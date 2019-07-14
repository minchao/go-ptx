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

	strfmt "github.com/go-openapi/strfmt"
)

// NewTrainLiveBoardAPIControllerGet1Params creates a new TrainLiveBoardAPIControllerGet1Params object
// with the default values initialized.
func NewTrainLiveBoardAPIControllerGet1Params() *TrainLiveBoardAPIControllerGet1Params {
	var (
		dollarTopDefault = string("30")
	)
	return &TrainLiveBoardAPIControllerGet1Params{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTrainLiveBoardAPIControllerGet1ParamsWithTimeout creates a new TrainLiveBoardAPIControllerGet1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewTrainLiveBoardAPIControllerGet1ParamsWithTimeout(timeout time.Duration) *TrainLiveBoardAPIControllerGet1Params {
	var (
		dollarTopDefault = string("30")
	)
	return &TrainLiveBoardAPIControllerGet1Params{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTrainLiveBoardAPIControllerGet1ParamsWithContext creates a new TrainLiveBoardAPIControllerGet1Params object
// with the default values initialized, and the ability to set a context for a request
func NewTrainLiveBoardAPIControllerGet1ParamsWithContext(ctx context.Context) *TrainLiveBoardAPIControllerGet1Params {
	var (
		dollarTopDefault = string("30")
	)
	return &TrainLiveBoardAPIControllerGet1Params{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTrainLiveBoardAPIControllerGet1ParamsWithHTTPClient creates a new TrainLiveBoardAPIControllerGet1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTrainLiveBoardAPIControllerGet1ParamsWithHTTPClient(client *http.Client) *TrainLiveBoardAPIControllerGet1Params {
	var (
		dollarTopDefault = string("30")
	)
	return &TrainLiveBoardAPIControllerGet1Params{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*TrainLiveBoardAPIControllerGet1Params contains all the parameters to send to the API endpoint
for the train live board Api controller get 1 operation typically these are written to a http.Request
*/
type TrainLiveBoardAPIControllerGet1Params struct {

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
	DollarTop *string
	/*TrainNo
	  欲查詢車次的代碼

	*/
	TrainNo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) WithTimeout(timeout time.Duration) *TrainLiveBoardAPIControllerGet1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) WithContext(ctx context.Context) *TrainLiveBoardAPIControllerGet1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) WithHTTPClient(client *http.Client) *TrainLiveBoardAPIControllerGet1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) WithDollarCount(dollarCount *string) *TrainLiveBoardAPIControllerGet1Params {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) SetDollarCount(dollarCount *string) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) WithDollarFilter(dollarFilter *string) *TrainLiveBoardAPIControllerGet1Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) WithDollarFormat(dollarFormat string) *TrainLiveBoardAPIControllerGet1Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) WithDollarOrderby(dollarOrderby *string) *TrainLiveBoardAPIControllerGet1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) WithDollarSelect(dollarSelect *string) *TrainLiveBoardAPIControllerGet1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) WithDollarSkip(dollarSkip *string) *TrainLiveBoardAPIControllerGet1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) WithDollarTop(dollarTop *string) *TrainLiveBoardAPIControllerGet1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) SetDollarTop(dollarTop *string) {
	o.DollarTop = dollarTop
}

// WithTrainNo adds the trainNo to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) WithTrainNo(trainNo string) *TrainLiveBoardAPIControllerGet1Params {
	o.SetTrainNo(trainNo)
	return o
}

// SetTrainNo adds the trainNo to the train live board Api controller get 1 params
func (o *TrainLiveBoardAPIControllerGet1Params) SetTrainNo(trainNo string) {
	o.TrainNo = trainNo
}

// WriteToRequest writes these params to a swagger request
func (o *TrainLiveBoardAPIControllerGet1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param TrainNo
	if err := r.SetPathParam("TrainNo", o.TrainNo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
