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

// PTXServiceDTORailSpecificationV3TRAStationFacilityFacilityInfo FacilityInfo
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.StationFacility.FacilityInfo
type PTXServiceDTORailSpecificationV3TRAStationFacilityFacilityInfo struct {

	// String
	//
	// 位置描述
	// Required: true
	Description *string `json:"Description" xml:"Description"`

	// String
	//
	// 樓層
	FloorLevel string `json:"FloorLevel,omitempty" xml:"FloorLevel,omitempty"`
}

// Validate validates this p t x service d t o rail specification v3 t r a station facility facility info
func (m *PTXServiceDTORailSpecificationV3TRAStationFacilityFacilityInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationFacilityFacilityInfo) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o rail specification v3 t r a station facility facility info based on context it is used
func (m *PTXServiceDTORailSpecificationV3TRAStationFacilityFacilityInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAStationFacilityFacilityInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAStationFacilityFacilityInfo) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRAStationFacilityFacilityInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
