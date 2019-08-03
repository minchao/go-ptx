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

// ServiceDTOVersion2BusBusVehicleInfo BusVehicleInfo
//
// 公車車輛基本資料
// swagger:model Service.DTO.Version2.Bus.BusVehicleInfo
type ServiceDTOVersion2BusBusVehicleInfo struct {

	// 車牌號碼
	// Required: true
	PlateNumb *string `json:"PlateNumb"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// integer
	//
	// 車輛種類 : [0:'一般',1:'低地板',2:'復康巴士',3:'小型巴士']
	// Required: true
	VehicleType *int32 `json:"VehicleType"`
}

// Validate validates this service d t o version2 bus bus vehicle info
func (m *ServiceDTOVersion2BusBusVehicleInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlateNumb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVehicleType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion2BusBusVehicleInfo) validatePlateNumb(formats strfmt.Registry) error {

	if err := validate.Required("PlateNumb", "body", m.PlateNumb); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BusBusVehicleInfo) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BusBusVehicleInfo) validateVehicleType(formats strfmt.Registry) error {

	if err := validate.Required("VehicleType", "body", m.VehicleType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2BusBusVehicleInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2BusBusVehicleInfo) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2BusBusVehicleInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
