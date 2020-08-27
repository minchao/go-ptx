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

// TaiwanTripBusAPIStopOfRoute1Reader is a Reader for the TaiwanTripBusAPIStopOfRoute1 structure.
type TaiwanTripBusAPIStopOfRoute1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaiwanTripBusAPIStopOfRoute1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaiwanTripBusAPIStopOfRoute1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTaiwanTripBusAPIStopOfRoute1OK creates a TaiwanTripBusAPIStopOfRoute1OK with default headers values
func NewTaiwanTripBusAPIStopOfRoute1OK() *TaiwanTripBusAPIStopOfRoute1OK {
	return &TaiwanTripBusAPIStopOfRoute1OK{}
}

/*TaiwanTripBusAPIStopOfRoute1OK handles this case with default header values.

OK
*/
type TaiwanTripBusAPIStopOfRoute1OK struct {
	Payload []*models.ServiceDTOVersion2TaiwanTripBusBusStopOfRoute
}

func (o *TaiwanTripBusAPIStopOfRoute1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/StopOfRoute/TaiwanTrip/{TaiwanTripName}][%d] taiwanTripBusApiStopOfRoute1OK  %+v", 200, o.Payload)
}

func (o *TaiwanTripBusAPIStopOfRoute1OK) GetPayload() []*models.ServiceDTOVersion2TaiwanTripBusBusStopOfRoute {
	return o.Payload
}

func (o *TaiwanTripBusAPIStopOfRoute1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
