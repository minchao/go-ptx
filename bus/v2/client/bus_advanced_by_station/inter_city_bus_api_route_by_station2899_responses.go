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

// InterCityBusAPIRouteByStation2899Reader is a Reader for the InterCityBusAPIRouteByStation2899 structure.
type InterCityBusAPIRouteByStation2899Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIRouteByStation2899Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIRouteByStation2899OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIRouteByStation2899Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIRouteByStation2899NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIRouteByStation2899OK creates a InterCityBusAPIRouteByStation2899OK with default headers values
func NewInterCityBusAPIRouteByStation2899OK() *InterCityBusAPIRouteByStation2899OK {
	return &InterCityBusAPIRouteByStation2899OK{}
}

/* InterCityBusAPIRouteByStation2899OK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIRouteByStation2899OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusRoute
}

func (o *InterCityBusAPIRouteByStation2899OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiRouteByStation2899OK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIRouteByStation2899OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusRoute {
	return o.Payload
}

func (o *InterCityBusAPIRouteByStation2899OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIRouteByStation2899Status299 creates a InterCityBusAPIRouteByStation2899Status299 with default headers values
func NewInterCityBusAPIRouteByStation2899Status299() *InterCityBusAPIRouteByStation2899Status299 {
	return &InterCityBusAPIRouteByStation2899Status299{}
}

/* InterCityBusAPIRouteByStation2899Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIRouteByStation2899Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIRouteByStation2899Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiRouteByStation2899Status299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIRouteByStation2899Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIRouteByStation2899Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIRouteByStation2899NotModified creates a InterCityBusAPIRouteByStation2899NotModified with default headers values
func NewInterCityBusAPIRouteByStation2899NotModified() *InterCityBusAPIRouteByStation2899NotModified {
	return &InterCityBusAPIRouteByStation2899NotModified{}
}

/* InterCityBusAPIRouteByStation2899NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIRouteByStation2899NotModified struct {
}

func (o *InterCityBusAPIRouteByStation2899NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiRouteByStation2899NotModified ", 304)
}

func (o *InterCityBusAPIRouteByStation2899NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
