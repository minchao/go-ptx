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

// PTXServiceDTORailSpecificationV2THSROldAvailableSeat AvailableSeat
//
// 高鐵對號座位狀態資訊
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.THSR.Old.AvailableSeat
type PTXServiceDTORailSpecificationV2THSROldAvailableSeat struct {

	// String
	//
	// 發車時間(格式: HH:mm)
	// Required: true
	DepartureTime *string `json:"DepartureTime"`

	// integer
	//
	// 方向 : [0:'南下',1:'北上']
	// Required: true
	Direction *int32 `json:"Direction"`

	// String
	//
	// 終點車站代碼
	// Required: true
	EndingStationID *string `json:"EndingStationID"`

	// NameType
	//
	// 終點車站名稱
	// Required: true
	EndingStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"EndingStationName"`

	// DateTime
	//
	// 來源平台更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// String
	//
	// 查詢車站代碼
	// Required: true
	StationID *string `json:"StationID"`

	// NameType
	//
	// 查詢車站名稱
	// Required: true
	StationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StationName"`

	// Array
	//
	// 車次停靠站點組合
	// Required: true
	StopStations []*PTXServiceDTORailSpecificationV2THSRStopStation `json:"StopStations"`

	// String
	//
	// 車次號碼
	// Required: true
	TrainNo *string `json:"TrainNo"`
}

// Validate validates this p t x service d t o rail specification v2 t h s r old available seat
func (m *PTXServiceDTORailSpecificationV2THSROldAvailableSeat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDepartureTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndingStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndingStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopStations(formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV2THSROldAvailableSeat) validateDepartureTime(formats strfmt.Registry) error {

	if err := validate.Required("DepartureTime", "body", m.DepartureTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSROldAvailableSeat) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSROldAvailableSeat) validateEndingStationID(formats strfmt.Registry) error {

	if err := validate.Required("EndingStationID", "body", m.EndingStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSROldAvailableSeat) validateEndingStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSROldAvailableSeat) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSROldAvailableSeat) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSROldAvailableSeat) validateStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSROldAvailableSeat) validateStopStations(formats strfmt.Registry) error {

	if err := validate.Required("StopStations", "body", m.StopStations); err != nil {
		return err
	}

	for i := 0; i < len(m.StopStations); i++ {
		if swag.IsZero(m.StopStations[i]) { // not required
			continue
		}

		if m.StopStations[i] != nil {
			if err := m.StopStations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StopStations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSROldAvailableSeat) validateTrainNo(formats strfmt.Registry) error {

	if err := validate.Required("TrainNo", "body", m.TrainNo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 t h s r old available seat based on the context it is used
func (m *PTXServiceDTORailSpecificationV2THSROldAvailableSeat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndingStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStopStations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSROldAvailableSeat) contextValidateEndingStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSROldAvailableSeat) contextValidateStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSROldAvailableSeat) contextValidateStopStations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StopStations); i++ {

		if m.StopStations[i] != nil {
			if err := m.StopStations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StopStations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2THSROldAvailableSeat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2THSROldAvailableSeat) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2THSROldAvailableSeat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
