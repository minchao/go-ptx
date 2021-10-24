// Code generated by go-swagger; DO NOT EDIT.

package international_ship

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/ship/v3/models"
)

// ShipInternationalSpecificScheduleReader is a Reader for the ShipInternationalSpecificSchedule structure.
type ShipInternationalSpecificScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShipInternationalSpecificScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShipInternationalSpecificScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewShipInternationalSpecificScheduleStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewShipInternationalSpecificScheduleNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShipInternationalSpecificScheduleOK creates a ShipInternationalSpecificScheduleOK with default headers values
func NewShipInternationalSpecificScheduleOK() *ShipInternationalSpecificScheduleOK {
	return &ShipInternationalSpecificScheduleOK{}
}

/* ShipInternationalSpecificScheduleOK describes a response with status code 200, with default header values.

Success
*/
type ShipInternationalSpecificScheduleOK struct {
	Payload *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3SpecificSchedule
}

func (o *ShipInternationalSpecificScheduleOK) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/SpecificSchedule/International][%d] shipInternationalSpecificScheduleOK  %+v", 200, o.Payload)
}
func (o *ShipInternationalSpecificScheduleOK) GetPayload() *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3SpecificSchedule {
	return o.Payload
}

func (o *ShipInternationalSpecificScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3SpecificSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipInternationalSpecificScheduleStatus299 creates a ShipInternationalSpecificScheduleStatus299 with default headers values
func NewShipInternationalSpecificScheduleStatus299() *ShipInternationalSpecificScheduleStatus299 {
	return &ShipInternationalSpecificScheduleStatus299{}
}

/* ShipInternationalSpecificScheduleStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type ShipInternationalSpecificScheduleStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *ShipInternationalSpecificScheduleStatus299) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/SpecificSchedule/International][%d] shipInternationalSpecificScheduleStatus299  %+v", 299, o.Payload)
}
func (o *ShipInternationalSpecificScheduleStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *ShipInternationalSpecificScheduleStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipInternationalSpecificScheduleNotModified creates a ShipInternationalSpecificScheduleNotModified with default headers values
func NewShipInternationalSpecificScheduleNotModified() *ShipInternationalSpecificScheduleNotModified {
	return &ShipInternationalSpecificScheduleNotModified{}
}

/* ShipInternationalSpecificScheduleNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type ShipInternationalSpecificScheduleNotModified struct {
}

func (o *ShipInternationalSpecificScheduleNotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/SpecificSchedule/International][%d] shipInternationalSpecificScheduleNotModified ", 304)
}

func (o *ShipInternationalSpecificScheduleNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
