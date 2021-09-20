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

// DailyTrainTimeTableAPIControllerGet5Reader is a Reader for the DailyTrainTimeTableAPIControllerGet5 structure.
type DailyTrainTimeTableAPIControllerGet5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DailyTrainTimeTableAPIControllerGet5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDailyTrainTimeTableAPIControllerGet5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewDailyTrainTimeTableAPIControllerGet5Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewDailyTrainTimeTableAPIControllerGet5NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDailyTrainTimeTableAPIControllerGet5OK creates a DailyTrainTimeTableAPIControllerGet5OK with default headers values
func NewDailyTrainTimeTableAPIControllerGet5OK() *DailyTrainTimeTableAPIControllerGet5OK {
	return &DailyTrainTimeTableAPIControllerGet5OK{}
}

/* DailyTrainTimeTableAPIControllerGet5OK describes a response with status code 200, with default header values.

Success
*/
type DailyTrainTimeTableAPIControllerGet5OK struct {
	Payload *models.PTXAPIRailModelTraDailyTrainWrapperPTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainTimetable
}

func (o *DailyTrainTimeTableAPIControllerGet5OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/DailyTrainTimetable/OD/Inclusive/{OriginStationID}/to/{DestinationStationID}/{TrainDate}][%d] dailyTrainTimeTableApiControllerGet5OK  %+v", 200, o.Payload)
}
func (o *DailyTrainTimeTableAPIControllerGet5OK) GetPayload() *models.PTXAPIRailModelTraDailyTrainWrapperPTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainTimetable {
	return o.Payload
}

func (o *DailyTrainTimeTableAPIControllerGet5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTraDailyTrainWrapperPTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainTimetable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDailyTrainTimeTableAPIControllerGet5Status299 creates a DailyTrainTimeTableAPIControllerGet5Status299 with default headers values
func NewDailyTrainTimeTableAPIControllerGet5Status299() *DailyTrainTimeTableAPIControllerGet5Status299 {
	return &DailyTrainTimeTableAPIControllerGet5Status299{}
}

/* DailyTrainTimeTableAPIControllerGet5Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type DailyTrainTimeTableAPIControllerGet5Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *DailyTrainTimeTableAPIControllerGet5Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/DailyTrainTimetable/OD/Inclusive/{OriginStationID}/to/{DestinationStationID}/{TrainDate}][%d] dailyTrainTimeTableApiControllerGet5Status299  %+v", 299, o.Payload)
}
func (o *DailyTrainTimeTableAPIControllerGet5Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *DailyTrainTimeTableAPIControllerGet5Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDailyTrainTimeTableAPIControllerGet5NotModified creates a DailyTrainTimeTableAPIControllerGet5NotModified with default headers values
func NewDailyTrainTimeTableAPIControllerGet5NotModified() *DailyTrainTimeTableAPIControllerGet5NotModified {
	return &DailyTrainTimeTableAPIControllerGet5NotModified{}
}

/* DailyTrainTimeTableAPIControllerGet5NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type DailyTrainTimeTableAPIControllerGet5NotModified struct {
}

func (o *DailyTrainTimeTableAPIControllerGet5NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/DailyTrainTimetable/OD/Inclusive/{OriginStationID}/to/{DestinationStationID}/{TrainDate}][%d] dailyTrainTimeTableApiControllerGet5NotModified ", 304)
}

func (o *DailyTrainTimeTableAPIControllerGet5NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
