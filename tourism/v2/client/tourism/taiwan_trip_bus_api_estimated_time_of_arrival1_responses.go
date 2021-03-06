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

// TaiwanTripBusAPIEstimatedTimeOfArrival1Reader is a Reader for the TaiwanTripBusAPIEstimatedTimeOfArrival1 structure.
type TaiwanTripBusAPIEstimatedTimeOfArrival1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaiwanTripBusAPIEstimatedTimeOfArrival1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTaiwanTripBusAPIEstimatedTimeOfArrival1OK creates a TaiwanTripBusAPIEstimatedTimeOfArrival1OK with default headers values
func NewTaiwanTripBusAPIEstimatedTimeOfArrival1OK() *TaiwanTripBusAPIEstimatedTimeOfArrival1OK {
	return &TaiwanTripBusAPIEstimatedTimeOfArrival1OK{}
}

/* TaiwanTripBusAPIEstimatedTimeOfArrival1OK describes a response with status code 200, with default header values.

Success
*/
type TaiwanTripBusAPIEstimatedTimeOfArrival1OK struct {
	Payload []*models.PTXServiceDTOTourismSpecificationV2BusN1EstimateTime
}

func (o *TaiwanTripBusAPIEstimatedTimeOfArrival1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/EstimatedTimeOfArrival/TaiwanTrip/{TaiwanTripName}][%d] taiwanTripBusApiEstimatedTimeOfArrival1OK  %+v", 200, o.Payload)
}
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival1OK) GetPayload() []*models.PTXServiceDTOTourismSpecificationV2BusN1EstimateTime {
	return o.Payload
}

func (o *TaiwanTripBusAPIEstimatedTimeOfArrival1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
