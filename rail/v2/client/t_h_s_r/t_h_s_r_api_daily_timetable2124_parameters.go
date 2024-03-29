// Code generated by go-swagger; DO NOT EDIT.

package t_h_s_r

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

// NewTHSRAPIDailyTimetable2124Params creates a new THSRAPIDailyTimetable2124Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTHSRAPIDailyTimetable2124Params() *THSRAPIDailyTimetable2124Params {
	return &THSRAPIDailyTimetable2124Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewTHSRAPIDailyTimetable2124ParamsWithTimeout creates a new THSRAPIDailyTimetable2124Params object
// with the ability to set a timeout on a request.
func NewTHSRAPIDailyTimetable2124ParamsWithTimeout(timeout time.Duration) *THSRAPIDailyTimetable2124Params {
	return &THSRAPIDailyTimetable2124Params{
		timeout: timeout,
	}
}

// NewTHSRAPIDailyTimetable2124ParamsWithContext creates a new THSRAPIDailyTimetable2124Params object
// with the ability to set a context for a request.
func NewTHSRAPIDailyTimetable2124ParamsWithContext(ctx context.Context) *THSRAPIDailyTimetable2124Params {
	return &THSRAPIDailyTimetable2124Params{
		Context: ctx,
	}
}

// NewTHSRAPIDailyTimetable2124ParamsWithHTTPClient creates a new THSRAPIDailyTimetable2124Params object
// with the ability to set a custom HTTPClient for a request.
func NewTHSRAPIDailyTimetable2124ParamsWithHTTPClient(client *http.Client) *THSRAPIDailyTimetable2124Params {
	return &THSRAPIDailyTimetable2124Params{
		HTTPClient: client,
	}
}

/* THSRAPIDailyTimetable2124Params contains all the parameters to send to the API endpoint
   for the t h s r Api daily timetable 2124 operation.

   Typically these are written to a http.Request.
*/
type THSRAPIDailyTimetable2124Params struct {

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

	/* Health.

	   加入參數'?health=true'即可查詢此API服務的健康狀態
	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the t h s r Api daily timetable 2124 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *THSRAPIDailyTimetable2124Params) WithDefaults() *THSRAPIDailyTimetable2124Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the t h s r Api daily timetable 2124 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *THSRAPIDailyTimetable2124Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := THSRAPIDailyTimetable2124Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) WithTimeout(timeout time.Duration) *THSRAPIDailyTimetable2124Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) WithContext(ctx context.Context) *THSRAPIDailyTimetable2124Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) WithHTTPClient(client *http.Client) *THSRAPIDailyTimetable2124Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) WithDollarFilter(dollarFilter *string) *THSRAPIDailyTimetable2124Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) WithDollarFormat(dollarFormat string) *THSRAPIDailyTimetable2124Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) WithDollarOrderby(dollarOrderby *string) *THSRAPIDailyTimetable2124Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) WithDollarSelect(dollarSelect *string) *THSRAPIDailyTimetable2124Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) WithDollarSkip(dollarSkip *string) *THSRAPIDailyTimetable2124Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) WithDollarTop(dollarTop *int64) *THSRAPIDailyTimetable2124Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithHealth adds the health to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) WithHealth(health *string) *THSRAPIDailyTimetable2124Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the t h s r Api daily timetable 2124 params
func (o *THSRAPIDailyTimetable2124Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *THSRAPIDailyTimetable2124Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Health != nil {

		// query param health
		var qrHealth string

		if o.Health != nil {
			qrHealth = *o.Health
		}
		qHealth := qrHealth
		if qHealth != "" {

			if err := r.SetQueryParam("health", qHealth); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
