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

// PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute BusVehicleRouteList
//
// 版本 wrapper (具有版本資訊)
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.Wrapper.BusVWrapper[PTX.Service.DTO.Bus.Specification.V3.VehicleRoute]
type PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute struct {

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode" xml:"String"`

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

	// Array
	//
	// 資料列表
	// Required: true
	VehicleRoutes []*PTXServiceDTOBusSpecificationV3VehicleRoute "json:\"VehicleRoutes\" xml:\"List`1\""

	// Int32
	//
	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this p t x service d t o bus specification v3 wrapper bus v wrapper p t x service d t o bus specification v3 vehicle route
func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute) Validate(formats strfmt.Registry) error {
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

	if err := m.validateUpdateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVehicleRoutes(formats); err != nil {
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

func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute) validateVehicleRoutes(formats strfmt.Registry) error {

	if err := validate.Required("VehicleRoutes", "body", m.VehicleRoutes); err != nil {
		return err
	}

	for i := 0; i < len(m.VehicleRoutes); i++ {
		if swag.IsZero(m.VehicleRoutes[i]) { // not required
			continue
		}

		if m.VehicleRoutes[i] != nil {
			if err := m.VehicleRoutes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("VehicleRoutes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v3 wrapper bus v wrapper p t x service d t o bus specification v3 vehicle route based on the context it is used
func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVehicleRoutes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute) contextValidateVehicleRoutes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VehicleRoutes); i++ {

		if m.VehicleRoutes[i] != nil {
			if err := m.VehicleRoutes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("VehicleRoutes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3VehicleRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
