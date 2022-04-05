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

// InterCityBusAPIRealTimeByFrequencyUDP2057Reader is a Reader for the InterCityBusAPIRealTimeByFrequencyUDP2057 structure.
type InterCityBusAPIRealTimeByFrequencyUDP2057Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIRealTimeByFrequencyUDP2057Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIRealTimeByFrequencyUdp2057OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIRealTimeByFrequencyUdp2057Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIRealTimeByFrequencyUdp2057NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIRealTimeByFrequencyUdp2057OK creates a InterCityBusAPIRealTimeByFrequencyUdp2057OK with default headers values
func NewInterCityBusAPIRealTimeByFrequencyUdp2057OK() *InterCityBusAPIRealTimeByFrequencyUdp2057OK {
	return &InterCityBusAPIRealTimeByFrequencyUdp2057OK{}
}

/* InterCityBusAPIRealTimeByFrequencyUdp2057OK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIRealTimeByFrequencyUdp2057OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusA1Data
}

func (o *InterCityBusAPIRealTimeByFrequencyUdp2057OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/Streaming/InterCity][%d] interCityBusApiRealTimeByFrequencyUdp2057OK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIRealTimeByFrequencyUdp2057OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusA1Data {
	return o.Payload
}

func (o *InterCityBusAPIRealTimeByFrequencyUdp2057OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIRealTimeByFrequencyUdp2057Status299 creates a InterCityBusAPIRealTimeByFrequencyUdp2057Status299 with default headers values
func NewInterCityBusAPIRealTimeByFrequencyUdp2057Status299() *InterCityBusAPIRealTimeByFrequencyUdp2057Status299 {
	return &InterCityBusAPIRealTimeByFrequencyUdp2057Status299{}
}

/* InterCityBusAPIRealTimeByFrequencyUdp2057Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIRealTimeByFrequencyUdp2057Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIRealTimeByFrequencyUdp2057Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/Streaming/InterCity][%d] interCityBusApiRealTimeByFrequencyUdp2057Status299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIRealTimeByFrequencyUdp2057Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIRealTimeByFrequencyUdp2057Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIRealTimeByFrequencyUdp2057NotModified creates a InterCityBusAPIRealTimeByFrequencyUdp2057NotModified with default headers values
func NewInterCityBusAPIRealTimeByFrequencyUdp2057NotModified() *InterCityBusAPIRealTimeByFrequencyUdp2057NotModified {
	return &InterCityBusAPIRealTimeByFrequencyUdp2057NotModified{}
}

/* InterCityBusAPIRealTimeByFrequencyUdp2057NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIRealTimeByFrequencyUdp2057NotModified struct {
}

func (o *InterCityBusAPIRealTimeByFrequencyUdp2057NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/Streaming/InterCity][%d] interCityBusApiRealTimeByFrequencyUdp2057NotModified ", 304)
}

func (o *InterCityBusAPIRealTimeByFrequencyUdp2057NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
