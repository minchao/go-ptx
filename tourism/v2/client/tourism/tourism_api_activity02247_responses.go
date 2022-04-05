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

// TourismAPIActivity02247Reader is a Reader for the TourismAPIActivity02247 structure.
type TourismAPIActivity02247Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TourismAPIActivity02247Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTourismAPIActivity02247OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTourismAPIActivity02247Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTourismAPIActivity02247NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTourismAPIActivity02247OK creates a TourismAPIActivity02247OK with default headers values
func NewTourismAPIActivity02247OK() *TourismAPIActivity02247OK {
	return &TourismAPIActivity02247OK{}
}

/* TourismAPIActivity02247OK describes a response with status code 200, with default header values.

Success
*/
type TourismAPIActivity02247OK struct {
	Payload []*models.PTXServiceDTOTourismSpecificationV2ActivityTourismInfo
}

func (o *TourismAPIActivity02247OK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Activity/{City}][%d] tourismApiActivity02247OK  %+v", 200, o.Payload)
}
func (o *TourismAPIActivity02247OK) GetPayload() []*models.PTXServiceDTOTourismSpecificationV2ActivityTourismInfo {
	return o.Payload
}

func (o *TourismAPIActivity02247OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTourismAPIActivity02247Status299 creates a TourismAPIActivity02247Status299 with default headers values
func NewTourismAPIActivity02247Status299() *TourismAPIActivity02247Status299 {
	return &TourismAPIActivity02247Status299{}
}

/* TourismAPIActivity02247Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type TourismAPIActivity02247Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *TourismAPIActivity02247Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Activity/{City}][%d] tourismApiActivity02247Status299  %+v", 299, o.Payload)
}
func (o *TourismAPIActivity02247Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *TourismAPIActivity02247Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTourismAPIActivity02247NotModified creates a TourismAPIActivity02247NotModified with default headers values
func NewTourismAPIActivity02247NotModified() *TourismAPIActivity02247NotModified {
	return &TourismAPIActivity02247NotModified{}
}

/* TourismAPIActivity02247NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type TourismAPIActivity02247NotModified struct {
}

func (o *TourismAPIActivity02247NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Activity/{City}][%d] tourismApiActivity02247NotModified ", 304)
}

func (o *TourismAPIActivity02247NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
