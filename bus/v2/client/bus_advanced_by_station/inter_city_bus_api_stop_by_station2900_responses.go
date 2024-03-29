// Code generated by go-swagger; DO NOT EDIT.

package bus_advanced_by_station

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v2/models"
)

// InterCityBusAPIStopByStation2900Reader is a Reader for the InterCityBusAPIStopByStation2900 structure.
type InterCityBusAPIStopByStation2900Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIStopByStation2900Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIStopByStation2900OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIStopByStation2900Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIStopByStation2900NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIStopByStation2900OK creates a InterCityBusAPIStopByStation2900OK with default headers values
func NewInterCityBusAPIStopByStation2900OK() *InterCityBusAPIStopByStation2900OK {
	return &InterCityBusAPIStopByStation2900OK{}
}

/* InterCityBusAPIStopByStation2900OK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIStopByStation2900OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusStop
}

func (o *InterCityBusAPIStopByStation2900OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Stop/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiStopByStation2900OK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIStopByStation2900OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusStop {
	return o.Payload
}

func (o *InterCityBusAPIStopByStation2900OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIStopByStation2900Status299 creates a InterCityBusAPIStopByStation2900Status299 with default headers values
func NewInterCityBusAPIStopByStation2900Status299() *InterCityBusAPIStopByStation2900Status299 {
	return &InterCityBusAPIStopByStation2900Status299{}
}

/* InterCityBusAPIStopByStation2900Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIStopByStation2900Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIStopByStation2900Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Stop/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiStopByStation2900Status299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIStopByStation2900Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIStopByStation2900Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIStopByStation2900NotModified creates a InterCityBusAPIStopByStation2900NotModified with default headers values
func NewInterCityBusAPIStopByStation2900NotModified() *InterCityBusAPIStopByStation2900NotModified {
	return &InterCityBusAPIStopByStation2900NotModified{}
}

/* InterCityBusAPIStopByStation2900NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIStopByStation2900NotModified struct {
}

func (o *InterCityBusAPIStopByStation2900NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Stop/InterCity/PassThrough/Station/{StationID}][%d] interCityBusApiStopByStation2900NotModified ", 304)
}

func (o *InterCityBusAPIStopByStation2900NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
