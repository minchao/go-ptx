// Code generated by go-swagger; DO NOT EDIT.

package air

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new air API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for air API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AirAPIAirline(params *AirAPIAirlineParams) (*AirAPIAirlineOK, error)

	AirAPIAirline1(params *AirAPIAirline1Params) (*AirAPIAirline1OK, error)

	AirAPIAirport(params *AirAPIAirportParams) (*AirAPIAirportOK, error)

	AirAPIAirport1(params *AirAPIAirport1Params) (*AirAPIAirport1OK, error)

	AirAPIArrival(params *AirAPIArrivalParams) (*AirAPIArrivalOK, error)

	AirAPIArrival1(params *AirAPIArrival1Params) (*AirAPIArrival1OK, error)

	AirAPIDeparture(params *AirAPIDepartureParams) (*AirAPIDepartureOK, error)

	AirAPIDeparture1(params *AirAPIDeparture1Params) (*AirAPIDeparture1OK, error)

	AirAPIDomestic(params *AirAPIDomesticParams) (*AirAPIDomesticOK, error)

	AirAPIFIDS(params *AirAPIFIDSParams) (*AirAPIFIDSOK, error)

	AirAPIFIDS1(params *AirAPIFIDS1Params) (*AirAPIFIDS1OK, error)

	AirAPIFlight(params *AirAPIFlightParams) (*AirAPIFlightOK, error)

	AirAPIFlight1(params *AirAPIFlight1Params) (*AirAPIFlight1OK, error)

	AirAPIInternational(params *AirAPIInternationalParams) (*AirAPIInternationalOK, error)

	AirAPIMETAR(params *AirAPIMETARParams) (*AirAPIMETAROK, error)

	AirAPIMETAR1(params *AirAPIMETAR1Params) (*AirAPIMETAR1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AirAPIAirline 取得指定s 航空公司 資料

  取得指定[航空公司]資料
*/
func (a *Client) AirAPIAirline(params *AirAPIAirlineParams) (*AirAPIAirlineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAirAPIAirlineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AirApi_Airline",
		Method:             "GET",
		PathPattern:        "/v2/Air/Airline/{IATA}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AirAPIAirlineReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AirAPIAirlineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AirApi_Airline: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AirAPIAirline1 取得航空公司資料s

  取得所有航空公司資料
*/
func (a *Client) AirAPIAirline1(params *AirAPIAirline1Params) (*AirAPIAirline1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAirAPIAirline1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AirApi_Airline_1",
		Method:             "GET",
		PathPattern:        "/v2/Air/Airline",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AirAPIAirline1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AirAPIAirline1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AirApi_Airline_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AirAPIAirport 取得機場資料s

  取得所有機場資料
*/
func (a *Client) AirAPIAirport(params *AirAPIAirportParams) (*AirAPIAirportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAirAPIAirportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AirApi_Airport",
		Method:             "GET",
		PathPattern:        "/v2/Air/Airport",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AirAPIAirportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AirAPIAirportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AirApi_Airport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AirAPIAirport1 取得指定s 機場 資料

  取得指定[機場]資料
*/
func (a *Client) AirAPIAirport1(params *AirAPIAirport1Params) (*AirAPIAirport1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAirAPIAirport1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AirApi_Airport_1",
		Method:             "GET",
		PathPattern:        "/v2/Air/Airport/{IATA}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AirAPIAirport1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AirAPIAirport1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AirApi_Airport_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AirAPIArrival 取得機場的即時入境航班s

  取得機場的即時入境航班
*/
func (a *Client) AirAPIArrival(params *AirAPIArrivalParams) (*AirAPIArrivalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAirAPIArrivalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AirApi_Arrival",
		Method:             "GET",
		PathPattern:        "/v2/Air/FIDS/Airport/Arrival",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AirAPIArrivalReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AirAPIArrivalOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AirApi_Arrival: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AirAPIArrival1 取得指定s 機場的即時入境航班

  取得[指定機場]的即時入境航班
*/
func (a *Client) AirAPIArrival1(params *AirAPIArrival1Params) (*AirAPIArrival1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAirAPIArrival1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AirApi_Arrival_1",
		Method:             "GET",
		PathPattern:        "/v2/Air/FIDS/Airport/Arrival/{IATA}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AirAPIArrival1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AirAPIArrival1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AirApi_Arrival_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AirAPIDeparture 取得機場的即時出境航班s

  取得機場的即時出境航班
*/
func (a *Client) AirAPIDeparture(params *AirAPIDepartureParams) (*AirAPIDepartureOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAirAPIDepartureParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AirApi_Departure",
		Method:             "GET",
		PathPattern:        "/v2/Air/FIDS/Airport/Departure",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AirAPIDepartureReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AirAPIDepartureOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AirApi_Departure: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AirAPIDeparture1 取得指定s 機場的即時出境航班

  取得指定的[機場即時出境航班]
*/
func (a *Client) AirAPIDeparture1(params *AirAPIDeparture1Params) (*AirAPIDeparture1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAirAPIDeparture1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AirApi_Departure_1",
		Method:             "GET",
		PathPattern:        "/v2/Air/FIDS/Airport/Departure/{IATA}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AirAPIDeparture1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AirAPIDeparture1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AirApi_Departure_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AirAPIDomestic 取得國內航空定期時刻表s
*/
func (a *Client) AirAPIDomestic(params *AirAPIDomesticParams) (*AirAPIDomesticOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAirAPIDomesticParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AirApi_Domestic",
		Method:             "GET",
		PathPattern:        "/v2/Air/GeneralSchedule/Domestic",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AirAPIDomesticReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AirAPIDomesticOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AirApi_Domestic: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AirAPIFIDS 取得機場的即時航班資料s

  取得即時航班資料
*/
func (a *Client) AirAPIFIDS(params *AirAPIFIDSParams) (*AirAPIFIDSOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAirAPIFIDSParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AirApi_FIDS",
		Method:             "GET",
		PathPattern:        "/v2/Air/FIDS/Airport",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AirAPIFIDSReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AirAPIFIDSOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AirApi_FIDS: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AirAPIFIDS1 取得指定s 機場的即時航班 資料

  取得指定[機場的即時航班]資料
*/
func (a *Client) AirAPIFIDS1(params *AirAPIFIDS1Params) (*AirAPIFIDS1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAirAPIFIDS1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AirApi_FIDS_1",
		Method:             "GET",
		PathPattern:        "/v2/Air/FIDS/Airport/{IATA}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AirAPIFIDS1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AirAPIFIDS1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AirApi_FIDS_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AirAPIFlight 取得即時航班資料s

  取得即時航班資料
*/
func (a *Client) AirAPIFlight(params *AirAPIFlightParams) (*AirAPIFlightOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAirAPIFlightParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AirApi_Flight",
		Method:             "GET",
		PathPattern:        "/v2/Air/FIDS/Flight",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AirAPIFlightReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AirAPIFlightOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AirApi_Flight: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AirAPIFlight1 取得指定s 即時航班 資料
*/
func (a *Client) AirAPIFlight1(params *AirAPIFlight1Params) (*AirAPIFlight1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAirAPIFlight1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AirApi_Flight_1",
		Method:             "GET",
		PathPattern:        "/v2/Air/FIDS/Flight/{FlightNo}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AirAPIFlight1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AirAPIFlight1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AirApi_Flight_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AirAPIInternational 取得國際航空定期時刻表s
*/
func (a *Client) AirAPIInternational(params *AirAPIInternationalParams) (*AirAPIInternationalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAirAPIInternationalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AirApi_International",
		Method:             "GET",
		PathPattern:        "/v2/Air/GeneralSchedule/International",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AirAPIInternationalReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AirAPIInternationalOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AirApi_International: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AirAPIMETAR 取得s 國內機場 氣象資訊觀測資料

  取得[國內機場]氣象資訊觀測資料
<br />目前提供之機場氣象觀測站代碼清單
<br />RCSS:松山機場    RCKH:高雄機場
<br />RCTP:桃園機場    RCMT:北竿機場
<br />RCFG:南竿機場    RCBS:金門機場
<br />RCQC:馬公機場    RCMQ:臺中機場
<br />RCKU:嘉義機場    RCNN:臺南機場
<br />RCYU:花蓮機場    RCFN:臺東機場
<br />RCLY:蘭嶼機場    RCGI:綠島機場
*/
func (a *Client) AirAPIMETAR(params *AirAPIMETARParams) (*AirAPIMETAROK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAirAPIMETARParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AirApi_METAR",
		Method:             "GET",
		PathPattern:        "/v2/Air/METAR/Airport",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AirAPIMETARReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AirAPIMETAROK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AirApi_METAR: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AirAPIMETAR1 取得指定s 國內機場 氣象資訊觀測資料

  取得指定[國內機場]氣象資訊觀測資料
*/
func (a *Client) AirAPIMETAR1(params *AirAPIMETAR1Params) (*AirAPIMETAR1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAirAPIMETAR1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AirApi_METAR_1",
		Method:             "GET",
		PathPattern:        "/v2/Air/METAR/Airport/{IATA}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AirAPIMETAR1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AirAPIMETAR1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AirApi_METAR_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
