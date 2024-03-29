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

// PTXAPIRailModelMRTRealTimeWrapperPTXServiceDTORailSpecificationV2MetroMRTNewsListNews MRTNewsList
//
// swagger:model PTX.API.Rail.Model.MRTRealTimeWrapper[PTX.Service.DTO.Rail.Specification.V2.Metro.MRTNewsList.News]
type PTXAPIRailModelMRTRealTimeWrapperPTXServiceDTORailSpecificationV2MetroMRTNewsListNews struct {

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode" xml:"AuthorityCode"`

	// 資料總筆數
	Count int64 `json:"Count,omitempty"`

	// Array
	//
	// 資料(陣列)
	// Required: true
	Newses []*PTXServiceDTORailSpecificationV2MetroMRTNewsListNews "json:\"Newses\" xml:\"List`1\""

	// Int32
	//
	// 來源端平台資料更新週期(秒)['-1: 不定期更新']
	// Required: true
	SrcUpdateInterval *int32 `json:"SrcUpdateInterval"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	SrcUpdateTime *strfmt.DateTime `json:"SrcUpdateTime"`

	// Int32
	//
	// 本平台資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x API rail model m r t real time wrapper p t x service d t o rail specification v2 metro m r t news list news
func (m *PTXAPIRailModelMRTRealTimeWrapperPTXServiceDTORailSpecificationV2MetroMRTNewsListNews) Validate(formats strfmt.Registry) error {
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

func (m *PTXAPIRailModelMRTRealTimeWrapperPTXServiceDTORailSpecificationV2MetroMRTNewsListNews) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelMRTRealTimeWrapperPTXServiceDTORailSpecificationV2MetroMRTNewsListNews) validateNewses(formats strfmt.Registry) error {

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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Newses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXAPIRailModelMRTRealTimeWrapperPTXServiceDTORailSpecificationV2MetroMRTNewsListNews) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelMRTRealTimeWrapperPTXServiceDTORailSpecificationV2MetroMRTNewsListNews) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelMRTRealTimeWrapperPTXServiceDTORailSpecificationV2MetroMRTNewsListNews) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelMRTRealTimeWrapperPTXServiceDTORailSpecificationV2MetroMRTNewsListNews) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x API rail model m r t real time wrapper p t x service d t o rail specification v2 metro m r t news list news based on the context it is used
func (m *PTXAPIRailModelMRTRealTimeWrapperPTXServiceDTORailSpecificationV2MetroMRTNewsListNews) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNewses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXAPIRailModelMRTRealTimeWrapperPTXServiceDTORailSpecificationV2MetroMRTNewsListNews) contextValidateNewses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Newses); i++ {

		if m.Newses[i] != nil {
			if err := m.Newses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Newses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Newses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXAPIRailModelMRTRealTimeWrapperPTXServiceDTORailSpecificationV2MetroMRTNewsListNews) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXAPIRailModelMRTRealTimeWrapperPTXServiceDTORailSpecificationV2MetroMRTNewsListNews) UnmarshalBinary(b []byte) error {
	var res PTXAPIRailModelMRTRealTimeWrapperPTXServiceDTORailSpecificationV2MetroMRTNewsListNews
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
