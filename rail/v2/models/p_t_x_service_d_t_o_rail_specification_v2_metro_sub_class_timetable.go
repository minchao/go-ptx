// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTORailSpecificationV2MetroSubClassTimetable Timetable
//
// 車站發車時刻資訊
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.Metro.SubClass.Timetable
type PTXServiceDTORailSpecificationV2MetroSubClassTimetable struct {

	// String
	//
	// 到站時刻(hh:mm)
	// Required: true
	ArrivalTime *string `json:"ArrivalTime"`

	// String
	//
	// 發車時刻(hh:mm)
	// Required: true
	DepartureTime *string `json:"DepartureTime"`

	// Int32
	//
	// 發車順序
	// Required: true
	Sequence *int32 `json:"Sequence"`

	// String
	//
	// 車次號碼(捷運通常沒有TrainNo車次資訊)
	// Required: true
	TrainNo *string `json:"TrainNo"`

	// 車種(0:不分車種, 1:普通車, 2:直達車)
	TrainType int32 `json:"TrainType,omitempty"`
}

// Validate validates this p t x service d t o rail specification v2 metro sub class timetable
func (m *PTXServiceDTORailSpecificationV2MetroSubClassTimetable) Validate(formats strfmt.Registry) error {
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

	if err := m.validateTrainNo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassTimetable) validateArrivalTime(formats strfmt.Registry) error {

	if err := validate.Required("ArrivalTime", "body", m.ArrivalTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassTimetable) validateDepartureTime(formats strfmt.Registry) error {

	if err := validate.Required("DepartureTime", "body", m.DepartureTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassTimetable) validateSequence(formats strfmt.Registry) error {

	if err := validate.Required("Sequence", "body", m.Sequence); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassTimetable) validateTrainNo(formats strfmt.Registry) error {

	if err := validate.Required("TrainNo", "body", m.TrainNo); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroSubClassTimetable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroSubClassTimetable) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2MetroSubClassTimetable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}