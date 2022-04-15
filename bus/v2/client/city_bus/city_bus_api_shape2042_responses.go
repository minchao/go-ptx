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

// CityBusAPIShape2042Reader is a Reader for the CityBusAPIShape2042 structure.
type CityBusAPIShape2042Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIShape2042Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIShape2042OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIShape2042Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIShape2042NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIShape2042OK creates a CityBusAPIShape2042OK with default headers values
func NewCityBusAPIShape2042OK() *CityBusAPIShape2042OK {
	return &CityBusAPIShape2042OK{}
}

/* CityBusAPIShape2042OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIShape2042OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusShape
}

func (o *CityBusAPIShape2042OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Shape/City/{City}][%d] cityBusApiShape2042OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIShape2042OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusShape {
	return o.Payload
}

func (o *CityBusAPIShape2042OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIShape2042Status299 creates a CityBusAPIShape2042Status299 with default headers values
func NewCityBusAPIShape2042Status299() *CityBusAPIShape2042Status299 {
	return &CityBusAPIShape2042Status299{}
}

/* CityBusAPIShape2042Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIShape2042Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIShape2042Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Shape/City/{City}][%d] cityBusApiShape2042Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIShape2042Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIShape2042Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIShape2042NotModified creates a CityBusAPIShape2042NotModified with default headers values
func NewCityBusAPIShape2042NotModified() *CityBusAPIShape2042NotModified {
	return &CityBusAPIShape2042NotModified{}
}

/* CityBusAPIShape2042NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIShape2042NotModified struct {
}

func (o *CityBusAPIShape2042NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Shape/City/{City}][%d] cityBusApiShape2042NotModified ", 304)
}

func (o *CityBusAPIShape2042NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}