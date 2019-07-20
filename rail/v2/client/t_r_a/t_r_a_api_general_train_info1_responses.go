// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/rail/v2/models"
)

// TRAAPIGeneralTrainInfo1Reader is a Reader for the TRAAPIGeneralTrainInfo1 structure.
type TRAAPIGeneralTrainInfo1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TRAAPIGeneralTrainInfo1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTRAAPIGeneralTrainInfo1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTRAAPIGeneralTrainInfo1OK creates a TRAAPIGeneralTrainInfo1OK with default headers values
func NewTRAAPIGeneralTrainInfo1OK() *TRAAPIGeneralTrainInfo1OK {
	return &TRAAPIGeneralTrainInfo1OK{}
}

/*TRAAPIGeneralTrainInfo1OK handles this case with default header values.

OK
*/
type TRAAPIGeneralTrainInfo1OK struct {
	Payload []*models.ServiceDTOVersion2RailTRARailGeneralTrainInfo
}

func (o *TRAAPIGeneralTrainInfo1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/GeneralTrainInfo/TrainNo/{TrainNo}][%d] tRAApiGeneralTrainInfo1OK  %+v", 200, o.Payload)
}

func (o *TRAAPIGeneralTrainInfo1OK) GetPayload() []*models.ServiceDTOVersion2RailTRARailGeneralTrainInfo {
	return o.Payload
}

func (o *TRAAPIGeneralTrainInfo1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
