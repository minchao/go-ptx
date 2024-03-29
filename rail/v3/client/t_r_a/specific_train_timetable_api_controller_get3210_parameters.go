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

// NewSpecificTrainTimetableAPIControllerGet3210Params creates a new SpecificTrainTimetableAPIControllerGet3210Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSpecificTrainTimetableAPIControllerGet3210Params() *SpecificTrainTimetableAPIControllerGet3210Params {
	return &SpecificTrainTimetableAPIControllerGet3210Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewSpecificTrainTimetableAPIControllerGet3210ParamsWithTimeout creates a new SpecificTrainTimetableAPIControllerGet3210Params object
// with the ability to set a timeout on a request.
func NewSpecificTrainTimetableAPIControllerGet3210ParamsWithTimeout(timeout time.Duration) *SpecificTrainTimetableAPIControllerGet3210Params {
	return &SpecificTrainTimetableAPIControllerGet3210Params{
		timeout: timeout,
	}
}

// NewSpecificTrainTimetableAPIControllerGet3210ParamsWithContext creates a new SpecificTrainTimetableAPIControllerGet3210Params object
// with the ability to set a context for a request.
func NewSpecificTrainTimetableAPIControllerGet3210ParamsWithContext(ctx context.Context) *SpecificTrainTimetableAPIControllerGet3210Params {
	return &SpecificTrainTimetableAPIControllerGet3210Params{
		Context: ctx,
	}
}

// NewSpecificTrainTimetableAPIControllerGet3210ParamsWithHTTPClient creates a new SpecificTrainTimetableAPIControllerGet3210Params object
// with the ability to set a custom HTTPClient for a request.
func NewSpecificTrainTimetableAPIControllerGet3210ParamsWithHTTPClient(client *http.Client) *SpecificTrainTimetableAPIControllerGet3210Params {
	return &SpecificTrainTimetableAPIControllerGet3210Params{
		HTTPClient: client,
	}
}

/* SpecificTrainTimetableAPIControllerGet3210Params contains all the parameters to send to the API endpoint
   for the specific train timetable Api controller get 3210 operation.

   Typically these are written to a http.Request.
*/
type SpecificTrainTimetableAPIControllerGet3210Params struct {

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

	/* Health.

	   加入參數'?health=true'即可查詢此API服務的健康狀態
	*/
	Health *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the specific train timetable Api controller get 3210 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SpecificTrainTimetableAPIControllerGet3210Params) WithDefaults() *SpecificTrainTimetableAPIControllerGet3210Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the specific train timetable Api controller get 3210 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SpecificTrainTimetableAPIControllerGet3210Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := SpecificTrainTimetableAPIControllerGet3210Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) WithTimeout(timeout time.Duration) *SpecificTrainTimetableAPIControllerGet3210Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) WithContext(ctx context.Context) *SpecificTrainTimetableAPIControllerGet3210Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) WithHTTPClient(client *http.Client) *SpecificTrainTimetableAPIControllerGet3210Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) WithDollarCount(dollarCount *bool) *SpecificTrainTimetableAPIControllerGet3210Params {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) WithDollarFilter(dollarFilter *string) *SpecificTrainTimetableAPIControllerGet3210Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) WithDollarFormat(dollarFormat string) *SpecificTrainTimetableAPIControllerGet3210Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) WithDollarOrderby(dollarOrderby *string) *SpecificTrainTimetableAPIControllerGet3210Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) WithDollarSelect(dollarSelect *string) *SpecificTrainTimetableAPIControllerGet3210Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) WithDollarSkip(dollarSkip *string) *SpecificTrainTimetableAPIControllerGet3210Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) WithDollarTop(dollarTop *int64) *SpecificTrainTimetableAPIControllerGet3210Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithHealth adds the health to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) WithHealth(health *string) *SpecificTrainTimetableAPIControllerGet3210Params {
	o.SetHealth(health)
	return o
}

// SetHealth adds the health to the specific train timetable Api controller get 3210 params
func (o *SpecificTrainTimetableAPIControllerGet3210Params) SetHealth(health *string) {
	o.Health = health
}

// WriteToRequest writes these params to a swagger request
func (o *SpecificTrainTimetableAPIControllerGet3210Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
