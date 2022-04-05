// Code generated by go-swagger; DO NOT EDIT.

package bus_advanced_by_station

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v2/models"
)

// InterCityBusAPIScheduleByStation2902Reader is a Reader for the InterCityBusAPIScheduleByStation2902 structure.
type InterCityBusAPIScheduleByStation2902Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIScheduleByStation2902Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIScheduleByStation2902OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIScheduleByStation2902Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIScheduleByStation2902NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIScheduleByStation2902OK creates a InterCityBusAPIScheduleByStation2902OK with default headers values
func NewInterCityBusAPIScheduleByStation2902OK() *InterCityBusAPIScheduleByStation2902OK {
	return &InterCityBusAPIScheduleByStation2902OK{}
}

/* InterCityBusAPIScheduleByStation2902OK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIScheduleByStation2902OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusSchedule
}

func (o *InterCityBusAPIScheduleByStation2902OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Schedule/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiScheduleByStation2902OK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIScheduleByStation2902OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusSchedule {
	return o.Payload
}

func (o *InterCityBusAPIScheduleByStation2902OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIScheduleByStation2902Status299 creates a InterCityBusAPIScheduleByStation2902Status299 with default headers values
func NewInterCityBusAPIScheduleByStation2902Status299() *InterCityBusAPIScheduleByStation2902Status299 {
	return &InterCityBusAPIScheduleByStation2902Status299{}
}

/* InterCityBusAPIScheduleByStation2902Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIScheduleByStation2902Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIScheduleByStation2902Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Schedule/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiScheduleByStation2902Status299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIScheduleByStation2902Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIScheduleByStation2902Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIScheduleByStation2902NotModified creates a InterCityBusAPIScheduleByStation2902NotModified with default headers values
func NewInterCityBusAPIScheduleByStation2902NotModified() *InterCityBusAPIScheduleByStation2902NotModified {
	return &InterCityBusAPIScheduleByStation2902NotModified{}
}

/* InterCityBusAPIScheduleByStation2902NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIScheduleByStation2902NotModified struct {
}

func (o *InterCityBusAPIScheduleByStation2902NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Schedule/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiScheduleByStation2902NotModified ", 304)
}

func (o *InterCityBusAPIScheduleByStation2902NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
