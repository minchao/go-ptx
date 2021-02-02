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

// InterCityBusAPIStopReader is a Reader for the InterCityBusAPIStop structure.
type InterCityBusAPIStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIStopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIStopStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIStopNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIStopOK creates a InterCityBusAPIStopOK with default headers values
func NewInterCityBusAPIStopOK() *InterCityBusAPIStopOK {
	return &InterCityBusAPIStopOK{}
}

/* InterCityBusAPIStopOK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIStopOK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusStop
}

func (o *InterCityBusAPIStopOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Stop/InterCity][%d] interCityBusApiStopOK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIStopOK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusStop {
	return o.Payload
}

func (o *InterCityBusAPIStopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIStopStatus299 creates a InterCityBusAPIStopStatus299 with default headers values
func NewInterCityBusAPIStopStatus299() *InterCityBusAPIStopStatus299 {
	return &InterCityBusAPIStopStatus299{}
}

/* InterCityBusAPIStopStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIStopStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIStopStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Stop/InterCity][%d] interCityBusApiStopStatus299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIStopStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIStopStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIStopNotModified creates a InterCityBusAPIStopNotModified with default headers values
func NewInterCityBusAPIStopNotModified() *InterCityBusAPIStopNotModified {
	return &InterCityBusAPIStopNotModified{}
}

/* InterCityBusAPIStopNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIStopNotModified struct {
}

func (o *InterCityBusAPIStopNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Stop/InterCity][%d] interCityBusApiStopNotModified ", 304)
}

func (o *InterCityBusAPIStopNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
