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

// CityBusAPIVehicleReader is a Reader for the CityBusAPIVehicle structure.
type CityBusAPIVehicleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIVehicleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIVehicleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIVehicleStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIVehicleOK creates a CityBusAPIVehicleOK with default headers values
func NewCityBusAPIVehicleOK() *CityBusAPIVehicleOK {
	return &CityBusAPIVehicleOK{}
}

/*CityBusAPIVehicleOK handles this case with default header values.

Success
*/
type CityBusAPIVehicleOK struct {
	Payload *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Vehicle
}

func (o *CityBusAPIVehicleOK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Vehicle/City/{City}][%d] cityBusApiVehicleOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIVehicleOK) GetPayload() *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Vehicle {
	return o.Payload
}

func (o *CityBusAPIVehicleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Vehicle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIVehicleStatus299 creates a CityBusAPIVehicleStatus299 with default headers values
func NewCityBusAPIVehicleStatus299() *CityBusAPIVehicleStatus299 {
	return &CityBusAPIVehicleStatus299{}
}

/*CityBusAPIVehicleStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIVehicleStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIVehicleStatus299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Vehicle/City/{City}][%d] cityBusApiVehicleStatus299  %+v", 299, o.Payload)
}

func (o *CityBusAPIVehicleStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIVehicleStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}