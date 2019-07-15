// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion2RailMetroSubClassStation Station
//
// 路線車站資訊
// swagger:model Service.DTO.Version2.Rail.Metro.SubClass.Station
type ServiceDTOVersion2RailMetroSubClassStation struct {

	// 已累積之里程距離(公里)
	CumulativeDistance float32 `json:"CumulativeDistance,omitempty"`

	// 站序
	// Required: true
	Sequence *int32 `json:"Sequence"`

	// 車站代碼
	// Required: true
	StationID *string `json:"StationID"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName *ServiceDTOVersion2BaseNameType `json:"StationName"`
}

// Validate validates this service d t o version2 rail metro sub class station
func (m *ServiceDTOVersion2RailMetroSubClassStation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSequence(formats); err != nil {
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

func (m *ServiceDTOVersion2RailMetroSubClassStation) validateSequence(formats strfmt.Registry) error {

	if err := validate.Required("Sequence", "body", m.Sequence); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroSubClassStation) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroSubClassStation) validateStationName(formats strfmt.Registry) error {

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
func (m *ServiceDTOVersion2RailMetroSubClassStation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroSubClassStation) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailMetroSubClassStation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
