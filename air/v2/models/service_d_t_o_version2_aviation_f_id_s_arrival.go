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

// ServiceDTOVersion2AviationFIDSArrival FIDSArrival
//
// 抵達航班顯示資料
//
// swagger:model Service.DTO.Version2.Aviation.FIDSArrival
type ServiceDTOVersion2AviationFIDSArrival struct {

	// 航空器型號
	AcType string `json:"AcType,omitempty"`

	// DateTime
	//
	// 實際抵達時間(ISO8601格式:yyyy-MM-ddTHH:mm)
	ActualArrivalTime string `json:"ActualArrivalTime,omitempty"`

	// integer
	//
	// 航線種類(目前民航局與桃機的FIDS系統都尚未提供此欄位資料) : [-2:'特殊',1:'國際',2:'國內',3:'兩岸',4:'國際包機',5:'國內包機',6:'兩岸包機']
	AirRouteType int32 `json:"AirRouteType,omitempty"`

	// 航空公司IATA國際代碼
	// Required: true
	AirlineID *string `json:"AirlineID"`

	// 停機坪(僅貨機提供)
	Apron string `json:"Apron,omitempty"`

	// 目的地機場IATA國際代碼
	// Required: true
	ArrivalAirportID *string `json:"ArrivalAirportID"`

	// 航班屬性狀態,為該機場觀點的狀態
	ArrivalRemark string `json:"ArrivalRemark,omitempty"`

	// 航班屬性狀態(英文)
	ArrivalRemarkEn string `json:"ArrivalRemarkEn,omitempty"`

	// 行李轉盤(到站FIDS可能有「行李轉盤」資訊, 離站FIDS不會有, 貨機則無此資訊)
	BaggageClaim string `json:"BaggageClaim,omitempty"`

	// 報到櫃檯(離站FIDS可能有「報到櫃台」資訊, 到站FIDS不會有, 貨機則無此資訊)
	CheckCounter string `json:"CheckCounter,omitempty"`

	// 航班共用班號
	CodeShare string `json:"CodeShare,omitempty"`

	// 起點機場IATA國際代碼
	// Required: true
	DepartureAirportID *string `json:"DepartureAirportID"`

	// DateTime
	//
	// 預估抵達時間(ISO8601格式:yyyy-MM-ddTHH:mm)
	EstimatedArrivalTime string `json:"EstimatedArrivalTime,omitempty"`

	// DateTime
	//
	// 航班日期(ISO8601格式:yyyy-MM-dd)
	// Required: true
	FlightDate *string `json:"FlightDate"`

	// 航機班號(不包含航空公司的AirlineID，僅有班號數字)
	// Required: true
	FlightNumber *string `json:"FlightNumber"`

	// 登機門(僅客機提供)
	Gate string `json:"Gate,omitempty"`

	// 是否為貨機
	IsCargo bool `json:"IsCargo,omitempty"`

	// DateTime
	//
	// 表訂抵達時間(ISO8601格式:yyyy-MM-ddTHH:mm)
	ScheduleArrivalTime string `json:"ScheduleArrivalTime,omitempty"`

	// 航廈
	Terminal string `json:"Terminal,omitempty"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this service d t o version2 aviation f ID s arrival
func (m *ServiceDTOVersion2AviationFIDSArrival) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAirlineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArrivalAirportID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartureAirportID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlightDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlightNumber(formats); err != nil {
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

func (m *ServiceDTOVersion2AviationFIDSArrival) validateAirlineID(formats strfmt.Registry) error {

	if err := validate.Required("AirlineID", "body", m.AirlineID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2AviationFIDSArrival) validateArrivalAirportID(formats strfmt.Registry) error {

	if err := validate.Required("ArrivalAirportID", "body", m.ArrivalAirportID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2AviationFIDSArrival) validateDepartureAirportID(formats strfmt.Registry) error {

	if err := validate.Required("DepartureAirportID", "body", m.DepartureAirportID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2AviationFIDSArrival) validateFlightDate(formats strfmt.Registry) error {

	if err := validate.Required("FlightDate", "body", m.FlightDate); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2AviationFIDSArrival) validateFlightNumber(formats strfmt.Registry) error {

	if err := validate.Required("FlightNumber", "body", m.FlightNumber); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2AviationFIDSArrival) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2AviationFIDSArrival) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2AviationFIDSArrival) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2AviationFIDSArrival
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
