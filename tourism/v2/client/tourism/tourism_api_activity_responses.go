// Code generated by go-swagger; DO NOT EDIT.

package tourism

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/tourism/v2/models"
)

// TourismAPIActivityReader is a Reader for the TourismAPIActivity structure.
type TourismAPIActivityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TourismAPIActivityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTourismAPIActivityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTourismAPIActivityOK creates a TourismAPIActivityOK with default headers values
func NewTourismAPIActivityOK() *TourismAPIActivityOK {
	return &TourismAPIActivityOK{}
}

/*TourismAPIActivityOK handles this case with default header values.

OK
*/
type TourismAPIActivityOK struct {
	Payload []*models.ServiceDTOVersion2ApplicationActivityTourismInfo
}

func (o *TourismAPIActivityOK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Activity][%d] tourismApiActivityOK  %+v", 200, o.Payload)
}

func (o *TourismAPIActivityOK) GetPayload() []*models.ServiceDTOVersion2ApplicationActivityTourismInfo {
	return o.Payload
}

func (o *TourismAPIActivityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
