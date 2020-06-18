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

// NewShipAPIRouteParams creates a new ShipAPIRouteParams object
// with the default values initialized.
func NewShipAPIRouteParams() *ShipAPIRouteParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &ShipAPIRouteParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewShipAPIRouteParamsWithTimeout creates a new ShipAPIRouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShipAPIRouteParamsWithTimeout(timeout time.Duration) *ShipAPIRouteParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &ShipAPIRouteParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewShipAPIRouteParamsWithContext creates a new ShipAPIRouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewShipAPIRouteParamsWithContext(ctx context.Context) *ShipAPIRouteParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &ShipAPIRouteParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewShipAPIRouteParamsWithHTTPClient creates a new ShipAPIRouteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShipAPIRouteParamsWithHTTPClient(client *http.Client) *ShipAPIRouteParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &ShipAPIRouteParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*ShipAPIRouteParams contains all the parameters to send to the API endpoint
for the ship Api route operation typically these are written to a http.Request
*/
type ShipAPIRouteParams struct {

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

// WithTimeout adds the timeout to the ship Api route params
func (o *ShipAPIRouteParams) WithTimeout(timeout time.Duration) *ShipAPIRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ship Api route params
func (o *ShipAPIRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ship Api route params
func (o *ShipAPIRouteParams) WithContext(ctx context.Context) *ShipAPIRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ship Api route params
func (o *ShipAPIRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ship Api route params
func (o *ShipAPIRouteParams) WithHTTPClient(client *http.Client) *ShipAPIRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ship Api route params
func (o *ShipAPIRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the ship Api route params
func (o *ShipAPIRouteParams) WithDollarCount(dollarCount *bool) *ShipAPIRouteParams {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the ship Api route params
func (o *ShipAPIRouteParams) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the ship Api route params
func (o *ShipAPIRouteParams) WithDollarFilter(dollarFilter *string) *ShipAPIRouteParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the ship Api route params
func (o *ShipAPIRouteParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the ship Api route params
func (o *ShipAPIRouteParams) WithDollarFormat(dollarFormat string) *ShipAPIRouteParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the ship Api route params
func (o *ShipAPIRouteParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the ship Api route params
func (o *ShipAPIRouteParams) WithDollarOrderby(dollarOrderby *string) *ShipAPIRouteParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the ship Api route params
func (o *ShipAPIRouteParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the ship Api route params
func (o *ShipAPIRouteParams) WithDollarSelect(dollarSelect *string) *ShipAPIRouteParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the ship Api route params
func (o *ShipAPIRouteParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the ship Api route params
func (o *ShipAPIRouteParams) WithDollarSkip(dollarSkip *string) *ShipAPIRouteParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the ship Api route params
func (o *ShipAPIRouteParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the ship Api route params
func (o *ShipAPIRouteParams) WithDollarTop(dollarTop *int64) *ShipAPIRouteParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the ship Api route params
func (o *ShipAPIRouteParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *ShipAPIRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
