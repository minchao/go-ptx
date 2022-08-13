// Code generated by go-swagger; DO NOT EDIT.

package metro

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v2/models"
)

// MetroAPIFirstLastTimetable2099Reader is a Reader for the MetroAPIFirstLastTimetable2099 structure.
type MetroAPIFirstLastTimetable2099Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroAPIFirstLastTimetable2099Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroAPIFirstLastTimetable2099OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewMetroAPIFirstLastTimetable2099Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewMetroAPIFirstLastTimetable2099NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMetroAPIFirstLastTimetable2099OK creates a MetroAPIFirstLastTimetable2099OK with default headers values
func NewMetroAPIFirstLastTimetable2099OK() *MetroAPIFirstLastTimetable2099OK {
	return &MetroAPIFirstLastTimetable2099OK{}
}

/* MetroAPIFirstLastTimetable2099OK describes a response with status code 200, with default header values.

Success
*/
type MetroAPIFirstLastTimetable2099OK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2MetroFirstLastTimetable
}

func (o *MetroAPIFirstLastTimetable2099OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/FirstLastTimetable/{RailSystem}][%d] metroApiFirstLastTimetable2099OK  %+v", 200, o.Payload)
}
func (o *MetroAPIFirstLastTimetable2099OK) GetPayload() []*models.PTXServiceDTORailSpecificationV2MetroFirstLastTimetable {
	return o.Payload
}

func (o *MetroAPIFirstLastTimetable2099OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroAPIFirstLastTimetable2099Status299 creates a MetroAPIFirstLastTimetable2099Status299 with default headers values
func NewMetroAPIFirstLastTimetable2099Status299() *MetroAPIFirstLastTimetable2099Status299 {
	return &MetroAPIFirstLastTimetable2099Status299{}
}

/* MetroAPIFirstLastTimetable2099Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type MetroAPIFirstLastTimetable2099Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *MetroAPIFirstLastTimetable2099Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/FirstLastTimetable/{RailSystem}][%d] metroApiFirstLastTimetable2099Status299  %+v", 299, o.Payload)
}
func (o *MetroAPIFirstLastTimetable2099Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *MetroAPIFirstLastTimetable2099Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroAPIFirstLastTimetable2099NotModified creates a MetroAPIFirstLastTimetable2099NotModified with default headers values
func NewMetroAPIFirstLastTimetable2099NotModified() *MetroAPIFirstLastTimetable2099NotModified {
	return &MetroAPIFirstLastTimetable2099NotModified{}
}

/* MetroAPIFirstLastTimetable2099NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type MetroAPIFirstLastTimetable2099NotModified struct {
}

func (o *MetroAPIFirstLastTimetable2099NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/FirstLastTimetable/{RailSystem}][%d] metroApiFirstLastTimetable2099NotModified ", 304)
}

func (o *MetroAPIFirstLastTimetable2099NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
