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

// PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo RailGeneralTrainInfo
//
// 臺鐵車次定期資料型別
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.TRA.TimeInfo.RailGeneralTrainInfo
type PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo struct {

	// integer
	//
	// 是否人車同行班次(置於攜車袋之自行車各級列車均可乘車) : [0:'否',1:'是']
	// Required: true
	BikeFlag *string `json:"BikeFlag"`

	// integer
	//
	// 是否設有哺(集)乳室車廂 : [0:'否',1:'是']
	// Required: true
	BreastFeedingFlag *string `json:"BreastFeedingFlag"`

	// integer
	//
	// 是否為每日行駛 : [0:'否',1:'是']
	// Required: true
	DailyFlag *string `json:"DailyFlag"`

	// integer
	//
	// 是否提供訂便當服務 : [0:'否',1:'是']
	// Required: true
	DiningFlag *string `json:"DiningFlag"`

	// integer
	//
	// 順逆行 : [0:'順行',1:'逆行']
	// Required: true
	Direction *string `json:"Direction"`

	// String
	//
	// 列車終點車站代號
	EndingStationID string `json:"EndingStationID,omitempty" xml:"String,omitempty"`

	// NameType
	//
	// 列車終點車站名稱
	EndingStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"EndingStationName,omitempty" xml:"NameType,omitempty"`

	// NameType
	//
	// 附註說明
	Note struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"Note,omitempty" xml:"NameType,omitempty"`

	// String
	//
	// 跨夜車站代碼
	OverNightStationID string `json:"OverNightStationID,omitempty" xml:"String,omitempty"`

	// integer
	//
	// 是否提供行李服務 : [0:'否',1:'是']
	// Required: true
	PackageServiceFlag *string `json:"PackageServiceFlag"`

	// String
	//
	// 列車起點車站代號
	StartingStationID string `json:"StartingStationID,omitempty" xml:"String,omitempty"`

	// NameType
	//
	// 列車起點車站名稱
	StartingStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StartingStationName,omitempty" xml:"NameType,omitempty"`

	// String
	//
	// 車次代碼
	// Required: true
	TrainNo *string `json:"TrainNo" xml:"String"`

	// String
	//
	// 列車車種簡碼
	TrainTypeCode string `json:"TrainTypeCode,omitempty" xml:"String,omitempty"`

	// String
	//
	// 列車車種代碼
	TrainTypeID string `json:"TrainTypeID,omitempty" xml:"String,omitempty"`

	// NameType
	//
	// 列車車種名稱
	TrainTypeName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"TrainTypeName,omitempty" xml:"NameType,omitempty"`

	// integer
	//
	// 山海線類型 : [0:'不經山海線',1:'山線',2:'海線',3:'成追線']
	TripLine string `json:"TripLine,omitempty"`

	// integer
	//
	// 是否設身障旅客專用座位車 : [0:'否',1:'是']
	// Required: true
	WheelchairFlag *string `json:"WheelchairFlag"`
}

// Validate validates this p t x service d t o rail specification v2 t r a time info rail general train info
func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBikeFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBreastFeedingFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDailyFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiningFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndingStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageServiceFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartingStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainNo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainTypeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWheelchairFlag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) validateBikeFlag(formats strfmt.Registry) error {

	if err := validate.Required("BikeFlag", "body", m.BikeFlag); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) validateBreastFeedingFlag(formats strfmt.Registry) error {

	if err := validate.Required("BreastFeedingFlag", "body", m.BreastFeedingFlag); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) validateDailyFlag(formats strfmt.Registry) error {

	if err := validate.Required("DailyFlag", "body", m.DailyFlag); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) validateDiningFlag(formats strfmt.Registry) error {

	if err := validate.Required("DiningFlag", "body", m.DiningFlag); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) validateEndingStationName(formats strfmt.Registry) error {
	if swag.IsZero(m.EndingStationName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) validateNote(formats strfmt.Registry) error {
	if swag.IsZero(m.Note) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) validatePackageServiceFlag(formats strfmt.Registry) error {

	if err := validate.Required("PackageServiceFlag", "body", m.PackageServiceFlag); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) validateStartingStationName(formats strfmt.Registry) error {
	if swag.IsZero(m.StartingStationName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) validateTrainNo(formats strfmt.Registry) error {

	if err := validate.Required("TrainNo", "body", m.TrainNo); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) validateTrainTypeName(formats strfmt.Registry) error {
	if swag.IsZero(m.TrainTypeName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) validateWheelchairFlag(formats strfmt.Registry) error {

	if err := validate.Required("WheelchairFlag", "body", m.WheelchairFlag); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 t r a time info rail general train info based on the context it is used
func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndingStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartingStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrainTypeName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) contextValidateEndingStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) contextValidateNote(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) contextValidateStartingStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) contextValidateTrainTypeName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2TRATimeInfoRailGeneralTrainInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
