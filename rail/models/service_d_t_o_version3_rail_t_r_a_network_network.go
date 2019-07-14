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

// ServiceDTOVersion3RailTRANetworkNetwork Network
//
// 臺鐵路網資料
// swagger:model Service.DTO.Version3.Rail.TRA.Network.Network
type ServiceDTOVersion3RailTRANetworkNetwork struct {

	// 臺鐵路線資訊
	// Required: true
	Lines []*ServiceDTOVersion3RailTRANetworkLine `json:"Lines"`

	// 臺鐵路網代碼
	// Required: true
	NetworkID *string `json:"NetworkID"`

	// MapNameType
	//
	// 臺鐵路網圖網址URL
	// Required: true
	NetworkMapURL *ServiceDTOVersion3RailTRANetworkMapNameType `json:"NetworkMapURL"`

	// NameType
	//
	// 臺鐵路網名稱
	// Required: true
	NetworkName *ServiceDTOVersion3BaseNameType `json:"NetworkName"`

	// 營運業者代碼
	// Required: true
	OperatorCode *string `json:"OperatorCode"`

	// NameType
	//
	// 營運業者名稱
	// Required: true
	OperatorName *ServiceDTOVersion3BaseNameType `json:"OperatorName"`
}

// Validate validates this service d t o version3 rail t r a network network
func (m *ServiceDTOVersion3RailTRANetworkNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkMapURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3RailTRANetworkNetwork) validateLines(formats strfmt.Registry) error {

	if err := validate.Required("Lines", "body", m.Lines); err != nil {
		return err
	}

	for i := 0; i < len(m.Lines); i++ {
		if swag.IsZero(m.Lines[i]) { // not required
			continue
		}

		if m.Lines[i] != nil {
			if err := m.Lines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Lines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceDTOVersion3RailTRANetworkNetwork) validateNetworkID(formats strfmt.Registry) error {

	if err := validate.Required("NetworkID", "body", m.NetworkID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRANetworkNetwork) validateNetworkMapURL(formats strfmt.Registry) error {

	if err := validate.Required("NetworkMapURL", "body", m.NetworkMapURL); err != nil {
		return err
	}

	if m.NetworkMapURL != nil {
		if err := m.NetworkMapURL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NetworkMapURL")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRANetworkNetwork) validateNetworkName(formats strfmt.Registry) error {

	if err := validate.Required("NetworkName", "body", m.NetworkName); err != nil {
		return err
	}

	if m.NetworkName != nil {
		if err := m.NetworkName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NetworkName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRANetworkNetwork) validateOperatorCode(formats strfmt.Registry) error {

	if err := validate.Required("OperatorCode", "body", m.OperatorCode); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRANetworkNetwork) validateOperatorName(formats strfmt.Registry) error {

	if err := validate.Required("OperatorName", "body", m.OperatorName); err != nil {
		return err
	}

	if m.OperatorName != nil {
		if err := m.OperatorName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OperatorName")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRANetworkNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRANetworkNetwork) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3RailTRANetworkNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
