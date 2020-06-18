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

// AirAPIAirline1Reader is a Reader for the AirAPIAirline1 structure.
type AirAPIAirline1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AirAPIAirline1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAirAPIAirline1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAirAPIAirline1OK creates a AirAPIAirline1OK with default headers values
func NewAirAPIAirline1OK() *AirAPIAirline1OK {
	return &AirAPIAirline1OK{}
}

/*AirAPIAirline1OK handles this case with default header values.

OK
*/
type AirAPIAirline1OK struct {
	Payload []*models.ServiceDTOVersion2AviationAirline
}

func (o *AirAPIAirline1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Air/Airline][%d] airApiAirline1OK  %+v", 200, o.Payload)
}

func (o *AirAPIAirline1OK) GetPayload() []*models.ServiceDTOVersion2AviationAirline {
	return o.Payload
}

func (o *AirAPIAirline1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
