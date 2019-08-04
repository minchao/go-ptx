// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServiceDTOVersion3BusScope Scope
//
// 通阻影響範圍
// swagger:model Service.DTO.Version3.Bus.Scope
type ServiceDTOVersion3BusScope struct {

	// 路線資料
	Routes []*ServiceDTOVersion3BusScopeRoute `json:"Routes"`

	// 站牌資料
	Stops []*ServiceDTOVersion3BusScopeStop `json:"Stops"`

	// 附屬路線資料
	SubRoutes []*ServiceDTOVersion3BusScopeSubRoute `json:"SubRoutes"`

	// 班次代碼資料
	TripIDs []*ServiceDTOVersion3BusScopeTrip `json:"TripIDs"`
}

// Validate validates this service d t o version3 bus scope
func (m *ServiceDTOVersion3BusScope) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubRoutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTripIDs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3BusScope) validateRoutes(formats strfmt.Registry) error {

	if swag.IsZero(m.Routes) { // not required
		return nil
	}

	for i := 0; i < len(m.Routes); i++ {
		if swag.IsZero(m.Routes[i]) { // not required
			continue
		}

		if m.Routes[i] != nil {
			if err := m.Routes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Routes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceDTOVersion3BusScope) validateStops(formats strfmt.Registry) error {

	if swag.IsZero(m.Stops) { // not required
		return nil
	}

	for i := 0; i < len(m.Stops); i++ {
		if swag.IsZero(m.Stops[i]) { // not required
			continue
		}

		if m.Stops[i] != nil {
			if err := m.Stops[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Stops" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceDTOVersion3BusScope) validateSubRoutes(formats strfmt.Registry) error {

	if swag.IsZero(m.SubRoutes) { // not required
		return nil
	}

	for i := 0; i < len(m.SubRoutes); i++ {
		if swag.IsZero(m.SubRoutes[i]) { // not required
			continue
		}

		if m.SubRoutes[i] != nil {
			if err := m.SubRoutes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SubRoutes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceDTOVersion3BusScope) validateTripIDs(formats strfmt.Registry) error {

	if swag.IsZero(m.TripIDs) { // not required
		return nil
	}

	for i := 0; i < len(m.TripIDs); i++ {
		if swag.IsZero(m.TripIDs[i]) { // not required
			continue
		}

		if m.TripIDs[i] != nil {
			if err := m.TripIDs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TripIDs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3BusScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3BusScope) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3BusScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
