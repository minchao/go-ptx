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

// TourismAPIHotelReader is a Reader for the TourismAPIHotel structure.
type TourismAPIHotelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TourismAPIHotelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTourismAPIHotelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTourismAPIHotelNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTourismAPIHotelOK creates a TourismAPIHotelOK with default headers values
func NewTourismAPIHotelOK() *TourismAPIHotelOK {
	return &TourismAPIHotelOK{}
}

/* TourismAPIHotelOK describes a response with status code 200, with default header values.

Success
*/
type TourismAPIHotelOK struct {
	Payload []*models.PTXServiceDTOTourismSpecificationV2HotelTourismInfo
}

func (o *TourismAPIHotelOK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Hotel][%d] tourismApiHotelOK  %+v", 200, o.Payload)
}
func (o *TourismAPIHotelOK) GetPayload() []*models.PTXServiceDTOTourismSpecificationV2HotelTourismInfo {
	return o.Payload
}

func (o *TourismAPIHotelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTourismAPIHotelNotModified creates a TourismAPIHotelNotModified with default headers values
func NewTourismAPIHotelNotModified() *TourismAPIHotelNotModified {
	return &TourismAPIHotelNotModified{}
}

/* TourismAPIHotelNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type TourismAPIHotelNotModified struct {
}

func (o *TourismAPIHotelNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Hotel][%d] tourismApiHotelNotModified ", 304)
}

func (o *TourismAPIHotelNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
