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

// InterCityBusAPIStopByOperator2890Reader is a Reader for the InterCityBusAPIStopByOperator2890 structure.
type InterCityBusAPIStopByOperator2890Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIStopByOperator2890Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIStopByOperator2890OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIStopByOperator2890Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewInterCityBusAPIStopByOperator2890NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIStopByOperator2890OK creates a InterCityBusAPIStopByOperator2890OK with default headers values
func NewInterCityBusAPIStopByOperator2890OK() *InterCityBusAPIStopByOperator2890OK {
	return &InterCityBusAPIStopByOperator2890OK{}
}

/* InterCityBusAPIStopByOperator2890OK describes a response with status code 200, with default header values.

Success
*/
type InterCityBusAPIStopByOperator2890OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusStop
}

func (o *InterCityBusAPIStopByOperator2890OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Stop/InterCity/OperatorNo/{OperatorNo}][%d] interCityBusApiStopByOperator2890OK  %+v", 200, o.Payload)
}
func (o *InterCityBusAPIStopByOperator2890OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusStop {
	return o.Payload
}

func (o *InterCityBusAPIStopByOperator2890OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIStopByOperator2890Status299 creates a InterCityBusAPIStopByOperator2890Status299 with default headers values
func NewInterCityBusAPIStopByOperator2890Status299() *InterCityBusAPIStopByOperator2890Status299 {
	return &InterCityBusAPIStopByOperator2890Status299{}
}

/* InterCityBusAPIStopByOperator2890Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIStopByOperator2890Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIStopByOperator2890Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Stop/InterCity/OperatorNo/{OperatorNo}][%d] interCityBusApiStopByOperator2890Status299  %+v", 299, o.Payload)
}
func (o *InterCityBusAPIStopByOperator2890Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIStopByOperator2890Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIStopByOperator2890NotModified creates a InterCityBusAPIStopByOperator2890NotModified with default headers values
func NewInterCityBusAPIStopByOperator2890NotModified() *InterCityBusAPIStopByOperator2890NotModified {
	return &InterCityBusAPIStopByOperator2890NotModified{}
}

/* InterCityBusAPIStopByOperator2890NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type InterCityBusAPIStopByOperator2890NotModified struct {
}

func (o *InterCityBusAPIStopByOperator2890NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Stop/InterCity/OperatorNo/{OperatorNo}][%d] interCityBusApiStopByOperator2890NotModified ", 304)
}

func (o *InterCityBusAPIStopByOperator2890NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
