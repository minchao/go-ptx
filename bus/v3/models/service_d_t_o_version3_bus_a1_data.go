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

// ServiceDTOVersion3BusA1Data A1Data
// swagger:model Service.DTO.Version3.Bus.A1Data
type ServiceDTOVersion3BusA1Data struct {

	// 方位角
	// Required: true
	Azimuth *float64 `json:"Azimuth"`

	// PointType
	//
	// 車輛位置經度
	// Required: true
	BusPosition *ServiceDTOVersion3BasePointType `json:"BusPosition"`

	// integer
	//
	// 行車狀況 : [0:'正常',1:'車禍',2:'故障',3:'塞車',4:'緊急求援',5:'加油',90:'不明',91:'去回不明',98:'偏移路線',99:'非營運狀態',100:'客滿',101:'包車出租',255:'未知']
	// Required: true
	BusStatus *int32 `json:"BusStatus"`

	// integer
	//
	// 車輛去返程 : [0:'去程',1:'返程',2:'迴圈',255:'未知']
	// Required: true
	Direction *int32 `json:"Direction"`

	// integer
	//
	// 勤務狀態 : [0:'正常',1:'開始',2:'結束']
	// Required: true
	DutyStatus *int32 `json:"DutyStatus"`

	// DateTime
	//
	// 車機系統紀錄時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	GPSTime string `json:"GPSTime,omitempty"`

	// DateTime
	//
	// 車機系統傳送日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	GPSTransTime string `json:"GPSTransTime,omitempty"`

	// integer
	//
	// 資料型態種類 : [0:'未知',1:'定期',2:'非定期']
	MessageType int32 `json:"MessageType,omitempty"`

	// 營運業者簡碼
	// Required: true
	OperatorCode *string `json:"OperatorCode"`

	// 營運業者代碼
	// Required: true
	OperatorID *string `json:"OperatorID"`

	// NameType
	//
	// 營運業者名稱
	OperatorName *ServiceDTOVersion3BaseNameType `json:"OperatorName,omitempty"`

	// 車牌號碼
	// Required: true
	PlateNumb *string `json:"PlateNumb"`

	// DateTime
	//
	// 來源端平台接收時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	RecTime *string `json:"RecTime"`

	// 地區既用中之路線代碼(為原資料內碼)
	// Required: true
	RouteID *string `json:"RouteID"`

	// NameType
	//
	// 路線名稱
	RouteName *ServiceDTOVersion3BaseNameType `json:"RouteName,omitempty"`

	// 路線唯一識別代碼，規則為 {業管機關簡碼} + {RouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	RouteUID string `json:"RouteUID,omitempty"`

	// 行駛速度(kph)
	// Required: true
	Speed *float64 `json:"Speed"`

	// 地區既用中之附屬路線代碼(為原資料內碼)
	SubRouteID string `json:"SubRouteID,omitempty"`

	// NameType
	//
	// 附屬路線名稱
	SubRouteName *ServiceDTOVersion3BaseNameType `json:"SubRouteName,omitempty"`

	// 附屬路線唯一識別代碼，規則為 {業管機關簡碼} + {SubRouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	SubRouteUID string `json:"SubRouteUID,omitempty"`

	// DateTime
	//
	// 來源端平台資料傳出時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	TransTime *string `json:"TransTime"`

	// 班次代碼
	TripID string `json:"TripID,omitempty"`

	// integer
	//
	// 車輛種類 : [1:'一般',2:'復康巴士',3:'專車',4:'其他']
	VehicleType int32 `json:"VehicleType,omitempty"`
}

// Validate validates this service d t o version3 bus a1 data
func (m *ServiceDTOVersion3BusA1Data) Validate(formats strfmt.Registry) error {
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

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDutyStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlateNumb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubRouteName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3BusA1Data) validateAzimuth(formats strfmt.Registry) error {

	if err := validate.Required("Azimuth", "body", m.Azimuth); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusA1Data) validateBusPosition(formats strfmt.Registry) error {

	if err := validate.Required("BusPosition", "body", m.BusPosition); err != nil {
		return err
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

func (m *ServiceDTOVersion3BusA1Data) validateBusStatus(formats strfmt.Registry) error {

	if err := validate.Required("BusStatus", "body", m.BusStatus); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusA1Data) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusA1Data) validateDutyStatus(formats strfmt.Registry) error {

	if err := validate.Required("DutyStatus", "body", m.DutyStatus); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusA1Data) validateOperatorCode(formats strfmt.Registry) error {

	if err := validate.Required("OperatorCode", "body", m.OperatorCode); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusA1Data) validateOperatorID(formats strfmt.Registry) error {

	if err := validate.Required("OperatorID", "body", m.OperatorID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusA1Data) validateOperatorName(formats strfmt.Registry) error {

	if swag.IsZero(m.OperatorName) { // not required
		return nil
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

func (m *ServiceDTOVersion3BusA1Data) validatePlateNumb(formats strfmt.Registry) error {

	if err := validate.Required("PlateNumb", "body", m.PlateNumb); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusA1Data) validateRecTime(formats strfmt.Registry) error {

	if err := validate.Required("RecTime", "body", m.RecTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusA1Data) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusA1Data) validateRouteName(formats strfmt.Registry) error {

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

func (m *ServiceDTOVersion3BusA1Data) validateSpeed(formats strfmt.Registry) error {

	if err := validate.Required("Speed", "body", m.Speed); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusA1Data) validateSubRouteName(formats strfmt.Registry) error {

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

func (m *ServiceDTOVersion3BusA1Data) validateTransTime(formats strfmt.Registry) error {

	if err := validate.Required("TransTime", "body", m.TransTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3BusA1Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3BusA1Data) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3BusA1Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
