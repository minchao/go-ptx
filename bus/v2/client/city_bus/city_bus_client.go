// Code generated by go-swagger; DO NOT EDIT.

package city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new city bus API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for city bus API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CityBusAPIDataVersion 取得指定s 縣市 目前資料的最新版本資訊

版本詳細資訊
*/
func (a *Client) CityBusAPIDataVersion(params *CityBusAPIDataVersionParams) (*CityBusAPIDataVersionOK, *CityBusAPIDataVersionStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIDataVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_DataVersion",
		Method:             "GET",
		PathPattern:        "/v2/Bus/DataVersion/City/{City}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIDataVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIDataVersionOK:
		return value, nil, nil
	case *CityBusAPIDataVersionStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIDisplayStopOfRoute 取得指定s 縣市 的市區公車顯示用路線站序資料

市區公車之顯示用路線站序資料，僅台北市與新北市可查詢
*/
func (a *Client) CityBusAPIDisplayStopOfRoute(params *CityBusAPIDisplayStopOfRouteParams) (*CityBusAPIDisplayStopOfRouteOK, *CityBusAPIDisplayStopOfRouteStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIDisplayStopOfRouteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_DisplayStopOfRoute",
		Method:             "GET",
		PathPattern:        "/v2/Bus/DisplayStopOfRoute/City/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIDisplayStopOfRouteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIDisplayStopOfRouteOK:
		return value, nil, nil
	case *CityBusAPIDisplayStopOfRouteStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIDisplayStopOfRoute1 取得指定s 縣市 路線名稱 的市區公車顯示用路線站序資料

市區公車之顯示用路線站序資料，僅台北市與新北市可查詢
*/
func (a *Client) CityBusAPIDisplayStopOfRoute1(params *CityBusAPIDisplayStopOfRoute1Params) (*CityBusAPIDisplayStopOfRoute1OK, *CityBusAPIDisplayStopOfRoute1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIDisplayStopOfRoute1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_DisplayStopOfRoute_1",
		Method:             "GET",
		PathPattern:        "/v2/Bus/DisplayStopOfRoute/City/{City}/{RouteName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIDisplayStopOfRoute1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIDisplayStopOfRoute1OK:
		return value, nil, nil
	case *CityBusAPIDisplayStopOfRoute1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIEstimatedTimeOfArrival 取得指定s 縣市 的公車預估到站資料 n1 批次更新

 ### 市區公車之預估到站資料(N1) ###
[部分縣市] 當 StopStatus = 1(尚未發車) 且 EstimateTime &gt; 0 (有值) 的情形, 屬正常情形, 雖目前尚未發車, 但提供EstimateTime值為預計多久後開始發車之時間。
*/
func (a *Client) CityBusAPIEstimatedTimeOfArrival(params *CityBusAPIEstimatedTimeOfArrivalParams) (*CityBusAPIEstimatedTimeOfArrivalOK, *CityBusAPIEstimatedTimeOfArrivalStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIEstimatedTimeOfArrivalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_EstimatedTimeOfArrival",
		Method:             "GET",
		PathPattern:        "/v2/Bus/EstimatedTimeOfArrival/City/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIEstimatedTimeOfArrivalReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIEstimatedTimeOfArrivalOK:
		return value, nil, nil
	case *CityBusAPIEstimatedTimeOfArrivalStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIEstimatedTimeOfArrival1 取得指定s 縣市 路線名稱 的公車預估到站資料 n1 批次更新

 ### 市區公車之預估到站資料(N1) ###
[部分縣市] 當 StopStatus = 1(尚未發車) 且 EstimateTime &gt; 0 (有值) 的情形, 屬正常情形, 雖目前尚未發車, 但提供EstimateTime值為預計多久後開始發車之時間。
*/
func (a *Client) CityBusAPIEstimatedTimeOfArrival1(params *CityBusAPIEstimatedTimeOfArrival1Params) (*CityBusAPIEstimatedTimeOfArrival1OK, *CityBusAPIEstimatedTimeOfArrival1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIEstimatedTimeOfArrival1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_EstimatedTimeOfArrival_1",
		Method:             "GET",
		PathPattern:        "/v2/Bus/EstimatedTimeOfArrival/City/{City}/{RouteName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIEstimatedTimeOfArrival1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIEstimatedTimeOfArrival1OK:
		return value, nil, nil
	case *CityBusAPIEstimatedTimeOfArrival1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPINews 取得指定s 縣市 的市區公車最新消息資料

市區公車最新消息資料
*/
func (a *Client) CityBusAPINews(params *CityBusAPINewsParams) (*CityBusAPINewsOK, *CityBusAPINewsStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPINewsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_News",
		Method:             "GET",
		PathPattern:        "/v2/Bus/News/City/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPINewsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPINewsOK:
		return value, nil, nil
	case *CityBusAPINewsStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIOperator 取得指定s 縣市 的市區公車營運業者資料

市區公車之營運業者資料
*/
func (a *Client) CityBusAPIOperator(params *CityBusAPIOperatorParams) (*CityBusAPIOperatorOK, *CityBusAPIOperatorStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIOperatorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_Operator",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Operator/City/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIOperatorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIOperatorOK:
		return value, nil, nil
	case *CityBusAPIOperatorStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIRealTimeByFrequency 取得指定s 縣市 的公車動態定時資料 a1 批次更新

### 市區公車之定時資料(A1) ###
*/
func (a *Client) CityBusAPIRealTimeByFrequency(params *CityBusAPIRealTimeByFrequencyParams) (*CityBusAPIRealTimeByFrequencyOK, *CityBusAPIRealTimeByFrequencyStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIRealTimeByFrequencyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_RealTimeByFrequency",
		Method:             "GET",
		PathPattern:        "/v2/Bus/RealTimeByFrequency/City/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIRealTimeByFrequencyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIRealTimeByFrequencyOK:
		return value, nil, nil
	case *CityBusAPIRealTimeByFrequencyStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIRealTimeByFrequency1 取得指定s 縣市 路線名稱 的公車動態定時資料 a1 批次更新

### 市區公車之定時資料(A1) ###
*/
func (a *Client) CityBusAPIRealTimeByFrequency1(params *CityBusAPIRealTimeByFrequency1Params) (*CityBusAPIRealTimeByFrequency1OK, *CityBusAPIRealTimeByFrequency1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIRealTimeByFrequency1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_RealTimeByFrequency_1",
		Method:             "GET",
		PathPattern:        "/v2/Bus/RealTimeByFrequency/City/{City}/{RouteName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIRealTimeByFrequency1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIRealTimeByFrequency1OK:
		return value, nil, nil
	case *CityBusAPIRealTimeByFrequency1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIRealTimeNearStop 取得指定s 縣市 的公車動態定點資料 a2 批次更新

### 市區公車之定點資料(A2) ###
<param name="City">欲查詢縣市</param><returns>公車動態定點資料(A2)</returns>
- [逐筆更新]與[批次更新]之差異請詳見資料使用葵花寶典([連結](https://ptxmotc.gitbooks.io/ptx-api-documentation/content/api-zi-liao-shi-yong-zhu-yi-shi-xiang/bus.html))
*/
func (a *Client) CityBusAPIRealTimeNearStop(params *CityBusAPIRealTimeNearStopParams) (*CityBusAPIRealTimeNearStopOK, *CityBusAPIRealTimeNearStopStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIRealTimeNearStopParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_RealTimeNearStop",
		Method:             "GET",
		PathPattern:        "/v2/Bus/RealTimeNearStop/City/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIRealTimeNearStopReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIRealTimeNearStopOK:
		return value, nil, nil
	case *CityBusAPIRealTimeNearStopStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIRealTimeNearStop1 取得指定s 縣市 路線名稱 的公車動態定點資料 a2 批次更新

### 市區公車之定點資料(A2) ###
*/
func (a *Client) CityBusAPIRealTimeNearStop1(params *CityBusAPIRealTimeNearStop1Params) (*CityBusAPIRealTimeNearStop1OK, *CityBusAPIRealTimeNearStop1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIRealTimeNearStop1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_RealTimeNearStop_1",
		Method:             "GET",
		PathPattern:        "/v2/Bus/RealTimeNearStop/City/{City}/{RouteName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIRealTimeNearStop1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIRealTimeNearStop1OK:
		return value, nil, nil
	case *CityBusAPIRealTimeNearStop1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIRoute 取得指定s 縣市 的市區公車路線資料

市區公車之路線資料
*/
func (a *Client) CityBusAPIRoute(params *CityBusAPIRouteParams) (*CityBusAPIRouteOK, *CityBusAPIRouteStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIRouteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_Route",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Route/City/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIRouteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIRouteOK:
		return value, nil, nil
	case *CityBusAPIRouteStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIRouteFare 取得指定s 縣市 的市區公車路線票價資料

市區公車路線票價資料
*/
func (a *Client) CityBusAPIRouteFare(params *CityBusAPIRouteFareParams) (*CityBusAPIRouteFareOK, *CityBusAPIRouteFareStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIRouteFareParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_RouteFare",
		Method:             "GET",
		PathPattern:        "/v2/Bus/RouteFare/City/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIRouteFareReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIRouteFareOK:
		return value, nil, nil
	case *CityBusAPIRouteFareStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIRouteFare1 取得指定s 縣市 路線名稱 的的市區公車路線票價資料

市區公車路線票價資料
*/
func (a *Client) CityBusAPIRouteFare1(params *CityBusAPIRouteFare1Params) (*CityBusAPIRouteFare1OK, *CityBusAPIRouteFare1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIRouteFare1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_RouteFare_1",
		Method:             "GET",
		PathPattern:        "/v2/Bus/RouteFare/City/{City}/{RouteName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIRouteFare1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIRouteFare1OK:
		return value, nil, nil
	case *CityBusAPIRouteFare1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIRoute1 取得指定s 縣市 路線名稱 的路線資料

市區公車之路線資料
*/
func (a *Client) CityBusAPIRoute1(params *CityBusAPIRoute1Params) (*CityBusAPIRoute1OK, *CityBusAPIRoute1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIRoute1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_Route_1",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Route/City/{City}/{RouteName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIRoute1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIRoute1OK:
		return value, nil, nil
	case *CityBusAPIRoute1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPISchedule 取得指定s 縣市 的市區公車路線班表資料

市區公車之班表及班距資料。一般市區公車班次較多時會採用【班距】式時刻表；班次較少時會採用【班表】式時刻表
*/
func (a *Client) CityBusAPISchedule(params *CityBusAPIScheduleParams) (*CityBusAPIScheduleOK, *CityBusAPIScheduleStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIScheduleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_Schedule",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Schedule/City/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIScheduleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIScheduleOK:
		return value, nil, nil
	case *CityBusAPIScheduleStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPISchedule1 取得指定s 縣市 路線名稱 的市區公車路線班表資料

市區公車之預定班表及班距資料。一般市區公車班次較多時會採用【班距】式時刻表；班次較少時會採用【班表】式時刻表
*/
func (a *Client) CityBusAPISchedule1(params *CityBusAPISchedule1Params) (*CityBusAPISchedule1OK, *CityBusAPISchedule1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPISchedule1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_Schedule_1",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Schedule/City/{City}/{RouteName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPISchedule1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPISchedule1OK:
		return value, nil, nil
	case *CityBusAPISchedule1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIShape 取得指定s 縣市 的市區公車線型資料

市區公車之線型資料
*/
func (a *Client) CityBusAPIShape(params *CityBusAPIShapeParams) (*CityBusAPIShapeOK, *CityBusAPIShapeStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIShapeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_Shape",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Shape/City/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIShapeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIShapeOK:
		return value, nil, nil
	case *CityBusAPIShapeStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIShape1 取得指定s 縣市 路線名稱 的市區公車線型資料

市區公車之線型資料
*/
func (a *Client) CityBusAPIShape1(params *CityBusAPIShape1Params) (*CityBusAPIShape1OK, *CityBusAPIShape1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIShape1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_Shape_1",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Shape/City/{City}/{RouteName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIShape1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIShape1OK:
		return value, nil, nil
	case *CityBusAPIShape1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIStation 取得指定s 縣市 的市區公車站位資料

市區公車之各站牌所屬的站位資料
*/
func (a *Client) CityBusAPIStation(params *CityBusAPIStationParams) (*CityBusAPIStationOK, *CityBusAPIStationStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIStationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_Station",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Station/City/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIStationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIStationOK:
		return value, nil, nil
	case *CityBusAPIStationStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIStationName 取得指定s 縣市 的市區公車站名碼資料

市區公車之各站牌所屬的站名碼資料
*/
func (a *Client) CityBusAPIStationName(params *CityBusAPIStationNameParams) (*CityBusAPIStationNameOK, *CityBusAPIStationNameStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIStationNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_StationName",
		Method:             "GET",
		PathPattern:        "/v2/Bus/StationName/City/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIStationNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIStationNameOK:
		return value, nil, nil
	case *CityBusAPIStationNameStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIStop 取得指定s 縣市 的市區公車站牌資料

市區公車之站牌資料
*/
func (a *Client) CityBusAPIStop(params *CityBusAPIStopParams) (*CityBusAPIStopOK, *CityBusAPIStopStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIStopParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_Stop",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Stop/City/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIStopReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIStopOK:
		return value, nil, nil
	case *CityBusAPIStopStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIStopOfRoute 取得指定s 縣市 的市區公車路線站序資料

市區公車之路線站序資料
*/
func (a *Client) CityBusAPIStopOfRoute(params *CityBusAPIStopOfRouteParams) (*CityBusAPIStopOfRouteOK, *CityBusAPIStopOfRouteStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIStopOfRouteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_StopOfRoute",
		Method:             "GET",
		PathPattern:        "/v2/Bus/StopOfRoute/City/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIStopOfRouteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIStopOfRouteOK:
		return value, nil, nil
	case *CityBusAPIStopOfRouteStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIStopOfRoute1 取得指定s 縣市 路線名稱 的市區公車路線站序資料

市區公車之路線站序資料
*/
func (a *Client) CityBusAPIStopOfRoute1(params *CityBusAPIStopOfRoute1Params) (*CityBusAPIStopOfRoute1OK, *CityBusAPIStopOfRoute1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIStopOfRoute1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_StopOfRoute_1",
		Method:             "GET",
		PathPattern:        "/v2/Bus/StopOfRoute/City/{City}/{RouteName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIStopOfRoute1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIStopOfRoute1OK:
		return value, nil, nil
	case *CityBusAPIStopOfRoute1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CityBusAPIVehicle 取得指定s 縣市 的市區公車車輛資料

市區公車之車輛資料
*/
func (a *Client) CityBusAPIVehicle(params *CityBusAPIVehicleParams) (*CityBusAPIVehicleOK, *CityBusAPIVehicleStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCityBusAPIVehicleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CityBusApi_Vehicle",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Vehicle/City/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CityBusAPIVehicleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CityBusAPIVehicleOK:
		return value, nil, nil
	case *CityBusAPIVehicleStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
