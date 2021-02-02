// Code generated by go-swagger; DO NOT EDIT.

package city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v2/models"
)

// CityBusAPIDisplayStopOfRoute1Reader is a Reader for the CityBusAPIDisplayStopOfRoute1 structure.
type CityBusAPIDisplayStopOfRoute1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIDisplayStopOfRoute1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIDisplayStopOfRoute1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIDisplayStopOfRoute1Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIDisplayStopOfRoute1NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIDisplayStopOfRoute1OK creates a CityBusAPIDisplayStopOfRoute1OK with default headers values
func NewCityBusAPIDisplayStopOfRoute1OK() *CityBusAPIDisplayStopOfRoute1OK {
	return &CityBusAPIDisplayStopOfRoute1OK{}
}

/* CityBusAPIDisplayStopOfRoute1OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIDisplayStopOfRoute1OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusDisplayStopOfRoute
}

func (o *CityBusAPIDisplayStopOfRoute1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/DisplayStopOfRoute/City/{City}/{RouteName}][%d] cityBusApiDisplayStopOfRoute1OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIDisplayStopOfRoute1OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusDisplayStopOfRoute {
	return o.Payload
}

func (o *CityBusAPIDisplayStopOfRoute1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIDisplayStopOfRoute1Status299 creates a CityBusAPIDisplayStopOfRoute1Status299 with default headers values
func NewCityBusAPIDisplayStopOfRoute1Status299() *CityBusAPIDisplayStopOfRoute1Status299 {
	return &CityBusAPIDisplayStopOfRoute1Status299{}
}

/* CityBusAPIDisplayStopOfRoute1Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIDisplayStopOfRoute1Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIDisplayStopOfRoute1Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/DisplayStopOfRoute/City/{City}/{RouteName}][%d] cityBusApiDisplayStopOfRoute1Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIDisplayStopOfRoute1Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIDisplayStopOfRoute1Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIDisplayStopOfRoute1NotModified creates a CityBusAPIDisplayStopOfRoute1NotModified with default headers values
func NewCityBusAPIDisplayStopOfRoute1NotModified() *CityBusAPIDisplayStopOfRoute1NotModified {
	return &CityBusAPIDisplayStopOfRoute1NotModified{}
}

/* CityBusAPIDisplayStopOfRoute1NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIDisplayStopOfRoute1NotModified struct {
}

func (o *CityBusAPIDisplayStopOfRoute1NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/DisplayStopOfRoute/City/{City}/{RouteName}][%d] cityBusApiDisplayStopOfRoute1NotModified ", 304)
}

func (o *CityBusAPIDisplayStopOfRoute1NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
