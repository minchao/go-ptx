// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceDTOVersion3BusScopeSubRoute SubRoute
//
// 附屬路線資料
//
// swagger:model Service.DTO.Version3.Bus.Scope.SubRoute
type ServiceDTOVersion3BusScopeSubRoute struct {

	// 地區既用中之附屬班次代碼(為原資料內碼)
	SubRouteID string `json:"SubRouteID,omitempty"`

	// NameType
	//
	// 附屬路線名稱
	SubRouteName *ServiceDTOVersion3BaseNameType `json:"SubRouteName,omitempty"`
}

// Validate validates this service d t o version3 bus scope sub route
func (m *ServiceDTOVersion3BusScopeSubRoute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubRouteName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3BusScopeSubRoute) validateSubRouteName(formats strfmt.Registry) error {

	if swag.IsZero(m.SubRouteName) { // not required
		return nil
	}

	if m.SubRouteName != nil {
		if err := m.SubRouteName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SubRouteName")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3BusScopeSubRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3BusScopeSubRoute) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3BusScopeSubRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
