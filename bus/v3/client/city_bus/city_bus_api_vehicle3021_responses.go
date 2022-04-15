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

// CityBusAPIVehicle3021Reader is a Reader for the CityBusAPIVehicle3021 structure.
type CityBusAPIVehicle3021Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIVehicle3021Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIVehicle3021OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIVehicle3021Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIVehicle3021NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIVehicle3021OK creates a CityBusAPIVehicle3021OK with default headers values
func NewCityBusAPIVehicle3021OK() *CityBusAPIVehicle3021OK {
	return &CityBusAPIVehicle3021OK{}
}

/* CityBusAPIVehicle3021OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIVehicle3021OK struct {
	Payload *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Vehicle
}

func (o *CityBusAPIVehicle3021OK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Vehicle/City/{City}][%d] cityBusApiVehicle3021OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIVehicle3021OK) GetPayload() *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Vehicle {
	return o.Payload
}

func (o *CityBusAPIVehicle3021OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Vehicle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIVehicle3021Status299 creates a CityBusAPIVehicle3021Status299 with default headers values
func NewCityBusAPIVehicle3021Status299() *CityBusAPIVehicle3021Status299 {
	return &CityBusAPIVehicle3021Status299{}
}

/* CityBusAPIVehicle3021Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIVehicle3021Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIVehicle3021Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Vehicle/City/{City}][%d] cityBusApiVehicle3021Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIVehicle3021Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIVehicle3021Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIVehicle3021NotModified creates a CityBusAPIVehicle3021NotModified with default headers values
func NewCityBusAPIVehicle3021NotModified() *CityBusAPIVehicle3021NotModified {
	return &CityBusAPIVehicle3021NotModified{}
}

/* CityBusAPIVehicle3021NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIVehicle3021NotModified struct {
}

func (o *CityBusAPIVehicle3021NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Vehicle/City/{City}][%d] cityBusApiVehicle3021NotModified ", 304)
}

func (o *CityBusAPIVehicle3021NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}