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

// LiteTrainLine3283Reader is a Reader for the LiteTrainLine3283 structure.
type LiteTrainLine3283Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LiteTrainLine3283Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLiteTrainLine3283OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewLiteTrainLine3283Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLiteTrainLine3283OK creates a LiteTrainLine3283OK with default headers values
func NewLiteTrainLine3283OK() *LiteTrainLine3283OK {
	return &LiteTrainLine3283OK{}
}

/* LiteTrainLine3283OK describes a response with status code 200, with default header values.

Success
*/
type LiteTrainLine3283OK struct {
	Payload *models.PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainLine
}

func (o *LiteTrainLine3283OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/Line][%d] liteTrainLine3283OK  %+v", 200, o.Payload)
}
func (o *LiteTrainLine3283OK) GetPayload() *models.PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainLine {
	return o.Payload
}

func (o *LiteTrainLine3283OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainLine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLiteTrainLine3283Status299 creates a LiteTrainLine3283Status299 with default headers values
func NewLiteTrainLine3283Status299() *LiteTrainLine3283Status299 {
	return &LiteTrainLine3283Status299{}
}

/* LiteTrainLine3283Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type LiteTrainLine3283Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *LiteTrainLine3283Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/Line][%d] liteTrainLine3283Status299  %+v", 299, o.Payload)
}
func (o *LiteTrainLine3283Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *LiteTrainLine3283Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
