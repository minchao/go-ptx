// Code generated by go-swagger; DO NOT EDIT.

package a_f_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new a f r API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for a f r API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	LiteTrainGeneralTrainTimetable(params *LiteTrainGeneralTrainTimetableParams, opts ...ClientOption) (*LiteTrainGeneralTrainTimetableOK, *LiteTrainGeneralTrainTimetableStatus299, error)

	LiteTrainGeneralTrainTimetable1(params *LiteTrainGeneralTrainTimetable1Params, opts ...ClientOption) (*LiteTrainGeneralTrainTimetable1OK, *LiteTrainGeneralTrainTimetable1Status299, error)

	LiteTrainLine(params *LiteTrainLineParams, opts ...ClientOption) (*LiteTrainLineOK, *LiteTrainLineStatus299, error)

	LiteTrainNetwork(params *LiteTrainNetworkParams, opts ...ClientOption) (*LiteTrainNetworkOK, *LiteTrainNetworkStatus299, error)

	LiteTrainNews(params *LiteTrainNewsParams, opts ...ClientOption) (*LiteTrainNewsOK, *LiteTrainNewsStatus299, error)

	LiteTrainODFare(params *LiteTrainODFareParams, opts ...ClientOption) (*LiteTrainODFareOK, *LiteTrainODFareStatus299, error)

	LiteTrainODFare1(params *LiteTrainODFare1Params, opts ...ClientOption) (*LiteTrainODFare1OK, *LiteTrainODFare1Status299, error)

	LiteTrainOperator(params *LiteTrainOperatorParams, opts ...ClientOption) (*LiteTrainOperatorOK, *LiteTrainOperatorStatus299, error)

	LiteTrainRoute(params *LiteTrainRouteParams, opts ...ClientOption) (*LiteTrainRouteOK, *LiteTrainRouteStatus299, error)

	LiteTrainStation(params *LiteTrainStationParams, opts ...ClientOption) (*LiteTrainStationOK, *LiteTrainStationStatus299, error)

	LiteTrainStationOfLine(params *LiteTrainStationOfLineParams, opts ...ClientOption) (*LiteTrainStationOfLineOK, *LiteTrainStationOfLineStatus299, error)

	LiteTrainStationOfRoute(params *LiteTrainStationOfRouteParams, opts ...ClientOption) (*LiteTrainStationOfRouteOK, *LiteTrainStationOfRouteStatus299, error)

	LiteTrainTrainType(params *LiteTrainTrainTypeParams, opts ...ClientOption) (*LiteTrainTrainTypeOK, *LiteTrainTrainTypeStatus299, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  LiteTrainGeneralTrainTimetable 取得所有車次的定期時刻表資料s

  取得所有車次的定期時刻表資料
*/
func (a *Client) LiteTrainGeneralTrainTimetable(params *LiteTrainGeneralTrainTimetableParams, opts ...ClientOption) (*LiteTrainGeneralTrainTimetableOK, *LiteTrainGeneralTrainTimetableStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLiteTrainGeneralTrainTimetableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LiteTrain_GeneralTrainTimetable",
		Method:             "GET",
		PathPattern:        "/v3/Rail/AFR/GeneralTrainTimetable",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LiteTrainGeneralTrainTimetableReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *LiteTrainGeneralTrainTimetableOK:
		return value, nil, nil
	case *LiteTrainGeneralTrainTimetableStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for a_f_r: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LiteTrainGeneralTrainTimetable1 取得指定s 車次 的定期時刻表資料

  取得指定[車次]的定期時刻表資料
*/
func (a *Client) LiteTrainGeneralTrainTimetable1(params *LiteTrainGeneralTrainTimetable1Params, opts ...ClientOption) (*LiteTrainGeneralTrainTimetable1OK, *LiteTrainGeneralTrainTimetable1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLiteTrainGeneralTrainTimetable1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "LiteTrain_GeneralTrainTimetable_1",
		Method:             "GET",
		PathPattern:        "/v3/Rail/AFR/GeneralTrainTimetable/TrainNo/{TrainNo}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LiteTrainGeneralTrainTimetable1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *LiteTrainGeneralTrainTimetable1OK:
		return value, nil, nil
	case *LiteTrainGeneralTrainTimetable1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for a_f_r: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LiteTrainLine 取得路線基本資料s

  取得路線基本資料
*/
func (a *Client) LiteTrainLine(params *LiteTrainLineParams, opts ...ClientOption) (*LiteTrainLineOK, *LiteTrainLineStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLiteTrainLineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LiteTrain_Line",
		Method:             "GET",
		PathPattern:        "/v3/Rail/AFR/Line",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LiteTrainLineReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *LiteTrainLineOK:
		return value, nil, nil
	case *LiteTrainLineStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for a_f_r: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LiteTrainNetwork 取得路網資料s

  取得路網資料
*/
func (a *Client) LiteTrainNetwork(params *LiteTrainNetworkParams, opts ...ClientOption) (*LiteTrainNetworkOK, *LiteTrainNetworkStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLiteTrainNetworkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LiteTrain_Network",
		Method:             "GET",
		PathPattern:        "/v3/Rail/AFR/Network",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LiteTrainNetworkReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *LiteTrainNetworkOK:
		return value, nil, nil
	case *LiteTrainNetworkStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for a_f_r: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LiteTrainNews 取得最新消息s

  取得最新消息
*/
func (a *Client) LiteTrainNews(params *LiteTrainNewsParams, opts ...ClientOption) (*LiteTrainNewsOK, *LiteTrainNewsStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLiteTrainNewsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LiteTrain_News",
		Method:             "GET",
		PathPattern:        "/v3/Rail/AFR/News",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LiteTrainNewsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *LiteTrainNewsOK:
		return value, nil, nil
	case *LiteTrainNewsStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for a_f_r: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LiteTrainODFare 取得所有票價資料s

  取得所有票價資料
*/
func (a *Client) LiteTrainODFare(params *LiteTrainODFareParams, opts ...ClientOption) (*LiteTrainODFareOK, *LiteTrainODFareStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLiteTrainODFareParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LiteTrain_ODFare",
		Method:             "GET",
		PathPattern:        "/v3/Rail/AFR/ODFare",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LiteTrainODFareReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *LiteTrainODFareOK:
		return value, nil, nil
	case *LiteTrainODFareStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for a_f_r: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LiteTrainODFare1 取得指定起迄站間票價資料s

  取得指定起迄站間票價資料
*/
func (a *Client) LiteTrainODFare1(params *LiteTrainODFare1Params, opts ...ClientOption) (*LiteTrainODFare1OK, *LiteTrainODFare1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLiteTrainODFare1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "LiteTrain_ODFare_1",
		Method:             "GET",
		PathPattern:        "/v3/Rail/AFR/ODFare/{OriginStationID}/to/{DestinationStationID}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LiteTrainODFare1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *LiteTrainODFare1OK:
		return value, nil, nil
	case *LiteTrainODFare1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for a_f_r: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LiteTrainOperator 取得營運業者基本資料s

  取得營運業者基本資料
*/
func (a *Client) LiteTrainOperator(params *LiteTrainOperatorParams, opts ...ClientOption) (*LiteTrainOperatorOK, *LiteTrainOperatorStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLiteTrainOperatorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LiteTrain_Operator",
		Method:             "GET",
		PathPattern:        "/v3/Rail/AFR/Operator",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LiteTrainOperatorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *LiteTrainOperatorOK:
		return value, nil, nil
	case *LiteTrainOperatorStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for a_f_r: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LiteTrainRoute 取得營運路線基本資料s

  取得營運路線基本資料
*/
func (a *Client) LiteTrainRoute(params *LiteTrainRouteParams, opts ...ClientOption) (*LiteTrainRouteOK, *LiteTrainRouteStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLiteTrainRouteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LiteTrain_Route",
		Method:             "GET",
		PathPattern:        "/v3/Rail/AFR/Route",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LiteTrainRouteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *LiteTrainRouteOK:
		return value, nil, nil
	case *LiteTrainRouteStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for a_f_r: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LiteTrainStation 取得車站基本資料s

  取得車站基本資料
*/
func (a *Client) LiteTrainStation(params *LiteTrainStationParams, opts ...ClientOption) (*LiteTrainStationOK, *LiteTrainStationStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLiteTrainStationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LiteTrain_Station",
		Method:             "GET",
		PathPattern:        "/v3/Rail/AFR/Station",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LiteTrainStationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *LiteTrainStationOK:
		return value, nil, nil
	case *LiteTrainStationStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for a_f_r: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LiteTrainStationOfLine 取得路線車站基本資料s

  取得路線車站基本資料
*/
func (a *Client) LiteTrainStationOfLine(params *LiteTrainStationOfLineParams, opts ...ClientOption) (*LiteTrainStationOfLineOK, *LiteTrainStationOfLineStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLiteTrainStationOfLineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LiteTrain_StationOfLine",
		Method:             "GET",
		PathPattern:        "/v3/Rail/AFR/StationOfLine",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LiteTrainStationOfLineReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *LiteTrainStationOfLineOK:
		return value, nil, nil
	case *LiteTrainStationOfLineStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for a_f_r: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LiteTrainStationOfRoute 取得營運路線車站基本資料s

  取得營運路線車站基本資料
*/
func (a *Client) LiteTrainStationOfRoute(params *LiteTrainStationOfRouteParams, opts ...ClientOption) (*LiteTrainStationOfRouteOK, *LiteTrainStationOfRouteStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLiteTrainStationOfRouteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LiteTrain_StationOfRoute",
		Method:             "GET",
		PathPattern:        "/v3/Rail/AFR/StationOfRoute",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LiteTrainStationOfRouteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *LiteTrainStationOfRouteOK:
		return value, nil, nil
	case *LiteTrainStationOfRouteStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for a_f_r: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LiteTrainTrainType 取得所有列車車種資料s

  取得所有列車車種資料
*/
func (a *Client) LiteTrainTrainType(params *LiteTrainTrainTypeParams, opts ...ClientOption) (*LiteTrainTrainTypeOK, *LiteTrainTrainTypeStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLiteTrainTrainTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LiteTrain_TrainType",
		Method:             "GET",
		PathPattern:        "/v3/Rail/AFR/TrainType",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LiteTrainTrainTypeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *LiteTrainTrainTypeOK:
		return value, nil, nil
	case *LiteTrainTrainTypeStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for a_f_r: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
