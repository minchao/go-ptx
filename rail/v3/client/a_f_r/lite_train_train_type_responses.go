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

// LiteTrainTrainTypeReader is a Reader for the LiteTrainTrainType structure.
type LiteTrainTrainTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LiteTrainTrainTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLiteTrainTrainTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewLiteTrainTrainTypeStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewLiteTrainTrainTypeNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLiteTrainTrainTypeOK creates a LiteTrainTrainTypeOK with default headers values
func NewLiteTrainTrainTypeOK() *LiteTrainTrainTypeOK {
	return &LiteTrainTrainTypeOK{}
}

/* LiteTrainTrainTypeOK describes a response with status code 200, with default header values.

Success
*/
type LiteTrainTrainTypeOK struct {
	Payload *models.PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainTrainType
}

func (o *LiteTrainTrainTypeOK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/TrainType][%d] liteTrainTrainTypeOK  %+v", 200, o.Payload)
}
func (o *LiteTrainTrainTypeOK) GetPayload() *models.PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainTrainType {
	return o.Payload
}

func (o *LiteTrainTrainTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainTrainType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLiteTrainTrainTypeStatus299 creates a LiteTrainTrainTypeStatus299 with default headers values
func NewLiteTrainTrainTypeStatus299() *LiteTrainTrainTypeStatus299 {
	return &LiteTrainTrainTypeStatus299{}
}

/* LiteTrainTrainTypeStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type LiteTrainTrainTypeStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *LiteTrainTrainTypeStatus299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/TrainType][%d] liteTrainTrainTypeStatus299  %+v", 299, o.Payload)
}
func (o *LiteTrainTrainTypeStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *LiteTrainTrainTypeStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLiteTrainTrainTypeNotModified creates a LiteTrainTrainTypeNotModified with default headers values
func NewLiteTrainTrainTypeNotModified() *LiteTrainTrainTypeNotModified {
	return &LiteTrainTrainTypeNotModified{}
}

/* LiteTrainTrainTypeNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type LiteTrainTrainTypeNotModified struct {
}

func (o *LiteTrainTrainTypeNotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/TrainType][%d] liteTrainTrainTypeNotModified ", 304)
}

func (o *LiteTrainTrainTypeNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
