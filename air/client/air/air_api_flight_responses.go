// Code generated by go-swagger; DO NOT EDIT.

package air

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/air/models"
)

// AirAPIFlightReader is a Reader for the AirAPIFlight structure.
type AirAPIFlightReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AirAPIFlightReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAirAPIFlightOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAirAPIFlightOK creates a AirAPIFlightOK with default headers values
func NewAirAPIFlightOK() *AirAPIFlightOK {
	return &AirAPIFlightOK{}
}

/*AirAPIFlightOK handles this case with default header values.

OK
*/
type AirAPIFlightOK struct {
	Payload []*models.ServiceDTOVersion2AviationFIDS
}

func (o *AirAPIFlightOK) Error() string {
	return fmt.Sprintf("[GET /v2/Air/FIDS/Flight][%d] airApiFlightOK  %+v", 200, o.Payload)
}

func (o *AirAPIFlightOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
