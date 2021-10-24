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

// CityBusAPIStopOfRouteByOperatorReader is a Reader for the CityBusAPIStopOfRouteByOperator structure.
type CityBusAPIStopOfRouteByOperatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIStopOfRouteByOperatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIStopOfRouteByOperatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIStopOfRouteByOperatorStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIStopOfRouteByOperatorNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIStopOfRouteByOperatorOK creates a CityBusAPIStopOfRouteByOperatorOK with default headers values
func NewCityBusAPIStopOfRouteByOperatorOK() *CityBusAPIStopOfRouteByOperatorOK {
	return &CityBusAPIStopOfRouteByOperatorOK{}
}

/* CityBusAPIStopOfRouteByOperatorOK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIStopOfRouteByOperatorOK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusStopOfRoute
}

func (o *CityBusAPIStopOfRouteByOperatorOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/StopOfRoute/City/{City}/OperatorNo/{OperatorNo}][%d] cityBusApiStopOfRouteByOperatorOK  %+v", 200, o.Payload)
}
func (o *CityBusAPIStopOfRouteByOperatorOK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusStopOfRoute {
	return o.Payload
}

func (o *CityBusAPIStopOfRouteByOperatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIStopOfRouteByOperatorStatus299 creates a CityBusAPIStopOfRouteByOperatorStatus299 with default headers values
func NewCityBusAPIStopOfRouteByOperatorStatus299() *CityBusAPIStopOfRouteByOperatorStatus299 {
	return &CityBusAPIStopOfRouteByOperatorStatus299{}
}

/* CityBusAPIStopOfRouteByOperatorStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIStopOfRouteByOperatorStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIStopOfRouteByOperatorStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/StopOfRoute/City/{City}/OperatorNo/{OperatorNo}][%d] cityBusApiStopOfRouteByOperatorStatus299  %+v", 299, o.Payload)
}
func (o *CityBusAPIStopOfRouteByOperatorStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIStopOfRouteByOperatorStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIStopOfRouteByOperatorNotModified creates a CityBusAPIStopOfRouteByOperatorNotModified with default headers values
func NewCityBusAPIStopOfRouteByOperatorNotModified() *CityBusAPIStopOfRouteByOperatorNotModified {
	return &CityBusAPIStopOfRouteByOperatorNotModified{}
}

/* CityBusAPIStopOfRouteByOperatorNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIStopOfRouteByOperatorNotModified struct {
}

func (o *CityBusAPIStopOfRouteByOperatorNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/StopOfRoute/City/{City}/OperatorNo/{OperatorNo}][%d] cityBusApiStopOfRouteByOperatorNotModified ", 304)
}

func (o *CityBusAPIStopOfRouteByOperatorNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
