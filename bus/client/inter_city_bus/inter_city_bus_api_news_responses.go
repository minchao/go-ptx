// Code generated by go-swagger; DO NOT EDIT.

package inter_city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/bus/models"
)

// InterCityBusAPINewsReader is a Reader for the InterCityBusAPINews structure.
type InterCityBusAPINewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPINewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPINewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInterCityBusAPINewsOK creates a InterCityBusAPINewsOK with default headers values
func NewInterCityBusAPINewsOK() *InterCityBusAPINewsOK {
	return &InterCityBusAPINewsOK{}
}

/*InterCityBusAPINewsOK handles this case with default header values.

OK
*/
type InterCityBusAPINewsOK struct {
	Payload []*models.ServiceDTOVersion2BusBusNews
}

func (o *InterCityBusAPINewsOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/News/InterCity][%d] interCityBusApiNewsOK  %+v", 200, o.Payload)
}

func (o *InterCityBusAPINewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
