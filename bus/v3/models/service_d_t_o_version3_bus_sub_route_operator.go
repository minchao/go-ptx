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

// ServiceDTOVersion3BusSubRouteOperator Operator
//
// swagger:model Service.DTO.Version3.Bus.SubRoute.Operator
type ServiceDTOVersion3BusSubRouteOperator struct {

	// 平台代碼
	// Required: true
	OperatorCode *string `json:"OperatorCode"`

	// 營運業者代碼
	// Required: true
	OperatorID *string `json:"OperatorID"`

	// NameType
	//
	// 營運業者代碼
	// Required: true
	OperatorName *ServiceDTOVersion3BaseNameType `json:"OperatorName"`

	// 營運業者編號[交通部票證資料系統定義]
	// Required: true
	OperatorNo *string `json:"OperatorNo"`
}

// Validate validates this service d t o version3 bus sub route operator
func (m *ServiceDTOVersion3BusSubRouteOperator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperatorCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorNo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3BusSubRouteOperator) validateOperatorCode(formats strfmt.Registry) error {

	if err := validate.Required("OperatorCode", "body", m.OperatorCode); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusSubRouteOperator) validateOperatorID(formats strfmt.Registry) error {

	if err := validate.Required("OperatorID", "body", m.OperatorID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusSubRouteOperator) validateOperatorName(formats strfmt.Registry) error {

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

func (m *ServiceDTOVersion3BusSubRouteOperator) validateOperatorNo(formats strfmt.Registry) error {

	if err := validate.Required("OperatorNo", "body", m.OperatorNo); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3BusSubRouteOperator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3BusSubRouteOperator) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3BusSubRouteOperator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
