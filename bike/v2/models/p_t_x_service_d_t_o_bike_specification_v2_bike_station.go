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

// PTXServiceDTOBikeSpecificationV2BikeStation BikeStation
//
// 自行車站點資訊
//
// swagger:model PTX.Service.DTO.Bike.Specification.V2.BikeStation
type PTXServiceDTOBikeSpecificationV2BikeStation struct {

	// String
	//
	// 業管單位代碼
	AuthorityID string `json:"AuthorityID,omitempty" xml:"String,omitempty"`

	// Int32
	//
	// 可容納之自行車總數
	BikesCapacity int32 `json:"BikesCapacity,omitempty"`

	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Format: date-time
	SrcUpdateTime strfmt.DateTime `json:"SrcUpdateTime,omitempty"`

	// NameType
	//
	// 站點地址
	StationAddress struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StationAddress,omitempty" xml:"NameType,omitempty"`

	// String
	//
	// 站點代碼
	StationID string `json:"StationID,omitempty" xml:"String,omitempty"`

	// NameType
	//
	// 站點名稱
	StationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StationName,omitempty" xml:"NameType,omitempty"`

	// PointType
	//
	// 站點位置
	StationPosition struct {
		PTXServiceDTOSharedSpecificationV2BasePointType
	} `json:"StationPosition,omitempty" xml:"PointType,omitempty"`

	// String
	//
	// 站點唯一識別代碼，規則為 {業管機關代碼} + {StationID}，其中 {業管機關代碼} 可於Authority API中的AuthorityCode欄位查詢
	StationUID string `json:"StationUID,omitempty" xml:"String,omitempty"`

	// String
	//
	// 站點描述
	StopDescription string `json:"StopDescription,omitempty" xml:"String,omitempty"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o bike specification v2 bike station
func (m *PTXServiceDTOBikeSpecificationV2BikeStation) Validate(formats strfmt.Registry) error {
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

func (m *PTXServiceDTOBikeSpecificationV2BikeStation) validateSrcUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.SrcUpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBikeSpecificationV2BikeStation) validateStationAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.StationAddress) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOBikeSpecificationV2BikeStation) validateStationName(formats strfmt.Registry) error {
	if swag.IsZero(m.StationName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOBikeSpecificationV2BikeStation) validateStationPosition(formats strfmt.Registry) error {
	if swag.IsZero(m.StationPosition) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOBikeSpecificationV2BikeStation) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bike specification v2 bike station based on the context it is used
func (m *PTXServiceDTOBikeSpecificationV2BikeStation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStationAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

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

func (m *PTXServiceDTOBikeSpecificationV2BikeStation) contextValidateStationAddress(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBikeSpecificationV2BikeStation) contextValidateStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBikeSpecificationV2BikeStation) contextValidateStationPosition(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBikeSpecificationV2BikeStation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBikeSpecificationV2BikeStation) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBikeSpecificationV2BikeStation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
