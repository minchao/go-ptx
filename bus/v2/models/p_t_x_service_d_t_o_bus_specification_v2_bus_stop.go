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

// PTXServiceDTOBusSpecificationV2BusStop BusStop
//
// 站牌/位資料型別
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.BusStop
type PTXServiceDTOBusSpecificationV2BusStop struct {

	// String
	//
	// 業管機關代碼
	// Required: true
	AuthorityID *string `json:"AuthorityID" xml:"String"`

	// String
	//
	// 方位角，E:東行;W:西行;S:南行;N:北行;SE:東南行;NE:東北行;SW:西南行;NW:西北行
	Bearing string `json:"Bearing,omitempty" xml:"String,omitempty"`

	// String
	//
	// 站牌權管所屬縣市(相當於市區公車API的City參數)[若為公路/國道客運路線則為空值]
	City string `json:"City,omitempty" xml:"String,omitempty"`

	// String
	//
	// 站牌權管所屬縣市之代碼(國際ISO 3166-2 三碼城市代碼)[若為公路/國道客運路線則為空值]
	CityCode string `json:"CityCode,omitempty" xml:"String,omitempty"`

	// String
	//
	// 站牌位置縣市之代碼(國際ISO 3166-2 三碼城市代碼)[若為公路/國道客運路線則為空值]
	LocationCityCode string `json:"LocationCityCode,omitempty" xml:"String,omitempty"`

	// String
	//
	// 站牌所屬的組站位ID
	// Required: true
	StationGroupID *string `json:"StationGroupID" xml:"String"`

	// String
	//
	// 站牌所屬的站位ID
	StationID string `json:"StationID,omitempty" xml:"String,omitempty"`

	// String
	//
	// 站牌地址
	StopAddress string `json:"StopAddress,omitempty" xml:"String,omitempty"`

	// String
	//
	// 站牌詳細說明描述
	StopDescription string `json:"StopDescription,omitempty" xml:"String,omitempty"`

	// String
	//
	// 地區既用中之站牌代碼(為原資料內碼)
	// Required: true
	StopID *string `json:"StopID" xml:"String"`

	// NameType
	//
	// 站牌名稱
	// Required: true
	StopName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StopName" xml:"NameType"`

	// PointType
	//
	// 站牌位置
	// Required: true
	StopPosition struct {
		PTXServiceDTOSharedSpecificationV2BasePointType
	} `json:"StopPosition" xml:"PointType"`

	// String
	//
	// 站牌唯一識別代碼，規則為 {業管機關簡碼} + {StopID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	StopUID *string `json:"StopUID" xml:"String"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// Int32
	//
	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this p t x service d t o bus specification v2 bus stop
func (m *PTXServiceDTOBusSpecificationV2BusStop) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopUID(formats); err != nil {
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

func (m *PTXServiceDTOBusSpecificationV2BusStop) validateAuthorityID(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityID", "body", m.AuthorityID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStop) validateStationGroupID(formats strfmt.Registry) error {

	if err := validate.Required("StationGroupID", "body", m.StationGroupID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStop) validateStopID(formats strfmt.Registry) error {

	if err := validate.Required("StopID", "body", m.StopID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStop) validateStopName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStop) validateStopPosition(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStop) validateStopUID(formats strfmt.Registry) error {

	if err := validate.Required("StopUID", "body", m.StopUID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStop) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStop) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v2 bus stop based on the context it is used
func (m *PTXServiceDTOBusSpecificationV2BusStop) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStopName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStopPosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStop) contextValidateStopName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStop) contextValidateStopPosition(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusStop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusStop) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2BusStop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
