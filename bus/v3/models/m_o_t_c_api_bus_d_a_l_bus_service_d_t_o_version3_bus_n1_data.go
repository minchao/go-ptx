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

// MOTCAPIBusDALBusServiceDTOVersion3BusN1Data BusN1DataList
//
// swagger:model MOTC.API.Bus.DAL.Bus[Service.DTO.Version3.Bus.N1Data]
type MOTCAPIBusDALBusServiceDTOVersion3BusN1Data struct {

	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode"`

	// 資料總筆數
	Count int64 `json:"Count,omitempty"`

	// 資料(陣列)
	// Required: true
	N1Datas []*ServiceDTOVersion3BusN1Data `json:"N1Datas"`

	// 來源端平台資料更新週期(秒)['-1: 不定期更新']
	// Required: true
	SrcUpdateInterval *int32 `json:"SrcUpdateInterval"`

	// DateTime
	//
	// [來源端平臺]資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// [平臺]資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// DateTime
	//
	// [平臺]資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// 資料版本編號
	VersionID int32 `json:"VersionID,omitempty"`
}

// Validate validates this m o t c API bus d a l bus service d t o version3 bus n1 data
func (m *MOTCAPIBusDALBusServiceDTOVersion3BusN1Data) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateN1Datas(formats); err != nil {
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

func (m *MOTCAPIBusDALBusServiceDTOVersion3BusN1Data) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIBusDALBusServiceDTOVersion3BusN1Data) validateN1Datas(formats strfmt.Registry) error {

	if err := validate.Required("N1Datas", "body", m.N1Datas); err != nil {
		return err
	}

	for i := 0; i < len(m.N1Datas); i++ {
		if swag.IsZero(m.N1Datas[i]) { // not required
			continue
		}

		if m.N1Datas[i] != nil {
			if err := m.N1Datas[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("N1Datas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MOTCAPIBusDALBusServiceDTOVersion3BusN1Data) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIBusDALBusServiceDTOVersion3BusN1Data) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIBusDALBusServiceDTOVersion3BusN1Data) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIBusDALBusServiceDTOVersion3BusN1Data) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MOTCAPIBusDALBusServiceDTOVersion3BusN1Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MOTCAPIBusDALBusServiceDTOVersion3BusN1Data) UnmarshalBinary(b []byte) error {
	var res MOTCAPIBusDALBusServiceDTOVersion3BusN1Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
