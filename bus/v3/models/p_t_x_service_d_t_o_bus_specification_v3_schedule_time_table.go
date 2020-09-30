// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTOBusSpecificationV3ScheduleTimeTable TimeTable
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.Schedule+TimeTable
type PTXServiceDTOBusSpecificationV3ScheduleTimeTable struct {

	// String
	//
	// 該路線班次是否使用低地板公車車輛
	IsLowFloor string `json:"IsLowFloor,omitempty"`

	// ServiceDay
	//
	// 服務日型態
	// Required: true
	ServiceDay struct {
		PTXServiceDTOBusSpecificationV3ScheduleServiceDay
	} `json:"ServiceDay"`

	// Array
	//
	// 例外營運日
	// Required: true
	SpecialDays []*PTXServiceDTOBusSpecificationV3ScheduleSpecialDay `json:"SpecialDays"`

	// Array
	//
	// 公車停靠時間資料
	// Required: true
	StopTimes []*PTXServiceDTOBusSpecificationV3ScheduleStopTime `json:"StopTimes"`

	// String
	//
	// 班次代碼，為無意義之編碼
	TripID string `json:"TripID,omitempty"`
}

// Validate validates this p t x service d t o bus specification v3 schedule time table
func (m *PTXServiceDTOBusSpecificationV3ScheduleTimeTable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceDay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpecialDays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopTimes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3ScheduleTimeTable) validateServiceDay(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3ScheduleTimeTable) validateSpecialDays(formats strfmt.Registry) error {

	if err := validate.Required("SpecialDays", "body", m.SpecialDays); err != nil {
		return err
	}

	for i := 0; i < len(m.SpecialDays); i++ {
		if swag.IsZero(m.SpecialDays[i]) { // not required
			continue
		}

		if m.SpecialDays[i] != nil {
			if err := m.SpecialDays[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SpecialDays" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3ScheduleTimeTable) validateStopTimes(formats strfmt.Registry) error {

	if err := validate.Required("StopTimes", "body", m.StopTimes); err != nil {
		return err
	}

	for i := 0; i < len(m.StopTimes); i++ {
		if swag.IsZero(m.StopTimes[i]) { // not required
			continue
		}

		if m.StopTimes[i] != nil {
			if err := m.StopTimes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StopTimes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3ScheduleTimeTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3ScheduleTimeTable) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3ScheduleTimeTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}