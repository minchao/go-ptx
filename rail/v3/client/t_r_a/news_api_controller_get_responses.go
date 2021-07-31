// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v3/models"
)

// NewsAPIControllerGetReader is a Reader for the NewsAPIControllerGet structure.
type NewsAPIControllerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NewsAPIControllerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNewsAPIControllerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewNewsAPIControllerGetStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewNewsAPIControllerGetNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNewsAPIControllerGetOK creates a NewsAPIControllerGetOK with default headers values
func NewNewsAPIControllerGetOK() *NewsAPIControllerGetOK {
	return &NewsAPIControllerGetOK{}
}

/* NewsAPIControllerGetOK describes a response with status code 200, with default header values.

Success
*/
type NewsAPIControllerGetOK struct {
	Payload *models.PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRANewsListNews
}

func (o *NewsAPIControllerGetOK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/News][%d] newsApiControllerGetOK  %+v", 200, o.Payload)
}
func (o *NewsAPIControllerGetOK) GetPayload() *models.PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRANewsListNews {
	return o.Payload
}

func (o *NewsAPIControllerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRANewsListNews)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNewsAPIControllerGetStatus299 creates a NewsAPIControllerGetStatus299 with default headers values
func NewNewsAPIControllerGetStatus299() *NewsAPIControllerGetStatus299 {
	return &NewsAPIControllerGetStatus299{}
}

/* NewsAPIControllerGetStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type NewsAPIControllerGetStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *NewsAPIControllerGetStatus299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/News][%d] newsApiControllerGetStatus299  %+v", 299, o.Payload)
}
func (o *NewsAPIControllerGetStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *NewsAPIControllerGetStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNewsAPIControllerGetNotModified creates a NewsAPIControllerGetNotModified with default headers values
func NewNewsAPIControllerGetNotModified() *NewsAPIControllerGetNotModified {
	return &NewsAPIControllerGetNotModified{}
}

/* NewsAPIControllerGetNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type NewsAPIControllerGetNotModified struct {
}

func (o *NewsAPIControllerGetNotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/News][%d] newsApiControllerGetNotModified ", 304)
}

func (o *NewsAPIControllerGetNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
