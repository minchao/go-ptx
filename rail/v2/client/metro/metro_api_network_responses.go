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

// MetroAPINetworkReader is a Reader for the MetroAPINetwork structure.
type MetroAPINetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroAPINetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroAPINetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewMetroAPINetworkStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewMetroAPINetworkNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMetroAPINetworkOK creates a MetroAPINetworkOK with default headers values
func NewMetroAPINetworkOK() *MetroAPINetworkOK {
	return &MetroAPINetworkOK{}
}

/* MetroAPINetworkOK describes a response with status code 200, with default header values.

Success
*/
type MetroAPINetworkOK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2MetroNetwork
}

func (o *MetroAPINetworkOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/Network/{Operator}][%d] metroApiNetworkOK  %+v", 200, o.Payload)
}
func (o *MetroAPINetworkOK) GetPayload() []*models.PTXServiceDTORailSpecificationV2MetroNetwork {
	return o.Payload
}

func (o *MetroAPINetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroAPINetworkStatus299 creates a MetroAPINetworkStatus299 with default headers values
func NewMetroAPINetworkStatus299() *MetroAPINetworkStatus299 {
	return &MetroAPINetworkStatus299{}
}

/* MetroAPINetworkStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type MetroAPINetworkStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *MetroAPINetworkStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/Network/{Operator}][%d] metroApiNetworkStatus299  %+v", 299, o.Payload)
}
func (o *MetroAPINetworkStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *MetroAPINetworkStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroAPINetworkNotModified creates a MetroAPINetworkNotModified with default headers values
func NewMetroAPINetworkNotModified() *MetroAPINetworkNotModified {
	return &MetroAPINetworkNotModified{}
}

/* MetroAPINetworkNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type MetroAPINetworkNotModified struct {
}

func (o *MetroAPINetworkNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/Network/{Operator}][%d] metroApiNetworkNotModified ", 304)
}

func (o *MetroAPINetworkNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
