// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTOBusSpecificationV2ScopeOperator Operator
//
// 營運業者資料
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.Scope+Operator
type PTXServiceDTOBusSpecificationV2ScopeOperator struct {

	// String
	//
	// 營運業者代碼 ,
	OperatorID string `json:"OperatorID,omitempty" xml:"OperatorID,omitempty"`

	// NameType
	//
	// 營運業者名稱
	OperatorName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"OperatorName,omitempty" xml:"NameType,omitempty"`
}

// Validate validates this p t x service d t o bus specification v2 scope operator
func (m *PTXServiceDTOBusSpecificationV2ScopeOperator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperatorName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2ScopeOperator) validateOperatorName(formats strfmt.Registry) error {
	if swag.IsZero(m.OperatorName) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v2 scope operator based on the context it is used
func (m *PTXServiceDTOBusSpecificationV2ScopeOperator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperatorName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2ScopeOperator) contextValidateOperatorName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2ScopeOperator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2ScopeOperator) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2ScopeOperator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
