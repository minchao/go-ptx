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

// THSRAPIAvailableSeatStatusListStationReader is a Reader for the THSRAPIAvailableSeatStatusListStation structure.
type THSRAPIAvailableSeatStatusListStationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIAvailableSeatStatusListStationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIAvailableSeatStatusListStationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTHSRAPIAvailableSeatStatusListStationStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTHSRAPIAvailableSeatStatusListStationNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTHSRAPIAvailableSeatStatusListStationOK creates a THSRAPIAvailableSeatStatusListStationOK with default headers values
func NewTHSRAPIAvailableSeatStatusListStationOK() *THSRAPIAvailableSeatStatusListStationOK {
	return &THSRAPIAvailableSeatStatusListStationOK{}
}

/* THSRAPIAvailableSeatStatusListStationOK describes a response with status code 200, with default header values.

Success
*/
type THSRAPIAvailableSeatStatusListStationOK struct {
	Payload *models.PTXAPIRailModelV2THSRAvailableSeatStatusOldWrapperPTXServiceDTORailSpecificationV2THSROldAvailableSeat
}

func (o *THSRAPIAvailableSeatStatusListStationOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/AvailableSeatStatusList][%d] tHSRApiAvailableSeatStatusListStationOK  %+v", 200, o.Payload)
}
func (o *THSRAPIAvailableSeatStatusListStationOK) GetPayload() *models.PTXAPIRailModelV2THSRAvailableSeatStatusOldWrapperPTXServiceDTORailSpecificationV2THSROldAvailableSeat {
	return o.Payload
}

func (o *THSRAPIAvailableSeatStatusListStationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelV2THSRAvailableSeatStatusOldWrapperPTXServiceDTORailSpecificationV2THSROldAvailableSeat)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPIAvailableSeatStatusListStationStatus299 creates a THSRAPIAvailableSeatStatusListStationStatus299 with default headers values
func NewTHSRAPIAvailableSeatStatusListStationStatus299() *THSRAPIAvailableSeatStatusListStationStatus299 {
	return &THSRAPIAvailableSeatStatusListStationStatus299{}
}

/* THSRAPIAvailableSeatStatusListStationStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type THSRAPIAvailableSeatStatusListStationStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *THSRAPIAvailableSeatStatusListStationStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/AvailableSeatStatusList][%d] tHSRApiAvailableSeatStatusListStationStatus299  %+v", 299, o.Payload)
}
func (o *THSRAPIAvailableSeatStatusListStationStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *THSRAPIAvailableSeatStatusListStationStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPIAvailableSeatStatusListStationNotModified creates a THSRAPIAvailableSeatStatusListStationNotModified with default headers values
func NewTHSRAPIAvailableSeatStatusListStationNotModified() *THSRAPIAvailableSeatStatusListStationNotModified {
	return &THSRAPIAvailableSeatStatusListStationNotModified{}
}

/* THSRAPIAvailableSeatStatusListStationNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type THSRAPIAvailableSeatStatusListStationNotModified struct {
}

func (o *THSRAPIAvailableSeatStatusListStationNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/AvailableSeatStatusList][%d] tHSRApiAvailableSeatStatusListStationNotModified ", 304)
}

func (o *THSRAPIAvailableSeatStatusListStationNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
