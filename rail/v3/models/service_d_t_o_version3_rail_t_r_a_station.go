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

// ServiceDTOVersion3RailTRAStation Station
//
// 台鐵車站資料
// swagger:model Service.DTO.Version3.Rail.TRA.Station
type ServiceDTOVersion3RailTRAStation struct {

	// 訂票車站代碼
	ReservationCode string `json:"ReservationCode,omitempty"`

	// 車站地址
	StationAddress string `json:"StationAddress,omitempty"`

	// 車站級別 = ['0: 特等', '1: 一等', '2: 二等', '3: 三等', '4: 簡易', '5: 招呼', '6: 號誌', 'A: 貨運', 'B: 基地', 'X: 非車']
	StationClass string `json:"StationClass,omitempty"`

	// 臺鐵車站代碼
	// Required: true
	StationID *string `json:"StationID"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName *ServiceDTOVersion3BaseNameType `json:"StationName"`

	// 車站聯絡電話
	StationPhone string `json:"StationPhone,omitempty"`

	// PointType
	//
	// 車站座標(WGS84)
	// Required: true
	StationPosition *ServiceDTOVersion3BasePointType `json:"StationPosition"`

	// 臺鐵車站唯一識別代碼
	// Required: true
	StationUID *string `json:"StationUID"`

	// 車站資訊說明網址
	StationURL string `json:"StationURL,omitempty"`
}

// Validate validates this service d t o version3 rail t r a station
func (m *ServiceDTOVersion3RailTRAStation) Validate(formats strfmt.Registry) error {
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3RailTRAStation) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRAStation) validateStationName(formats strfmt.Registry) error {

	if err := validate.Required("StationName", "body", m.StationName); err != nil {
		return err
	}

	if m.StationName != nil {
		if err := m.StationName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StationName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRAStation) validateStationPosition(formats strfmt.Registry) error {

	if err := validate.Required("StationPosition", "body", m.StationPosition); err != nil {
		return err
	}

	if m.StationPosition != nil {
		if err := m.StationPosition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StationPosition")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRAStation) validateStationUID(formats strfmt.Registry) error {

	if err := validate.Required("StationUID", "body", m.StationUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRAStation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRAStation) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3RailTRAStation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
