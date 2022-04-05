// Code generated by go-swagger; DO NOT EDIT.

package ship_by_route_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/ship/v3/models"
)

// ShipRouteTypeDailySchedule3259Reader is a Reader for the ShipRouteTypeDailySchedule3259 structure.
type ShipRouteTypeDailySchedule3259Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShipRouteTypeDailySchedule3259Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShipRouteTypeDailySchedule3259OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewShipRouteTypeDailySchedule3259Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewShipRouteTypeDailySchedule3259NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShipRouteTypeDailySchedule3259OK creates a ShipRouteTypeDailySchedule3259OK with default headers values
func NewShipRouteTypeDailySchedule3259OK() *ShipRouteTypeDailySchedule3259OK {
	return &ShipRouteTypeDailySchedule3259OK{}
}

/* ShipRouteTypeDailySchedule3259OK describes a response with status code 200, with default header values.

Success
*/
type ShipRouteTypeDailySchedule3259OK struct {
	Payload []*models.PTXServiceDTOShipSpecificationV3ShipDailySchedule
}

func (o *ShipRouteTypeDailySchedule3259OK) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/DailySchedule/RouteType/{RouteType}][%d] shipRouteTypeDailySchedule3259OK  %+v", 200, o.Payload)
}
func (o *ShipRouteTypeDailySchedule3259OK) GetPayload() []*models.PTXServiceDTOShipSpecificationV3ShipDailySchedule {
	return o.Payload
}

func (o *ShipRouteTypeDailySchedule3259OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipRouteTypeDailySchedule3259Status299 creates a ShipRouteTypeDailySchedule3259Status299 with default headers values
func NewShipRouteTypeDailySchedule3259Status299() *ShipRouteTypeDailySchedule3259Status299 {
	return &ShipRouteTypeDailySchedule3259Status299{}
}

/* ShipRouteTypeDailySchedule3259Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type ShipRouteTypeDailySchedule3259Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *ShipRouteTypeDailySchedule3259Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/DailySchedule/RouteType/{RouteType}][%d] shipRouteTypeDailySchedule3259Status299  %+v", 299, o.Payload)
}
func (o *ShipRouteTypeDailySchedule3259Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *ShipRouteTypeDailySchedule3259Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipRouteTypeDailySchedule3259NotModified creates a ShipRouteTypeDailySchedule3259NotModified with default headers values
func NewShipRouteTypeDailySchedule3259NotModified() *ShipRouteTypeDailySchedule3259NotModified {
	return &ShipRouteTypeDailySchedule3259NotModified{}
}

/* ShipRouteTypeDailySchedule3259NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type ShipRouteTypeDailySchedule3259NotModified struct {
}

func (o *ShipRouteTypeDailySchedule3259NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/DailySchedule/RouteType/{RouteType}][%d] shipRouteTypeDailySchedule3259NotModified ", 304)
}

func (o *ShipRouteTypeDailySchedule3259NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
