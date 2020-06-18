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

// AirAPIFIDSReader is a Reader for the AirAPIFIDS structure.
type AirAPIFIDSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AirAPIFIDSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAirAPIFIDSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAirAPIFIDSOK creates a AirAPIFIDSOK with default headers values
func NewAirAPIFIDSOK() *AirAPIFIDSOK {
	return &AirAPIFIDSOK{}
}

/*AirAPIFIDSOK handles this case with default header values.

OK
*/
type AirAPIFIDSOK struct {
	Payload []*models.ServiceDTOVersion2AviationAirportFIDS
}

func (o *AirAPIFIDSOK) Error() string {
	return fmt.Sprintf("[GET /v2/Air/FIDS/Airport][%d] airApiFIdSOK  %+v", 200, o.Payload)
}

func (o *AirAPIFIDSOK) GetPayload() []*models.ServiceDTOVersion2AviationAirportFIDS {
	return o.Payload
}

func (o *AirAPIFIDSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
