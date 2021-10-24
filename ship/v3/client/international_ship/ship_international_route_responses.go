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

// ShipInternationalRouteReader is a Reader for the ShipInternationalRoute structure.
type ShipInternationalRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShipInternationalRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShipInternationalRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewShipInternationalRouteStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewShipInternationalRouteNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShipInternationalRouteOK creates a ShipInternationalRouteOK with default headers values
func NewShipInternationalRouteOK() *ShipInternationalRouteOK {
	return &ShipInternationalRouteOK{}
}

/* ShipInternationalRouteOK describes a response with status code 200, with default header values.

Success
*/
type ShipInternationalRouteOK struct {
	Payload *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Route
}

func (o *ShipInternationalRouteOK) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/Route/International][%d] shipInternationalRouteOK  %+v", 200, o.Payload)
}
func (o *ShipInternationalRouteOK) GetPayload() *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Route {
	return o.Payload
}

func (o *ShipInternationalRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Route)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipInternationalRouteStatus299 creates a ShipInternationalRouteStatus299 with default headers values
func NewShipInternationalRouteStatus299() *ShipInternationalRouteStatus299 {
	return &ShipInternationalRouteStatus299{}
}

/* ShipInternationalRouteStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type ShipInternationalRouteStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *ShipInternationalRouteStatus299) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/Route/International][%d] shipInternationalRouteStatus299  %+v", 299, o.Payload)
}
func (o *ShipInternationalRouteStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *ShipInternationalRouteStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipInternationalRouteNotModified creates a ShipInternationalRouteNotModified with default headers values
func NewShipInternationalRouteNotModified() *ShipInternationalRouteNotModified {
	return &ShipInternationalRouteNotModified{}
}

/* ShipInternationalRouteNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type ShipInternationalRouteNotModified struct {
}

func (o *ShipInternationalRouteNotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/Route/International][%d] shipInternationalRouteNotModified ", 304)
}

func (o *ShipInternationalRouteNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
