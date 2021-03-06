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

// CityBusAPIOperatorReader is a Reader for the CityBusAPIOperator structure.
type CityBusAPIOperatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIOperatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIOperatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIOperatorStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIOperatorNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIOperatorOK creates a CityBusAPIOperatorOK with default headers values
func NewCityBusAPIOperatorOK() *CityBusAPIOperatorOK {
	return &CityBusAPIOperatorOK{}
}

/* CityBusAPIOperatorOK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIOperatorOK struct {
	Payload []*models.PTXServiceDTOSharedSpecificationV2BaseOperator
}

func (o *CityBusAPIOperatorOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Operator/City/{City}][%d] cityBusApiOperatorOK  %+v", 200, o.Payload)
}
func (o *CityBusAPIOperatorOK) GetPayload() []*models.PTXServiceDTOSharedSpecificationV2BaseOperator {
	return o.Payload
}

func (o *CityBusAPIOperatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIOperatorStatus299 creates a CityBusAPIOperatorStatus299 with default headers values
func NewCityBusAPIOperatorStatus299() *CityBusAPIOperatorStatus299 {
	return &CityBusAPIOperatorStatus299{}
}

/* CityBusAPIOperatorStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIOperatorStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIOperatorStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Operator/City/{City}][%d] cityBusApiOperatorStatus299  %+v", 299, o.Payload)
}
func (o *CityBusAPIOperatorStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIOperatorStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIOperatorNotModified creates a CityBusAPIOperatorNotModified with default headers values
func NewCityBusAPIOperatorNotModified() *CityBusAPIOperatorNotModified {
	return &CityBusAPIOperatorNotModified{}
}

/* CityBusAPIOperatorNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIOperatorNotModified struct {
}

func (o *CityBusAPIOperatorNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Operator/City/{City}][%d] cityBusApiOperatorNotModified ", 304)
}

func (o *CityBusAPIOperatorNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
