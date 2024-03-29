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

// PTXServiceDTOBusSpecificationV2BusShape BusShape
//
// 公車線型資料
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.BusShape
type PTXServiceDTOBusSpecificationV2BusShape struct {

	// Int32
	//
	// 去返程，若無值則表示來源尚無區分去返程 : [0:'去程',1:'返程',2:'迴圈',255:'未知']
	Direction int64 `json:"Direction,omitempty"`

	// String
	//
	// 路線軌跡編碼(encoded polyline)
	// Required: true
	EncodedPolyline *string `json:"EncodedPolyline" xml:"EncodedPolyline"`

	// String
	//
	// well-known text，為路線軌跡資料
	// Required: true
	Geometry *string `json:"Geometry" xml:"Geometry"`

	// String
	//
	// 地區既用中之路線代碼(為原資料內碼)
	// Required: true
	RouteID *string `json:"RouteID" xml:"RouteID"`

	// NameType
	//
	// 路線名稱
	// Required: true
	RouteName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"RouteName" xml:"NameType"`

	// String
	//
	// 路線唯一識別代碼，規則為 {業管機關簡碼} + {RouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	RouteUID *string `json:"RouteUID" xml:"RouteUID"`

	// String
	//
	// 附屬路線唯一識別代碼，規則為 {業管機關簡碼} + {SubRouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	SubRouteID string `json:"SubRouteID,omitempty" xml:"SubRouteID,omitempty"`

	// NameType
	//
	// 附屬路線名稱
	SubRouteName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"SubRouteName,omitempty" xml:"NameType,omitempty"`

	// String
	//
	// 附屬路線唯一識別代碼，規則為 {業管機關簡碼} + {SubRouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	SubRouteUID *string `json:"SubRouteUID" xml:"SubRouteUID"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`

	// Int32
	//
	// 資料版本編號(由於該服務資料不再版控，固定帶入版號0)
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this p t x service d t o bus specification v2 bus shape
func (m *PTXServiceDTOBusSpecificationV2BusShape) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEncodedPolyline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeometry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubRouteName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubRouteUID(formats); err != nil {
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

func (m *PTXServiceDTOBusSpecificationV2BusShape) validateEncodedPolyline(formats strfmt.Registry) error {

	if err := validate.Required("EncodedPolyline", "body", m.EncodedPolyline); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusShape) validateGeometry(formats strfmt.Registry) error {

	if err := validate.Required("Geometry", "body", m.Geometry); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusShape) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusShape) validateRouteName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusShape) validateRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("RouteUID", "body", m.RouteUID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusShape) validateSubRouteName(formats strfmt.Registry) error {
	if swag.IsZero(m.SubRouteName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusShape) validateSubRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("SubRouteUID", "body", m.SubRouteUID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusShape) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusShape) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v2 bus shape based on the context it is used
func (m *PTXServiceDTOBusSpecificationV2BusShape) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRouteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubRouteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusShape) contextValidateRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusShape) contextValidateSubRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusShape) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusShape) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2BusShape
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
