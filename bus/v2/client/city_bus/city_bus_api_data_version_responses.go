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

// CityBusAPIDataVersionReader is a Reader for the CityBusAPIDataVersion structure.
type CityBusAPIDataVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIDataVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIDataVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCityBusAPIDataVersionOK creates a CityBusAPIDataVersionOK with default headers values
func NewCityBusAPIDataVersionOK() *CityBusAPIDataVersionOK {
	return &CityBusAPIDataVersionOK{}
}

/*CityBusAPIDataVersionOK handles this case with default header values.

OK
*/
type CityBusAPIDataVersionOK struct {
	Payload *models.ServiceDTOVersion2BusBusVersion
}

func (o *CityBusAPIDataVersionOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/DataVersion/City/{City}][%d] cityBusApiDataVersionOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIDataVersionOK) GetPayload() *models.ServiceDTOVersion2BusBusVersion {
	return o.Payload
}

func (o *CityBusAPIDataVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion2BusBusVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
