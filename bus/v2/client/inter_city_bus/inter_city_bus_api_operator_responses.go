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

// InterCityBusAPIOperatorReader is a Reader for the InterCityBusAPIOperator structure.
type InterCityBusAPIOperatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIOperatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIOperatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIOperatorStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIOperatorNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIOperatorOK creates a InterCityBusAPIOperatorOK with default headers values
func NewInterCityBusAPIOperatorOK() *InterCityBusAPIOperatorOK {
	return &InterCityBusAPIOperatorOK{}
}

/* InterCityBusAPIOperatorOK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIOperatorOK struct {
	Payload []*models.PTXServiceDTOSharedSpecificationV2BaseOperator
}

func (o *InterCityBusAPIOperatorOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Operator/InterCity][%d] interCityBusApiOperatorOK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIOperatorOK) GetPayload() []*models.PTXServiceDTOSharedSpecificationV2BaseOperator {
	return o.Payload
}

func (o *InterCityBusAPIOperatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIOperatorStatus299 creates a InterCityBusAPIOperatorStatus299 with default headers values
func NewInterCityBusAPIOperatorStatus299() *InterCityBusAPIOperatorStatus299 {
	return &InterCityBusAPIOperatorStatus299{}
}

/* InterCityBusAPIOperatorStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIOperatorStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIOperatorStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Operator/InterCity][%d] interCityBusApiOperatorStatus299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIOperatorStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIOperatorStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIOperatorNotModified creates a InterCityBusAPIOperatorNotModified with default headers values
func NewInterCityBusAPIOperatorNotModified() *InterCityBusAPIOperatorNotModified {
	return &InterCityBusAPIOperatorNotModified{}
}

/* InterCityBusAPIOperatorNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIOperatorNotModified struct {
}

func (o *InterCityBusAPIOperatorNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Operator/InterCity][%d] interCityBusApiOperatorNotModified ", 304)
}

func (o *InterCityBusAPIOperatorNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
