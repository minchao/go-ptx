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

// PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop BusStopList
//
// 版本 wrapper (具有版本資訊)
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.Wrapper.BusVWrapper[PTX.Service.DTO.Bus.Specification.V3.Stop]
type PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop struct {

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode" xml:"AuthorityCode"`

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
	// Format: date-time
	SrcUpdateTime *strfmt.DateTime `json:"SrcUpdateTime"`

	// Array
	//
	// 資料列表
	// Required: true
	Stops []*PTXServiceDTOBusSpecificationV3Stop "json:\"Stops\" xml:\"List`1\""

	// Int32
	//
	// [平臺]資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// DateTime
	//
	// [平臺]資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`

	// Int32
	//
	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this p t x service d t o bus specification v3 wrapper bus v wrapper p t x service d t o bus specification v3 stop
func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStops(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop) validateStops(formats strfmt.Registry) error {

	if err := validate.Required("Stops", "body", m.Stops); err != nil {
		return err
	}

	for i := 0; i < len(m.Stops); i++ {
		if swag.IsZero(m.Stops[i]) { // not required
			continue
		}

		if m.Stops[i] != nil {
			if err := m.Stops[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Stops" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Stops" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v3 wrapper bus v wrapper p t x service d t o bus specification v3 stop based on the context it is used
func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStops(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop) contextValidateStops(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Stops); i++ {

		if m.Stops[i] != nil {
			if err := m.Stops[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Stops" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Stops" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3Stop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
