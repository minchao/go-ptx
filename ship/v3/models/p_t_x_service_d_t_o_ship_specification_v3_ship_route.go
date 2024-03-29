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

// PTXServiceDTOShipSpecificationV3ShipRoute ShipRoute
//
// swagger:model PTX.Service.DTO.Ship.Specification.V3.ShipRoute
type PTXServiceDTOShipSpecificationV3ShipRoute struct {

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode" xml:"AuthorityCode"`

	// String
	//
	// 航線描述
	// Required: true
	Description *string `json:"Description" xml:"Description"`

	// String
	//
	// 航線迄點港口代碼
	// Required: true
	EndPortID *string `json:"EndPortID" xml:"EndPortID"`

	// String
	//
	// 航線迄點港口名稱
	// Required: true
	EndPortName *string `json:"EndPortName" xml:"EndPortName"`

	// Array
	//
	// 營運業者
	// Required: true
	Operators []*PTXServiceDTOShipSpecificationV3Operators "json:\"Operators\" xml:\"List`1\""

	// Single
	//
	// 航線哩程
	// Required: true
	RouteDistance *float32 `json:"RouteDistance"`

	// String
	//
	// 航線代碼
	// Required: true
	RouteID *string `json:"RouteID" xml:"RouteID"`

	// String
	//
	// 航運路線簡圖網址
	RouteMapURL string `json:"RouteMapURL,omitempty" xml:"RouteMapURL,omitempty"`

	// NameType
	//
	// 航線名稱
	// Required: true
	RouteName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"RouteName" xml:"NameType"`

	// Int32
	//
	// 航線種類 : [1:'國內航線(渡輪、藍色公路)',2:'離島航線',3:'兩岸航線',4:'小三通航線',254:'其他']
	// Required: true
	RouteType *int64 `json:"RouteType"`

	// DateTime
	//
	// [來源端平臺]資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	SrcUpdateTime *strfmt.DateTime `json:"SrcUpdateTime"`

	// String
	//
	// 航線起點港口代碼
	// Required: true
	StartPortID *string `json:"StartPortID" xml:"StartPortID"`

	// String
	//
	// 航線起點港口名稱
	// Required: true
	StartPortName *string `json:"StartPortName" xml:"StartPortName"`

	// NameType
	//
	// 票價描述
	// Required: true
	TicketPriceDescription struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"TicketPriceDescription" xml:"NameType"`

	// DateTime
	//
	// [平臺]資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`

	// Array
	//
	// 船舶資訊
	Vessels []*PTXServiceDTOShipSpecificationV3Vessels "json:\"Vessels\" xml:\"List`1\""

	// String
	//
	// 氣象預報連結
	// Required: true
	WeatherURL *string `json:"WeatherURL" xml:"WeatherURL"`
}

// Validate validates this p t x service d t o ship specification v3 ship route
func (m *PTXServiceDTOShipSpecificationV3ShipRoute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndPortID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndPortName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteDistance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartPortID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartPortName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTicketPriceDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVessels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeatherURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) validateEndPortID(formats strfmt.Registry) error {

	if err := validate.Required("EndPortID", "body", m.EndPortID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) validateEndPortName(formats strfmt.Registry) error {

	if err := validate.Required("EndPortName", "body", m.EndPortName); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) validateOperators(formats strfmt.Registry) error {

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

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) validateRouteDistance(formats strfmt.Registry) error {

	if err := validate.Required("RouteDistance", "body", m.RouteDistance); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) validateRouteName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) validateRouteType(formats strfmt.Registry) error {

	if err := validate.Required("RouteType", "body", m.RouteType); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) validateStartPortID(formats strfmt.Registry) error {

	if err := validate.Required("StartPortID", "body", m.StartPortID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) validateStartPortName(formats strfmt.Registry) error {

	if err := validate.Required("StartPortName", "body", m.StartPortName); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) validateTicketPriceDescription(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) validateVessels(formats strfmt.Registry) error {
	if swag.IsZero(m.Vessels) { // not required
		return nil
	}

	for i := 0; i < len(m.Vessels); i++ {
		if swag.IsZero(m.Vessels[i]) { // not required
			continue
		}

		if m.Vessels[i] != nil {
			if err := m.Vessels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Vessels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Vessels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) validateWeatherURL(formats strfmt.Registry) error {

	if err := validate.Required("WeatherURL", "body", m.WeatherURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o ship specification v3 ship route based on the context it is used
func (m *PTXServiceDTOShipSpecificationV3ShipRoute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRouteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTicketPriceDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVessels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) contextValidateOperators(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) contextValidateRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) contextValidateTicketPriceDescription(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipRoute) contextValidateVessels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Vessels); i++ {

		if m.Vessels[i] != nil {
			if err := m.Vessels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Vessels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Vessels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3ShipRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3ShipRoute) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOShipSpecificationV3ShipRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
