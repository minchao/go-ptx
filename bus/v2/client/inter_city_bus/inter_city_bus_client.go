// Code generated by go-swagger; DO NOT EDIT.

package inter_city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new inter city bus API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for inter city bus API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
InterCityBusAPIDataVersion 取得公路客運目前資料的最新版本資訊s

版本詳細資訊
*/
func (a *Client) InterCityBusAPIDataVersion(params *InterCityBusAPIDataVersionParams) (*InterCityBusAPIDataVersionOK, *InterCityBusAPIDataVersionStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIDataVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_DataVersion",
		Method:             "GET",
		PathPattern:        "/v2/Bus/DataVersion/InterCity",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIDataVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIDataVersionOK:
		return value, nil, nil
	case *InterCityBusAPIDataVersionStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPIEstimatedTimeOfArrival 取得公路客運的預估到站資料s n1 批次更新

###公路客運之預估到站資料(N1)###
不保留[現在時間]超過[本平台資料更新時間]兩分鐘的資料
*/
func (a *Client) InterCityBusAPIEstimatedTimeOfArrival(params *InterCityBusAPIEstimatedTimeOfArrivalParams) (*InterCityBusAPIEstimatedTimeOfArrivalOK, *InterCityBusAPIEstimatedTimeOfArrivalStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIEstimatedTimeOfArrivalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_EstimatedTimeOfArrival",
		Method:             "GET",
		PathPattern:        "/v2/Bus/EstimatedTimeOfArrival/InterCity",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIEstimatedTimeOfArrivalReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIEstimatedTimeOfArrivalOK:
		return value, nil, nil
	case *InterCityBusAPIEstimatedTimeOfArrivalStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPIEstimatedTimeOfArrival1 取得指定s 路線名稱 的公路客運預估到站資料 n1 批次更新

### 公路客運之預估到站資料(N1) ###
不保留[現在時間]超過[本平台資料更新時間]兩分鐘的資料
*/
func (a *Client) InterCityBusAPIEstimatedTimeOfArrival1(params *InterCityBusAPIEstimatedTimeOfArrival1Params) (*InterCityBusAPIEstimatedTimeOfArrival1OK, *InterCityBusAPIEstimatedTimeOfArrival1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIEstimatedTimeOfArrival1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_EstimatedTimeOfArrival_1",
		Method:             "GET",
		PathPattern:        "/v2/Bus/EstimatedTimeOfArrival/InterCity/{RouteName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIEstimatedTimeOfArrival1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIEstimatedTimeOfArrival1OK:
		return value, nil, nil
	case *InterCityBusAPIEstimatedTimeOfArrival1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPINews 取得公路客運之最新消息s

公路客運之最新消息
*/
func (a *Client) InterCityBusAPINews(params *InterCityBusAPINewsParams) (*InterCityBusAPINewsOK, *InterCityBusAPINewsStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPINewsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_News",
		Method:             "GET",
		PathPattern:        "/v2/Bus/News/InterCity",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPINewsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPINewsOK:
		return value, nil, nil
	case *InterCityBusAPINewsStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPIOperator 取得公路客運的營運業者資料s

公路客運之營運業者資料
*/
func (a *Client) InterCityBusAPIOperator(params *InterCityBusAPIOperatorParams) (*InterCityBusAPIOperatorOK, *InterCityBusAPIOperatorStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIOperatorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_Operator",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Operator/InterCity",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIOperatorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIOperatorOK:
		return value, nil, nil
	case *InterCityBusAPIOperatorStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPIRealTimeByFrequency 取得公路客運的動態定時資料s a1 批次更新

### 公路客運之定時資料(A1) ###
*/
func (a *Client) InterCityBusAPIRealTimeByFrequency(params *InterCityBusAPIRealTimeByFrequencyParams) (*InterCityBusAPIRealTimeByFrequencyOK, *InterCityBusAPIRealTimeByFrequencyStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIRealTimeByFrequencyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_RealTimeByFrequency",
		Method:             "GET",
		PathPattern:        "/v2/Bus/RealTimeByFrequency/InterCity",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIRealTimeByFrequencyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIRealTimeByFrequencyOK:
		return value, nil, nil
	case *InterCityBusAPIRealTimeByFrequencyStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPIRealTimeByFrequency1 取得指定s 路線名稱 的公路客運動態定時資料 a1 批次更新

### 公路客運之定時資料(A1) ###
*/
func (a *Client) InterCityBusAPIRealTimeByFrequency1(params *InterCityBusAPIRealTimeByFrequency1Params) (*InterCityBusAPIRealTimeByFrequency1OK, *InterCityBusAPIRealTimeByFrequency1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIRealTimeByFrequency1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_RealTimeByFrequency_1",
		Method:             "GET",
		PathPattern:        "/v2/Bus/RealTimeByFrequency/InterCity/{RouteName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIRealTimeByFrequency1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIRealTimeByFrequency1OK:
		return value, nil, nil
	case *InterCityBusAPIRealTimeByFrequency1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPIRealTimeNearStop 取得公路客運的動態定點資料s a2 批次更新

### 公路客運之定點資料(A2) ###
<returns>公路客運動態定點資料(A2)</returns>
- [逐筆更新]與[批次更新]之差異請詳見資料使用葵花寶典([連結](https://ptxmotc.gitbooks.io/ptx-api-documentation/content/api-zi-liao-shi-yong-zhu-yi-shi-xiang/bus.html))
*/
func (a *Client) InterCityBusAPIRealTimeNearStop(params *InterCityBusAPIRealTimeNearStopParams) (*InterCityBusAPIRealTimeNearStopOK, *InterCityBusAPIRealTimeNearStopStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIRealTimeNearStopParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_RealTimeNearStop",
		Method:             "GET",
		PathPattern:        "/v2/Bus/RealTimeNearStop/InterCity",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIRealTimeNearStopReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIRealTimeNearStopOK:
		return value, nil, nil
	case *InterCityBusAPIRealTimeNearStopStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPIRealTimeNearStop1 取得指定s 路線名稱 的公路客運動態定點資料 a2 批次更新

### 公路客運之定點資料(A2) ###
*/
func (a *Client) InterCityBusAPIRealTimeNearStop1(params *InterCityBusAPIRealTimeNearStop1Params) (*InterCityBusAPIRealTimeNearStop1OK, *InterCityBusAPIRealTimeNearStop1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIRealTimeNearStop1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_RealTimeNearStop_1",
		Method:             "GET",
		PathPattern:        "/v2/Bus/RealTimeNearStop/InterCity/{RouteName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIRealTimeNearStop1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIRealTimeNearStop1OK:
		return value, nil, nil
	case *InterCityBusAPIRealTimeNearStop1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPIRoute 取得公路客運路線資料s

公路客運之路線資料
*/
func (a *Client) InterCityBusAPIRoute(params *InterCityBusAPIRouteParams) (*InterCityBusAPIRouteOK, *InterCityBusAPIRouteStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIRouteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_Route",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Route/InterCity",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIRouteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIRouteOK:
		return value, nil, nil
	case *InterCityBusAPIRouteStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPIRouteFare 取得公路客運之路線票價資料s

公路客運之路線票價資料
*/
func (a *Client) InterCityBusAPIRouteFare(params *InterCityBusAPIRouteFareParams) (*InterCityBusAPIRouteFareOK, *InterCityBusAPIRouteFareStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIRouteFareParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_RouteFare",
		Method:             "GET",
		PathPattern:        "/v2/Bus/RouteFare/InterCity",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIRouteFareReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIRouteFareOK:
		return value, nil, nil
	case *InterCityBusAPIRouteFareStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPIRouteFare1 取得指定s 路線名稱 的公路客運路線資料

公路客運之路線資料
*/
func (a *Client) InterCityBusAPIRouteFare1(params *InterCityBusAPIRouteFare1Params) (*InterCityBusAPIRouteFare1OK, *InterCityBusAPIRouteFare1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIRouteFare1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_RouteFare_1",
		Method:             "GET",
		PathPattern:        "/v2/Bus/RouteFare/InterCity/{RouteName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIRouteFare1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIRouteFare1OK:
		return value, nil, nil
	case *InterCityBusAPIRouteFare1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPIRoute1 取得指定s 路線名稱 的公路客運路線資料

公路客運之路線資料
*/
func (a *Client) InterCityBusAPIRoute1(params *InterCityBusAPIRoute1Params) (*InterCityBusAPIRoute1OK, *InterCityBusAPIRoute1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIRoute1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_Route_1",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Route/InterCity/{RouteName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIRoute1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIRoute1OK:
		return value, nil, nil
	case *InterCityBusAPIRoute1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPISchedule 取得公路客運路線班表資料s

公路客運之預定班表及班距資料。公路及國道客運多採用【班表】式時刻表
*/
func (a *Client) InterCityBusAPISchedule(params *InterCityBusAPIScheduleParams) (*InterCityBusAPIScheduleOK, *InterCityBusAPIScheduleStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIScheduleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_Schedule",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Schedule/InterCity",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIScheduleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIScheduleOK:
		return value, nil, nil
	case *InterCityBusAPIScheduleStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPISchedule1 取得指定s 路線名稱 的公路客運路線班表資料

公路客運之預定班表及班距資料。公路及國道客運多採用【班表】式時刻表
*/
func (a *Client) InterCityBusAPISchedule1(params *InterCityBusAPISchedule1Params) (*InterCityBusAPISchedule1OK, *InterCityBusAPISchedule1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPISchedule1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_Schedule_1",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Schedule/InterCity/{RouteName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPISchedule1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPISchedule1OK:
		return value, nil, nil
	case *InterCityBusAPISchedule1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPIStation 取得公路客運站位資料s

公路客運之各站牌所屬的站位資料
*/
func (a *Client) InterCityBusAPIStation(params *InterCityBusAPIStationParams) (*InterCityBusAPIStationOK, *InterCityBusAPIStationStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIStationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_Station",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Station/InterCity",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIStationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIStationOK:
		return value, nil, nil
	case *InterCityBusAPIStationStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPIStationName 取得公路客運站名碼資料s

公路客運之各站牌所屬的站名碼資料
*/
func (a *Client) InterCityBusAPIStationName(params *InterCityBusAPIStationNameParams) (*InterCityBusAPIStationNameOK, *InterCityBusAPIStationNameStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIStationNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_StationName",
		Method:             "GET",
		PathPattern:        "/v2/Bus/StationName/InterCity",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIStationNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIStationNameOK:
		return value, nil, nil
	case *InterCityBusAPIStationNameStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPIStop 取得公路客運站牌資料s

公路客運之站牌資料
*/
func (a *Client) InterCityBusAPIStop(params *InterCityBusAPIStopParams) (*InterCityBusAPIStopOK, *InterCityBusAPIStopStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIStopParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_Stop",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Stop/InterCity",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIStopReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIStopOK:
		return value, nil, nil
	case *InterCityBusAPIStopStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPIStopOfRoute 取得公路客運路線與站牌資料s

公路客運之路線與站牌資料
*/
func (a *Client) InterCityBusAPIStopOfRoute(params *InterCityBusAPIStopOfRouteParams) (*InterCityBusAPIStopOfRouteOK, *InterCityBusAPIStopOfRouteStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIStopOfRouteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_StopOfRoute",
		Method:             "GET",
		PathPattern:        "/v2/Bus/StopOfRoute/InterCity",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIStopOfRouteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIStopOfRouteOK:
		return value, nil, nil
	case *InterCityBusAPIStopOfRouteStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPIStopOfRoute1 取得指定s 路線名稱 的公路客運路線與站牌資料

公路客運之路線與站牌資料
*/
func (a *Client) InterCityBusAPIStopOfRoute1(params *InterCityBusAPIStopOfRoute1Params) (*InterCityBusAPIStopOfRoute1OK, *InterCityBusAPIStopOfRoute1Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIStopOfRoute1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_StopOfRoute_1",
		Method:             "GET",
		PathPattern:        "/v2/Bus/StopOfRoute/InterCity/{RouteName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIStopOfRoute1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIStopOfRoute1OK:
		return value, nil, nil
	case *InterCityBusAPIStopOfRoute1Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InterCityBusAPIVehicle 取得公路客運之車輛資料s

公路客運之車輛資料
*/
func (a *Client) InterCityBusAPIVehicle(params *InterCityBusAPIVehicleParams) (*InterCityBusAPIVehicleOK, *InterCityBusAPIVehicleStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInterCityBusAPIVehicleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InterCityBusApi_Vehicle",
		Method:             "GET",
		PathPattern:        "/v2/Bus/Vehicle/InterCity",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InterCityBusAPIVehicleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *InterCityBusAPIVehicleOK:
		return value, nil, nil
	case *InterCityBusAPIVehicleStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inter_city_bus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
