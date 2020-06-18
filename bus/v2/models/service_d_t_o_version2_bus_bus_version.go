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

// ServiceDTOVersion2BusBusVersion BusVersion
//
// 公車版本資料型別
//
// swagger:model Service.DTO.Version2.Bus.BusVersion
type ServiceDTOVersion2BusBusVersion struct {

	// DateTime
	//
	// 此資料版本最後檢查更新之日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateCheckTime *string `json:"UpdateCheckTime"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this service d t o version2 bus bus version
func (m *ServiceDTOVersion2BusBusVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpdateCheckTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion2BusBusVersion) validateUpdateCheckTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateCheckTime", "body", m.UpdateCheckTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BusBusVersion) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BusBusVersion) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2BusBusVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2BusBusVersion) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2BusBusVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
