// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable StationTimetable
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.DailyStationTimeTable.StationTimetable
type PTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable struct {

	// integer
	//
	// 行駛方向 : [0:'順行',1:'逆行']
	Direction int32 `json:"Direction,omitempty"`

	// String
	//
	// 營運路線代碼
	RouteID string `json:"RouteID,omitempty"`

	// String
	//
	// 車站代碼
	StationID string `json:"StationID,omitempty"`

	// NameType
	//
	// 車站名稱
	StationName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"StationName,omitempty"`

	// Array
	//
	// 目的站車站名稱
	// Required: true
	TimeTables []*PTXServiceDTORailSpecificationV3TRADailyStationTimeTableTimeTable `json:"TimeTables"`
}

// Validate validates this p t x service d t o rail specification v3 t r a daily station time table station timetable
func (m *PTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeTables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable) validateStationName(formats strfmt.Registry) error {

	if swag.IsZero(m.StationName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable) validateTimeTables(formats strfmt.Registry) error {

	if err := validate.Required("TimeTables", "body", m.TimeTables); err != nil {
		return err
	}

	for i := 0; i < len(m.TimeTables); i++ {
		if swag.IsZero(m.TimeTables[i]) { // not required
			continue
		}

		if m.TimeTables[i] != nil {
			if err := m.TimeTables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TimeTables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}