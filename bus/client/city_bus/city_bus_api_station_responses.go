// Code generated by go-swagger; DO NOT EDIT.

package city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/bus/models"
)

// CityBusAPIStationReader is a Reader for the CityBusAPIStation structure.
type CityBusAPIStationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIStationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIStationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCityBusAPIStationOK creates a CityBusAPIStationOK with default headers values
func NewCityBusAPIStationOK() *CityBusAPIStationOK {
	return &CityBusAPIStationOK{}
}

/*CityBusAPIStationOK handles this case with default header values.

OK
*/
type CityBusAPIStationOK struct {
	Payload []*models.ServiceDTOVersion2BusBusStation
}

func (o *CityBusAPIStationOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Station/City/{City}][%d] cityBusApiStationOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIStationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
