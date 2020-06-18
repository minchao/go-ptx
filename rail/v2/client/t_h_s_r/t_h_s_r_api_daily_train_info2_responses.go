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

// THSRAPIDailyTrainInfo2Reader is a Reader for the THSRAPIDailyTrainInfo2 structure.
type THSRAPIDailyTrainInfo2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIDailyTrainInfo2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIDailyTrainInfo2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTHSRAPIDailyTrainInfo2OK creates a THSRAPIDailyTrainInfo2OK with default headers values
func NewTHSRAPIDailyTrainInfo2OK() *THSRAPIDailyTrainInfo2OK {
	return &THSRAPIDailyTrainInfo2OK{}
}

/*THSRAPIDailyTrainInfo2OK handles this case with default header values.

OK
*/
type THSRAPIDailyTrainInfo2OK struct {
	Payload []*models.ServiceDTOVersion2RailTHSRRailDailyTrainInfo
}

func (o *THSRAPIDailyTrainInfo2OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTrainInfo/TrainDate/{TrainDate}][%d] tHSRApiDailyTrainInfo2OK  %+v", 200, o.Payload)
}

func (o *THSRAPIDailyTrainInfo2OK) GetPayload() []*models.ServiceDTOVersion2RailTHSRRailDailyTrainInfo {
	return o.Payload
}

func (o *THSRAPIDailyTrainInfo2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
