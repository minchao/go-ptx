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

// PTXServiceDTOBusSpecificationV2StationStop StationStop
//
// 站牌與路線資料型別
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.StationStop
type PTXServiceDTOBusSpecificationV2StationStop struct {

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
	// 地區既用中之站牌代碼(為原資料內碼)
	// Required: true
	StopID *string `json:"StopID" xml:"StopID"`

	// NameType
	//
	// 站牌名稱
	// Required: true
	StopName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StopName" xml:"NameType"`

	// String
	//
	// 站牌唯一識別代碼，規則為 {業管機關簡碼} + {StopID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	StopUID *string `json:"StopUID" xml:"StopUID"`
}

// Validate validates this p t x service d t o bus specification v2 station stop
func (m *PTXServiceDTOBusSpecificationV2StationStop) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2StationStop) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2StationStop) validateRouteName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2StationStop) validateRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("RouteUID", "body", m.RouteUID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2StationStop) validateStopID(formats strfmt.Registry) error {

	if err := validate.Required("StopID", "body", m.StopID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2StationStop) validateStopName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2StationStop) validateStopUID(formats strfmt.Registry) error {

	if err := validate.Required("StopUID", "body", m.StopUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v2 station stop based on the context it is used
func (m *PTXServiceDTOBusSpecificationV2StationStop) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRouteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStopName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2StationStop) contextValidateRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2StationStop) contextValidateStopName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2StationStop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2StationStop) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2StationStop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
