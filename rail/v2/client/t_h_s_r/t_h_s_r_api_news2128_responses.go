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

// THSRAPINews2128Reader is a Reader for the THSRAPINews2128 structure.
type THSRAPINews2128Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPINews2128Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPINews2128OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTHSRAPINews2128Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTHSRAPINews2128NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTHSRAPINews2128OK creates a THSRAPINews2128OK with default headers values
func NewTHSRAPINews2128OK() *THSRAPINews2128OK {
	return &THSRAPINews2128OK{}
}

/* THSRAPINews2128OK describes a response with status code 200, with default header values.

Success
*/
type THSRAPINews2128OK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2THSRNews
}

func (o *THSRAPINews2128OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/News][%d] tHSRApiNews2128OK  %+v", 200, o.Payload)
}
func (o *THSRAPINews2128OK) GetPayload() []*models.PTXServiceDTORailSpecificationV2THSRNews {
	return o.Payload
}

func (o *THSRAPINews2128OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPINews2128Status299 creates a THSRAPINews2128Status299 with default headers values
func NewTHSRAPINews2128Status299() *THSRAPINews2128Status299 {
	return &THSRAPINews2128Status299{}
}

/* THSRAPINews2128Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type THSRAPINews2128Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *THSRAPINews2128Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/News][%d] tHSRApiNews2128Status299  %+v", 299, o.Payload)
}
func (o *THSRAPINews2128Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *THSRAPINews2128Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPINews2128NotModified creates a THSRAPINews2128NotModified with default headers values
func NewTHSRAPINews2128NotModified() *THSRAPINews2128NotModified {
	return &THSRAPINews2128NotModified{}
}

/* THSRAPINews2128NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type THSRAPINews2128NotModified struct {
}

func (o *THSRAPINews2128NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/News][%d] tHSRApiNews2128NotModified ", 304)
}

func (o *THSRAPINews2128NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
