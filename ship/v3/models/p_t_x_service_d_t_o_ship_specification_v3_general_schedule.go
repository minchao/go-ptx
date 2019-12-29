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

// PTXServiceDTOShipSpecificationV3GeneralSchedule GeneralSchedule
// swagger:model PTX.Service.DTO.Ship.Specification.V3.GeneralSchedule
type PTXServiceDTOShipSpecificationV3GeneralSchedule struct {

	// Int32
	//
	// 方向性描述(0=去程, 1=返程,2=迴圈)
	Direction int32 `json:"Direction,omitempty"`

	// String
	//
	// 營運業者代碼
	OperatorID string `json:"OperatorID,omitempty"`

	// String
	//
	// 航線代碼
	RouteID string `json:"RouteID,omitempty"`

	// NameType
	//
	// 航線名稱
	RouteName *PTXServiceDTOSharedSpecificationV3BaseNameType `json:"RouteName,omitempty"`

	// Array
	Timetables []*PTXServiceDTOShipSpecificationV3Timetable `json:"Timetables"`

	// Array
	//
	// 船舶資訊
	Vessels []*PTXServiceDTOShipSpecificationV3Vessel `json:"Vessels"`
}

// Validate validates this p t x service d t o ship specification v3 general schedule
func (m *PTXServiceDTOShipSpecificationV3GeneralSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRouteName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimetables(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVessels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOShipSpecificationV3GeneralSchedule) validateRouteName(formats strfmt.Registry) error {

	if swag.IsZero(m.RouteName) { // not required
		return nil
	}

	if m.RouteName != nil {
		if err := m.RouteName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RouteName")
			}
			return err
		}
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3GeneralSchedule) validateTimetables(formats strfmt.Registry) error {

	if swag.IsZero(m.Timetables) { // not required
		return nil
	}

	for i := 0; i < len(m.Timetables); i++ {
		if swag.IsZero(m.Timetables[i]) { // not required
			continue
		}

		if m.Timetables[i] != nil {
			if err := m.Timetables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Timetables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3GeneralSchedule) validateVessels(formats strfmt.Registry) error {

	if swag.IsZero(m.Vessels) { // not required
		return nil
	}

	for i := 0; i < len(m.Vessels); i++ {
		if swag.IsZero(m.Vessels[i]) { // not required
			continue
		}

		if m.Vessels[i] != nil {
			if err := m.Vessels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Vessels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3GeneralSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3GeneralSchedule) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOShipSpecificationV3GeneralSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
