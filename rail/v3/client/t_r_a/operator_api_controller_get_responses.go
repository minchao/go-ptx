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

// OperatorAPIControllerGetReader is a Reader for the OperatorAPIControllerGet structure.
type OperatorAPIControllerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperatorAPIControllerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOperatorAPIControllerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewOperatorAPIControllerGetStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewOperatorAPIControllerGetNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOperatorAPIControllerGetOK creates a OperatorAPIControllerGetOK with default headers values
func NewOperatorAPIControllerGetOK() *OperatorAPIControllerGetOK {
	return &OperatorAPIControllerGetOK{}
}

/* OperatorAPIControllerGetOK describes a response with status code 200, with default header values.

Success
*/
type OperatorAPIControllerGetOK struct {
	Payload *models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAOperator
}

func (o *OperatorAPIControllerGetOK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/Operator][%d] operatorApiControllerGetOK  %+v", 200, o.Payload)
}
func (o *OperatorAPIControllerGetOK) GetPayload() *models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAOperator {
	return o.Payload
}

func (o *OperatorAPIControllerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAOperator)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperatorAPIControllerGetStatus299 creates a OperatorAPIControllerGetStatus299 with default headers values
func NewOperatorAPIControllerGetStatus299() *OperatorAPIControllerGetStatus299 {
	return &OperatorAPIControllerGetStatus299{}
}

/* OperatorAPIControllerGetStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type OperatorAPIControllerGetStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *OperatorAPIControllerGetStatus299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/Operator][%d] operatorApiControllerGetStatus299  %+v", 299, o.Payload)
}
func (o *OperatorAPIControllerGetStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *OperatorAPIControllerGetStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperatorAPIControllerGetNotModified creates a OperatorAPIControllerGetNotModified with default headers values
func NewOperatorAPIControllerGetNotModified() *OperatorAPIControllerGetNotModified {
	return &OperatorAPIControllerGetNotModified{}
}

/* OperatorAPIControllerGetNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type OperatorAPIControllerGetNotModified struct {
}

func (o *OperatorAPIControllerGetNotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/Operator][%d] operatorApiControllerGetNotModified ", 304)
}

func (o *OperatorAPIControllerGetNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
