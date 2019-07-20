// Code generated by go-swagger; DO NOT EDIT.

package taiwan_trip_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/tourism/v2/models"
)

// TaiwanTripBusAPIScheduleReader is a Reader for the TaiwanTripBusAPISchedule structure.
type TaiwanTripBusAPIScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaiwanTripBusAPIScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaiwanTripBusAPIScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTaiwanTripBusAPIScheduleOK creates a TaiwanTripBusAPIScheduleOK with default headers values
func NewTaiwanTripBusAPIScheduleOK() *TaiwanTripBusAPIScheduleOK {
	return &TaiwanTripBusAPIScheduleOK{}
}

/*TaiwanTripBusAPIScheduleOK handles this case with default header values.

OK
*/
type TaiwanTripBusAPIScheduleOK struct {
	Payload []*models.ServiceDTOVersion2TaiwanTripBusBusSchedule
}

func (o *TaiwanTripBusAPIScheduleOK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/Schedule/TaiwanTrip][%d] taiwanTripBusApiScheduleOK  %+v", 200, o.Payload)
}

func (o *TaiwanTripBusAPIScheduleOK) GetPayload() []*models.ServiceDTOVersion2TaiwanTripBusBusSchedule {
	return o.Payload
}

func (o *TaiwanTripBusAPIScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
