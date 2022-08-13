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

// PTXServiceDTORailSpecificationV3TRANetworkNetwork Network
//
// 路網資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.Network.Network
type PTXServiceDTORailSpecificationV3TRANetworkNetwork struct {

	// Array
	//
	// 路線資訊
	// Required: true
	Lines []*PTXServiceDTORailSpecificationV3TRANetworkLine "json:\"Lines\" xml:\"List`1\""

	// String
	//
	// 路網代碼
	// Required: true
	NetworkID *string `json:"NetworkID" xml:"NetworkID"`

	// MapNameType
	//
	// 路網圖網址URL
	// Required: true
	NetworkMapURL struct {
		PTXServiceDTORailSpecificationV3TRANetworkMapNameType
	} `json:"NetworkMapURL" xml:"MapNameType"`

	// NameType
	//
	// 路網名稱
	// Required: true
	NetworkName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"NetworkName" xml:"NameType"`

	// String
	//
	// 營運業者代碼
	// Required: true
	OperatorCode *string `json:"OperatorCode" xml:"OperatorCode"`

	// NameType
	//
	// 營運業者名稱
	// Required: true
	OperatorName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"OperatorName" xml:"NameType"`
}

// Validate validates this p t x service d t o rail specification v3 t r a network network
func (m *PTXServiceDTORailSpecificationV3TRANetworkNetwork) Validate(formats strfmt.Registry) error {
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

	if err := m.validateNetworkName(formats); err != nil {
		res = append(res, err)
	}

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

func (m *PTXServiceDTORailSpecificationV3TRANetworkNetwork) validateLines(formats strfmt.Registry) error {

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

func (m *PTXServiceDTORailSpecificationV3TRANetworkNetwork) validateNetworkID(formats strfmt.Registry) error {

	if err := validate.Required("NetworkID", "body", m.NetworkID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRANetworkNetwork) validateNetworkMapURL(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRANetworkNetwork) validateNetworkName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRANetworkNetwork) validateOperatorCode(formats strfmt.Registry) error {

	if err := validate.Required("OperatorCode", "body", m.OperatorCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRANetworkNetwork) validateOperatorName(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a network network based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRANetworkNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLines(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkMapURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperatorName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRANetworkNetwork) contextValidateLines(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PTXServiceDTORailSpecificationV3TRANetworkNetwork) contextValidateNetworkMapURL(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRANetworkNetwork) contextValidateNetworkName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRANetworkNetwork) contextValidateOperatorName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRANetworkNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRANetworkNetwork) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRANetworkNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
