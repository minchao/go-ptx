// Code generated by go-swagger; DO NOT EDIT.

package advanced

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v2/models"
)

// BusAPIShapeByOperatorReader is a Reader for the BusAPIShapeByOperator structure.
type BusAPIShapeByOperatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BusAPIShapeByOperatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBusAPIShapeByOperatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewBusAPIShapeByOperatorStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewBusAPIShapeByOperatorNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBusAPIShapeByOperatorOK creates a BusAPIShapeByOperatorOK with default headers values
func NewBusAPIShapeByOperatorOK() *BusAPIShapeByOperatorOK {
	return &BusAPIShapeByOperatorOK{}
}

/* BusAPIShapeByOperatorOK describes a response with status code 200, with default header values.

Success
*/
type BusAPIShapeByOperatorOK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusShape
}

func (o *BusAPIShapeByOperatorOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Shape/OperatorNo/{OperatorNo}][%d] busApiShapeByOperatorOK  %+v", 200, o.Payload)
}
func (o *BusAPIShapeByOperatorOK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusShape {
	return o.Payload
}

func (o *BusAPIShapeByOperatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBusAPIShapeByOperatorStatus299 creates a BusAPIShapeByOperatorStatus299 with default headers values
func NewBusAPIShapeByOperatorStatus299() *BusAPIShapeByOperatorStatus299 {
	return &BusAPIShapeByOperatorStatus299{}
}

/* BusAPIShapeByOperatorStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type BusAPIShapeByOperatorStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *BusAPIShapeByOperatorStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Shape/OperatorNo/{OperatorNo}][%d] busApiShapeByOperatorStatus299  %+v", 299, o.Payload)
}
func (o *BusAPIShapeByOperatorStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *BusAPIShapeByOperatorStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBusAPIShapeByOperatorNotModified creates a BusAPIShapeByOperatorNotModified with default headers values
func NewBusAPIShapeByOperatorNotModified() *BusAPIShapeByOperatorNotModified {
	return &BusAPIShapeByOperatorNotModified{}
}

/* BusAPIShapeByOperatorNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type BusAPIShapeByOperatorNotModified struct {
}

func (o *BusAPIShapeByOperatorNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Shape/OperatorNo/{OperatorNo}][%d] busApiShapeByOperatorNotModified ", 304)
}

func (o *BusAPIShapeByOperatorNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
