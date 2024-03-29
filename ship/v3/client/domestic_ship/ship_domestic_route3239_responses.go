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

// ShipDomesticRoute3239Reader is a Reader for the ShipDomesticRoute3239 structure.
type ShipDomesticRoute3239Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShipDomesticRoute3239Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShipDomesticRoute3239OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewShipDomesticRoute3239Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewShipDomesticRoute3239NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShipDomesticRoute3239OK creates a ShipDomesticRoute3239OK with default headers values
func NewShipDomesticRoute3239OK() *ShipDomesticRoute3239OK {
	return &ShipDomesticRoute3239OK{}
}

/* ShipDomesticRoute3239OK describes a response with status code 200, with default header values.

Success
*/
type ShipDomesticRoute3239OK struct {
	Payload *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Route
}

func (o *ShipDomesticRoute3239OK) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/Route/Domestic/City/{City}][%d] shipDomesticRoute3239OK  %+v", 200, o.Payload)
}
func (o *ShipDomesticRoute3239OK) GetPayload() *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Route {
	return o.Payload
}

func (o *ShipDomesticRoute3239OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Route)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipDomesticRoute3239Status299 creates a ShipDomesticRoute3239Status299 with default headers values
func NewShipDomesticRoute3239Status299() *ShipDomesticRoute3239Status299 {
	return &ShipDomesticRoute3239Status299{}
}

/* ShipDomesticRoute3239Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type ShipDomesticRoute3239Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *ShipDomesticRoute3239Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/Route/Domestic/City/{City}][%d] shipDomesticRoute3239Status299  %+v", 299, o.Payload)
}
func (o *ShipDomesticRoute3239Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *ShipDomesticRoute3239Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipDomesticRoute3239NotModified creates a ShipDomesticRoute3239NotModified with default headers values
func NewShipDomesticRoute3239NotModified() *ShipDomesticRoute3239NotModified {
	return &ShipDomesticRoute3239NotModified{}
}

/* ShipDomesticRoute3239NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type ShipDomesticRoute3239NotModified struct {
}

func (o *ShipDomesticRoute3239NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/Route/Domestic/City/{City}][%d] shipDomesticRoute3239NotModified ", 304)
}

func (o *ShipDomesticRoute3239NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
