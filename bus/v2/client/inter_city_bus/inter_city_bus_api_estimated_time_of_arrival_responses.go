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

// InterCityBusAPIEstimatedTimeOfArrivalReader is a Reader for the InterCityBusAPIEstimatedTimeOfArrival structure.
type InterCityBusAPIEstimatedTimeOfArrivalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIEstimatedTimeOfArrivalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIEstimatedTimeOfArrivalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIEstimatedTimeOfArrivalStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIEstimatedTimeOfArrivalNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIEstimatedTimeOfArrivalOK creates a InterCityBusAPIEstimatedTimeOfArrivalOK with default headers values
func NewInterCityBusAPIEstimatedTimeOfArrivalOK() *InterCityBusAPIEstimatedTimeOfArrivalOK {
	return &InterCityBusAPIEstimatedTimeOfArrivalOK{}
}

/* InterCityBusAPIEstimatedTimeOfArrivalOK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIEstimatedTimeOfArrivalOK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusN1EstimateTime
}

func (o *InterCityBusAPIEstimatedTimeOfArrivalOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/EstimatedTimeOfArrival/InterCity][%d] interCityBusApiEstimatedTimeOfArrivalOK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIEstimatedTimeOfArrivalOK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusN1EstimateTime {
	return o.Payload
}

func (o *InterCityBusAPIEstimatedTimeOfArrivalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIEstimatedTimeOfArrivalStatus299 creates a InterCityBusAPIEstimatedTimeOfArrivalStatus299 with default headers values
func NewInterCityBusAPIEstimatedTimeOfArrivalStatus299() *InterCityBusAPIEstimatedTimeOfArrivalStatus299 {
	return &InterCityBusAPIEstimatedTimeOfArrivalStatus299{}
}

/* InterCityBusAPIEstimatedTimeOfArrivalStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIEstimatedTimeOfArrivalStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIEstimatedTimeOfArrivalStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/EstimatedTimeOfArrival/InterCity][%d] interCityBusApiEstimatedTimeOfArrivalStatus299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIEstimatedTimeOfArrivalStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIEstimatedTimeOfArrivalStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIEstimatedTimeOfArrivalNotModified creates a InterCityBusAPIEstimatedTimeOfArrivalNotModified with default headers values
func NewInterCityBusAPIEstimatedTimeOfArrivalNotModified() *InterCityBusAPIEstimatedTimeOfArrivalNotModified {
	return &InterCityBusAPIEstimatedTimeOfArrivalNotModified{}
}

/* InterCityBusAPIEstimatedTimeOfArrivalNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIEstimatedTimeOfArrivalNotModified struct {
}

func (o *InterCityBusAPIEstimatedTimeOfArrivalNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/EstimatedTimeOfArrival/InterCity][%d] interCityBusApiEstimatedTimeOfArrivalNotModified ", 304)
}

func (o *InterCityBusAPIEstimatedTimeOfArrivalNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
