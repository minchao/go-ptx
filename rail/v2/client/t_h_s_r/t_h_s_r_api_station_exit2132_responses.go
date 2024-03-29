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

// THSRAPIStationExit2132Reader is a Reader for the THSRAPIStationExit2132 structure.
type THSRAPIStationExit2132Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIStationExit2132Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIStationExit2132OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTHSRAPIStationExit2132Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTHSRAPIStationExit2132NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTHSRAPIStationExit2132OK creates a THSRAPIStationExit2132OK with default headers values
func NewTHSRAPIStationExit2132OK() *THSRAPIStationExit2132OK {
	return &THSRAPIStationExit2132OK{}
}

/* THSRAPIStationExit2132OK describes a response with status code 200, with default header values.

Success
*/
type THSRAPIStationExit2132OK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2THSRStationExit
}

func (o *THSRAPIStationExit2132OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/StationExit][%d] tHSRApiStationExit2132OK  %+v", 200, o.Payload)
}
func (o *THSRAPIStationExit2132OK) GetPayload() []*models.PTXServiceDTORailSpecificationV2THSRStationExit {
	return o.Payload
}

func (o *THSRAPIStationExit2132OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPIStationExit2132Status299 creates a THSRAPIStationExit2132Status299 with default headers values
func NewTHSRAPIStationExit2132Status299() *THSRAPIStationExit2132Status299 {
	return &THSRAPIStationExit2132Status299{}
}

/* THSRAPIStationExit2132Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type THSRAPIStationExit2132Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *THSRAPIStationExit2132Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/StationExit][%d] tHSRApiStationExit2132Status299  %+v", 299, o.Payload)
}
func (o *THSRAPIStationExit2132Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *THSRAPIStationExit2132Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPIStationExit2132NotModified creates a THSRAPIStationExit2132NotModified with default headers values
func NewTHSRAPIStationExit2132NotModified() *THSRAPIStationExit2132NotModified {
	return &THSRAPIStationExit2132NotModified{}
}

/* THSRAPIStationExit2132NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type THSRAPIStationExit2132NotModified struct {
}

func (o *THSRAPIStationExit2132NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/StationExit][%d] tHSRApiStationExit2132NotModified ", 304)
}

func (o *THSRAPIStationExit2132NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
