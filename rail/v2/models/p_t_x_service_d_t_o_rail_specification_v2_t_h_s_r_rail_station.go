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

// PTXServiceDTORailSpecificationV2THSRRailStation RailStation
//
// 高鐵車站資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.THSR.RailStation
type PTXServiceDTORailSpecificationV2THSRRailStation struct {

	// String
	//
	// 車站位置所屬縣市
	LocationCity string `json:"LocationCity,omitempty" xml:"LocationCity,omitempty"`

	// String
	//
	// 車站位置所屬縣市代碼
	LocationCityCode string `json:"LocationCityCode,omitempty" xml:"LocationCityCode,omitempty"`

	// String
	//
	// 車站位置所屬鄉鎮
	LocationTown string `json:"LocationTown,omitempty" xml:"LocationTown,omitempty"`

	// String
	//
	// 車站位置所屬鄉鎮代碼
	LocationTownCode string `json:"LocationTownCode,omitempty" xml:"LocationTownCode,omitempty"`

	// String
	//
	// 營運業者代碼
	// Required: true
	OperatorID *string `json:"OperatorID" xml:"OperatorID"`

	// String
	//
	// 車站地址
	// Required: true
	StationAddress *string `json:"StationAddress" xml:"StationAddress"`

	// String
	//
	// 車站簡碼(訂票系統用)
	// Required: true
	StationCode *string `json:"StationCode" xml:"StationCode"`

	// String
	//
	// 車站代碼
	// Required: true
	StationID *string `json:"StationID" xml:"StationID"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StationName" xml:"NameType"`

	// String
	//
	// 車站聯絡電話
	StationPhone string `json:"StationPhone,omitempty" xml:"StationPhone,omitempty"`

	// PointType
	//
	// 車站位置
	StationPosition struct {
		PTXServiceDTORailSpecificationV2PointType
	} `json:"StationPosition,omitempty" xml:"PointType,omitempty"`

	// String
	//
	// 車站唯一識別代碼
	// Required: true
	StationUID *string `json:"StationUID" xml:"StationUID"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`

	// Int32
	//
	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this p t x service d t o rail specification v2 t h s r rail station
func (m *PTXServiceDTORailSpecificationV2THSRRailStation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationCode(formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV2THSRRailStation) validateOperatorID(formats strfmt.Registry) error {

	if err := validate.Required("OperatorID", "body", m.OperatorID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStation) validateStationAddress(formats strfmt.Registry) error {

	if err := validate.Required("StationAddress", "body", m.StationAddress); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStation) validateStationCode(formats strfmt.Registry) error {

	if err := validate.Required("StationCode", "body", m.StationCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStation) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStation) validateStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStation) validateStationPosition(formats strfmt.Registry) error {
	if swag.IsZero(m.StationPosition) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStation) validateStationUID(formats strfmt.Registry) error {

	if err := validate.Required("StationUID", "body", m.StationUID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStation) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStation) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 t h s r rail station based on the context it is used
func (m *PTXServiceDTORailSpecificationV2THSRRailStation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PTXServiceDTORailSpecificationV2THSRRailStation) contextValidateStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStation) contextValidateStationPosition(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2THSRRailStation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2THSRRailStation) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2THSRRailStation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
