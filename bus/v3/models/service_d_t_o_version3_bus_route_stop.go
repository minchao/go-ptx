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

// ServiceDTOVersion3BusRouteStop Stop
// swagger:model Service.DTO.Version3.Bus.Route.Stop
type ServiceDTOVersion3BusRouteStop struct {

	// 站牌代碼
	// Required: true
	StopID *string `json:"StopID"`

	// NameType
	//
	// 站牌名稱
	// Required: true
	StopName *ServiceDTOVersion3BaseNameType `json:"StopName"`
}

// Validate validates this service d t o version3 bus route stop
func (m *ServiceDTOVersion3BusRouteStop) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStopID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3BusRouteStop) validateStopID(formats strfmt.Registry) error {

	if err := validate.Required("StopID", "body", m.StopID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusRouteStop) validateStopName(formats strfmt.Registry) error {

	if err := validate.Required("StopName", "body", m.StopName); err != nil {
		return err
	}

	if m.StopName != nil {
		if err := m.StopName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StopName")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3BusRouteStop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3BusRouteStop) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3BusRouteStop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
