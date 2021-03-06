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

// PTXAPIShipModelScheduleWrapperPTXServiceDTOShipSpecificationV3GeneralSchedule ShipGeneralScheduleList
//
// swagger:model PTX.API.Ship.Model.ScheduleWrapper[PTX.Service.DTO.Ship.Specification.V3.GeneralSchedule]
type PTXAPIShipModelScheduleWrapperPTXServiceDTOShipSpecificationV3GeneralSchedule struct {

	// String
	//
	// 業管機關簡碼
	AuthorityCode string `json:"AuthorityCode,omitempty"`

	// 資料總筆數<span class="emphasis fas fa-pen" rel="與來源Inbound XML不同，為提供資料的總筆數[該欄位由本平台自動產製]"></span>
	Count int64 `json:"Count,omitempty"`

	// DateTime
	//
	// 有效起始日期(資料表欄位)
	EffectiveDate string `json:"EffectiveDate,omitempty"`

	// DateTime
	//
	// YYYY-MM-DD(本欄位通常為空值)
	ExpireDate string `json:"ExpireDate,omitempty"`

	// Array
	//
	// 資料(陣列)
	GeneralSchedules []*PTXServiceDTOShipSpecificationV3GeneralSchedule `json:"GeneralSchedules"`

	// Int32
	//
	// 本平台資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// DateTime
	//
	// XML更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	UpdateTime string `json:"UpdateTime,omitempty"`

	// Int32
	//
	// 資料版本編號<span class="emphasis fas fa-pen" rel="與來源Inbound XML不同，為提供資料的版本編號[該欄位由本平台自動產製]"></span>
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this p t x API ship model schedule wrapper p t x service d t o ship specification v3 general schedule
func (m *PTXAPIShipModelScheduleWrapperPTXServiceDTOShipSpecificationV3GeneralSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGeneralSchedules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXAPIShipModelScheduleWrapperPTXServiceDTOShipSpecificationV3GeneralSchedule) validateGeneralSchedules(formats strfmt.Registry) error {
	if swag.IsZero(m.GeneralSchedules) { // not required
		return nil
	}

	for i := 0; i < len(m.GeneralSchedules); i++ {
		if swag.IsZero(m.GeneralSchedules[i]) { // not required
			continue
		}

		if m.GeneralSchedules[i] != nil {
			if err := m.GeneralSchedules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GeneralSchedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXAPIShipModelScheduleWrapperPTXServiceDTOShipSpecificationV3GeneralSchedule) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIShipModelScheduleWrapperPTXServiceDTOShipSpecificationV3GeneralSchedule) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x API ship model schedule wrapper p t x service d t o ship specification v3 general schedule based on the context it is used
func (m *PTXAPIShipModelScheduleWrapperPTXServiceDTOShipSpecificationV3GeneralSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGeneralSchedules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXAPIShipModelScheduleWrapperPTXServiceDTOShipSpecificationV3GeneralSchedule) contextValidateGeneralSchedules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GeneralSchedules); i++ {

		if m.GeneralSchedules[i] != nil {
			if err := m.GeneralSchedules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GeneralSchedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXAPIShipModelScheduleWrapperPTXServiceDTOShipSpecificationV3GeneralSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXAPIShipModelScheduleWrapperPTXServiceDTOShipSpecificationV3GeneralSchedule) UnmarshalBinary(b []byte) error {
	var res PTXAPIShipModelScheduleWrapperPTXServiceDTOShipSpecificationV3GeneralSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
