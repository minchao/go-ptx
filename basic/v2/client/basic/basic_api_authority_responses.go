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

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBasicAPIAuthorityOK creates a BasicAPIAuthorityOK with default headers values
func NewBasicAPIAuthorityOK() *BasicAPIAuthorityOK {
	return &BasicAPIAuthorityOK{}
}

/*BasicAPIAuthorityOK handles this case with default header values.

OK
*/
type BasicAPIAuthorityOK struct {
	Payload []*models.ServiceDTOVersion2BaseAuthority
}

func (o *BasicAPIAuthorityOK) Error() string {
	return fmt.Sprintf("[GET /v2/Basic/Authority][%d] basicApiAuthorityOK  %+v", 200, o.Payload)
}

func (o *BasicAPIAuthorityOK) GetPayload() []*models.ServiceDTOVersion2BaseAuthority {
	return o.Payload
}

func (o *BasicAPIAuthorityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
