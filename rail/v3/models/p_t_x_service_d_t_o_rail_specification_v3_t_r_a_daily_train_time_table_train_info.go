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

// PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo TrainInfo
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.DailyTrainTimeTable.TrainInfo
type PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo struct {

	// integer
	//
	// 是否人車同行班次(置於攜車袋之自行車各級列車均可乘車) : [0:'否',1:'是']
	// Required: true
	BikeFlag *string `json:"BikeFlag"`

	// integer
	//
	// 是否設有哺(集)乳室車廂 : [0:'否',1:'是']
	// Required: true
	BreastFeedFlag *string `json:"BreastFeedFlag"`

	// integer
	//
	// 是否提供小汽車上火車服務 : [0:'否',1:'是']
	CarFlag string `json:"CarFlag,omitempty"`

	// integer
	//
	// 是否每日行駛 : [0:'否',1:'是']
	// Required: true
	DailyFlag *string `json:"DailyFlag"`

	// integer
	//
	// 是否提供訂便當服務 : [0:'否',1:'是']
	// Required: true
	DiningFlag *string `json:"DiningFlag"`

	// DiningFlagSection[]
	//
	// 提供訂便當服務之車站區間
	DiningFlagSections []*PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableDiningFlagSection `json:"DiningFlagSections" xml:"DiningFlagSection[]"`

	// integer
	//
	// 行駛方向 : [0:'順行',1:'逆行']
	// Required: true
	Direction *string `json:"Direction"`

	// String
	//
	// 車次之終點站車站代號
	EndingStationID string `json:"EndingStationID,omitempty" xml:"String,omitempty"`

	// NameType
	//
	// 車次之終點站車站名稱
	// Required: true
	EndingStationName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"EndingStationName" xml:"NameType"`

	// integer
	//
	// 是否為加班車 : [0:'否',1:'是']
	// Required: true
	ExtraTrainFlag *string `json:"ExtraTrainFlag"`

	// String
	//
	// 附註說明
	Note string `json:"Note,omitempty" xml:"String,omitempty"`

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
	// 營運路線代碼
	// Required: true
	RouteID *string `json:"RouteID" xml:"String"`

	// String
	//
	// 車次之起始站車站代號
	StartingStationID string `json:"StartingStationID,omitempty" xml:"String,omitempty"`

	// NameType
	//
	// 車次之起始站車站名稱
	// Required: true
	StartingStationName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"StartingStationName" xml:"NameType"`

	// String
	//
	// 車次代碼
	// Required: true
	TrainNo *string `json:"TrainNo" xml:"String"`

	// String
	//
	// 車種簡碼 = ['1: 太魯閣', '2: 普悠瑪', '3: 自強', '4: 莒光', '5: 復興', '6: 區間', '7: 普快', '10: 區間快']
	TrainTypeCode string `json:"TrainTypeCode,omitempty" xml:"String,omitempty"`

	// String
	//
	// 車種代嗎
	TrainTypeID string `json:"TrainTypeID,omitempty" xml:"String,omitempty"`

	// NameType
	//
	// 車種名稱
	// Required: true
	TrainTypeName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"TrainTypeName" xml:"NameType"`

	// String
	//
	// 車次之目的地方向描述
	TripHeadSign string `json:"TripHeadSign,omitempty" xml:"String,omitempty"`

	// integer
	//
	// 山海線類型 : [0:'不經山海線',1:'山線',2:'海線',3:'成追線']
	TripLine string `json:"TripLine,omitempty"`

	// integer
	//
	// 是否設身障旅客專用座位車 : [0:'否',1:'是']
	// Required: true
	WheelChairFlag *string `json:"WheelChairFlag"`
}

// Validate validates this p t x service d t o rail specification v3 t r a daily train time table train info
func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBikeFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBreastFeedFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDailyFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiningFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiningFlagSections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndingStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtraTrainFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageServiceFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
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

	if err := m.validateWheelChairFlag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) validateBikeFlag(formats strfmt.Registry) error {

	if err := validate.Required("BikeFlag", "body", m.BikeFlag); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) validateBreastFeedFlag(formats strfmt.Registry) error {

	if err := validate.Required("BreastFeedFlag", "body", m.BreastFeedFlag); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) validateDailyFlag(formats strfmt.Registry) error {

	if err := validate.Required("DailyFlag", "body", m.DailyFlag); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) validateDiningFlag(formats strfmt.Registry) error {

	if err := validate.Required("DiningFlag", "body", m.DiningFlag); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) validateDiningFlagSections(formats strfmt.Registry) error {
	if swag.IsZero(m.DiningFlagSections) { // not required
		return nil
	}

	for i := 0; i < len(m.DiningFlagSections); i++ {
		if swag.IsZero(m.DiningFlagSections[i]) { // not required
			continue
		}

		if m.DiningFlagSections[i] != nil {
			if err := m.DiningFlagSections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DiningFlagSections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) validateEndingStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) validateExtraTrainFlag(formats strfmt.Registry) error {

	if err := validate.Required("ExtraTrainFlag", "body", m.ExtraTrainFlag); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) validatePackageServiceFlag(formats strfmt.Registry) error {

	if err := validate.Required("PackageServiceFlag", "body", m.PackageServiceFlag); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) validateStartingStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) validateTrainNo(formats strfmt.Registry) error {

	if err := validate.Required("TrainNo", "body", m.TrainNo); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) validateTrainTypeName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) validateWheelChairFlag(formats strfmt.Registry) error {

	if err := validate.Required("WheelChairFlag", "body", m.WheelChairFlag); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a daily train time table train info based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDiningFlagSections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndingStationName(ctx, formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) contextValidateDiningFlagSections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DiningFlagSections); i++ {

		if m.DiningFlagSections[i] != nil {
			if err := m.DiningFlagSections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DiningFlagSections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) contextValidateEndingStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) contextValidateStartingStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) contextValidateTrainTypeName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
