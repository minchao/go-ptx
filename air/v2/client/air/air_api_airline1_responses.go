// Code generated by go-swagger; DO NOT EDIT.

package air

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/air/v2/models"
)

// AirAPIAirline1Reader is a Reader for the AirAPIAirline1 structure.
type AirAPIAirline1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AirAPIAirline1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAirAPIAirline1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewAirAPIAirline1NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAirAPIAirline1OK creates a AirAPIAirline1OK with default headers values
func NewAirAPIAirline1OK() *AirAPIAirline1OK {
	return &AirAPIAirline1OK{}
}

/* AirAPIAirline1OK describes a response with status code 200, with default header values.

Success
*/
type AirAPIAirline1OK struct {
	Payload []*models.PTXServiceDTOAirSpecificationV2Airline
}

func (o *AirAPIAirline1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Air/Airline][%d] airApiAirline1OK  %+v", 200, o.Payload)
}
func (o *AirAPIAirline1OK) GetPayload() []*models.PTXServiceDTOAirSpecificationV2Airline {
	return o.Payload
}

func (o *AirAPIAirline1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAirAPIAirline1NotModified creates a AirAPIAirline1NotModified with default headers values
func NewAirAPIAirline1NotModified() *AirAPIAirline1NotModified {
	return &AirAPIAirline1NotModified{}
}

/* AirAPIAirline1NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type AirAPIAirline1NotModified struct {
}

func (o *AirAPIAirline1NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Air/Airline][%d] airApiAirline1NotModified ", 304)
}

func (o *AirAPIAirline1NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
