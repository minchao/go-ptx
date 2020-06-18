// Code generated by go-swagger; DO NOT EDIT.

package tourism

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/tourism/v2/models"
)

// TaiwanTripBusAPIS2TravelTimeReader is a Reader for the TaiwanTripBusAPIS2TravelTime structure.
type TaiwanTripBusAPIS2TravelTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaiwanTripBusAPIS2TravelTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaiwanTripBusAPIS2TravelTimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTaiwanTripBusAPIS2TravelTimeOK creates a TaiwanTripBusAPIS2TravelTimeOK with default headers values
func NewTaiwanTripBusAPIS2TravelTimeOK() *TaiwanTripBusAPIS2TravelTimeOK {
	return &TaiwanTripBusAPIS2TravelTimeOK{}
}

/*TaiwanTripBusAPIS2TravelTimeOK handles this case with default header values.

OK
*/
type TaiwanTripBusAPIS2TravelTimeOK struct {
	Payload []*models.ServiceDTOVersion2TaiwanTripBusBusTaiwanTripS2STravelTime
}

func (o *TaiwanTripBusAPIS2TravelTimeOK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/S2TravelTime/TaiwanTrip][%d] taiwanTripBusApiS2TravelTimeOK  %+v", 200, o.Payload)
}

func (o *TaiwanTripBusAPIS2TravelTimeOK) GetPayload() []*models.ServiceDTOVersion2TaiwanTripBusBusTaiwanTripS2STravelTime {
	return o.Payload
}

func (o *TaiwanTripBusAPIS2TravelTimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
