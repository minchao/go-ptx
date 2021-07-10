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

// PTXServiceDTOAirSpecificationV2FIDSArrival FIDSArrival
//
// 抵達航班顯示資料
//
// swagger:model PTX.Service.DTO.Air.Specification.V2.FIDSArrival
type PTXServiceDTOAirSpecificationV2FIDSArrival struct {

	// String
	//
	// 航空器型號
	AcType string `json:"AcType,omitempty" xml:"String,omitempty"`

	// 實際抵達時間(ISO8601格式:yyyy-MM-ddTHH:mm)
	// Format: date-time
	ActualArrivalTime strfmt.DateTime `json:"ActualArrivalTime,omitempty"`

	// Int32
	//
	// 航線種類(目前民航局與桃機的FIDS系統都尚未提供此欄位資料) : [-2:'特殊',1:'國際',2:'國內',3:'兩岸',4:'國際包機',5:'國內包機',6:'兩岸包機']
	AirRouteType int64 `json:"AirRouteType,omitempty"`

	// String
	//
	// 航空公司IATA國際代碼
	// Required: true
	AirlineID *string `json:"AirlineID" xml:"String"`

	// String
	//
	// 停機坪(僅貨機提供)
	Apron string `json:"Apron,omitempty" xml:"String,omitempty"`

	// String
	//
	// 目的地機場IATA國際代碼
	// Required: true
	ArrivalAirportID *string `json:"ArrivalAirportID" xml:"String"`

	// String
	//
	// 航班屬性狀態,為該機場觀點的狀態
	ArrivalRemark string `json:"ArrivalRemark,omitempty" xml:"String,omitempty"`

	// String
	//
	// 航班屬性狀態(英文)
	ArrivalRemarkEn string `json:"ArrivalRemarkEn,omitempty" xml:"String,omitempty"`

	// String
	//
	// 行李轉盤(到站FIDS可能有「行李轉盤」資訊, 離站FIDS不會有, 貨機則無此資訊)
	BaggageClaim string `json:"BaggageClaim,omitempty" xml:"String,omitempty"`

	// String
	//
	// 報到櫃檯(離站FIDS可能有「報到櫃台」資訊, 到站FIDS不會有, 貨機則無此資訊)
	CheckCounter string `json:"CheckCounter,omitempty" xml:"String,omitempty"`

	// String
	//
	// 航班共用班號
	CodeShare string `json:"CodeShare,omitempty" xml:"String,omitempty"`

	// String
	//
	// 起點機場IATA國際代碼
	// Required: true
	DepartureAirportID *string `json:"DepartureAirportID" xml:"String"`

	// 預估抵達時間(ISO8601格式:yyyy-MM-ddTHH:mm)
	// Format: date-time
	EstimatedArrivalTime strfmt.DateTime `json:"EstimatedArrivalTime,omitempty"`

	// 航班日期(ISO8601格式:yyyy-MM-dd)
	// Required: true
	// Format: date-time
	FlightDate *strfmt.DateTime `json:"FlightDate"`

	// String
	//
	// 航機班號(不包含航空公司的AirlineID，僅有班號數字)
	// Required: true
	FlightNumber *string `json:"FlightNumber" xml:"String"`

	// String
	//
	// 登機門(僅客機提供)
	Gate string `json:"Gate,omitempty" xml:"String,omitempty"`

	// Boolean
	//
	// 是否為貨機
	IsCargo bool `json:"IsCargo,omitempty"`

	// 表訂抵達時間(ISO8601格式:yyyy-MM-ddTHH:mm)
	// Format: date-time
	ScheduleArrivalTime strfmt.DateTime `json:"ScheduleArrivalTime,omitempty"`

	// String
	//
	// 航廈
	Terminal string `json:"Terminal,omitempty" xml:"String,omitempty"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o air specification v2 f ID s arrival
func (m *PTXServiceDTOAirSpecificationV2FIDSArrival) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActualArrivalTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAirlineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArrivalAirportID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartureAirportID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEstimatedArrivalTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlightDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlightNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduleArrivalTime(formats); err != nil {
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

func (m *PTXServiceDTOAirSpecificationV2FIDSArrival) validateActualArrivalTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ActualArrivalTime) { // not required
		return nil
	}

	if err := validate.FormatOf("ActualArrivalTime", "body", "date-time", m.ActualArrivalTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2FIDSArrival) validateAirlineID(formats strfmt.Registry) error {

	if err := validate.Required("AirlineID", "body", m.AirlineID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2FIDSArrival) validateArrivalAirportID(formats strfmt.Registry) error {

	if err := validate.Required("ArrivalAirportID", "body", m.ArrivalAirportID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2FIDSArrival) validateDepartureAirportID(formats strfmt.Registry) error {

	if err := validate.Required("DepartureAirportID", "body", m.DepartureAirportID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2FIDSArrival) validateEstimatedArrivalTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EstimatedArrivalTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EstimatedArrivalTime", "body", "date-time", m.EstimatedArrivalTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2FIDSArrival) validateFlightDate(formats strfmt.Registry) error {

	if err := validate.Required("FlightDate", "body", m.FlightDate); err != nil {
		return err
	}

	if err := validate.FormatOf("FlightDate", "body", "date-time", m.FlightDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2FIDSArrival) validateFlightNumber(formats strfmt.Registry) error {

	if err := validate.Required("FlightNumber", "body", m.FlightNumber); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2FIDSArrival) validateScheduleArrivalTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ScheduleArrivalTime) { // not required
		return nil
	}

	if err := validate.FormatOf("ScheduleArrivalTime", "body", "date-time", m.ScheduleArrivalTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2FIDSArrival) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o air specification v2 f ID s arrival based on context it is used
func (m *PTXServiceDTOAirSpecificationV2FIDSArrival) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOAirSpecificationV2FIDSArrival) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOAirSpecificationV2FIDSArrival) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOAirSpecificationV2FIDSArrival
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
