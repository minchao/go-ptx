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

// MetroAPIFrequencyReader is a Reader for the MetroAPIFrequency structure.
type MetroAPIFrequencyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroAPIFrequencyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroAPIFrequencyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMetroAPIFrequencyOK creates a MetroAPIFrequencyOK with default headers values
func NewMetroAPIFrequencyOK() *MetroAPIFrequencyOK {
	return &MetroAPIFrequencyOK{}
}

/*MetroAPIFrequencyOK handles this case with default header values.

OK
*/
type MetroAPIFrequencyOK struct {
	Payload []*models.ServiceDTOVersion2RailMetroFrequency
}

func (o *MetroAPIFrequencyOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/Frequency/{Operator}][%d] metroApiFrequencyOK  %+v", 200, o.Payload)
}

func (o *MetroAPIFrequencyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
