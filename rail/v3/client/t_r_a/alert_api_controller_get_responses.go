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

// AlertAPIControllerGetReader is a Reader for the AlertAPIControllerGet structure.
type AlertAPIControllerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertAPIControllerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAlertAPIControllerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewAlertAPIControllerGetStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewAlertAPIControllerGetNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAlertAPIControllerGetOK creates a AlertAPIControllerGetOK with default headers values
func NewAlertAPIControllerGetOK() *AlertAPIControllerGetOK {
	return &AlertAPIControllerGetOK{}
}

/* AlertAPIControllerGetOK describes a response with status code 200, with default header values.

Success
*/
type AlertAPIControllerGetOK struct {
	Payload *models.PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAAlertListAlert
}

func (o *AlertAPIControllerGetOK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/Alert][%d] alertApiControllerGetOK  %+v", 200, o.Payload)
}
func (o *AlertAPIControllerGetOK) GetPayload() *models.PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAAlertListAlert {
	return o.Payload
}

func (o *AlertAPIControllerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAAlertListAlert)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertAPIControllerGetStatus299 creates a AlertAPIControllerGetStatus299 with default headers values
func NewAlertAPIControllerGetStatus299() *AlertAPIControllerGetStatus299 {
	return &AlertAPIControllerGetStatus299{}
}

/* AlertAPIControllerGetStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type AlertAPIControllerGetStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *AlertAPIControllerGetStatus299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/Alert][%d] alertApiControllerGetStatus299  %+v", 299, o.Payload)
}
func (o *AlertAPIControllerGetStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *AlertAPIControllerGetStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertAPIControllerGetNotModified creates a AlertAPIControllerGetNotModified with default headers values
func NewAlertAPIControllerGetNotModified() *AlertAPIControllerGetNotModified {
	return &AlertAPIControllerGetNotModified{}
}

/* AlertAPIControllerGetNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type AlertAPIControllerGetNotModified struct {
}

func (o *AlertAPIControllerGetNotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/Alert][%d] alertApiControllerGetNotModified ", 304)
}

func (o *AlertAPIControllerGetNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
