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

// NewTRAAPIDailyTrainInfoParams creates a new TRAAPIDailyTrainInfoParams object
// with the default values initialized.
func NewTRAAPIDailyTrainInfoParams() *TRAAPIDailyTrainInfoParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &TRAAPIDailyTrainInfoParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTRAAPIDailyTrainInfoParamsWithTimeout creates a new TRAAPIDailyTrainInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTRAAPIDailyTrainInfoParamsWithTimeout(timeout time.Duration) *TRAAPIDailyTrainInfoParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &TRAAPIDailyTrainInfoParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTRAAPIDailyTrainInfoParamsWithContext creates a new TRAAPIDailyTrainInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewTRAAPIDailyTrainInfoParamsWithContext(ctx context.Context) *TRAAPIDailyTrainInfoParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &TRAAPIDailyTrainInfoParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTRAAPIDailyTrainInfoParamsWithHTTPClient creates a new TRAAPIDailyTrainInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTRAAPIDailyTrainInfoParamsWithHTTPClient(client *http.Client) *TRAAPIDailyTrainInfoParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &TRAAPIDailyTrainInfoParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*TRAAPIDailyTrainInfoParams contains all the parameters to send to the API endpoint
for the t r a Api daily train info operation typically these are written to a http.Request
*/
type TRAAPIDailyTrainInfoParams struct {

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

// WithTimeout adds the timeout to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) WithTimeout(timeout time.Duration) *TRAAPIDailyTrainInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) WithContext(ctx context.Context) *TRAAPIDailyTrainInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) WithHTTPClient(client *http.Client) *TRAAPIDailyTrainInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) WithDollarFilter(dollarFilter *string) *TRAAPIDailyTrainInfoParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) WithDollarFormat(dollarFormat string) *TRAAPIDailyTrainInfoParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) WithDollarOrderby(dollarOrderby *string) *TRAAPIDailyTrainInfoParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) WithDollarSelect(dollarSelect *string) *TRAAPIDailyTrainInfoParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) WithDollarSkip(dollarSkip *string) *TRAAPIDailyTrainInfoParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) WithDollarTop(dollarTop *int64) *TRAAPIDailyTrainInfoParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the t r a Api daily train info params
func (o *TRAAPIDailyTrainInfoParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *TRAAPIDailyTrainInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
