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

// CityBusAPIDailyStopTimeTable24371Reader is a Reader for the CityBusAPIDailyStopTimeTable24371 structure.
type CityBusAPIDailyStopTimeTable24371Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIDailyStopTimeTable24371Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIDailyStopTimeTable24371OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIDailyStopTimeTable24371Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIDailyStopTimeTable24371NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIDailyStopTimeTable24371OK creates a CityBusAPIDailyStopTimeTable24371OK with default headers values
func NewCityBusAPIDailyStopTimeTable24371OK() *CityBusAPIDailyStopTimeTable24371OK {
	return &CityBusAPIDailyStopTimeTable24371OK{}
}

/* CityBusAPIDailyStopTimeTable24371OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIDailyStopTimeTable24371OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable
}

func (o *CityBusAPIDailyStopTimeTable24371OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/DailyStopTimeTable/City/{City}/{RouteName}][%d] cityBusApiDailyStopTimeTable24371OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIDailyStopTimeTable24371OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable {
	return o.Payload
}

func (o *CityBusAPIDailyStopTimeTable24371OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIDailyStopTimeTable24371Status299 creates a CityBusAPIDailyStopTimeTable24371Status299 with default headers values
func NewCityBusAPIDailyStopTimeTable24371Status299() *CityBusAPIDailyStopTimeTable24371Status299 {
	return &CityBusAPIDailyStopTimeTable24371Status299{}
}

/* CityBusAPIDailyStopTimeTable24371Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIDailyStopTimeTable24371Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIDailyStopTimeTable24371Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/DailyStopTimeTable/City/{City}/{RouteName}][%d] cityBusApiDailyStopTimeTable24371Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIDailyStopTimeTable24371Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIDailyStopTimeTable24371Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIDailyStopTimeTable24371NotModified creates a CityBusAPIDailyStopTimeTable24371NotModified with default headers values
func NewCityBusAPIDailyStopTimeTable24371NotModified() *CityBusAPIDailyStopTimeTable24371NotModified {
	return &CityBusAPIDailyStopTimeTable24371NotModified{}
}

/* CityBusAPIDailyStopTimeTable24371NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIDailyStopTimeTable24371NotModified struct {
}

func (o *CityBusAPIDailyStopTimeTable24371NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/DailyStopTimeTable/City/{City}/{RouteName}][%d] cityBusApiDailyStopTimeTable24371NotModified ", 304)
}

func (o *CityBusAPIDailyStopTimeTable24371NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
