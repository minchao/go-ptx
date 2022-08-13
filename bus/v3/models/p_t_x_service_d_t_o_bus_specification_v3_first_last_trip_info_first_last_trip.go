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

// PTXServiceDTOBusSpecificationV3FirstLastTripInfoFirstLastTrip FirstLastTrip
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.FirstLastTripInfo+FirstLastTrip
type PTXServiceDTOBusSpecificationV3FirstLastTripInfoFirstLastTrip struct {

	// String
	//
	// 頭班車發車時間(HH:mm)
	// Required: true
	FirstTripDepTime *string `json:"FirstTripDepTime" xml:"FirstTripDepTime"`

	// String
	//
	// 末班車發車時間(HH:mm)
	// Required: true
	LastTripDepTime *string `json:"LastTripDepTime" xml:"LastTripDepTime"`

	// ServiceDay
	//
	// 服務日型態
	// Required: true
	ServiceDay struct {
		PTXServiceDTOBusSpecificationV3FirstLastTripInfoServiceDay
	} `json:"ServiceDay" xml:"ServiceDay"`
}

// Validate validates this p t x service d t o bus specification v3 first last trip info first last trip
func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoFirstLastTrip) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirstTripDepTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastTripDepTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceDay(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoFirstLastTrip) validateFirstTripDepTime(formats strfmt.Registry) error {

	if err := validate.Required("FirstTripDepTime", "body", m.FirstTripDepTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoFirstLastTrip) validateLastTripDepTime(formats strfmt.Registry) error {

	if err := validate.Required("LastTripDepTime", "body", m.LastTripDepTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoFirstLastTrip) validateServiceDay(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v3 first last trip info first last trip based on the context it is used
func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoFirstLastTrip) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceDay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoFirstLastTrip) contextValidateServiceDay(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoFirstLastTrip) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3FirstLastTripInfoFirstLastTrip) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3FirstLastTripInfoFirstLastTrip
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
