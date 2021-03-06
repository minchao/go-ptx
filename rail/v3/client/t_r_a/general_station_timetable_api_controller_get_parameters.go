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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGeneralStationTimetableAPIControllerGetParams creates a new GeneralStationTimetableAPIControllerGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGeneralStationTimetableAPIControllerGetParams() *GeneralStationTimetableAPIControllerGetParams {
	return &GeneralStationTimetableAPIControllerGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGeneralStationTimetableAPIControllerGetParamsWithTimeout creates a new GeneralStationTimetableAPIControllerGetParams object
// with the ability to set a timeout on a request.
func NewGeneralStationTimetableAPIControllerGetParamsWithTimeout(timeout time.Duration) *GeneralStationTimetableAPIControllerGetParams {
	return &GeneralStationTimetableAPIControllerGetParams{
		timeout: timeout,
	}
}

// NewGeneralStationTimetableAPIControllerGetParamsWithContext creates a new GeneralStationTimetableAPIControllerGetParams object
// with the ability to set a context for a request.
func NewGeneralStationTimetableAPIControllerGetParamsWithContext(ctx context.Context) *GeneralStationTimetableAPIControllerGetParams {
	return &GeneralStationTimetableAPIControllerGetParams{
		Context: ctx,
	}
}

// NewGeneralStationTimetableAPIControllerGetParamsWithHTTPClient creates a new GeneralStationTimetableAPIControllerGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGeneralStationTimetableAPIControllerGetParamsWithHTTPClient(client *http.Client) *GeneralStationTimetableAPIControllerGetParams {
	return &GeneralStationTimetableAPIControllerGetParams{
		HTTPClient: client,
	}
}

/* GeneralStationTimetableAPIControllerGetParams contains all the parameters to send to the API endpoint
   for the general station timetable Api controller get operation.

   Typically these are written to a http.Request.
*/
type GeneralStationTimetableAPIControllerGetParams struct {

	/* DollarCount.

	   查詢數量
	*/
	DollarCount *bool

	/* DollarFilter.

	   過濾
	*/
	DollarFilter *string

	/* DollarFormat.

	   指定來源格式
	*/
	DollarFormat string

	/* DollarOrderby.

	   排序
	*/
	DollarOrderby *string

	/* DollarSelect.

	   挑選
	*/
	DollarSelect *string

	/* DollarSkip.

	   跳過前幾筆
	*/
	DollarSkip *string

	/* DollarTop.

	   取前幾筆

	   Default: 30
	*/
	DollarTop *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the general station timetable Api controller get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GeneralStationTimetableAPIControllerGetParams) WithDefaults() *GeneralStationTimetableAPIControllerGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the general station timetable Api controller get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GeneralStationTimetableAPIControllerGetParams) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := GeneralStationTimetableAPIControllerGetParams{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) WithTimeout(timeout time.Duration) *GeneralStationTimetableAPIControllerGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) WithContext(ctx context.Context) *GeneralStationTimetableAPIControllerGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) WithHTTPClient(client *http.Client) *GeneralStationTimetableAPIControllerGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) WithDollarCount(dollarCount *bool) *GeneralStationTimetableAPIControllerGetParams {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) WithDollarFilter(dollarFilter *string) *GeneralStationTimetableAPIControllerGetParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) WithDollarFormat(dollarFormat string) *GeneralStationTimetableAPIControllerGetParams {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) WithDollarOrderby(dollarOrderby *string) *GeneralStationTimetableAPIControllerGetParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) WithDollarSelect(dollarSelect *string) *GeneralStationTimetableAPIControllerGetParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) WithDollarSkip(dollarSkip *string) *GeneralStationTimetableAPIControllerGetParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) WithDollarTop(dollarTop *int64) *GeneralStationTimetableAPIControllerGetParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the general station timetable Api controller get params
func (o *GeneralStationTimetableAPIControllerGetParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WriteToRequest writes these params to a swagger request
func (o *GeneralStationTimetableAPIControllerGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
