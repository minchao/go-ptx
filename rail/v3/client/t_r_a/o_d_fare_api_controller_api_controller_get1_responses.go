// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v3/models"
)

// ODFareAPIControllerAPIControllerGet1Reader is a Reader for the ODFareAPIControllerAPIControllerGet1 structure.
type ODFareAPIControllerAPIControllerGet1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ODFareAPIControllerAPIControllerGet1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewODFareAPIControllerAPIControllerGet1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewODFareAPIControllerAPIControllerGet1NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewODFareAPIControllerAPIControllerGet1OK creates a ODFareAPIControllerAPIControllerGet1OK with default headers values
func NewODFareAPIControllerAPIControllerGet1OK() *ODFareAPIControllerAPIControllerGet1OK {
	return &ODFareAPIControllerAPIControllerGet1OK{}
}

/* ODFareAPIControllerAPIControllerGet1OK describes a response with status code 200, with default header values.

Success
*/
type ODFareAPIControllerAPIControllerGet1OK struct {
	Payload *models.PTXAPIRailModelTRAODFareWrapperPTXServiceDTORailSpecificationV3TRAODFareODFare
}

func (o *ODFareAPIControllerAPIControllerGet1OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/ODFare/{OriginStationID}/to/{DestinationStationID}][%d] oDFareApiControllerApiControllerGet1OK  %+v", 200, o.Payload)
}
func (o *ODFareAPIControllerAPIControllerGet1OK) GetPayload() *models.PTXAPIRailModelTRAODFareWrapperPTXServiceDTORailSpecificationV3TRAODFareODFare {
	return o.Payload
}

func (o *ODFareAPIControllerAPIControllerGet1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTRAODFareWrapperPTXServiceDTORailSpecificationV3TRAODFareODFare)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewODFareAPIControllerAPIControllerGet1NotModified creates a ODFareAPIControllerAPIControllerGet1NotModified with default headers values
func NewODFareAPIControllerAPIControllerGet1NotModified() *ODFareAPIControllerAPIControllerGet1NotModified {
	return &ODFareAPIControllerAPIControllerGet1NotModified{}
}

/* ODFareAPIControllerAPIControllerGet1NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type ODFareAPIControllerAPIControllerGet1NotModified struct {
}

func (o *ODFareAPIControllerAPIControllerGet1NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/ODFare/{OriginStationID}/to/{DestinationStationID}][%d] oDFareApiControllerApiControllerGet1NotModified ", 304)
}

func (o *ODFareAPIControllerAPIControllerGet1NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
