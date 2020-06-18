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

// ServiceDTOVersion2TaiwanTripBusBusA1Data BusA1Data
//
// 定時車機資料型別
//
// swagger:model Service.DTO.Version2.TaiwanTripBus.BusA1Data
type ServiceDTOVersion2TaiwanTripBusBusA1Data struct {

	// 方位角
	// Required: true
	Azimuth *float64 `json:"Azimuth"`

	// PointType
	//
	// 車輛位置經度
	BusPosition *ServiceDTOVersion2BasePointType `json:"BusPosition,omitempty"`

	// integer
	//
	// 行車狀況 : [0:'正常',1:'車禍',2:'故障',3:'塞車',4:'緊急求援',5:'加油',90:'不明',91:'去回不明',98:'偏移路線',99:'非營運狀態',100:'客滿',101:'包車出租',255:'未知']
	// Required: true
	BusStatus *int32 `json:"BusStatus"`

	// integer
	//
	// 去返程 : [0:'去程',1:'返程',2:'迴圈',255:'未知']
	Direction int32 `json:"Direction,omitempty"`

	// integer
	//
	// 勤務狀態 : [0:'正常',1:'開始',2:'結束']
	// Required: true
	DutyStatus *int32 `json:"DutyStatus"`

	// DateTime
	//
	// 車機時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	GPSTime *string `json:"GPSTime"`

	// integer
	//
	// 資料型態種類 : [0:'未知',1:'定期',2:'非定期']
	MessageType int32 `json:"MessageType,omitempty"`

	// 營運業者代碼
	OperatorID string `json:"OperatorID,omitempty"`

	// 車牌號碼
	// Required: true
	PlateNumb *string `json:"PlateNumb"`

	// 地區既用中之路線代碼(為原資料內碼)
	RouteID string `json:"RouteID,omitempty"`

	// 路線唯一識別代碼，規則為 {業管機關代碼} + {RouteID}，其中 {業管機關代碼} 可於Authority API中的AuthorityCode欄位查詢
	RouteUID string `json:"RouteUID,omitempty"`

	// 行駛速度(kph)
	// Required: true
	Speed *float64 `json:"Speed"`

	// DateTime
	//
	// 來源端平台接收時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	SrcRecTime string `json:"SrcRecTime,omitempty"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	SrcUpdateTime string `json:"SrcUpdateTime,omitempty"`

	// 地區既用中之子路線代碼(為原資料內碼)
	SubRouteID string `json:"SubRouteID,omitempty"`

	// NameType
	//
	// 子路線名稱
	SubRouteName *ServiceDTOVersion2BaseNameType `json:"SubRouteName,omitempty"`

	// 子路線唯一識別代碼，規則為 {業管機關代碼} + {SubRouteID}，其中 {業管機關代碼} 可於Authority API中的AuthorityCode欄位查詢
	SubRouteUID string `json:"SubRouteUID,omitempty"`

	// NameType
	//
	// 台灣好行路線名稱
	TaiwanTripName *ServiceDTOVersion2BaseNameType `json:"TaiwanTripName,omitempty"`

	// DateTime
	//
	// 車機資料傳輸時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	TransTime string `json:"TransTime,omitempty"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this service d t o version2 taiwan trip bus bus a1 data
func (m *ServiceDTOVersion2TaiwanTripBusBusA1Data) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzimuth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBusPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBusStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDutyStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGPSTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlateNumb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubRouteName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaiwanTripName(formats); err != nil {
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

func (m *ServiceDTOVersion2TaiwanTripBusBusA1Data) validateAzimuth(formats strfmt.Registry) error {

	if err := validate.Required("Azimuth", "body", m.Azimuth); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2TaiwanTripBusBusA1Data) validateBusPosition(formats strfmt.Registry) error {

	if swag.IsZero(m.BusPosition) { // not required
		return nil
	}

	if m.BusPosition != nil {
		if err := m.BusPosition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BusPosition")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2TaiwanTripBusBusA1Data) validateBusStatus(formats strfmt.Registry) error {

	if err := validate.Required("BusStatus", "body", m.BusStatus); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2TaiwanTripBusBusA1Data) validateDutyStatus(formats strfmt.Registry) error {

	if err := validate.Required("DutyStatus", "body", m.DutyStatus); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2TaiwanTripBusBusA1Data) validateGPSTime(formats strfmt.Registry) error {

	if err := validate.Required("GPSTime", "body", m.GPSTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2TaiwanTripBusBusA1Data) validatePlateNumb(formats strfmt.Registry) error {

	if err := validate.Required("PlateNumb", "body", m.PlateNumb); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2TaiwanTripBusBusA1Data) validateSpeed(formats strfmt.Registry) error {

	if err := validate.Required("Speed", "body", m.Speed); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2TaiwanTripBusBusA1Data) validateSubRouteName(formats strfmt.Registry) error {

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

func (m *ServiceDTOVersion2TaiwanTripBusBusA1Data) validateTaiwanTripName(formats strfmt.Registry) error {

	if swag.IsZero(m.TaiwanTripName) { // not required
		return nil
	}

	if m.TaiwanTripName != nil {
		if err := m.TaiwanTripName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TaiwanTripName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2TaiwanTripBusBusA1Data) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2TaiwanTripBusBusA1Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2TaiwanTripBusBusA1Data) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2TaiwanTripBusBusA1Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
