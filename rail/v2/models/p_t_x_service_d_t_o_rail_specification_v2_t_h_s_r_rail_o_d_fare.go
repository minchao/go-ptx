// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTORailSpecificationV2THSRRailODFare RailODFare
//
// 高鐵起訖站票價資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.THSR.RailODFare
type PTXServiceDTORailSpecificationV2THSRRailODFare struct {

	// String
	//
	// 迄點車站代碼
	// Required: true
	DestinationStationID *string `json:"DestinationStationID" xml:"String"`

	// NameType
	//
	// 迄點車站名稱
	// Required: true
	DestinationStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"DestinationStationName" xml:"NameType"`

	// Int32
	//
	// 行駛方向 : [0:'南下',1:'北上']
	// Required: true
	Direction *int64 `json:"Direction"`

	// Array
	//
	// 票價收費資訊(本項僅列標準、商務及自由之基本票價，其他優待票及團體票之折扣計算請參考高鐵網站票價產品一覽表http://www.thsrc.com.tw/tw/Article/ArticleContent/caa6fac8-b875-4ad6-b1e6-96c2902d12a6 說明)
	// Required: true
	Fares []*PTXServiceDTORailSpecificationV2THSRODFare "json:\"Fares\" xml:\"List`1\""

	// String
	//
	// 起點車站代碼
	// Required: true
	OriginStationID *string `json:"OriginStationID" xml:"String"`

	// NameType
	//
	// 起點車站名稱
	// Required: true
	OriginStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"OriginStationName" xml:"NameType"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	SrcUpdateTime *strfmt.DateTime `json:"SrcUpdateTime"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`

	// Int32
	//
	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this p t x service d t o rail specification v2 t h s r rail o d fare
func (m *PTXServiceDTORailSpecificationV2THSRRailODFare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFares(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV2THSRRailODFare) validateDestinationStationID(formats strfmt.Registry) error {

	if err := validate.Required("DestinationStationID", "body", m.DestinationStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailODFare) validateDestinationStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailODFare) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailODFare) validateFares(formats strfmt.Registry) error {

	if err := validate.Required("Fares", "body", m.Fares); err != nil {
		return err
	}

	for i := 0; i < len(m.Fares); i++ {
		if swag.IsZero(m.Fares[i]) { // not required
			continue
		}

		if m.Fares[i] != nil {
			if err := m.Fares[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Fares" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Fares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailODFare) validateOriginStationID(formats strfmt.Registry) error {

	if err := validate.Required("OriginStationID", "body", m.OriginStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailODFare) validateOriginStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailODFare) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailODFare) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailODFare) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 t h s r rail o d fare based on the context it is used
func (m *PTXServiceDTORailSpecificationV2THSRRailODFare) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestinationStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFares(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailODFare) contextValidateDestinationStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailODFare) contextValidateFares(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Fares); i++ {

		if m.Fares[i] != nil {
			if err := m.Fares[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Fares" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Fares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailODFare) contextValidateOriginStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2THSRRailODFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2THSRRailODFare) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2THSRRailODFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
