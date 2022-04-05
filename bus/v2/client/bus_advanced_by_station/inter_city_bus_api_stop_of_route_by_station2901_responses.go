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

// InterCityBusAPIStopOfRouteByStation2901Reader is a Reader for the InterCityBusAPIStopOfRouteByStation2901 structure.
type InterCityBusAPIStopOfRouteByStation2901Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIStopOfRouteByStation2901Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIStopOfRouteByStation2901OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIStopOfRouteByStation2901Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIStopOfRouteByStation2901NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIStopOfRouteByStation2901OK creates a InterCityBusAPIStopOfRouteByStation2901OK with default headers values
func NewInterCityBusAPIStopOfRouteByStation2901OK() *InterCityBusAPIStopOfRouteByStation2901OK {
	return &InterCityBusAPIStopOfRouteByStation2901OK{}
}

/* InterCityBusAPIStopOfRouteByStation2901OK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIStopOfRouteByStation2901OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusStopOfRoute
}

func (o *InterCityBusAPIStopOfRouteByStation2901OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/StopOfRoute/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiStopOfRouteByStation2901OK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIStopOfRouteByStation2901OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusStopOfRoute {
	return o.Payload
}

func (o *InterCityBusAPIStopOfRouteByStation2901OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIStopOfRouteByStation2901Status299 creates a InterCityBusAPIStopOfRouteByStation2901Status299 with default headers values
func NewInterCityBusAPIStopOfRouteByStation2901Status299() *InterCityBusAPIStopOfRouteByStation2901Status299 {
	return &InterCityBusAPIStopOfRouteByStation2901Status299{}
}

/* InterCityBusAPIStopOfRouteByStation2901Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIStopOfRouteByStation2901Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIStopOfRouteByStation2901Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/StopOfRoute/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiStopOfRouteByStation2901Status299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIStopOfRouteByStation2901Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIStopOfRouteByStation2901Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIStopOfRouteByStation2901NotModified creates a InterCityBusAPIStopOfRouteByStation2901NotModified with default headers values
func NewInterCityBusAPIStopOfRouteByStation2901NotModified() *InterCityBusAPIStopOfRouteByStation2901NotModified {
	return &InterCityBusAPIStopOfRouteByStation2901NotModified{}
}

/* InterCityBusAPIStopOfRouteByStation2901NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIStopOfRouteByStation2901NotModified struct {
}

func (o *InterCityBusAPIStopOfRouteByStation2901NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/StopOfRoute/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiStopOfRouteByStation2901NotModified ", 304)
}

func (o *InterCityBusAPIStopOfRouteByStation2901NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
