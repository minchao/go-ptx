// Code generated by go-swagger; DO NOT EDIT.

package ship_by_route_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ship by route type API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ship by route type API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ShipRouteTypeDailySchedule23260(params *ShipRouteTypeDailySchedule23260Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShipRouteTypeDailySchedule23260OK, *ShipRouteTypeDailySchedule23260Status299, error)

	ShipRouteTypeDailySchedule3259(params *ShipRouteTypeDailySchedule3259Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShipRouteTypeDailySchedule3259OK, *ShipRouteTypeDailySchedule3259Status299, error)

	ShipRouteTypeRoute3257(params *ShipRouteTypeRoute3257Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShipRouteTypeRoute3257OK, *ShipRouteTypeRoute3257Status299, error)

	ShipRouteTypeSpecificSchedule23262(params *ShipRouteTypeSpecificSchedule23262Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShipRouteTypeSpecificSchedule23262OK, *ShipRouteTypeSpecificSchedule23262Status299, error)

	ShipRouteTypeSpecificSchedule3261(params *ShipRouteTypeSpecificSchedule3261Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShipRouteTypeSpecificSchedule3261OK, *ShipRouteTypeSpecificSchedule3261Status299, error)

	ShipRouteTypeStopOfRoute3258(params *ShipRouteTypeStopOfRoute3258Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShipRouteTypeStopOfRoute3258OK, *ShipRouteTypeStopOfRoute3258Status299, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ShipRouteTypeDailySchedule23260 取得指定s 航線 的全臺航運之每日班表資料
*/
func (a *Client) ShipRouteTypeDailySchedule23260(params *ShipRouteTypeDailySchedule23260Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShipRouteTypeDailySchedule23260OK, *ShipRouteTypeDailySchedule23260Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShipRouteTypeDailySchedule23260Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "Ship_RouteType_DailySchedule2_3260",
		Method:             "GET",
		PathPattern:        "/v3/Ship/DailySchedule/RouteID/{RouteID}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ShipRouteTypeDailySchedule23260Reader{formats: a.formats},
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
	case *ShipRouteTypeDailySchedule23260OK:
		return value, nil, nil
	case *ShipRouteTypeDailySchedule23260Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ship_by_route_type: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ShipRouteTypeDailySchedule3259 取得指定s 航線類別 的全臺航運之每日班表資料
*/
func (a *Client) ShipRouteTypeDailySchedule3259(params *ShipRouteTypeDailySchedule3259Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShipRouteTypeDailySchedule3259OK, *ShipRouteTypeDailySchedule3259Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShipRouteTypeDailySchedule3259Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "Ship_RouteType_DailySchedule_3259",
		Method:             "GET",
		PathPattern:        "/v3/Ship/DailySchedule/RouteType/{RouteType}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ShipRouteTypeDailySchedule3259Reader{formats: a.formats},
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
	case *ShipRouteTypeDailySchedule3259OK:
		return value, nil, nil
	case *ShipRouteTypeDailySchedule3259Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ship_by_route_type: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ShipRouteTypeRoute3257 取得指定s 航線類別 的全臺航運之航線資料
*/
func (a *Client) ShipRouteTypeRoute3257(params *ShipRouteTypeRoute3257Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShipRouteTypeRoute3257OK, *ShipRouteTypeRoute3257Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShipRouteTypeRoute3257Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "Ship_RouteType_Route_3257",
		Method:             "GET",
		PathPattern:        "/v3/Ship/Route/RouteType/{RouteType}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ShipRouteTypeRoute3257Reader{formats: a.formats},
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
	case *ShipRouteTypeRoute3257OK:
		return value, nil, nil
	case *ShipRouteTypeRoute3257Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ship_by_route_type: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ShipRouteTypeSpecificSchedule23262 取得指定s 航線 的全臺航運之特殊班表資料
*/
func (a *Client) ShipRouteTypeSpecificSchedule23262(params *ShipRouteTypeSpecificSchedule23262Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShipRouteTypeSpecificSchedule23262OK, *ShipRouteTypeSpecificSchedule23262Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShipRouteTypeSpecificSchedule23262Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "Ship_RouteType_SpecificSchedule2_3262",
		Method:             "GET",
		PathPattern:        "/v3/Ship/SpecificSchedule/RouteID/{RouteID}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ShipRouteTypeSpecificSchedule23262Reader{formats: a.formats},
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
	case *ShipRouteTypeSpecificSchedule23262OK:
		return value, nil, nil
	case *ShipRouteTypeSpecificSchedule23262Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ship_by_route_type: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ShipRouteTypeSpecificSchedule3261 取得指定s 航線類別 的全臺航運之特殊班表資料
*/
func (a *Client) ShipRouteTypeSpecificSchedule3261(params *ShipRouteTypeSpecificSchedule3261Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShipRouteTypeSpecificSchedule3261OK, *ShipRouteTypeSpecificSchedule3261Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShipRouteTypeSpecificSchedule3261Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "Ship_RouteType_SpecificSchedule_3261",
		Method:             "GET",
		PathPattern:        "/v3/Ship/SpecificSchedule/RouteType/{RouteType}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ShipRouteTypeSpecificSchedule3261Reader{formats: a.formats},
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
	case *ShipRouteTypeSpecificSchedule3261OK:
		return value, nil, nil
	case *ShipRouteTypeSpecificSchedule3261Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ship_by_route_type: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ShipRouteTypeStopOfRoute3258 取得指定s 航線類別 之航線靠港順序資料
*/
func (a *Client) ShipRouteTypeStopOfRoute3258(params *ShipRouteTypeStopOfRoute3258Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShipRouteTypeStopOfRoute3258OK, *ShipRouteTypeStopOfRoute3258Status299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShipRouteTypeStopOfRoute3258Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "Ship_RouteType_StopOfRoute_3258",
		Method:             "GET",
		PathPattern:        "/v3/Ship/StopOfRoute/RouteType/{RouteType}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ShipRouteTypeStopOfRoute3258Reader{formats: a.formats},
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
	case *ShipRouteTypeStopOfRoute3258OK:
		return value, nil, nil
	case *ShipRouteTypeStopOfRoute3258Status299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ship_by_route_type: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
