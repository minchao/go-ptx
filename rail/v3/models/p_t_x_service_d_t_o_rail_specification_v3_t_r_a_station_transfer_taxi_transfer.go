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

// PTXServiceDTORailSpecificationV3TRAStationTransferTaxiTransfer TaxiTransfer
//
// 計程車轉乘資訊
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.StationTransfer.TaxiTransfer
type PTXServiceDTORailSpecificationV3TRAStationTransferTaxiTransfer struct {

	// String
	//
	// 轉乘描述
	Description string `json:"Description,omitempty" xml:"Description,omitempty"`

	// String
	//
	// 轉乘樓層
	FloorLevel string `json:"FloorLevel,omitempty" xml:"FloorLevel,omitempty"`

	// 是否為站內或站外轉乘
	IsOnSiteTransfer bool `json:"IsOnSiteTransfer,omitempty"`

	// 最短轉乘時間
	MinTransferTime float64 `json:"MinTransferTime,omitempty"`

	// String
	//
	// 運具種類代碼
	// Required: true
	Mode *string `json:"Mode" xml:"Mode"`

	// String
	//
	// 計程車招呼站代碼
	TaxiStopID string `json:"TaxiStopID,omitempty" xml:"TaxiStopID,omitempty"`

	// NameType
	//
	// 計程車招呼站名稱
	// Required: true
	TaxiStopName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"TaxiStopName" xml:"NameType"`
}

// Validate validates this p t x service d t o rail specification v3 t r a station transfer taxi transfer
func (m *PTXServiceDTORailSpecificationV3TRAStationTransferTaxiTransfer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxiStopName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationTransferTaxiTransfer) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("Mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationTransferTaxiTransfer) validateTaxiStopName(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a station transfer taxi transfer based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRAStationTransferTaxiTransfer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTaxiStopName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationTransferTaxiTransfer) contextValidateTaxiStopName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAStationTransferTaxiTransfer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAStationTransferTaxiTransfer) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRAStationTransferTaxiTransfer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
