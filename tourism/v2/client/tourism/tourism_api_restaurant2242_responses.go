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

// TourismAPIRestaurant2242Reader is a Reader for the TourismAPIRestaurant2242 structure.
type TourismAPIRestaurant2242Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TourismAPIRestaurant2242Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTourismAPIRestaurant2242OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTourismAPIRestaurant2242Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTourismAPIRestaurant2242NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTourismAPIRestaurant2242OK creates a TourismAPIRestaurant2242OK with default headers values
func NewTourismAPIRestaurant2242OK() *TourismAPIRestaurant2242OK {
	return &TourismAPIRestaurant2242OK{}
}

/* TourismAPIRestaurant2242OK describes a response with status code 200, with default header values.

Success
*/
type TourismAPIRestaurant2242OK struct {
	Payload []*models.PTXServiceDTOTourismSpecificationV2RestaurantTourismInfo
}

func (o *TourismAPIRestaurant2242OK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Restaurant][%d] tourismApiRestaurant2242OK  %+v", 200, o.Payload)
}
func (o *TourismAPIRestaurant2242OK) GetPayload() []*models.PTXServiceDTOTourismSpecificationV2RestaurantTourismInfo {
	return o.Payload
}

func (o *TourismAPIRestaurant2242OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTourismAPIRestaurant2242Status299 creates a TourismAPIRestaurant2242Status299 with default headers values
func NewTourismAPIRestaurant2242Status299() *TourismAPIRestaurant2242Status299 {
	return &TourismAPIRestaurant2242Status299{}
}

/* TourismAPIRestaurant2242Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type TourismAPIRestaurant2242Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *TourismAPIRestaurant2242Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Restaurant][%d] tourismApiRestaurant2242Status299  %+v", 299, o.Payload)
}
func (o *TourismAPIRestaurant2242Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *TourismAPIRestaurant2242Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTourismAPIRestaurant2242NotModified creates a TourismAPIRestaurant2242NotModified with default headers values
func NewTourismAPIRestaurant2242NotModified() *TourismAPIRestaurant2242NotModified {
	return &TourismAPIRestaurant2242NotModified{}
}

/* TourismAPIRestaurant2242NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type TourismAPIRestaurant2242NotModified struct {
}

func (o *TourismAPIRestaurant2242NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Restaurant][%d] tourismApiRestaurant2242NotModified ", 304)
}

func (o *TourismAPIRestaurant2242NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
