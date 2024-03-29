// Code generated by go-swagger; DO NOT EDIT.

package bus_advanced_near_by

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v2/models"
)

// BusAPIRouteNearBy2852Reader is a Reader for the BusAPIRouteNearBy2852 structure.
type BusAPIRouteNearBy2852Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BusAPIRouteNearBy2852Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBusAPIRouteNearBy2852OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewBusAPIRouteNearBy2852Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewBusAPIRouteNearBy2852NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBusAPIRouteNearBy2852OK creates a BusAPIRouteNearBy2852OK with default headers values
func NewBusAPIRouteNearBy2852OK() *BusAPIRouteNearBy2852OK {
	return &BusAPIRouteNearBy2852OK{}
}

/* BusAPIRouteNearBy2852OK describes a response with status code 200, with default header values.

Success
*/
type BusAPIRouteNearBy2852OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusRoute
}

func (o *BusAPIRouteNearBy2852OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/NearBy][%d] busApiRouteNearBy2852OK  %+v", 200, o.Payload)
}
func (o *BusAPIRouteNearBy2852OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusRoute {
	return o.Payload
}

func (o *BusAPIRouteNearBy2852OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBusAPIRouteNearBy2852Status299 creates a BusAPIRouteNearBy2852Status299 with default headers values
func NewBusAPIRouteNearBy2852Status299() *BusAPIRouteNearBy2852Status299 {
	return &BusAPIRouteNearBy2852Status299{}
}

/* BusAPIRouteNearBy2852Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type BusAPIRouteNearBy2852Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *BusAPIRouteNearBy2852Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/NearBy][%d] busApiRouteNearBy2852Status299  %+v", 299, o.Payload)
}
func (o *BusAPIRouteNearBy2852Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *BusAPIRouteNearBy2852Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBusAPIRouteNearBy2852NotModified creates a BusAPIRouteNearBy2852NotModified with default headers values
func NewBusAPIRouteNearBy2852NotModified() *BusAPIRouteNearBy2852NotModified {
	return &BusAPIRouteNearBy2852NotModified{}
}

/* BusAPIRouteNearBy2852NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type BusAPIRouteNearBy2852NotModified struct {
}

func (o *BusAPIRouteNearBy2852NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/NearBy][%d] busApiRouteNearBy2852NotModified ", 304)
}

func (o *BusAPIRouteNearBy2852NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
