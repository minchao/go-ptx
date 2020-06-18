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

// ServiceDTOVersion3BusScheduleStopTime StopTime
//
// swagger:model Service.DTO.Version3.Bus.Schedule.StopTime
type ServiceDTOVersion3BusScheduleStopTime struct {

	// 到站時間，格式為:HH:mm
	// Required: true
	ArrivalTime *string `json:"ArrivalTime"`

	// 離站時間，格式為:HH:mm
	// Required: true
	DepartureTime *string `json:"DepartureTime"`

	// 地區既用中之站牌代碼(為原資料內碼)
	// Required: true
	StopID *string `json:"StopID"`

	// NameType
	//
	// 站牌名稱
	StopName *ServiceDTOVersion3BaseNameType `json:"StopName,omitempty"`

	// 路線經過站牌之順序(由1開始)
	// Required: true
	StopSequence *int32 `json:"StopSequence"`

	// 站牌唯一識別代碼，規則為 {業管機關代碼} + {StopID}，其中 {業管機關代碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	StopUID *string `json:"StopUID"`
}

// Validate validates this service d t o version3 bus schedule stop time
func (m *ServiceDTOVersion3BusScheduleStopTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArrivalTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartureTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopSequence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3BusScheduleStopTime) validateArrivalTime(formats strfmt.Registry) error {

	if err := validate.Required("ArrivalTime", "body", m.ArrivalTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusScheduleStopTime) validateDepartureTime(formats strfmt.Registry) error {

	if err := validate.Required("DepartureTime", "body", m.DepartureTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusScheduleStopTime) validateStopID(formats strfmt.Registry) error {

	if err := validate.Required("StopID", "body", m.StopID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusScheduleStopTime) validateStopName(formats strfmt.Registry) error {

	if swag.IsZero(m.StopName) { // not required
		return nil
	}

	if m.StopName != nil {
		if err := m.StopName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StopName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion3BusScheduleStopTime) validateStopSequence(formats strfmt.Registry) error {

	if err := validate.Required("StopSequence", "body", m.StopSequence); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusScheduleStopTime) validateStopUID(formats strfmt.Registry) error {

	if err := validate.Required("StopUID", "body", m.StopUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3BusScheduleStopTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3BusScheduleStopTime) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3BusScheduleStopTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}