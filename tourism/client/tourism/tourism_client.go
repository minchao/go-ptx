// Code generated by go-swagger; DO NOT EDIT.

package tourism

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new tourism API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tourism API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
TourismAPIActivity 取得所有觀光活動資料s

取得所有觀光活動資料
*/
func (a *Client) TourismAPIActivity(params *TourismAPIActivityParams) (*TourismAPIActivityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTourismAPIActivityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TourismApi_Activity",
		Method:             "GET",
		PathPattern:        "/v2/Tourism/Activity",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TourismAPIActivityReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TourismAPIActivityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TourismApi_Activity: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TourismAPIActivity0 取得指定s 縣市 觀光活動資料

取得指定[縣市]觀光活動資料
*/
func (a *Client) TourismAPIActivity0(params *TourismAPIActivity0Params) (*TourismAPIActivity0OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTourismAPIActivity0Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TourismApi_Activity_0",
		Method:             "GET",
		PathPattern:        "/v2/Tourism/Activity/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TourismAPIActivity0Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TourismAPIActivity0OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TourismApi_Activity_0: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TourismAPIHotel 取得所有觀光旅宿資料s

取得所有觀光旅宿資料
*/
func (a *Client) TourismAPIHotel(params *TourismAPIHotelParams) (*TourismAPIHotelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTourismAPIHotelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TourismApi_Hotel",
		Method:             "GET",
		PathPattern:        "/v2/Tourism/Hotel",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TourismAPIHotelReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TourismAPIHotelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TourismApi_Hotel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TourismAPIHotel0 取得指定s 縣市 觀光旅宿資料

取得指定[縣市]觀光旅宿資料
*/
func (a *Client) TourismAPIHotel0(params *TourismAPIHotel0Params) (*TourismAPIHotel0OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTourismAPIHotel0Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TourismApi_Hotel_0",
		Method:             "GET",
		PathPattern:        "/v2/Tourism/Hotel/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TourismAPIHotel0Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TourismAPIHotel0OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TourismApi_Hotel_0: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TourismAPIRestaurant 取得所有觀光餐飲資料s

取得所有觀光餐飲資料
*/
func (a *Client) TourismAPIRestaurant(params *TourismAPIRestaurantParams) (*TourismAPIRestaurantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTourismAPIRestaurantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TourismApi_Restaurant",
		Method:             "GET",
		PathPattern:        "/v2/Tourism/Restaurant",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TourismAPIRestaurantReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TourismAPIRestaurantOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TourismApi_Restaurant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TourismAPIRestaurant0 取得指定s 縣市 觀光餐飲資料

取得指定[縣市]觀光餐飲資料
*/
func (a *Client) TourismAPIRestaurant0(params *TourismAPIRestaurant0Params) (*TourismAPIRestaurant0OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTourismAPIRestaurant0Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TourismApi_Restaurant_0",
		Method:             "GET",
		PathPattern:        "/v2/Tourism/Restaurant/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TourismAPIRestaurant0Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TourismAPIRestaurant0OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TourismApi_Restaurant_0: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TourismAPIScenicSpot 取得所有觀光景點資料s

取得所有觀光景點資料
*/
func (a *Client) TourismAPIScenicSpot(params *TourismAPIScenicSpotParams) (*TourismAPIScenicSpotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTourismAPIScenicSpotParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TourismApi_ScenicSpot",
		Method:             "GET",
		PathPattern:        "/v2/Tourism/ScenicSpot",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TourismAPIScenicSpotReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TourismAPIScenicSpotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TourismApi_ScenicSpot: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TourismAPIScenicSpot0 取得指定s 縣市 觀光景點資料

取得指定[縣市]觀光景點資料
*/
func (a *Client) TourismAPIScenicSpot0(params *TourismAPIScenicSpot0Params) (*TourismAPIScenicSpot0OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTourismAPIScenicSpot0Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TourismApi_ScenicSpot_0",
		Method:             "GET",
		PathPattern:        "/v2/Tourism/ScenicSpot/{City}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TourismAPIScenicSpot0Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TourismAPIScenicSpot0OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TourismApi_ScenicSpot_0: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
