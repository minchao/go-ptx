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

// PTXServiceDTOShipSpecificationV3Stoptime Stoptime
//
// swagger:model PTX.Service.DTO.Ship.Specification.V3.Stoptime
type PTXServiceDTOShipSpecificationV3Stoptime struct {

	// String
	//
	// 港口進港時間
	// Required: true
	ArrivalTime *string `json:"ArrivalTime" xml:"String"`

	// String
	//
	// 港口離港時間
	// Required: true
	DepartureTime *string `json:"DepartureTime" xml:"String"`

	// String
	//
	// 停靠港口代碼
	// Required: true
	PortID *string `json:"PortID" xml:"String"`

	// NameType
	//
	// 港口名稱
	// Required: true
	PortName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"PortName" xml:"NameType"`

	// Int32
	//
	// 航線停靠站序
	// Required: true
	StopSequence *int32 `json:"StopSequence"`

	// Int32
	//
	// 預估航行時間
	// Required: true
	TravelTime *int32 `json:"TravelTime"`
}

// Validate validates this p t x service d t o ship specification v3 stoptime
func (m *PTXServiceDTOShipSpecificationV3Stoptime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArrivalTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartureTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopSequence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTravelTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOShipSpecificationV3Stoptime) validateArrivalTime(formats strfmt.Registry) error {

	if err := validate.Required("ArrivalTime", "body", m.ArrivalTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3Stoptime) validateDepartureTime(formats strfmt.Registry) error {

	if err := validate.Required("DepartureTime", "body", m.DepartureTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3Stoptime) validatePortID(formats strfmt.Registry) error {

	if err := validate.Required("PortID", "body", m.PortID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3Stoptime) validatePortName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3Stoptime) validateStopSequence(formats strfmt.Registry) error {

	if err := validate.Required("StopSequence", "body", m.StopSequence); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3Stoptime) validateTravelTime(formats strfmt.Registry) error {

	if err := validate.Required("TravelTime", "body", m.TravelTime); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o ship specification v3 stoptime based on the context it is used
func (m *PTXServiceDTOShipSpecificationV3Stoptime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePortName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOShipSpecificationV3Stoptime) contextValidatePortName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3Stoptime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3Stoptime) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOShipSpecificationV3Stoptime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
