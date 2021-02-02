// Code generated by go-swagger; DO NOT EDIT.

package tourism

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/tourism/v2/models"
)

// TaiwanTripBusAPINewsReader is a Reader for the TaiwanTripBusAPINews structure.
type TaiwanTripBusAPINewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaiwanTripBusAPINewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaiwanTripBusAPINewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTaiwanTripBusAPINewsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTaiwanTripBusAPINewsOK creates a TaiwanTripBusAPINewsOK with default headers values
func NewTaiwanTripBusAPINewsOK() *TaiwanTripBusAPINewsOK {
	return &TaiwanTripBusAPINewsOK{}
}

/* TaiwanTripBusAPINewsOK describes a response with status code 200, with default header values.

Success
*/
type TaiwanTripBusAPINewsOK struct {
	Payload []*models.PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews
}

func (o *TaiwanTripBusAPINewsOK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/News/TaiwanTrip][%d] taiwanTripBusApiNewsOK  %+v", 200, o.Payload)
}
func (o *TaiwanTripBusAPINewsOK) GetPayload() []*models.PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews {
	return o.Payload
}

func (o *TaiwanTripBusAPINewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaiwanTripBusAPINewsNotModified creates a TaiwanTripBusAPINewsNotModified with default headers values
func NewTaiwanTripBusAPINewsNotModified() *TaiwanTripBusAPINewsNotModified {
	return &TaiwanTripBusAPINewsNotModified{}
}

/* TaiwanTripBusAPINewsNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type TaiwanTripBusAPINewsNotModified struct {
}

func (o *TaiwanTripBusAPINewsNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/News/TaiwanTrip][%d] taiwanTripBusApiNewsNotModified ", 304)
}

func (o *TaiwanTripBusAPINewsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
