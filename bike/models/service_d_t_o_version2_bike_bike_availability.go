// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion2BikeBikeAvailability BikeAvailability
//
// 自行車站點剩餘數量資訊
// swagger:model Service.DTO.Version2.Bike.BikeAvailability
type ServiceDTOVersion2BikeBikeAvailability struct {

	// 可租借車數
	AvailableRentBikes int32 `json:"AvailableRentBikes,omitempty"`

	// 可歸還車數
	AvailableReturnBikes int32 `json:"AvailableReturnBikes,omitempty"`

	// 服務狀態
	// Enum: [0: 停止營運 1: 正常營運]
	ServieAvailable string `json:"ServieAvailable,omitempty"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// 站點代碼
	StationID string `json:"StationID,omitempty"`

	// 站點唯一識別代碼，規則為 {業管機關代碼} + {StationID}，其中 {業管機關代碼} 可於Authority API中的AuthorityCode欄位查詢
	StationUID string `json:"StationUID,omitempty"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this service d t o version2 bike bike availability
func (m *ServiceDTOVersion2BikeBikeAvailability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServieAvailable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serviceDTOVersion2BikeBikeAvailabilityTypeServieAvailablePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["0: 停止營運","1: 正常營運"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceDTOVersion2BikeBikeAvailabilityTypeServieAvailablePropEnum = append(serviceDTOVersion2BikeBikeAvailabilityTypeServieAvailablePropEnum, v)
	}
}

const (

	// ServiceDTOVersion2BikeBikeAvailabilityServieAvailableNr0停止營運 captures enum value "0: 停止營運"
	ServiceDTOVersion2BikeBikeAvailabilityServieAvailableNr0停止營運 string = "0: 停止營運"

	// ServiceDTOVersion2BikeBikeAvailabilityServieAvailableNr1正常營運 captures enum value "1: 正常營運"
	ServiceDTOVersion2BikeBikeAvailabilityServieAvailableNr1正常營運 string = "1: 正常營運"
)

// prop value enum
func (m *ServiceDTOVersion2BikeBikeAvailability) validateServieAvailableEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serviceDTOVersion2BikeBikeAvailabilityTypeServieAvailablePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceDTOVersion2BikeBikeAvailability) validateServieAvailable(formats strfmt.Registry) error {

	if swag.IsZero(m.ServieAvailable) { // not required
		return nil
	}

	// value enum
	if err := m.validateServieAvailableEnum("ServieAvailable", "body", m.ServieAvailable); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BikeBikeAvailability) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BikeBikeAvailability) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2BikeBikeAvailability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2BikeBikeAvailability) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2BikeBikeAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
