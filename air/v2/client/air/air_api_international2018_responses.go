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

// AirAPIInternational2018Reader is a Reader for the AirAPIInternational2018 structure.
type AirAPIInternational2018Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AirAPIInternational2018Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAirAPIInternational2018OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewAirAPIInternational2018Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewAirAPIInternational2018NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAirAPIInternational2018OK creates a AirAPIInternational2018OK with default headers values
func NewAirAPIInternational2018OK() *AirAPIInternational2018OK {
	return &AirAPIInternational2018OK{}
}

/* AirAPIInternational2018OK describes a response with status code 200, with default header values.

Success
*/
type AirAPIInternational2018OK struct {
	Payload []*models.PTXServiceDTOAirSpecificationV2GeneralFlightSchedule
}

func (o *AirAPIInternational2018OK) Error() string {
	return fmt.Sprintf("[GET /v2/Air/GeneralSchedule/International][%d] airApiInternational2018OK  %+v", 200, o.Payload)
}
func (o *AirAPIInternational2018OK) GetPayload() []*models.PTXServiceDTOAirSpecificationV2GeneralFlightSchedule {
	return o.Payload
}

func (o *AirAPIInternational2018OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAirAPIInternational2018Status299 creates a AirAPIInternational2018Status299 with default headers values
func NewAirAPIInternational2018Status299() *AirAPIInternational2018Status299 {
	return &AirAPIInternational2018Status299{}
}

/* AirAPIInternational2018Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type AirAPIInternational2018Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *AirAPIInternational2018Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Air/GeneralSchedule/International][%d] airApiInternational2018Status299  %+v", 299, o.Payload)
}
func (o *AirAPIInternational2018Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *AirAPIInternational2018Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAirAPIInternational2018NotModified creates a AirAPIInternational2018NotModified with default headers values
func NewAirAPIInternational2018NotModified() *AirAPIInternational2018NotModified {
	return &AirAPIInternational2018NotModified{}
}

/* AirAPIInternational2018NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type AirAPIInternational2018NotModified struct {
}

func (o *AirAPIInternational2018NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Air/GeneralSchedule/International][%d] airApiInternational2018NotModified ", 304)
}

func (o *AirAPIInternational2018NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}