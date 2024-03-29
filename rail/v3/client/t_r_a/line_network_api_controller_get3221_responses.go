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

// LineNetworkAPIControllerGet3221Reader is a Reader for the LineNetworkAPIControllerGet3221 structure.
type LineNetworkAPIControllerGet3221Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LineNetworkAPIControllerGet3221Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLineNetworkAPIControllerGet3221OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewLineNetworkAPIControllerGet3221Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLineNetworkAPIControllerGet3221OK creates a LineNetworkAPIControllerGet3221OK with default headers values
func NewLineNetworkAPIControllerGet3221OK() *LineNetworkAPIControllerGet3221OK {
	return &LineNetworkAPIControllerGet3221OK{}
}

/* LineNetworkAPIControllerGet3221OK describes a response with status code 200, with default header values.

Success
*/
type LineNetworkAPIControllerGet3221OK struct {
	Payload *models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRALineNetworkLineNetwork
}

func (o *LineNetworkAPIControllerGet3221OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/LineNetwork][%d] lineNetworkApiControllerGet3221OK  %+v", 200, o.Payload)
}
func (o *LineNetworkAPIControllerGet3221OK) GetPayload() *models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRALineNetworkLineNetwork {
	return o.Payload
}

func (o *LineNetworkAPIControllerGet3221OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRALineNetworkLineNetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLineNetworkAPIControllerGet3221Status299 creates a LineNetworkAPIControllerGet3221Status299 with default headers values
func NewLineNetworkAPIControllerGet3221Status299() *LineNetworkAPIControllerGet3221Status299 {
	return &LineNetworkAPIControllerGet3221Status299{}
}

/* LineNetworkAPIControllerGet3221Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type LineNetworkAPIControllerGet3221Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *LineNetworkAPIControllerGet3221Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/LineNetwork][%d] lineNetworkApiControllerGet3221Status299  %+v", 299, o.Payload)
}
func (o *LineNetworkAPIControllerGet3221Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *LineNetworkAPIControllerGet3221Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
