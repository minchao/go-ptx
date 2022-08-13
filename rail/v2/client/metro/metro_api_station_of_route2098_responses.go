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

// MetroAPIStationOfRoute2098Reader is a Reader for the MetroAPIStationOfRoute2098 structure.
type MetroAPIStationOfRoute2098Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroAPIStationOfRoute2098Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroAPIStationOfRoute2098OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewMetroAPIStationOfRoute2098Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewMetroAPIStationOfRoute2098NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMetroAPIStationOfRoute2098OK creates a MetroAPIStationOfRoute2098OK with default headers values
func NewMetroAPIStationOfRoute2098OK() *MetroAPIStationOfRoute2098OK {
	return &MetroAPIStationOfRoute2098OK{}
}

/* MetroAPIStationOfRoute2098OK describes a response with status code 200, with default header values.

Success
*/
type MetroAPIStationOfRoute2098OK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2MetroStationOfRoute
}

func (o *MetroAPIStationOfRoute2098OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/StationOfRoute/{RailSystem}][%d] metroApiStationOfRoute2098OK  %+v", 200, o.Payload)
}
func (o *MetroAPIStationOfRoute2098OK) GetPayload() []*models.PTXServiceDTORailSpecificationV2MetroStationOfRoute {
	return o.Payload
}

func (o *MetroAPIStationOfRoute2098OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroAPIStationOfRoute2098Status299 creates a MetroAPIStationOfRoute2098Status299 with default headers values
func NewMetroAPIStationOfRoute2098Status299() *MetroAPIStationOfRoute2098Status299 {
	return &MetroAPIStationOfRoute2098Status299{}
}

/* MetroAPIStationOfRoute2098Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type MetroAPIStationOfRoute2098Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *MetroAPIStationOfRoute2098Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/StationOfRoute/{RailSystem}][%d] metroApiStationOfRoute2098Status299  %+v", 299, o.Payload)
}
func (o *MetroAPIStationOfRoute2098Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *MetroAPIStationOfRoute2098Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroAPIStationOfRoute2098NotModified creates a MetroAPIStationOfRoute2098NotModified with default headers values
func NewMetroAPIStationOfRoute2098NotModified() *MetroAPIStationOfRoute2098NotModified {
	return &MetroAPIStationOfRoute2098NotModified{}
}

/* MetroAPIStationOfRoute2098NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type MetroAPIStationOfRoute2098NotModified struct {
}

func (o *MetroAPIStationOfRoute2098NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/StationOfRoute/{RailSystem}][%d] metroApiStationOfRoute2098NotModified ", 304)
}

func (o *MetroAPIStationOfRoute2098NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
