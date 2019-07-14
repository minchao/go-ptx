// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion3RailTRAStationTransferAirportTransfer AirportTransfer
//
// 航空運具轉乘資訊
// swagger:model Service.DTO.Version3.Rail.TRA.StationTransfer.AirportTransfer
type ServiceDTOVersion3RailTRAStationTransferAirportTransfer struct {

	// 機場代碼
	// Required: true
	AirportID *string `json:"AirportID"`

	// NameType
	//
	// 機場名稱
	// Required: true
	AirportName *ServiceDTOVersion3BaseNameType `json:"AirportName"`

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

	// 機場營運業者簡碼
	// Required: true
	OperatorCode *string `json:"OperatorCode"`

	// NameType
	//
	// 機場營運業者名稱
	// Required: true
	OperatorName *ServiceDTOVersion3BaseNameType `json:"OperatorName"`
}

// Validate validates this service d t o version3 rail t r a station transfer airport transfer
func (m *ServiceDTOVersion3RailTRAStationTransferAirportTransfer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAirportID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAirportName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
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

func (m *ServiceDTOVersion3RailTRAStationTransferAirportTransfer) validateAirportID(formats strfmt.Registry) error {

	if err := validate.Required("AirportID", "body", m.AirportID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRAStationTransferAirportTransfer) validateAirportName(formats strfmt.Registry) error {

	if err := validate.Required("AirportName", "body", m.AirportName); err != nil {
		return err
	}

	if m.AirportName != nil {
		if err := m.AirportName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AirportName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRAStationTransferAirportTransfer) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("Mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRAStationTransferAirportTransfer) validateOperatorCode(formats strfmt.Registry) error {

	if err := validate.Required("OperatorCode", "body", m.OperatorCode); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRAStationTransferAirportTransfer) validateOperatorName(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRAStationTransferAirportTransfer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRAStationTransferAirportTransfer) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3RailTRAStationTransferAirportTransfer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
