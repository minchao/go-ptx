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

// PTXServiceDTOBusSpecificationV2BusS2STravelTime BusS2STravelTime
//
// 站間旅行時間資料型別
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.BusS2STravelTime
type PTXServiceDTOBusSpecificationV2BusS2STravelTime struct {

	// Int32
	//
	// 去返程 : [0:'去程',1:'返程',2:'迴圈',255:'未知']
	// Required: true
	Direction *int64 `json:"Direction"`

	// String
	//
	// 地區既用中之路線代碼(為原資料內碼)
	// Required: true
	RouteID *string `json:"RouteID" xml:"RouteID"`

	// String
	//
	// 路線唯一識別代碼，規則為 {業管機關簡碼} + {RouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	RouteUID *string `json:"RouteUID" xml:"RouteUID"`

	// String
	//
	// 地區既用中之附屬路線代碼(為原資料內碼)
	// Required: true
	SubRouteID *string `json:"SubRouteID" xml:"SubRouteID"`

	// String
	//
	// 附屬路線唯一識別代碼，規則為 {業管機關簡碼} + {SubRouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	SubRouteUID *string `json:"SubRouteUID" xml:"SubRouteUID"`

	// Array
	//
	// 旅行時間資訊
	// Required: true
	TravelTimes []*PTXServiceDTOBusSpecificationV2ServiceTime "json:\"TravelTimes\" xml:\"List`1\""

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o bus specification v2 bus s2 s travel time
func (m *PTXServiceDTOBusSpecificationV2BusS2STravelTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubRouteUID(formats); err != nil {
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

func (m *PTXServiceDTOBusSpecificationV2BusS2STravelTime) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusS2STravelTime) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusS2STravelTime) validateRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("RouteUID", "body", m.RouteUID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusS2STravelTime) validateSubRouteID(formats strfmt.Registry) error {

	if err := validate.Required("SubRouteID", "body", m.SubRouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusS2STravelTime) validateSubRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("SubRouteUID", "body", m.SubRouteUID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusS2STravelTime) validateTravelTimes(formats strfmt.Registry) error {

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

func (m *PTXServiceDTOBusSpecificationV2BusS2STravelTime) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v2 bus s2 s travel time based on the context it is used
func (m *PTXServiceDTOBusSpecificationV2BusS2STravelTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTravelTimes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusS2STravelTime) contextValidateTravelTimes(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PTXServiceDTOBusSpecificationV2BusS2STravelTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusS2STravelTime) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2BusS2STravelTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
