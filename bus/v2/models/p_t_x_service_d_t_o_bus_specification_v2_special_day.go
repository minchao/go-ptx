// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTOBusSpecificationV2SpecialDay SpecialDay
//
// 特殊營運日資料
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.SpecialDay
type PTXServiceDTOBusSpecificationV2SpecialDay struct {

	// DatePeriod
	//
	// 連續特殊日期
	DatePeriod struct {
		PTXServiceDTOBusSpecificationV2DatePeriod
	} `json:"DatePeriod,omitempty" xml:"DatePeriod,omitempty"`

	// Array
	//
	// 不連續特殊日期
	Dates []string "json:\"Dates\" xml:\"List`1\""

	// String
	//
	// 特殊營運描述
	Description string `json:"Description,omitempty" xml:"Description,omitempty"`

	// Int32
	//
	// 營運服務狀態代碼 : [0:'正常營運',1:'加班營運',2:'取消/停駛營運']
	ServiceStatus int64 `json:"ServiceStatus,omitempty"`
}

// Validate validates this p t x service d t o bus specification v2 special day
func (m *PTXServiceDTOBusSpecificationV2SpecialDay) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatePeriod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2SpecialDay) validateDatePeriod(formats strfmt.Registry) error {
	if swag.IsZero(m.DatePeriod) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v2 special day based on the context it is used
func (m *PTXServiceDTOBusSpecificationV2SpecialDay) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatePeriod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2SpecialDay) contextValidateDatePeriod(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2SpecialDay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2SpecialDay) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2SpecialDay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
