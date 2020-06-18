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

// ServiceDTOVersion3RailTRAStationTransferRailTransfer RailTransfer
//
// 軌道轉乘資訊
//
// swagger:model Service.DTO.Version3.Rail.TRA.StationTransfer.RailTransfer
type ServiceDTOVersion3RailTRAStationTransferRailTransfer struct {

	// 轉乘描述
	Description string `json:"Description,omitempty"`

	// 轉乘樓層
	FloorLevel string `json:"FloorLevel,omitempty"`

	// 是否為站內或站外轉乘
	IsOnSiteTransfer bool `json:"IsOnSiteTransfer,omitempty"`

	// 最短轉乘時間
	MinTransferTime int32 `json:"MinTransferTime,omitempty"`

	// 運具種類代碼
	// Required: true
	Mode *string `json:"Mode"`

	// 軌道營運業者簡碼
	// Required: true
	OperatorCode *string `json:"OperatorCode"`

	// NameType
	//
	// 軌道營運業者名稱
	// Required: true
	OperatorName *ServiceDTOVersion3BaseNameType `json:"OperatorName"`

	// 軌道場站代碼
	// Required: true
	StationID *string `json:"StationID"`

	// NameType
	//
	// 軌道場站名稱
	// Required: true
	StationName *ServiceDTOVersion3BaseNameType `json:"StationName"`
}

// Validate validates this service d t o version3 rail t r a station transfer rail transfer
func (m *ServiceDTOVersion3RailTRAStationTransferRailTransfer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3RailTRAStationTransferRailTransfer) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("Mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRAStationTransferRailTransfer) validateOperatorCode(formats strfmt.Registry) error {

	if err := validate.Required("OperatorCode", "body", m.OperatorCode); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRAStationTransferRailTransfer) validateOperatorName(formats strfmt.Registry) error {

	if err := validate.Required("OperatorName", "body", m.OperatorName); err != nil {
		return err
	}

	if m.OperatorName != nil {
		if err := m.OperatorName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OperatorName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRAStationTransferRailTransfer) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRAStationTransferRailTransfer) validateStationName(formats strfmt.Registry) error {

	if err := validate.Required("StationName", "body", m.StationName); err != nil {
		return err
	}

	if m.StationName != nil {
		if err := m.StationName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StationName")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRAStationTransferRailTransfer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRAStationTransferRailTransfer) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3RailTRAStationTransferRailTransfer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
