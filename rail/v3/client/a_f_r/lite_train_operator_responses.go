// Code generated by go-swagger; DO NOT EDIT.

package a_f_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v3/models"
)

// LiteTrainOperatorReader is a Reader for the LiteTrainOperator structure.
type LiteTrainOperatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LiteTrainOperatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLiteTrainOperatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewLiteTrainOperatorNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLiteTrainOperatorOK creates a LiteTrainOperatorOK with default headers values
func NewLiteTrainOperatorOK() *LiteTrainOperatorOK {
	return &LiteTrainOperatorOK{}
}

/* LiteTrainOperatorOK describes a response with status code 200, with default header values.

Success
*/
type LiteTrainOperatorOK struct {
	Payload *models.PTXAPIRailModelLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainOperator
}

func (o *LiteTrainOperatorOK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/Operator][%d] liteTrainOperatorOK  %+v", 200, o.Payload)
}
func (o *LiteTrainOperatorOK) GetPayload() *models.PTXAPIRailModelLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainOperator {
	return o.Payload
}

func (o *LiteTrainOperatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainOperator)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLiteTrainOperatorNotModified creates a LiteTrainOperatorNotModified with default headers values
func NewLiteTrainOperatorNotModified() *LiteTrainOperatorNotModified {
	return &LiteTrainOperatorNotModified{}
}

/* LiteTrainOperatorNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type LiteTrainOperatorNotModified struct {
}

func (o *LiteTrainOperatorNotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/Operator][%d] liteTrainOperatorNotModified ", 304)
}

func (o *LiteTrainOperatorNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
