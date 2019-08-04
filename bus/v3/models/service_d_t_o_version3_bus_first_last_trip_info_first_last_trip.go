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

// ServiceDTOVersion3BusFirstLastTripInfoFirstLastTrip FirstLastTrip
// swagger:model Service.DTO.Version3.Bus.FirstLastTripInfo.FirstLastTrip
type ServiceDTOVersion3BusFirstLastTripInfoFirstLastTrip struct {

	// 頭班車發車時間(HH:mm)
	// Required: true
	FirstTripDepTime *string `json:"FirstTripDepTime"`

	// 末班車發車時間(HH:mm)
	// Required: true
	LastTripDepTime *string `json:"LastTripDepTime"`

	// ServiceDay
	//
	// 服務日型態
	// Required: true
	ServiceDay *ServiceDTOVersion3BusFirstLastTripInfoServiceDay `json:"ServiceDay"`
}

// Validate validates this service d t o version3 bus first last trip info first last trip
func (m *ServiceDTOVersion3BusFirstLastTripInfoFirstLastTrip) Validate(formats strfmt.Registry) error {
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

func (m *ServiceDTOVersion3BusFirstLastTripInfoFirstLastTrip) validateFirstTripDepTime(formats strfmt.Registry) error {

	if err := validate.Required("FirstTripDepTime", "body", m.FirstTripDepTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusFirstLastTripInfoFirstLastTrip) validateLastTripDepTime(formats strfmt.Registry) error {

	if err := validate.Required("LastTripDepTime", "body", m.LastTripDepTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusFirstLastTripInfoFirstLastTrip) validateServiceDay(formats strfmt.Registry) error {

	if err := validate.Required("ServiceDay", "body", m.ServiceDay); err != nil {
		return err
	}

	if m.ServiceDay != nil {
		if err := m.ServiceDay.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ServiceDay")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3BusFirstLastTripInfoFirstLastTrip) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3BusFirstLastTripInfoFirstLastTrip) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3BusFirstLastTripInfoFirstLastTrip
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}