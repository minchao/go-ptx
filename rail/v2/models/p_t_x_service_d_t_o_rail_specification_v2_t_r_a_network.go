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

// PTXServiceDTORailSpecificationV2TRANetwork Network
//
// 臺鐵路網資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.TRA.Network
type PTXServiceDTORailSpecificationV2TRANetwork struct {

	// Array
	//
	// 臺鐵路線資訊
	// Required: true
	Lines []*PTXServiceDTORailSpecificationV2TRALineSimple "json:\"Lines\" xml:\"List`1\""

	// String
	//
	// 臺鐵路網代碼
	// Required: true
	NetworkID *string `json:"NetworkID" xml:"String"`

	// String
	//
	// 臺鐵路網圖網址URL
	// Required: true
	NetworkMapURL *string `json:"NetworkMapUrl" xml:"String"`

	// String
	//
	// 臺鐵路網英文名稱
	// Required: true
	NetworkNameEn *string `json:"NetworkNameEn" xml:"String"`

	// String
	//
	// 臺鐵路網中文名稱
	// Required: true
	NetworkNameZh *string `json:"NetworkNameZh" xml:"String"`

	// String
	//
	// 營運業者代碼
	// Required: true
	OperatorID *string `json:"OperatorID" xml:"String"`

	// String
	//
	// 臺鐵路網英文名稱
	// Required: true
	OperatorNameEn *string `json:"OperatorNameEn" xml:"String"`

	// String
	//
	// 臺鐵路網中文名稱
	// Required: true
	OperatorNameZh *string `json:"OperatorNameZh" xml:"String"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o rail specification v2 t r a network
func (m *PTXServiceDTORailSpecificationV2TRANetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkMapURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkNameEn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkNameZh(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorNameEn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorNameZh(formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV2TRANetwork) validateLines(formats strfmt.Registry) error {

	if err := validate.Required("Lines", "body", m.Lines); err != nil {
		return err
	}

	for i := 0; i < len(m.Lines); i++ {
		if swag.IsZero(m.Lines[i]) { // not required
			continue
		}

		if m.Lines[i] != nil {
			if err := m.Lines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Lines" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Lines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRANetwork) validateNetworkID(formats strfmt.Registry) error {

	if err := validate.Required("NetworkID", "body", m.NetworkID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRANetwork) validateNetworkMapURL(formats strfmt.Registry) error {

	if err := validate.Required("NetworkMapUrl", "body", m.NetworkMapURL); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRANetwork) validateNetworkNameEn(formats strfmt.Registry) error {

	if err := validate.Required("NetworkNameEn", "body", m.NetworkNameEn); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRANetwork) validateNetworkNameZh(formats strfmt.Registry) error {

	if err := validate.Required("NetworkNameZh", "body", m.NetworkNameZh); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRANetwork) validateOperatorID(formats strfmt.Registry) error {

	if err := validate.Required("OperatorID", "body", m.OperatorID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRANetwork) validateOperatorNameEn(formats strfmt.Registry) error {

	if err := validate.Required("OperatorNameEn", "body", m.OperatorNameEn); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRANetwork) validateOperatorNameZh(formats strfmt.Registry) error {

	if err := validate.Required("OperatorNameZh", "body", m.OperatorNameZh); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRANetwork) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 t r a network based on the context it is used
func (m *PTXServiceDTORailSpecificationV2TRANetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLines(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRANetwork) contextValidateLines(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Lines); i++ {

		if m.Lines[i] != nil {
			if err := m.Lines[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Lines" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Lines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRANetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRANetwork) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2TRANetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
