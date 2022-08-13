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

// PTXServiceDTORailSpecificationV3TRAOperator Operator
//
// 營運業者基本資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.Operator
type PTXServiceDTORailSpecificationV3TRAOperator struct {

	// String
	//
	// 營運業者票價查詢網站連結
	FareURL string `json:"FareURL,omitempty" xml:"FareURL,omitempty"`

	// String
	//
	// 營運業者簡碼
	// Required: true
	OperatorCode *string `json:"OperatorCode" xml:"OperatorCode"`

	// String
	//
	// 營運業者電子信箱
	OperatorEmail string `json:"OperatorEmail,omitempty" xml:"OperatorEmail,omitempty"`

	// String
	//
	// 營運業者Logo網址
	OperatorLogoURL string `json:"OperatorLogoURL,omitempty" xml:"OperatorLogoURL,omitempty"`

	// NameType
	//
	// 營運業者名稱
	// Required: true
	OperatorName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"OperatorName" xml:"NameType"`

	// String
	//
	// 營運業者連絡電話
	OperatorPhone string `json:"OperatorPhone,omitempty" xml:"OperatorPhone,omitempty"`

	// String
	//
	// 營運業者網址連結
	OperatorURL string `json:"OperatorURL,omitempty" xml:"OperatorURL,omitempty"`

	// String
	//
	// 營運業者訂票電話
	ReservationPhone string `json:"ReservationPhone,omitempty" xml:"ReservationPhone,omitempty"`

	// String
	//
	// 營運業者訂票網站
	ReservationURL string `json:"ReservationURL,omitempty" xml:"ReservationURL,omitempty"`
}

// Validate validates this p t x service d t o rail specification v3 t r a operator
func (m *PTXServiceDTORailSpecificationV3TRAOperator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperatorCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAOperator) validateOperatorCode(formats strfmt.Registry) error {

	if err := validate.Required("OperatorCode", "body", m.OperatorCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAOperator) validateOperatorName(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a operator based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRAOperator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperatorName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAOperator) contextValidateOperatorName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAOperator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAOperator) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRAOperator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
