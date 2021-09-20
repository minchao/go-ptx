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

// CityBusAPIFirstLastTripInfo1Reader is a Reader for the CityBusAPIFirstLastTripInfo1 structure.
type CityBusAPIFirstLastTripInfo1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIFirstLastTripInfo1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIFirstLastTripInfo1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIFirstLastTripInfo1Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIFirstLastTripInfo1NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIFirstLastTripInfo1OK creates a CityBusAPIFirstLastTripInfo1OK with default headers values
func NewCityBusAPIFirstLastTripInfo1OK() *CityBusAPIFirstLastTripInfo1OK {
	return &CityBusAPIFirstLastTripInfo1OK{}
}

/* CityBusAPIFirstLastTripInfo1OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIFirstLastTripInfo1OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusFirstLastTripInfo
}

func (o *CityBusAPIFirstLastTripInfo1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/FirstLastTripInfo/City/{City}][%d] cityBusApiFirstLastTripInfo1OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIFirstLastTripInfo1OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusFirstLastTripInfo {
	return o.Payload
}

func (o *CityBusAPIFirstLastTripInfo1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIFirstLastTripInfo1Status299 creates a CityBusAPIFirstLastTripInfo1Status299 with default headers values
func NewCityBusAPIFirstLastTripInfo1Status299() *CityBusAPIFirstLastTripInfo1Status299 {
	return &CityBusAPIFirstLastTripInfo1Status299{}
}

/* CityBusAPIFirstLastTripInfo1Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIFirstLastTripInfo1Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIFirstLastTripInfo1Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/FirstLastTripInfo/City/{City}][%d] cityBusApiFirstLastTripInfo1Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIFirstLastTripInfo1Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIFirstLastTripInfo1Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIFirstLastTripInfo1NotModified creates a CityBusAPIFirstLastTripInfo1NotModified with default headers values
func NewCityBusAPIFirstLastTripInfo1NotModified() *CityBusAPIFirstLastTripInfo1NotModified {
	return &CityBusAPIFirstLastTripInfo1NotModified{}
}

/* CityBusAPIFirstLastTripInfo1NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIFirstLastTripInfo1NotModified struct {
}

func (o *CityBusAPIFirstLastTripInfo1NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/FirstLastTripInfo/City/{City}][%d] cityBusApiFirstLastTripInfo1NotModified ", 304)
}

func (o *CityBusAPIFirstLastTripInfo1NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
