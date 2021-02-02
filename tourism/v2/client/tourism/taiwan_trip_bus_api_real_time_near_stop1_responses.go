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

// TaiwanTripBusAPIRealTimeNearStop1Reader is a Reader for the TaiwanTripBusAPIRealTimeNearStop1 structure.
type TaiwanTripBusAPIRealTimeNearStop1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaiwanTripBusAPIRealTimeNearStop1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaiwanTripBusAPIRealTimeNearStop1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTaiwanTripBusAPIRealTimeNearStop1OK creates a TaiwanTripBusAPIRealTimeNearStop1OK with default headers values
func NewTaiwanTripBusAPIRealTimeNearStop1OK() *TaiwanTripBusAPIRealTimeNearStop1OK {
	return &TaiwanTripBusAPIRealTimeNearStop1OK{}
}

/* TaiwanTripBusAPIRealTimeNearStop1OK describes a response with status code 200, with default header values.

Success
*/
type TaiwanTripBusAPIRealTimeNearStop1OK struct {
	Payload []*models.PTXServiceDTOTourismSpecificationV2BusA2Data
}

func (o *TaiwanTripBusAPIRealTimeNearStop1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/RealTimeNearStop/TaiwanTrip/{TaiwanTripName}][%d] taiwanTripBusApiRealTimeNearStop1OK  %+v", 200, o.Payload)
}
func (o *TaiwanTripBusAPIRealTimeNearStop1OK) GetPayload() []*models.PTXServiceDTOTourismSpecificationV2BusA2Data {
	return o.Payload
}

func (o *TaiwanTripBusAPIRealTimeNearStop1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
