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

// ServiceDTOVersion3RailTRANetworkLine Line
//
// 路線基本資料
//
// swagger:model Service.DTO.Version3.Rail.TRA.Network.Line
type ServiceDTOVersion3RailTRANetworkLine struct {

	// 路線代碼
	// Required: true
	LineID *string `json:"LineID"`

	// 路線編號
	// Required: true
	LineNo *string `json:"LineNo"`
}

// Validate validates this service d t o version3 rail t r a network line
func (m *ServiceDTOVersion3RailTRANetworkLine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineNo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3RailTRANetworkLine) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRANetworkLine) validateLineNo(formats strfmt.Registry) error {

	if err := validate.Required("LineNo", "body", m.LineNo); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRANetworkLine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRANetworkLine) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3RailTRANetworkLine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
