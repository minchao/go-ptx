// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTOAirSpecificationV2GeneralFlightSchedule GeneralFlightSchedule
//
// 航空定期時刻表資料
//
// swagger:model PTX.Service.DTO.Air.Specification.V2.GeneralFlightSchedule
type PTXServiceDTOAirSpecificationV2GeneralFlightSchedule struct {

	// String
	//
	// 航空公司IATA國際代碼
	// Required: true
	AirlineID *string `json:"AirlineID" xml:"AirlineID"`

	// String
	//
	// 目的地機場IATA國際代碼
	// Required: true
	ArrivalAirportID *string `json:"ArrivalAirportID" xml:"ArrivalAirportID"`

	// String
	//
	// 終點機場抵達時間 (格式: HH:mm 當地時間，跨日以+1 表示)
	// Required: true
	ArrivalTime *string `json:"ArrivalTime" xml:"ArrivalTime"`

	// Array
	//
	// 共用班號
	CodeShare []*PTXServiceDTOAirSpecificationV2CodeShare "json:\"CodeShare\" xml:\"List`1\""

	// String
	//
	// 起點機場IATA國際代碼
	// Required: true
	DepartureAirportID *string `json:"DepartureAirportID" xml:"DepartureAirportID"`

	// String
	//
	// 起點機場出發時間 (格式: HH:mm 當地時間，跨日以+1 表示)
	// Required: true
	DepartureTime *string `json:"DepartureTime" xml:"DepartureTime"`

	// String
	//
	// 班機號碼(包含航空公司的AirlineID，結構為AirlineID加上3~4碼航機班號數字；若班號僅有兩碼，其結構會加上0補足三碼，"AirlineID + 0 + 兩碼班號")
	// Required: true
	FlightNumber *string `json:"FlightNumber" xml:"FlightNumber"`

	// Boolean
	//
	// 週五飛行與否
	// Required: true
	Friday *bool `json:"Friday"`

	// Boolean
	//
	// 週一飛行與否
	// Required: true
	Monday *bool `json:"Monday"`

	// Boolean
	//
	// 週六飛行與否
	// Required: true
	Saturday *bool `json:"Saturday"`

	// DateTime
	//
	// 班表結束日期(ISO8601格式:yyyy-MM-dd)
	// Required: true
	// Format: date-time
	ScheduleEndDate *strfmt.DateTime `json:"ScheduleEndDate"`

	// DateTime
	//
	// 班表開始日期(ISO8601格式:yyyy-MM-dd)
	// Required: true
	// Format: date-time
	ScheduleStartDate *strfmt.DateTime `json:"ScheduleStartDate"`

	// Boolean
	//
	// 週日飛行與否
	// Required: true
	Sunday *bool `json:"Sunday"`

	// String
	//
	// 航廈
	Terminal string `json:"Terminal,omitempty" xml:"Terminal,omitempty"`

	// Boolean
	//
	// 週四飛行與否
	// Required: true
	Thursday *bool `json:"Thursday"`

	// Boolean
	//
	// 週二飛行與否
	// Required: true
	Tuesday *bool `json:"Tuesday"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`

	// Int32
	//
	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`

	// Boolean
	//
	// 週三飛行與否
	// Required: true
	Wednesday *bool `json:"Wednesday"`
}

// Validate validates this p t x service d t o air specification v2 general flight schedule
func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAirlineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArrivalAirportID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArrivalTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCodeShare(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartureAirportID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartureTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlightNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFriday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSaturday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduleEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduleStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSunday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThursday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTuesday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWednesday(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateAirlineID(formats strfmt.Registry) error {

	if err := validate.Required("AirlineID", "body", m.AirlineID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateArrivalAirportID(formats strfmt.Registry) error {

	if err := validate.Required("ArrivalAirportID", "body", m.ArrivalAirportID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateArrivalTime(formats strfmt.Registry) error {

	if err := validate.Required("ArrivalTime", "body", m.ArrivalTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateCodeShare(formats strfmt.Registry) error {
	if swag.IsZero(m.CodeShare) { // not required
		return nil
	}

	for i := 0; i < len(m.CodeShare); i++ {
		if swag.IsZero(m.CodeShare[i]) { // not required
			continue
		}

		if m.CodeShare[i] != nil {
			if err := m.CodeShare[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CodeShare" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CodeShare" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateDepartureAirportID(formats strfmt.Registry) error {

	if err := validate.Required("DepartureAirportID", "body", m.DepartureAirportID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateDepartureTime(formats strfmt.Registry) error {

	if err := validate.Required("DepartureTime", "body", m.DepartureTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateFlightNumber(formats strfmt.Registry) error {

	if err := validate.Required("FlightNumber", "body", m.FlightNumber); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateFriday(formats strfmt.Registry) error {

	if err := validate.Required("Friday", "body", m.Friday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateMonday(formats strfmt.Registry) error {

	if err := validate.Required("Monday", "body", m.Monday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateSaturday(formats strfmt.Registry) error {

	if err := validate.Required("Saturday", "body", m.Saturday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateScheduleEndDate(formats strfmt.Registry) error {

	if err := validate.Required("ScheduleEndDate", "body", m.ScheduleEndDate); err != nil {
		return err
	}

	if err := validate.FormatOf("ScheduleEndDate", "body", "date-time", m.ScheduleEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateScheduleStartDate(formats strfmt.Registry) error {

	if err := validate.Required("ScheduleStartDate", "body", m.ScheduleStartDate); err != nil {
		return err
	}

	if err := validate.FormatOf("ScheduleStartDate", "body", "date-time", m.ScheduleStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateSunday(formats strfmt.Registry) error {

	if err := validate.Required("Sunday", "body", m.Sunday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateThursday(formats strfmt.Registry) error {

	if err := validate.Required("Thursday", "body", m.Thursday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateTuesday(formats strfmt.Registry) error {

	if err := validate.Required("Tuesday", "body", m.Tuesday); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) validateWednesday(formats strfmt.Registry) error {

	if err := validate.Required("Wednesday", "body", m.Wednesday); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o air specification v2 general flight schedule based on the context it is used
func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCodeShare(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) contextValidateCodeShare(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CodeShare); i++ {

		if m.CodeShare[i] != nil {
			if err := m.CodeShare[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CodeShare" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CodeShare" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOAirSpecificationV2GeneralFlightSchedule) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOAirSpecificationV2GeneralFlightSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
