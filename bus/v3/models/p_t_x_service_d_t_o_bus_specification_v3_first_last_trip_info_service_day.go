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

// PTXServiceDTOBusSpecificationV3FirstLastTripInfoServiceDay ServiceDay
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.FirstLastTripInfo+ServiceDay
type PTXServiceDTOBusSpecificationV3FirstLastTripInfoServiceDay struct {

	// 國定假日後一日營運與否
	DayAfterNationalHoliday bool `json:"DayAfterNationalHoliday,omitempty"`

	// 國定假日前一日營運與否
	DayBeforeNationalHoliday bool `json:"DayBeforeNationalHoliday,omitempty"`

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
	ServiceTag string `json:"ServiceTag,omitempty"`

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

// Validate validates this p t x service d t o bus specification v3 first last trip info service day
func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoServiceDay) Validate(formats strfmt.Registry) error {
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

func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoServiceDay) validateFriday(formats strfmt.Registry) error {

	if err := validate.Required("Friday", "body", m.Friday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoServiceDay) validateMonday(formats strfmt.Registry) error {

	if err := validate.Required("Monday", "body", m.Monday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoServiceDay) validateSaturday(formats strfmt.Registry) error {

	if err := validate.Required("Saturday", "body", m.Saturday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoServiceDay) validateSunday(formats strfmt.Registry) error {

	if err := validate.Required("Sunday", "body", m.Sunday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoServiceDay) validateThursday(formats strfmt.Registry) error {

	if err := validate.Required("Thursday", "body", m.Thursday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoServiceDay) validateTuesday(formats strfmt.Registry) error {

	if err := validate.Required("Tuesday", "body", m.Tuesday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoServiceDay) validateWednesday(formats strfmt.Registry) error {

	if err := validate.Required("Wednesday", "body", m.Wednesday); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoServiceDay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoServiceDay) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3FirstLastTripInfoServiceDay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}