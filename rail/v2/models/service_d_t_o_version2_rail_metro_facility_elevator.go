// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion2RailMetroFacilityElevator Elevator
//
// 無障礙電梯位置資訊
// swagger:model Service.DTO.Version2.Rail.Metro.Facility.Elevator
type ServiceDTOVersion2RailMetroFacilityElevator struct {

	// 位置描述
	// Required: true
	Description *string `json:"Description"`

	// 樓層
	// Required: true
	FloorLevel *string `json:"FloorLevel"`
}

// Validate validates this service d t o version2 rail metro facility elevator
func (m *ServiceDTOVersion2RailMetroFacilityElevator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFloorLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion2RailMetroFacilityElevator) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFacilityElevator) validateFloorLevel(formats strfmt.Registry) error {

	if err := validate.Required("FloorLevel", "body", m.FloorLevel); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroFacilityElevator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroFacilityElevator) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailMetroFacilityElevator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
