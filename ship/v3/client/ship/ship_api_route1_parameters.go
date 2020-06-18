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

// NewShipAPIRoute1Params creates a new ShipAPIRoute1Params object
// with the default values initialized.
func NewShipAPIRoute1Params() *ShipAPIRoute1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &ShipAPIRoute1Params{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewShipAPIRoute1ParamsWithTimeout creates a new ShipAPIRoute1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewShipAPIRoute1ParamsWithTimeout(timeout time.Duration) *ShipAPIRoute1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &ShipAPIRoute1Params{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewShipAPIRoute1ParamsWithContext creates a new ShipAPIRoute1Params object
// with the default values initialized, and the ability to set a context for a request
func NewShipAPIRoute1ParamsWithContext(ctx context.Context) *ShipAPIRoute1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &ShipAPIRoute1Params{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewShipAPIRoute1ParamsWithHTTPClient creates a new ShipAPIRoute1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShipAPIRoute1ParamsWithHTTPClient(client *http.Client) *ShipAPIRoute1Params {
	var (
		dollarTopDefault = int64(30)
	)
	return &ShipAPIRoute1Params{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*ShipAPIRoute1Params contains all the parameters to send to the API endpoint
for the ship Api route 1 operation typically these are written to a http.Request
*/
type ShipAPIRoute1Params struct {

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
	/*RouteID
	  指定航運路線代碼

	*/
	RouteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ship Api route 1 params
func (o *ShipAPIRoute1Params) WithTimeout(timeout time.Duration) *ShipAPIRoute1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ship Api route 1 params
func (o *ShipAPIRoute1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ship Api route 1 params
func (o *ShipAPIRoute1Params) WithContext(ctx context.Context) *ShipAPIRoute1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ship Api route 1 params
func (o *ShipAPIRoute1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ship Api route 1 params
func (o *ShipAPIRoute1Params) WithHTTPClient(client *http.Client) *ShipAPIRoute1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ship Api route 1 params
func (o *ShipAPIRoute1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the ship Api route 1 params
func (o *ShipAPIRoute1Params) WithDollarCount(dollarCount *bool) *ShipAPIRoute1Params {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the ship Api route 1 params
func (o *ShipAPIRoute1Params) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the ship Api route 1 params
func (o *ShipAPIRoute1Params) WithDollarFilter(dollarFilter *string) *ShipAPIRoute1Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the ship Api route 1 params
func (o *ShipAPIRoute1Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the ship Api route 1 params
func (o *ShipAPIRoute1Params) WithDollarFormat(dollarFormat string) *ShipAPIRoute1Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the ship Api route 1 params
func (o *ShipAPIRoute1Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the ship Api route 1 params
func (o *ShipAPIRoute1Params) WithDollarOrderby(dollarOrderby *string) *ShipAPIRoute1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the ship Api route 1 params
func (o *ShipAPIRoute1Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the ship Api route 1 params
func (o *ShipAPIRoute1Params) WithDollarSelect(dollarSelect *string) *ShipAPIRoute1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the ship Api route 1 params
func (o *ShipAPIRoute1Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the ship Api route 1 params
func (o *ShipAPIRoute1Params) WithDollarSkip(dollarSkip *string) *ShipAPIRoute1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the ship Api route 1 params
func (o *ShipAPIRoute1Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the ship Api route 1 params
func (o *ShipAPIRoute1Params) WithDollarTop(dollarTop *int64) *ShipAPIRoute1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the ship Api route 1 params
func (o *ShipAPIRoute1Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithRouteID adds the routeID to the ship Api route 1 params
func (o *ShipAPIRoute1Params) WithRouteID(routeID string) *ShipAPIRoute1Params {
	o.SetRouteID(routeID)
	return o
}

// SetRouteID adds the routeId to the ship Api route 1 params
func (o *ShipAPIRoute1Params) SetRouteID(routeID string) {
	o.RouteID = routeID
}

// WriteToRequest writes these params to a swagger request
func (o *ShipAPIRoute1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param RouteID
	if err := r.SetPathParam("RouteID", o.RouteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
