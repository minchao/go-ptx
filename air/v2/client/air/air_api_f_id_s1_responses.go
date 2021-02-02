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

// AirAPIFIDS1Reader is a Reader for the AirAPIFIDS1 structure.
type AirAPIFIDS1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AirAPIFIDS1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAirAPIFIDS1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewAirAPIFIDS1NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAirAPIFIDS1OK creates a AirAPIFIDS1OK with default headers values
func NewAirAPIFIDS1OK() *AirAPIFIDS1OK {
	return &AirAPIFIDS1OK{}
}

/* AirAPIFIDS1OK describes a response with status code 200, with default header values.

Success
*/
type AirAPIFIDS1OK struct {
	Payload []*models.PTXServiceDTOAirSpecificationV2AirportFIDS
}

func (o *AirAPIFIDS1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Air/FIDS/Airport/{IATA}][%d] airApiFIdS1OK  %+v", 200, o.Payload)
}
func (o *AirAPIFIDS1OK) GetPayload() []*models.PTXServiceDTOAirSpecificationV2AirportFIDS {
	return o.Payload
}

func (o *AirAPIFIDS1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAirAPIFIDS1NotModified creates a AirAPIFIDS1NotModified with default headers values
func NewAirAPIFIDS1NotModified() *AirAPIFIDS1NotModified {
	return &AirAPIFIDS1NotModified{}
}

/* AirAPIFIDS1NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type AirAPIFIDS1NotModified struct {
}

func (o *AirAPIFIDS1NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Air/FIDS/Airport/{IATA}][%d] airApiFIdS1NotModified ", 304)
}

func (o *AirAPIFIDS1NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
