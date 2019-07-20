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

// CityBusAPIStationNameReader is a Reader for the CityBusAPIStationName structure.
type CityBusAPIStationNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIStationNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIStationNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCityBusAPIStationNameOK creates a CityBusAPIStationNameOK with default headers values
func NewCityBusAPIStationNameOK() *CityBusAPIStationNameOK {
	return &CityBusAPIStationNameOK{}
}

/*CityBusAPIStationNameOK handles this case with default header values.

OK
*/
type CityBusAPIStationNameOK struct {
	Payload []*models.ServiceDTOVersion2BusBusStationName
}

func (o *CityBusAPIStationNameOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/StationName/City/{City}][%d] cityBusApiStationNameOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIStationNameOK) GetPayload() []*models.ServiceDTOVersion2BusBusStationName {
	return o.Payload
}

func (o *CityBusAPIStationNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
