// Code generated by go-swagger; DO NOT EDIT.

package city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v2/models"
)

// CityBusAPIRealTimeNearStopUDP20471Reader is a Reader for the CityBusAPIRealTimeNearStopUDP20471 structure.
type CityBusAPIRealTimeNearStopUDP20471Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRealTimeNearStopUDP20471Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRealTimeNearStopUdp20471OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIRealTimeNearStopUdp20471Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIRealTimeNearStopUdp20471NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIRealTimeNearStopUdp20471OK creates a CityBusAPIRealTimeNearStopUdp20471OK with default headers values
func NewCityBusAPIRealTimeNearStopUdp20471OK() *CityBusAPIRealTimeNearStopUdp20471OK {
	return &CityBusAPIRealTimeNearStopUdp20471OK{}
}

/* CityBusAPIRealTimeNearStopUdp20471OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIRealTimeNearStopUdp20471OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusA2Data
}

func (o *CityBusAPIRealTimeNearStopUdp20471OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeNearStop/Streaming/City/{City}/{RouteName}][%d] cityBusApiRealTimeNearStopUdp20471OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIRealTimeNearStopUdp20471OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusA2Data {
	return o.Payload
}

func (o *CityBusAPIRealTimeNearStopUdp20471OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRealTimeNearStopUdp20471Status299 creates a CityBusAPIRealTimeNearStopUdp20471Status299 with default headers values
func NewCityBusAPIRealTimeNearStopUdp20471Status299() *CityBusAPIRealTimeNearStopUdp20471Status299 {
	return &CityBusAPIRealTimeNearStopUdp20471Status299{}
}

/* CityBusAPIRealTimeNearStopUdp20471Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIRealTimeNearStopUdp20471Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIRealTimeNearStopUdp20471Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeNearStop/Streaming/City/{City}/{RouteName}][%d] cityBusApiRealTimeNearStopUdp20471Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIRealTimeNearStopUdp20471Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIRealTimeNearStopUdp20471Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRealTimeNearStopUdp20471NotModified creates a CityBusAPIRealTimeNearStopUdp20471NotModified with default headers values
func NewCityBusAPIRealTimeNearStopUdp20471NotModified() *CityBusAPIRealTimeNearStopUdp20471NotModified {
	return &CityBusAPIRealTimeNearStopUdp20471NotModified{}
}

/* CityBusAPIRealTimeNearStopUdp20471NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIRealTimeNearStopUdp20471NotModified struct {
}

func (o *CityBusAPIRealTimeNearStopUdp20471NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeNearStop/Streaming/City/{City}/{RouteName}][%d] cityBusApiRealTimeNearStopUdp20471NotModified ", 304)
}

func (o *CityBusAPIRealTimeNearStopUdp20471NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
