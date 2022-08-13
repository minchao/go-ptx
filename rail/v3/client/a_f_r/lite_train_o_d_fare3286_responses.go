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

// LiteTrainODFare3286Reader is a Reader for the LiteTrainODFare3286 structure.
type LiteTrainODFare3286Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LiteTrainODFare3286Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLiteTrainODFare3286OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewLiteTrainODFare3286Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLiteTrainODFare3286OK creates a LiteTrainODFare3286OK with default headers values
func NewLiteTrainODFare3286OK() *LiteTrainODFare3286OK {
	return &LiteTrainODFare3286OK{}
}

/* LiteTrainODFare3286OK describes a response with status code 200, with default header values.

Success
*/
type LiteTrainODFare3286OK struct {
	Payload *models.PTXAPIRailModelLiteTrainLiteTrainODFareWrapperPTXServiceDTORailSpecificationV3LiteTrainODFareODFare
}

func (o *LiteTrainODFare3286OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/ODFare][%d] liteTrainODFare3286OK  %+v", 200, o.Payload)
}
func (o *LiteTrainODFare3286OK) GetPayload() *models.PTXAPIRailModelLiteTrainLiteTrainODFareWrapperPTXServiceDTORailSpecificationV3LiteTrainODFareODFare {
	return o.Payload
}

func (o *LiteTrainODFare3286OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelLiteTrainLiteTrainODFareWrapperPTXServiceDTORailSpecificationV3LiteTrainODFareODFare)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLiteTrainODFare3286Status299 creates a LiteTrainODFare3286Status299 with default headers values
func NewLiteTrainODFare3286Status299() *LiteTrainODFare3286Status299 {
	return &LiteTrainODFare3286Status299{}
}

/* LiteTrainODFare3286Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type LiteTrainODFare3286Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *LiteTrainODFare3286Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/ODFare][%d] liteTrainODFare3286Status299  %+v", 299, o.Payload)
}
func (o *LiteTrainODFare3286Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *LiteTrainODFare3286Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
