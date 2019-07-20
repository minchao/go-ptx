// Code generated by go-swagger; DO NOT EDIT.

package city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/bus/v2/models"
)

// CityBusAPIRouteReader is a Reader for the CityBusAPIRoute structure.
type CityBusAPIRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCityBusAPIRouteOK creates a CityBusAPIRouteOK with default headers values
func NewCityBusAPIRouteOK() *CityBusAPIRouteOK {
	return &CityBusAPIRouteOK{}
}

/*CityBusAPIRouteOK handles this case with default header values.

OK
*/
type CityBusAPIRouteOK struct {
	Payload []*models.ServiceDTOVersion2BusBusRoute
}

func (o *CityBusAPIRouteOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/City/{City}][%d] cityBusApiRouteOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIRouteOK) GetPayload() []*models.ServiceDTOVersion2BusBusRoute {
	return o.Payload
}

func (o *CityBusAPIRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
