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

// PTXServiceDTOBusSpecificationV2BusDiscountPeriods BusDiscountPeriods
//
// 優惠時段
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.BusDiscountPeriods
type PTXServiceDTOBusSpecificationV2BusDiscountPeriods struct {

	// String
	//
	// 結束時間(HH:mm制)
	// Required: true
	EndTime *string `json:"EndTime"`

	// ServiceDay
	//
	// 營運日型態
	// Required: true
	ServiceDay struct {
		PTXServiceDTOBusSpecificationV2EmbeddedServiceDay
	} `json:"ServiceDay"`

	// String
	//
	// 開始時間(HH:mm制)
	// Required: true
	StartTime *string `json:"StartTime"`
}

// Validate validates this p t x service d t o bus specification v2 bus discount periods
func (m *PTXServiceDTOBusSpecificationV2BusDiscountPeriods) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceDay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDiscountPeriods) validateEndTime(formats strfmt.Registry) error {

	if err := validate.Required("EndTime", "body", m.EndTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDiscountPeriods) validateServiceDay(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDiscountPeriods) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("StartTime", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v2 bus discount periods based on the context it is used
func (m *PTXServiceDTOBusSpecificationV2BusDiscountPeriods) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceDay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDiscountPeriods) contextValidateServiceDay(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusDiscountPeriods) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusDiscountPeriods) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2BusDiscountPeriods
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
