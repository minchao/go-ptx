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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewMetroAPINetworkParams creates a new MetroAPINetworkParams object
// with the default values initialized.
func NewMetroAPINetworkParams() *MetroAPINetworkParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &MetroAPINetworkParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewMetroAPINetworkParamsWithTimeout creates a new MetroAPINetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMetroAPINetworkParamsWithTimeout(timeout time.Duration) *MetroAPINetworkParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &MetroAPINetworkParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewMetroAPINetworkParamsWithContext creates a new MetroAPINetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewMetroAPINetworkParamsWithContext(ctx context.Context) *MetroAPINetworkParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &MetroAPINetworkParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewMetroAPINetworkParamsWithHTTPClient creates a new MetroAPINetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMetroAPINetworkParamsWithHTTPClient(client *http.Client) *MetroAPINetworkParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &MetroAPINetworkParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*MetroAPINetworkParams contains all the parameters to send to the API endpoint
for the metro Api network operation typically these are written to a http.Request
*/
type MetroAPINetworkParams struct {

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
	/*Operator
	  欲查詢縣市

	*/
	Operator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the metro Api network params
func (o *MetroAPINetworkParams) WithTimeout(timeout time.Duration) *MetroAPINetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the metro Api network params
func (o *MetroAPINetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the metro Api network params
func (o *MetroAPINetworkParams) WithContext(ctx context.Context) *MetroAPINetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the metro Api network params
func (o *MetroAPINetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the metro Api network params
func (o *MetroAPINetworkParams) WithHTTPClient(client *http.Client) *MetroAPINetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the metro Api network params
func (o *MetroAPINetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the metro Api network params
func (o *MetroAPINetworkParams) WithDollarFilter(dollarFilter *string) *MetroAPINetworkParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the metro Api network params
func (o *MetroAPINetworkParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the metro Api network params
func (o *MetroAPINetworkParams) WithDollarFormat(dollarFormat string) *MetroAPINetworkParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the metro Api network params
func (o *MetroAPINetworkParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the metro Api network params
func (o *MetroAPINetworkParams) WithDollarOrderby(dollarOrderby *string) *MetroAPINetworkParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the metro Api network params
func (o *MetroAPINetworkParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the metro Api network params
func (o *MetroAPINetworkParams) WithDollarSelect(dollarSelect *string) *MetroAPINetworkParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the metro Api network params
func (o *MetroAPINetworkParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the metro Api network params
func (o *MetroAPINetworkParams) WithDollarSkip(dollarSkip *string) *MetroAPINetworkParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the metro Api network params
func (o *MetroAPINetworkParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the metro Api network params
func (o *MetroAPINetworkParams) WithDollarTop(dollarTop *int64) *MetroAPINetworkParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the metro Api network params
func (o *MetroAPINetworkParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithOperator adds the operator to the metro Api network params
func (o *MetroAPINetworkParams) WithOperator(operator string) *MetroAPINetworkParams {
	o.SetOperator(operator)
	return o
}

// SetOperator adds the operator to the metro Api network params
func (o *MetroAPINetworkParams) SetOperator(operator string) {
	o.Operator = operator
}

// WriteToRequest writes these params to a swagger request
func (o *MetroAPINetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param Operator
	if err := r.SetPathParam("Operator", o.Operator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
