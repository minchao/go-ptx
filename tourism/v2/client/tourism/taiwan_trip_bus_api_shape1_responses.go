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

// TaiwanTripBusAPIShape1Reader is a Reader for the TaiwanTripBusAPIShape1 structure.
type TaiwanTripBusAPIShape1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaiwanTripBusAPIShape1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaiwanTripBusAPIShape1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTaiwanTripBusAPIShape1OK creates a TaiwanTripBusAPIShape1OK with default headers values
func NewTaiwanTripBusAPIShape1OK() *TaiwanTripBusAPIShape1OK {
	return &TaiwanTripBusAPIShape1OK{}
}

/*TaiwanTripBusAPIShape1OK handles this case with default header values.

OK
*/
type TaiwanTripBusAPIShape1OK struct {
	Payload []*models.ServiceDTOVersion2TaiwanTripBusBusTaiwanTripShape
}

func (o *TaiwanTripBusAPIShape1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/Shape/TaiwanTrip/{TaiwanTripName}][%d] taiwanTripBusApiShape1OK  %+v", 200, o.Payload)
}

func (o *TaiwanTripBusAPIShape1OK) GetPayload() []*models.ServiceDTOVersion2TaiwanTripBusBusTaiwanTripShape {
	return o.Payload
}

func (o *TaiwanTripBusAPIShape1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}