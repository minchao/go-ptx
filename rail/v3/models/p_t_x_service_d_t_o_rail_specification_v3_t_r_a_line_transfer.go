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

// PTXServiceDTORailSpecificationV3TRALineTransfer LineTransfer
//
// 路線站間轉乘基本資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.LineTransfer
type PTXServiceDTORailSpecificationV3TRALineTransfer struct {

	// String
	//
	// 路線間轉乘(起)之路線代碼
	// Required: true
	FromLineID *string `json:"FromLineID" xml:"FromLineID"`

	// NameType
	//
	// 路線間轉乘(起)之路線名稱
	// Required: true
	FromLineName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"FromLineName" xml:"NameType"`

	// String
	//
	// 路線間轉乘(起)之車站代碼
	// Required: true
	FromStationID *string `json:"FromStationID" xml:"FromStationID"`

	// NameType
	//
	// 路線間轉乘(起)之車站名稱
	// Required: true
	FromStationName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"FromStationName" xml:"NameType"`

	// Double
	//
	// 轉乘耗時參考時間(分)
	// Required: true
	MinTransferTime *float64 `json:"MinTransferTime"`

	// String
	//
	// 路線間轉乘(迄)之路線代碼
	// Required: true
	ToLineID *string `json:"ToLineID" xml:"ToLineID"`

	// NameType
	//
	// 路線間轉乘(迄)之路線名稱
	// Required: true
	ToLineName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"ToLineName" xml:"NameType"`

	// String
	//
	// 路線間轉乘(迄)之車站代碼
	// Required: true
	ToStationID *string `json:"ToStationID" xml:"ToStationID"`

	// NameType
	//
	// 路線間轉乘(迄)之車站名稱
	// Required: true
	ToStationName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"ToStationName" xml:"NameType"`

	// String
	//
	// 轉乘方式文字描述
	// Required: true
	TransferDescription *string `json:"TransferDescription" xml:"TransferDescription"`
}

// Validate validates this p t x service d t o rail specification v3 t r a line transfer
func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFromLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromLineName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinTransferTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToLineName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransferDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) validateFromLineID(formats strfmt.Registry) error {

	if err := validate.Required("FromLineID", "body", m.FromLineID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) validateFromLineName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) validateFromStationID(formats strfmt.Registry) error {

	if err := validate.Required("FromStationID", "body", m.FromStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) validateFromStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) validateMinTransferTime(formats strfmt.Registry) error {

	if err := validate.Required("MinTransferTime", "body", m.MinTransferTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) validateToLineID(formats strfmt.Registry) error {

	if err := validate.Required("ToLineID", "body", m.ToLineID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) validateToLineName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) validateToStationID(formats strfmt.Registry) error {

	if err := validate.Required("ToStationID", "body", m.ToStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) validateToStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) validateTransferDescription(formats strfmt.Registry) error {

	if err := validate.Required("TransferDescription", "body", m.TransferDescription); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a line transfer based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFromLineName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFromStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateToLineName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateToStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) contextValidateFromLineName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) contextValidateFromStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) contextValidateToLineName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) contextValidateToStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRALineTransfer) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRALineTransfer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
