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

// ServiceDTOVersion2RailMetroSubClassHeadway Headway
//
// 班距頻率資訊
// swagger:model Service.DTO.Version2.Rail.Metro.SubClass.Headway
type ServiceDTOVersion2RailMetroSubClassHeadway struct {

	// 結束時間
	// Required: true
	EndTime *string `json:"EndTime"`

	// 最大班距時間(分)
	// Required: true
	MaxHeadwayMins *int32 `json:"MaxHeadwayMins"`

	// 最小班距時間(分)
	// Required: true
	MinHeadwayMins *int32 `json:"MinHeadwayMins"`

	// 尖峰/離峰狀態(0:離峰, 1:尖峰)
	// Required: true
	PeakFlag *string `json:"PeakFlag"`

	// 開始時間
	// Required: true
	StartTime *string `json:"StartTime"`
}

// Validate validates this service d t o version2 rail metro sub class headway
func (m *ServiceDTOVersion2RailMetroSubClassHeadway) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxHeadwayMins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinHeadwayMins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeakFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion2RailMetroSubClassHeadway) validateEndTime(formats strfmt.Registry) error {

	if err := validate.Required("EndTime", "body", m.EndTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroSubClassHeadway) validateMaxHeadwayMins(formats strfmt.Registry) error {

	if err := validate.Required("MaxHeadwayMins", "body", m.MaxHeadwayMins); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroSubClassHeadway) validateMinHeadwayMins(formats strfmt.Registry) error {

	if err := validate.Required("MinHeadwayMins", "body", m.MinHeadwayMins); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroSubClassHeadway) validatePeakFlag(formats strfmt.Registry) error {

	if err := validate.Required("PeakFlag", "body", m.PeakFlag); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroSubClassHeadway) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("StartTime", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroSubClassHeadway) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroSubClassHeadway) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailMetroSubClassHeadway
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
