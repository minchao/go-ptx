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

// PTXServiceDTOBusSpecificationV3RouteFareStageFareServiceDay ServiceDay
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.RouteFare+StageFare+ServiceDay
type PTXServiceDTOBusSpecificationV3RouteFareStageFareServiceDay struct {

	// integer
	//
	// [0:'否',1:'是']
	DayAfterHoliday int32 `json:"DayAfterHoliday,omitempty"`

	// integer
	//
	// [0:'否',1:'是']
	DayBeforeHoliday int32 `json:"DayBeforeHoliday,omitempty"`

	// integer
	//
	// [0:'否',1:'是']
	// Required: true
	Friday *int32 `json:"Friday"`

	// integer
	//
	// [0:'否',1:'是']
	// Required: true
	Monday *int32 `json:"Monday"`

	// integer
	//
	// [0:'否',1:'是']
	NationalHolidays int32 `json:"NationalHolidays,omitempty"`

	// integer
	//
	// [0:'否',1:'是']
	// Required: true
	Saturday *int32 `json:"Saturday"`

	// String
	// Required: true
	ServiceTag *string `json:"ServiceTag"`

	// integer
	//
	// [0:'否',1:'是']
	// Required: true
	Sunday *int32 `json:"Sunday"`

	// integer
	//
	// [0:'否',1:'是']
	// Required: true
	Thursday *int32 `json:"Thursday"`

	// integer
	//
	// [0:'否',1:'是']
	// Required: true
	Tuesday *int32 `json:"Tuesday"`

	// integer
	//
	// [0:'否',1:'是']
	TyphoonDay int32 `json:"TyphoonDay,omitempty"`

	// integer
	//
	// [0:'否',1:'是']
	// Required: true
	Wednesday *int32 `json:"Wednesday"`
}

// Validate validates this p t x service d t o bus specification v3 route fare stage fare service day
func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareServiceDay) Validate(formats strfmt.Registry) error {
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

	if err := m.validateServiceTag(formats); err != nil {
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

func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareServiceDay) validateFriday(formats strfmt.Registry) error {

	if err := validate.Required("Friday", "body", m.Friday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareServiceDay) validateMonday(formats strfmt.Registry) error {

	if err := validate.Required("Monday", "body", m.Monday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareServiceDay) validateSaturday(formats strfmt.Registry) error {

	if err := validate.Required("Saturday", "body", m.Saturday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareServiceDay) validateServiceTag(formats strfmt.Registry) error {

	if err := validate.Required("ServiceTag", "body", m.ServiceTag); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareServiceDay) validateSunday(formats strfmt.Registry) error {

	if err := validate.Required("Sunday", "body", m.Sunday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareServiceDay) validateThursday(formats strfmt.Registry) error {

	if err := validate.Required("Thursday", "body", m.Thursday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareServiceDay) validateTuesday(formats strfmt.Registry) error {

	if err := validate.Required("Tuesday", "body", m.Tuesday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareServiceDay) validateWednesday(formats strfmt.Registry) error {

	if err := validate.Required("Wednesday", "body", m.Wednesday); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o bus specification v3 route fare stage fare service day based on context it is used
func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareServiceDay) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareServiceDay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareServiceDay) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3RouteFareStageFareServiceDay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
