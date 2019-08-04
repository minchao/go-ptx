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

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCityBusAPIVehicleRouteOK creates a CityBusAPIVehicleRouteOK with default headers values
func NewCityBusAPIVehicleRouteOK() *CityBusAPIVehicleRouteOK {
	return &CityBusAPIVehicleRouteOK{}
}

/*CityBusAPIVehicleRouteOK handles this case with default header values.

OK
*/
type CityBusAPIVehicleRouteOK struct {
	Payload *models.MOTCAPIBusDALBusServiceDTOVersion3BusVehicleRoute
}

func (o *CityBusAPIVehicleRouteOK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/VehicleRoute/City/{City}][%d] cityBusApiVehicleRouteOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIVehicleRouteOK) GetPayload() *models.MOTCAPIBusDALBusServiceDTOVersion3BusVehicleRoute {
	return o.Payload
}

func (o *CityBusAPIVehicleRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MOTCAPIBusDALBusServiceDTOVersion3BusVehicleRoute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}