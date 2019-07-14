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

// ServiceDTOVersion2BikeBikeStation BikeStation
//
// 自行車站點資訊
// swagger:model Service.DTO.Version2.Bike.BikeStation
type ServiceDTOVersion2BikeBikeStation struct {

	// 業管單位代碼
	AuthorityID string `json:"AuthorityID,omitempty"`

	// 可容納之自行車總數
	BikesCapacity int32 `json:"BikesCapacity,omitempty"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// NameType
	//
	// 站點地址
	StationAddress *ServiceDTOVersion2BaseNameType `json:"StationAddress,omitempty"`

	// 站點代碼
	StationID string `json:"StationID,omitempty"`

	// NameType
	//
	// 站點名稱
	StationName *ServiceDTOVersion2BaseNameType `json:"StationName,omitempty"`

	// PointType
	//
	// 站點位置
	StationPosition *ServiceDTOVersion2BasePointType `json:"StationPosition,omitempty"`

	// 站點唯一識別代碼，規則為 {業管機關代碼} + {StationID}，其中 {業管機關代碼} 可於Authority API中的AuthorityCode欄位查詢
	StationUID string `json:"StationUID,omitempty"`

	// 站點描述
	StopDescription string `json:"StopDescription,omitempty"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this service d t o version2 bike bike station
func (m *ServiceDTOVersion2BikeBikeStation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationPosition(formats); err != nil {
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

func (m *ServiceDTOVersion2BikeBikeStation) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BikeBikeStation) validateStationAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.StationAddress) { // not required
		return nil
	}

	if m.StationAddress != nil {
		if err := m.StationAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StationAddress")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2BikeBikeStation) validateStationName(formats strfmt.Registry) error {

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

func (m *ServiceDTOVersion2BikeBikeStation) validateStationPosition(formats strfmt.Registry) error {

	if swag.IsZero(m.StationPosition) { // not required
		return nil
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

func (m *ServiceDTOVersion2BikeBikeStation) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2BikeBikeStation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2BikeBikeStation) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2BikeBikeStation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
