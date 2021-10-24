// Code generated by go-swagger; DO NOT EDIT.

package domestic_ship

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/ship/v3/models"
)

// ShipDomesticRouteReader is a Reader for the ShipDomesticRoute structure.
type ShipDomesticRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShipDomesticRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShipDomesticRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewShipDomesticRouteStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewShipDomesticRouteNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShipDomesticRouteOK creates a ShipDomesticRouteOK with default headers values
func NewShipDomesticRouteOK() *ShipDomesticRouteOK {
	return &ShipDomesticRouteOK{}
}

/* ShipDomesticRouteOK describes a response with status code 200, with default header values.

Success
*/
type ShipDomesticRouteOK struct {
	Payload *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Route
}

func (o *ShipDomesticRouteOK) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/Route/Domestic/City/{City}][%d] shipDomesticRouteOK  %+v", 200, o.Payload)
}
func (o *ShipDomesticRouteOK) GetPayload() *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Route {
	return o.Payload
}

func (o *ShipDomesticRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Route)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipDomesticRouteStatus299 creates a ShipDomesticRouteStatus299 with default headers values
func NewShipDomesticRouteStatus299() *ShipDomesticRouteStatus299 {
	return &ShipDomesticRouteStatus299{}
}

/* ShipDomesticRouteStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type ShipDomesticRouteStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *ShipDomesticRouteStatus299) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/Route/Domestic/City/{City}][%d] shipDomesticRouteStatus299  %+v", 299, o.Payload)
}
func (o *ShipDomesticRouteStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *ShipDomesticRouteStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipDomesticRouteNotModified creates a ShipDomesticRouteNotModified with default headers values
func NewShipDomesticRouteNotModified() *ShipDomesticRouteNotModified {
	return &ShipDomesticRouteNotModified{}
}

/* ShipDomesticRouteNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type ShipDomesticRouteNotModified struct {
}

func (o *ShipDomesticRouteNotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/Route/Domestic/City/{City}][%d] shipDomesticRouteNotModified ", 304)
}

func (o *ShipDomesticRouteNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
