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

// PTXServiceDTOTourismSpecificationV2BusRoute BusRoute
//
// 路線資料型別
//
// swagger:model PTX.Service.DTO.Tourism.Specification.V2.BusRoute
type PTXServiceDTOTourismSpecificationV2BusRoute struct {

	// String
	//
	// 業管單位代碼
	// Required: true
	AuthorityID *string `json:"AuthorityID" xml:"AuthorityID"`

	// Int32
	//
	// 公車路線類別 : [11:'市區公車',12:'公路客運',13:'國道客運',14:'接駁車']
	// Required: true
	BusRouteType *int64 `json:"BusRouteType"`

	// String
	//
	// 路線權管所屬縣市(相當於市區公車API的City參數)[若為公路/國道客運路線則為空值]
	City string `json:"City,omitempty" xml:"City,omitempty"`

	// String
	//
	// 路線權管所屬縣市之代碼(國際ISO 3166-2 三碼城市代碼)[若為公路/國道客運路線則為空值]
	CityCode string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`

	// String
	//
	// 起站英文名稱
	DepartureStopNameEn string `json:"DepartureStopNameEn,omitempty" xml:"DepartureStopNameEn,omitempty"`

	// String
	//
	// 起站中文名稱
	DepartureStopNameZh string `json:"DepartureStopNameZh,omitempty" xml:"DepartureStopNameZh,omitempty"`

	// String
	//
	// 終點站英文名稱
	DestinationStopNameEn string `json:"DestinationStopNameEn,omitempty" xml:"DestinationStopNameEn,omitempty"`

	// String
	//
	// 終點站中文名稱
	DestinationStopNameZh string `json:"DestinationStopNameZh,omitempty" xml:"DestinationStopNameZh,omitempty"`

	// String
	//
	// 收費緩衝區英文敘述
	FareBufferZoneDescriptionEn string `json:"FareBufferZoneDescriptionEn,omitempty" xml:"FareBufferZoneDescriptionEn,omitempty"`

	// String
	//
	// 收費緩衝區中文敘述
	FareBufferZoneDescriptionZh string `json:"FareBufferZoneDescriptionZh,omitempty" xml:"FareBufferZoneDescriptionZh,omitempty"`

	// Boolean
	//
	// 實際上是否有多條附屬路線。(此欄位值與SubRoutes結構並無強烈的絕對關聯。詳細說明請參閱swagger上方的【資料服務使用注意事項】)
	// Required: true
	HasSubRoutes *bool `json:"HasSubRoutes"`

	// String
	//
	// 路線公車動態資訊網址
	LiveBusURL string `json:"LiveBusUrl,omitempty" xml:"LiveBusUrl,omitempty"`

	// Array
	//
	// 營運業者代碼
	// Required: true
	OperatorIDs []string "json:\"OperatorIDs\" xml:\"List`1\""

	// Array
	//
	// 營運業者
	// Required: true
	Operators []*PTXServiceDTOTourismSpecificationV2RouteOperator "json:\"Operators\" xml:\"List`1\""

	// String
	//
	// 資料提供平台代碼
	// Required: true
	ProviderID *string `json:"ProviderID" xml:"ProviderID"`

	// String
	//
	// 地區既用中之路線代碼(為原資料內碼)
	// Required: true
	RouteID *string `json:"RouteID" xml:"RouteID"`

	// String
	//
	// 路線簡圖網址
	RouteMapImageURL string `json:"RouteMapImageUrl,omitempty" xml:"RouteMapImageUrl,omitempty"`

	// String
	//
	// 路線唯一識別代碼，規則為 {業管機關代碼} + {RouteID}，其中 {業管機關代碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	RouteUID *string `json:"RouteUID" xml:"RouteUID"`

	// String
	//
	// 路線資訊說明網址
	RouteURL string `json:"RouteUrl,omitempty" xml:"RouteUrl,omitempty"`

	// Array
	//
	// 附屬路線資料(如果原始資料並無提供附屬路線ID，而本平台基於跨來源資料之一致性，會以SubRouteID=RouteID產製一份相對應的附屬路線資料(若有去返程，則會有兩筆))
	SubRoutes []*PTXServiceDTOTourismSpecificationV2BusSubRoute "json:\"SubRoutes\" xml:\"List`1\""

	// NameType
	//
	// 台灣好行路線名稱
	// Required: true
	TaiwanTripName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"TaiwanTripName" xml:"NameType"`

	// String
	//
	// 票價英文敘述
	TicketPriceDescriptionEn string `json:"TicketPriceDescriptionEn,omitempty" xml:"TicketPriceDescriptionEn,omitempty"`

	// String
	//
	// 票價中文敘述
	TicketPriceDescriptionZh string `json:"TicketPriceDescriptionZh,omitempty" xml:"TicketPriceDescriptionZh,omitempty"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o tourism specification v2 bus route
func (m *PTXServiceDTOTourismSpecificationV2BusRoute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBusRouteType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasSubRoutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorIDs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubRoutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaiwanTripName(formats); err != nil {
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

func (m *PTXServiceDTOTourismSpecificationV2BusRoute) validateAuthorityID(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityID", "body", m.AuthorityID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusRoute) validateBusRouteType(formats strfmt.Registry) error {

	if err := validate.Required("BusRouteType", "body", m.BusRouteType); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusRoute) validateHasSubRoutes(formats strfmt.Registry) error {

	if err := validate.Required("HasSubRoutes", "body", m.HasSubRoutes); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusRoute) validateOperatorIDs(formats strfmt.Registry) error {

	if err := validate.Required("OperatorIDs", "body", m.OperatorIDs); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusRoute) validateOperators(formats strfmt.Registry) error {

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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Operators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusRoute) validateProviderID(formats strfmt.Registry) error {

	if err := validate.Required("ProviderID", "body", m.ProviderID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusRoute) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusRoute) validateRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("RouteUID", "body", m.RouteUID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusRoute) validateSubRoutes(formats strfmt.Registry) error {
	if swag.IsZero(m.SubRoutes) { // not required
		return nil
	}

	for i := 0; i < len(m.SubRoutes); i++ {
		if swag.IsZero(m.SubRoutes[i]) { // not required
			continue
		}

		if m.SubRoutes[i] != nil {
			if err := m.SubRoutes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SubRoutes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SubRoutes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusRoute) validateTaiwanTripName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusRoute) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o tourism specification v2 bus route based on the context it is used
func (m *PTXServiceDTOTourismSpecificationV2BusRoute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubRoutes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaiwanTripName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusRoute) contextValidateOperators(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Operators); i++ {

		if m.Operators[i] != nil {
			if err := m.Operators[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Operators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Operators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusRoute) contextValidateSubRoutes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SubRoutes); i++ {

		if m.SubRoutes[i] != nil {
			if err := m.SubRoutes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SubRoutes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SubRoutes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusRoute) contextValidateTaiwanTripName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2BusRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2BusRoute) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOTourismSpecificationV2BusRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
