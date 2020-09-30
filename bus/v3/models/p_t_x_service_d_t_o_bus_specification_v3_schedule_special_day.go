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

// PTXServiceDTOBusSpecificationV3ScheduleSpecialDay SpecialDay
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.Schedule+SpecialDay
type PTXServiceDTOBusSpecificationV3ScheduleSpecialDay struct {

	// DatePeriod
	//
	// 連續特殊日期
	// Required: true
	DatePeriod struct {
		PTXServiceDTOBusSpecificationV3ScheduleDatePeriod
	} `json:"DatePeriod"`

	// Array
	//
	// 不連續特殊日期
	// Required: true
	Dates []string `json:"Dates"`

	// String
	//
	// 特殊營運描述
	// Required: true
	Description *string `json:"Description"`

	// integer
	//
	// 營運服務狀態代碼0=停止營運, 1=正常營運 , 2=加班營運 : [0:'正常營運',1:'加班營運',2:'取消/停駛營運']
	// Required: true
	ServiceStatus *int32 `json:"ServiceStatus"`
}

// Validate validates this p t x service d t o bus specification v3 schedule special day
func (m *PTXServiceDTOBusSpecificationV3ScheduleSpecialDay) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatePeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3ScheduleSpecialDay) validateDatePeriod(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3ScheduleSpecialDay) validateDates(formats strfmt.Registry) error {

	if err := validate.Required("Dates", "body", m.Dates); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3ScheduleSpecialDay) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3ScheduleSpecialDay) validateServiceStatus(formats strfmt.Registry) error {

	if err := validate.Required("ServiceStatus", "body", m.ServiceStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3ScheduleSpecialDay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3ScheduleSpecialDay) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3ScheduleSpecialDay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}