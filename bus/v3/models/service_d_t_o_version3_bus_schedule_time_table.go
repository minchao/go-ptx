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

// ServiceDTOVersion3BusScheduleTimeTable TimeTable
//
// swagger:model Service.DTO.Version3.Bus.Schedule.TimeTable
type ServiceDTOVersion3BusScheduleTimeTable struct {

	// 該路線班次是否使用低地板公車車輛
	IsLowFloor string `json:"IsLowFloor,omitempty"`

	// ServiceDay
	//
	// 服務日型態
	// Required: true
	ServiceDay *ServiceDTOVersion3BusScheduleServiceDay `json:"ServiceDay"`

	// 例外營運日
	// Required: true
	SpecialDays []*ServiceDTOVersion3BusScheduleSpecialDay `json:"SpecialDays"`

	// 公車停靠時間資料
	// Required: true
	StopTimes []*ServiceDTOVersion3BusScheduleStopTime `json:"StopTimes"`

	// 班次代碼，為無意義之編碼
	TripID string `json:"TripID,omitempty"`
}

// Validate validates this service d t o version3 bus schedule time table
func (m *ServiceDTOVersion3BusScheduleTimeTable) Validate(formats strfmt.Registry) error {
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

func (m *ServiceDTOVersion3BusScheduleTimeTable) validateServiceDay(formats strfmt.Registry) error {

	if err := validate.Required("ServiceDay", "body", m.ServiceDay); err != nil {
		return err
	}

	if m.ServiceDay != nil {
		if err := m.ServiceDay.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ServiceDay")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion3BusScheduleTimeTable) validateSpecialDays(formats strfmt.Registry) error {

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

func (m *ServiceDTOVersion3BusScheduleTimeTable) validateStopTimes(formats strfmt.Registry) error {

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
func (m *ServiceDTOVersion3BusScheduleTimeTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3BusScheduleTimeTable) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3BusScheduleTimeTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
