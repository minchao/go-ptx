// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer BusTransfer
//
// 公車運具轉乘資訊
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.StationTransfer.BusTransfer
type PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer struct {

	// String
	//
	// 轉乘描述
	Description string `json:"Description,omitempty" xml:"Description,omitempty"`

	// String
	//
	// 轉乘公車開往方向
	// Required: true
	Destination *string `json:"Destination" xml:"Destination"`

	// String
	//
	// 轉乘樓層
	FloorLevel string `json:"FloorLevel,omitempty" xml:"FloorLevel,omitempty"`

	// 是否為站內或站外轉乘
	IsOnSiteTransfer bool `json:"IsOnSiteTransfer,omitempty"`

	// 最短轉乘時間
	MinTransferTime float64 `json:"MinTransferTime,omitempty"`

	// String
	//
	// 運具種類代碼
	// Required: true
	Mode *string `json:"Mode" xml:"Mode"`

	// String
	//
	// 公車營運業者簡碼
	// Required: true
	OperatorCode *string `json:"OperatorCode" xml:"OperatorCode"`

	// NameType
	//
	// 公車營運業者名稱
	// Required: true
	OperatorName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"OperatorName" xml:"NameType"`

	// String
	//
	// 轉乘公車所在月台
	Platform string `json:"Platform,omitempty" xml:"Platform,omitempty"`

	// String
	//
	// 轉乘公車路線代碼
	// Required: true
	RouteID *string `json:"RouteID" xml:"RouteID"`

	// NameType
	//
	// 轉乘公車路線名稱
	// Required: true
	RouteName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"RouteName" xml:"NameType"`

	// String
	//
	// 轉乘公車站牌代碼
	// Required: true
	StopID *string `json:"StopID" xml:"StopID"`

	// NameType
	//
	// 轉乘公車站牌名稱
	// Required: true
	StopName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"StopName" xml:"NameType"`
}

// Validate validates this p t x service d t o rail specification v3 t r a station transfer bus transfer
func (m *PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("Destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("Mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer) validateOperatorCode(formats strfmt.Registry) error {

	if err := validate.Required("OperatorCode", "body", m.OperatorCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer) validateOperatorName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer) validateRouteName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer) validateStopID(formats strfmt.Registry) error {

	if err := validate.Required("StopID", "body", m.StopID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer) validateStopName(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a station transfer bus transfer based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperatorName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRouteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStopName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer) contextValidateOperatorName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer) contextValidateRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer) contextValidateStopName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRAStationTransferBusTransfer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
