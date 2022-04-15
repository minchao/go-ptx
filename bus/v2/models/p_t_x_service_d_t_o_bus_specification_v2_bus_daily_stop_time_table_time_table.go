// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTOBusSpecificationV2BusDailyStopTimeTableTimeTable TimeTable
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.BusDailyStopTimeTable+TimeTable
type PTXServiceDTOBusSpecificationV2BusDailyStopTimeTableTimeTable struct {

	// String
	//
	// 到站時間，格式為:HH:mm
	// Required: true
	ArrivalTime *string `json:"ArrivalTime" xml:"String"`

	// String
	//
	// 離站時間，格式為:HH:mm
	// Required: true
	DepartureTime *string `json:"DepartureTime" xml:"String"`

	// 該路線班次是否使用低地板公車車輛 , 0 = 否 , 1 = 是
	IsLowFloor bool `json:"IsLowFloor,omitempty"`

	// Int32
	//
	// 發車順序(由1開始)
	// Required: true
	Sequence *int32 `json:"Sequence"`

	// Int32
	//
	// 時刻表時間類型 : [0:'估算時間',1:'表訂時間']
	// Required: true
	TimeType *int64 `json:"TimeType"`

	// String
	//
	// 班次代碼
	TripID string `json:"TripID,omitempty" xml:"String,omitempty"`
}

// Validate validates this p t x service d t o bus specification v2 bus daily stop time table time table
func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTableTimeTable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArrivalTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartureTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSequence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTableTimeTable) validateArrivalTime(formats strfmt.Registry) error {

	if err := validate.Required("ArrivalTime", "body", m.ArrivalTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTableTimeTable) validateDepartureTime(formats strfmt.Registry) error {

	if err := validate.Required("DepartureTime", "body", m.DepartureTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTableTimeTable) validateSequence(formats strfmt.Registry) error {

	if err := validate.Required("Sequence", "body", m.Sequence); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTableTimeTable) validateTimeType(formats strfmt.Registry) error {

	if err := validate.Required("TimeType", "body", m.TimeType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o bus specification v2 bus daily stop time table time table based on context it is used
func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTableTimeTable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTableTimeTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTableTimeTable) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2BusDailyStopTimeTableTimeTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}