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

// ServiceDTOVersion2RailMetroSubClassOperationTime OperationTime
//
// 營運時間資訊
// swagger:model Service.DTO.Version2.Rail.Metro.SubClass.OperationTime
type ServiceDTOVersion2RailMetroSubClassOperationTime struct {

	// 營運結束時間
	// Required: true
	EndTime *string `json:"EndTime"`

	// 營運開始時間
	// Required: true
	StartTime *string `json:"StartTime"`
}

// Validate validates this service d t o version2 rail metro sub class operation time
func (m *ServiceDTOVersion2RailMetroSubClassOperationTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
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

func (m *ServiceDTOVersion2RailMetroSubClassOperationTime) validateEndTime(formats strfmt.Registry) error {

	if err := validate.Required("EndTime", "body", m.EndTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroSubClassOperationTime) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("StartTime", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroSubClassOperationTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroSubClassOperationTime) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailMetroSubClassOperationTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
