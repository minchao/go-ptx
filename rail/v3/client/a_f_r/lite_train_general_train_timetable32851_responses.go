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

// LiteTrainGeneralTrainTimetable32851Reader is a Reader for the LiteTrainGeneralTrainTimetable32851 structure.
type LiteTrainGeneralTrainTimetable32851Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LiteTrainGeneralTrainTimetable32851Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLiteTrainGeneralTrainTimetable32851OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewLiteTrainGeneralTrainTimetable32851Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLiteTrainGeneralTrainTimetable32851OK creates a LiteTrainGeneralTrainTimetable32851OK with default headers values
func NewLiteTrainGeneralTrainTimetable32851OK() *LiteTrainGeneralTrainTimetable32851OK {
	return &LiteTrainGeneralTrainTimetable32851OK{}
}

/* LiteTrainGeneralTrainTimetable32851OK describes a response with status code 200, with default header values.

Success
*/
type LiteTrainGeneralTrainTimetable32851OK struct {
	Payload *models.PTXAPIRailModelLiteTrainLiteTrainGeneralTrainWrapperPTXServiceDTORailSpecificationV3LiteTrainGeneralTrainTimetable
}

func (o *LiteTrainGeneralTrainTimetable32851OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/GeneralTrainTimetable/TrainNo/{TrainNo}][%d] liteTrainGeneralTrainTimetable32851OK  %+v", 200, o.Payload)
}
func (o *LiteTrainGeneralTrainTimetable32851OK) GetPayload() *models.PTXAPIRailModelLiteTrainLiteTrainGeneralTrainWrapperPTXServiceDTORailSpecificationV3LiteTrainGeneralTrainTimetable {
	return o.Payload
}

func (o *LiteTrainGeneralTrainTimetable32851OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelLiteTrainLiteTrainGeneralTrainWrapperPTXServiceDTORailSpecificationV3LiteTrainGeneralTrainTimetable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLiteTrainGeneralTrainTimetable32851Status299 creates a LiteTrainGeneralTrainTimetable32851Status299 with default headers values
func NewLiteTrainGeneralTrainTimetable32851Status299() *LiteTrainGeneralTrainTimetable32851Status299 {
	return &LiteTrainGeneralTrainTimetable32851Status299{}
}

/* LiteTrainGeneralTrainTimetable32851Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type LiteTrainGeneralTrainTimetable32851Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *LiteTrainGeneralTrainTimetable32851Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/GeneralTrainTimetable/TrainNo/{TrainNo}][%d] liteTrainGeneralTrainTimetable32851Status299  %+v", 299, o.Payload)
}
func (o *LiteTrainGeneralTrainTimetable32851Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *LiteTrainGeneralTrainTimetable32851Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
