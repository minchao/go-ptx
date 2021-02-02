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
	case 304:
		result := NewTourismAPIActivity0NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTourismAPIActivity0OK creates a TourismAPIActivity0OK with default headers values
func NewTourismAPIActivity0OK() *TourismAPIActivity0OK {
	return &TourismAPIActivity0OK{}
}

/* TourismAPIActivity0OK describes a response with status code 200, with default header values.

Success
*/
type TourismAPIActivity0OK struct {
	Payload []*models.PTXServiceDTOTourismSpecificationV2ActivityTourismInfo
}

func (o *TourismAPIActivity0OK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Activity/{City}][%d] tourismApiActivity0OK  %+v", 200, o.Payload)
}
func (o *TourismAPIActivity0OK) GetPayload() []*models.PTXServiceDTOTourismSpecificationV2ActivityTourismInfo {
	return o.Payload
}

func (o *TourismAPIActivity0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTourismAPIActivity0NotModified creates a TourismAPIActivity0NotModified with default headers values
func NewTourismAPIActivity0NotModified() *TourismAPIActivity0NotModified {
	return &TourismAPIActivity0NotModified{}
}

/* TourismAPIActivity0NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type TourismAPIActivity0NotModified struct {
}

func (o *TourismAPIActivity0NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Activity/{City}][%d] tourismApiActivity0NotModified ", 304)
}

func (o *TourismAPIActivity0NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
