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

// NewTRAAPIODFareStation1Params creates a new TRAAPIODFareStation1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTRAAPIODFareStation1Params() *TRAAPIODFareStation1Params {
	return &TRAAPIODFareStation1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewTRAAPIODFareStation1ParamsWithTimeout creates a new TRAAPIODFareStation1Params object
// with the ability to set a timeout on a request.
func NewTRAAPIODFareStation1ParamsWithTimeout(timeout time.Duration) *TRAAPIODFareStation1Params {
	return &TRAAPIODFareStation1Params{
		timeout: timeout,
	}
}

// NewTRAAPIODFareStation1ParamsWithContext creates a new TRAAPIODFareStation1Params object
// with the ability to set a context for a request.
func NewTRAAPIODFareStation1ParamsWithContext(ctx context.Context) *TRAAPIODFareStation1Params {
	return &TRAAPIODFareStation1Params{
		Context: ctx,
	}
}

// NewTRAAPIODFareStation1ParamsWithHTTPClient creates a new TRAAPIODFareStation1Params object
// with the ability to set a custom HTTPClient for a request.
func NewTRAAPIODFareStation1ParamsWithHTTPClient(client *http.Client) *TRAAPIODFareStation1Params {
	return &TRAAPIODFareStation1Params{
		HTTPClient: client,
	}
}

/* TRAAPIODFareStation1Params contains all the parameters to send to the API endpoint
   for the t r a Api o d fare station 1 operation.

   Typically these are written to a http.Request.
*/
type TRAAPIODFareStation1Params struct {

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

	/* DestinationStationID.

	   迄點車站代碼
	*/
	DestinationStationID string

	/* OriginStationID.

	   起點車站代碼
	*/
	OriginStationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the t r a Api o d fare station 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TRAAPIODFareStation1Params) WithDefaults() *TRAAPIODFareStation1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the t r a Api o d fare station 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TRAAPIODFareStation1Params) SetDefaults() {
	var (
		dollarTopDefault = int64(30)
	)

	val := TRAAPIODFareStation1Params{
		DollarTop: &dollarTopDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) WithTimeout(timeout time.Duration) *TRAAPIODFareStation1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) WithContext(ctx context.Context) *TRAAPIODFareStation1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) WithHTTPClient(client *http.Client) *TRAAPIODFareStation1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) WithDollarFilter(dollarFilter *string) *TRAAPIODFareStation1Params {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarFormat adds the dollarFormat to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) WithDollarFormat(dollarFormat string) *TRAAPIODFareStation1Params {
	o.SetDollarFormat(dollarFormat)
	return o
}

// SetDollarFormat adds the dollarFormat to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) SetDollarFormat(dollarFormat string) {
	o.DollarFormat = dollarFormat
}

// WithDollarOrderby adds the dollarOrderby to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) WithDollarOrderby(dollarOrderby *string) *TRAAPIODFareStation1Params {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) WithDollarSelect(dollarSelect *string) *TRAAPIODFareStation1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) WithDollarSkip(dollarSkip *string) *TRAAPIODFareStation1Params {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) WithDollarTop(dollarTop *int64) *TRAAPIODFareStation1Params {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithDestinationStationID adds the destinationStationID to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) WithDestinationStationID(destinationStationID string) *TRAAPIODFareStation1Params {
	o.SetDestinationStationID(destinationStationID)
	return o
}

// SetDestinationStationID adds the destinationStationId to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) SetDestinationStationID(destinationStationID string) {
	o.DestinationStationID = destinationStationID
}

// WithOriginStationID adds the originStationID to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) WithOriginStationID(originStationID string) *TRAAPIODFareStation1Params {
	o.SetOriginStationID(originStationID)
	return o
}

// SetOriginStationID adds the originStationId to the t r a Api o d fare station 1 params
func (o *TRAAPIODFareStation1Params) SetOriginStationID(originStationID string) {
	o.OriginStationID = originStationID
}

// WriteToRequest writes these params to a swagger request
func (o *TRAAPIODFareStation1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param DestinationStationID
	if err := r.SetPathParam("DestinationStationID", o.DestinationStationID); err != nil {
		return err
	}

	// path param OriginStationID
	if err := r.SetPathParam("OriginStationID", o.OriginStationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
