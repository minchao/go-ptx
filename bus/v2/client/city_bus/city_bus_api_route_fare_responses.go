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

// CityBusAPIRouteFareReader is a Reader for the CityBusAPIRouteFare structure.
type CityBusAPIRouteFareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRouteFareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRouteFareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCityBusAPIRouteFareOK creates a CityBusAPIRouteFareOK with default headers values
func NewCityBusAPIRouteFareOK() *CityBusAPIRouteFareOK {
	return &CityBusAPIRouteFareOK{}
}

/*CityBusAPIRouteFareOK handles this case with default header values.

OK
*/
type CityBusAPIRouteFareOK struct {
	Payload []*models.ServiceDTOVersion2BusBusRouteFare
}

func (o *CityBusAPIRouteFareOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RouteFare/City/{City}][%d] cityBusApiRouteFareOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIRouteFareOK) GetPayload() []*models.ServiceDTOVersion2BusBusRouteFare {
	return o.Payload
}

func (o *CityBusAPIRouteFareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
