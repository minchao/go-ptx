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

// ServiceDTOVersion2RailMetroFrequency Frequency
//
// 路線發車班距頻率資料
// swagger:model Service.DTO.Version2.Rail.Metro.Frequency
type ServiceDTOVersion2RailMetroFrequency struct {

	// 班距頻率資訊
	// Required: true
	Headways []*ServiceDTOVersion2RailMetroSubClassHeadway `json:"Headways"`

	// 營運路線所屬之路線代碼
	// Required: true
	LineID *string `json:"LineID"`

	// 營運路線所屬之路線編號
	LineNo string `json:"LineNo,omitempty"`

	// OperationTime
	//
	// 營運時間資訊
	// Required: true
	OperationTime *ServiceDTOVersion2RailMetroSubClassOperationTime `json:"OperationTime"`

	// 營運路線代碼
	// Required: true
	RouteID *string `json:"RouteID"`

	// ServiceDays
	//
	// 服務日型態
	// Required: true
	ServiceDays *ServiceDTOVersion2RailMetroSubClassServiceDays `json:"ServiceDays"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// 車種(0:不分車種, 1:普通車, 2:直達車)
	TrainType int32 `json:"TrainType,omitempty"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this service d t o version2 rail metro frequency
func (m *ServiceDTOVersion2RailMetroFrequency) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeadways(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceDays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion2RailMetroFrequency) validateHeadways(formats strfmt.Registry) error {

	if err := validate.Required("Headways", "body", m.Headways); err != nil {
		return err
	}

	for i := 0; i < len(m.Headways); i++ {
		if swag.IsZero(m.Headways[i]) { // not required
			continue
		}

		if m.Headways[i] != nil {
			if err := m.Headways[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Headways" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFrequency) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFrequency) validateOperationTime(formats strfmt.Registry) error {

	if err := validate.Required("OperationTime", "body", m.OperationTime); err != nil {
		return err
	}

	if m.OperationTime != nil {
		if err := m.OperationTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OperationTime")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFrequency) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFrequency) validateServiceDays(formats strfmt.Registry) error {

	if err := validate.Required("ServiceDays", "body", m.ServiceDays); err != nil {
		return err
	}

	if m.ServiceDays != nil {
		if err := m.ServiceDays.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ServiceDays")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFrequency) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFrequency) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFrequency) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroFrequency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroFrequency) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailMetroFrequency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
