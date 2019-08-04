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

// CityBusAPIEstimatedTimeOfArrival2Reader is a Reader for the CityBusAPIEstimatedTimeOfArrival2 structure.
type CityBusAPIEstimatedTimeOfArrival2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIEstimatedTimeOfArrival2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIEstimatedTimeOfArrival2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCityBusAPIEstimatedTimeOfArrival2OK creates a CityBusAPIEstimatedTimeOfArrival2OK with default headers values
func NewCityBusAPIEstimatedTimeOfArrival2OK() *CityBusAPIEstimatedTimeOfArrival2OK {
	return &CityBusAPIEstimatedTimeOfArrival2OK{}
}

/*CityBusAPIEstimatedTimeOfArrival2OK handles this case with default header values.

OK
*/
type CityBusAPIEstimatedTimeOfArrival2OK struct {
	Payload *models.MOTCAPIBusDALBusServiceDTOVersion3BusN1Data
}

func (o *CityBusAPIEstimatedTimeOfArrival2OK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/EstimatedTimeOfArrival/City/{City}][%d] cityBusApiEstimatedTimeOfArrival2OK  %+v", 200, o.Payload)
}

func (o *CityBusAPIEstimatedTimeOfArrival2OK) GetPayload() *models.MOTCAPIBusDALBusServiceDTOVersion3BusN1Data {
	return o.Payload
}

func (o *CityBusAPIEstimatedTimeOfArrival2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MOTCAPIBusDALBusServiceDTOVersion3BusN1Data)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
