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
	case 304:
		result := NewTourismAPIHotel0NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTourismAPIHotel0OK creates a TourismAPIHotel0OK with default headers values
func NewTourismAPIHotel0OK() *TourismAPIHotel0OK {
	return &TourismAPIHotel0OK{}
}

/* TourismAPIHotel0OK describes a response with status code 200, with default header values.

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

// NewTourismAPIHotel0NotModified creates a TourismAPIHotel0NotModified with default headers values
func NewTourismAPIHotel0NotModified() *TourismAPIHotel0NotModified {
	return &TourismAPIHotel0NotModified{}
}

/* TourismAPIHotel0NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type TourismAPIHotel0NotModified struct {
}

func (o *TourismAPIHotel0NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Hotel/{City}][%d] tourismApiHotel0NotModified ", 304)
}

func (o *TourismAPIHotel0NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
