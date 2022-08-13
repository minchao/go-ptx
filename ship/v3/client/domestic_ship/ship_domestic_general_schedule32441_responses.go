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

// ShipDomesticGeneralSchedule32441Reader is a Reader for the ShipDomesticGeneralSchedule32441 structure.
type ShipDomesticGeneralSchedule32441Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShipDomesticGeneralSchedule32441Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShipDomesticGeneralSchedule32441OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewShipDomesticGeneralSchedule32441Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewShipDomesticGeneralSchedule32441NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShipDomesticGeneralSchedule32441OK creates a ShipDomesticGeneralSchedule32441OK with default headers values
func NewShipDomesticGeneralSchedule32441OK() *ShipDomesticGeneralSchedule32441OK {
	return &ShipDomesticGeneralSchedule32441OK{}
}

/* ShipDomesticGeneralSchedule32441OK describes a response with status code 200, with default header values.

Success
*/
type ShipDomesticGeneralSchedule32441OK struct {
	Payload *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3GeneralSchedule
}

func (o *ShipDomesticGeneralSchedule32441OK) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/GeneralSchedule/Domestic/City/{City}/{RouteID}][%d] shipDomesticGeneralSchedule32441OK  %+v", 200, o.Payload)
}
func (o *ShipDomesticGeneralSchedule32441OK) GetPayload() *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3GeneralSchedule {
	return o.Payload
}

func (o *ShipDomesticGeneralSchedule32441OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3GeneralSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipDomesticGeneralSchedule32441Status299 creates a ShipDomesticGeneralSchedule32441Status299 with default headers values
func NewShipDomesticGeneralSchedule32441Status299() *ShipDomesticGeneralSchedule32441Status299 {
	return &ShipDomesticGeneralSchedule32441Status299{}
}

/* ShipDomesticGeneralSchedule32441Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type ShipDomesticGeneralSchedule32441Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *ShipDomesticGeneralSchedule32441Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/GeneralSchedule/Domestic/City/{City}/{RouteID}][%d] shipDomesticGeneralSchedule32441Status299  %+v", 299, o.Payload)
}
func (o *ShipDomesticGeneralSchedule32441Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *ShipDomesticGeneralSchedule32441Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipDomesticGeneralSchedule32441NotModified creates a ShipDomesticGeneralSchedule32441NotModified with default headers values
func NewShipDomesticGeneralSchedule32441NotModified() *ShipDomesticGeneralSchedule32441NotModified {
	return &ShipDomesticGeneralSchedule32441NotModified{}
}

/* ShipDomesticGeneralSchedule32441NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type ShipDomesticGeneralSchedule32441NotModified struct {
}

func (o *ShipDomesticGeneralSchedule32441NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/GeneralSchedule/Domestic/City/{City}/{RouteID}][%d] shipDomesticGeneralSchedule32441NotModified ", 304)
}

func (o *ShipDomesticGeneralSchedule32441NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}