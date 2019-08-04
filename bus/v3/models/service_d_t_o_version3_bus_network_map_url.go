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

// ServiceDTOVersion3BusNetworkMapURL NetworkMapURL
// swagger:model Service.DTO.Version3.Bus.NetworkMapURL
type ServiceDTOVersion3BusNetworkMapURL struct {

	// 路網圖名稱
	// Required: true
	MapName *string `json:"MapName"`

	// NameType
	//
	// 路網圖網址URL
	// Required: true
	MapNameURL *ServiceDTOVersion3BaseNameType `json:"MapNameURL"`
}

// Validate validates this service d t o version3 bus network map URL
func (m *ServiceDTOVersion3BusNetworkMapURL) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMapName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapNameURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3BusNetworkMapURL) validateMapName(formats strfmt.Registry) error {

	if err := validate.Required("MapName", "body", m.MapName); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusNetworkMapURL) validateMapNameURL(formats strfmt.Registry) error {

	if err := validate.Required("MapNameURL", "body", m.MapNameURL); err != nil {
		return err
	}

	if m.MapNameURL != nil {
		if err := m.MapNameURL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MapNameURL")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3BusNetworkMapURL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3BusNetworkMapURL) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3BusNetworkMapURL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}