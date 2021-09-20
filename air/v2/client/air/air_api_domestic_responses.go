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

// AirAPIDomesticReader is a Reader for the AirAPIDomestic structure.
type AirAPIDomesticReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AirAPIDomesticReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAirAPIDomesticOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewAirAPIDomesticStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewAirAPIDomesticNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAirAPIDomesticOK creates a AirAPIDomesticOK with default headers values
func NewAirAPIDomesticOK() *AirAPIDomesticOK {
	return &AirAPIDomesticOK{}
}

/* AirAPIDomesticOK describes a response with status code 200, with default header values.

Success
*/
type AirAPIDomesticOK struct {
	Payload []*models.PTXServiceDTOAirSpecificationV2GeneralFlightSchedule
}

func (o *AirAPIDomesticOK) Error() string {
	return fmt.Sprintf("[GET /v2/Air/GeneralSchedule/Domestic][%d] airApiDomesticOK  %+v", 200, o.Payload)
}
func (o *AirAPIDomesticOK) GetPayload() []*models.PTXServiceDTOAirSpecificationV2GeneralFlightSchedule {
	return o.Payload
}

func (o *AirAPIDomesticOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAirAPIDomesticStatus299 creates a AirAPIDomesticStatus299 with default headers values
func NewAirAPIDomesticStatus299() *AirAPIDomesticStatus299 {
	return &AirAPIDomesticStatus299{}
}

/* AirAPIDomesticStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type AirAPIDomesticStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *AirAPIDomesticStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Air/GeneralSchedule/Domestic][%d] airApiDomesticStatus299  %+v", 299, o.Payload)
}
func (o *AirAPIDomesticStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *AirAPIDomesticStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAirAPIDomesticNotModified creates a AirAPIDomesticNotModified with default headers values
func NewAirAPIDomesticNotModified() *AirAPIDomesticNotModified {
	return &AirAPIDomesticNotModified{}
}

/* AirAPIDomesticNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type AirAPIDomesticNotModified struct {
}

func (o *AirAPIDomesticNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Air/GeneralSchedule/Domestic][%d] airApiDomesticNotModified ", 304)
}

func (o *AirAPIDomesticNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
