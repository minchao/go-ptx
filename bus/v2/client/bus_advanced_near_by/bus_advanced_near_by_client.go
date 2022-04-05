// Code generated by go-swagger; DO NOT EDIT.

package bus_advanced_near_by

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new bus advanced near by API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bus advanced near by API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BusAPIEstimatedTimeOfArrivalNearBy2855(params *BusAPIEstimatedTimeOfArrivalNearBy2855Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BusAPIEstimatedTimeOfArrivalNearBy2855OK, *BusAPIEstimatedTimeOfArrivalNearBy2855Status299, error)

	BusAPIRealTimeByFrequencyNearBy2853(params *BusAPIRealTimeByFrequencyNearBy2853Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BusAPIRealTimeByFrequencyNearBy2853OK, *BusAPIRealTimeByFrequencyNearBy2853Status299, error)

	BusAPIRealTimeNearStopNearBy2854(params *BusAPIRealTimeNearStopNearBy2854Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BusAPIRealTimeNearStopNearBy2854OK, *BusAPIRealTimeNearStopNearBy2854Status299, error)

	BusAPIRouteNearBy2852(params *BusAPIRouteNearBy2852Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BusAPIRouteNearBy2852OK, *BusAPIRouteNearBy2852Status299, error)

	BusAPIStationNearBy2851(params *BusAPIStationNearBy2851Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BusAPIStationNearBy2851OK, *BusAPIStationNearBy2851Status299, error)

	BusAPIStopNearBy2850(params *BusAPIStopNearBy2850Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BusAPIStopNearBy2850OK, *BusAPIStopNearBy2850Status299, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BusAPIEstimatedTimeOfArrivalNearBy2855 取得指定s 位置 範圍 的全臺公車預估到站資料 n1

  取得指定[位置,範圍]的全臺公車預估到站資料(N1)
*/
func (a *Client) BusAPIEstimatedTimeOfArrivalNearBy2855(params *BusAPIEstimatedTimeOfArrivalNearBy2855Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BusAPIEstimatedTimeOfArrivalNearBy2855OK, *BusAPIEstimatedTimeOfArrivalNearBy2855Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBusAPIEstimatedTimeOfArrivalNearBy2855Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "BusApi_EstimatedTimeOfArrival_NearBy_2855",
		Method:             "GET",
		PathPattern:        "/v2/Bus/EstimatedTimeOfArrival/NearBy",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BusAPIEstimatedTimeOfArrivalNearBy2855Reader{formats: a.formats},
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
	case *BusAPIEstimatedTimeOfArrivalNearBy2855OK:
		return value, nil, nil
	case *BusAPIEstimatedTimeOfArrivalNearBy2855Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for bus_advanced_near_by: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BusAPIRealTimeByFrequencyNearBy2853 取得指定s 位置 範圍 的全臺公車動態定時資料 a1

  取得指定[位置,範圍]的全臺公車動態定時資料(A1)
*/
func (a *Client) BusAPIRealTimeByFrequencyNearBy2853(params *BusAPIRealTimeByFrequencyNearBy2853Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BusAPIRealTimeByFrequencyNearBy2853OK, *BusAPIRealTimeByFrequencyNearBy2853Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBusAPIRealTimeByFrequencyNearBy2853Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "BusApi_RealTimeByFrequency_NearBy_2853",
		Method:             "GET",
		PathPattern:        "/v2/Bus/RealTimeByFrequency/NearBy",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BusAPIRealTimeByFrequencyNearBy2853Reader{formats: a.formats},
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
	case *BusAPIRealTimeByFrequencyNearBy2853OK:
		return value, nil, nil
	case *BusAPIRealTimeByFrequencyNearBy2853Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for bus_advanced_near_by: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BusAPIRealTimeNearStopNearBy2854 取得指定s 位置 範圍 的全臺公車動態定點資料 a2

  取得指定[位置,範圍]的全臺公車動態定點資料(A2)
*/
func (a *Client) BusAPIRealTimeNearStopNearBy2854(params *BusAPIRealTimeNearStopNearBy2854Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BusAPIRealTimeNearStopNearBy2854OK, *BusAPIRealTimeNearStopNearBy2854Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBusAPIRealTimeNearStopNearBy2854Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "BusApi_RealTimeNearStop_NearBy_2854",
		Method:             "GET",
		PathPattern:        "/v2/Bus/RealTimeNearStop/NearBy",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BusAPIRealTimeNearStopNearBy2854Reader{formats: a.formats},
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
	case *BusAPIRealTimeNearStopNearBy2854OK:
		return value, nil, nil
	case *BusAPIRealTimeNearStopNearBy2854Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for bus_advanced_near_by: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BusAPIRouteNearBy2852 取得指定s 位置 範圍 的全臺公車路線資料

  取得指定[位置,範圍]的全臺公車路線資料
*/
func (a *Client) BusAPIRouteNearBy2852(params *BusAPIRouteNearBy2852Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BusAPIRouteNearBy2852OK, *BusAPIRouteNearBy2852Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBusAPIRouteNearBy2852Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "BusApi_Route_NearBy_2852",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Route/NearBy",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BusAPIRouteNearBy2852Reader{formats: a.formats},
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
	case *BusAPIRouteNearBy2852OK:
		return value, nil, nil
	case *BusAPIRouteNearBy2852Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for bus_advanced_near_by: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BusAPIStationNearBy2851 取得指定s 位置 範圍 的全臺公車站位資料

  取得指定[位置,範圍]的全臺公車站位資料
*/
func (a *Client) BusAPIStationNearBy2851(params *BusAPIStationNearBy2851Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BusAPIStationNearBy2851OK, *BusAPIStationNearBy2851Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBusAPIStationNearBy2851Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "BusApi_Station_NearBy_2851",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Station/NearBy",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BusAPIStationNearBy2851Reader{formats: a.formats},
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
	case *BusAPIStationNearBy2851OK:
		return value, nil, nil
	case *BusAPIStationNearBy2851Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for bus_advanced_near_by: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BusAPIStopNearBy2850 取得指定s 位置 範圍 的全臺公車站牌資料

  取得指定[位置,範圍]的全臺公車站牌資料
*/
func (a *Client) BusAPIStopNearBy2850(params *BusAPIStopNearBy2850Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BusAPIStopNearBy2850OK, *BusAPIStopNearBy2850Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBusAPIStopNearBy2850Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "BusApi_Stop_NearBy_2850",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Stop/NearBy",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BusAPIStopNearBy2850Reader{formats: a.formats},
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
	case *BusAPIStopNearBy2850OK:
		return value, nil, nil
	case *BusAPIStopNearBy2850Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for bus_advanced_near_by: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
