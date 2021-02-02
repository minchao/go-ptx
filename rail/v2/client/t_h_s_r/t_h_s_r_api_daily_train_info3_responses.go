// Code generated by go-swagger; DO NOT EDIT.

package t_h_s_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v2/models"
)

// THSRAPIDailyTrainInfo3Reader is a Reader for the THSRAPIDailyTrainInfo3 structure.
type THSRAPIDailyTrainInfo3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIDailyTrainInfo3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIDailyTrainInfo3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTHSRAPIDailyTrainInfo3NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTHSRAPIDailyTrainInfo3OK creates a THSRAPIDailyTrainInfo3OK with default headers values
func NewTHSRAPIDailyTrainInfo3OK() *THSRAPIDailyTrainInfo3OK {
	return &THSRAPIDailyTrainInfo3OK{}
}

/* THSRAPIDailyTrainInfo3OK describes a response with status code 200, with default header values.

Success
*/
type THSRAPIDailyTrainInfo3OK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2THSRRailDailyTrainInfo
}

func (o *THSRAPIDailyTrainInfo3OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTrainInfo/TrainNo/{TrainNo}/TrainDate/{TrainDate}][%d] tHSRApiDailyTrainInfo3OK  %+v", 200, o.Payload)
}
func (o *THSRAPIDailyTrainInfo3OK) GetPayload() []*models.PTXServiceDTORailSpecificationV2THSRRailDailyTrainInfo {
	return o.Payload
}

func (o *THSRAPIDailyTrainInfo3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPIDailyTrainInfo3NotModified creates a THSRAPIDailyTrainInfo3NotModified with default headers values
func NewTHSRAPIDailyTrainInfo3NotModified() *THSRAPIDailyTrainInfo3NotModified {
	return &THSRAPIDailyTrainInfo3NotModified{}
}

/* THSRAPIDailyTrainInfo3NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type THSRAPIDailyTrainInfo3NotModified struct {
}

func (o *THSRAPIDailyTrainInfo3NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTrainInfo/TrainNo/{TrainNo}/TrainDate/{TrainDate}][%d] tHSRApiDailyTrainInfo3NotModified ", 304)
}

func (o *THSRAPIDailyTrainInfo3NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
