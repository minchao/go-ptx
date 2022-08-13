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

// DailyStationTimeTableAPIControllerGet32121Reader is a Reader for the DailyStationTimeTableAPIControllerGet32121 structure.
type DailyStationTimeTableAPIControllerGet32121Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DailyStationTimeTableAPIControllerGet32121Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDailyStationTimeTableAPIControllerGet32121OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewDailyStationTimeTableAPIControllerGet32121Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDailyStationTimeTableAPIControllerGet32121OK creates a DailyStationTimeTableAPIControllerGet32121OK with default headers values
func NewDailyStationTimeTableAPIControllerGet32121OK() *DailyStationTimeTableAPIControllerGet32121OK {
	return &DailyStationTimeTableAPIControllerGet32121OK{}
}

/* DailyStationTimeTableAPIControllerGet32121OK describes a response with status code 200, with default header values.

Success
*/
type DailyStationTimeTableAPIControllerGet32121OK struct {
	Payload *models.PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable
}

func (o *DailyStationTimeTableAPIControllerGet32121OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/DailyStationTimetable/Today/Station/{StationID}][%d] dailyStationTimeTableApiControllerGet32121OK  %+v", 200, o.Payload)
}
func (o *DailyStationTimeTableAPIControllerGet32121OK) GetPayload() *models.PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable {
	return o.Payload
}

func (o *DailyStationTimeTableAPIControllerGet32121OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDailyStationTimeTableAPIControllerGet32121Status299 creates a DailyStationTimeTableAPIControllerGet32121Status299 with default headers values
func NewDailyStationTimeTableAPIControllerGet32121Status299() *DailyStationTimeTableAPIControllerGet32121Status299 {
	return &DailyStationTimeTableAPIControllerGet32121Status299{}
}

/* DailyStationTimeTableAPIControllerGet32121Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type DailyStationTimeTableAPIControllerGet32121Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *DailyStationTimeTableAPIControllerGet32121Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/DailyStationTimetable/Today/Station/{StationID}][%d] dailyStationTimeTableApiControllerGet32121Status299  %+v", 299, o.Payload)
}
func (o *DailyStationTimeTableAPIControllerGet32121Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *DailyStationTimeTableAPIControllerGet32121Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
