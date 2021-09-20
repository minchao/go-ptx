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
	case 299:
		result := NewTourismAPIActivityStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTourismAPIActivityNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTourismAPIActivityOK creates a TourismAPIActivityOK with default headers values
func NewTourismAPIActivityOK() *TourismAPIActivityOK {
	return &TourismAPIActivityOK{}
}

/* TourismAPIActivityOK describes a response with status code 200, with default header values.

Success
*/
type TourismAPIActivityOK struct {
	Payload []*models.PTXServiceDTOTourismSpecificationV2ActivityTourismInfo
}

func (o *TourismAPIActivityOK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Activity][%d] tourismApiActivityOK  %+v", 200, o.Payload)
}
func (o *TourismAPIActivityOK) GetPayload() []*models.PTXServiceDTOTourismSpecificationV2ActivityTourismInfo {
	return o.Payload
}

func (o *TourismAPIActivityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTourismAPIActivityStatus299 creates a TourismAPIActivityStatus299 with default headers values
func NewTourismAPIActivityStatus299() *TourismAPIActivityStatus299 {
	return &TourismAPIActivityStatus299{}
}

/* TourismAPIActivityStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type TourismAPIActivityStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *TourismAPIActivityStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Activity][%d] tourismApiActivityStatus299  %+v", 299, o.Payload)
}
func (o *TourismAPIActivityStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *TourismAPIActivityStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTourismAPIActivityNotModified creates a TourismAPIActivityNotModified with default headers values
func NewTourismAPIActivityNotModified() *TourismAPIActivityNotModified {
	return &TourismAPIActivityNotModified{}
}

/* TourismAPIActivityNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type TourismAPIActivityNotModified struct {
}

func (o *TourismAPIActivityNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Activity][%d] tourismApiActivityNotModified ", 304)
}

func (o *TourismAPIActivityNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
