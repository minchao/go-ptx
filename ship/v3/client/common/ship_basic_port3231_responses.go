// Code generated by go-swagger; DO NOT EDIT.

package common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/ship/v3/models"
)

// ShipBasicPort3231Reader is a Reader for the ShipBasicPort3231 structure.
type ShipBasicPort3231Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShipBasicPort3231Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShipBasicPort3231OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewShipBasicPort3231Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewShipBasicPort3231NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShipBasicPort3231OK creates a ShipBasicPort3231OK with default headers values
func NewShipBasicPort3231OK() *ShipBasicPort3231OK {
	return &ShipBasicPort3231OK{}
}

/* ShipBasicPort3231OK describes a response with status code 200, with default header values.

Success
*/
type ShipBasicPort3231OK struct {
	Payload *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Port
}

func (o *ShipBasicPort3231OK) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/Port][%d] shipBasicPort3231OK  %+v", 200, o.Payload)
}
func (o *ShipBasicPort3231OK) GetPayload() *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Port {
	return o.Payload
}

func (o *ShipBasicPort3231OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Port)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipBasicPort3231Status299 creates a ShipBasicPort3231Status299 with default headers values
func NewShipBasicPort3231Status299() *ShipBasicPort3231Status299 {
	return &ShipBasicPort3231Status299{}
}

/* ShipBasicPort3231Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type ShipBasicPort3231Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *ShipBasicPort3231Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/Port][%d] shipBasicPort3231Status299  %+v", 299, o.Payload)
}
func (o *ShipBasicPort3231Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *ShipBasicPort3231Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipBasicPort3231NotModified creates a ShipBasicPort3231NotModified with default headers values
func NewShipBasicPort3231NotModified() *ShipBasicPort3231NotModified {
	return &ShipBasicPort3231NotModified{}
}

/* ShipBasicPort3231NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type ShipBasicPort3231NotModified struct {
}

func (o *ShipBasicPort3231NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/Port][%d] shipBasicPort3231NotModified ", 304)
}

func (o *ShipBasicPort3231NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
