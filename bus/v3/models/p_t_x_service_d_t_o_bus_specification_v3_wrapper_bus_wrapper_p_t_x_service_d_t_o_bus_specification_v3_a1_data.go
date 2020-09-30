// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3A1Data BusA1DataList
//
// 基本 wrapper
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.Wrapper.BusWrapper[PTX.Service.DTO.Bus.Specification.V3.A1Data]
type PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3A1Data struct {

	// Array
	//
	// 資料列表
	// Required: true
	A1Datas []*PTXServiceDTOBusSpecificationV3A1Data `json:"A1Datas"`

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode"`

	// 資料總筆數
	Count int64 `json:"Count,omitempty"`

	// Int32
	//
	// 來源端平台資料更新週期(秒)['-1: 不定期更新']
	// Required: true
	SrcUpdateInterval *int32 `json:"SrcUpdateInterval"`

	// DateTime
	//
	// [來源端平臺]資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// Int32
	//
	// [平臺]資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// DateTime
	//
	// [平臺]資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this p t x service d t o bus specification v3 wrapper bus wrapper p t x service d t o bus specification v3 a1 data
func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3A1Data) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateA1Datas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorityCode(formats); err != nil {
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

func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3A1Data) validateA1Datas(formats strfmt.Registry) error {

	if err := validate.Required("A1Datas", "body", m.A1Datas); err != nil {
		return err
	}

	for i := 0; i < len(m.A1Datas); i++ {
		if swag.IsZero(m.A1Datas[i]) { // not required
			continue
		}

		if m.A1Datas[i] != nil {
			if err := m.A1Datas[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("A1Datas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3A1Data) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3A1Data) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3A1Data) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3A1Data) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3A1Data) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3A1Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3A1Data) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3A1Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}