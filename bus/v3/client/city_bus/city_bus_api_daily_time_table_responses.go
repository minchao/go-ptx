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

// CityBusAPIDailyTimeTableReader is a Reader for the CityBusAPIDailyTimeTable structure.
type CityBusAPIDailyTimeTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIDailyTimeTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIDailyTimeTableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIDailyTimeTableStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIDailyTimeTableNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIDailyTimeTableOK creates a CityBusAPIDailyTimeTableOK with default headers values
func NewCityBusAPIDailyTimeTableOK() *CityBusAPIDailyTimeTableOK {
	return &CityBusAPIDailyTimeTableOK{}
}

/* CityBusAPIDailyTimeTableOK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIDailyTimeTableOK struct {
	Payload *models.PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable
}

func (o *CityBusAPIDailyTimeTableOK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/DailyTimeTable/City/{City}][%d] cityBusApiDailyTimeTableOK  %+v", 200, o.Payload)
}
func (o *CityBusAPIDailyTimeTableOK) GetPayload() *models.PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable {
	return o.Payload
}

func (o *CityBusAPIDailyTimeTableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIDailyTimeTableStatus299 creates a CityBusAPIDailyTimeTableStatus299 with default headers values
func NewCityBusAPIDailyTimeTableStatus299() *CityBusAPIDailyTimeTableStatus299 {
	return &CityBusAPIDailyTimeTableStatus299{}
}

/* CityBusAPIDailyTimeTableStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIDailyTimeTableStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIDailyTimeTableStatus299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/DailyTimeTable/City/{City}][%d] cityBusApiDailyTimeTableStatus299  %+v", 299, o.Payload)
}
func (o *CityBusAPIDailyTimeTableStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIDailyTimeTableStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIDailyTimeTableNotModified creates a CityBusAPIDailyTimeTableNotModified with default headers values
func NewCityBusAPIDailyTimeTableNotModified() *CityBusAPIDailyTimeTableNotModified {
	return &CityBusAPIDailyTimeTableNotModified{}
}

/* CityBusAPIDailyTimeTableNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIDailyTimeTableNotModified struct {
}

func (o *CityBusAPIDailyTimeTableNotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/DailyTimeTable/City/{City}][%d] cityBusApiDailyTimeTableNotModified ", 304)
}

func (o *CityBusAPIDailyTimeTableNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
