// Code generated by go-swagger; DO NOT EDIT.

package bus_advanced_by_operator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v2/models"
)

// InterCityBusAPIStopOfRouteByOperator2892Reader is a Reader for the InterCityBusAPIStopOfRouteByOperator2892 structure.
type InterCityBusAPIStopOfRouteByOperator2892Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIStopOfRouteByOperator2892Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIStopOfRouteByOperator2892OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIStopOfRouteByOperator2892Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIStopOfRouteByOperator2892NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIStopOfRouteByOperator2892OK creates a InterCityBusAPIStopOfRouteByOperator2892OK with default headers values
func NewInterCityBusAPIStopOfRouteByOperator2892OK() *InterCityBusAPIStopOfRouteByOperator2892OK {
	return &InterCityBusAPIStopOfRouteByOperator2892OK{}
}

/* InterCityBusAPIStopOfRouteByOperator2892OK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIStopOfRouteByOperator2892OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusStopOfRoute
}

func (o *InterCityBusAPIStopOfRouteByOperator2892OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/StopOfRoute/InterCity/OperatorNo/{OperatorNo}][%d] interCityBusApiStopOfRouteByOperator2892OK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIStopOfRouteByOperator2892OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusStopOfRoute {
	return o.Payload
}

func (o *InterCityBusAPIStopOfRouteByOperator2892OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIStopOfRouteByOperator2892Status299 creates a InterCityBusAPIStopOfRouteByOperator2892Status299 with default headers values
func NewInterCityBusAPIStopOfRouteByOperator2892Status299() *InterCityBusAPIStopOfRouteByOperator2892Status299 {
	return &InterCityBusAPIStopOfRouteByOperator2892Status299{}
}

/* InterCityBusAPIStopOfRouteByOperator2892Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIStopOfRouteByOperator2892Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIStopOfRouteByOperator2892Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/StopOfRoute/InterCity/OperatorNo/{OperatorNo}][%d] interCityBusApiStopOfRouteByOperator2892Status299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIStopOfRouteByOperator2892Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIStopOfRouteByOperator2892Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIStopOfRouteByOperator2892NotModified creates a InterCityBusAPIStopOfRouteByOperator2892NotModified with default headers values
func NewInterCityBusAPIStopOfRouteByOperator2892NotModified() *InterCityBusAPIStopOfRouteByOperator2892NotModified {
	return &InterCityBusAPIStopOfRouteByOperator2892NotModified{}
}

/* InterCityBusAPIStopOfRouteByOperator2892NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIStopOfRouteByOperator2892NotModified struct {
}

func (o *InterCityBusAPIStopOfRouteByOperator2892NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/StopOfRoute/InterCity/OperatorNo/{OperatorNo}][%d] interCityBusApiStopOfRouteByOperator2892NotModified ", 304)
}

func (o *InterCityBusAPIStopOfRouteByOperator2892NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
