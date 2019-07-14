// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRAStation TRAStationList
// swagger:model MOTC.API.Rail.Models.TRABaseWrapper[Service.DTO.Version3.Rail.TRA.Station]
type MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRAStation struct {

	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode"`

	// 資料總筆數
	Count int64 `json:"Count,omitempty"`

	// 來源端平台資料更新週期(秒)
	// Required: true
	SrcUpdateInterval *int32 `json:"SrcUpdateInterval"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// 資料(陣列)
	// Required: true
	Stations []*ServiceDTOVersion3RailTRAStation `json:"Stations"`

	// 本平台資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	UpdateTime string `json:"UpdateTime,omitempty"`
}

// Validate validates this m o t c API rail models t r a base wrapper service d t o version3 rail t r a station
func (m *MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRAStation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRAStation) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRAStation) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRAStation) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRAStation) validateStations(formats strfmt.Registry) error {

	if err := validate.Required("Stations", "body", m.Stations); err != nil {
		return err
	}

	for i := 0; i < len(m.Stations); i++ {
		if swag.IsZero(m.Stations[i]) { // not required
			continue
		}

		if m.Stations[i] != nil {
			if err := m.Stations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Stations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRAStation) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRAStation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRAStation) UnmarshalBinary(b []byte) error {
	var res MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRAStation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
