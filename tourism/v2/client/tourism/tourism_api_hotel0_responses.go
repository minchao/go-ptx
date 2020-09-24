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

// TourismAPIHotel0Reader is a Reader for the TourismAPIHotel0 structure.
type TourismAPIHotel0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TourismAPIHotel0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTourismAPIHotel0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTourismAPIHotel0OK creates a TourismAPIHotel0OK with default headers values
func NewTourismAPIHotel0OK() *TourismAPIHotel0OK {
	return &TourismAPIHotel0OK{}
}

/*TourismAPIHotel0OK handles this case with default header values.

Success
*/
type TourismAPIHotel0OK struct {
	Payload []*models.PTXServiceDTOTourismSpecificationV2HotelTourismInfo
}

func (o *TourismAPIHotel0OK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Hotel/{City}][%d] tourismApiHotel0OK  %+v", 200, o.Payload)
}

func (o *TourismAPIHotel0OK) GetPayload() []*models.PTXServiceDTOTourismSpecificationV2HotelTourismInfo {
	return o.Payload
}

func (o *TourismAPIHotel0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
