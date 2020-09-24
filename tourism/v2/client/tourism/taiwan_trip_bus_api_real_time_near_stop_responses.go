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

// TaiwanTripBusAPIRealTimeNearStopReader is a Reader for the TaiwanTripBusAPIRealTimeNearStop structure.
type TaiwanTripBusAPIRealTimeNearStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaiwanTripBusAPIRealTimeNearStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaiwanTripBusAPIRealTimeNearStopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTaiwanTripBusAPIRealTimeNearStopOK creates a TaiwanTripBusAPIRealTimeNearStopOK with default headers values
func NewTaiwanTripBusAPIRealTimeNearStopOK() *TaiwanTripBusAPIRealTimeNearStopOK {
	return &TaiwanTripBusAPIRealTimeNearStopOK{}
}

/*TaiwanTripBusAPIRealTimeNearStopOK handles this case with default header values.

Success
*/
type TaiwanTripBusAPIRealTimeNearStopOK struct {
	Payload []*models.PTXServiceDTOTourismSpecificationV2BusA2Data
}

func (o *TaiwanTripBusAPIRealTimeNearStopOK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/RealTimeNearStop/TaiwanTrip][%d] taiwanTripBusApiRealTimeNearStopOK  %+v", 200, o.Payload)
}

func (o *TaiwanTripBusAPIRealTimeNearStopOK) GetPayload() []*models.PTXServiceDTOTourismSpecificationV2BusA2Data {
	return o.Payload
}

func (o *TaiwanTripBusAPIRealTimeNearStopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
