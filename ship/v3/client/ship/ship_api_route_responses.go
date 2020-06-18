// Code generated by go-swagger; DO NOT EDIT.

package ship

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/ship/v3/models"
)

// ShipAPIRouteReader is a Reader for the ShipAPIRoute structure.
type ShipAPIRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShipAPIRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShipAPIRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewShipAPIRouteOK creates a ShipAPIRouteOK with default headers values
func NewShipAPIRouteOK() *ShipAPIRouteOK {
	return &ShipAPIRouteOK{}
}

/*ShipAPIRouteOK handles this case with default header values.

Success
*/
type ShipAPIRouteOK struct {
	Payload *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Route
}

func (o *ShipAPIRouteOK) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/Route][%d] shipApiRouteOK  %+v", 200, o.Payload)
}

func (o *ShipAPIRouteOK) GetPayload() *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Route {
	return o.Payload
}

func (o *ShipAPIRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Route)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
