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

// CityBusAPIRealTimeByFrequencyUDP1Reader is a Reader for the CityBusAPIRealTimeByFrequencyUDP1 structure.
type CityBusAPIRealTimeByFrequencyUDP1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRealTimeByFrequencyUDP1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRealTimeByFrequencyUdp1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIRealTimeByFrequencyUdp1Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIRealTimeByFrequencyUdp1NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIRealTimeByFrequencyUdp1OK creates a CityBusAPIRealTimeByFrequencyUdp1OK with default headers values
func NewCityBusAPIRealTimeByFrequencyUdp1OK() *CityBusAPIRealTimeByFrequencyUdp1OK {
	return &CityBusAPIRealTimeByFrequencyUdp1OK{}
}

/* CityBusAPIRealTimeByFrequencyUdp1OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIRealTimeByFrequencyUdp1OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusA1Data
}

func (o *CityBusAPIRealTimeByFrequencyUdp1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/Streaming/City/{City}/{RouteName}][%d] cityBusApiRealTimeByFrequencyUdp1OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIRealTimeByFrequencyUdp1OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusA1Data {
	return o.Payload
}

func (o *CityBusAPIRealTimeByFrequencyUdp1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRealTimeByFrequencyUdp1Status299 creates a CityBusAPIRealTimeByFrequencyUdp1Status299 with default headers values
func NewCityBusAPIRealTimeByFrequencyUdp1Status299() *CityBusAPIRealTimeByFrequencyUdp1Status299 {
	return &CityBusAPIRealTimeByFrequencyUdp1Status299{}
}

/* CityBusAPIRealTimeByFrequencyUdp1Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIRealTimeByFrequencyUdp1Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIRealTimeByFrequencyUdp1Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/Streaming/City/{City}/{RouteName}][%d] cityBusApiRealTimeByFrequencyUdp1Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIRealTimeByFrequencyUdp1Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIRealTimeByFrequencyUdp1Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRealTimeByFrequencyUdp1NotModified creates a CityBusAPIRealTimeByFrequencyUdp1NotModified with default headers values
func NewCityBusAPIRealTimeByFrequencyUdp1NotModified() *CityBusAPIRealTimeByFrequencyUdp1NotModified {
	return &CityBusAPIRealTimeByFrequencyUdp1NotModified{}
}

/* CityBusAPIRealTimeByFrequencyUdp1NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIRealTimeByFrequencyUdp1NotModified struct {
}

func (o *CityBusAPIRealTimeByFrequencyUdp1NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/Streaming/City/{City}/{RouteName}][%d] cityBusApiRealTimeByFrequencyUdp1NotModified ", 304)
}

func (o *CityBusAPIRealTimeByFrequencyUdp1NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
