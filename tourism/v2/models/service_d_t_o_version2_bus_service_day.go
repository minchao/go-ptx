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

// ServiceDTOVersion2BusServiceDay ServiceDay
//
// swagger:model Service.DTO.Version2.Bus.ServiceDay
type ServiceDTOVersion2BusServiceDay struct {

	// integer
	//
	// 星期五是否營運 : [0:'否',1:'是']
	// Required: true
	Friday *int32 `json:"Friday"`

	// integer
	//
	// 星期一是否營運 : [0:'否',1:'是']
	// Required: true
	Monday *int32 `json:"Monday"`

	// integer
	//
	// 國定假日是否營運 : [0:'否',1:'是']
	NationalHolidays int32 `json:"NationalHolidays,omitempty"`

	// integer
	//
	// 星期六是否營運 : [0:'否',1:'是']
	// Required: true
	Saturday *int32 `json:"Saturday"`

	// integer
	//
	// 星期日是否營運 : [0:'否',1:'是']
	// Required: true
	Sunday *int32 `json:"Sunday"`

	// integer
	//
	// 星期四是否營運 : [0:'否',1:'是']
	// Required: true
	Thursday *int32 `json:"Thursday"`

	// integer
	//
	// 星期二是否營運 : [0:'否',1:'是']
	// Required: true
	Tuesday *int32 `json:"Tuesday"`

	// integer
	//
	// 星期三是否營運 : [0:'否',1:'是']
	// Required: true
	Wednesday *int32 `json:"Wednesday"`
}

// Validate validates this service d t o version2 bus service day
func (m *ServiceDTOVersion2BusServiceDay) Validate(formats strfmt.Registry) error {
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

func (m *ServiceDTOVersion2BusServiceDay) validateFriday(formats strfmt.Registry) error {

	if err := validate.Required("Friday", "body", m.Friday); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BusServiceDay) validateMonday(formats strfmt.Registry) error {

	if err := validate.Required("Monday", "body", m.Monday); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BusServiceDay) validateSaturday(formats strfmt.Registry) error {

	if err := validate.Required("Saturday", "body", m.Saturday); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BusServiceDay) validateSunday(formats strfmt.Registry) error {

	if err := validate.Required("Sunday", "body", m.Sunday); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BusServiceDay) validateThursday(formats strfmt.Registry) error {

	if err := validate.Required("Thursday", "body", m.Thursday); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BusServiceDay) validateTuesday(formats strfmt.Registry) error {

	if err := validate.Required("Tuesday", "body", m.Tuesday); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BusServiceDay) validateWednesday(formats strfmt.Registry) error {

	if err := validate.Required("Wednesday", "body", m.Wednesday); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2BusServiceDay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2BusServiceDay) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2BusServiceDay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
