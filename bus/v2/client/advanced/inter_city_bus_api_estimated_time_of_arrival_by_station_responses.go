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

// InterCityBusAPIEstimatedTimeOfArrivalByStationReader is a Reader for the InterCityBusAPIEstimatedTimeOfArrivalByStation structure.
type InterCityBusAPIEstimatedTimeOfArrivalByStationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIEstimatedTimeOfArrivalByStationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIEstimatedTimeOfArrivalByStationStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIEstimatedTimeOfArrivalByStationNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIEstimatedTimeOfArrivalByStationOK creates a InterCityBusAPIEstimatedTimeOfArrivalByStationOK with default headers values
func NewInterCityBusAPIEstimatedTimeOfArrivalByStationOK() *InterCityBusAPIEstimatedTimeOfArrivalByStationOK {
	return &InterCityBusAPIEstimatedTimeOfArrivalByStationOK{}
}

/* InterCityBusAPIEstimatedTimeOfArrivalByStationOK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIEstimatedTimeOfArrivalByStationOK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusN1EstimateTime
}

func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/EstimatedTimeOfArrival/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiEstimatedTimeOfArrivalByStationOK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationOK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusN1EstimateTime {
	return o.Payload
}

func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIEstimatedTimeOfArrivalByStationStatus299 creates a InterCityBusAPIEstimatedTimeOfArrivalByStationStatus299 with default headers values
func NewInterCityBusAPIEstimatedTimeOfArrivalByStationStatus299() *InterCityBusAPIEstimatedTimeOfArrivalByStationStatus299 {
	return &InterCityBusAPIEstimatedTimeOfArrivalByStationStatus299{}
}

/* InterCityBusAPIEstimatedTimeOfArrivalByStationStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIEstimatedTimeOfArrivalByStationStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/EstimatedTimeOfArrival/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiEstimatedTimeOfArrivalByStationStatus299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIEstimatedTimeOfArrivalByStationNotModified creates a InterCityBusAPIEstimatedTimeOfArrivalByStationNotModified with default headers values
func NewInterCityBusAPIEstimatedTimeOfArrivalByStationNotModified() *InterCityBusAPIEstimatedTimeOfArrivalByStationNotModified {
	return &InterCityBusAPIEstimatedTimeOfArrivalByStationNotModified{}
}

/* InterCityBusAPIEstimatedTimeOfArrivalByStationNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIEstimatedTimeOfArrivalByStationNotModified struct {
}

func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/EstimatedTimeOfArrival/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiEstimatedTimeOfArrivalByStationNotModified ", 304)
}

func (o *InterCityBusAPIEstimatedTimeOfArrivalByStationNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
