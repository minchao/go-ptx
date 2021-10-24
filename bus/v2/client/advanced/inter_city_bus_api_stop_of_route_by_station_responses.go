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

// InterCityBusAPIStopOfRouteByStationReader is a Reader for the InterCityBusAPIStopOfRouteByStation structure.
type InterCityBusAPIStopOfRouteByStationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIStopOfRouteByStationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIStopOfRouteByStationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIStopOfRouteByStationStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIStopOfRouteByStationNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIStopOfRouteByStationOK creates a InterCityBusAPIStopOfRouteByStationOK with default headers values
func NewInterCityBusAPIStopOfRouteByStationOK() *InterCityBusAPIStopOfRouteByStationOK {
	return &InterCityBusAPIStopOfRouteByStationOK{}
}

/* InterCityBusAPIStopOfRouteByStationOK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIStopOfRouteByStationOK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusStopOfRoute
}

func (o *InterCityBusAPIStopOfRouteByStationOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/StopOfRoute/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiStopOfRouteByStationOK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIStopOfRouteByStationOK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusStopOfRoute {
	return o.Payload
}

func (o *InterCityBusAPIStopOfRouteByStationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIStopOfRouteByStationStatus299 creates a InterCityBusAPIStopOfRouteByStationStatus299 with default headers values
func NewInterCityBusAPIStopOfRouteByStationStatus299() *InterCityBusAPIStopOfRouteByStationStatus299 {
	return &InterCityBusAPIStopOfRouteByStationStatus299{}
}

/* InterCityBusAPIStopOfRouteByStationStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIStopOfRouteByStationStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIStopOfRouteByStationStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/StopOfRoute/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiStopOfRouteByStationStatus299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIStopOfRouteByStationStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIStopOfRouteByStationStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIStopOfRouteByStationNotModified creates a InterCityBusAPIStopOfRouteByStationNotModified with default headers values
func NewInterCityBusAPIStopOfRouteByStationNotModified() *InterCityBusAPIStopOfRouteByStationNotModified {
	return &InterCityBusAPIStopOfRouteByStationNotModified{}
}

/* InterCityBusAPIStopOfRouteByStationNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIStopOfRouteByStationNotModified struct {
}

func (o *InterCityBusAPIStopOfRouteByStationNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/StopOfRoute/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiStopOfRouteByStationNotModified ", 304)
}

func (o *InterCityBusAPIStopOfRouteByStationNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
