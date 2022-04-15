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

// PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Vessel ShipVesselList
//
// swagger:model PTX.API.Ship.Model.ShipWrapper[PTX.Service.DTO.Ship.Specification.V3.Vessel]
type PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Vessel struct {

	// String
	//
	// 業管機關簡碼
	AuthorityCode string `json:"AuthorityCode,omitempty" xml:"String,omitempty"`

	// 資料總筆數<span class="emphasis fas fa-pen" rel="與來源Inbound XML不同，為提供資料的總筆數[該欄位由本平台自動產製]"></span>
	Count int64 `json:"Count,omitempty"`

	// [來源端平臺]資料更新週期(秒)['-1: 不定期更新']<span class="emphasis fas fa-pen" rel="為來源Inbound XML中的UpdateInterval [該欄位由本平台自動產製]"></span>
	// Required: true
	SrcUpdateInterval *int32 `json:"SrcUpdateInterval"`

	// [來源端平臺]資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)<span class="emphasis fas fa-pen" rel="為來源Inbound XML中的UpdateTime [該欄位由本平台自動產製]"></span>
	// Required: true
	// Format: date-time
	SrcUpdateTime *strfmt.DateTime `json:"SrcUpdateTime"`

	// Int32
	//
	// 本平台資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// XML更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Format: date-time
	UpdateTime strfmt.DateTime `json:"UpdateTime,omitempty"`

	// 資料版本編號<span class="emphasis fas fa-pen" rel="與來源Inbound XML不同，為提供資料的版本編號[該欄位由本平台自動產製]"></span>
	// Required: true
	VersionID *int32 `json:"VersionID"`

	// Array
	//
	// 資料(陣列)
	Vessels []*PTXServiceDTOShipSpecificationV3Vessel "json:\"Vessels\" xml:\"List`1\""
}

// Validate validates this p t x API ship model ship wrapper p t x service d t o ship specification v3 vessel
func (m *PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Vessel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSrcUpdateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVessels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Vessel) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Vessel) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Vessel) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Vessel) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Vessel) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Vessel) validateVessels(formats strfmt.Registry) error {
	if swag.IsZero(m.Vessels) { // not required
		return nil
	}

	for i := 0; i < len(m.Vessels); i++ {
		if swag.IsZero(m.Vessels[i]) { // not required
			continue
		}

		if m.Vessels[i] != nil {
			if err := m.Vessels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Vessels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Vessels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this p t x API ship model ship wrapper p t x service d t o ship specification v3 vessel based on the context it is used
func (m *PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Vessel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVessels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Vessel) contextValidateVessels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Vessels); i++ {

		if m.Vessels[i] != nil {
			if err := m.Vessels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Vessels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Vessels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Vessel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Vessel) UnmarshalBinary(b []byte) error {
	var res PTXAPIShipModelShipWrapperPTXServiceDTOShipSpecificationV3Vessel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}