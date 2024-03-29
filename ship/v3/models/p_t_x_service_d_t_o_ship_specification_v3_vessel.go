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

// PTXServiceDTOShipSpecificationV3Vessel Vessel
//
// swagger:model PTX.Service.DTO.Ship.Specification.V3.Vessel
type PTXServiceDTOShipSpecificationV3Vessel struct {

	// String
	//
	// 船呼
	CallSign string `json:"CallSign,omitempty" xml:"CallSign,omitempty"`

	// 船舶總噸位
	GrossTonnage float32 `json:"GrossTonnage,omitempty"`

	// String
	//
	// 國際船舶編號
	IMO string `json:"IMO,omitempty" xml:"IMO,omitempty"`

	// 船長
	Length float32 `json:"Length,omitempty"`

	// 吃水深度
	LoadDraft float32 `json:"LoadDraft,omitempty"`

	// String
	//
	// AIS船舶編號
	MMSI string `json:"MMSI,omitempty" xml:"MMSI,omitempty"`

	// 最大載客數
	MaxPassengerCount int32 `json:"MaxPassengerCount,omitempty"`

	// String
	//
	// 船舶所屬國籍
	// Required: true
	Nationality *string `json:"Nationality" xml:"Nationality"`

	// String
	//
	// 備註
	Note string `json:"Note,omitempty" xml:"Note,omitempty"`

	// String
	//
	// 船舶所屬營運業者代碼
	// Required: true
	OperatorID *string `json:"OperatorID" xml:"OperatorID"`

	// String
	//
	// 船舶分類
	VesselClass string `json:"VesselClass,omitempty" xml:"VesselClass,omitempty"`

	// String
	//
	// 船舶代碼
	// Required: true
	VesselID *string `json:"VesselID" xml:"VesselID"`

	// NameType
	//
	// 船舶名稱
	// Required: true
	VesselName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"VesselName" xml:"NameType"`

	// String
	//
	// 台灣船舶號數
	VesselNo string `json:"VesselNo,omitempty" xml:"VesselNo,omitempty"`

	// 船寬
	Width float32 `json:"Width,omitempty"`
}

// Validate validates this p t x service d t o ship specification v3 vessel
func (m *PTXServiceDTOShipSpecificationV3Vessel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNationality(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVesselID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVesselName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOShipSpecificationV3Vessel) validateNationality(formats strfmt.Registry) error {

	if err := validate.Required("Nationality", "body", m.Nationality); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3Vessel) validateOperatorID(formats strfmt.Registry) error {

	if err := validate.Required("OperatorID", "body", m.OperatorID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3Vessel) validateVesselID(formats strfmt.Registry) error {

	if err := validate.Required("VesselID", "body", m.VesselID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3Vessel) validateVesselName(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this p t x service d t o ship specification v3 vessel based on the context it is used
func (m *PTXServiceDTOShipSpecificationV3Vessel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVesselName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOShipSpecificationV3Vessel) contextValidateVesselName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3Vessel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3Vessel) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOShipSpecificationV3Vessel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
