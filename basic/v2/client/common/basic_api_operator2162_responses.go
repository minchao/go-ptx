// Code generated by go-swagger; DO NOT EDIT.

package common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/basic/v2/models"
)

// BasicAPIOperator2162Reader is a Reader for the BasicAPIOperator2162 structure.
type BasicAPIOperator2162Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BasicAPIOperator2162Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBasicAPIOperator2162OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewBasicAPIOperator2162NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBasicAPIOperator2162OK creates a BasicAPIOperator2162OK with default headers values
func NewBasicAPIOperator2162OK() *BasicAPIOperator2162OK {
	return &BasicAPIOperator2162OK{}
}

/* BasicAPIOperator2162OK describes a response with status code 200, with default header values.

Success
*/
type BasicAPIOperator2162OK struct {
	Payload []*models.PTXServiceDTOSharedSpecificationV2BaseOperator
}

func (o *BasicAPIOperator2162OK) Error() string {
	return fmt.Sprintf("[GET /v2/Basic/Operator][%d] basicApiOperator2162OK  %+v", 200, o.Payload)
}
func (o *BasicAPIOperator2162OK) GetPayload() []*models.PTXServiceDTOSharedSpecificationV2BaseOperator {
	return o.Payload
}

func (o *BasicAPIOperator2162OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBasicAPIOperator2162NotModified creates a BasicAPIOperator2162NotModified with default headers values
func NewBasicAPIOperator2162NotModified() *BasicAPIOperator2162NotModified {
	return &BasicAPIOperator2162NotModified{}
}

/* BasicAPIOperator2162NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type BasicAPIOperator2162NotModified struct {
}

func (o *BasicAPIOperator2162NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Basic/Operator][%d] basicApiOperator2162NotModified ", 304)
}

func (o *BasicAPIOperator2162NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
