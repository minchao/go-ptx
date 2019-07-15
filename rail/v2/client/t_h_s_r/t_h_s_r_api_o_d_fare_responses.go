// Code generated by go-swagger; DO NOT EDIT.

package t_h_s_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/rail/v2/models"
)

// THSRAPIODFareReader is a Reader for the THSRAPIODFare structure.
type THSRAPIODFareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIODFareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIODFareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTHSRAPIODFareOK creates a THSRAPIODFareOK with default headers values
func NewTHSRAPIODFareOK() *THSRAPIODFareOK {
	return &THSRAPIODFareOK{}
}

/*THSRAPIODFareOK handles this case with default header values.

OK
*/
type THSRAPIODFareOK struct {
	Payload []*models.ServiceDTOVersion2RailTHSRRailODFare
}

func (o *THSRAPIODFareOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/ODFare][%d] tHSRApiODFareOK  %+v", 200, o.Payload)
}

func (o *THSRAPIODFareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
