// Code generated by go-swagger; DO NOT EDIT.

package basic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/basic/v2/models"
)

// BasicAPIAuthorityReader is a Reader for the BasicAPIAuthority structure.
type BasicAPIAuthorityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BasicAPIAuthorityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBasicAPIAuthorityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewBasicAPIAuthorityNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBasicAPIAuthorityOK creates a BasicAPIAuthorityOK with default headers values
func NewBasicAPIAuthorityOK() *BasicAPIAuthorityOK {
	return &BasicAPIAuthorityOK{}
}

/* BasicAPIAuthorityOK describes a response with status code 200, with default header values.

Success
*/
type BasicAPIAuthorityOK struct {
	Payload []*models.PTXServiceDTOSharedSpecificationV2BaseAuthority
}

func (o *BasicAPIAuthorityOK) Error() string {
	return fmt.Sprintf("[GET /v2/Basic/Authority][%d] basicApiAuthorityOK  %+v", 200, o.Payload)
}
func (o *BasicAPIAuthorityOK) GetPayload() []*models.PTXServiceDTOSharedSpecificationV2BaseAuthority {
	return o.Payload
}

func (o *BasicAPIAuthorityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBasicAPIAuthorityNotModified creates a BasicAPIAuthorityNotModified with default headers values
func NewBasicAPIAuthorityNotModified() *BasicAPIAuthorityNotModified {
	return &BasicAPIAuthorityNotModified{}
}

/* BasicAPIAuthorityNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type BasicAPIAuthorityNotModified struct {
}

func (o *BasicAPIAuthorityNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Basic/Authority][%d] basicApiAuthorityNotModified ", 304)
}

func (o *BasicAPIAuthorityNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
