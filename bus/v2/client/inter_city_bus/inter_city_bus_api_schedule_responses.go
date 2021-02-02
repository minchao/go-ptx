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

// InterCityBusAPIScheduleReader is a Reader for the InterCityBusAPISchedule structure.
type InterCityBusAPIScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIScheduleStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIScheduleNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIScheduleOK creates a InterCityBusAPIScheduleOK with default headers values
func NewInterCityBusAPIScheduleOK() *InterCityBusAPIScheduleOK {
	return &InterCityBusAPIScheduleOK{}
}

/* InterCityBusAPIScheduleOK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIScheduleOK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusSchedule
}

func (o *InterCityBusAPIScheduleOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Schedule/InterCity][%d] interCityBusApiScheduleOK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIScheduleOK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusSchedule {
	return o.Payload
}

func (o *InterCityBusAPIScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIScheduleStatus299 creates a InterCityBusAPIScheduleStatus299 with default headers values
func NewInterCityBusAPIScheduleStatus299() *InterCityBusAPIScheduleStatus299 {
	return &InterCityBusAPIScheduleStatus299{}
}

/* InterCityBusAPIScheduleStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIScheduleStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIScheduleStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Schedule/InterCity][%d] interCityBusApiScheduleStatus299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIScheduleStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIScheduleStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIScheduleNotModified creates a InterCityBusAPIScheduleNotModified with default headers values
func NewInterCityBusAPIScheduleNotModified() *InterCityBusAPIScheduleNotModified {
	return &InterCityBusAPIScheduleNotModified{}
}

/* InterCityBusAPIScheduleNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIScheduleNotModified struct {
}

func (o *InterCityBusAPIScheduleNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Schedule/InterCity][%d] interCityBusApiScheduleNotModified ", 304)
}

func (o *InterCityBusAPIScheduleNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
