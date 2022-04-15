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

// LiteTrainNews3291Reader is a Reader for the LiteTrainNews3291 structure.
type LiteTrainNews3291Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LiteTrainNews3291Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLiteTrainNews3291OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewLiteTrainNews3291Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewLiteTrainNews3291NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLiteTrainNews3291OK creates a LiteTrainNews3291OK with default headers values
func NewLiteTrainNews3291OK() *LiteTrainNews3291OK {
	return &LiteTrainNews3291OK{}
}

/* LiteTrainNews3291OK describes a response with status code 200, with default header values.

Success
*/
type LiteTrainNews3291OK struct {
	Payload *models.PTXAPIRailModelLiteTrainLiteTrainRealTimeWrapperPTXServiceDTORailSpecificationV3TRATRANewsListNews
}

func (o *LiteTrainNews3291OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/News][%d] liteTrainNews3291OK  %+v", 200, o.Payload)
}
func (o *LiteTrainNews3291OK) GetPayload() *models.PTXAPIRailModelLiteTrainLiteTrainRealTimeWrapperPTXServiceDTORailSpecificationV3TRATRANewsListNews {
	return o.Payload
}

func (o *LiteTrainNews3291OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelLiteTrainLiteTrainRealTimeWrapperPTXServiceDTORailSpecificationV3TRATRANewsListNews)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLiteTrainNews3291Status299 creates a LiteTrainNews3291Status299 with default headers values
func NewLiteTrainNews3291Status299() *LiteTrainNews3291Status299 {
	return &LiteTrainNews3291Status299{}
}

/* LiteTrainNews3291Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type LiteTrainNews3291Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *LiteTrainNews3291Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/News][%d] liteTrainNews3291Status299  %+v", 299, o.Payload)
}
func (o *LiteTrainNews3291Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *LiteTrainNews3291Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLiteTrainNews3291NotModified creates a LiteTrainNews3291NotModified with default headers values
func NewLiteTrainNews3291NotModified() *LiteTrainNews3291NotModified {
	return &LiteTrainNews3291NotModified{}
}

/* LiteTrainNews3291NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type LiteTrainNews3291NotModified struct {
}

func (o *LiteTrainNews3291NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/News][%d] liteTrainNews3291NotModified ", 304)
}

func (o *LiteTrainNews3291NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}