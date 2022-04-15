// Code generated by go-swagger; DO NOT EDIT.

package city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v3/models"
)

// CityBusAPIVehicleDepot3022Reader is a Reader for the CityBusAPIVehicleDepot3022 structure.
type CityBusAPIVehicleDepot3022Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIVehicleDepot3022Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIVehicleDepot3022OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIVehicleDepot3022Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIVehicleDepot3022NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIVehicleDepot3022OK creates a CityBusAPIVehicleDepot3022OK with default headers values
func NewCityBusAPIVehicleDepot3022OK() *CityBusAPIVehicleDepot3022OK {
	return &CityBusAPIVehicleDepot3022OK{}
}

/* CityBusAPIVehicleDepot3022OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIVehicleDepot3022OK struct {
	Payload *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleDepot
}

func (o *CityBusAPIVehicleDepot3022OK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/VehicleDepot/City/{City}][%d] cityBusApiVehicleDepot3022OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIVehicleDepot3022OK) GetPayload() *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleDepot {
	return o.Payload
}

func (o *CityBusAPIVehicleDepot3022OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleDepot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIVehicleDepot3022Status299 creates a CityBusAPIVehicleDepot3022Status299 with default headers values
func NewCityBusAPIVehicleDepot3022Status299() *CityBusAPIVehicleDepot3022Status299 {
	return &CityBusAPIVehicleDepot3022Status299{}
}

/* CityBusAPIVehicleDepot3022Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIVehicleDepot3022Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIVehicleDepot3022Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/VehicleDepot/City/{City}][%d] cityBusApiVehicleDepot3022Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIVehicleDepot3022Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIVehicleDepot3022Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIVehicleDepot3022NotModified creates a CityBusAPIVehicleDepot3022NotModified with default headers values
func NewCityBusAPIVehicleDepot3022NotModified() *CityBusAPIVehicleDepot3022NotModified {
	return &CityBusAPIVehicleDepot3022NotModified{}
}

/* CityBusAPIVehicleDepot3022NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIVehicleDepot3022NotModified struct {
}

func (o *CityBusAPIVehicleDepot3022NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/VehicleDepot/City/{City}][%d] cityBusApiVehicleDepot3022NotModified ", 304)
}

func (o *CityBusAPIVehicleDepot3022NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}