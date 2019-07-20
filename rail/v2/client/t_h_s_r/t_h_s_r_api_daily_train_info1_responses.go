// Code generated by go-swagger; DO NOT EDIT.

package t_h_s_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/rail/v2/models"
)

// THSRAPIDailyTrainInfo1Reader is a Reader for the THSRAPIDailyTrainInfo1 structure.
type THSRAPIDailyTrainInfo1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIDailyTrainInfo1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIDailyTrainInfo1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTHSRAPIDailyTrainInfo1OK creates a THSRAPIDailyTrainInfo1OK with default headers values
func NewTHSRAPIDailyTrainInfo1OK() *THSRAPIDailyTrainInfo1OK {
	return &THSRAPIDailyTrainInfo1OK{}
}

/*THSRAPIDailyTrainInfo1OK handles this case with default header values.

OK
*/
type THSRAPIDailyTrainInfo1OK struct {
	Payload []*models.ServiceDTOVersion2RailTHSRRailDailyTrainInfo
}

func (o *THSRAPIDailyTrainInfo1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTrainInfo/Today/TrainNo/{TrainNo}][%d] tHSRApiDailyTrainInfo1OK  %+v", 200, o.Payload)
}

func (o *THSRAPIDailyTrainInfo1OK) GetPayload() []*models.ServiceDTOVersion2RailTHSRRailDailyTrainInfo {
	return o.Payload
}

func (o *THSRAPIDailyTrainInfo1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
