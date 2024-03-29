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

// CityBusAPIEstimatedTimeOfArrival3017Reader is a Reader for the CityBusAPIEstimatedTimeOfArrival3017 structure.
type CityBusAPIEstimatedTimeOfArrival3017Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIEstimatedTimeOfArrival3017Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIEstimatedTimeOfArrival3017OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIEstimatedTimeOfArrival3017Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIEstimatedTimeOfArrival3017NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIEstimatedTimeOfArrival3017OK creates a CityBusAPIEstimatedTimeOfArrival3017OK with default headers values
func NewCityBusAPIEstimatedTimeOfArrival3017OK() *CityBusAPIEstimatedTimeOfArrival3017OK {
	return &CityBusAPIEstimatedTimeOfArrival3017OK{}
}

/* CityBusAPIEstimatedTimeOfArrival3017OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIEstimatedTimeOfArrival3017OK struct {
	Payload *models.PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3N1Data
}

func (o *CityBusAPIEstimatedTimeOfArrival3017OK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/EstimatedTimeOfArrival/City/{City}][%d] cityBusApiEstimatedTimeOfArrival3017OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIEstimatedTimeOfArrival3017OK) GetPayload() *models.PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3N1Data {
	return o.Payload
}

func (o *CityBusAPIEstimatedTimeOfArrival3017OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3N1Data)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIEstimatedTimeOfArrival3017Status299 creates a CityBusAPIEstimatedTimeOfArrival3017Status299 with default headers values
func NewCityBusAPIEstimatedTimeOfArrival3017Status299() *CityBusAPIEstimatedTimeOfArrival3017Status299 {
	return &CityBusAPIEstimatedTimeOfArrival3017Status299{}
}

/* CityBusAPIEstimatedTimeOfArrival3017Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIEstimatedTimeOfArrival3017Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIEstimatedTimeOfArrival3017Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/EstimatedTimeOfArrival/City/{City}][%d] cityBusApiEstimatedTimeOfArrival3017Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIEstimatedTimeOfArrival3017Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIEstimatedTimeOfArrival3017Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIEstimatedTimeOfArrival3017NotModified creates a CityBusAPIEstimatedTimeOfArrival3017NotModified with default headers values
func NewCityBusAPIEstimatedTimeOfArrival3017NotModified() *CityBusAPIEstimatedTimeOfArrival3017NotModified {
	return &CityBusAPIEstimatedTimeOfArrival3017NotModified{}
}

/* CityBusAPIEstimatedTimeOfArrival3017NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIEstimatedTimeOfArrival3017NotModified struct {
}

func (o *CityBusAPIEstimatedTimeOfArrival3017NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/EstimatedTimeOfArrival/City/{City}][%d] cityBusApiEstimatedTimeOfArrival3017NotModified ", 304)
}

func (o *CityBusAPIEstimatedTimeOfArrival3017NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
