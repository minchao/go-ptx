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

// PTXServiceDTOAirSpecificationV2CodeShare CodeShare
//
// 共用班號
//
// swagger:model PTX.Service.DTO.Air.Specification.V2.CodeShare
type PTXServiceDTOAirSpecificationV2CodeShare struct {

	// String
	//
	// 航空公司IATA國際代碼
	// Required: true
	AirlineID *string `json:"AirlineID" xml:"String"`

	// String
	//
	// 航機班號(不包含航空公司的AirlineID，僅有班號數字)
	// Required: true
	FlightNumber *string `json:"FlightNumber" xml:"String"`
}

// Validate validates this p t x service d t o air specification v2 code share
func (m *PTXServiceDTOAirSpecificationV2CodeShare) Validate(formats strfmt.Registry) error {
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

func (m *PTXServiceDTOAirSpecificationV2CodeShare) validateAirlineID(formats strfmt.Registry) error {

	if err := validate.Required("AirlineID", "body", m.AirlineID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2CodeShare) validateFlightNumber(formats strfmt.Registry) error {

	if err := validate.Required("FlightNumber", "body", m.FlightNumber); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o air specification v2 code share based on context it is used
func (m *PTXServiceDTOAirSpecificationV2CodeShare) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOAirSpecificationV2CodeShare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOAirSpecificationV2CodeShare) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOAirSpecificationV2CodeShare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
