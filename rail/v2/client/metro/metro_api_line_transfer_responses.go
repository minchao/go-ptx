// Code generated by go-swagger; DO NOT EDIT.

package metro

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v2/models"
)

// MetroAPILineTransferReader is a Reader for the MetroAPILineTransfer structure.
type MetroAPILineTransferReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroAPILineTransferReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroAPILineTransferOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMetroAPILineTransferOK creates a MetroAPILineTransferOK with default headers values
func NewMetroAPILineTransferOK() *MetroAPILineTransferOK {
	return &MetroAPILineTransferOK{}
}

/*MetroAPILineTransferOK handles this case with default header values.

Success
*/
type MetroAPILineTransferOK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2MetroLineTransfer
}

func (o *MetroAPILineTransferOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/LineTransfer/{Operator}][%d] metroApiLineTransferOK  %+v", 200, o.Payload)
}

func (o *MetroAPILineTransferOK) GetPayload() []*models.PTXServiceDTORailSpecificationV2MetroLineTransfer {
	return o.Payload
}

func (o *MetroAPILineTransferOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
