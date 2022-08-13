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

// PTXServiceDTOBusSpecificationV2EmbeddedRouteOperator RouteOperator
//
// 營運業者資訊
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.Embedded.RouteOperator
type PTXServiceDTOBusSpecificationV2EmbeddedRouteOperator struct {

	// String
	//
	// 營運業者簡碼
	// Required: true
	OperatorCode *string `json:"OperatorCode" xml:"OperatorCode"`

	// String
	//
	// 營運業者代碼
	// Required: true
	OperatorID *string `json:"OperatorID" xml:"OperatorID"`

	// NameType
	//
	// 營運業者名稱
	// Required: true
	OperatorName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"OperatorName" xml:"NameType"`

	// String
	//
	// 營運業者編號[交通部票證資料系統定義]
	// Required: true
	OperatorNo *string `json:"OperatorNo" xml:"OperatorNo"`
}

// Validate validates this p t x service d t o bus specification v2 embedded route operator
func (m *PTXServiceDTOBusSpecificationV2EmbeddedRouteOperator) Validate(formats strfmt.Registry) error {
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

func (m *PTXServiceDTOBusSpecificationV2EmbeddedRouteOperator) validateOperatorCode(formats strfmt.Registry) error {

	if err := validate.Required("OperatorCode", "body", m.OperatorCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2EmbeddedRouteOperator) validateOperatorID(formats strfmt.Registry) error {

	if err := validate.Required("OperatorID", "body", m.OperatorID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2EmbeddedRouteOperator) validateOperatorName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2EmbeddedRouteOperator) validateOperatorNo(formats strfmt.Registry) error {

	if err := validate.Required("OperatorNo", "body", m.OperatorNo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v2 embedded route operator based on the context it is used
func (m *PTXServiceDTOBusSpecificationV2EmbeddedRouteOperator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperatorName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2EmbeddedRouteOperator) contextValidateOperatorName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2EmbeddedRouteOperator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2EmbeddedRouteOperator) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2EmbeddedRouteOperator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
