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

// ShipAPIRouteFare1Reader is a Reader for the ShipAPIRouteFare1 structure.
type ShipAPIRouteFare1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShipAPIRouteFare1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShipAPIRouteFare1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewShipAPIRouteFare1NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShipAPIRouteFare1OK creates a ShipAPIRouteFare1OK with default headers values
func NewShipAPIRouteFare1OK() *ShipAPIRouteFare1OK {
	return &ShipAPIRouteFare1OK{}
}

/* ShipAPIRouteFare1OK describes a response with status code 200, with default header values.

Success
*/
type ShipAPIRouteFare1OK struct {
	Payload *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3RouteFare
}

func (o *ShipAPIRouteFare1OK) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/RouteFare/{RouteID}][%d] shipApiRouteFare1OK  %+v", 200, o.Payload)
}
func (o *ShipAPIRouteFare1OK) GetPayload() *models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3RouteFare {
	return o.Payload
}

func (o *ShipAPIRouteFare1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3RouteFare)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShipAPIRouteFare1NotModified creates a ShipAPIRouteFare1NotModified with default headers values
func NewShipAPIRouteFare1NotModified() *ShipAPIRouteFare1NotModified {
	return &ShipAPIRouteFare1NotModified{}
}

/* ShipAPIRouteFare1NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type ShipAPIRouteFare1NotModified struct {
}

func (o *ShipAPIRouteFare1NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Ship/RouteFare/{RouteID}][%d] shipApiRouteFare1NotModified ", 304)
}

func (o *ShipAPIRouteFare1NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
