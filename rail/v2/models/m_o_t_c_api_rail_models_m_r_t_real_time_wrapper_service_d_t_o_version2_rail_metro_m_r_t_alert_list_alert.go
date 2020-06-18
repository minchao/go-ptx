// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTAlertListAlert MRTAlertList
//
// swagger:model MOTC.API.Rail.Models.MRTRealTimeWrapper[Service.DTO.Version2.Rail.Metro.MRTAlertList.Alert]
type MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTAlertListAlert struct {

	// 資料(陣列)
	// Required: true
	Alerts []*ServiceDTOVersion2RailMetroMRTAlertListAlert `json:"Alerts"`

	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode"`

	// 資料總筆數
	Count int64 `json:"Count,omitempty"`

	// 來源端平台資料更新週期(秒)['-1: 不定期更新']
	// Required: true
	SrcUpdateInterval *int32 `json:"SrcUpdateInterval"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// 本平台資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this m o t c API rail models m r t real time wrapper service d t o version2 rail metro m r t alert list alert
func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTAlertListAlert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlerts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateInterval(formats); err != nil {
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

func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTAlertListAlert) validateAlerts(formats strfmt.Registry) error {

	if err := validate.Required("Alerts", "body", m.Alerts); err != nil {
		return err
	}

	for i := 0; i < len(m.Alerts); i++ {
		if swag.IsZero(m.Alerts[i]) { // not required
			continue
		}

		if m.Alerts[i] != nil {
			if err := m.Alerts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Alerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTAlertListAlert) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTAlertListAlert) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTAlertListAlert) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTAlertListAlert) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTAlertListAlert) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTAlertListAlert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTAlertListAlert) UnmarshalBinary(b []byte) error {
	var res MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTAlertListAlert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
