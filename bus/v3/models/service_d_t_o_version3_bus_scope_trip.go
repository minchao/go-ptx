// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ServiceDTOVersion3BusScopeTrip Trip
//
// 班次代碼資料
// swagger:model Service.DTO.Version3.Bus.Scope.Trip
type ServiceDTOVersion3BusScopeTrip struct {

	// 地區既用中之班次代碼(為原資料內碼)
	TripID string `json:"TripID,omitempty"`
}

// Validate validates this service d t o version3 bus scope trip
func (m *ServiceDTOVersion3BusScopeTrip) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3BusScopeTrip) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3BusScopeTrip) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3BusScopeTrip
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
