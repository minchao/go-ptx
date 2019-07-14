// Code generated by go-swagger; DO NOT EDIT.

package taiwan_trip_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/tourism/models"
)

// TaiwanTripBusAPINewsReader is a Reader for the TaiwanTripBusAPINews structure.
type TaiwanTripBusAPINewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaiwanTripBusAPINewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaiwanTripBusAPINewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTaiwanTripBusAPINewsOK creates a TaiwanTripBusAPINewsOK with default headers values
func NewTaiwanTripBusAPINewsOK() *TaiwanTripBusAPINewsOK {
	return &TaiwanTripBusAPINewsOK{}
}

/*TaiwanTripBusAPINewsOK handles this case with default header values.

OK
*/
type TaiwanTripBusAPINewsOK struct {
	Payload []*models.ServiceDTOVersion2TaiwanTripBusBusTaiwanTripNews
}

func (o *TaiwanTripBusAPINewsOK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/News/TaiwanTrip][%d] taiwanTripBusApiNewsOK  %+v", 200, o.Payload)
}

func (o *TaiwanTripBusAPINewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
