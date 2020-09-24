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

// MetroAPIRouteReader is a Reader for the MetroAPIRoute structure.
type MetroAPIRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroAPIRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroAPIRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMetroAPIRouteOK creates a MetroAPIRouteOK with default headers values
func NewMetroAPIRouteOK() *MetroAPIRouteOK {
	return &MetroAPIRouteOK{}
}

/*MetroAPIRouteOK handles this case with default header values.

Success
*/
type MetroAPIRouteOK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2MetroRoute
}

func (o *MetroAPIRouteOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/Route/{Operator}][%d] metroApiRouteOK  %+v", 200, o.Payload)
}

func (o *MetroAPIRouteOK) GetPayload() []*models.PTXServiceDTORailSpecificationV2MetroRoute {
	return o.Payload
}

func (o *MetroAPIRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
