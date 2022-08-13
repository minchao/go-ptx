// Code generated by go-swagger; DO NOT EDIT.

package ship_advanced_by_route_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/ship/v3/models"
)

// ShipRouteTypeGeneralSchedule23269Reader is a Reader for the ShipRouteTypeGeneralSchedule23269 structure.
type ShipRouteTypeGeneralSchedule23269Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShipRouteTypeGeneralSchedule23269Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShipRouteTypeGeneralSchedule23269OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewShipRouteTypeGeneralSchedule23269Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewShipRouteTypeGeneralSchedule23269NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShipRouteTypeGeneralSchedule23269OK creates a ShipRouteTypeGeneralSchedule23269OK with default headers values
func NewShipRouteTypeGeneralSchedule23269OK() *ShipRouteTypeGeneralSchedule23269OK {
	return &ShipRouteTypeGeneralSchedule23269OK{}
}

/* ShipRouteTypeGeneralSchedule23269OK describes a response with status code 200, with default header values.

Success
*/
type ShipRouteTypeGeneralSchedule23269OK struct {
	Payload []*models.PTXServiceDTOShipSpecificationV3ShipGeneralSchedule
}

func (o *ShipRouteTypeGeneralSchedule23269OK) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/GeneralSchedule/RouteID/{RouteID}][%d] shipRouteTypeGeneralSchedule23269OK  %+v", 200, o.Payload)
}
func (o *ShipRouteTypeGeneralSchedule23269OK) GetPayload() []*models.PTXServiceDTOShipSpecificationV3ShipGeneralSchedule {
	return o.Payload
}

func (o *ShipRouteTypeGeneralSchedule23269OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipRouteTypeGeneralSchedule23269Status299 creates a ShipRouteTypeGeneralSchedule23269Status299 with default headers values
func NewShipRouteTypeGeneralSchedule23269Status299() *ShipRouteTypeGeneralSchedule23269Status299 {
	return &ShipRouteTypeGeneralSchedule23269Status299{}
}

/* ShipRouteTypeGeneralSchedule23269Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type ShipRouteTypeGeneralSchedule23269Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *ShipRouteTypeGeneralSchedule23269Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/GeneralSchedule/RouteID/{RouteID}][%d] shipRouteTypeGeneralSchedule23269Status299  %+v", 299, o.Payload)
}
func (o *ShipRouteTypeGeneralSchedule23269Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *ShipRouteTypeGeneralSchedule23269Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipRouteTypeGeneralSchedule23269NotModified creates a ShipRouteTypeGeneralSchedule23269NotModified with default headers values
func NewShipRouteTypeGeneralSchedule23269NotModified() *ShipRouteTypeGeneralSchedule23269NotModified {
	return &ShipRouteTypeGeneralSchedule23269NotModified{}
}

/* ShipRouteTypeGeneralSchedule23269NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type ShipRouteTypeGeneralSchedule23269NotModified struct {
}

func (o *ShipRouteTypeGeneralSchedule23269NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/GeneralSchedule/RouteID/{RouteID}][%d] shipRouteTypeGeneralSchedule23269NotModified ", 304)
}

func (o *ShipRouteTypeGeneralSchedule23269NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}