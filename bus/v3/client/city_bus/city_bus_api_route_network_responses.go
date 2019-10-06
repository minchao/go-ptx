// Code generated by go-swagger; DO NOT EDIT.

package city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/bus/v3/models"
)

// CityBusAPIRouteNetworkReader is a Reader for the CityBusAPIRouteNetwork structure.
type CityBusAPIRouteNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRouteNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRouteNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIRouteNetworkStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCityBusAPIRouteNetworkOK creates a CityBusAPIRouteNetworkOK with default headers values
func NewCityBusAPIRouteNetworkOK() *CityBusAPIRouteNetworkOK {
	return &CityBusAPIRouteNetworkOK{}
}

/*CityBusAPIRouteNetworkOK handles this case with default header values.

OK
*/
type CityBusAPIRouteNetworkOK struct {
	Payload *models.MOTCAPIBusDALBusServiceDTOVersion3BusRouteNetwork
}

func (o *CityBusAPIRouteNetworkOK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/RouteNetwork/City/{City}][%d] cityBusApiRouteNetworkOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIRouteNetworkOK) GetPayload() *models.MOTCAPIBusDALBusServiceDTOVersion3BusRouteNetwork {
	return o.Payload
}

func (o *CityBusAPIRouteNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MOTCAPIBusDALBusServiceDTOVersion3BusRouteNetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRouteNetworkStatus299 creates a CityBusAPIRouteNetworkStatus299 with default headers values
func NewCityBusAPIRouteNetworkStatus299() *CityBusAPIRouteNetworkStatus299 {
	return &CityBusAPIRouteNetworkStatus299{}
}

/*CityBusAPIRouteNetworkStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIRouteNetworkStatus299 struct {
	Payload *models.ServiceDTOVersion3BaseHealth
}

func (o *CityBusAPIRouteNetworkStatus299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/RouteNetwork/City/{City}][%d] cityBusApiRouteNetworkStatus299  %+v", 299, o.Payload)
}

func (o *CityBusAPIRouteNetworkStatus299) GetPayload() *models.ServiceDTOVersion3BaseHealth {
	return o.Payload
}

func (o *CityBusAPIRouteNetworkStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion3BaseHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
