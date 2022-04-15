// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v2/models"
)

// TRAAPILiveBoard21531Reader is a Reader for the TRAAPILiveBoard21531 structure.
type TRAAPILiveBoard21531Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TRAAPILiveBoard21531Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTRAAPILiveBoard21531OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTRAAPILiveBoard21531Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTRAAPILiveBoard21531NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTRAAPILiveBoard21531OK creates a TRAAPILiveBoard21531OK with default headers values
func NewTRAAPILiveBoard21531OK() *TRAAPILiveBoard21531OK {
	return &TRAAPILiveBoard21531OK{}
}

/* TRAAPILiveBoard21531OK describes a response with status code 200, with default header values.

Success
*/
type TRAAPILiveBoard21531OK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2TRARailLiveBoard
}

func (o *TRAAPILiveBoard21531OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/LiveBoard/Station/{StationID}][%d] tRAApiLiveBoard21531OK  %+v", 200, o.Payload)
}
func (o *TRAAPILiveBoard21531OK) GetPayload() []*models.PTXServiceDTORailSpecificationV2TRARailLiveBoard {
	return o.Payload
}

func (o *TRAAPILiveBoard21531OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTRAAPILiveBoard21531Status299 creates a TRAAPILiveBoard21531Status299 with default headers values
func NewTRAAPILiveBoard21531Status299() *TRAAPILiveBoard21531Status299 {
	return &TRAAPILiveBoard21531Status299{}
}

/* TRAAPILiveBoard21531Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type TRAAPILiveBoard21531Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *TRAAPILiveBoard21531Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/LiveBoard/Station/{StationID}][%d] tRAApiLiveBoard21531Status299  %+v", 299, o.Payload)
}
func (o *TRAAPILiveBoard21531Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *TRAAPILiveBoard21531Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTRAAPILiveBoard21531NotModified creates a TRAAPILiveBoard21531NotModified with default headers values
func NewTRAAPILiveBoard21531NotModified() *TRAAPILiveBoard21531NotModified {
	return &TRAAPILiveBoard21531NotModified{}
}

/* TRAAPILiveBoard21531NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type TRAAPILiveBoard21531NotModified struct {
}

func (o *TRAAPILiveBoard21531NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/LiveBoard/Station/{StationID}][%d] tRAApiLiveBoard21531NotModified ", 304)
}

func (o *TRAAPILiveBoard21531NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}