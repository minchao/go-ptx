// Code generated by go-swagger; DO NOT EDIT.

package bike

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new bike API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bike API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	BikeAPIAvailability(params *BikeAPIAvailabilityParams) (*BikeAPIAvailabilityOK, error)

	BikeAPIStation(params *BikeAPIStationParams) (*BikeAPIStationOK, error)

	CyclingAPIShape(params *CyclingAPIShapeParams) (*CyclingAPIShapeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BikeAPIAvailability 取得動態指定s 縣市 的公共自行車即時車位資料

  取得動態指定[縣市]的公共自行車即時車位資料<br />[更新頻率]2分鐘更新一次
*/
func (a *Client) BikeAPIAvailability(params *BikeAPIAvailabilityParams) (*BikeAPIAvailabilityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBikeAPIAvailabilityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BikeApi_Availability",
		Method:             "GET",
		PathPattern:        "/v2/Bike/Availability/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BikeAPIAvailabilityReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BikeAPIAvailabilityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BikeApi_Availability: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BikeAPIStation 取得指定s 縣市 的公共自行車租借站位資料

  取得指定[縣市]的公共自行車租借站位資料<br />[更新頻率]2分鐘更新一次
*/
func (a *Client) BikeAPIStation(params *BikeAPIStationParams) (*BikeAPIStationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBikeAPIStationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BikeApi_Station",
		Method:             "GET",
		PathPattern:        "/v2/Bike/Station/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BikeAPIStationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BikeAPIStationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BikeApi_Station: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CyclingAPIShape 取得指定縣市之自行車道路網圖資s

  取得指定縣市之自行車道路網圖資
*/
func (a *Client) CyclingAPIShape(params *CyclingAPIShapeParams) (*CyclingAPIShapeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCyclingAPIShapeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CyclingApi_Shape",
		Method:             "GET",
		PathPattern:        "/v2/Cycling/Shape/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CyclingAPIShapeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CyclingAPIShapeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CyclingApi_Shape: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
