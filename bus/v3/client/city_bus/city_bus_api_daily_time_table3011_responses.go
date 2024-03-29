// Code generated by go-swagger; DO NOT EDIT.

package city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v3/models"
)

// CityBusAPIDailyTimeTable3011Reader is a Reader for the CityBusAPIDailyTimeTable3011 structure.
type CityBusAPIDailyTimeTable3011Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIDailyTimeTable3011Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIDailyTimeTable3011OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIDailyTimeTable3011Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIDailyTimeTable3011NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIDailyTimeTable3011OK creates a CityBusAPIDailyTimeTable3011OK with default headers values
func NewCityBusAPIDailyTimeTable3011OK() *CityBusAPIDailyTimeTable3011OK {
	return &CityBusAPIDailyTimeTable3011OK{}
}

/* CityBusAPIDailyTimeTable3011OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIDailyTimeTable3011OK struct {
	Payload *models.PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable
}

func (o *CityBusAPIDailyTimeTable3011OK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/DailyTimeTable/City/{City}][%d] cityBusApiDailyTimeTable3011OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIDailyTimeTable3011OK) GetPayload() *models.PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable {
	return o.Payload
}

func (o *CityBusAPIDailyTimeTable3011OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIDailyTimeTable3011Status299 creates a CityBusAPIDailyTimeTable3011Status299 with default headers values
func NewCityBusAPIDailyTimeTable3011Status299() *CityBusAPIDailyTimeTable3011Status299 {
	return &CityBusAPIDailyTimeTable3011Status299{}
}

/* CityBusAPIDailyTimeTable3011Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIDailyTimeTable3011Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIDailyTimeTable3011Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/DailyTimeTable/City/{City}][%d] cityBusApiDailyTimeTable3011Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIDailyTimeTable3011Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIDailyTimeTable3011Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIDailyTimeTable3011NotModified creates a CityBusAPIDailyTimeTable3011NotModified with default headers values
func NewCityBusAPIDailyTimeTable3011NotModified() *CityBusAPIDailyTimeTable3011NotModified {
	return &CityBusAPIDailyTimeTable3011NotModified{}
}

/* CityBusAPIDailyTimeTable3011NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIDailyTimeTable3011NotModified struct {
}

func (o *CityBusAPIDailyTimeTable3011NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/DailyTimeTable/City/{City}][%d] cityBusApiDailyTimeTable3011NotModified ", 304)
}

func (o *CityBusAPIDailyTimeTable3011NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
