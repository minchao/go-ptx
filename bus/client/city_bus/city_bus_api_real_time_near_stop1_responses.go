// Code generated by go-swagger; DO NOT EDIT.

package city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/bus/models"
)

// CityBusAPIRealTimeNearStop1Reader is a Reader for the CityBusAPIRealTimeNearStop1 structure.
type CityBusAPIRealTimeNearStop1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRealTimeNearStop1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRealTimeNearStop1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCityBusAPIRealTimeNearStop1OK creates a CityBusAPIRealTimeNearStop1OK with default headers values
func NewCityBusAPIRealTimeNearStop1OK() *CityBusAPIRealTimeNearStop1OK {
	return &CityBusAPIRealTimeNearStop1OK{}
}

/*CityBusAPIRealTimeNearStop1OK handles this case with default header values.

OK
*/
type CityBusAPIRealTimeNearStop1OK struct {
	Payload []*models.ServiceDTOVersion2BusBusA2Data
}

func (o *CityBusAPIRealTimeNearStop1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeNearStop/City/{City}/{RouteName}][%d] cityBusApiRealTimeNearStop1OK  %+v", 200, o.Payload)
}

func (o *CityBusAPIRealTimeNearStop1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
