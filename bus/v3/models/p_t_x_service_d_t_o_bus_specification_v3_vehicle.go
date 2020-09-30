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

// PTXServiceDTOBusSpecificationV3Vehicle Vehicle
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.Vehicle
type PTXServiceDTOBusSpecificationV3Vehicle struct {

	// Int32
	//
	// 讀卡機配置(1:前門刷卡;2:前後門刷卡)
	// Required: true
	CardReaderLayout *int32 `json:"CardReaderLayout"`

	// Boolean
	//
	// 是否有升降或斜坡板設備
	// Required: true
	HasLiftOrRamp *bool `json:"HasLiftOrRamp"`

	// Boolean
	//
	// 是否有提供Wifi服務
	// Required: true
	HasWifi *bool `json:"HasWifi"`

	// String
	//
	// 車機代號
	InBoxID string `json:"InBoxID,omitempty"`

	// Boolean
	//
	// 是否為電動公車
	// Required: true
	IsElectric *bool `json:"IsElectric"`

	// Boolean
	//
	// 是否為油電混合公車
	// Required: true
	IsHybrid *bool `json:"IsHybrid"`

	// Boolean
	//
	// 是否為低地板
	// Required: true
	IsLowFloor *bool `json:"IsLowFloor"`

	// String
	//
	// 營運業者簡碼
	// Required: true
	OperatorCode *string `json:"OperatorCode"`

	// String
	//
	// 營運業者代碼
	// Required: true
	OperatorID *string `json:"OperatorID"`

	// String
	//
	// 車牌號碼
	// Required: true
	PlateNumb *string `json:"PlateNumb"`

	// DateTime
	//
	// 購入時間
	PurchaseTime string `json:"PurchaseTime,omitempty"`

	// integer
	//
	// 車輛型別 : [1:'大型巴士',2:'中型巴士',3:'小型巴士',4:'雙層巴士',5:'雙節巴士',6:'計程車']
	// Required: true
	VehicleClass *int32 `json:"VehicleClass"`

	// integer
	//
	// 車輛種類 : [1:'一般',2:'復康巴士',3:'專車',4:'其他']
	// Required: true
	VehicleType *int32 `json:"VehicleType"`
}

// Validate validates this p t x service d t o bus specification v3 vehicle
func (m *PTXServiceDTOBusSpecificationV3Vehicle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCardReaderLayout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasLiftOrRamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasWifi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsElectric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsHybrid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsLowFloor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlateNumb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVehicleClass(formats); err != nil {
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

func (m *PTXServiceDTOBusSpecificationV3Vehicle) validateCardReaderLayout(formats strfmt.Registry) error {

	if err := validate.Required("CardReaderLayout", "body", m.CardReaderLayout); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Vehicle) validateHasLiftOrRamp(formats strfmt.Registry) error {

	if err := validate.Required("HasLiftOrRamp", "body", m.HasLiftOrRamp); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Vehicle) validateHasWifi(formats strfmt.Registry) error {

	if err := validate.Required("HasWifi", "body", m.HasWifi); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Vehicle) validateIsElectric(formats strfmt.Registry) error {

	if err := validate.Required("IsElectric", "body", m.IsElectric); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Vehicle) validateIsHybrid(formats strfmt.Registry) error {

	if err := validate.Required("IsHybrid", "body", m.IsHybrid); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Vehicle) validateIsLowFloor(formats strfmt.Registry) error {

	if err := validate.Required("IsLowFloor", "body", m.IsLowFloor); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Vehicle) validateOperatorCode(formats strfmt.Registry) error {

	if err := validate.Required("OperatorCode", "body", m.OperatorCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Vehicle) validateOperatorID(formats strfmt.Registry) error {

	if err := validate.Required("OperatorID", "body", m.OperatorID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Vehicle) validatePlateNumb(formats strfmt.Registry) error {

	if err := validate.Required("PlateNumb", "body", m.PlateNumb); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Vehicle) validateVehicleClass(formats strfmt.Registry) error {

	if err := validate.Required("VehicleClass", "body", m.VehicleClass); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Vehicle) validateVehicleType(formats strfmt.Registry) error {

	if err := validate.Required("VehicleType", "body", m.VehicleType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3Vehicle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3Vehicle) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3Vehicle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}