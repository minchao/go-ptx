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

// CityBusAPIRouteNetwork20551Reader is a Reader for the CityBusAPIRouteNetwork20551 structure.
type CityBusAPIRouteNetwork20551Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRouteNetwork20551Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRouteNetwork20551OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIRouteNetwork20551Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIRouteNetwork20551NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIRouteNetwork20551OK creates a CityBusAPIRouteNetwork20551OK with default headers values
func NewCityBusAPIRouteNetwork20551OK() *CityBusAPIRouteNetwork20551OK {
	return &CityBusAPIRouteNetwork20551OK{}
}

/* CityBusAPIRouteNetwork20551OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIRouteNetwork20551OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusRouteNetwork
}

func (o *CityBusAPIRouteNetwork20551OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RouteNetwork/City/{City}/{RouteName}][%d] cityBusApiRouteNetwork20551OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIRouteNetwork20551OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusRouteNetwork {
	return o.Payload
}

func (o *CityBusAPIRouteNetwork20551OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRouteNetwork20551Status299 creates a CityBusAPIRouteNetwork20551Status299 with default headers values
func NewCityBusAPIRouteNetwork20551Status299() *CityBusAPIRouteNetwork20551Status299 {
	return &CityBusAPIRouteNetwork20551Status299{}
}

/* CityBusAPIRouteNetwork20551Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIRouteNetwork20551Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIRouteNetwork20551Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RouteNetwork/City/{City}/{RouteName}][%d] cityBusApiRouteNetwork20551Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIRouteNetwork20551Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIRouteNetwork20551Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRouteNetwork20551NotModified creates a CityBusAPIRouteNetwork20551NotModified with default headers values
func NewCityBusAPIRouteNetwork20551NotModified() *CityBusAPIRouteNetwork20551NotModified {
	return &CityBusAPIRouteNetwork20551NotModified{}
}

/* CityBusAPIRouteNetwork20551NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIRouteNetwork20551NotModified struct {
}

func (o *CityBusAPIRouteNetwork20551NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RouteNetwork/City/{City}/{RouteName}][%d] cityBusApiRouteNetwork20551NotModified ", 304)
}

func (o *CityBusAPIRouteNetwork20551NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}