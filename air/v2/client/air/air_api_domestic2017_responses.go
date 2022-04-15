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

// AirAPIDomestic2017Reader is a Reader for the AirAPIDomestic2017 structure.
type AirAPIDomestic2017Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AirAPIDomestic2017Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAirAPIDomestic2017OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewAirAPIDomestic2017Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewAirAPIDomestic2017NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAirAPIDomestic2017OK creates a AirAPIDomestic2017OK with default headers values
func NewAirAPIDomestic2017OK() *AirAPIDomestic2017OK {
	return &AirAPIDomestic2017OK{}
}

/* AirAPIDomestic2017OK describes a response with status code 200, with default header values.

Success
*/
type AirAPIDomestic2017OK struct {
	Payload []*models.PTXServiceDTOAirSpecificationV2GeneralFlightSchedule
}

func (o *AirAPIDomestic2017OK) Error() string {
	return fmt.Sprintf("[GET /v2/Air/GeneralSchedule/Domestic][%d] airApiDomestic2017OK  %+v", 200, o.Payload)
}
func (o *AirAPIDomestic2017OK) GetPayload() []*models.PTXServiceDTOAirSpecificationV2GeneralFlightSchedule {
	return o.Payload
}

func (o *AirAPIDomestic2017OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAirAPIDomestic2017Status299 creates a AirAPIDomestic2017Status299 with default headers values
func NewAirAPIDomestic2017Status299() *AirAPIDomestic2017Status299 {
	return &AirAPIDomestic2017Status299{}
}

/* AirAPIDomestic2017Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type AirAPIDomestic2017Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *AirAPIDomestic2017Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Air/GeneralSchedule/Domestic][%d] airApiDomestic2017Status299  %+v", 299, o.Payload)
}
func (o *AirAPIDomestic2017Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *AirAPIDomestic2017Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAirAPIDomestic2017NotModified creates a AirAPIDomestic2017NotModified with default headers values
func NewAirAPIDomestic2017NotModified() *AirAPIDomestic2017NotModified {
	return &AirAPIDomestic2017NotModified{}
}

/* AirAPIDomestic2017NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type AirAPIDomestic2017NotModified struct {
}

func (o *AirAPIDomestic2017NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Air/GeneralSchedule/Domestic][%d] airApiDomestic2017NotModified ", 304)
}

func (o *AirAPIDomestic2017NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}