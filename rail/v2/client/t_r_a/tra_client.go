// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new t r a API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for t r a API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	TRAAPIDailyTimetable(params *TRAAPIDailyTimetableParams, opts ...ClientOption) (*TRAAPIDailyTimetableOK, error)

	TRAAPIDailyTimetable1(params *TRAAPIDailyTimetable1Params, opts ...ClientOption) (*TRAAPIDailyTimetable1OK, error)

	TRAAPIDailyTimetable2(params *TRAAPIDailyTimetable2Params, opts ...ClientOption) (*TRAAPIDailyTimetable2OK, error)

	TRAAPIDailyTimetable3(params *TRAAPIDailyTimetable3Params, opts ...ClientOption) (*TRAAPIDailyTimetable3OK, error)

	TRAAPIDailyTrainInfo(params *TRAAPIDailyTrainInfoParams, opts ...ClientOption) (*TRAAPIDailyTrainInfoOK, error)

	TRAAPIDailyTrainInfo1(params *TRAAPIDailyTrainInfo1Params, opts ...ClientOption) (*TRAAPIDailyTrainInfo1OK, error)

	TRAAPIDailyTrainInfo2(params *TRAAPIDailyTrainInfo2Params, opts ...ClientOption) (*TRAAPIDailyTrainInfo2OK, error)

	TRAAPIDailyTrainInfo3(params *TRAAPIDailyTrainInfo3Params, opts ...ClientOption) (*TRAAPIDailyTrainInfo3OK, error)

	TRAAPIGeneralTimetable(params *TRAAPIGeneralTimetableParams, opts ...ClientOption) (*TRAAPIGeneralTimetableOK, error)

	TRAAPIGeneralTimetable1(params *TRAAPIGeneralTimetable1Params, opts ...ClientOption) (*TRAAPIGeneralTimetable1OK, error)

	TRAAPIGeneralTrainInfo(params *TRAAPIGeneralTrainInfoParams, opts ...ClientOption) (*TRAAPIGeneralTrainInfoOK, error)

	TRAAPIGeneralTrainInfo1(params *TRAAPIGeneralTrainInfo1Params, opts ...ClientOption) (*TRAAPIGeneralTrainInfo1OK, error)

	TRAAPILine(params *TRAAPILineParams, opts ...ClientOption) (*TRAAPILineOK, error)

	TRAAPILiveBoard(params *TRAAPILiveBoardParams, opts ...ClientOption) (*TRAAPILiveBoardOK, error)

	TRAAPILiveBoard1(params *TRAAPILiveBoard1Params, opts ...ClientOption) (*TRAAPILiveBoard1OK, error)

	TRAAPILiveTrainDelay(params *TRAAPILiveTrainDelayParams, opts ...ClientOption) (*TRAAPILiveTrainDelayOK, error)

	TRAAPINetwork(params *TRAAPINetworkParams, opts ...ClientOption) (*TRAAPINetworkOK, error)

	TRAAPIODDailyTimetable(params *TRAAPIODDailyTimetableParams, opts ...ClientOption) (*TRAAPIODDailyTimetableOK, error)

	TRAAPIODFareStation(params *TRAAPIODFareStationParams, opts ...ClientOption) (*TRAAPIODFareStationOK, error)

	TRAAPIODFareStation1(params *TRAAPIODFareStation1Params, opts ...ClientOption) (*TRAAPIODFareStation1OK, error)

	TRAAPIShape(params *TRAAPIShapeParams, opts ...ClientOption) (*TRAAPIShapeOK, error)

	TRAAPIStation(params *TRAAPIStationParams, opts ...ClientOption) (*TRAAPIStationOK, error)

	TRAAPIStationOfLine(params *TRAAPIStationOfLineParams, opts ...ClientOption) (*TRAAPIStationOfLineOK, error)

	TRAAPIStationTimetable(params *TRAAPIStationTimetableParams, opts ...ClientOption) (*TRAAPIStationTimetableOK, error)

	TRAAPITrainType(params *TRAAPITrainTypeParams, opts ...ClientOption) (*TRAAPITrainTypeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  TRAAPIDailyTimetable 取得當天所有車次的時刻表資料s

  取得當天所有車次的時刻表資料
*/
func (a *Client) TRAAPIDailyTimetable(params *TRAAPIDailyTimetableParams, opts ...ClientOption) (*TRAAPIDailyTimetableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIDailyTimetableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_DailyTimetable",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/DailyTimetable/Today",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIDailyTimetableReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIDailyTimetableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_DailyTimetable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIDailyTimetable1 取得當天指定s 車次 的時刻表資料

  取得當天指定[車次]的時刻表資料
*/
func (a *Client) TRAAPIDailyTimetable1(params *TRAAPIDailyTimetable1Params, opts ...ClientOption) (*TRAAPIDailyTimetable1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIDailyTimetable1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_DailyTimetable_1",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/DailyTimetable/Today/TrainNo/{TrainNo}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIDailyTimetable1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIDailyTimetable1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_DailyTimetable_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIDailyTimetable2 取得指定s 日期 所有車次的時刻表資料

  取得指定[日期]所有車次的時刻表資料(台鐵提供近60天每日時刻表)
*/
func (a *Client) TRAAPIDailyTimetable2(params *TRAAPIDailyTimetable2Params, opts ...ClientOption) (*TRAAPIDailyTimetable2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIDailyTimetable2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_DailyTimetable_2",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/DailyTimetable/TrainDate/{TrainDate}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIDailyTimetable2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIDailyTimetable2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_DailyTimetable_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIDailyTimetable3 取得指定s 日期 車次 的時刻表資料

  取得指定[日期],[車次]的時刻表資料(台鐵提供近60天每日時刻表)
*/
func (a *Client) TRAAPIDailyTimetable3(params *TRAAPIDailyTimetable3Params, opts ...ClientOption) (*TRAAPIDailyTimetable3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIDailyTimetable3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_DailyTimetable_3",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/DailyTimetable/TrainNo/{TrainNo}/TrainDate/{TrainDate}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIDailyTimetable3Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIDailyTimetable3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_DailyTimetable_3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIDailyTrainInfo 取得當天所有車次的車次資料s

  取得當天所有車次的車次資料
*/
func (a *Client) TRAAPIDailyTrainInfo(params *TRAAPIDailyTrainInfoParams, opts ...ClientOption) (*TRAAPIDailyTrainInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIDailyTrainInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_DailyTrainInfo",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/DailyTrainInfo/Today",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIDailyTrainInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIDailyTrainInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_DailyTrainInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIDailyTrainInfo1 取得當天指定s 車次 的車次資料

  取得當天指定[車次]的車次資料
*/
func (a *Client) TRAAPIDailyTrainInfo1(params *TRAAPIDailyTrainInfo1Params, opts ...ClientOption) (*TRAAPIDailyTrainInfo1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIDailyTrainInfo1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_DailyTrainInfo_1",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/DailyTrainInfo/Today/TrainNo/{TrainNo}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIDailyTrainInfo1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIDailyTrainInfo1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_DailyTrainInfo_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIDailyTrainInfo2 取得指定s 日期 所有車次的車次資料

  取得指定[日期]所有車次的車次資料(台鐵提供近60天每日時刻表)
*/
func (a *Client) TRAAPIDailyTrainInfo2(params *TRAAPIDailyTrainInfo2Params, opts ...ClientOption) (*TRAAPIDailyTrainInfo2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIDailyTrainInfo2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_DailyTrainInfo_2",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/DailyTrainInfo/TrainDate/{TrainDate}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIDailyTrainInfo2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIDailyTrainInfo2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_DailyTrainInfo_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIDailyTrainInfo3 取得指定s 日期 與 車次 的車次資料

  取得指定[日期]與[車次]的車次資料(台鐵提供近60天每日時刻表)
*/
func (a *Client) TRAAPIDailyTrainInfo3(params *TRAAPIDailyTrainInfo3Params, opts ...ClientOption) (*TRAAPIDailyTrainInfo3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIDailyTrainInfo3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_DailyTrainInfo_3",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/DailyTrainInfo/TrainNo/{TrainNo}/TrainDate/{TrainDate}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIDailyTrainInfo3Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIDailyTrainInfo3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_DailyTrainInfo_3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIGeneralTimetable 取得所有車次的定期時刻表資料s

  取得所有車次的定期時刻表資料
*/
func (a *Client) TRAAPIGeneralTimetable(params *TRAAPIGeneralTimetableParams, opts ...ClientOption) (*TRAAPIGeneralTimetableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIGeneralTimetableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_GeneralTimetable",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/GeneralTimetable",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIGeneralTimetableReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIGeneralTimetableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_GeneralTimetable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIGeneralTimetable1 取得指定s 車次 的定期時刻表資料

  取得指定[車次]的定期時刻表資料
*/
func (a *Client) TRAAPIGeneralTimetable1(params *TRAAPIGeneralTimetable1Params, opts ...ClientOption) (*TRAAPIGeneralTimetable1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIGeneralTimetable1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_GeneralTimetable_1",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/GeneralTimetable/TrainNo/{TrainNo}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIGeneralTimetable1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIGeneralTimetable1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_GeneralTimetable_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIGeneralTrainInfo 取得所有車次的定期車次資料s

  取得所有車次的定期車次資料
*/
func (a *Client) TRAAPIGeneralTrainInfo(params *TRAAPIGeneralTrainInfoParams, opts ...ClientOption) (*TRAAPIGeneralTrainInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIGeneralTrainInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_GeneralTrainInfo",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/GeneralTrainInfo",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIGeneralTrainInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIGeneralTrainInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_GeneralTrainInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIGeneralTrainInfo1 取得指定s 車次 的定期車次資料

  取得指定[車次]的定期車次資料
*/
func (a *Client) TRAAPIGeneralTrainInfo1(params *TRAAPIGeneralTrainInfo1Params, opts ...ClientOption) (*TRAAPIGeneralTrainInfo1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIGeneralTrainInfo1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_GeneralTrainInfo_1",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/GeneralTrainInfo/TrainNo/{TrainNo}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIGeneralTrainInfo1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIGeneralTrainInfo1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_GeneralTrainInfo_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPILine 取得路線基本資料s

  取得路線基本資料
*/
func (a *Client) TRAAPILine(params *TRAAPILineParams, opts ...ClientOption) (*TRAAPILineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPILineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_Line",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/Line",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPILineReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPILineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_Line: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPILiveBoard 取得車站別列車即時到離站電子看板s 動態前後30分鐘的車次

  取得車站別列車即時到離站電子看板(動態前後30分鐘的車次)。更新頻率：2分鐘。此資料已過濾離站車次資訊
*/
func (a *Client) TRAAPILiveBoard(params *TRAAPILiveBoardParams, opts ...ClientOption) (*TRAAPILiveBoardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPILiveBoardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_LiveBoard",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/LiveBoard",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPILiveBoardReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPILiveBoardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_LiveBoard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPILiveBoard1 取得指定s 車站 列車即時到離站電子看板 動態前後30分鐘的車次

  取得指定[車站]列車即時到離站電子看板(動態前後30分鐘的車次)。更新頻率：2分鐘。此資料已過濾離站車次資訊
*/
func (a *Client) TRAAPILiveBoard1(params *TRAAPILiveBoard1Params, opts ...ClientOption) (*TRAAPILiveBoard1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPILiveBoard1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_LiveBoard_1",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/LiveBoard/Station/{StationID}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPILiveBoard1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPILiveBoard1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_LiveBoard_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPILiveTrainDelay 取得列車即時準點s 延誤時間資料

  取得列車即時準點/延誤時間資料。更新頻率：2分鐘
*/
func (a *Client) TRAAPILiveTrainDelay(params *TRAAPILiveTrainDelayParams, opts ...ClientOption) (*TRAAPILiveTrainDelayOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPILiveTrainDelayParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_LiveTrainDelay",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/LiveTrainDelay",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPILiveTrainDelayReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPILiveTrainDelayOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_LiveTrainDelay: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPINetwork 取得臺鐵路網資料s

  取得臺鐵路網資料
*/
func (a *Client) TRAAPINetwork(params *TRAAPINetworkParams, opts ...ClientOption) (*TRAAPINetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPINetworkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_Network",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/Network",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPINetworkReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPINetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_Network: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIODDailyTimetable 取得指定s 日期 起迄站間 之站間時刻表資料

  取得指定[日期],[起迄站間]之站間時刻表資料
*/
func (a *Client) TRAAPIODDailyTimetable(params *TRAAPIODDailyTimetableParams, opts ...ClientOption) (*TRAAPIODDailyTimetableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIODDailyTimetableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_ODDailyTimetable",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/DailyTimetable/OD/{OriginStationID}/to/{DestinationStationID}/{TrainDate}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIODDailyTimetableReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIODDailyTimetableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_ODDailyTimetable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIODFareStation 取得指定s 起訖站間 之票價資料

  取得指定[起訖站間]之票價資料
*/
func (a *Client) TRAAPIODFareStation(params *TRAAPIODFareStationParams, opts ...ClientOption) (*TRAAPIODFareStationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIODFareStationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_ODFareStation",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/ODFare/{OriginStationID}/to/{DestinationStationID}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIODFareStationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIODFareStationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_ODFareStation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIODFareStation1 取得票價資料s

  取得票價資料
*/
func (a *Client) TRAAPIODFareStation1(params *TRAAPIODFareStation1Params, opts ...ClientOption) (*TRAAPIODFareStation1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIODFareStation1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_ODFareStation_1",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/ODFare",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIODFareStation1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIODFareStation1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_ODFareStation_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIShape 取得軌道路網實體路線圖資資料s

  取得軌道路網實體路線圖資資料
*/
func (a *Client) TRAAPIShape(params *TRAAPIShapeParams, opts ...ClientOption) (*TRAAPIShapeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIShapeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_Shape",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/Shape",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIShapeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIShapeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_Shape: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIStation 取得車站基本資料s

  取得車站基本資料
*/
func (a *Client) TRAAPIStation(params *TRAAPIStationParams, opts ...ClientOption) (*TRAAPIStationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIStationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_Station",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/Station",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIStationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIStationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_Station: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIStationOfLine 取得路線車站基本資料s

  取得路線車站基本資料
*/
func (a *Client) TRAAPIStationOfLine(params *TRAAPIStationOfLineParams, opts ...ClientOption) (*TRAAPIStationOfLineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIStationOfLineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_StationOfLine",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/StationOfLine",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIStationOfLineReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIStationOfLineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_StationOfLine: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPIStationTimetable 取得指定s 日期 車站 的站別時刻表資料

  取得指定[日期],[車站]的站別時刻表資料
*/
func (a *Client) TRAAPIStationTimetable(params *TRAAPIStationTimetableParams, opts ...ClientOption) (*TRAAPIStationTimetableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPIStationTimetableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_StationTimetable",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/DailyTimetable/Station/{StationID}/{TrainDate}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPIStationTimetableReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPIStationTimetableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_StationTimetable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TRAAPITrainType 取得所有列車車種資料s

  取得所有列車車種資料
*/
func (a *Client) TRAAPITrainType(params *TRAAPITrainTypeParams, opts ...ClientOption) (*TRAAPITrainTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTRAAPITrainTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TRAApi_TrainType",
		Method:             "GET",
		PathPattern:        "/v2/Rail/TRA/TrainType",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TRAAPITrainTypeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TRAAPITrainTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TRAApi_TrainType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
