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

// PTXServiceDTOSharedSpecificationV3BaseAuthority Authority
//
// 業管機關
//
// swagger:model PTX.Service.DTO.Shared.Specification.V3.Base.Authority
type PTXServiceDTOSharedSpecificationV3BaseAuthority struct {

	// String
	//
	// 業管機關地址
	// Required: true
	AuthorityAddress *string `json:"AuthorityAddress" xml:"String"`

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode" xml:"String"`

	// String
	//
	// 業管機關電子信箱
	// Required: true
	AuthorityEmail *string `json:"AuthorityEmail" xml:"String"`

	// NameType
	//
	// 業管機關名稱
	// Required: true
	AuthorityName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"AuthorityName" xml:"NameType"`

	// String
	//
	// 業管機關識別代碼(可參閱: https://oid.nat.gov.tw/OIDWeb/)
	// Required: true
	AuthorityOID *string `json:"AuthorityOID" xml:"String"`

	// String
	//
	// 業管機關連絡電話
	// Required: true
	AuthorityPhone *string `json:"AuthorityPhone" xml:"String"`

	// String
	//
	// 業管機關官網網址
	AuthorityURL string `json:"AuthorityUrl,omitempty" xml:"String,omitempty"`

	// String
	//
	// 業管機關Logo網址
	LogoURL string `json:"LogoURL,omitempty" xml:"String,omitempty"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o shared specification v3 base authority
func (m *PTXServiceDTOSharedSpecificationV3BaseAuthority) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorityEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorityName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorityOID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorityPhone(formats); err != nil {
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

func (m *PTXServiceDTOSharedSpecificationV3BaseAuthority) validateAuthorityAddress(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityAddress", "body", m.AuthorityAddress); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOSharedSpecificationV3BaseAuthority) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOSharedSpecificationV3BaseAuthority) validateAuthorityEmail(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityEmail", "body", m.AuthorityEmail); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOSharedSpecificationV3BaseAuthority) validateAuthorityName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOSharedSpecificationV3BaseAuthority) validateAuthorityOID(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityOID", "body", m.AuthorityOID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOSharedSpecificationV3BaseAuthority) validateAuthorityPhone(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityPhone", "body", m.AuthorityPhone); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOSharedSpecificationV3BaseAuthority) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o shared specification v3 base authority based on the context it is used
func (m *PTXServiceDTOSharedSpecificationV3BaseAuthority) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthorityName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOSharedSpecificationV3BaseAuthority) contextValidateAuthorityName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOSharedSpecificationV3BaseAuthority) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOSharedSpecificationV3BaseAuthority) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOSharedSpecificationV3BaseAuthority
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
