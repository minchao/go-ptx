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

// CityBusAPIS2STravelTimeDetail2056Reader is a Reader for the CityBusAPIS2STravelTimeDetail2056 structure.
type CityBusAPIS2STravelTimeDetail2056Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIS2STravelTimeDetail2056Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIS2STravelTimeDetail2056OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIS2STravelTimeDetail2056Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIS2STravelTimeDetail2056NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIS2STravelTimeDetail2056OK creates a CityBusAPIS2STravelTimeDetail2056OK with default headers values
func NewCityBusAPIS2STravelTimeDetail2056OK() *CityBusAPIS2STravelTimeDetail2056OK {
	return &CityBusAPIS2STravelTimeDetail2056OK{}
}

/* CityBusAPIS2STravelTimeDetail2056OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIS2STravelTimeDetail2056OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusS2STravelTime
}

func (o *CityBusAPIS2STravelTimeDetail2056OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/S2STravelTime/City/{City}/{RouteID}][%d] cityBusApiS2STravelTimeDetail2056OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIS2STravelTimeDetail2056OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusS2STravelTime {
	return o.Payload
}

func (o *CityBusAPIS2STravelTimeDetail2056OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIS2STravelTimeDetail2056Status299 creates a CityBusAPIS2STravelTimeDetail2056Status299 with default headers values
func NewCityBusAPIS2STravelTimeDetail2056Status299() *CityBusAPIS2STravelTimeDetail2056Status299 {
	return &CityBusAPIS2STravelTimeDetail2056Status299{}
}

/* CityBusAPIS2STravelTimeDetail2056Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIS2STravelTimeDetail2056Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIS2STravelTimeDetail2056Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/S2STravelTime/City/{City}/{RouteID}][%d] cityBusApiS2STravelTimeDetail2056Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIS2STravelTimeDetail2056Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIS2STravelTimeDetail2056Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIS2STravelTimeDetail2056NotModified creates a CityBusAPIS2STravelTimeDetail2056NotModified with default headers values
func NewCityBusAPIS2STravelTimeDetail2056NotModified() *CityBusAPIS2STravelTimeDetail2056NotModified {
	return &CityBusAPIS2STravelTimeDetail2056NotModified{}
}

/* CityBusAPIS2STravelTimeDetail2056NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIS2STravelTimeDetail2056NotModified struct {
}

func (o *CityBusAPIS2STravelTimeDetail2056NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/S2STravelTime/City/{City}/{RouteID}][%d] cityBusApiS2STravelTimeDetail2056NotModified ", 304)
}

func (o *CityBusAPIS2STravelTimeDetail2056NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}