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

// ServiceDTOVersion3BusStop Stop
// swagger:model Service.DTO.Version3.Bus.Stop
type ServiceDTOVersion3BusStop struct {

	// 方位角，E:東行;W:西行;S:南行;N:北行;SE:東南行;NE:東北行;SW:西南行;NW:西北行
	Bearing string `json:"Bearing,omitempty"`

	// 站牌權管所屬縣市之代碼(國際ISO 3166-2 三碼城市代碼)[若為公路/國道客運路線則為空值]
	CityCode string `json:"CityCode,omitempty"`

	// 路名
	RoadName string `json:"RoadName,omitempty"`

	// 站位代碼
	StationID string `json:"StationID,omitempty"`

	// 站位唯一識別代碼，規則為 {業管機關簡碼} + {StationID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	StationUID string `json:"StationUID,omitempty"`

	// 站牌地址
	StopAddress string `json:"StopAddress,omitempty"`

	// 站牌簡碼
	StopCode string `json:"StopCode,omitempty"`

	// 站牌詳細說明描述
	StopDescription string `json:"StopDescription,omitempty"`

	// 地區既用中之站牌代碼(為原資料內碼)
	// Required: true
	StopID *string `json:"StopID"`

	// NameType
	//
	// 站牌名稱
	// Required: true
	StopName *ServiceDTOVersion3BaseNameType `json:"StopName"`

	// PointType
	//
	// 站牌位置
	// Required: true
	StopPosition *ServiceDTOVersion3BasePointType `json:"StopPosition"`

	// 站牌唯一識別代碼，規則為 {業管機關簡碼} + {StopID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	StopUID *string `json:"StopUID"`

	// 站牌資訊說明網址
	StopURL string `json:"StopURL,omitempty"`
}

// Validate validates this service d t o version3 bus stop
func (m *ServiceDTOVersion3BusStop) Validate(formats strfmt.Registry) error {
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3BusStop) validateStopID(formats strfmt.Registry) error {

	if err := validate.Required("StopID", "body", m.StopID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusStop) validateStopName(formats strfmt.Registry) error {

	if err := validate.Required("StopName", "body", m.StopName); err != nil {
		return err
	}

	if m.StopName != nil {
		if err := m.StopName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StopName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion3BusStop) validateStopPosition(formats strfmt.Registry) error {

	if err := validate.Required("StopPosition", "body", m.StopPosition); err != nil {
		return err
	}

	if m.StopPosition != nil {
		if err := m.StopPosition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StopPosition")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion3BusStop) validateStopUID(formats strfmt.Registry) error {

	if err := validate.Required("StopUID", "body", m.StopUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3BusStop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3BusStop) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3BusStop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}