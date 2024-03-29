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

// PTXServiceDTORailSpecificationV3TRAGeneralStationTimetableServiceDay ServiceDay
//
// 服務日型態
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.GeneralStationTimetable.ServiceDay
type PTXServiceDTORailSpecificationV3TRAGeneralStationTimetableServiceDay struct {

	// 假日後一日營運與否
	DayAfterHoliday bool `json:"DayAfterHoliday,omitempty"`

	// 假日前一日營運與否
	DayBeforeHoliday bool `json:"DayBeforeHoliday,omitempty"`

	// Boolean
	//
	// 星期五營運與否
	// Required: true
	Friday *bool `json:"Friday"`

	// Boolean
	//
	// 星期一營運與否
	// Required: true
	Monday *bool `json:"Monday"`

	// 國定假日營運與否
	NationalHolidays bool `json:"NationalHolidays,omitempty"`

	// Boolean
	//
	// 星期六營運與否
	// Required: true
	Saturday *bool `json:"Saturday"`

	// String
	//
	// 服務日標籤
	ServiceTag string `json:"ServiceTag,omitempty" xml:"ServiceTag,omitempty"`

	// Boolean
	//
	// 星期日營運與否
	// Required: true
	Sunday *bool `json:"Sunday"`

	// Boolean
	//
	// 星期四營運與否
	// Required: true
	Thursday *bool `json:"Thursday"`

	// Boolean
	//
	// 星期二營運與否
	// Required: true
	Tuesday *bool `json:"Tuesday"`

	// 颱風停止上班上課期間營運與否
	TyphoonDay bool `json:"TyphoonDay,omitempty"`

	// Boolean
	//
	// 星期三營運與否
	// Required: true
	Wednesday *bool `json:"Wednesday"`
}

// Validate validates this p t x service d t o rail specification v3 t r a general station timetable service day
func (m *PTXServiceDTORailSpecificationV3TRAGeneralStationTimetableServiceDay) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFriday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSaturday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSunday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThursday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTuesday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWednesday(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAGeneralStationTimetableServiceDay) validateFriday(formats strfmt.Registry) error {

	if err := validate.Required("Friday", "body", m.Friday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAGeneralStationTimetableServiceDay) validateMonday(formats strfmt.Registry) error {

	if err := validate.Required("Monday", "body", m.Monday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAGeneralStationTimetableServiceDay) validateSaturday(formats strfmt.Registry) error {

	if err := validate.Required("Saturday", "body", m.Saturday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAGeneralStationTimetableServiceDay) validateSunday(formats strfmt.Registry) error {

	if err := validate.Required("Sunday", "body", m.Sunday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAGeneralStationTimetableServiceDay) validateThursday(formats strfmt.Registry) error {

	if err := validate.Required("Thursday", "body", m.Thursday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAGeneralStationTimetableServiceDay) validateTuesday(formats strfmt.Registry) error {

	if err := validate.Required("Tuesday", "body", m.Tuesday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAGeneralStationTimetableServiceDay) validateWednesday(formats strfmt.Registry) error {

	if err := validate.Required("Wednesday", "body", m.Wednesday); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o rail specification v3 t r a general station timetable service day based on context it is used
func (m *PTXServiceDTORailSpecificationV3TRAGeneralStationTimetableServiceDay) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAGeneralStationTimetableServiceDay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAGeneralStationTimetableServiceDay) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRAGeneralStationTimetableServiceDay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
