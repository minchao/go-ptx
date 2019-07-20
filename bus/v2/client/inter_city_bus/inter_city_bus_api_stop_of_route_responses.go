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

// InterCityBusAPIStopOfRouteReader is a Reader for the InterCityBusAPIStopOfRoute structure.
type InterCityBusAPIStopOfRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIStopOfRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIStopOfRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInterCityBusAPIStopOfRouteOK creates a InterCityBusAPIStopOfRouteOK with default headers values
func NewInterCityBusAPIStopOfRouteOK() *InterCityBusAPIStopOfRouteOK {
	return &InterCityBusAPIStopOfRouteOK{}
}

/*InterCityBusAPIStopOfRouteOK handles this case with default header values.

OK
*/
type InterCityBusAPIStopOfRouteOK struct {
	Payload []*models.ServiceDTOVersion2BusBusStopOfRoute
}

func (o *InterCityBusAPIStopOfRouteOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/StopOfRoute/InterCity][%d] interCityBusApiStopOfRouteOK  %+v", 200, o.Payload)
}

func (o *InterCityBusAPIStopOfRouteOK) GetPayload() []*models.ServiceDTOVersion2BusBusStopOfRoute {
	return o.Payload
}

func (o *InterCityBusAPIStopOfRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
