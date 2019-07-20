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

// CityBusAPIRouteFare1Reader is a Reader for the CityBusAPIRouteFare1 structure.
type CityBusAPIRouteFare1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRouteFare1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRouteFare1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCityBusAPIRouteFare1OK creates a CityBusAPIRouteFare1OK with default headers values
func NewCityBusAPIRouteFare1OK() *CityBusAPIRouteFare1OK {
	return &CityBusAPIRouteFare1OK{}
}

/*CityBusAPIRouteFare1OK handles this case with default header values.

OK
*/
type CityBusAPIRouteFare1OK struct {
	Payload []*models.ServiceDTOVersion2BusBusRouteFare
}

func (o *CityBusAPIRouteFare1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RouteFare/City/{City}/{RouteName}][%d] cityBusApiRouteFare1OK  %+v", 200, o.Payload)
}

func (o *CityBusAPIRouteFare1OK) GetPayload() []*models.ServiceDTOVersion2BusBusRouteFare {
	return o.Payload
}

func (o *CityBusAPIRouteFare1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
