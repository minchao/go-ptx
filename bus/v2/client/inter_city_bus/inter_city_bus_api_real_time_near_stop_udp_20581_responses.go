// Code generated by go-swagger; DO NOT EDIT.

package inter_city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v2/models"
)

// InterCityBusAPIRealTimeNearStopUDP20581Reader is a Reader for the InterCityBusAPIRealTimeNearStopUDP20581 structure.
type InterCityBusAPIRealTimeNearStopUDP20581Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIRealTimeNearStopUDP20581Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIRealTimeNearStopUdp20581OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIRealTimeNearStopUdp20581Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIRealTimeNearStopUdp20581NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIRealTimeNearStopUdp20581OK creates a InterCityBusAPIRealTimeNearStopUdp20581OK with default headers values
func NewInterCityBusAPIRealTimeNearStopUdp20581OK() *InterCityBusAPIRealTimeNearStopUdp20581OK {
	return &InterCityBusAPIRealTimeNearStopUdp20581OK{}
}

/* InterCityBusAPIRealTimeNearStopUdp20581OK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIRealTimeNearStopUdp20581OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusA2Data
}

func (o *InterCityBusAPIRealTimeNearStopUdp20581OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeNearStop/Streaming/InterCity/{RouteName}][%d] interCityBusApiRealTimeNearStopUdp20581OK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIRealTimeNearStopUdp20581OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusA2Data {
	return o.Payload
}

func (o *InterCityBusAPIRealTimeNearStopUdp20581OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIRealTimeNearStopUdp20581Status299 creates a InterCityBusAPIRealTimeNearStopUdp20581Status299 with default headers values
func NewInterCityBusAPIRealTimeNearStopUdp20581Status299() *InterCityBusAPIRealTimeNearStopUdp20581Status299 {
	return &InterCityBusAPIRealTimeNearStopUdp20581Status299{}
}

/* InterCityBusAPIRealTimeNearStopUdp20581Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIRealTimeNearStopUdp20581Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIRealTimeNearStopUdp20581Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeNearStop/Streaming/InterCity/{RouteName}][%d] interCityBusApiRealTimeNearStopUdp20581Status299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIRealTimeNearStopUdp20581Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIRealTimeNearStopUdp20581Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIRealTimeNearStopUdp20581NotModified creates a InterCityBusAPIRealTimeNearStopUdp20581NotModified with default headers values
func NewInterCityBusAPIRealTimeNearStopUdp20581NotModified() *InterCityBusAPIRealTimeNearStopUdp20581NotModified {
	return &InterCityBusAPIRealTimeNearStopUdp20581NotModified{}
}

/* InterCityBusAPIRealTimeNearStopUdp20581NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIRealTimeNearStopUdp20581NotModified struct {
}

func (o *InterCityBusAPIRealTimeNearStopUdp20581NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeNearStop/Streaming/InterCity/{RouteName}][%d] interCityBusApiRealTimeNearStopUdp20581NotModified ", 304)
}

func (o *InterCityBusAPIRealTimeNearStopUdp20581NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
