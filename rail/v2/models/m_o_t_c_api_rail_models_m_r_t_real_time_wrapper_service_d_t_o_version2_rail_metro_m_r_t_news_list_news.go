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

// MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTNewsListNews MRTNewsList
//
// swagger:model MOTC.API.Rail.Models.MRTRealTimeWrapper[Service.DTO.Version2.Rail.Metro.MRTNewsList.News]
type MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTNewsListNews struct {

	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode"`

	// 資料總筆數
	Count int64 `json:"Count,omitempty"`

	// 資料(陣列)
	// Required: true
	Newses []*ServiceDTOVersion2RailMetroMRTNewsListNews `json:"Newses"`

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

// Validate validates this m o t c API rail models m r t real time wrapper service d t o version2 rail metro m r t news list news
func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTNewsListNews) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewses(formats); err != nil {
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

func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTNewsListNews) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTNewsListNews) validateNewses(formats strfmt.Registry) error {

	if err := validate.Required("Newses", "body", m.Newses); err != nil {
		return err
	}

	for i := 0; i < len(m.Newses); i++ {
		if swag.IsZero(m.Newses[i]) { // not required
			continue
		}

		if m.Newses[i] != nil {
			if err := m.Newses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Newses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTNewsListNews) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTNewsListNews) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTNewsListNews) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTNewsListNews) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTNewsListNews) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTNewsListNews) UnmarshalBinary(b []byte) error {
	var res MOTCAPIRailModelsMRTRealTimeWrapperServiceDTOVersion2RailMetroMRTNewsListNews
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
