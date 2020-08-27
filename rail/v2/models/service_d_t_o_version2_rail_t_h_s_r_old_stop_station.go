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

// ServiceDTOVersion2RailTHSROldStopStation StopStation
//
// 車次停靠站點組合
//
// swagger:model Service.DTO.Version2.Rail.THSR.Old.StopStation
type ServiceDTOVersion2RailTHSROldStopStation struct {

	// 商務席剩餘座位狀態 =  ['Available: 尚有座位' or 'Limited: 座位有限' or 'Full: 已無座位']
	// Required: true
	BusinessSeatStatus *string `json:"BusinessSeatStatus"`

	// 標準席剩餘座位狀態 = ['Available: 尚有座位' or 'Limited: 座位有限' or 'Full: 已無座位']
	// Required: true
	StandardSeatStatus *string `json:"StandardSeatStatus"`

	// 車站代碼
	// Required: true
	StationID *string `json:"StationID"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName *ServiceDTOVersion2BaseNameType `json:"StationName"`
}

// Validate validates this service d t o version2 rail t h s r old stop station
func (m *ServiceDTOVersion2RailTHSROldStopStation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBusinessSeatStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardSeatStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion2RailTHSROldStopStation) validateBusinessSeatStatus(formats strfmt.Registry) error {

	if err := validate.Required("BusinessSeatStatus", "body", m.BusinessSeatStatus); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTHSROldStopStation) validateStandardSeatStatus(formats strfmt.Registry) error {

	if err := validate.Required("StandardSeatStatus", "body", m.StandardSeatStatus); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTHSROldStopStation) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTHSROldStopStation) validateStationName(formats strfmt.Registry) error {

	if err := validate.Required("StationName", "body", m.StationName); err != nil {
		return err
	}

	if m.StationName != nil {
		if err := m.StationName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StationName")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTHSROldStopStation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTHSROldStopStation) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailTHSROldStopStation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}