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

// PTXServiceDTOBusSpecificationV3StopOfRouteStop Stop
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.StopOfRoute+Stop
type PTXServiceDTOBusSpecificationV3StopOfRouteStop struct {

	// integer
	//
	// 上下車站別 : [-1:'可下車',0:'可上下車',1:'可上車']
	// Required: true
	BoardingType *string `json:"BoardingType"`

	// String
	//
	// 累積行駛距離
	CumulativeDistance string `json:"CumulativeDistance,omitempty" xml:"String,omitempty"`

	// Boolean
	//
	// 是否為分段點
	IsSectionPoint bool `json:"IsSectionPoint,omitempty"`

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
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"StopName" xml:"NameType"`

	// PointType
	//
	// 站牌位置
	// Required: true
	StopPosition struct {
		PTXServiceDTOSharedSpecificationV3BasePointType
	} `json:"StopPosition" xml:"PointType"`

	// Int32
	//
	// 路線經過站牌之順序
	// Required: true
	StopSequence *int32 `json:"StopSequence"`

	// String
	//
	// 站牌唯一識別代碼，規則為 {業管機關簡碼} + {StopID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	StopUID *string `json:"StopUID" xml:"String"`
}

// Validate validates this p t x service d t o bus specification v3 stop of route stop
func (m *PTXServiceDTOBusSpecificationV3StopOfRouteStop) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoardingType(formats); err != nil {
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

	if err := m.validateStopSequence(formats); err != nil {
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

func (m *PTXServiceDTOBusSpecificationV3StopOfRouteStop) validateBoardingType(formats strfmt.Registry) error {

	if err := validate.Required("BoardingType", "body", m.BoardingType); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3StopOfRouteStop) validateStopID(formats strfmt.Registry) error {

	if err := validate.Required("StopID", "body", m.StopID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3StopOfRouteStop) validateStopName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3StopOfRouteStop) validateStopPosition(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3StopOfRouteStop) validateStopSequence(formats strfmt.Registry) error {

	if err := validate.Required("StopSequence", "body", m.StopSequence); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3StopOfRouteStop) validateStopUID(formats strfmt.Registry) error {

	if err := validate.Required("StopUID", "body", m.StopUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v3 stop of route stop based on the context it is used
func (m *PTXServiceDTOBusSpecificationV3StopOfRouteStop) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PTXServiceDTOBusSpecificationV3StopOfRouteStop) contextValidateStopName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3StopOfRouteStop) contextValidateStopPosition(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3StopOfRouteStop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3StopOfRouteStop) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3StopOfRouteStop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
