// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTOBusSpecificationV3A1Data A1Data
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.A1Data
type PTXServiceDTOBusSpecificationV3A1Data struct {

	// Double
	//
	// 方位角
	// Required: true
	Azimuth *float64 `json:"Azimuth"`

	// PointType
	//
	// 車輛位置經度
	// Required: true
	BusPosition struct {
		PTXServiceDTOSharedSpecificationV3BasePointType
	} `json:"BusPosition" xml:"PointType"`

	// Int32
	//
	// 行車狀況 : [0:'正常',1:'車禍',2:'故障',3:'塞車',4:'緊急求援',5:'加油',90:'不明',91:'去回不明',98:'偏移路線',99:'非營運狀態',100:'客滿',101:'包車出租',255:'未知']
	// Required: true
	BusStatus *int64 `json:"BusStatus"`

	// Int32
	//
	// 車輛去返程 : [0:'去程',1:'返程',2:'迴圈',255:'未知']
	// Required: true
	Direction *int64 `json:"Direction"`

	// Int32
	//
	// 勤務狀態 : [0:'正常',1:'開始',2:'結束']
	// Required: true
	DutyStatus *int64 `json:"DutyStatus"`

	// 車機系統紀錄時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Format: date-time
	GPSTime strfmt.DateTime `json:"GPSTime,omitempty"`

	// 車機系統傳送日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Format: date-time
	GPSTransTime strfmt.DateTime `json:"GPSTransTime,omitempty"`

	// Int32
	//
	// 資料型態種類 : [0:'未知',1:'定期',2:'非定期']
	MessageType int64 `json:"MessageType,omitempty"`

	// String
	//
	// 營運業者簡碼
	// Required: true
	OperatorCode *string `json:"OperatorCode" xml:"String"`

	// String
	//
	// 營運業者代碼
	// Required: true
	OperatorID *string `json:"OperatorID" xml:"String"`

	// NameType
	//
	// 營運業者名稱
	OperatorName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"OperatorName,omitempty" xml:"NameType,omitempty"`

	// String
	//
	// 車牌號碼
	// Required: true
	PlateNumb *string `json:"PlateNumb" xml:"String"`

	// DateTime
	//
	// 來源端平台接收時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	RecTime *strfmt.DateTime `json:"RecTime"`

	// String
	//
	// 地區既用中之路線代碼(為原資料內碼)
	// Required: true
	RouteID *string `json:"RouteID" xml:"String"`

	// NameType
	//
	// 路線名稱
	RouteName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"RouteName,omitempty" xml:"NameType,omitempty"`

	// String
	//
	// 路線唯一識別代碼，規則為 {業管機關簡碼} + {RouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	RouteUID string `json:"RouteUID,omitempty" xml:"String,omitempty"`

	// Double
	//
	// 行駛速度(kph)
	// Required: true
	Speed *float64 `json:"Speed"`

	// String
	//
	// 地區既用中之附屬路線代碼(為原資料內碼)
	SubRouteID string `json:"SubRouteID,omitempty" xml:"String,omitempty"`

	// NameType
	//
	// 附屬路線名稱
	SubRouteName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"SubRouteName,omitempty" xml:"NameType,omitempty"`

	// String
	//
	// 附屬路線唯一識別代碼，規則為 {業管機關簡碼} + {SubRouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	SubRouteUID string `json:"SubRouteUID,omitempty" xml:"String,omitempty"`

	// DateTime
	//
	// 來源端平台資料傳出時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	TransTime *strfmt.DateTime `json:"TransTime"`

	// String
	//
	// 班次代碼
	TripID string `json:"TripID,omitempty" xml:"String,omitempty"`

	// Int32
	//
	// 車輛種類 : [1:'一般',2:'復康巴士',3:'專車',4:'其他']
	VehicleType int64 `json:"VehicleType,omitempty"`
}

// Validate validates this p t x service d t o bus specification v3 a1 data
func (m *PTXServiceDTOBusSpecificationV3A1Data) Validate(formats strfmt.Registry) error {
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

	if err := m.validateGPSTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGPSTransTime(formats); err != nil {
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

func (m *PTXServiceDTOBusSpecificationV3A1Data) validateAzimuth(formats strfmt.Registry) error {

	if err := validate.Required("Azimuth", "body", m.Azimuth); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) validateBusPosition(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) validateBusStatus(formats strfmt.Registry) error {

	if err := validate.Required("BusStatus", "body", m.BusStatus); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) validateDutyStatus(formats strfmt.Registry) error {

	if err := validate.Required("DutyStatus", "body", m.DutyStatus); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) validateGPSTime(formats strfmt.Registry) error {
	if swag.IsZero(m.GPSTime) { // not required
		return nil
	}

	if err := validate.FormatOf("GPSTime", "body", "date-time", m.GPSTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) validateGPSTransTime(formats strfmt.Registry) error {
	if swag.IsZero(m.GPSTransTime) { // not required
		return nil
	}

	if err := validate.FormatOf("GPSTransTime", "body", "date-time", m.GPSTransTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) validateOperatorCode(formats strfmt.Registry) error {

	if err := validate.Required("OperatorCode", "body", m.OperatorCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) validateOperatorID(formats strfmt.Registry) error {

	if err := validate.Required("OperatorID", "body", m.OperatorID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) validateOperatorName(formats strfmt.Registry) error {
	if swag.IsZero(m.OperatorName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) validatePlateNumb(formats strfmt.Registry) error {

	if err := validate.Required("PlateNumb", "body", m.PlateNumb); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) validateRecTime(formats strfmt.Registry) error {

	if err := validate.Required("RecTime", "body", m.RecTime); err != nil {
		return err
	}

	if err := validate.FormatOf("RecTime", "body", "date-time", m.RecTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) validateRouteName(formats strfmt.Registry) error {
	if swag.IsZero(m.RouteName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) validateSpeed(formats strfmt.Registry) error {

	if err := validate.Required("Speed", "body", m.Speed); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) validateSubRouteName(formats strfmt.Registry) error {
	if swag.IsZero(m.SubRouteName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) validateTransTime(formats strfmt.Registry) error {

	if err := validate.Required("TransTime", "body", m.TransTime); err != nil {
		return err
	}

	if err := validate.FormatOf("TransTime", "body", "date-time", m.TransTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v3 a1 data based on the context it is used
func (m *PTXServiceDTOBusSpecificationV3A1Data) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBusPosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperatorName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRouteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubRouteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) contextValidateBusPosition(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) contextValidateOperatorName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) contextValidateRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3A1Data) contextValidateSubRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3A1Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3A1Data) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3A1Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
