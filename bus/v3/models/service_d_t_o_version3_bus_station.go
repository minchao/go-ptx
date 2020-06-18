// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion3BusStation Station
//
// swagger:model Service.DTO.Version3.Bus.Station
type ServiceDTOVersion3BusStation struct {

	// 方位角，E:東行;W:西行;S:南行;N:北行;SE:東南行;NE:東北行;SW:西南行;NW:西北行
	Bearing string `json:"Bearing,omitempty"`

	// 站牌所在道路上之路名。
	RoadName string `json:"RoadName,omitempty"`

	// 站位地址
	StationAddress string `json:"StationAddress,omitempty"`

	// 站牌詳細說明描述
	StationDescription string `json:"StationDescription,omitempty"`

	// 站位代碼
	// Required: true
	StationID *string `json:"StationID"`

	// NameType
	//
	// 站位名稱
	// Required: true
	StationName *ServiceDTOVersion3BaseNameType `json:"StationName"`

	// PointType
	//
	// 站位位置
	// Required: true
	StationPosition *ServiceDTOVersion3BasePointType `json:"StationPosition"`

	// 站位唯一識別代碼，規則為 {業管機關簡碼} + {StationID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	StationUID *string `json:"StationUID"`
}

// Validate validates this service d t o version3 bus station
func (m *ServiceDTOVersion3BusStation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3BusStation) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusStation) validateStationName(formats strfmt.Registry) error {

	if err := validate.Required("StationName", "body", m.StationName); err != nil {
		return err
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

func (m *ServiceDTOVersion3BusStation) validateStationPosition(formats strfmt.Registry) error {

	if err := validate.Required("StationPosition", "body", m.StationPosition); err != nil {
		return err
	}

	if m.StationPosition != nil {
		if err := m.StationPosition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StationPosition")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion3BusStation) validateStationUID(formats strfmt.Registry) error {

	if err := validate.Required("StationUID", "body", m.StationUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3BusStation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3BusStation) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3BusStation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
