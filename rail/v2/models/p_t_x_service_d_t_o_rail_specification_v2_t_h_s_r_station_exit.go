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

// PTXServiceDTORailSpecificationV2THSRStationExit StationExit
//
// 高鐵車站出入口資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.THSR.StationExit
type PTXServiceDTORailSpecificationV2THSRStationExit struct {

	// 是否有電梯
	// Required: true
	Elevator *bool `json:"Elevator"`

	// 是否有電扶梯
	// Required: true
	Escalator *bool `json:"Escalator"`

	// String
	//
	// 出入口代碼
	// Required: true
	ExitID *string `json:"ExitID" xml:"ExitID"`

	// NameType
	//
	// 出入口名稱
	// Required: true
	ExitName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"ExitName" xml:"NameType"`

	// PointType
	//
	// 出入口座標
	// Required: true
	ExitPosition struct {
		PTXServiceDTOSharedSpecificationV2BasePointType
	} `json:"ExitPosition" xml:"PointType"`

	// String
	//
	// 地址描述
	// Required: true
	LocationDescription *string `json:"LocationDescription" xml:"LocationDescription"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	SrcUpdateTime *strfmt.DateTime `json:"SrcUpdateTime"`

	// 是否有樓梯
	// Required: true
	Stair *bool `json:"Stair"`

	// String
	//
	// 車站代碼
	// Required: true
	StationID *string `json:"StationID" xml:"StationID"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StationName" xml:"NameType"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o rail specification v2 t h s r station exit
func (m *PTXServiceDTORailSpecificationV2THSRStationExit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElevator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEscalator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExitID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExitName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExitPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStair(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV2THSRStationExit) validateElevator(formats strfmt.Registry) error {

	if err := validate.Required("Elevator", "body", m.Elevator); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRStationExit) validateEscalator(formats strfmt.Registry) error {

	if err := validate.Required("Escalator", "body", m.Escalator); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRStationExit) validateExitID(formats strfmt.Registry) error {

	if err := validate.Required("ExitID", "body", m.ExitID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRStationExit) validateExitName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRStationExit) validateExitPosition(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRStationExit) validateLocationDescription(formats strfmt.Registry) error {

	if err := validate.Required("LocationDescription", "body", m.LocationDescription); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRStationExit) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRStationExit) validateStair(formats strfmt.Registry) error {

	if err := validate.Required("Stair", "body", m.Stair); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRStationExit) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRStationExit) validateStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRStationExit) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 t h s r station exit based on the context it is used
func (m *PTXServiceDTORailSpecificationV2THSRStationExit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExitName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExitPosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRStationExit) contextValidateExitName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRStationExit) contextValidateExitPosition(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRStationExit) contextValidateStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2THSRStationExit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2THSRStationExit) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2THSRStationExit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
