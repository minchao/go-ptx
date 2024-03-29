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

// CityBusAPIStop3001Reader is a Reader for the CityBusAPIStop3001 structure.
type CityBusAPIStop3001Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIStop3001Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIStop3001OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIStop3001Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIStop3001NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIStop3001OK creates a CityBusAPIStop3001OK with default headers values
func NewCityBusAPIStop3001OK() *CityBusAPIStop3001OK {
	return &CityBusAPIStop3001OK{}
}

/* CityBusAPIStop3001OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIStop3001OK struct {
	Payload *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop
}

func (o *CityBusAPIStop3001OK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Stop/City/{City}][%d] cityBusApiStop3001OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIStop3001OK) GetPayload() *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop {
	return o.Payload
}

func (o *CityBusAPIStop3001OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIStop3001Status299 creates a CityBusAPIStop3001Status299 with default headers values
func NewCityBusAPIStop3001Status299() *CityBusAPIStop3001Status299 {
	return &CityBusAPIStop3001Status299{}
}

/* CityBusAPIStop3001Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIStop3001Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIStop3001Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Stop/City/{City}][%d] cityBusApiStop3001Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIStop3001Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIStop3001Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIStop3001NotModified creates a CityBusAPIStop3001NotModified with default headers values
func NewCityBusAPIStop3001NotModified() *CityBusAPIStop3001NotModified {
	return &CityBusAPIStop3001NotModified{}
}

/* CityBusAPIStop3001NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIStop3001NotModified struct {
}

func (o *CityBusAPIStop3001NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Stop/City/{City}][%d] cityBusApiStop3001NotModified ", 304)
}

func (o *CityBusAPIStop3001NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
