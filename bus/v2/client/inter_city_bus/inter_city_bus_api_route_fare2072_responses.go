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

// InterCityBusAPIRouteFare2072Reader is a Reader for the InterCityBusAPIRouteFare2072 structure.
type InterCityBusAPIRouteFare2072Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIRouteFare2072Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIRouteFare2072OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIRouteFare2072Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIRouteFare2072NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIRouteFare2072OK creates a InterCityBusAPIRouteFare2072OK with default headers values
func NewInterCityBusAPIRouteFare2072OK() *InterCityBusAPIRouteFare2072OK {
	return &InterCityBusAPIRouteFare2072OK{}
}

/* InterCityBusAPIRouteFare2072OK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIRouteFare2072OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusRouteFare
}

func (o *InterCityBusAPIRouteFare2072OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RouteFare/InterCity][%d] interCityBusApiRouteFare2072OK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIRouteFare2072OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusRouteFare {
	return o.Payload
}

func (o *InterCityBusAPIRouteFare2072OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIRouteFare2072Status299 creates a InterCityBusAPIRouteFare2072Status299 with default headers values
func NewInterCityBusAPIRouteFare2072Status299() *InterCityBusAPIRouteFare2072Status299 {
	return &InterCityBusAPIRouteFare2072Status299{}
}

/* InterCityBusAPIRouteFare2072Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIRouteFare2072Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIRouteFare2072Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RouteFare/InterCity][%d] interCityBusApiRouteFare2072Status299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIRouteFare2072Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIRouteFare2072Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIRouteFare2072NotModified creates a InterCityBusAPIRouteFare2072NotModified with default headers values
func NewInterCityBusAPIRouteFare2072NotModified() *InterCityBusAPIRouteFare2072NotModified {
	return &InterCityBusAPIRouteFare2072NotModified{}
}

/* InterCityBusAPIRouteFare2072NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIRouteFare2072NotModified struct {
}

func (o *InterCityBusAPIRouteFare2072NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RouteFare/InterCity][%d] interCityBusApiRouteFare2072NotModified ", 304)
}

func (o *InterCityBusAPIRouteFare2072NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
