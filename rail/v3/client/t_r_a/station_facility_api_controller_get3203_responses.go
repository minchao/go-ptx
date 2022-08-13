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

// StationFacilityAPIControllerGet3203Reader is a Reader for the StationFacilityAPIControllerGet3203 structure.
type StationFacilityAPIControllerGet3203Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StationFacilityAPIControllerGet3203Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStationFacilityAPIControllerGet3203OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewStationFacilityAPIControllerGet3203Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStationFacilityAPIControllerGet3203OK creates a StationFacilityAPIControllerGet3203OK with default headers values
func NewStationFacilityAPIControllerGet3203OK() *StationFacilityAPIControllerGet3203OK {
	return &StationFacilityAPIControllerGet3203OK{}
}

/* StationFacilityAPIControllerGet3203OK describes a response with status code 200, with default header values.

Success
*/
type StationFacilityAPIControllerGet3203OK struct {
	Payload *models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility
}

func (o *StationFacilityAPIControllerGet3203OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/StationFacility][%d] stationFacilityApiControllerGet3203OK  %+v", 200, o.Payload)
}
func (o *StationFacilityAPIControllerGet3203OK) GetPayload() *models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility {
	return o.Payload
}

func (o *StationFacilityAPIControllerGet3203OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStationFacilityAPIControllerGet3203Status299 creates a StationFacilityAPIControllerGet3203Status299 with default headers values
func NewStationFacilityAPIControllerGet3203Status299() *StationFacilityAPIControllerGet3203Status299 {
	return &StationFacilityAPIControllerGet3203Status299{}
}

/* StationFacilityAPIControllerGet3203Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type StationFacilityAPIControllerGet3203Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *StationFacilityAPIControllerGet3203Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/StationFacility][%d] stationFacilityApiControllerGet3203Status299  %+v", 299, o.Payload)
}
func (o *StationFacilityAPIControllerGet3203Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *StationFacilityAPIControllerGet3203Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
