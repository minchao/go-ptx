// Code generated by go-swagger; DO NOT EDIT.

package tourism

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

// NewTaiwanTripBusAPIRealTimeNearStopParams creates a new TaiwanTripBusAPIRealTimeNearStopParams object
// with the default values initialized.
func NewTaiwanTripBusAPIRealTimeNearStopParams() *TaiwanTripBusAPIRealTimeNearStopParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &TaiwanTripBusAPIRealTimeNearStopParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTaiwanTripBusAPIRealTimeNearStopParamsWithTimeout creates a new TaiwanTripBusAPIRealTimeNearStopParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTaiwanTripBusAPIRealTimeNearStopParamsWithTimeout(timeout time.Duration) *TaiwanTripBusAPIRealTimeNearStopParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &TaiwanTripBusAPIRealTimeNearStopParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTaiwanTripBusAPIRealTimeNearStopParamsWithContext creates a new TaiwanTripBusAPIRealTimeNearStopParams object
// with the default values initialized, and the ability to set a context for a request
func NewTaiwanTripBusAPIRealTimeNearStopParamsWithContext(ctx context.Context) *TaiwanTripBusAPIRealTimeNearStopParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &TaiwanTripBusAPIRealTimeNearStopParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTaiwanTripBusAPIRealTimeNearStopParamsWithHTTPClient creates a new TaiwanTripBusAPIRealTimeNearStopParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTaiwanTripBusAPIRealTimeNearStopParamsWithHTTPClient(client *http.Client) *TaiwanTripBusAPIRealTimeNearStopParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &TaiwanTripBusAPIRealTimeNearStopParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*TaiwanTripBusAPIRealTimeNearStopParams contains all the parameters to send to the API endpoint
for the taiwan trip bus Api real time near stop operation typically these are written to a http.Request
*/
type TaiwanTripBusAPIRealTimeNearStopParams struct {

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

// WithTimeout adds the timeout to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) WithTimeout(timeout time.Duration) *TaiwanTripBusAPIRealTimeNearStopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) WithContext(ctx context.Context) *TaiwanTripBusAPIRealTimeNearStopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) WithHTTPClient(client *http.Client) *TaiwanTripBusAPIRealTimeNearStopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) WithDollarFilter(dollarFilter *string) *TaiwanTripBusAPIRealTimeNearStopParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) WithDollarFormat(dollarFormat string) *TaiwanTripBusAPIRealTimeNearStopParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) WithDollarOrderby(dollarOrderby *string) *TaiwanTripBusAPIRealTimeNearStopParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) WithDollarSelect(dollarSelect *string) *TaiwanTripBusAPIRealTimeNearStopParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) WithDollarSkip(dollarSkip *string) *TaiwanTripBusAPIRealTimeNearStopParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) WithDollarTop(dollarTop *int64) *TaiwanTripBusAPIRealTimeNearStopParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the taiwan trip bus Api real time near stop params
func (o *TaiwanTripBusAPIRealTimeNearStopParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *TaiwanTripBusAPIRealTimeNearStopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
