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

// ServiceDTOVersion2AviationCodeShare CodeShare
//
// 共用班號
// swagger:model Service.DTO.Version2.Aviation.CodeShare
type ServiceDTOVersion2AviationCodeShare struct {

	// 航空公司IATA國際代碼
	// Required: true
	AirlineID *string `json:"AirlineID"`

	// 班機號碼
	// Required: true
	FlightNumber *string `json:"FlightNumber"`
}

// Validate validates this service d t o version2 aviation code share
func (m *ServiceDTOVersion2AviationCodeShare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAirlineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlightNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion2AviationCodeShare) validateAirlineID(formats strfmt.Registry) error {

	if err := validate.Required("AirlineID", "body", m.AirlineID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2AviationCodeShare) validateFlightNumber(formats strfmt.Registry) error {

	if err := validate.Required("FlightNumber", "body", m.FlightNumber); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2AviationCodeShare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2AviationCodeShare) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2AviationCodeShare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
