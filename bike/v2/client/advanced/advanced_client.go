// Code generated by go-swagger; DO NOT EDIT.

package advanced

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new advanced API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for advanced API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BikeAPIAvailabilityAllCityNearBy(params *BikeAPIAvailabilityAllCityNearByParams, opts ...ClientOption) (*BikeAPIAvailabilityAllCityNearByOK, error)

	BikeAPIStationAllCityNearBy(params *BikeAPIStationAllCityNearByParams, opts ...ClientOption) (*BikeAPIStationAllCityNearByOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BikeAPIAvailabilityAllCityNearBy 取得指定s 位置 範圍 的全臺公共自行車即時車位資料

  取得指定[位置,範圍]的全臺公共自行車即時車位資料
*/
func (a *Client) BikeAPIAvailabilityAllCityNearBy(params *BikeAPIAvailabilityAllCityNearByParams, opts ...ClientOption) (*BikeAPIAvailabilityAllCityNearByOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBikeAPIAvailabilityAllCityNearByParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BikeApi_Availability_AllCity_NearBy",
		Method:             "GET",
		PathPattern:        "/v2/Bike/Availability/NearBy",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BikeAPIAvailabilityAllCityNearByReader{formats: a.formats},
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
	success, ok := result.(*BikeAPIAvailabilityAllCityNearByOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BikeApi_Availability_AllCity_NearBy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BikeAPIStationAllCityNearBy 取得指定s 位置 範圍 的全臺公共自行車租借站位資料

  取得指定[位置,範圍]的全臺公共自行車租借站位資料
*/
func (a *Client) BikeAPIStationAllCityNearBy(params *BikeAPIStationAllCityNearByParams, opts ...ClientOption) (*BikeAPIStationAllCityNearByOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBikeAPIStationAllCityNearByParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BikeApi_Station_AllCity_NearBy",
		Method:             "GET",
		PathPattern:        "/v2/Bike/Station/NearBy",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BikeAPIStationAllCityNearByReader{formats: a.formats},
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
	success, ok := result.(*BikeAPIStationAllCityNearByOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BikeApi_Station_AllCity_NearBy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
