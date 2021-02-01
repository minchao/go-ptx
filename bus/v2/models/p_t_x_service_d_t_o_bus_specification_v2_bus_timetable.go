// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTOBusSpecificationV2BusTimetable BusTimetable
//
// 班次資料
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.BusTimetable
type PTXServiceDTOBusSpecificationV2BusTimetable struct {

	// ServiceDay
	//
	// 週內營運日
	ServiceDay struct {
		PTXServiceDTOBusSpecificationV2EmbeddedServiceDay
	} `json:"ServiceDay,omitempty"`

	// Array
	//
	// 特殊營運日
	SpecialDays []*PTXServiceDTOBusSpecificationV2SpecialDay `json:"SpecialDays"`

	// Array
	//
	// 公車停靠時間資料
	// Required: true
	StopTimes []*PTXServiceDTOBusSpecificationV2BusStopTime `json:"StopTimes"`

	// String
	//
	// 班次代碼，為無意義之編碼
	TripID string `json:"TripID,omitempty"`
}

// Validate validates this p t x service d t o bus specification v2 bus timetable
func (m *PTXServiceDTOBusSpecificationV2BusTimetable) Validate(formats strfmt.Registry) error {
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

func (m *PTXServiceDTOBusSpecificationV2BusTimetable) validateServiceDay(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceDay) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusTimetable) validateSpecialDays(formats strfmt.Registry) error {
	if swag.IsZero(m.SpecialDays) { // not required
		return nil
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

func (m *PTXServiceDTOBusSpecificationV2BusTimetable) validateStopTimes(formats strfmt.Registry) error {

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

// ContextValidate validate this p t x service d t o bus specification v2 bus timetable based on the context it is used
func (m *PTXServiceDTOBusSpecificationV2BusTimetable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceDay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpecialDays(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStopTimes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusTimetable) contextValidateServiceDay(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusTimetable) contextValidateSpecialDays(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SpecialDays); i++ {

		if m.SpecialDays[i] != nil {
			if err := m.SpecialDays[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SpecialDays" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusTimetable) contextValidateStopTimes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StopTimes); i++ {

		if m.StopTimes[i] != nil {
			if err := m.StopTimes[i].ContextValidate(ctx, formats); err != nil {
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
func (m *PTXServiceDTOBusSpecificationV2BusTimetable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusTimetable) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2BusTimetable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
