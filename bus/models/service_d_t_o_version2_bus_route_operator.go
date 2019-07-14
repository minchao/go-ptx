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

// ServiceDTOVersion2BusRouteOperator RouteOperator
//
// 營運業者資訊
// swagger:model Service.DTO.Version2.Bus.RouteOperator
type ServiceDTOVersion2BusRouteOperator struct {

	// 營運業者簡碼
	// Required: true
	OperatorCode *string `json:"OperatorCode"`

	// 營運業者代碼
	// Required: true
	OperatorID *string `json:"OperatorID"`

	// NameType
	//
	// 營運業者名稱
	// Required: true
	OperatorName *ServiceDTOVersion2BaseNameType `json:"OperatorName"`

	// 營運業者編號[交通部票證資料系統定義]
	// Required: true
	OperatorNo *string `json:"OperatorNo"`
}

// Validate validates this service d t o version2 bus route operator
func (m *ServiceDTOVersion2BusRouteOperator) Validate(formats strfmt.Registry) error {
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

func (m *ServiceDTOVersion2BusRouteOperator) validateOperatorCode(formats strfmt.Registry) error {

	if err := validate.Required("OperatorCode", "body", m.OperatorCode); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BusRouteOperator) validateOperatorID(formats strfmt.Registry) error {

	if err := validate.Required("OperatorID", "body", m.OperatorID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BusRouteOperator) validateOperatorName(formats strfmt.Registry) error {

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

func (m *ServiceDTOVersion2BusRouteOperator) validateOperatorNo(formats strfmt.Registry) error {

	if err := validate.Required("OperatorNo", "body", m.OperatorNo); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2BusRouteOperator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2BusRouteOperator) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2BusRouteOperator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
