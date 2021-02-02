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

// CityBusAPISchedule2Reader is a Reader for the CityBusAPISchedule2 structure.
type CityBusAPISchedule2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPISchedule2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPISchedule2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPISchedule2Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPISchedule2NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPISchedule2OK creates a CityBusAPISchedule2OK with default headers values
func NewCityBusAPISchedule2OK() *CityBusAPISchedule2OK {
	return &CityBusAPISchedule2OK{}
}

/* CityBusAPISchedule2OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPISchedule2OK struct {
	Payload *models.PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule
}

func (o *CityBusAPISchedule2OK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Schedule/City/{City}][%d] cityBusApiSchedule2OK  %+v", 200, o.Payload)
}
func (o *CityBusAPISchedule2OK) GetPayload() *models.PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule {
	return o.Payload
}

func (o *CityBusAPISchedule2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPISchedule2Status299 creates a CityBusAPISchedule2Status299 with default headers values
func NewCityBusAPISchedule2Status299() *CityBusAPISchedule2Status299 {
	return &CityBusAPISchedule2Status299{}
}

/* CityBusAPISchedule2Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPISchedule2Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPISchedule2Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Schedule/City/{City}][%d] cityBusApiSchedule2Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPISchedule2Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPISchedule2Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPISchedule2NotModified creates a CityBusAPISchedule2NotModified with default headers values
func NewCityBusAPISchedule2NotModified() *CityBusAPISchedule2NotModified {
	return &CityBusAPISchedule2NotModified{}
}

/* CityBusAPISchedule2NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPISchedule2NotModified struct {
}

func (o *CityBusAPISchedule2NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Schedule/City/{City}][%d] cityBusApiSchedule2NotModified ", 304)
}

func (o *CityBusAPISchedule2NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
