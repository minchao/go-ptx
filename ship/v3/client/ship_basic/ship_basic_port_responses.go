// Code generated by go-swagger; DO NOT EDIT.

package ship_basic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/ship/v3/models"
)

// ShipBasicPortReader is a Reader for the ShipBasicPort structure.
type ShipBasicPortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShipBasicPortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShipBasicPortOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewShipBasicPortStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewShipBasicPortNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShipBasicPortOK creates a ShipBasicPortOK with default headers values
func NewShipBasicPortOK() *ShipBasicPortOK {
	return &ShipBasicPortOK{}
}

/* ShipBasicPortOK describes a response with status code 200, with default header values.

Success
*/
type ShipBasicPortOK struct {
	Payload *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Port
}

func (o *ShipBasicPortOK) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/Basic/Port][%d] shipBasicPortOK  %+v", 200, o.Payload)
}
func (o *ShipBasicPortOK) GetPayload() *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Port {
	return o.Payload
}

func (o *ShipBasicPortOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Port)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipBasicPortStatus299 creates a ShipBasicPortStatus299 with default headers values
func NewShipBasicPortStatus299() *ShipBasicPortStatus299 {
	return &ShipBasicPortStatus299{}
}

/* ShipBasicPortStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type ShipBasicPortStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *ShipBasicPortStatus299) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/Basic/Port][%d] shipBasicPortStatus299  %+v", 299, o.Payload)
}
func (o *ShipBasicPortStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *ShipBasicPortStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipBasicPortNotModified creates a ShipBasicPortNotModified with default headers values
func NewShipBasicPortNotModified() *ShipBasicPortNotModified {
	return &ShipBasicPortNotModified{}
}

/* ShipBasicPortNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type ShipBasicPortNotModified struct {
}

func (o *ShipBasicPortNotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/Basic/Port][%d] shipBasicPortNotModified ", 304)
}

func (o *ShipBasicPortNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
