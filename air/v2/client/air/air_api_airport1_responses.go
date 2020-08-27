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

// AirAPIAirport1Reader is a Reader for the AirAPIAirport1 structure.
type AirAPIAirport1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AirAPIAirport1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAirAPIAirport1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAirAPIAirport1OK creates a AirAPIAirport1OK with default headers values
func NewAirAPIAirport1OK() *AirAPIAirport1OK {
	return &AirAPIAirport1OK{}
}

/*AirAPIAirport1OK handles this case with default header values.

Success
*/
type AirAPIAirport1OK struct {
	Payload *models.PTXServiceDTOAirSpecificationV2Airport
}

func (o *AirAPIAirport1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Air/Airport/{IATA}][%d] airApiAirport1OK  %+v", 200, o.Payload)
}

func (o *AirAPIAirport1OK) GetPayload() *models.PTXServiceDTOAirSpecificationV2Airport {
	return o.Payload
}

func (o *AirAPIAirport1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOAirSpecificationV2Airport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
