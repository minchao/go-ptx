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

// CityBusAPIS2STravelTimeReader is a Reader for the CityBusAPIS2STravelTime structure.
type CityBusAPIS2STravelTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIS2STravelTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIS2STravelTimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCityBusAPIS2STravelTimeOK creates a CityBusAPIS2STravelTimeOK with default headers values
func NewCityBusAPIS2STravelTimeOK() *CityBusAPIS2STravelTimeOK {
	return &CityBusAPIS2STravelTimeOK{}
}

/*CityBusAPIS2STravelTimeOK handles this case with default header values.

OK
*/
type CityBusAPIS2STravelTimeOK struct {
	Payload *models.MOTCAPIBusDALBusServiceDTOVersion3BusS2STravelTime
}

func (o *CityBusAPIS2STravelTimeOK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/S2STravelTime/City/{City}][%d] cityBusApiS2STravelTimeOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIS2STravelTimeOK) GetPayload() *models.MOTCAPIBusDALBusServiceDTOVersion3BusS2STravelTime {
	return o.Payload
}

func (o *CityBusAPIS2STravelTimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MOTCAPIBusDALBusServiceDTOVersion3BusS2STravelTime)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}