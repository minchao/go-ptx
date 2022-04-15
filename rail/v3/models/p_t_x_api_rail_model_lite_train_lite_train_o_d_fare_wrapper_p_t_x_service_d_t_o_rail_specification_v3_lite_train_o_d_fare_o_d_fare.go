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

// PTXAPIRailModelLiteTrainLiteTrainODFareWrapperPTXServiceDTORailSpecificationV3LiteTrainODFareODFare LiteTrainODFareList
//
// swagger:model PTX.API.Rail.Model.LiteTrain.LiteTrainODFareWrapper[PTX.Service.DTO.Rail.Specification.V3.LiteTrain.ODFare.ODFare]
type PTXAPIRailModelLiteTrainLiteTrainODFareWrapperPTXServiceDTORailSpecificationV3LiteTrainODFareODFare struct {

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode" xml:"String"`

	// 資料總筆數
	Count int64 `json:"Count,omitempty"`

	// String
	//
	// 有效起始日期
	EffectiveDate string `json:"EffectiveDate,omitempty" xml:"String,omitempty"`

	// String
	//
	// 有效終止日期
	ExpireDate string `json:"ExpireDate,omitempty" xml:"String,omitempty"`

	// Array
	//
	// 資料(陣列)
	// Required: true
	ODFares []*PTXServiceDTORailSpecificationV3LiteTrainODFareODFare "json:\"ODFares\" xml:\"List`1\""

	// Int32
	//
	// 來源端平台資料更新週期(秒)['-1: 不定期更新']
	// Required: true
	SrcUpdateInterval *int32 `json:"SrcUpdateInterval"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	SrcUpdateTime *strfmt.DateTime `json:"SrcUpdateTime"`

	// String
	//
	// 資料版本
	SrcVersion string `json:"SrcVersion,omitempty" xml:"String,omitempty"`

	// Int32
	//
	// 本平台資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x API rail model lite train lite train o d fare wrapper p t x service d t o rail specification v3 lite train o d fare o d fare
func (m *PTXAPIRailModelLiteTrainLiteTrainODFareWrapperPTXServiceDTORailSpecificationV3LiteTrainODFareODFare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateODFares(formats); err != nil {
		res = append(res, err)
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXAPIRailModelLiteTrainLiteTrainODFareWrapperPTXServiceDTORailSpecificationV3LiteTrainODFareODFare) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelLiteTrainLiteTrainODFareWrapperPTXServiceDTORailSpecificationV3LiteTrainODFareODFare) validateODFares(formats strfmt.Registry) error {

	if err := validate.Required("ODFares", "body", m.ODFares); err != nil {
		return err
	}

	for i := 0; i < len(m.ODFares); i++ {
		if swag.IsZero(m.ODFares[i]) { // not required
			continue
		}

		if m.ODFares[i] != nil {
			if err := m.ODFares[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ODFares" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ODFares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXAPIRailModelLiteTrainLiteTrainODFareWrapperPTXServiceDTORailSpecificationV3LiteTrainODFareODFare) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelLiteTrainLiteTrainODFareWrapperPTXServiceDTORailSpecificationV3LiteTrainODFareODFare) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelLiteTrainLiteTrainODFareWrapperPTXServiceDTORailSpecificationV3LiteTrainODFareODFare) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelLiteTrainLiteTrainODFareWrapperPTXServiceDTORailSpecificationV3LiteTrainODFareODFare) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x API rail model lite train lite train o d fare wrapper p t x service d t o rail specification v3 lite train o d fare o d fare based on the context it is used
func (m *PTXAPIRailModelLiteTrainLiteTrainODFareWrapperPTXServiceDTORailSpecificationV3LiteTrainODFareODFare) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateODFares(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXAPIRailModelLiteTrainLiteTrainODFareWrapperPTXServiceDTORailSpecificationV3LiteTrainODFareODFare) contextValidateODFares(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ODFares); i++ {

		if m.ODFares[i] != nil {
			if err := m.ODFares[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ODFares" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ODFares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXAPIRailModelLiteTrainLiteTrainODFareWrapperPTXServiceDTORailSpecificationV3LiteTrainODFareODFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXAPIRailModelLiteTrainLiteTrainODFareWrapperPTXServiceDTORailSpecificationV3LiteTrainODFareODFare) UnmarshalBinary(b []byte) error {
	var res PTXAPIRailModelLiteTrainLiteTrainODFareWrapperPTXServiceDTORailSpecificationV3LiteTrainODFareODFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}