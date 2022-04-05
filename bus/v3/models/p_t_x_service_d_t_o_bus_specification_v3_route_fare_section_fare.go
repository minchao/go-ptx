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

// PTXServiceDTOBusSpecificationV3RouteFareSectionFare SectionFare
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.RouteFare+SectionFare
type PTXServiceDTOBusSpecificationV3RouteFareSectionFare struct {

	// Array
	//
	// 緩衝區資訊
	// Required: true
	BufferZones []*PTXServiceDTOBusSpecificationV3RouteFareSectionFareBufferZone "json:\"BufferZones\" xml:\"List`1\""

	// Array
	//
	// 每段收費資訊
	// Required: true
	Fares []*PTXServiceDTOBusSpecificationV3RouteFareSectionFareFareSection "json:\"Fares\" xml:\"List`1\""
}

// Validate validates this p t x service d t o bus specification v3 route fare section fare
func (m *PTXServiceDTOBusSpecificationV3RouteFareSectionFare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBufferZones(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFares(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteFareSectionFare) validateBufferZones(formats strfmt.Registry) error {

	if err := validate.Required("BufferZones", "body", m.BufferZones); err != nil {
		return err
	}

	for i := 0; i < len(m.BufferZones); i++ {
		if swag.IsZero(m.BufferZones[i]) { // not required
			continue
		}

		if m.BufferZones[i] != nil {
			if err := m.BufferZones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BufferZones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BufferZones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteFareSectionFare) validateFares(formats strfmt.Registry) error {

	if err := validate.Required("Fares", "body", m.Fares); err != nil {
		return err
	}

	for i := 0; i < len(m.Fares); i++ {
		if swag.IsZero(m.Fares[i]) { // not required
			continue
		}

		if m.Fares[i] != nil {
			if err := m.Fares[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Fares" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Fares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v3 route fare section fare based on the context it is used
func (m *PTXServiceDTOBusSpecificationV3RouteFareSectionFare) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBufferZones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFares(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteFareSectionFare) contextValidateBufferZones(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BufferZones); i++ {

		if m.BufferZones[i] != nil {
			if err := m.BufferZones[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BufferZones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BufferZones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteFareSectionFare) contextValidateFares(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Fares); i++ {

		if m.Fares[i] != nil {
			if err := m.Fares[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Fares" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Fares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3RouteFareSectionFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3RouteFareSectionFare) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3RouteFareSectionFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
