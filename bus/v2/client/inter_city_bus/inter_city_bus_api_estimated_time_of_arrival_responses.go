// Code generated by go-swagger; DO NOT EDIT.

package inter_city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/bus/v2/models"
)

// InterCityBusAPIEstimatedTimeOfArrivalReader is a Reader for the InterCityBusAPIEstimatedTimeOfArrival structure.
type InterCityBusAPIEstimatedTimeOfArrivalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIEstimatedTimeOfArrivalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIEstimatedTimeOfArrivalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInterCityBusAPIEstimatedTimeOfArrivalOK creates a InterCityBusAPIEstimatedTimeOfArrivalOK with default headers values
func NewInterCityBusAPIEstimatedTimeOfArrivalOK() *InterCityBusAPIEstimatedTimeOfArrivalOK {
	return &InterCityBusAPIEstimatedTimeOfArrivalOK{}
}

/*InterCityBusAPIEstimatedTimeOfArrivalOK handles this case with default header values.

OK
*/
type InterCityBusAPIEstimatedTimeOfArrivalOK struct {
	Payload []*models.ServiceDTOVersion2BusBusN1EstimateTime
}

func (o *InterCityBusAPIEstimatedTimeOfArrivalOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/EstimatedTimeOfArrival/InterCity][%d] interCityBusApiEstimatedTimeOfArrivalOK  %+v", 200, o.Payload)
}

func (o *InterCityBusAPIEstimatedTimeOfArrivalOK) GetPayload() []*models.ServiceDTOVersion2BusBusN1EstimateTime {
	return o.Payload
}

func (o *InterCityBusAPIEstimatedTimeOfArrivalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
