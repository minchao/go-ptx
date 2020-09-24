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

// CityBusAPIVehicleRouteReader is a Reader for the CityBusAPIVehicleRoute structure.
type CityBusAPIVehicleRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIVehicleRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIVehicleRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIVehicleRouteStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIVehicleRouteOK creates a CityBusAPIVehicleRouteOK with default headers values
func NewCityBusAPIVehicleRouteOK() *CityBusAPIVehicleRouteOK {
	return &CityBusAPIVehicleRouteOK{}
}

/*CityBusAPIVehicleRouteOK handles this case with default header values.

Success
*/
type CityBusAPIVehicleRouteOK struct {
	Payload *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute
}

func (o *CityBusAPIVehicleRouteOK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/VehicleRoute/City/{City}][%d] cityBusApiVehicleRouteOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIVehicleRouteOK) GetPayload() *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute {
	return o.Payload
}

func (o *CityBusAPIVehicleRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIVehicleRouteStatus299 creates a CityBusAPIVehicleRouteStatus299 with default headers values
func NewCityBusAPIVehicleRouteStatus299() *CityBusAPIVehicleRouteStatus299 {
	return &CityBusAPIVehicleRouteStatus299{}
}

/*CityBusAPIVehicleRouteStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIVehicleRouteStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIVehicleRouteStatus299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/VehicleRoute/City/{City}][%d] cityBusApiVehicleRouteStatus299  %+v", 299, o.Payload)
}

func (o *CityBusAPIVehicleRouteStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIVehicleRouteStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
