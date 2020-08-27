// Code generated by go-swagger; DO NOT EDIT.

package tourism

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/tourism/v2/models"
)

// TourismAPIActivity0Reader is a Reader for the TourismAPIActivity0 structure.
type TourismAPIActivity0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TourismAPIActivity0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTourismAPIActivity0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTourismAPIActivity0OK creates a TourismAPIActivity0OK with default headers values
func NewTourismAPIActivity0OK() *TourismAPIActivity0OK {
	return &TourismAPIActivity0OK{}
}

/*TourismAPIActivity0OK handles this case with default header values.

OK
*/
type TourismAPIActivity0OK struct {
	Payload []*models.ServiceDTOVersion2ApplicationActivityTourismInfo
}

func (o *TourismAPIActivity0OK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Activity/{City}][%d] tourismApiActivity0OK  %+v", 200, o.Payload)
}

func (o *TourismAPIActivity0OK) GetPayload() []*models.ServiceDTOVersion2ApplicationActivityTourismInfo {
	return o.Payload
}

func (o *TourismAPIActivity0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
