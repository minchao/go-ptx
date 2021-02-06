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

// PTXServiceDTOBusSpecificationV2EmbeddedStop Stop
//
// 站牌代碼資料
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.Embedded.Stop
type PTXServiceDTOBusSpecificationV2EmbeddedStop struct {

	// String
	//
	// 站牌位置縣市之代碼(國際ISO 3166-2 三碼城市代碼)[若為公路/國道客運路線則為空值]
	LocationCityCode string `json:"LocationCityCode,omitempty"`

	// String
	//
	// 站牌所屬的組站位ID
	// Required: true
	StationGroupID *string `json:"StationGroupID"`

	// String
	//
	// 站牌所屬的站位ID
	StationID string `json:"StationID,omitempty"`

	// integer
	//
	// 上下車站別 : [-1:'可下車',0:'可上下車',1:'可上車']
	StopBoarding int32 `json:"StopBoarding,omitempty"`

	// String
	//
	// 地區既用中之站牌代碼(為原資料內碼)
	// Required: true
	StopID *string `json:"StopID"`

	// NameType
	//
	// 站牌名稱
	// Required: true
	StopName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StopName"`

	// PointType
	//
	// 站牌位置
	// Required: true
	StopPosition struct {
		PTXServiceDTOSharedSpecificationV2BasePointType
	} `json:"StopPosition"`

	// Int32
	//
	// 路線經過站牌之順序
	// Required: true
	StopSequence *int32 `json:"StopSequence"`

	// String
	//
	// 站牌唯一識別代碼，規則為 {業管機關簡碼} + {StopID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	StopUID *string `json:"StopUID"`
}

// Validate validates this p t x service d t o bus specification v2 embedded stop
func (m *PTXServiceDTOBusSpecificationV2EmbeddedStop) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *PTXServiceDTOBusSpecificationV2EmbeddedStop) validateStationGroupID(formats strfmt.Registry) error {

	if err := validate.Required("StationGroupID", "body", m.StationGroupID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2EmbeddedStop) validateStopID(formats strfmt.Registry) error {

	if err := validate.Required("StopID", "body", m.StopID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2EmbeddedStop) validateStopName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2EmbeddedStop) validateStopPosition(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2EmbeddedStop) validateStopSequence(formats strfmt.Registry) error {

	if err := validate.Required("StopSequence", "body", m.StopSequence); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2EmbeddedStop) validateStopUID(formats strfmt.Registry) error {

	if err := validate.Required("StopUID", "body", m.StopUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v2 embedded stop based on the context it is used
func (m *PTXServiceDTOBusSpecificationV2EmbeddedStop) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PTXServiceDTOBusSpecificationV2EmbeddedStop) contextValidateStopName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2EmbeddedStop) contextValidateStopPosition(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2EmbeddedStop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2EmbeddedStop) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2EmbeddedStop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}