// Code generated by go-swagger; DO NOT EDIT.

package bus_advanced_by_operator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v2/models"
)

// InterCityBusAPIStationByOperator2891Reader is a Reader for the InterCityBusAPIStationByOperator2891 structure.
type InterCityBusAPIStationByOperator2891Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIStationByOperator2891Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIStationByOperator2891OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIStationByOperator2891Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIStationByOperator2891NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIStationByOperator2891OK creates a InterCityBusAPIStationByOperator2891OK with default headers values
func NewInterCityBusAPIStationByOperator2891OK() *InterCityBusAPIStationByOperator2891OK {
	return &InterCityBusAPIStationByOperator2891OK{}
}

/* InterCityBusAPIStationByOperator2891OK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIStationByOperator2891OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusStation
}

func (o *InterCityBusAPIStationByOperator2891OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Station/InterCity/OperatorNo/{OperatorNo}][%d] interCityBusApiStationByOperator2891OK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIStationByOperator2891OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusStation {
	return o.Payload
}

func (o *InterCityBusAPIStationByOperator2891OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIStationByOperator2891Status299 creates a InterCityBusAPIStationByOperator2891Status299 with default headers values
func NewInterCityBusAPIStationByOperator2891Status299() *InterCityBusAPIStationByOperator2891Status299 {
	return &InterCityBusAPIStationByOperator2891Status299{}
}

/* InterCityBusAPIStationByOperator2891Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIStationByOperator2891Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIStationByOperator2891Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Station/InterCity/OperatorNo/{OperatorNo}][%d] interCityBusApiStationByOperator2891Status299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIStationByOperator2891Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIStationByOperator2891Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIStationByOperator2891NotModified creates a InterCityBusAPIStationByOperator2891NotModified with default headers values
func NewInterCityBusAPIStationByOperator2891NotModified() *InterCityBusAPIStationByOperator2891NotModified {
	return &InterCityBusAPIStationByOperator2891NotModified{}
}

/* InterCityBusAPIStationByOperator2891NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIStationByOperator2891NotModified struct {
}

func (o *InterCityBusAPIStationByOperator2891NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Station/InterCity/OperatorNo/{OperatorNo}][%d] interCityBusApiStationByOperator2891NotModified ", 304)
}

func (o *InterCityBusAPIStationByOperator2891NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}