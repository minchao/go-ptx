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

// ServiceDTOVersion2RailMetroFacilityInformationSpot InformationSpot
//
// 詢問處位置資訊
// swagger:model Service.DTO.Version2.Rail.Metro.Facility.InformationSpot
type ServiceDTOVersion2RailMetroFacilityInformationSpot struct {

	// 位置描述
	// Required: true
	Description *string `json:"Description"`

	// 樓層
	// Required: true
	FloorLevel *string `json:"FloorLevel"`
}

// Validate validates this service d t o version2 rail metro facility information spot
func (m *ServiceDTOVersion2RailMetroFacilityInformationSpot) Validate(formats strfmt.Registry) error {
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

func (m *ServiceDTOVersion2RailMetroFacilityInformationSpot) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFacilityInformationSpot) validateFloorLevel(formats strfmt.Registry) error {

	if err := validate.Required("FloorLevel", "body", m.FloorLevel); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroFacilityInformationSpot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroFacilityInformationSpot) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailMetroFacilityInformationSpot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
