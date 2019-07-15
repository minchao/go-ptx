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

// ServiceDTOVersion2RailMetroS2STravelTime S2STravelTime
//
// 列車站間運行時間基本資料
// swagger:model Service.DTO.Version2.Rail.Metro.S2STravelTime
type ServiceDTOVersion2RailMetroS2STravelTime struct {

	// 營運路線所屬之路線代碼
	// Required: true
	LineID *string `json:"LineID"`

	// 營運路線所屬之路線編號
	LineNo string `json:"LineNo,omitempty"`

	// 營運路線代碼
	// Required: true
	RouteID *string `json:"RouteID"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// 車種(0:不分車種, 1:普通車, 2:直達車)
	TrainType int32 `json:"TrainType,omitempty"`

	// 站間運行時間資訊
	// Required: true
	TravelTimes []*ServiceDTOVersion2RailMetroSubClassTravelTime `json:"TravelTimes"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this service d t o version2 rail metro s2 s travel time
func (m *ServiceDTOVersion2RailMetroS2STravelTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTravelTimes(formats); err != nil {
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

func (m *ServiceDTOVersion2RailMetroS2STravelTime) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroS2STravelTime) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroS2STravelTime) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroS2STravelTime) validateTravelTimes(formats strfmt.Registry) error {

	if err := validate.Required("TravelTimes", "body", m.TravelTimes); err != nil {
		return err
	}

	for i := 0; i < len(m.TravelTimes); i++ {
		if swag.IsZero(m.TravelTimes[i]) { // not required
			continue
		}

		if m.TravelTimes[i] != nil {
			if err := m.TravelTimes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TravelTimes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroS2STravelTime) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroS2STravelTime) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroS2STravelTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroS2STravelTime) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailMetroS2STravelTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
