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

// PTXServiceDTOBusSpecificationV2ServiceTime ServiceTime
//
// 行駛時間資料型別
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.ServiceTime
type PTXServiceDTOBusSpecificationV2ServiceTime struct {

	// Int32
	//
	// 結束時區(小時)
	// Required: true
	EndHour *int32 `json:"EndHour"`

	// Array
	//
	// 站間行駛時間資訊
	// Required: true
	S2STimes []*PTXServiceDTOBusSpecificationV2TravelTime `json:"S2STimes"`

	// Int32
	//
	// 起始時區(小時)
	// Required: true
	StartHour *int32 `json:"StartHour"`

	// Int32
	//
	// 星期
	// Required: true
	Weekday *int32 `json:"Weekday"`
}

// Validate validates this p t x service d t o bus specification v2 service time
func (m *PTXServiceDTOBusSpecificationV2ServiceTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndHour(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS2STimes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartHour(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeekday(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2ServiceTime) validateEndHour(formats strfmt.Registry) error {

	if err := validate.Required("EndHour", "body", m.EndHour); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2ServiceTime) validateS2STimes(formats strfmt.Registry) error {

	if err := validate.Required("S2STimes", "body", m.S2STimes); err != nil {
		return err
	}

	for i := 0; i < len(m.S2STimes); i++ {
		if swag.IsZero(m.S2STimes[i]) { // not required
			continue
		}

		if m.S2STimes[i] != nil {
			if err := m.S2STimes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("S2STimes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2ServiceTime) validateStartHour(formats strfmt.Registry) error {

	if err := validate.Required("StartHour", "body", m.StartHour); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2ServiceTime) validateWeekday(formats strfmt.Registry) error {

	if err := validate.Required("Weekday", "body", m.Weekday); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v2 service time based on the context it is used
func (m *PTXServiceDTOBusSpecificationV2ServiceTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateS2STimes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2ServiceTime) contextValidateS2STimes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.S2STimes); i++ {

		if m.S2STimes[i] != nil {
			if err := m.S2STimes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("S2STimes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2ServiceTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2ServiceTime) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2ServiceTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
