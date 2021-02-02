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

// PTXServiceDTORailSpecificationV2THSRRailGeneralTrainInfo RailGeneralTrainInfo
//
// 高鐵車次定期資料型別
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.THSR.RailGeneralTrainInfo
type PTXServiceDTORailSpecificationV2THSRRailGeneralTrainInfo struct {

	// integer
	//
	// 行駛方向 : [0:'南下',1:'北上']
	// Required: true
	Direction *int32 `json:"Direction"`

	// String
	//
	// 列車終點車站代號
	EndingStationID string `json:"EndingStationID,omitempty"`

	// NameType
	//
	// 列車終點車站名稱
	EndingStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"EndingStationName,omitempty"`

	// NameType
	//
	// 附註說明
	Note struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"Note,omitempty"`

	// String
	//
	// 列車起點車站代號
	StartingStationID string `json:"StartingStationID,omitempty"`

	// NameType
	//
	// 列車起點車站名稱
	StartingStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StartingStationName,omitempty"`

	// String
	//
	// 車次代碼
	// Required: true
	TrainNo *string `json:"TrainNo"`
}

// Validate validates this p t x service d t o rail specification v2 t h s r rail general train info
func (m *PTXServiceDTORailSpecificationV2THSRRailGeneralTrainInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndingStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartingStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainNo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailGeneralTrainInfo) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailGeneralTrainInfo) validateEndingStationName(formats strfmt.Registry) error {
	if swag.IsZero(m.EndingStationName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailGeneralTrainInfo) validateNote(formats strfmt.Registry) error {
	if swag.IsZero(m.Note) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailGeneralTrainInfo) validateStartingStationName(formats strfmt.Registry) error {
	if swag.IsZero(m.StartingStationName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailGeneralTrainInfo) validateTrainNo(formats strfmt.Registry) error {

	if err := validate.Required("TrainNo", "body", m.TrainNo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 t h s r rail general train info based on the context it is used
func (m *PTXServiceDTORailSpecificationV2THSRRailGeneralTrainInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndingStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartingStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailGeneralTrainInfo) contextValidateEndingStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailGeneralTrainInfo) contextValidateNote(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailGeneralTrainInfo) contextValidateStartingStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2THSRRailGeneralTrainInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2THSRRailGeneralTrainInfo) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2THSRRailGeneralTrainInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
