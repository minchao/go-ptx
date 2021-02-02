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

// PTXServiceDTORailSpecificationV2TRARailStation RailStation
//
// 台鐵車站資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.TRA.RailStation
type PTXServiceDTORailSpecificationV2TRARailStation struct {

	// String
	//
	// 車站位置所屬縣市
	LocationCity string `json:"LocationCity,omitempty"`

	// String
	//
	// 車站位置所屬縣市代碼
	LocationCityCode string `json:"LocationCityCode,omitempty"`

	// String
	//
	// 車站位置所屬鄉鎮
	LocationTown string `json:"LocationTown,omitempty"`

	// String
	//
	// 車站位置所屬鄉鎮代碼
	LocationTownCode string `json:"LocationTownCode,omitempty"`

	// String
	//
	// 營運業者代碼
	// Required: true
	OperatorID *string `json:"OperatorID"`

	// String
	//
	// 票價用站牌代碼
	ReservationCode string `json:"ReservationCode,omitempty"`

	// String
	//
	// 車站地址
	// Required: true
	StationAddress *string `json:"StationAddress"`

	// String
	//
	// 車站級別
	StationClass string `json:"StationClass,omitempty"`

	// String
	//
	// 車站代碼
	// Required: true
	StationID *string `json:"StationID"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StationName"`

	// String
	//
	// 車站聯絡電話
	StationPhone string `json:"StationPhone,omitempty"`

	// PointType
	//
	// 車站位置
	StationPosition struct {
		PTXServiceDTORailSpecificationV2PointType
	} `json:"StationPosition,omitempty"`

	// String
	//
	// 車站唯一識別代碼
	// Required: true
	StationUID *string `json:"StationUID"`

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

// Validate validates this p t x service d t o rail specification v2 t r a rail station
func (m *PTXServiceDTORailSpecificationV2TRARailStation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationAddress(formats); err != nil {
		res = append(res, err)
	}

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

func (m *PTXServiceDTORailSpecificationV2TRARailStation) validateOperatorID(formats strfmt.Registry) error {

	if err := validate.Required("OperatorID", "body", m.OperatorID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailStation) validateStationAddress(formats strfmt.Registry) error {

	if err := validate.Required("StationAddress", "body", m.StationAddress); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailStation) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailStation) validateStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailStation) validateStationPosition(formats strfmt.Registry) error {
	if swag.IsZero(m.StationPosition) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailStation) validateStationUID(formats strfmt.Registry) error {

	if err := validate.Required("StationUID", "body", m.StationUID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailStation) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailStation) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 t r a rail station based on the context it is used
func (m *PTXServiceDTORailSpecificationV2TRARailStation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStationPosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailStation) contextValidateStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailStation) contextValidateStationPosition(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRARailStation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRARailStation) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2TRARailStation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
