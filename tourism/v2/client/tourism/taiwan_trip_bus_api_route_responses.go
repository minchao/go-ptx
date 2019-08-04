// Code generated by go-swagger; DO NOT EDIT.

package tourism

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/tourism/v2/models"
)

// TaiwanTripBusAPIRouteReader is a Reader for the TaiwanTripBusAPIRoute structure.
type TaiwanTripBusAPIRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaiwanTripBusAPIRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaiwanTripBusAPIRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTaiwanTripBusAPIRouteOK creates a TaiwanTripBusAPIRouteOK with default headers values
func NewTaiwanTripBusAPIRouteOK() *TaiwanTripBusAPIRouteOK {
	return &TaiwanTripBusAPIRouteOK{}
}

/*TaiwanTripBusAPIRouteOK handles this case with default header values.

OK
*/
type TaiwanTripBusAPIRouteOK struct {
	Payload []*models.ServiceDTOVersion2TaiwanTripBusBusRoute
}

func (o *TaiwanTripBusAPIRouteOK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/Route/TaiwanTrip][%d] taiwanTripBusApiRouteOK  %+v", 200, o.Payload)
}

func (o *TaiwanTripBusAPIRouteOK) GetPayload() []*models.ServiceDTOVersion2TaiwanTripBusBusRoute {
	return o.Payload
}

func (o *TaiwanTripBusAPIRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}