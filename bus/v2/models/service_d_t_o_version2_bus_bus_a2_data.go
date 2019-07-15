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

// ServiceDTOVersion2BusBusA2Data BusA2Data
//
// 定點車機資料型別
// swagger:model Service.DTO.Version2.Bus.BusA2Data
type ServiceDTOVersion2BusBusA2Data struct {

	// 進站離站
	// Enum: [0 1]
	A2EventType int64 `json:"A2EventType,omitempty"`

	// 行車狀況
	// Enum: [0 1 2 3 4 5 90 91 98 99 100 101 255]
	BusStatus int64 `json:"BusStatus,omitempty"`

	// 去返程
	// Required: true
	// Enum: [0 1 2 255]
	Direction *int64 `json:"Direction"`

	// 勤務狀態
	// Enum: [0 1 2]
	DutyStatus int64 `json:"DutyStatus,omitempty"`

	// DateTime
	//
	// 車機時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz) [觸發到離站的GPS時間]
	// Required: true
	GPSTime *string `json:"GPSTime"`

	// 資料型態種類
	// Enum: [0 1 2]
	MessageType int64 `json:"MessageType,omitempty"`

	// 營運業者代碼
	OperatorID string `json:"OperatorID,omitempty"`

	// 車牌號碼
	// Required: true
	PlateNumb *string `json:"PlateNumb"`

	// 地區既用中之路線代碼(為原資料內碼)
	RouteID string `json:"RouteID,omitempty"`

	// NameType
	//
	// 路線名
	RouteName *ServiceDTOVersion2BaseNameType `json:"RouteName,omitempty"`

	// 路線唯一識別代碼，規則為 {業管機關簡碼} + {RouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	RouteUID string `json:"RouteUID,omitempty"`

	// DateTime
	//
	// 來源端平台接收時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	SrcRecTime string `json:"SrcRecTime,omitempty"`

	// DateTime
	//
	// 來源端平台資料傳出時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)[公總使用TCP動態即時推播故有提供此欄位, 而非公總系統因使用整包資料更新, 故沒有提供此欄位]
	SrcTransTime string `json:"SrcTransTime,omitempty"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)[公總使用TCP動態即時推播故沒有提供此欄位, 而非公總系統因提供整包資料更新, 故有提供此欄]
	SrcUpdateTime string `json:"SrcUpdateTime,omitempty"`

	// 地區既用中之站牌代號(為原資料內碼)
	StopID string `json:"StopID,omitempty"`

	// NameType
	//
	// 站牌名
	StopName *ServiceDTOVersion2BaseNameType `json:"StopName,omitempty"`

	// 路線經過站牌之順序
	StopSequence int32 `json:"StopSequence,omitempty"`

	// 站牌唯一識別代碼，規則為 {業管機關簡碼} + {StopID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	StopUID string `json:"StopUID,omitempty"`

	// 地區既用中之子路線代碼(為原資料內碼)
	SubRouteID string `json:"SubRouteID,omitempty"`

	// NameType
	//
	// 子路線名稱
	SubRouteName *ServiceDTOVersion2BaseNameType `json:"SubRouteName,omitempty"`

	// 子路線唯一識別代碼，規則為 {業管機關簡碼} + {SubRouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	SubRouteUID string `json:"SubRouteUID,omitempty"`

	// DateTime
	//
	// 車機資料傳輸時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)[多數單位沒有提供此欄位資訊]
	TransTime string `json:"TransTime,omitempty"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this service d t o version2 bus bus a2 data
func (m *ServiceDTOVersion2BusBusA2Data) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateA2EventType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBusStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDutyStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGPSTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessageType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlateNumb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubRouteName(formats); err != nil {
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

var serviceDTOVersion2BusBusA2DataTypeA2EventTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceDTOVersion2BusBusA2DataTypeA2EventTypePropEnum = append(serviceDTOVersion2BusBusA2DataTypeA2EventTypePropEnum, v)
	}
}

