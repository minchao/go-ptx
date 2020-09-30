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

// PTXServiceDTORailSpecificationV2MetroFrequency Frequency
//
// 路線發車班距頻率資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.Metro.Frequency
type PTXServiceDTORailSpecificationV2MetroFrequency struct {

	// Array
	//
	// 班距頻率資訊
	// Required: true
	Headways []*PTXServiceDTORailSpecificationV2MetroSubClassHeadway `json:"Headways"`

	// String
	//
	// 營運路線所屬之路線代碼
	// Required: true
	LineID *string `json:"LineID"`

	// String
	//
	// 營運路線所屬之路線編號
	LineNo string `json:"LineNo,omitempty"`

	// OperationTime
	//
	// 營運時間資訊
	// Required: true
	OperationTime struct {
		PTXServiceDTORailSpecificationV2MetroSubClassOperationTime
	} `json:"OperationTime"`

	// String
	//
	// 營運路線代碼
	// Required: true
	RouteID *string `json:"RouteID"`

	// ServiceDays
	//
	// 服務日型態
	// Required: true
	ServiceDays struct {
		PTXServiceDTORailSpecificationV2MetroSubClassServiceDays
	} `json:"ServiceDays"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// 車種(0:不分車種, 1:普通車, 2:直達車)
	TrainType int32 `json:"TrainType,omitempty"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// Int32
	//
	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this p t x service d t o rail specification v2 metro frequency
func (m *PTXServiceDTORailSpecificationV2MetroFrequency) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeadways(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceDays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFrequency) validateHeadways(formats strfmt.Registry) error {

	if err := validate.Required("Headways", "body", m.Headways); err != nil {
		return err
	}

	for i := 0; i < len(m.Headways); i++ {
		if swag.IsZero(m.Headways[i]) { // not required
			continue
		}

		if m.Headways[i] != nil {
			if err := m.Headways[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Headways" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFrequency) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFrequency) validateOperationTime(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFrequency) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFrequency) validateServiceDays(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFrequency) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFrequency) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFrequency) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroFrequency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroFrequency) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2MetroFrequency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}