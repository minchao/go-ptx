// Code generated by go-swagger; DO NOT EDIT.

package air

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

// NewAirAPIFIDS1Params creates a new AirAPIFIDS1Params object
// with the default values initialized.
func NewAirAPIFIDS1Params() *AirAPIFIDS1Params {
	var (
		dollarTopDefault = string("30")
	)
	return &AirAPIFIDS1Params{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewAirAPIFIDS1ParamsWithTimeout creates a new AirAPIFIDS1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAirAPIFIDS1ParamsWithTimeout(timeout time.Duration) *AirAPIFIDS1Params {
	var (
		dollarTopDefault = string("30")
	)
	return &AirAPIFIDS1Params{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewAirAPIFIDS1ParamsWithContext creates a new AirAPIFIDS1Params object
// with the default values initialized, and the ability to set a context for a request
func NewAirAPIFIDS1ParamsWithContext(ctx context.Context) *AirAPIFIDS1Params {
	var (
		dollarTopDefault = string("30")
	)
	return &AirAPIFIDS1Params{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewAirAPIFIDS1ParamsWithHTTPClient creates a new AirAPIFIDS1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAirAPIFIDS1ParamsWithHTTPClient(client *http.Client) *AirAPIFIDS1Params {
	var (
		dollarTopDefault = string("30")
	)
	return &AirAPIFIDS1Params{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*AirAPIFIDS1Params contains all the parameters to send to the API endpoint
for the air Api f ID s 1 operation typically these are written to a http.Request
*/
type AirAPIFIDS1Params struct {

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
	/*IATA
	  機場代碼

	*/
	IATA string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) WithTimeout(timeout time.Duration) *AirAPIFIDS1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) WithContext(ctx context.Context) *AirAPIFIDS1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) WithHTTPClient(client *http.Client) *AirAPIFIDS1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) WithDollarFilter(dollarFilter *string) *AirAPIFIDS1Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) WithDollarFormat(dollarFormat string) *AirAPIFIDS1Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) WithDollarOrderby(dollarOrderby *string) *AirAPIFIDS1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) WithDollarSelect(dollarSelect *string) *AirAPIFIDS1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) WithDollarSkip(dollarSkip *string) *AirAPIFIDS1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) WithDollarTop(dollarTop *string) *AirAPIFIDS1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) SetDollarTop(dollarTop *string) {
	o.DollarTop = dollarTop
}

// WithIATA adds the iATA to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) WithIATA(iATA string) *AirAPIFIDS1Params {
	o.SetIATA(iATA)
	return o
}

// SetIATA adds the iATA to the air Api f ID s 1 params
func (o *AirAPIFIDS1Params) SetIATA(iATA string) {
	o.IATA = iATA
}

// WriteToRequest writes these params to a swagger request
func (o *AirAPIFIDS1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param IATA
	if err := r.SetPathParam("IATA", o.IATA); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
