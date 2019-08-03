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

// ServiceDTOVersion2TaiwanTripBusBusTaiwanTripS2STravelTime BusTaiwanTripS2STravelTime
// swagger:model Service.DTO.Version2.TaiwanTripBus.BusTaiwanTripS2STravelTime
type ServiceDTOVersion2TaiwanTripBusBusTaiwanTripS2STravelTime struct {

	// 路線代碼
	// Required: true
	RouteID *string `json:"RouteID"`

	// 路線唯一識別代碼，規則為 {業管機關簡碼} + {RouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	RouteUID *string `json:"RouteUID"`

	// 附屬路線代碼
	SubRouteID string `json:"SubRouteID,omitempty"`

	// 附屬路線唯一識別代碼，規則為 {業管機關簡碼} + {SubRouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	SubRouteUID string `json:"SubRouteUID,omitempty"`

	// 站間運行時間資訊
	// Required: true
	TravelTimes []*ServiceDTOVersion2TaiwanTripBusBusTaiwanTripS2STravelTimeTravelTime `json:"TravelTimes"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this service d t o version2 taiwan trip bus bus taiwan trip s2 s travel time
func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripS2STravelTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTravelTimes(formats); err != nil {
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

func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripS2STravelTime) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripS2STravelTime) validateRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("RouteUID", "body", m.RouteUID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripS2STravelTime) validateTravelTimes(formats strfmt.Registry) error {

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

func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripS2STravelTime) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripS2STravelTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripS2STravelTime) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2TaiwanTripBusBusTaiwanTripS2STravelTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
