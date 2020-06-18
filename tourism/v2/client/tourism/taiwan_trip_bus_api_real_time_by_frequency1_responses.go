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

// TaiwanTripBusAPIRealTimeByFrequency1Reader is a Reader for the TaiwanTripBusAPIRealTimeByFrequency1 structure.
type TaiwanTripBusAPIRealTimeByFrequency1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaiwanTripBusAPIRealTimeByFrequency1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaiwanTripBusAPIRealTimeByFrequency1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTaiwanTripBusAPIRealTimeByFrequency1OK creates a TaiwanTripBusAPIRealTimeByFrequency1OK with default headers values
func NewTaiwanTripBusAPIRealTimeByFrequency1OK() *TaiwanTripBusAPIRealTimeByFrequency1OK {
	return &TaiwanTripBusAPIRealTimeByFrequency1OK{}
}

/*TaiwanTripBusAPIRealTimeByFrequency1OK handles this case with default header values.

OK
*/
type TaiwanTripBusAPIRealTimeByFrequency1OK struct {
	Payload []*models.ServiceDTOVersion2TaiwanTripBusBusA1Data
}

func (o *TaiwanTripBusAPIRealTimeByFrequency1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/RealTimeByFrequency/TaiwanTrip/{TaiwanTripName}][%d] taiwanTripBusApiRealTimeByFrequency1OK  %+v", 200, o.Payload)
}

func (o *TaiwanTripBusAPIRealTimeByFrequency1OK) GetPayload() []*models.ServiceDTOVersion2TaiwanTripBusBusA1Data {
	return o.Payload
}

func (o *TaiwanTripBusAPIRealTimeByFrequency1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
