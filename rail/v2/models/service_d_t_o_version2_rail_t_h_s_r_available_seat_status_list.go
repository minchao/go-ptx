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

// ServiceDTOVersion2RailTHSRAvailableSeatStatusList AvailableSeatStatusList
//
// 高鐵對號座位狀態看板資料
// swagger:model Service.DTO.Version2.Rail.THSR.AvailableSeatStatusList
type ServiceDTOVersion2RailTHSRAvailableSeatStatusList struct {

	// 對號座位狀態資訊(依高鐵規定若營運狀態有異常狀況時，剩餘座位資訊將停留在最後正常運行時間之狀態不做更新，實際資料請參考高鐵各車站現場對號座剩餘座位資訊看板)
	// Required: true
	AvailableSeats []*ServiceDTOVersion2RailTHSRAvailableSeat `json:"AvailableSeats"`

	// DateTime
	//
	// 更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this service d t o version2 rail t h s r available seat status list
func (m *ServiceDTOVersion2RailTHSRAvailableSeatStatusList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableSeats(formats); err != nil {
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

func (m *ServiceDTOVersion2RailTHSRAvailableSeatStatusList) validateAvailableSeats(formats strfmt.Registry) error {

	if err := validate.Required("AvailableSeats", "body", m.AvailableSeats); err != nil {
		return err
	}

	for i := 0; i < len(m.AvailableSeats); i++ {
		if swag.IsZero(m.AvailableSeats[i]) { // not required
			continue
		}

		if m.AvailableSeats[i] != nil {
			if err := m.AvailableSeats[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AvailableSeats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceDTOVersion2RailTHSRAvailableSeatStatusList) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTHSRAvailableSeatStatusList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTHSRAvailableSeatStatusList) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailTHSRAvailableSeatStatusList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
