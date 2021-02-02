// Code generated by go-swagger; DO NOT EDIT.

package t_h_s_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v2/models"
)

// THSRAPIODFareReader is a Reader for the THSRAPIODFare structure.
type THSRAPIODFareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIODFareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIODFareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTHSRAPIODFareNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTHSRAPIODFareOK creates a THSRAPIODFareOK with default headers values
func NewTHSRAPIODFareOK() *THSRAPIODFareOK {
	return &THSRAPIODFareOK{}
}

/* THSRAPIODFareOK describes a response with status code 200, with default header values.

Success
*/
type THSRAPIODFareOK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2THSRRailODFare
}

func (o *THSRAPIODFareOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/ODFare][%d] tHSRApiODFareOK  %+v", 200, o.Payload)
}
func (o *THSRAPIODFareOK) GetPayload() []*models.PTXServiceDTORailSpecificationV2THSRRailODFare {
	return o.Payload
}

func (o *THSRAPIODFareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPIODFareNotModified creates a THSRAPIODFareNotModified with default headers values
func NewTHSRAPIODFareNotModified() *THSRAPIODFareNotModified {
	return &THSRAPIODFareNotModified{}
}

/* THSRAPIODFareNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type THSRAPIODFareNotModified struct {
}

func (o *THSRAPIODFareNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/ODFare][%d] tHSRApiODFareNotModified ", 304)
}

func (o *THSRAPIODFareNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
