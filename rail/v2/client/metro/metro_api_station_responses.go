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

// MetroAPIStationReader is a Reader for the MetroAPIStation structure.
type MetroAPIStationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroAPIStationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroAPIStationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMetroAPIStationOK creates a MetroAPIStationOK with default headers values
func NewMetroAPIStationOK() *MetroAPIStationOK {
	return &MetroAPIStationOK{}
}

/*MetroAPIStationOK handles this case with default header values.

OK
*/
type MetroAPIStationOK struct {
	Payload []*models.ServiceDTOVersion2RailMetroStation
}

func (o *MetroAPIStationOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/Station/{Operator}][%d] metroApiStationOK  %+v", 200, o.Payload)
}

func (o *MetroAPIStationOK) GetPayload() []*models.ServiceDTOVersion2RailMetroStation {
	return o.Payload
}

func (o *MetroAPIStationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
