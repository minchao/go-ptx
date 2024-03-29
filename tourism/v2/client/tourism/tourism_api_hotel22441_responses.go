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

// TourismAPIHotel22441Reader is a Reader for the TourismAPIHotel22441 structure.
type TourismAPIHotel22441Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TourismAPIHotel22441Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTourismAPIHotel22441OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTourismAPIHotel22441Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTourismAPIHotel22441NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTourismAPIHotel22441OK creates a TourismAPIHotel22441OK with default headers values
func NewTourismAPIHotel22441OK() *TourismAPIHotel22441OK {
	return &TourismAPIHotel22441OK{}
}

/* TourismAPIHotel22441OK describes a response with status code 200, with default header values.

Success
*/
type TourismAPIHotel22441OK struct {
	Payload []*models.PTXServiceDTOTourismSpecificationV2HotelTourismInfo
}

func (o *TourismAPIHotel22441OK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Hotel/{City}][%d] tourismApiHotel22441OK  %+v", 200, o.Payload)
}
func (o *TourismAPIHotel22441OK) GetPayload() []*models.PTXServiceDTOTourismSpecificationV2HotelTourismInfo {
	return o.Payload
}

func (o *TourismAPIHotel22441OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTourismAPIHotel22441Status299 creates a TourismAPIHotel22441Status299 with default headers values
func NewTourismAPIHotel22441Status299() *TourismAPIHotel22441Status299 {
	return &TourismAPIHotel22441Status299{}
}

/* TourismAPIHotel22441Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type TourismAPIHotel22441Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *TourismAPIHotel22441Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Hotel/{City}][%d] tourismApiHotel22441Status299  %+v", 299, o.Payload)
}
func (o *TourismAPIHotel22441Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *TourismAPIHotel22441Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTourismAPIHotel22441NotModified creates a TourismAPIHotel22441NotModified with default headers values
func NewTourismAPIHotel22441NotModified() *TourismAPIHotel22441NotModified {
	return &TourismAPIHotel22441NotModified{}
}

/* TourismAPIHotel22441NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type TourismAPIHotel22441NotModified struct {
}

func (o *TourismAPIHotel22441NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Hotel/{City}][%d] tourismApiHotel22441NotModified ", 304)
}

func (o *TourismAPIHotel22441NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
