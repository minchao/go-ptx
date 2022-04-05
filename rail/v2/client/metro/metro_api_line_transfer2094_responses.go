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

// MetroAPILineTransfer2094Reader is a Reader for the MetroAPILineTransfer2094 structure.
type MetroAPILineTransfer2094Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroAPILineTransfer2094Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroAPILineTransfer2094OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewMetroAPILineTransfer2094Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewMetroAPILineTransfer2094NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMetroAPILineTransfer2094OK creates a MetroAPILineTransfer2094OK with default headers values
func NewMetroAPILineTransfer2094OK() *MetroAPILineTransfer2094OK {
	return &MetroAPILineTransfer2094OK{}
}

/* MetroAPILineTransfer2094OK describes a response with status code 200, with default header values.

Success
*/
type MetroAPILineTransfer2094OK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2MetroLineTransfer
}

func (o *MetroAPILineTransfer2094OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/LineTransfer/{MetroSystem}][%d] metroApiLineTransfer2094OK  %+v", 200, o.Payload)
}
func (o *MetroAPILineTransfer2094OK) GetPayload() []*models.PTXServiceDTORailSpecificationV2MetroLineTransfer {
	return o.Payload
}

func (o *MetroAPILineTransfer2094OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroAPILineTransfer2094Status299 creates a MetroAPILineTransfer2094Status299 with default headers values
func NewMetroAPILineTransfer2094Status299() *MetroAPILineTransfer2094Status299 {
	return &MetroAPILineTransfer2094Status299{}
}

/* MetroAPILineTransfer2094Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type MetroAPILineTransfer2094Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *MetroAPILineTransfer2094Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/LineTransfer/{MetroSystem}][%d] metroApiLineTransfer2094Status299  %+v", 299, o.Payload)
}
func (o *MetroAPILineTransfer2094Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *MetroAPILineTransfer2094Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroAPILineTransfer2094NotModified creates a MetroAPILineTransfer2094NotModified with default headers values
func NewMetroAPILineTransfer2094NotModified() *MetroAPILineTransfer2094NotModified {
	return &MetroAPILineTransfer2094NotModified{}
}

/* MetroAPILineTransfer2094NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type MetroAPILineTransfer2094NotModified struct {
}

func (o *MetroAPILineTransfer2094NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/LineTransfer/{MetroSystem}][%d] metroApiLineTransfer2094NotModified ", 304)
}

func (o *MetroAPILineTransfer2094NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
