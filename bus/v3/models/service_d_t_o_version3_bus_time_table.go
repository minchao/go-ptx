// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion3BusTimeTable TimeTable
// swagger:model Service.DTO.Version3.Bus.TimeTable
type ServiceDTOVersion3BusTimeTable struct {

	// 公車停靠時間資料
	// Required: true
	StopTimes []*ServiceDTOVersion3BusStopTime `json:"StopTimes"`

	// 班次代碼，為無意義之編碼
	TripID string `json:"TripID,omitempty"`
}

// Validate validates this service d t o version3 bus time table
func (m *ServiceDTOVersion3BusTimeTable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStopTimes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3BusTimeTable) validateStopTimes(formats strfmt.Registry) error {

	if err := validate.Required("StopTimes", "body", m.StopTimes); err != nil {
		return err
	}

	for i := 0; i < len(m.StopTimes); i++ {
		if swag.IsZero(m.StopTimes[i]) { // not required
			continue
		}

		if m.StopTimes[i] != nil {
			if err := m.StopTimes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StopTimes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3BusTimeTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3BusTimeTable) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3BusTimeTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}