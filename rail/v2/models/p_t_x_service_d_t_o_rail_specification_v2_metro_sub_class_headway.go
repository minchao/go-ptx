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

// PTXServiceDTORailSpecificationV2MetroSubClassHeadway Headway
//
// 班距頻率資訊
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.Metro.SubClass.Headway
type PTXServiceDTORailSpecificationV2MetroSubClassHeadway struct {

	// String
	//
	// 結束時間
	// Required: true
	EndTime *string `json:"EndTime" xml:"EndTime"`

	// Int32
	//
	// 最大班距時間(分)
	// Required: true
	MaxHeadwayMins *int32 `json:"MaxHeadwayMins"`

	// Int32
	//
	// 最小班距時間(分)
	// Required: true
	MinHeadwayMins *int32 `json:"MinHeadwayMins"`

	// String
	//
	// 尖峰/離峰狀態(0:離峰, 1:尖峰)
	// Required: true
	PeakFlag *string `json:"PeakFlag" xml:"PeakFlag"`

	// String
	//
	// 開始時間
	// Required: true
	StartTime *string `json:"StartTime" xml:"StartTime"`
}

// Validate validates this p t x service d t o rail specification v2 metro sub class headway
func (m *PTXServiceDTORailSpecificationV2MetroSubClassHeadway) Validate(formats strfmt.Registry) error {
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

func (m *PTXServiceDTORailSpecificationV2MetroSubClassHeadway) validateEndTime(formats strfmt.Registry) error {

	if err := validate.Required("EndTime", "body", m.EndTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassHeadway) validateMaxHeadwayMins(formats strfmt.Registry) error {

	if err := validate.Required("MaxHeadwayMins", "body", m.MaxHeadwayMins); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassHeadway) validateMinHeadwayMins(formats strfmt.Registry) error {

	if err := validate.Required("MinHeadwayMins", "body", m.MinHeadwayMins); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassHeadway) validatePeakFlag(formats strfmt.Registry) error {

	if err := validate.Required("PeakFlag", "body", m.PeakFlag); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassHeadway) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("StartTime", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o rail specification v2 metro sub class headway based on context it is used
func (m *PTXServiceDTORailSpecificationV2MetroSubClassHeadway) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroSubClassHeadway) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroSubClassHeadway) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2MetroSubClassHeadway
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
