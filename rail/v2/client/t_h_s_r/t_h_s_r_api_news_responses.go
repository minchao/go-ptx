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

// THSRAPINewsReader is a Reader for the THSRAPINews structure.
type THSRAPINewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPINewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPINewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTHSRAPINewsStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTHSRAPINewsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTHSRAPINewsOK creates a THSRAPINewsOK with default headers values
func NewTHSRAPINewsOK() *THSRAPINewsOK {
	return &THSRAPINewsOK{}
}

/* THSRAPINewsOK describes a response with status code 200, with default header values.

Success
*/
type THSRAPINewsOK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2THSRNews
}

func (o *THSRAPINewsOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/News][%d] tHSRApiNewsOK  %+v", 200, o.Payload)
}
func (o *THSRAPINewsOK) GetPayload() []*models.PTXServiceDTORailSpecificationV2THSRNews {
	return o.Payload
}

func (o *THSRAPINewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPINewsStatus299 creates a THSRAPINewsStatus299 with default headers values
func NewTHSRAPINewsStatus299() *THSRAPINewsStatus299 {
	return &THSRAPINewsStatus299{}
}

/* THSRAPINewsStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type THSRAPINewsStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *THSRAPINewsStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/News][%d] tHSRApiNewsStatus299  %+v", 299, o.Payload)
}
func (o *THSRAPINewsStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *THSRAPINewsStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPINewsNotModified creates a THSRAPINewsNotModified with default headers values
func NewTHSRAPINewsNotModified() *THSRAPINewsNotModified {
	return &THSRAPINewsNotModified{}
}

/* THSRAPINewsNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type THSRAPINewsNotModified struct {
}

func (o *THSRAPINewsNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/News][%d] tHSRApiNewsNotModified ", 304)
}

func (o *THSRAPINewsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
