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

// PTXServiceDTOBusSpecificationV2BusStationGroup BusStationGroup
//
// 組站位資料型別
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.BusStationGroup
type PTXServiceDTOBusSpecificationV2BusStationGroup struct {

	// String
	//
	// 組站位代碼
	// Required: true
	StationGroupID *string `json:"StationGroupID" xml:"String"`

	// NameType
	//
	// 組站位名稱
	// Required: true
	StationGroupName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StationGroupName" xml:"NameType"`

	// PointType
	//
	// 組站位位置
	// Required: true
	StationGroupPosition struct {
		PTXServiceDTOSharedSpecificationV2BasePointType
	} `json:"StationGroupPosition" xml:"PointType"`

	// String
	//
	// 組站位唯一識別代碼，規則為 {業管機關簡碼} + {StationGroupID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	StationGroupUID *string `json:"StationGroupUID" xml:"String"`

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

// Validate validates this p t x service d t o bus specification v2 bus station group
func (m *PTXServiceDTOBusSpecificationV2BusStationGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStationGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationGroupName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationGroupPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationGroupUID(formats); err != nil {
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

func (m *PTXServiceDTOBusSpecificationV2BusStationGroup) validateStationGroupID(formats strfmt.Registry) error {

	if err := validate.Required("StationGroupID", "body", m.StationGroupID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStationGroup) validateStationGroupName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStationGroup) validateStationGroupPosition(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStationGroup) validateStationGroupUID(formats strfmt.Registry) error {

	if err := validate.Required("StationGroupUID", "body", m.StationGroupUID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStationGroup) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStationGroup) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v2 bus station group based on the context it is used
func (m *PTXServiceDTOBusSpecificationV2BusStationGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStationGroupName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStationGroupPosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStationGroup) contextValidateStationGroupName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStationGroup) contextValidateStationGroupPosition(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusStationGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusStationGroup) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2BusStationGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
