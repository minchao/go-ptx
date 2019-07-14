// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/rail/models"
)

// DailyStationTimeTableAPIControllerGet1Reader is a Reader for the DailyStationTimeTableAPIControllerGet1 structure.
type DailyStationTimeTableAPIControllerGet1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DailyStationTimeTableAPIControllerGet1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDailyStationTimeTableAPIControllerGet1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDailyStationTimeTableAPIControllerGet1OK creates a DailyStationTimeTableAPIControllerGet1OK with default headers values
func NewDailyStationTimeTableAPIControllerGet1OK() *DailyStationTimeTableAPIControllerGet1OK {
	return &DailyStationTimeTableAPIControllerGet1OK{}
}

/*DailyStationTimeTableAPIControllerGet1OK handles this case with default header values.

OK
*/
type DailyStationTimeTableAPIControllerGet1OK struct {
	Payload *models.MOTCAPIRailModelsTraDailyStationWrapperServiceDTOVersion3RailTRADailyStationTimeTableStationTimetable
}

func (o *DailyStationTimeTableAPIControllerGet1OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/DailyStationTimetable/Today/Station/{StationID}][%d] dailyStationTimeTableApiControllerGet1OK  %+v", 200, o.Payload)
}

func (o *DailyStationTimeTableAPIControllerGet1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MOTCAPIRailModelsTraDailyStationWrapperServiceDTOVersion3RailTRADailyStationTimeTableStationTimetable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
