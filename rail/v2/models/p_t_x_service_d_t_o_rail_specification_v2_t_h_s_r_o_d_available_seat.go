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

// PTXServiceDTORailSpecificationV2THSRODAvailableSeat AvailableSeat
//
// 高鐵站間對號座位狀態資訊
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.THSR.OD.AvailableSeat
type PTXServiceDTORailSpecificationV2THSRODAvailableSeat struct {

	// String
	//
	// 商務席剩餘座位狀態
	// Required: true
	BusinessSeatStatus *string `json:"BusinessSeatStatus" xml:"BusinessSeatStatus"`

	// String
	//
	// 迄站車站簡碼(訂票系統用)
	// Required: true
	DestinationStationCode *string `json:"DestinationStationCode" xml:"DestinationStationCode"`

	// String
	//
	// 迄點車站代碼
	// Required: true
	DestinationStationID *string `json:"DestinationStationID" xml:"DestinationStationID"`

	// NameType
	//
	// 迄點車站名稱
	// Required: true
	DestinationStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"DestinationStationName" xml:"NameType"`

	// Int32
	//
	// 行駛方向
	// Required: true
	Direction *int32 `json:"Direction"`

	// String
	//
	// 起站車站簡碼(訂票系統用)
	// Required: true
	OriginStationCode *string `json:"OriginStationCode" xml:"OriginStationCode"`

	// String
	//
	// 起點車站代碼
	// Required: true
	OriginStationID *string `json:"OriginStationID" xml:"OriginStationID"`

	// NameType
	//
	// 起點車站名稱
	// Required: true
	OriginStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"OriginStationName" xml:"NameType"`

	// String
	//
	// 標準席剩餘座位狀態
	// Required: true
	StandardSeatStatus *string `json:"StandardSeatStatus" xml:"StandardSeatStatus"`

	// String
	//
	// 車次號碼
	// Required: true
	TrainNo *string `json:"TrainNo" xml:"TrainNo"`
}

// Validate validates this p t x service d t o rail specification v2 t h s r o d available seat
func (m *PTXServiceDTORailSpecificationV2THSRODAvailableSeat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBusinessSeatStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationStationCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginStationCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardSeatStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainNo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRODAvailableSeat) validateBusinessSeatStatus(formats strfmt.Registry) error {

	if err := validate.Required("BusinessSeatStatus", "body", m.BusinessSeatStatus); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRODAvailableSeat) validateDestinationStationCode(formats strfmt.Registry) error {

	if err := validate.Required("DestinationStationCode", "body", m.DestinationStationCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRODAvailableSeat) validateDestinationStationID(formats strfmt.Registry) error {

	if err := validate.Required("DestinationStationID", "body", m.DestinationStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRODAvailableSeat) validateDestinationStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRODAvailableSeat) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRODAvailableSeat) validateOriginStationCode(formats strfmt.Registry) error {

	if err := validate.Required("OriginStationCode", "body", m.OriginStationCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRODAvailableSeat) validateOriginStationID(formats strfmt.Registry) error {

	if err := validate.Required("OriginStationID", "body", m.OriginStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRODAvailableSeat) validateOriginStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRODAvailableSeat) validateStandardSeatStatus(formats strfmt.Registry) error {

	if err := validate.Required("StandardSeatStatus", "body", m.StandardSeatStatus); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRODAvailableSeat) validateTrainNo(formats strfmt.Registry) error {

	if err := validate.Required("TrainNo", "body", m.TrainNo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 t h s r o d available seat based on the context it is used
func (m *PTXServiceDTORailSpecificationV2THSRODAvailableSeat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestinationStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRODAvailableSeat) contextValidateDestinationStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRODAvailableSeat) contextValidateOriginStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2THSRODAvailableSeat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2THSRODAvailableSeat) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2THSRODAvailableSeat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
