// Code generated by go-swagger; DO NOT EDIT.

package ship

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

// NewShipAPIAuthorityParams creates a new ShipAPIAuthorityParams object
// with the default values initialized.
func NewShipAPIAuthorityParams() *ShipAPIAuthorityParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &ShipAPIAuthorityParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewShipAPIAuthorityParamsWithTimeout creates a new ShipAPIAuthorityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShipAPIAuthorityParamsWithTimeout(timeout time.Duration) *ShipAPIAuthorityParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &ShipAPIAuthorityParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewShipAPIAuthorityParamsWithContext creates a new ShipAPIAuthorityParams object
// with the default values initialized, and the ability to set a context for a request
func NewShipAPIAuthorityParamsWithContext(ctx context.Context) *ShipAPIAuthorityParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &ShipAPIAuthorityParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewShipAPIAuthorityParamsWithHTTPClient creates a new ShipAPIAuthorityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShipAPIAuthorityParamsWithHTTPClient(client *http.Client) *ShipAPIAuthorityParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &ShipAPIAuthorityParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*ShipAPIAuthorityParams contains all the parameters to send to the API endpoint
for the ship Api authority operation typically these are written to a http.Request
*/
type ShipAPIAuthorityParams struct {

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

// WithTimeout adds the timeout to the ship Api authority params
func (o *ShipAPIAuthorityParams) WithTimeout(timeout time.Duration) *ShipAPIAuthorityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ship Api authority params
func (o *ShipAPIAuthorityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ship Api authority params
func (o *ShipAPIAuthorityParams) WithContext(ctx context.Context) *ShipAPIAuthorityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ship Api authority params
func (o *ShipAPIAuthorityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ship Api authority params
func (o *ShipAPIAuthorityParams) WithHTTPClient(client *http.Client) *ShipAPIAuthorityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ship Api authority params
func (o *ShipAPIAuthorityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the ship Api authority params
func (o *ShipAPIAuthorityParams) WithDollarCount(dollarCount *bool) *ShipAPIAuthorityParams {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the ship Api authority params
func (o *ShipAPIAuthorityParams) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the ship Api authority params
func (o *ShipAPIAuthorityParams) WithDollarFilter(dollarFilter *string) *ShipAPIAuthorityParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the ship Api authority params
func (o *ShipAPIAuthorityParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the ship Api authority params
func (o *ShipAPIAuthorityParams) WithDollarFormat(dollarFormat string) *ShipAPIAuthorityParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the ship Api authority params
func (o *ShipAPIAuthorityParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the ship Api authority params
func (o *ShipAPIAuthorityParams) WithDollarOrderby(dollarOrderby *string) *ShipAPIAuthorityParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the ship Api authority params
func (o *ShipAPIAuthorityParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the ship Api authority params
func (o *ShipAPIAuthorityParams) WithDollarSelect(dollarSelect *string) *ShipAPIAuthorityParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the ship Api authority params
func (o *ShipAPIAuthorityParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the ship Api authority params
func (o *ShipAPIAuthorityParams) WithDollarSkip(dollarSkip *string) *ShipAPIAuthorityParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the ship Api authority params
func (o *ShipAPIAuthorityParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the ship Api authority params
func (o *ShipAPIAuthorityParams) WithDollarTop(dollarTop *int64) *ShipAPIAuthorityParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the ship Api authority params
func (o *ShipAPIAuthorityParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *ShipAPIAuthorityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
