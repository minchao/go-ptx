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

// CityBusAPIRoute1Reader is a Reader for the CityBusAPIRoute1 structure.
type CityBusAPIRoute1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRoute1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRoute1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIRoute1Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIRoute1NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIRoute1OK creates a CityBusAPIRoute1OK with default headers values
func NewCityBusAPIRoute1OK() *CityBusAPIRoute1OK {
	return &CityBusAPIRoute1OK{}
}

/* CityBusAPIRoute1OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIRoute1OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusRoute
}

func (o *CityBusAPIRoute1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/City/{City}/{RouteName}][%d] cityBusApiRoute1OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIRoute1OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusRoute {
	return o.Payload
}

func (o *CityBusAPIRoute1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRoute1Status299 creates a CityBusAPIRoute1Status299 with default headers values
func NewCityBusAPIRoute1Status299() *CityBusAPIRoute1Status299 {
	return &CityBusAPIRoute1Status299{}
}

/* CityBusAPIRoute1Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIRoute1Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIRoute1Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/City/{City}/{RouteName}][%d] cityBusApiRoute1Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIRoute1Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIRoute1Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRoute1NotModified creates a CityBusAPIRoute1NotModified with default headers values
func NewCityBusAPIRoute1NotModified() *CityBusAPIRoute1NotModified {
	return &CityBusAPIRoute1NotModified{}
}

/* CityBusAPIRoute1NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIRoute1NotModified struct {
}

func (o *CityBusAPIRoute1NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/City/{City}/{RouteName}][%d] cityBusApiRoute1NotModified ", 304)
}

func (o *CityBusAPIRoute1NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
