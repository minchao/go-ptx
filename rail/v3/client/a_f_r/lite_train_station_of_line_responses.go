// Code generated by go-swagger; DO NOT EDIT.

package a_f_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v3/models"
)

// LiteTrainStationOfLineReader is a Reader for the LiteTrainStationOfLine structure.
type LiteTrainStationOfLineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LiteTrainStationOfLineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLiteTrainStationOfLineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewLiteTrainStationOfLineNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLiteTrainStationOfLineOK creates a LiteTrainStationOfLineOK with default headers values
func NewLiteTrainStationOfLineOK() *LiteTrainStationOfLineOK {
	return &LiteTrainStationOfLineOK{}
}

/* LiteTrainStationOfLineOK describes a response with status code 200, with default header values.

Success
*/
type LiteTrainStationOfLineOK struct {
	Payload *models.PTXAPIRailModelLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainStationOfLineStationOfLine
}

func (o *LiteTrainStationOfLineOK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/StationOfLine][%d] liteTrainStationOfLineOK  %+v", 200, o.Payload)
}
func (o *LiteTrainStationOfLineOK) GetPayload() *models.PTXAPIRailModelLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainStationOfLineStationOfLine {
	return o.Payload
}

func (o *LiteTrainStationOfLineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainStationOfLineStationOfLine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLiteTrainStationOfLineNotModified creates a LiteTrainStationOfLineNotModified with default headers values
func NewLiteTrainStationOfLineNotModified() *LiteTrainStationOfLineNotModified {
	return &LiteTrainStationOfLineNotModified{}
}

/* LiteTrainStationOfLineNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type LiteTrainStationOfLineNotModified struct {
}

func (o *LiteTrainStationOfLineNotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/StationOfLine][%d] liteTrainStationOfLineNotModified ", 304)
}

func (o *LiteTrainStationOfLineNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
