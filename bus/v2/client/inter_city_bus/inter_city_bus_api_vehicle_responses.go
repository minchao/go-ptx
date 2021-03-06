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

// InterCityBusAPIVehicleReader is a Reader for the InterCityBusAPIVehicle structure.
type InterCityBusAPIVehicleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIVehicleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIVehicleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIVehicleStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIVehicleNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIVehicleOK creates a InterCityBusAPIVehicleOK with default headers values
func NewInterCityBusAPIVehicleOK() *InterCityBusAPIVehicleOK {
	return &InterCityBusAPIVehicleOK{}
}

/* InterCityBusAPIVehicleOK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIVehicleOK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusVehicleInfo
}

func (o *InterCityBusAPIVehicleOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Vehicle/InterCity][%d] interCityBusApiVehicleOK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIVehicleOK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusVehicleInfo {
	return o.Payload
}

func (o *InterCityBusAPIVehicleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIVehicleStatus299 creates a InterCityBusAPIVehicleStatus299 with default headers values
func NewInterCityBusAPIVehicleStatus299() *InterCityBusAPIVehicleStatus299 {
	return &InterCityBusAPIVehicleStatus299{}
}

/* InterCityBusAPIVehicleStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIVehicleStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIVehicleStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Vehicle/InterCity][%d] interCityBusApiVehicleStatus299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIVehicleStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIVehicleStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIVehicleNotModified creates a InterCityBusAPIVehicleNotModified with default headers values
func NewInterCityBusAPIVehicleNotModified() *InterCityBusAPIVehicleNotModified {
	return &InterCityBusAPIVehicleNotModified{}
}

/* InterCityBusAPIVehicleNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIVehicleNotModified struct {
}

func (o *InterCityBusAPIVehicleNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Vehicle/InterCity][%d] interCityBusApiVehicleNotModified ", 304)
}

func (o *InterCityBusAPIVehicleNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
