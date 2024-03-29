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

// InterCityBusAPINews2070Reader is a Reader for the InterCityBusAPINews2070 structure.
type InterCityBusAPINews2070Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPINews2070Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPINews2070OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPINews2070Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPINews2070NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPINews2070OK creates a InterCityBusAPINews2070OK with default headers values
func NewInterCityBusAPINews2070OK() *InterCityBusAPINews2070OK {
	return &InterCityBusAPINews2070OK{}
}

/* InterCityBusAPINews2070OK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPINews2070OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusNews
}

func (o *InterCityBusAPINews2070OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/News/InterCity][%d] interCityBusApiNews2070OK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPINews2070OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusNews {
	return o.Payload
}

func (o *InterCityBusAPINews2070OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPINews2070Status299 creates a InterCityBusAPINews2070Status299 with default headers values
func NewInterCityBusAPINews2070Status299() *InterCityBusAPINews2070Status299 {
	return &InterCityBusAPINews2070Status299{}
}

/* InterCityBusAPINews2070Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPINews2070Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPINews2070Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/News/InterCity][%d] interCityBusApiNews2070Status299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPINews2070Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPINews2070Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPINews2070NotModified creates a InterCityBusAPINews2070NotModified with default headers values
func NewInterCityBusAPINews2070NotModified() *InterCityBusAPINews2070NotModified {
	return &InterCityBusAPINews2070NotModified{}
}

/* InterCityBusAPINews2070NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPINews2070NotModified struct {
}

func (o *InterCityBusAPINews2070NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/News/InterCity][%d] interCityBusApiNews2070NotModified ", 304)
}

func (o *InterCityBusAPINews2070NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
