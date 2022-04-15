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

// LiteTrainGeneralTrainTimetable3285Reader is a Reader for the LiteTrainGeneralTrainTimetable3285 structure.
type LiteTrainGeneralTrainTimetable3285Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LiteTrainGeneralTrainTimetable3285Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLiteTrainGeneralTrainTimetable3285OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewLiteTrainGeneralTrainTimetable3285Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewLiteTrainGeneralTrainTimetable3285NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLiteTrainGeneralTrainTimetable3285OK creates a LiteTrainGeneralTrainTimetable3285OK with default headers values
func NewLiteTrainGeneralTrainTimetable3285OK() *LiteTrainGeneralTrainTimetable3285OK {
	return &LiteTrainGeneralTrainTimetable3285OK{}
}

/* LiteTrainGeneralTrainTimetable3285OK describes a response with status code 200, with default header values.

Success
*/
type LiteTrainGeneralTrainTimetable3285OK struct {
	Payload *models.PTXAPIRailModelLiteTrainLiteTrainGeneralTrainWrapperPTXServiceDTORailSpecificationV3LiteTrainGeneralTrainTimetable
}

func (o *LiteTrainGeneralTrainTimetable3285OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/GeneralTrainTimetable][%d] liteTrainGeneralTrainTimetable3285OK  %+v", 200, o.Payload)
}
func (o *LiteTrainGeneralTrainTimetable3285OK) GetPayload() *models.PTXAPIRailModelLiteTrainLiteTrainGeneralTrainWrapperPTXServiceDTORailSpecificationV3LiteTrainGeneralTrainTimetable {
	return o.Payload
}

func (o *LiteTrainGeneralTrainTimetable3285OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelLiteTrainLiteTrainGeneralTrainWrapperPTXServiceDTORailSpecificationV3LiteTrainGeneralTrainTimetable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLiteTrainGeneralTrainTimetable3285Status299 creates a LiteTrainGeneralTrainTimetable3285Status299 with default headers values
func NewLiteTrainGeneralTrainTimetable3285Status299() *LiteTrainGeneralTrainTimetable3285Status299 {
	return &LiteTrainGeneralTrainTimetable3285Status299{}
}

/* LiteTrainGeneralTrainTimetable3285Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type LiteTrainGeneralTrainTimetable3285Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *LiteTrainGeneralTrainTimetable3285Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/GeneralTrainTimetable][%d] liteTrainGeneralTrainTimetable3285Status299  %+v", 299, o.Payload)
}
func (o *LiteTrainGeneralTrainTimetable3285Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *LiteTrainGeneralTrainTimetable3285Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLiteTrainGeneralTrainTimetable3285NotModified creates a LiteTrainGeneralTrainTimetable3285NotModified with default headers values
func NewLiteTrainGeneralTrainTimetable3285NotModified() *LiteTrainGeneralTrainTimetable3285NotModified {
	return &LiteTrainGeneralTrainTimetable3285NotModified{}
}

/* LiteTrainGeneralTrainTimetable3285NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type LiteTrainGeneralTrainTimetable3285NotModified struct {
}

func (o *LiteTrainGeneralTrainTimetable3285NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/GeneralTrainTimetable][%d] liteTrainGeneralTrainTimetable3285NotModified ", 304)
}

func (o *LiteTrainGeneralTrainTimetable3285NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}