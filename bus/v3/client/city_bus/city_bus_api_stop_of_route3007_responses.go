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

// CityBusAPIStopOfRoute3007Reader is a Reader for the CityBusAPIStopOfRoute3007 structure.
type CityBusAPIStopOfRoute3007Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIStopOfRoute3007Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIStopOfRoute3007OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIStopOfRoute3007Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIStopOfRoute3007NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIStopOfRoute3007OK creates a CityBusAPIStopOfRoute3007OK with default headers values
func NewCityBusAPIStopOfRoute3007OK() *CityBusAPIStopOfRoute3007OK {
	return &CityBusAPIStopOfRoute3007OK{}
}

/* CityBusAPIStopOfRoute3007OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIStopOfRoute3007OK struct {
	Payload *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3StopOfRoute
}

func (o *CityBusAPIStopOfRoute3007OK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/StopOfRoute/City/{City}][%d] cityBusApiStopOfRoute3007OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIStopOfRoute3007OK) GetPayload() *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3StopOfRoute {
	return o.Payload
}

func (o *CityBusAPIStopOfRoute3007OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3StopOfRoute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIStopOfRoute3007Status299 creates a CityBusAPIStopOfRoute3007Status299 with default headers values
func NewCityBusAPIStopOfRoute3007Status299() *CityBusAPIStopOfRoute3007Status299 {
	return &CityBusAPIStopOfRoute3007Status299{}
}

/* CityBusAPIStopOfRoute3007Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIStopOfRoute3007Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIStopOfRoute3007Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/StopOfRoute/City/{City}][%d] cityBusApiStopOfRoute3007Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIStopOfRoute3007Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIStopOfRoute3007Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIStopOfRoute3007NotModified creates a CityBusAPIStopOfRoute3007NotModified with default headers values
func NewCityBusAPIStopOfRoute3007NotModified() *CityBusAPIStopOfRoute3007NotModified {
	return &CityBusAPIStopOfRoute3007NotModified{}
}

/* CityBusAPIStopOfRoute3007NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIStopOfRoute3007NotModified struct {
}

func (o *CityBusAPIStopOfRoute3007NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/StopOfRoute/City/{City}][%d] cityBusApiStopOfRoute3007NotModified ", 304)
}

func (o *CityBusAPIStopOfRoute3007NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
