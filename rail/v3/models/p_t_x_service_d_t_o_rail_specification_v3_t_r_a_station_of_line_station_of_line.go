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

// PTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine StationOfLine
//
// 路線車站基本資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.StationOfLine.StationOfLine
type PTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine struct {

	// String
	//
	// 路線代碼
	// Required: true
	LineID *string `json:"LineID" xml:"String"`

	// String
	//
	// 路線編號
	// Required: true
	LineNo *string `json:"LineNo" xml:"String"`

	// Array
	//
	// 路線車站資訊
	// Required: true
	Stations []*PTXServiceDTORailSpecificationV3TRAStationOfLineLineStation "json:\"Stations\" xml:\"List`1\""
}

// Validate validates this p t x service d t o rail specification v3 t r a station of line station of line
func (m *PTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineNo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) validateLineNo(formats strfmt.Registry) error {

	if err := validate.Required("LineNo", "body", m.LineNo); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) validateStations(formats strfmt.Registry) error {

	if err := validate.Required("Stations", "body", m.Stations); err != nil {
		return err
	}

	for i := 0; i < len(m.Stations); i++ {
		if swag.IsZero(m.Stations[i]) { // not required
			continue
		}

		if m.Stations[i] != nil {
			if err := m.Stations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Stations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a station of line station of line based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) contextValidateStations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Stations); i++ {

		if m.Stations[i] != nil {
			if err := m.Stations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Stations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
