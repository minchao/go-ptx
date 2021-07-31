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

// PTXServiceDTORailSpecificationV3LiteTrainStation Station
//
// 小火車車站資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.LiteTrain.Station
type PTXServiceDTORailSpecificationV3LiteTrainStation struct {

	// String
	//
	// 訂票車站代碼
	ReservationCode string `json:"ReservationCode,omitempty" xml:"String,omitempty"`

	// String
	//
	// 車站地址
	StationAddress string `json:"StationAddress,omitempty" xml:"String,omitempty"`

	// String
	//
	// 車站級別 = ['0: 特等', '1: 一等', '2: 二等', '3: 三等', '4: 簡易', '5: 招呼', '6: 號誌', '7:其他']
	// Required: true
	StationClass *string `json:"StationClass" xml:"String"`

	// String
	//
	// 車站代碼
	// Required: true
	StationID *string `json:"StationID" xml:"String"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"StationName" xml:"NameType"`

	// String
	//
	// 車站聯絡電話
	StationPhone string `json:"StationPhone,omitempty" xml:"String,omitempty"`

	// PointType
	//
	// 車站座標(WGS84)
	// Required: true
	StationPosition struct {
		PTXServiceDTOSharedSpecificationV3BasePointType
	} `json:"StationPosition" xml:"PointType"`

	// String
	//
	// 車站唯一識別代碼
	// Required: true
	StationUID *string `json:"StationUID" xml:"String"`

	// String
	//
	// 車站資訊說明網址
	StationURL string `json:"StationURL,omitempty" xml:"String,omitempty"`
}

// Validate validates this p t x service d t o rail specification v3 lite train station
func (m *PTXServiceDTORailSpecificationV3LiteTrainStation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStationClass(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStation) validateStationClass(formats strfmt.Registry) error {

	if err := validate.Required("StationClass", "body", m.StationClass); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStation) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStation) validateStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStation) validateStationPosition(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStation) validateStationUID(formats strfmt.Registry) error {

	if err := validate.Required("StationUID", "body", m.StationUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 lite train station based on the context it is used
func (m *PTXServiceDTORailSpecificationV3LiteTrainStation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PTXServiceDTORailSpecificationV3LiteTrainStation) contextValidateStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStation) contextValidateStationPosition(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3LiteTrainStation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3LiteTrainStation) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3LiteTrainStation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
