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

// NewTRAAPILiveBoard1Params creates a new TRAAPILiveBoard1Params object
// with the default values initialized.
func NewTRAAPILiveBoard1Params() *TRAAPILiveBoard1Params {
	var (
		dollarTopDefault = string("30")
	)
	return &TRAAPILiveBoard1Params{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTRAAPILiveBoard1ParamsWithTimeout creates a new TRAAPILiveBoard1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewTRAAPILiveBoard1ParamsWithTimeout(timeout time.Duration) *TRAAPILiveBoard1Params {
	var (
		dollarTopDefault = string("30")
	)
	return &TRAAPILiveBoard1Params{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewTRAAPILiveBoard1ParamsWithContext creates a new TRAAPILiveBoard1Params object
// with the default values initialized, and the ability to set a context for a request
func NewTRAAPILiveBoard1ParamsWithContext(ctx context.Context) *TRAAPILiveBoard1Params {
	var (
		dollarTopDefault = string("30")
	)
	return &TRAAPILiveBoard1Params{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewTRAAPILiveBoard1ParamsWithHTTPClient creates a new TRAAPILiveBoard1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTRAAPILiveBoard1ParamsWithHTTPClient(client *http.Client) *TRAAPILiveBoard1Params {
	var (
		dollarTopDefault = string("30")
	)
	return &TRAAPILiveBoard1Params{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*TRAAPILiveBoard1Params contains all the parameters to send to the API endpoint
for the t r a Api live board 1 operation typically these are written to a http.Request
*/
type TRAAPILiveBoard1Params struct {

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
	/*StationID
	  車站代碼

	*/
	StationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) WithTimeout(timeout time.Duration) *TRAAPILiveBoard1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) WithContext(ctx context.Context) *TRAAPILiveBoard1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) WithHTTPClient(client *http.Client) *TRAAPILiveBoard1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) WithDollarFilter(dollarFilter *string) *TRAAPILiveBoard1Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) WithDollarFormat(dollarFormat string) *TRAAPILiveBoard1Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) WithDollarOrderby(dollarOrderby *string) *TRAAPILiveBoard1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) WithDollarSelect(dollarSelect *string) *TRAAPILiveBoard1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) WithDollarSkip(dollarSkip *string) *TRAAPILiveBoard1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) WithDollarTop(dollarTop *string) *TRAAPILiveBoard1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) SetDollarTop(dollarTop *string) {
	o.DollarTop = dollarTop
}

// WithStationID adds the stationID to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) WithStationID(stationID string) *TRAAPILiveBoard1Params {
	o.SetStationID(stationID)
	return o
}

// SetStationID adds the stationId to the t r a Api live board 1 params
func (o *TRAAPILiveBoard1Params) SetStationID(stationID string) {
	o.StationID = stationID
}

// WriteToRequest writes these params to a swagger request
func (o *TRAAPILiveBoard1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param StationID
	if err := r.SetPathParam("StationID", o.StationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
