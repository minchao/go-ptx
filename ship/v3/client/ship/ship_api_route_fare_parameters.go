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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewShipAPIRouteFareParams creates a new ShipAPIRouteFareParams object
// with the default values initialized.
func NewShipAPIRouteFareParams() *ShipAPIRouteFareParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &ShipAPIRouteFareParams{
		DollarTop: &dollarTopDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewShipAPIRouteFareParamsWithTimeout creates a new ShipAPIRouteFareParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShipAPIRouteFareParamsWithTimeout(timeout time.Duration) *ShipAPIRouteFareParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &ShipAPIRouteFareParams{
		DollarTop: &dollarTopDefault,

		timeout: timeout,
	}
}

// NewShipAPIRouteFareParamsWithContext creates a new ShipAPIRouteFareParams object
// with the default values initialized, and the ability to set a context for a request
func NewShipAPIRouteFareParamsWithContext(ctx context.Context) *ShipAPIRouteFareParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &ShipAPIRouteFareParams{
		DollarTop: &dollarTopDefault,

		Context: ctx,
	}
}

// NewShipAPIRouteFareParamsWithHTTPClient creates a new ShipAPIRouteFareParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShipAPIRouteFareParamsWithHTTPClient(client *http.Client) *ShipAPIRouteFareParams {
	var (
		dollarTopDefault = int64(30)
	)
	return &ShipAPIRouteFareParams{
		DollarTop:  &dollarTopDefault,
		HTTPClient: client,
	}
}

/*ShipAPIRouteFareParams contains all the parameters to send to the API endpoint
for the ship Api route fare operation typically these are written to a http.Request
*/
type ShipAPIRouteFareParams struct {

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

// WithTimeout adds the timeout to the ship Api route fare params
func (o *ShipAPIRouteFareParams) WithTimeout(timeout time.Duration) *ShipAPIRouteFareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ship Api route fare params
func (o *ShipAPIRouteFareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ship Api route fare params
func (o *ShipAPIRouteFareParams) WithContext(ctx context.Context) *ShipAPIRouteFareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ship Api route fare params
func (o *ShipAPIRouteFareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ship Api route fare params
func (o *ShipAPIRouteFareParams) WithHTTPClient(client *http.Client) *ShipAPIRouteFareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ship Api route fare params
func (o *ShipAPIRouteFareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the ship Api route fare params
func (o *ShipAPIRouteFareParams) WithDollarCount(dollarCount *string) *ShipAPIRouteFareParams {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the ship Api route fare params
func (o *ShipAPIRouteFareParams) SetDollarCount(dollarCount *string) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the ship Api route fare params
func (o *ShipAPIRouteFareParams) WithDollarFilter(dollarFilter *string) *ShipAPIRouteFareParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the ship Api route fare params
func (o *ShipAPIRouteFareParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the ship Api route fare params
func (o *ShipAPIRouteFareParams) WithDollarFormat(dollarFormat string) *ShipAPIRouteFareParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the ship Api route fare params
func (o *ShipAPIRouteFareParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the ship Api route fare params
func (o *ShipAPIRouteFareParams) WithDollarOrderby(dollarOrderby *string) *ShipAPIRouteFareParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the ship Api route fare params
func (o *ShipAPIRouteFareParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the ship Api route fare params
func (o *ShipAPIRouteFareParams) WithDollarSelect(dollarSelect *string) *ShipAPIRouteFareParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the ship Api route fare params
func (o *ShipAPIRouteFareParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the ship Api route fare params
func (o *ShipAPIRouteFareParams) WithDollarSkip(dollarSkip *string) *ShipAPIRouteFareParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the ship Api route fare params
func (o *ShipAPIRouteFareParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the ship Api route fare params
func (o *ShipAPIRouteFareParams) WithDollarTop(dollarTop *int64) *ShipAPIRouteFareParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the ship Api route fare params
func (o *ShipAPIRouteFareParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *ShipAPIRouteFareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
