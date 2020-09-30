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

// PTXServiceDTOBusSpecificationV3Route Route
//
// 路線資料型別
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.Route
type PTXServiceDTOBusSpecificationV3Route struct {

	// NameType
	//
	// 路線起站名稱
	// Required: true
	DepartureStopName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"DepartureStopName"`

	// NameType
	//
	// 路線迄站名稱
	// Required: true
	DestinationStopName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"DestinationStopName"`

	// Stop
	//
	// 終點站牌
	// Required: true
	EndStop struct {
		PTXServiceDTOBusSpecificationV3RouteStop
	} `json:"EndStop"`

	// NameType
	//
	// 計費緩衝區敘述
	FareBufferZoneDescription struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"FareBufferZoneDescription,omitempty"`

	// String
	//
	// 收費資訊說明網址
	FareURL string `json:"FareURL,omitempty"`

	// Boolean
	//
	// 實際上是否有多條附屬路線。(此欄位值與SubRoutes結構並無強烈的絕對關聯。詳細說明請參閱swagger上方的【資料服務使用注意事項】)
	// Required: true
	HasSubRoutes *bool `json:"HasSubRoutes"`

	// Boolean
	//
	// 是否為環狀線
	// Required: true
	IsCircular *bool `json:"IsCircular"`

	// String
	//
	// 路線公車動態資訊網址
	LiveBusURL string `json:"LiveBusURL,omitempty"`

	// Array
	//
	// 營運業者
	// Required: true
	Operators []*PTXServiceDTOBusSpecificationV3RouteOperator `json:"Operators"`

	// String
	//
	// 路線描述
	RouteDescription string `json:"RouteDescription,omitempty"`

	// 路線旅行長度
	RouteDistance float32 `json:"RouteDistance,omitempty"`

	// String
	//
	// 地區既用中之路線代碼(為原資料內碼)
	// Required: true
	RouteID *string `json:"RouteID"`

	// NameType
	//
	// 路線長名稱
	RouteLongName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"RouteLongName,omitempty"`

	// String
	//
	// 路線簡圖網址
	RouteMapImageURL string `json:"RouteMapImageURL,omitempty"`

	// NameType
	//
	// 路線名稱
	// Required: true
	RouteName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"RouteName"`

	// integer
	//
	// 公車路線類別 : [11:'市區公車',12:'公路客運',13:'國道客運']
	// Required: true
	RouteType *int32 `json:"RouteType"`

	// String
	//
	// 路線唯一識別代碼，規則為 {業管機關簡碼} + {RouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	RouteUID *string `json:"RouteUID"`

	// String
	//
	// 路線資訊說明網址
	RouteURL string `json:"RouteURL,omitempty"`

	// ServiceType
	//
	// 公車路線服務類型
	// Required: true
	ServiceType struct {
		PTXServiceDTOBusSpecificationV3ServiceType
	} `json:"ServiceType"`

	// Stop
	//
	// 起始站牌
	// Required: true
	StartStop struct {
		PTXServiceDTOBusSpecificationV3RouteStop
	} `json:"StartStop"`

	// NameType
	//
	// 票價描述
	// Required: true
	TicketPriceDescription struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"TicketPriceDescription"`
}

// Validate validates this p t x service d t o bus specification v3 route
func (m *PTXServiceDTOBusSpecificationV3Route) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDepartureStopName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationStopName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndStop(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFareBufferZoneDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasSubRoutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsCircular(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteLongName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartStop(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTicketPriceDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Route) validateDepartureStopName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Route) validateDestinationStopName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Route) validateEndStop(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Route) validateFareBufferZoneDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.FareBufferZoneDescription) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Route) validateHasSubRoutes(formats strfmt.Registry) error {

	if err := validate.Required("HasSubRoutes", "body", m.HasSubRoutes); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Route) validateIsCircular(formats strfmt.Registry) error {

	if err := validate.Required("IsCircular", "body", m.IsCircular); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Route) validateOperators(formats strfmt.Registry) error {

	if err := validate.Required("Operators", "body", m.Operators); err != nil {
		return err
	}

	for i := 0; i < len(m.Operators); i++ {
		if swag.IsZero(m.Operators[i]) { // not required
			continue
		}

		if m.Operators[i] != nil {
			if err := m.Operators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Operators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Route) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Route) validateRouteLongName(formats strfmt.Registry) error {

	if swag.IsZero(m.RouteLongName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Route) validateRouteName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Route) validateRouteType(formats strfmt.Registry) error {

	if err := validate.Required("RouteType", "body", m.RouteType); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Route) validateRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("RouteUID", "body", m.RouteUID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Route) validateServiceType(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Route) validateStartStop(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3Route) validateTicketPriceDescription(formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3Route) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3Route) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3Route
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}