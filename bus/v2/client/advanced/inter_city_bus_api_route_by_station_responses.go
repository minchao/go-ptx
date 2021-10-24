// Code generated by go-swagger; DO NOT EDIT.

package advanced

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v2/models"
)

// InterCityBusAPIRouteByStationReader is a Reader for the InterCityBusAPIRouteByStation structure.
type InterCityBusAPIRouteByStationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIRouteByStationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIRouteByStationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIRouteByStationStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIRouteByStationNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIRouteByStationOK creates a InterCityBusAPIRouteByStationOK with default headers values
func NewInterCityBusAPIRouteByStationOK() *InterCityBusAPIRouteByStationOK {
	return &InterCityBusAPIRouteByStationOK{}
}

/* InterCityBusAPIRouteByStationOK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIRouteByStationOK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusRoute
}

func (o *InterCityBusAPIRouteByStationOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiRouteByStationOK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIRouteByStationOK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusRoute {
	return o.Payload
}

func (o *InterCityBusAPIRouteByStationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIRouteByStationStatus299 creates a InterCityBusAPIRouteByStationStatus299 with default headers values
func NewInterCityBusAPIRouteByStationStatus299() *InterCityBusAPIRouteByStationStatus299 {
	return &InterCityBusAPIRouteByStationStatus299{}
}

/* InterCityBusAPIRouteByStationStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIRouteByStationStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIRouteByStationStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiRouteByStationStatus299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIRouteByStationStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIRouteByStationStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIRouteByStationNotModified creates a InterCityBusAPIRouteByStationNotModified with default headers values
func NewInterCityBusAPIRouteByStationNotModified() *InterCityBusAPIRouteByStationNotModified {
	return &InterCityBusAPIRouteByStationNotModified{}
}

/* InterCityBusAPIRouteByStationNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIRouteByStationNotModified struct {
}

func (o *InterCityBusAPIRouteByStationNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiRouteByStationNotModified ", 304)
}

func (o *InterCityBusAPIRouteByStationNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
