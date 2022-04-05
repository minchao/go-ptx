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

// CityBusAPIVehicleRoute3023Reader is a Reader for the CityBusAPIVehicleRoute3023 structure.
type CityBusAPIVehicleRoute3023Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIVehicleRoute3023Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIVehicleRoute3023OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIVehicleRoute3023Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIVehicleRoute3023NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIVehicleRoute3023OK creates a CityBusAPIVehicleRoute3023OK with default headers values
func NewCityBusAPIVehicleRoute3023OK() *CityBusAPIVehicleRoute3023OK {
	return &CityBusAPIVehicleRoute3023OK{}
}

/* CityBusAPIVehicleRoute3023OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIVehicleRoute3023OK struct {
	Payload *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute
}

func (o *CityBusAPIVehicleRoute3023OK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/VehicleRoute/City/{City}][%d] cityBusApiVehicleRoute3023OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIVehicleRoute3023OK) GetPayload() *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute {
	return o.Payload
}

func (o *CityBusAPIVehicleRoute3023OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIVehicleRoute3023Status299 creates a CityBusAPIVehicleRoute3023Status299 with default headers values
func NewCityBusAPIVehicleRoute3023Status299() *CityBusAPIVehicleRoute3023Status299 {
	return &CityBusAPIVehicleRoute3023Status299{}
}

/* CityBusAPIVehicleRoute3023Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIVehicleRoute3023Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIVehicleRoute3023Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/VehicleRoute/City/{City}][%d] cityBusApiVehicleRoute3023Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIVehicleRoute3023Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIVehicleRoute3023Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIVehicleRoute3023NotModified creates a CityBusAPIVehicleRoute3023NotModified with default headers values
func NewCityBusAPIVehicleRoute3023NotModified() *CityBusAPIVehicleRoute3023NotModified {
	return &CityBusAPIVehicleRoute3023NotModified{}
}

/* CityBusAPIVehicleRoute3023NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIVehicleRoute3023NotModified struct {
}

func (o *CityBusAPIVehicleRoute3023NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/VehicleRoute/City/{City}][%d] cityBusApiVehicleRoute3023NotModified ", 304)
}

func (o *CityBusAPIVehicleRoute3023NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
