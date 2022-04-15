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

// InterCityBusAPISchedule2069Reader is a Reader for the InterCityBusAPISchedule2069 structure.
type InterCityBusAPISchedule2069Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPISchedule2069Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPISchedule2069OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPISchedule2069Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPISchedule2069NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPISchedule2069OK creates a InterCityBusAPISchedule2069OK with default headers values
func NewInterCityBusAPISchedule2069OK() *InterCityBusAPISchedule2069OK {
	return &InterCityBusAPISchedule2069OK{}
}

/* InterCityBusAPISchedule2069OK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPISchedule2069OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusSchedule
}

func (o *InterCityBusAPISchedule2069OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Schedule/InterCity][%d] interCityBusApiSchedule2069OK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPISchedule2069OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusSchedule {
	return o.Payload
}

func (o *InterCityBusAPISchedule2069OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPISchedule2069Status299 creates a InterCityBusAPISchedule2069Status299 with default headers values
func NewInterCityBusAPISchedule2069Status299() *InterCityBusAPISchedule2069Status299 {
	return &InterCityBusAPISchedule2069Status299{}
}

/* InterCityBusAPISchedule2069Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPISchedule2069Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPISchedule2069Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Schedule/InterCity][%d] interCityBusApiSchedule2069Status299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPISchedule2069Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPISchedule2069Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPISchedule2069NotModified creates a InterCityBusAPISchedule2069NotModified with default headers values
func NewInterCityBusAPISchedule2069NotModified() *InterCityBusAPISchedule2069NotModified {
	return &InterCityBusAPISchedule2069NotModified{}
}

/* InterCityBusAPISchedule2069NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPISchedule2069NotModified struct {
}

func (o *InterCityBusAPISchedule2069NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Schedule/InterCity][%d] interCityBusApiSchedule2069NotModified ", 304)
}

func (o *InterCityBusAPISchedule2069NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}