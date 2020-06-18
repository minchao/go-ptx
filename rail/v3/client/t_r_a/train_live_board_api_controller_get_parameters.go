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

// NewTrainLiveBoardAPIControllerGetParams creates a new TrainLiveBoardAPIControllerGetParams object
// with the default values initialized.
func NewTrainLiveBoardAPIControllerGetParams() *TrainLiveBoardAPIControllerGetParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &TrainLiveBoardAPIControllerGetParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTrainLiveBoardAPIControllerGetParamsWithTimeout creates a new TrainLiveBoardAPIControllerGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTrainLiveBoardAPIControllerGetParamsWithTimeout(timeout time.Duration) *TrainLiveBoardAPIControllerGetParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &TrainLiveBoardAPIControllerGetParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTrainLiveBoardAPIControllerGetParamsWithContext creates a new TrainLiveBoardAPIControllerGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewTrainLiveBoardAPIControllerGetParamsWithContext(ctx context.Context) *TrainLiveBoardAPIControllerGetParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &TrainLiveBoardAPIControllerGetParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTrainLiveBoardAPIControllerGetParamsWithHTTPClient creates a new TrainLiveBoardAPIControllerGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTrainLiveBoardAPIControllerGetParamsWithHTTPClient(client *http.Client) *TrainLiveBoardAPIControllerGetParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &TrainLiveBoardAPIControllerGetParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*TrainLiveBoardAPIControllerGetParams contains all the parameters to send to the API endpoint
for the train live board Api controller get operation typically these are written to a http.Request
*/
type TrainLiveBoardAPIControllerGetParams struct {

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

// WithTimeout adds the timeout to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) WithTimeout(timeout time.Duration) *TrainLiveBoardAPIControllerGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) WithContext(ctx context.Context) *TrainLiveBoardAPIControllerGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) WithHTTPClient(client *http.Client) *TrainLiveBoardAPIControllerGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) WithDollarCount(dollarCount *string) *TrainLiveBoardAPIControllerGetParams {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) SetDollarCount(dollarCount *string) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) WithDollarFilter(dollarFilter *string) *TrainLiveBoardAPIControllerGetParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) WithDollarFormat(dollarFormat string) *TrainLiveBoardAPIControllerGetParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) WithDollarOrderby(dollarOrderby *string) *TrainLiveBoardAPIControllerGetParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) WithDollarSelect(dollarSelect *string) *TrainLiveBoardAPIControllerGetParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) WithDollarSkip(dollarSkip *string) *TrainLiveBoardAPIControllerGetParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) WithDollarTop(dollarTop *int64) *TrainLiveBoardAPIControllerGetParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the train live board Api controller get params
func (o *TrainLiveBoardAPIControllerGetParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *TrainLiveBoardAPIControllerGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
