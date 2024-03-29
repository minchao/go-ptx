// Code generated by go-swagger; DO NOT EDIT.

package metro

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new metro API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for metro API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	MetroAPIAlert2107(params *MetroAPIAlert2107Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIAlert2107OK, *MetroAPIAlert2107Status299, error)

	MetroAPIFirstLastTimetable2099(params *MetroAPIFirstLastTimetable2099Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIFirstLastTimetable2099OK, *MetroAPIFirstLastTimetable2099Status299, error)

	MetroAPIFrequency2100(params *MetroAPIFrequency2100Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIFrequency2100OK, *MetroAPIFrequency2100Status299, error)

	MetroAPILineTransfer2094(params *MetroAPILineTransfer2094Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPILineTransfer2094OK, *MetroAPILineTransfer2094Status299, error)

	MetroAPILine2091(params *MetroAPILine2091Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPILine2091OK, *MetroAPILine2091Status299, error)

	MetroAPILiveBoard2103(params *MetroAPILiveBoard2103Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPILiveBoard2103OK, *MetroAPILiveBoard2103Status299, error)

	MetroAPINetwork2090(params *MetroAPINetwork2090Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPINetwork2090OK, *MetroAPINetwork2090Status299, error)

	MetroAPINews2106(params *MetroAPINews2106Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPINews2106OK, *MetroAPINews2106Status299, error)

	MetroAPIODFare2102(params *MetroAPIODFare2102Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIODFare2102OK, *MetroAPIODFare2102Status299, error)

	MetroAPIRoute2097(params *MetroAPIRoute2097Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIRoute2097OK, *MetroAPIRoute2097Status299, error)

	MetroAPIS2STravelTime2101(params *MetroAPIS2STravelTime2101Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIS2STravelTime2101OK, *MetroAPIS2STravelTime2101Status299, error)

	MetroAPIShape2105(params *MetroAPIShape2105Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIShape2105OK, *MetroAPIShape2105Status299, error)

	MetroAPIStationExit2096(params *MetroAPIStationExit2096Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIStationExit2096OK, *MetroAPIStationExit2096Status299, error)

	MetroAPIStationFacility2095(params *MetroAPIStationFacility2095Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIStationFacility2095OK, *MetroAPIStationFacility2095Status299, error)

	MetroAPIStationOfLine2093(params *MetroAPIStationOfLine2093Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIStationOfLine2093OK, *MetroAPIStationOfLine2093Status299, error)

	MetroAPIStationOfRoute2098(params *MetroAPIStationOfRoute2098Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIStationOfRoute2098OK, *MetroAPIStationOfRoute2098Status299, error)

	MetroAPIStationTimeTable2104(params *MetroAPIStationTimeTable2104Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIStationTimeTable2104OK, *MetroAPIStationTimeTable2104Status299, error)

	MetroAPIStation2092(params *MetroAPIStation2092Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIStation2092OK, *MetroAPIStation2092Status299, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  MetroAPIAlert2107 取得營運通阻資料s

  取得營運通阻資料
*/
func (a *Client) MetroAPIAlert2107(params *MetroAPIAlert2107Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIAlert2107OK, *MetroAPIAlert2107Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPIAlert2107Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_Alert_2107",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/Alert/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPIAlert2107Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPIAlert2107OK:
		return value, nil, nil
	case *MetroAPIAlert2107Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetroAPIFirstLastTimetable2099 取得捷運首末班車時刻表資料s

  取得捷運首末班車時刻表資料
*/
func (a *Client) MetroAPIFirstLastTimetable2099(params *MetroAPIFirstLastTimetable2099Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIFirstLastTimetable2099OK, *MetroAPIFirstLastTimetable2099Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPIFirstLastTimetable2099Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_FirstLastTimetable_2099",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/FirstLastTimetable/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPIFirstLastTimetable2099Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPIFirstLastTimetable2099OK:
		return value, nil, nil
	case *MetroAPIFirstLastTimetable2099Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetroAPIFrequency2100 取得捷運路線發車班距頻率資料s

  取得捷運路線發車班距頻率資料
*/
func (a *Client) MetroAPIFrequency2100(params *MetroAPIFrequency2100Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIFrequency2100OK, *MetroAPIFrequency2100Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPIFrequency2100Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_Frequency_2100",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/Frequency/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPIFrequency2100Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPIFrequency2100OK:
		return value, nil, nil
	case *MetroAPIFrequency2100Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetroAPILineTransfer2094 取得捷運路線站間轉乘基本資料s

  取得捷運路線站間轉乘基本資料
*/
func (a *Client) MetroAPILineTransfer2094(params *MetroAPILineTransfer2094Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPILineTransfer2094OK, *MetroAPILineTransfer2094Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPILineTransfer2094Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_LineTransfer_2094",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/LineTransfer/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPILineTransfer2094Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPILineTransfer2094OK:
		return value, nil, nil
	case *MetroAPILineTransfer2094Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetroAPILine2091 取得捷運路線基本資料s

  取得捷運路線基本資料
*/
func (a *Client) MetroAPILine2091(params *MetroAPILine2091Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPILine2091OK, *MetroAPILine2091Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPILine2091Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_Line_2091",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/Line/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPILine2091Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPILine2091OK:
		return value, nil, nil
	case *MetroAPILine2091Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetroAPILiveBoard2103 取得捷運車站別列車即時到離站電子看板資訊s

  取得捷運車站別列車即時到離站電子看板資訊
*/
func (a *Client) MetroAPILiveBoard2103(params *MetroAPILiveBoard2103Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPILiveBoard2103OK, *MetroAPILiveBoard2103Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPILiveBoard2103Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_LiveBoard_2103",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/LiveBoard/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPILiveBoard2103Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPILiveBoard2103OK:
		return value, nil, nil
	case *MetroAPILiveBoard2103Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetroAPINetwork2090 取得捷運路網資料s

  取得捷運路網資料
*/
func (a *Client) MetroAPINetwork2090(params *MetroAPINetwork2090Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPINetwork2090OK, *MetroAPINetwork2090Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPINetwork2090Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_Network_2090",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/Network/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPINetwork2090Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPINetwork2090OK:
		return value, nil, nil
	case *MetroAPINetwork2090Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetroAPINews2106 取得最新消息s

  取得最新消息
*/
func (a *Client) MetroAPINews2106(params *MetroAPINews2106Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPINews2106OK, *MetroAPINews2106Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPINews2106Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_News_2106",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/News/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPINews2106Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPINews2106OK:
		return value, nil, nil
	case *MetroAPINews2106Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetroAPIODFare2102 取得捷運起迄站間票價資料s

  取得捷運起迄站間票價資料
*/
func (a *Client) MetroAPIODFare2102(params *MetroAPIODFare2102Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIODFare2102OK, *MetroAPIODFare2102Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPIODFare2102Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_ODFare_2102",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/ODFare/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPIODFare2102Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPIODFare2102OK:
		return value, nil, nil
	case *MetroAPIODFare2102Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetroAPIRoute2097 取得捷運營運路線基本資料s

  取得捷運營運路線基本資料
*/
func (a *Client) MetroAPIRoute2097(params *MetroAPIRoute2097Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIRoute2097OK, *MetroAPIRoute2097Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPIRoute2097Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_Route_2097",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/Route/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPIRoute2097Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPIRoute2097OK:
		return value, nil, nil
	case *MetroAPIRoute2097Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetroAPIS2STravelTime2101 取得捷運列車站間運行時間資料s

  取得捷運列車站間運行時間資料
*/
func (a *Client) MetroAPIS2STravelTime2101(params *MetroAPIS2STravelTime2101Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIS2STravelTime2101OK, *MetroAPIS2STravelTime2101Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPIS2STravelTime2101Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_S2STravelTime_2101",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/S2STravelTime/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPIS2STravelTime2101Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPIS2STravelTime2101OK:
		return value, nil, nil
	case *MetroAPIS2STravelTime2101Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetroAPIShape2105 取得指定營運業者之軌道路網實體路線圖資資料s

  取得指定營運業者之軌道路網實體路線圖資資料
*/
func (a *Client) MetroAPIShape2105(params *MetroAPIShape2105Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIShape2105OK, *MetroAPIShape2105Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPIShape2105Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_Shape_2105",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/Shape/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPIShape2105Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPIShape2105OK:
		return value, nil, nil
	case *MetroAPIShape2105Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetroAPIStationExit2096 取得捷運車站出入口基本資料s

  取得捷運車站出入口基本資料
*/
func (a *Client) MetroAPIStationExit2096(params *MetroAPIStationExit2096Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIStationExit2096OK, *MetroAPIStationExit2096Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPIStationExit2096Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_StationExit_2096",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/StationExit/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPIStationExit2096Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPIStationExit2096OK:
		return value, nil, nil
	case *MetroAPIStationExit2096Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetroAPIStationFacility2095 取得捷運車站設施資料s

  取得捷運車站設施資料
*/
func (a *Client) MetroAPIStationFacility2095(params *MetroAPIStationFacility2095Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIStationFacility2095OK, *MetroAPIStationFacility2095Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPIStationFacility2095Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_StationFacility_2095",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/StationFacility/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPIStationFacility2095Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPIStationFacility2095OK:
		return value, nil, nil
	case *MetroAPIStationFacility2095Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetroAPIStationOfLine2093 取得捷運路線車站基本資料s

  取得捷運路線車站基本資料
*/
func (a *Client) MetroAPIStationOfLine2093(params *MetroAPIStationOfLine2093Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIStationOfLine2093OK, *MetroAPIStationOfLine2093Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPIStationOfLine2093Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_StationOfLine_2093",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/StationOfLine/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPIStationOfLine2093Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPIStationOfLine2093OK:
		return value, nil, nil
	case *MetroAPIStationOfLine2093Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetroAPIStationOfRoute2098 取得捷運營運路線車站基本資料s

  取得捷運營運路線車站基本資料
*/
func (a *Client) MetroAPIStationOfRoute2098(params *MetroAPIStationOfRoute2098Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIStationOfRoute2098OK, *MetroAPIStationOfRoute2098Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPIStationOfRoute2098Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_StationOfRoute_2098",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/StationOfRoute/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPIStationOfRoute2098Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPIStationOfRoute2098OK:
		return value, nil, nil
	case *MetroAPIStationOfRoute2098Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetroAPIStationTimeTable2104 取得捷運站別時刻表資料s

  取得捷運站別時刻表資料

## 使用注意事項
臺北捷運目前無提供文湖線站別時刻表，建議您可使用［取得捷運路線發車班距頻率資料］取得文湖線列車相關資訊。
*/
func (a *Client) MetroAPIStationTimeTable2104(params *MetroAPIStationTimeTable2104Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIStationTimeTable2104OK, *MetroAPIStationTimeTable2104Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPIStationTimeTable2104Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_StationTimeTable_2104",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/StationTimeTable/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPIStationTimeTable2104Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPIStationTimeTable2104OK:
		return value, nil, nil
	case *MetroAPIStationTimeTable2104Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetroAPIStation2092 取得捷運車站基本資料s

  取得捷運車站基本資料
*/
func (a *Client) MetroAPIStation2092(params *MetroAPIStation2092Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetroAPIStation2092OK, *MetroAPIStation2092Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetroAPIStation2092Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "MetroApi_Station_2092",
		Method:             "GET",
		PathPattern:        "/v2/Rail/Metro/Station/{RailSystem}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetroAPIStation2092Reader{formats: a.formats},
		AuthInfo:           authInfo,
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
	case *MetroAPIStation2092OK:
		return value, nil, nil
	case *MetroAPIStation2092Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metro: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
