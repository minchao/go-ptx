// Code generated by go-swagger; DO NOT EDIT.

package air

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/air/v2/models"
)

// AirAPIAirlineReader is a Reader for the AirAPIAirline structure.
type AirAPIAirlineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AirAPIAirlineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAirAPIAirlineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAirAPIAirlineOK creates a AirAPIAirlineOK with default headers values
func NewAirAPIAirlineOK() *AirAPIAirlineOK {
	return &AirAPIAirlineOK{}
}

/*AirAPIAirlineOK handles this case with default header values.

OK
*/
type AirAPIAirlineOK struct {
	Payload *models.ServiceDTOVersion2AviationAirline
}

func (o *AirAPIAirlineOK) Error() string {
	return fmt.Sprintf("[GET /v2/Air/Airline/{IATA}][%d] airApiAirlineOK  %+v", 200, o.Payload)
}

func (o *AirAPIAirlineOK) GetPayload() *models.ServiceDTOVersion2AviationAirline {
	return o.Payload
}

func (o *AirAPIAirlineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion2AviationAirline)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
