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

// CityBusAPIFirstLastTripInfo2Reader is a Reader for the CityBusAPIFirstLastTripInfo2 structure.
type CityBusAPIFirstLastTripInfo2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIFirstLastTripInfo2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIFirstLastTripInfo2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIFirstLastTripInfo2Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIFirstLastTripInfo2NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIFirstLastTripInfo2OK creates a CityBusAPIFirstLastTripInfo2OK with default headers values
func NewCityBusAPIFirstLastTripInfo2OK() *CityBusAPIFirstLastTripInfo2OK {
	return &CityBusAPIFirstLastTripInfo2OK{}
}

/* CityBusAPIFirstLastTripInfo2OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIFirstLastTripInfo2OK struct {
	Payload *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3FirstLastTripInfo
}

func (o *CityBusAPIFirstLastTripInfo2OK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/FirstLastTripInfo/City/{City}][%d] cityBusApiFirstLastTripInfo2OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIFirstLastTripInfo2OK) GetPayload() *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3FirstLastTripInfo {
	return o.Payload
}

func (o *CityBusAPIFirstLastTripInfo2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3FirstLastTripInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIFirstLastTripInfo2Status299 creates a CityBusAPIFirstLastTripInfo2Status299 with default headers values
func NewCityBusAPIFirstLastTripInfo2Status299() *CityBusAPIFirstLastTripInfo2Status299 {
	return &CityBusAPIFirstLastTripInfo2Status299{}
}

/* CityBusAPIFirstLastTripInfo2Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIFirstLastTripInfo2Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIFirstLastTripInfo2Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/FirstLastTripInfo/City/{City}][%d] cityBusApiFirstLastTripInfo2Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIFirstLastTripInfo2Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIFirstLastTripInfo2Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIFirstLastTripInfo2NotModified creates a CityBusAPIFirstLastTripInfo2NotModified with default headers values
func NewCityBusAPIFirstLastTripInfo2NotModified() *CityBusAPIFirstLastTripInfo2NotModified {
	return &CityBusAPIFirstLastTripInfo2NotModified{}
}

/* CityBusAPIFirstLastTripInfo2NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIFirstLastTripInfo2NotModified struct {
}

func (o *CityBusAPIFirstLastTripInfo2NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/FirstLastTripInfo/City/{City}][%d] cityBusApiFirstLastTripInfo2NotModified ", 304)
}

func (o *CityBusAPIFirstLastTripInfo2NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
