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

// THSRAPIAvailableSeatStatusOD2Reader is a Reader for the THSRAPIAvailableSeatStatusOD2 structure.
type THSRAPIAvailableSeatStatusOD2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIAvailableSeatStatusOD2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIAvailableSeatStatusOD2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTHSRAPIAvailableSeatStatusOD2OK creates a THSRAPIAvailableSeatStatusOD2OK with default headers values
func NewTHSRAPIAvailableSeatStatusOD2OK() *THSRAPIAvailableSeatStatusOD2OK {
	return &THSRAPIAvailableSeatStatusOD2OK{}
}

/*THSRAPIAvailableSeatStatusOD2OK handles this case with default header values.

Success
*/
type THSRAPIAvailableSeatStatusOD2OK struct {
	Payload *models.PTXAPIRailModelV2THSRODAvaliableSeatStatusWrapperPTXServiceDTORailSpecificationV2THSRODAvailableSeat
}

func (o *THSRAPIAvailableSeatStatusOD2OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/AvailableSeatStatus/Train/OD/{OriginStationID}/to/{DestinationStationID}/TrainDate/{TrainDate}/TrainNo/{TrainNo}][%d] tHSRApiAvailableSeatStatusOD2OK  %+v", 200, o.Payload)
}

func (o *THSRAPIAvailableSeatStatusOD2OK) GetPayload() *models.PTXAPIRailModelV2THSRODAvaliableSeatStatusWrapperPTXServiceDTORailSpecificationV2THSRODAvailableSeat {
	return o.Payload
}

func (o *THSRAPIAvailableSeatStatusOD2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelV2THSRODAvaliableSeatStatusWrapperPTXServiceDTORailSpecificationV2THSRODAvailableSeat)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}