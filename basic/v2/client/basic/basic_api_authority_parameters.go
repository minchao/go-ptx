// Code generated by go-swagger; DO NOT EDIT.

package basic

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

// NewBasicAPIAuthorityParams creates a new BasicAPIAuthorityParams object
// with the default values initialized.
func NewBasicAPIAuthorityParams() *BasicAPIAuthorityParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &BasicAPIAuthorityParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewBasicAPIAuthorityParamsWithTimeout creates a new BasicAPIAuthorityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBasicAPIAuthorityParamsWithTimeout(timeout time.Duration) *BasicAPIAuthorityParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &BasicAPIAuthorityParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewBasicAPIAuthorityParamsWithContext creates a new BasicAPIAuthorityParams object
// with the default values initialized, and the ability to set a context for a request
func NewBasicAPIAuthorityParamsWithContext(ctx context.Context) *BasicAPIAuthorityParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &BasicAPIAuthorityParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewBasicAPIAuthorityParamsWithHTTPClient creates a new BasicAPIAuthorityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBasicAPIAuthorityParamsWithHTTPClient(client *http.Client) *BasicAPIAuthorityParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &BasicAPIAuthorityParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*BasicAPIAuthorityParams contains all the parameters to send to the API endpoint
for the basic Api authority operation typically these are written to a http.Request
*/
type BasicAPIAuthorityParams struct {

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

// WithTimeout adds the timeout to the basic Api authority params
func (o *BasicAPIAuthorityParams) WithTimeout(timeout time.Duration) *BasicAPIAuthorityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the basic Api authority params
func (o *BasicAPIAuthorityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the basic Api authority params
func (o *BasicAPIAuthorityParams) WithContext(ctx context.Context) *BasicAPIAuthorityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the basic Api authority params
func (o *BasicAPIAuthorityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the basic Api authority params
func (o *BasicAPIAuthorityParams) WithHTTPClient(client *http.Client) *BasicAPIAuthorityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the basic Api authority params
func (o *BasicAPIAuthorityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the basic Api authority params
func (o *BasicAPIAuthorityParams) WithDollarFilter(dollarFilter *string) *BasicAPIAuthorityParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the basic Api authority params
func (o *BasicAPIAuthorityParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the basic Api authority params
func (o *BasicAPIAuthorityParams) WithDollarFormat(dollarFormat string) *BasicAPIAuthorityParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the basic Api authority params
func (o *BasicAPIAuthorityParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the basic Api authority params
func (o *BasicAPIAuthorityParams) WithDollarOrderby(dollarOrderby *string) *BasicAPIAuthorityParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the basic Api authority params
func (o *BasicAPIAuthorityParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the basic Api authority params
func (o *BasicAPIAuthorityParams) WithDollarSelect(dollarSelect *string) *BasicAPIAuthorityParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the basic Api authority params
func (o *BasicAPIAuthorityParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the basic Api authority params
func (o *BasicAPIAuthorityParams) WithDollarSkip(dollarSkip *string) *BasicAPIAuthorityParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the basic Api authority params
func (o *BasicAPIAuthorityParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the basic Api authority params
func (o *BasicAPIAuthorityParams) WithDollarTop(dollarTop *int64) *BasicAPIAuthorityParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the basic Api authority params
func (o *BasicAPIAuthorityParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *BasicAPIAuthorityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
