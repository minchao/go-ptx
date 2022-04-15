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

// TourismAPIRestaurant02243Reader is a Reader for the TourismAPIRestaurant02243 structure.
type TourismAPIRestaurant02243Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TourismAPIRestaurant02243Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTourismAPIRestaurant02243OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTourismAPIRestaurant02243Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTourismAPIRestaurant02243NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTourismAPIRestaurant02243OK creates a TourismAPIRestaurant02243OK with default headers values
func NewTourismAPIRestaurant02243OK() *TourismAPIRestaurant02243OK {
	return &TourismAPIRestaurant02243OK{}
}

/* TourismAPIRestaurant02243OK describes a response with status code 200, with default header values.

Success
*/
type TourismAPIRestaurant02243OK struct {
	Payload []*models.PTXServiceDTOTourismSpecificationV2RestaurantTourismInfo
}

func (o *TourismAPIRestaurant02243OK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Restaurant/{City}][%d] tourismApiRestaurant02243OK  %+v", 200, o.Payload)
}
func (o *TourismAPIRestaurant02243OK) GetPayload() []*models.PTXServiceDTOTourismSpecificationV2RestaurantTourismInfo {
	return o.Payload
}

func (o *TourismAPIRestaurant02243OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTourismAPIRestaurant02243Status299 creates a TourismAPIRestaurant02243Status299 with default headers values
func NewTourismAPIRestaurant02243Status299() *TourismAPIRestaurant02243Status299 {
	return &TourismAPIRestaurant02243Status299{}
}

/* TourismAPIRestaurant02243Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type TourismAPIRestaurant02243Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *TourismAPIRestaurant02243Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Restaurant/{City}][%d] tourismApiRestaurant02243Status299  %+v", 299, o.Payload)
}
func (o *TourismAPIRestaurant02243Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *TourismAPIRestaurant02243Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTourismAPIRestaurant02243NotModified creates a TourismAPIRestaurant02243NotModified with default headers values
func NewTourismAPIRestaurant02243NotModified() *TourismAPIRestaurant02243NotModified {
	return &TourismAPIRestaurant02243NotModified{}
}

/* TourismAPIRestaurant02243NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type TourismAPIRestaurant02243NotModified struct {
}

func (o *TourismAPIRestaurant02243NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Restaurant/{City}][%d] tourismApiRestaurant02243NotModified ", 304)
}

func (o *TourismAPIRestaurant02243NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}