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

// ServiceDTOVersion2RailTRAStationOfLine StationOfLine
//
// 路線車站基本資料
// swagger:model Service.DTO.Version2.Rail.TRA.StationOfLine
type ServiceDTOVersion2RailTRAStationOfLine struct {

	// 路線代碼
	// Required: true
	LineID *string `json:"LineID"`

	// 路線編號
	// Required: true
	LineNo *string `json:"LineNo"`

	// 路線車站資訊
	// Required: true
	Stations []*ServiceDTOVersion2RailTRALineStation `json:"Stations"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this service d t o version2 rail t r a station of line
func (m *ServiceDTOVersion2RailTRAStationOfLine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineNo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion2RailTRAStationOfLine) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRAStationOfLine) validateLineNo(formats strfmt.Registry) error {

	if err := validate.Required("LineNo", "body", m.LineNo); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRAStationOfLine) validateStations(formats strfmt.Registry) error {

	if err := validate.Required("Stations", "body", m.Stations); err != nil {
		return err
	}

	for i := 0; i < len(m.Stations); i++ {
		if swag.IsZero(m.Stations[i]) { // not required
			continue
		}

		if m.Stations[i] != nil {
			if err := m.Stations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Stations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceDTOVersion2RailTRAStationOfLine) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTRAStationOfLine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTRAStationOfLine) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailTRAStationOfLine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
