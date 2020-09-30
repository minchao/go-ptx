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

// PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility TRAStationFacilityList
//
// swagger:model PTX.API.Rail.Model.TRABaseWrapper[PTX.Service.DTO.Rail.Specification.V3.TRA.StationFacility.StationFacility]
type PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility struct {

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode"`

	// 資料總筆數
	Count int64 `json:"Count,omitempty"`

	// Int32
	//
	// 來源端平台資料更新週期(秒)['-1: 不定期更新']
	// Required: true
	SrcUpdateInterval *int32 `json:"SrcUpdateInterval"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// Array
	//
	// 資料(陣列)
	// Required: true
	StationFacilities []*PTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility `json:"StationFacilities"`

	// Int32
	//
	// 本平台資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this p t x API rail model t r a base wrapper p t x service d t o rail specification v3 t r a station facility station facility
func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility) Validate(formats strfmt.Registry) error {
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

	if err := m.validateStationFacilities(formats); err != nil {
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

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility) validateStationFacilities(formats strfmt.Registry) error {

	if err := validate.Required("StationFacilities", "body", m.StationFacilities); err != nil {
		return err
	}

	for i := 0; i < len(m.StationFacilities); i++ {
		if swag.IsZero(m.StationFacilities[i]) { // not required
			continue
		}

		if m.StationFacilities[i] != nil {
			if err := m.StationFacilities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StationFacilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility) UnmarshalBinary(b []byte) error {
	var res PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}