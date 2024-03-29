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

// PTXServiceDTOBusSpecificationV3N1Data N1Data
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.N1Data
type PTXServiceDTOBusSpecificationV3N1Data struct {

	// String
	//
	// 車輛目前所在站牌代碼
	CurrentStop string `json:"CurrentStop,omitempty" xml:"CurrentStop,omitempty"`

	// 系統演算該筆預估到站資料的時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Format: date-time
	DataTime strfmt.DateTime `json:"DataTime,omitempty"`

	// String
	//
	// 迄點站站牌ID代碼
	DestinationStopID string `json:"DestinationStopID,omitempty" xml:"DestinationStopID,omitempty"`

	// NameType
	//
	// 迄點站站牌名稱
	DestinationStopName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"DestinationStopName,omitempty" xml:"NameType,omitempty"`

	// Int32
	//
	// 車輛去返程(該方向指的是此公車運具目前所在路線的去返程方向，非指站牌所在路線的去返程方向，使用時請加值業者多加注意) : [0:'去程',1:'返程',2:'迴圈',255:'未知']
	// Required: true
	Direction *int64 `json:"Direction"`

	// 到站時間預估(秒) [當StopStatus値為1~4或PlateNumb値為-1時，EstimateTime値為空値; 反之，EstimateTime有値]
	EstimateTime int32 `json:"EstimateTime,omitempty"`

	// 是否為末班車
	IsLastBus bool `json:"IsLastBus,omitempty"`

	// String
	//
	// 車牌號碼 [値為値為-1時，表示目前該站牌無車輛行駛]
	PlateNumb string `json:"PlateNumb,omitempty" xml:"PlateNumb,omitempty"`

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
	RouteID *string `json:"RouteID" xml:"RouteID"`

	// NameType
	//
	// 路線名稱
	RouteName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"RouteName,omitempty" xml:"NameType,omitempty"`

	// String
	//
	// 路線唯一識別代碼，規則為 {業管機關代碼} + {RouteID}，其中 {業管機關代碼} 可於Authority API中的AuthorityCode欄位查詢
	RouteUID string `json:"RouteUID,omitempty" xml:"RouteUID,omitempty"`

	// String
	//
	// 預排班表時間
	ScheduledTime string `json:"ScheduledTime,omitempty" xml:"ScheduledTime,omitempty"`

	// 路線經過站牌之順序
	StopCountDown int32 `json:"StopCountDown,omitempty"`

	// String
	//
	// 地區既用中之站牌代碼(為原資料內碼)
	// Required: true
	StopID *string `json:"StopID" xml:"StopID"`

	// NameType
	//
	// 站牌名稱
	StopName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"StopName,omitempty" xml:"NameType,omitempty"`

	// Int32
	//
	// 車輛狀態備註 : [0:'正常',1:'尚未發車',2:'交管不停靠',3:'末班車已過',4:'今日未營運']
	StopStatus int64 `json:"StopStatus,omitempty"`

	// String
	//
	// 站牌唯一識別代碼，規則為 {業管機關簡碼} + {StopID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	StopUID string `json:"StopUID,omitempty" xml:"StopUID,omitempty"`

	// String
	//
	// 地區既用中之附屬路線代碼(為原資料內碼)
	SubRouteID string `json:"SubRouteID,omitempty" xml:"SubRouteID,omitempty"`

	// NameType
	//
	// 附屬路線名稱
	SubRouteName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"SubRouteName,omitempty" xml:"NameType,omitempty"`

	// String
	//
	// 附屬路線唯一識別代碼，規則為 {業管機關簡碼} + {SubRouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	SubRouteUID string `json:"SubRouteUID,omitempty" xml:"SubRouteUID,omitempty"`

	// DateTime
	//
	// 來源端平台資料傳出時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	TransTime *strfmt.DateTime `json:"TransTime"`
}

// Validate validates this p t x service d t o bus specification v3 n1 data
func (m *PTXServiceDTOBusSpecificationV3N1Data) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationStopName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
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

	if err := m.validateStopID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopName(formats); err != nil {
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

func (m *PTXServiceDTOBusSpecificationV3N1Data) validateDataTime(formats strfmt.Registry) error {
	if swag.IsZero(m.DataTime) { // not required
		return nil
	}

	if err := validate.FormatOf("DataTime", "body", "date-time", m.DataTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3N1Data) validateDestinationStopName(formats strfmt.Registry) error {
	if swag.IsZero(m.DestinationStopName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3N1Data) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3N1Data) validateRecTime(formats strfmt.Registry) error {

	if err := validate.Required("RecTime", "body", m.RecTime); err != nil {
		return err
	}

	if err := validate.FormatOf("RecTime", "body", "date-time", m.RecTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3N1Data) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3N1Data) validateRouteName(formats strfmt.Registry) error {
	if swag.IsZero(m.RouteName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3N1Data) validateStopID(formats strfmt.Registry) error {

	if err := validate.Required("StopID", "body", m.StopID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3N1Data) validateStopName(formats strfmt.Registry) error {
	if swag.IsZero(m.StopName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3N1Data) validateSubRouteName(formats strfmt.Registry) error {
	if swag.IsZero(m.SubRouteName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3N1Data) validateTransTime(formats strfmt.Registry) error {

	if err := validate.Required("TransTime", "body", m.TransTime); err != nil {
		return err
	}

	if err := validate.FormatOf("TransTime", "body", "date-time", m.TransTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v3 n1 data based on the context it is used
func (m *PTXServiceDTOBusSpecificationV3N1Data) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestinationStopName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRouteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStopName(ctx, formats); err != nil {
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

func (m *PTXServiceDTOBusSpecificationV3N1Data) contextValidateDestinationStopName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3N1Data) contextValidateRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3N1Data) contextValidateStopName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3N1Data) contextValidateSubRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3N1Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3N1Data) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3N1Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
