// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v3/models"
)

// StationLiveBoardAPIControllerGet32131Reader is a Reader for the StationLiveBoardAPIControllerGet32131 structure.
type StationLiveBoardAPIControllerGet32131Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StationLiveBoardAPIControllerGet32131Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStationLiveBoardAPIControllerGet32131OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewStationLiveBoardAPIControllerGet32131Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewStationLiveBoardAPIControllerGet32131NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStationLiveBoardAPIControllerGet32131OK creates a StationLiveBoardAPIControllerGet32131OK with default headers values
func NewStationLiveBoardAPIControllerGet32131OK() *StationLiveBoardAPIControllerGet32131OK {
	return &StationLiveBoardAPIControllerGet32131OK{}
}

/* StationLiveBoardAPIControllerGet32131OK describes a response with status code 200, with default header values.

Success
*/
type StationLiveBoardAPIControllerGet32131OK struct {
	Payload *models.PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard
}

func (o *StationLiveBoardAPIControllerGet32131OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/StationLiveBoard/Station/{StationID}][%d] stationLiveBoardApiControllerGet32131OK  %+v", 200, o.Payload)
}
func (o *StationLiveBoardAPIControllerGet32131OK) GetPayload() *models.PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard {
	return o.Payload
}

func (o *StationLiveBoardAPIControllerGet32131OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStationLiveBoardAPIControllerGet32131Status299 creates a StationLiveBoardAPIControllerGet32131Status299 with default headers values
func NewStationLiveBoardAPIControllerGet32131Status299() *StationLiveBoardAPIControllerGet32131Status299 {
	return &StationLiveBoardAPIControllerGet32131Status299{}
}

/* StationLiveBoardAPIControllerGet32131Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type StationLiveBoardAPIControllerGet32131Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *StationLiveBoardAPIControllerGet32131Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/StationLiveBoard/Station/{StationID}][%d] stationLiveBoardApiControllerGet32131Status299  %+v", 299, o.Payload)
}
func (o *StationLiveBoardAPIControllerGet32131Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *StationLiveBoardAPIControllerGet32131Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStationLiveBoardAPIControllerGet32131NotModified creates a StationLiveBoardAPIControllerGet32131NotModified with default headers values
func NewStationLiveBoardAPIControllerGet32131NotModified() *StationLiveBoardAPIControllerGet32131NotModified {
	return &StationLiveBoardAPIControllerGet32131NotModified{}
}

/* StationLiveBoardAPIControllerGet32131NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type StationLiveBoardAPIControllerGet32131NotModified struct {
}

func (o *StationLiveBoardAPIControllerGet32131NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/StationLiveBoard/Station/{StationID}][%d] stationLiveBoardApiControllerGet32131NotModified ", 304)
}

func (o *StationLiveBoardAPIControllerGet32131NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}