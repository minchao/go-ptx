// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion3RailTRADailyStationTimeTableStationTimetable StationTimetable
// swagger:model Service.DTO.Version3.Rail.TRA.DailyStationTimeTable.StationTimetable
type ServiceDTOVersion3RailTRADailyStationTimeTableStationTimetable struct {

	// 目的站車站代號
	DestinationStationID string `json:"DestinationStationID,omitempty"`

	// NameType
	//
	// 目的站車站名稱
	DestinationStationName *ServiceDTOVersion3BaseNameType `json:"DestinationStationName,omitempty"`

	// 行駛方向
	// Enum: [0: 順行 1: 逆行]
	Direction string `json:"Direction,omitempty"`

	// 營運路線代碼
	RouteID string `json:"RouteID,omitempty"`

	// 車站代碼
	StationID string `json:"StationID,omitempty"`

	// NameType
	//
	// 車站名稱
	StationName *ServiceDTOVersion3BaseNameType `json:"StationName,omitempty"`

	// 目的站車站名稱
	// Required: true
	TimeTables []*ServiceDTOVersion3RailTRADailyStationTimeTableTimeTable `json:"TimeTables"`
}

// Validate validates this service d t o version3 rail t r a daily station time table station timetable
func (m *ServiceDTOVersion3RailTRADailyStationTimeTableStationTimetable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

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

func (m *ServiceDTOVersion3RailTRADailyStationTimeTableStationTimetable) validateDestinationStationName(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationStationName) { // not required
		return nil
	}

	if m.DestinationStationName != nil {
		if err := m.DestinationStationName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DestinationStationName")
			}
			return err
		}
	}

	return nil
}

var serviceDTOVersion3RailTRADailyStationTimeTableStationTimetableTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["0: 順行","1: 逆行"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceDTOVersion3RailTRADailyStationTimeTableStationTimetableTypeDirectionPropEnum = append(serviceDTOVersion3RailTRADailyStationTimeTableStationTimetableTypeDirectionPropEnum, v)
	}
}

const (

	// ServiceDTOVersion3RailTRADailyStationTimeTableStationTimetableDirectionNr0順行 captures enum value "0: 順行"
	ServiceDTOVersion3RailTRADailyStationTimeTableStationTimetableDirectionNr0順行 string = "0: 順行"

	// ServiceDTOVersion3RailTRADailyStationTimeTableStationTimetableDirectionNr1逆行 captures enum value "1: 逆行"
	ServiceDTOVersion3RailTRADailyStationTimeTableStationTimetableDirectionNr1逆行 string = "1: 逆行"
)

// prop value enum
func (m *ServiceDTOVersion3RailTRADailyStationTimeTableStationTimetable) validateDirectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serviceDTOVersion3RailTRADailyStationTimeTableStationTimetableTypeDirectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceDTOVersion3RailTRADailyStationTimeTableStationTimetable) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRADailyStationTimeTableStationTimetable) validateStationName(formats strfmt.Registry) error {

	if swag.IsZero(m.StationName) { // not required
		return nil
	}

	if m.StationName != nil {
		if err := m.StationName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StationName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRADailyStationTimeTableStationTimetable) validateTimeTables(formats strfmt.Registry) error {

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
func (m *ServiceDTOVersion3RailTRADailyStationTimeTableStationTimetable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRADailyStationTimeTableStationTimetable) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3RailTRADailyStationTimeTableStationTimetable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
