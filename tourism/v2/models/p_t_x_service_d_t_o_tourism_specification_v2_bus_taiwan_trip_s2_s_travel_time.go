// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTOTourismSpecificationV2BusTaiwanTripS2STravelTime BusTaiwanTripS2STravelTime
//
// swagger:model PTX.Service.DTO.Tourism.Specification.V2.BusTaiwanTripS2STravelTime
type PTXServiceDTOTourismSpecificationV2BusTaiwanTripS2STravelTime struct {

	// String
	//
	// 路線代碼
	// Required: true
	RouteID *string `json:"RouteID" xml:"RouteID"`

	// String
	//
	// 路線唯一識別代碼，規則為 {業管機關簡碼} + {RouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	RouteUID *string `json:"RouteUID" xml:"RouteUID"`

	// String
	//
	// 附屬路線代碼
	SubRouteID string `json:"SubRouteID,omitempty" xml:"SubRouteID,omitempty"`

	// String
	//
	// 附屬路線唯一識別代碼，規則為 {業管機關簡碼} + {SubRouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	SubRouteUID string `json:"SubRouteUID,omitempty" xml:"SubRouteUID,omitempty"`

	// Array
	//
	// 站間運行時間資訊
	// Required: true
	TravelTimes []*PTXServiceDTOTourismSpecificationV2BusTaiwanTripS2STravelTimeTravelTime "json:\"TravelTimes\" xml:\"List`1\""

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o tourism specification v2 bus taiwan trip s2 s travel time
func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripS2STravelTime) Validate(formats strfmt.Registry) error {
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

func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripS2STravelTime) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripS2STravelTime) validateRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("RouteUID", "body", m.RouteUID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripS2STravelTime) validateTravelTimes(formats strfmt.Registry) error {

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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("TravelTimes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripS2STravelTime) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o tourism specification v2 bus taiwan trip s2 s travel time based on the context it is used
func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripS2STravelTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTravelTimes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripS2STravelTime) contextValidateTravelTimes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TravelTimes); i++ {

		if m.TravelTimes[i] != nil {
			if err := m.TravelTimes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TravelTimes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("TravelTimes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripS2STravelTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripS2STravelTime) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOTourismSpecificationV2BusTaiwanTripS2STravelTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