// prop value enum
func (m *ServiceDTOVersion2BusBusA2Data) validateA2EventTypeEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, serviceDTOVersion2BusBusA2DataTypeA2EventTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceDTOVersion2BusBusA2Data) validateA2EventType(formats strfmt.Registry) error {

	if swag.IsZero(m.A2EventType) { // not required
		return nil
	}

	// value enum
	if err := m.validateA2EventTypeEnum("A2EventType", "body", m.A2EventType); err != nil {
		return err
	}

	return nil
}

var serviceDTOVersion2BusBusA2DataTypeBusStatusPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1,2,3,4,5,90,91,98,99,100,101,255]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceDTOVersion2BusBusA2DataTypeBusStatusPropEnum = append(serviceDTOVersion2BusBusA2DataTypeBusStatusPropEnum, v)
	}
}

// prop value enum
func (m *ServiceDTOVersion2BusBusA2Data) validateBusStatusEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, serviceDTOVersion2BusBusA2DataTypeBusStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceDTOVersion2BusBusA2Data) validateBusStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.BusStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateBusStatusEnum("BusStatus", "body", m.BusStatus); err != nil {
		return err
	}

	return nil
}

var serviceDTOVersion2BusBusA2DataTypeDirectionPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1,2,255]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceDTOVersion2BusBusA2DataTypeDirectionPropEnum = append(serviceDTOVersion2BusBusA2DataTypeDirectionPropEnum, v)
	}
}

// prop value enum
func (m *ServiceDTOVersion2BusBusA2Data) validateDirectionEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, serviceDTOVersion2BusBusA2DataTypeDirectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceDTOVersion2BusBusA2Data) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	// value enum
	if err := m.validateDirectionEnum("Direction", "body", *m.Direction); err != nil {
		return err
	}

	return nil
}

var serviceDTOVersion2BusBusA2DataTypeDutyStatusPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceDTOVersion2BusBusA2DataTypeDutyStatusPropEnum = append(serviceDTOVersion2BusBusA2DataTypeDutyStatusPropEnum, v)
	}
}

// prop value enum
func (m *ServiceDTOVersion2BusBusA2Data) validateDutyStatusEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, serviceDTOVersion2BusBusA2DataTypeDutyStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceDTOVersion2BusBusA2Data) validateDutyStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.DutyStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateDutyStatusEnum("DutyStatus", "body", m.DutyStatus); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BusBusA2Data) validateGPSTime(formats strfmt.Registry) error {

	if err := validate.Required("GPSTime", "body", m.GPSTime); err != nil {
		return err
	}

	return nil
}

var serviceDTOVersion2BusBusA2DataTypeMessageTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceDTOVersion2BusBusA2DataTypeMessageTypePropEnum = append(serviceDTOVersion2BusBusA2DataTypeMessageTypePropEnum, v)
	}
}

// prop value enum
func (m *ServiceDTOVersion2BusBusA2Data) validateMessageTypeEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, serviceDTOVersion2BusBusA2DataTypeMessageTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceDTOVersion2BusBusA2Data) validateMessageType(formats strfmt.Registry) error {

	if swag.IsZero(m.MessageType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMessageTypeEnum("MessageType", "body", m.MessageType); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BusBusA2Data) validatePlateNumb(formats strfmt.Registry) error {

	if err := validate.Required("PlateNumb", "body", m.PlateNumb); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BusBusA2Data) validateRouteName(formats strfmt.Registry) error {

	if swag.IsZero(m.RouteName) { // not required
		return nil
	}

	if m.RouteName != nil {
		if err := m.RouteName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RouteName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2BusBusA2Data) validateStopName(formats strfmt.Registry) error {

	if swag.IsZero(m.StopName) { // not required
		return nil
	}

	if m.StopName != nil {
		if err := m.StopName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StopName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2BusBusA2Data) validateSubRouteName(formats strfmt.Registry) error {

	if swag.IsZero(m.SubRouteName) { // not required
		return nil
	}

	if m.SubRouteName != nil {
		if err := m.SubRouteName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SubRouteName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2BusBusA2Data) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2BusBusA2Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2BusBusA2Data) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2BusBusA2Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
