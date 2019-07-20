// Code generated by go-swagger; DO NOT EDIT.

package metro

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/rail/v2/models"
)

// MetroAPINetworkReader is a Reader for the MetroAPINetwork structure.
type MetroAPINetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroAPINetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroAPINetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMetroAPINetworkOK creates a MetroAPINetworkOK with default headers values
func NewMetroAPINetworkOK() *MetroAPINetworkOK {
	return &MetroAPINetworkOK{}
}

/*MetroAPINetworkOK handles this case with default header values.

OK
*/
type MetroAPINetworkOK struct {
	Payload []*models.ServiceDTOVersion2RailMetroNetwork
}

func (o *MetroAPINetworkOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/Network][%d] metroApiNetworkOK  %+v", 200, o.Payload)
}

func (o *MetroAPINetworkOK) GetPayload() []*models.ServiceDTOVersion2RailMetroNetwork {
	return o.Payload
}

func (o *MetroAPINetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
