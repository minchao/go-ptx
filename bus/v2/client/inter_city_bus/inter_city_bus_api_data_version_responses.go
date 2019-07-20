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

// InterCityBusAPIDataVersionReader is a Reader for the InterCityBusAPIDataVersion structure.
type InterCityBusAPIDataVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIDataVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIDataVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInterCityBusAPIDataVersionOK creates a InterCityBusAPIDataVersionOK with default headers values
func NewInterCityBusAPIDataVersionOK() *InterCityBusAPIDataVersionOK {
	return &InterCityBusAPIDataVersionOK{}
}

/*InterCityBusAPIDataVersionOK handles this case with default header values.

OK
*/
type InterCityBusAPIDataVersionOK struct {
	Payload *models.ServiceDTOVersion2BusBusVersion
}

func (o *InterCityBusAPIDataVersionOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/DataVersion/InterCity][%d] interCityBusApiDataVersionOK  %+v", 200, o.Payload)
}

func (o *InterCityBusAPIDataVersionOK) GetPayload() *models.ServiceDTOVersion2BusBusVersion {
	return o.Payload
}

func (o *InterCityBusAPIDataVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion2BusBusVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
