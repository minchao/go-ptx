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

// PTXServiceDTOTourismSpecificationV2BusA2Data BusA2Data
//
// 定點車機資料型別
//
// swagger:model PTX.Service.DTO.Tourism.Specification.V2.BusA2Data
type PTXServiceDTOTourismSpecificationV2BusA2Data struct {

	// Int32
	//
	// 進站離站 : [0:'離站',1:'進站']
	A2EventType int64 `json:"A2EventType,omitempty"`

	// Int32
	//
	// 行車狀況 : [0:'正常',1:'車禍',2:'故障',3:'塞車',4:'緊急求援',5:'加油',90:'不明',91:'去回不明',98:'偏移路線',99:'非營運狀態',100:'客滿',101:'包車出租',255:'未知']
	BusStatus int64 `json:"BusStatus,omitempty"`

	// Int32
	//
	// 去返程 : [0:'去程',1:'返程',2:'迴圈',255:'未知']
	// Required: true
	Direction *int64 `json:"Direction"`

	// Int32
	//
	// 勤務狀態 : [0:'正常',1:'開始',2:'結束']
	DutyStatus int64 `json:"DutyStatus,omitempty"`

	// DateTime
	//
	// 車機時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	GPSTime *strfmt.DateTime `json:"GPSTime"`

	// Int32
	//
	// 資料型態種類 : [0:'未知',1:'定期',2:'非定期']
	MessageType int64 `json:"MessageType,omitempty"`

	// String
	//
	// 營運業者代碼
	OperatorID string `json:"OperatorID,omitempty" xml:"OperatorID,omitempty"`

	// String
	//
	// 車牌號碼
	// Required: true
	PlateNumb *string `json:"PlateNumb" xml:"PlateNumb"`

	// String
	//
	// 地區既用中之路線代碼(為原資料內碼)
	RouteID string `json:"RouteID,omitempty" xml:"RouteID,omitempty"`

	// String
	//
	// 路線唯一識別代碼，規則為 {業管機關代碼} + {RouteID}，其中 {業管機關代碼} 可於Authority API中的AuthorityCode欄位查詢
	RouteUID string `json:"RouteUID,omitempty" xml:"RouteUID,omitempty"`

	// 來源端平台接收時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Format: date-time
	SrcRecTime strfmt.DateTime `json:"SrcRecTime,omitempty"`

	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Format: date-time
	SrcUpdateTime strfmt.DateTime `json:"SrcUpdateTime,omitempty"`

	// String
	//
	// 地區既用中之站牌代號(為原資料內碼)
	StopID string `json:"StopID,omitempty" xml:"StopID,omitempty"`

	// NameType
	//
	// 站牌名
	StopName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StopName,omitempty" xml:"NameType,omitempty"`

	// String
	//
	// 站牌唯一識別代碼，規則為 {平台代碼} + {StopID}，其中 {平台代碼} 可於Provider API中的ProviderCode欄位查詢
	StopUID string `json:"StopUID,omitempty" xml:"StopUID,omitempty"`

	// String
	//
	// 地區既用中之子路線代碼(為原資料內碼)
	SubRouteID string `json:"SubRouteID,omitempty" xml:"SubRouteID,omitempty"`

	// NameType
	//
	// 子路線名稱
	SubRouteName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"SubRouteName,omitempty" xml:"NameType,omitempty"`

	// String
	//
	// 子路線唯一識別代碼，規則為 {業管機關代碼} + {SubRouteID}，其中 {業管機關代碼} 可於Authority API中的AuthorityCode欄位查詢
	SubRouteUID string `json:"SubRouteUID,omitempty" xml:"SubRouteUID,omitempty"`

	// NameType
	//
	// 台灣好行路線名稱
	TaiwanTripName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"TaiwanTripName,omitempty" xml:"NameType,omitempty"`

	// 車機資料傳輸時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Format: date-time
	TransTime strfmt.DateTime `json:"TransTime,omitempty"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o tourism specification v2 bus a2 data
func (m *PTXServiceDTOTourismSpecificationV2BusA2Data) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGPSTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlateNumb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcRecTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubRouteName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaiwanTripName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransTime(formats); err != nil {
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

func (m *PTXServiceDTOTourismSpecificationV2BusA2Data) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusA2Data) validateGPSTime(formats strfmt.Registry) error {

	if err := validate.Required("GPSTime", "body", m.GPSTime); err != nil {
		return err
	}

	if err := validate.FormatOf("GPSTime", "body", "date-time", m.GPSTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusA2Data) validatePlateNumb(formats strfmt.Registry) error {

	if err := validate.Required("PlateNumb", "body", m.PlateNumb); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusA2Data) validateSrcRecTime(formats strfmt.Registry) error {
	if swag.IsZero(m.SrcRecTime) { // not required
		return nil
	}

	if err := validate.FormatOf("SrcRecTime", "body", "date-time", m.SrcRecTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusA2Data) validateSrcUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.SrcUpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusA2Data) validateStopName(formats strfmt.Registry) error {
	if swag.IsZero(m.StopName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusA2Data) validateSubRouteName(formats strfmt.Registry) error {
	if swag.IsZero(m.SubRouteName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusA2Data) validateTaiwanTripName(formats strfmt.Registry) error {
	if swag.IsZero(m.TaiwanTripName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusA2Data) validateTransTime(formats strfmt.Registry) error {
	if swag.IsZero(m.TransTime) { // not required
		return nil
	}

	if err := validate.FormatOf("TransTime", "body", "date-time", m.TransTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusA2Data) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o tourism specification v2 bus a2 data based on the context it is used
func (m *PTXServiceDTOTourismSpecificationV2BusA2Data) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStopName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubRouteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaiwanTripName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusA2Data) contextValidateStopName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusA2Data) contextValidateSubRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusA2Data) contextValidateTaiwanTripName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2BusA2Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2BusA2Data) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOTourismSpecificationV2BusA2Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
