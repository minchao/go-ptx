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

// DailyTrainTimeTableAPIControllerGet32114Reader is a Reader for the DailyTrainTimeTableAPIControllerGet32114 structure.
type DailyTrainTimeTableAPIControllerGet32114Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DailyTrainTimeTableAPIControllerGet32114Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDailyTrainTimeTableAPIControllerGet32114OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewDailyTrainTimeTableAPIControllerGet32114Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewDailyTrainTimeTableAPIControllerGet32114NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDailyTrainTimeTableAPIControllerGet32114OK creates a DailyTrainTimeTableAPIControllerGet32114OK with default headers values
func NewDailyTrainTimeTableAPIControllerGet32114OK() *DailyTrainTimeTableAPIControllerGet32114OK {
	return &DailyTrainTimeTableAPIControllerGet32114OK{}
}

/* DailyTrainTimeTableAPIControllerGet32114OK describes a response with status code 200, with default header values.

Success
*/
type DailyTrainTimeTableAPIControllerGet32114OK struct {
	Payload *models.PTXAPIRailModelTraDailyTrainWrapperPTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainTimetable
}

func (o *DailyTrainTimeTableAPIControllerGet32114OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/DailyTrainTimetable/OD/{OriginStationID}/to/{DestinationStationID}/{TrainDate}][%d] dailyTrainTimeTableApiControllerGet32114OK  %+v", 200, o.Payload)
}
func (o *DailyTrainTimeTableAPIControllerGet32114OK) GetPayload() *models.PTXAPIRailModelTraDailyTrainWrapperPTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainTimetable {
	return o.Payload
}

func (o *DailyTrainTimeTableAPIControllerGet32114OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTraDailyTrainWrapperPTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainTimetable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDailyTrainTimeTableAPIControllerGet32114Status299 creates a DailyTrainTimeTableAPIControllerGet32114Status299 with default headers values
func NewDailyTrainTimeTableAPIControllerGet32114Status299() *DailyTrainTimeTableAPIControllerGet32114Status299 {
	return &DailyTrainTimeTableAPIControllerGet32114Status299{}
}

/* DailyTrainTimeTableAPIControllerGet32114Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type DailyTrainTimeTableAPIControllerGet32114Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *DailyTrainTimeTableAPIControllerGet32114Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/DailyTrainTimetable/OD/{OriginStationID}/to/{DestinationStationID}/{TrainDate}][%d] dailyTrainTimeTableApiControllerGet32114Status299  %+v", 299, o.Payload)
}
func (o *DailyTrainTimeTableAPIControllerGet32114Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *DailyTrainTimeTableAPIControllerGet32114Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDailyTrainTimeTableAPIControllerGet32114NotModified creates a DailyTrainTimeTableAPIControllerGet32114NotModified with default headers values
func NewDailyTrainTimeTableAPIControllerGet32114NotModified() *DailyTrainTimeTableAPIControllerGet32114NotModified {
	return &DailyTrainTimeTableAPIControllerGet32114NotModified{}
}

/* DailyTrainTimeTableAPIControllerGet32114NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type DailyTrainTimeTableAPIControllerGet32114NotModified struct {
}

func (o *DailyTrainTimeTableAPIControllerGet32114NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/DailyTrainTimetable/OD/{OriginStationID}/to/{DestinationStationID}/{TrainDate}][%d] dailyTrainTimeTableApiControllerGet32114NotModified ", 304)
}

func (o *DailyTrainTimeTableAPIControllerGet32114NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
