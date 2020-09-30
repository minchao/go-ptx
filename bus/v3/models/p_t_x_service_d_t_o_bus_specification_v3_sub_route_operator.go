// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTOBusSpecificationV3SubRouteOperator Operator
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.SubRoute+Operator
type PTXServiceDTOBusSpecificationV3SubRouteOperator struct {

	// String
	//
	// 平台代碼
	// Required: true
	OperatorCode *string `json:"OperatorCode"`

	// String
	//
	// 營運業者代碼
	// Required: true
	OperatorID *string `json:"OperatorID"`

	// NameType
	//
	// 營運業者代碼
	// Required: true
	OperatorName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"OperatorName"`

	// String
	//
	// 營運業者編號[交通部票證資料系統定義]
	// Required: true
	OperatorNo *string `json:"OperatorNo"`
}

// Validate validates this p t x service d t o bus specification v3 sub route operator
func (m *PTXServiceDTOBusSpecificationV3SubRouteOperator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperatorCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorNo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3SubRouteOperator) validateOperatorCode(formats strfmt.Registry) error {

	if err := validate.Required("OperatorCode", "body", m.OperatorCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3SubRouteOperator) validateOperatorID(formats strfmt.Registry) error {

	if err := validate.Required("OperatorID", "body", m.OperatorID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3SubRouteOperator) validateOperatorName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3SubRouteOperator) validateOperatorNo(formats strfmt.Registry) error {

	if err := validate.Required("OperatorNo", "body", m.OperatorNo); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3SubRouteOperator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3SubRouteOperator) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3SubRouteOperator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}