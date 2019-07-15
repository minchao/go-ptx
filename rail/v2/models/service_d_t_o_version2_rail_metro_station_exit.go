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

// ServiceDTOVersion2RailMetroStationExit StationExit
//
// 捷運車站出入口基本資料
// swagger:model Service.DTO.Version2.Rail.Metro.StationExit
type ServiceDTOVersion2RailMetroStationExit struct {

	// 是否有電梯
	// Required: true
	Elevator *bool `json:"Elevator"`

	// 是否有電扶梯(0:沒有,1:雙向,2:出站,3:入站)
	// Required: true
	Escalator *int32 `json:"Escalator"`

	// 出入口代碼
	// Required: true
	ExitID *string `json:"ExitID"`

	// NameType
	//
	// 出入口名稱
	// Required: true
	ExitName *ServiceDTOVersion2BaseNameType `json:"ExitName"`

	// PointType
	//
	// 出入口座標
	// Required: true
	ExitPosition *ServiceDTOVersion2BasePointType `json:"ExitPosition"`

	// 地址描述
	// Required: true
	LocationDescription *string `json:"LocationDescription"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// 是否有樓梯
	// Required: true
	Stair *bool `json:"Stair"`

	// 車站代碼
	// Required: true
	StationID *string `json:"StationID"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName *ServiceDTOVersion2BaseNameType `json:"StationName"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this service d t o version2 rail metro station exit
func (m *ServiceDTOVersion2RailMetroStationExit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElevator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEscalator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExitID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExitName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExitPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStair(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
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

func (m *ServiceDTOVersion2RailMetroStationExit) validateElevator(formats strfmt.Registry) error {

	if err := validate.Required("Elevator", "body", m.Elevator); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroStationExit) validateEscalator(formats strfmt.Registry) error {

	if err := validate.Required("Escalator", "body", m.Escalator); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroStationExit) validateExitID(formats strfmt.Registry) error {

	if err := validate.Required("ExitID", "body", m.ExitID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroStationExit) validateExitName(formats strfmt.Registry) error {

	if err := validate.Required("ExitName", "body", m.ExitName); err != nil {
		return err
	}

	if m.ExitName != nil {
		if err := m.ExitName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ExitName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroStationExit) validateExitPosition(formats strfmt.Registry) error {

	if err := validate.Required("ExitPosition", "body", m.ExitPosition); err != nil {
		return err
	}

	if m.ExitPosition != nil {
		if err := m.ExitPosition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ExitPosition")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroStationExit) validateLocationDescription(formats strfmt.Registry) error {

	if err := validate.Required("LocationDescription", "body", m.LocationDescription); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroStationExit) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroStationExit) validateStair(formats strfmt.Registry) error {

	if err := validate.Required("Stair", "body", m.Stair); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroStationExit) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroStationExit) validateStationName(formats strfmt.Registry) error {

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

func (m *ServiceDTOVersion2RailMetroStationExit) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroStationExit) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroStationExit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroStationExit) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailMetroStationExit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
