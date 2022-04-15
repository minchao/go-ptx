// Code generated by go-swagger; DO NOT EDIT.

package city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v3/models"
)

// CityBusAPIRoute3003Reader is a Reader for the CityBusAPIRoute3003 structure.
type CityBusAPIRoute3003Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRoute3003Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRoute3003OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIRoute3003Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIRoute3003NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIRoute3003OK creates a CityBusAPIRoute3003OK with default headers values
func NewCityBusAPIRoute3003OK() *CityBusAPIRoute3003OK {
	return &CityBusAPIRoute3003OK{}
}

/* CityBusAPIRoute3003OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIRoute3003OK struct {
	Payload *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Route
}

func (o *CityBusAPIRoute3003OK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Route/City/{City}][%d] cityBusApiRoute3003OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIRoute3003OK) GetPayload() *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Route {
	return o.Payload
}

func (o *CityBusAPIRoute3003OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Route)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRoute3003Status299 creates a CityBusAPIRoute3003Status299 with default headers values
func NewCityBusAPIRoute3003Status299() *CityBusAPIRoute3003Status299 {
	return &CityBusAPIRoute3003Status299{}
}

/* CityBusAPIRoute3003Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIRoute3003Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIRoute3003Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Route/City/{City}][%d] cityBusApiRoute3003Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIRoute3003Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIRoute3003Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRoute3003NotModified creates a CityBusAPIRoute3003NotModified with default headers values
func NewCityBusAPIRoute3003NotModified() *CityBusAPIRoute3003NotModified {
	return &CityBusAPIRoute3003NotModified{}
}

/* CityBusAPIRoute3003NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIRoute3003NotModified struct {
}

func (o *CityBusAPIRoute3003NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Route/City/{City}][%d] cityBusApiRoute3003NotModified ", 304)
}

func (o *CityBusAPIRoute3003NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}