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

// DailyTrainTimeTableAPIControllerGet2Reader is a Reader for the DailyTrainTimeTableAPIControllerGet2 structure.
type DailyTrainTimeTableAPIControllerGet2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DailyTrainTimeTableAPIControllerGet2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDailyTrainTimeTableAPIControllerGet2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDailyTrainTimeTableAPIControllerGet2OK creates a DailyTrainTimeTableAPIControllerGet2OK with default headers values
func NewDailyTrainTimeTableAPIControllerGet2OK() *DailyTrainTimeTableAPIControllerGet2OK {
	return &DailyTrainTimeTableAPIControllerGet2OK{}
}

/*DailyTrainTimeTableAPIControllerGet2OK handles this case with default header values.

OK
*/
type DailyTrainTimeTableAPIControllerGet2OK struct {
	Payload *models.MOTCAPIRailModelsTraDailyTrainWrapperServiceDTOVersion3RailTRADailyTrainTimeTableTrainTimetable
}

func (o *DailyTrainTimeTableAPIControllerGet2OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/DailyTrainTimetable/TrainDate/{TrainDate}][%d] dailyTrainTimeTableApiControllerGet2OK  %+v", 200, o.Payload)
}

func (o *DailyTrainTimeTableAPIControllerGet2OK) GetPayload() *models.MOTCAPIRailModelsTraDailyTrainWrapperServiceDTOVersion3RailTRADailyTrainTimeTableTrainTimetable {
	return o.Payload
}

func (o *DailyTrainTimeTableAPIControllerGet2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MOTCAPIRailModelsTraDailyTrainWrapperServiceDTOVersion3RailTRADailyTrainTimeTableTrainTimetable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
