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

// PTXServiceDTORailSpecificationV2TRALineStation LineStation
//
// 路線車站資訊
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.TRA.LineStation
type PTXServiceDTORailSpecificationV2TRALineStation struct {

	// Int32
	//
	// 站序
	// Required: true
	Sequence *int32 `json:"Sequence"`

	// String
	//
	// 車站代碼
	// Required: true
	StationID *string `json:"StationID" xml:"StationID"`

	// String
	//
	// 車站名稱
	// Required: true
	StationName *string `json:"StationName" xml:"StationName"`

	// Single
	//
	// 已累積之里程距離(公里)
	// Required: true
	TraveledDistance *float32 `json:"TraveledDistance"`
}

// Validate validates this p t x service d t o rail specification v2 t r a line station
func (m *PTXServiceDTORailSpecificationV2TRALineStation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSequence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraveledDistance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRALineStation) validateSequence(formats strfmt.Registry) error {

	if err := validate.Required("Sequence", "body", m.Sequence); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRALineStation) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRALineStation) validateStationName(formats strfmt.Registry) error {

	if err := validate.Required("StationName", "body", m.StationName); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRALineStation) validateTraveledDistance(formats strfmt.Registry) error {

	if err := validate.Required("TraveledDistance", "body", m.TraveledDistance); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o rail specification v2 t r a line station based on context it is used
func (m *PTXServiceDTORailSpecificationV2TRALineStation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRALineStation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRALineStation) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2TRALineStation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
