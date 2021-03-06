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

// SpecificTrainTimetableAPIControllerGet1Reader is a Reader for the SpecificTrainTimetableAPIControllerGet1 structure.
type SpecificTrainTimetableAPIControllerGet1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SpecificTrainTimetableAPIControllerGet1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSpecificTrainTimetableAPIControllerGet1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewSpecificTrainTimetableAPIControllerGet1NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSpecificTrainTimetableAPIControllerGet1OK creates a SpecificTrainTimetableAPIControllerGet1OK with default headers values
func NewSpecificTrainTimetableAPIControllerGet1OK() *SpecificTrainTimetableAPIControllerGet1OK {
	return &SpecificTrainTimetableAPIControllerGet1OK{}
}

/* SpecificTrainTimetableAPIControllerGet1OK describes a response with status code 200, with default header values.

Success
*/
type SpecificTrainTimetableAPIControllerGet1OK struct {
	Payload *models.PTXAPIRailModelTRASpecificWrapperPTXServiceDTORailSpecificationV3TRASpecificTrainTimetable
}

func (o *SpecificTrainTimetableAPIControllerGet1OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/SpecificTrainTimetable/TrainNo/{TrainNo}][%d] specificTrainTimetableApiControllerGet1OK  %+v", 200, o.Payload)
}
func (o *SpecificTrainTimetableAPIControllerGet1OK) GetPayload() *models.PTXAPIRailModelTRASpecificWrapperPTXServiceDTORailSpecificationV3TRASpecificTrainTimetable {
	return o.Payload
}

func (o *SpecificTrainTimetableAPIControllerGet1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTRASpecificWrapperPTXServiceDTORailSpecificationV3TRASpecificTrainTimetable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecificTrainTimetableAPIControllerGet1NotModified creates a SpecificTrainTimetableAPIControllerGet1NotModified with default headers values
func NewSpecificTrainTimetableAPIControllerGet1NotModified() *SpecificTrainTimetableAPIControllerGet1NotModified {
	return &SpecificTrainTimetableAPIControllerGet1NotModified{}
}

/* SpecificTrainTimetableAPIControllerGet1NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type SpecificTrainTimetableAPIControllerGet1NotModified struct {
}

func (o *SpecificTrainTimetableAPIControllerGet1NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/SpecificTrainTimetable/TrainNo/{TrainNo}][%d] specificTrainTimetableApiControllerGet1NotModified ", 304)
}

func (o *SpecificTrainTimetableAPIControllerGet1NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
