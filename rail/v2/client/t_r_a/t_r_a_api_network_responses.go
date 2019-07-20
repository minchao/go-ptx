// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/rail/v2/models"
)

// TRAAPINetworkReader is a Reader for the TRAAPINetwork structure.
type TRAAPINetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TRAAPINetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTRAAPINetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTRAAPINetworkOK creates a TRAAPINetworkOK with default headers values
func NewTRAAPINetworkOK() *TRAAPINetworkOK {
	return &TRAAPINetworkOK{}
}

/*TRAAPINetworkOK handles this case with default header values.

OK
*/
type TRAAPINetworkOK struct {
	Payload []*models.ServiceDTOVersion2RailTRANetwork
}

func (o *TRAAPINetworkOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/Network][%d] tRAApiNetworkOK  %+v", 200, o.Payload)
}

func (o *TRAAPINetworkOK) GetPayload() []*models.ServiceDTOVersion2RailTRANetwork {
	return o.Payload
}

func (o *TRAAPINetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
