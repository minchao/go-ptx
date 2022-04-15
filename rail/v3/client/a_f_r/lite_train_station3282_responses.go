// Code generated by go-swagger; DO NOT EDIT.

package a_f_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v3/models"
)

// LiteTrainStation3282Reader is a Reader for the LiteTrainStation3282 structure.
type LiteTrainStation3282Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LiteTrainStation3282Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLiteTrainStation3282OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewLiteTrainStation3282Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewLiteTrainStation3282NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLiteTrainStation3282OK creates a LiteTrainStation3282OK with default headers values
func NewLiteTrainStation3282OK() *LiteTrainStation3282OK {
	return &LiteTrainStation3282OK{}
}

/* LiteTrainStation3282OK describes a response with status code 200, with default header values.

Success
*/
type LiteTrainStation3282OK struct {
	Payload *models.PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainStation
}

func (o *LiteTrainStation3282OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/Station][%d] liteTrainStation3282OK  %+v", 200, o.Payload)
}
func (o *LiteTrainStation3282OK) GetPayload() *models.PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainStation {
	return o.Payload
}

func (o *LiteTrainStation3282OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainStation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLiteTrainStation3282Status299 creates a LiteTrainStation3282Status299 with default headers values
func NewLiteTrainStation3282Status299() *LiteTrainStation3282Status299 {
	return &LiteTrainStation3282Status299{}
}

/* LiteTrainStation3282Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type LiteTrainStation3282Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *LiteTrainStation3282Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/Station][%d] liteTrainStation3282Status299  %+v", 299, o.Payload)
}
func (o *LiteTrainStation3282Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *LiteTrainStation3282Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLiteTrainStation3282NotModified creates a LiteTrainStation3282NotModified with default headers values
func NewLiteTrainStation3282NotModified() *LiteTrainStation3282NotModified {
	return &LiteTrainStation3282NotModified{}
}

/* LiteTrainStation3282NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type LiteTrainStation3282NotModified struct {
}

func (o *LiteTrainStation3282NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/Station][%d] liteTrainStation3282NotModified ", 304)
}

func (o *LiteTrainStation3282NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}