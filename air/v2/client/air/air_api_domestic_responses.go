// Code generated by go-swagger; DO NOT EDIT.

package air

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/air/v2/models"
)

// AirAPIDomesticReader is a Reader for the AirAPIDomestic structure.
type AirAPIDomesticReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AirAPIDomesticReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAirAPIDomesticOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAirAPIDomesticOK creates a AirAPIDomesticOK with default headers values
func NewAirAPIDomesticOK() *AirAPIDomesticOK {
	return &AirAPIDomesticOK{}
}

/*AirAPIDomesticOK handles this case with default header values.

Success
*/
type AirAPIDomesticOK struct {
	Payload []*models.PTXServiceDTOAirSpecificationV2GeneralFlightSchedule
}

func (o *AirAPIDomesticOK) Error() string {
	return fmt.Sprintf("[GET /v2/Air/GeneralSchedule/Domestic][%d] airApiDomesticOK  %+v", 200, o.Payload)
}

func (o *AirAPIDomesticOK) GetPayload() []*models.PTXServiceDTOAirSpecificationV2GeneralFlightSchedule {
	return o.Payload
}

func (o *AirAPIDomesticOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
