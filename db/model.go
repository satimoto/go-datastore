// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/twpayne/go-geom"
)

type AuthMethodType string

const (
	AuthMethodTypeAUTHREQUEST AuthMethodType = "AUTH_REQUEST"
	AuthMethodTypeWHITELIST   AuthMethodType = "WHITELIST"
)

func (e *AuthMethodType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AuthMethodType(s)
	case string:
		*e = AuthMethodType(s)
	default:
		return fmt.Errorf("unsupported scan type for AuthMethodType: %T", src)
	}
	return nil
}

type AuthenticationActions string

const (
	AuthenticationActionsRegister AuthenticationActions = "register"
	AuthenticationActionsLogin    AuthenticationActions = "login"
	AuthenticationActionsLink     AuthenticationActions = "link"
	AuthenticationActionsAuth     AuthenticationActions = "auth"
)

func (e *AuthenticationActions) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AuthenticationActions(s)
	case string:
		*e = AuthenticationActions(s)
	default:
		return fmt.Errorf("unsupported scan type for AuthenticationActions: %T", src)
	}
	return nil
}

type ChannelRequestStatus string

const (
	ChannelRequestStatusREQUESTED        ChannelRequestStatus = "REQUESTED"
	ChannelRequestStatusAWAITINGPAYMENTS ChannelRequestStatus = "AWAITING_PAYMENTS"
	ChannelRequestStatusAWAITINGPREIMAGE ChannelRequestStatus = "AWAITING_PREIMAGE"
	ChannelRequestStatusSETTLINGHTLCS    ChannelRequestStatus = "SETTLING_HTLCS"
	ChannelRequestStatusOPENINGCHANNEL   ChannelRequestStatus = "OPENING_CHANNEL"
	ChannelRequestStatusCOMPLETED        ChannelRequestStatus = "COMPLETED"
	ChannelRequestStatusFAILED           ChannelRequestStatus = "FAILED"
)

func (e *ChannelRequestStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ChannelRequestStatus(s)
	case string:
		*e = ChannelRequestStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for ChannelRequestStatus: %T", src)
	}
	return nil
}

type ChargingPeriodDimensionType string

const (
	ChargingPeriodDimensionTypeENERGY      ChargingPeriodDimensionType = "ENERGY"
	ChargingPeriodDimensionTypeFLAT        ChargingPeriodDimensionType = "FLAT"
	ChargingPeriodDimensionTypeMAXCURRENT  ChargingPeriodDimensionType = "MAX_CURRENT"
	ChargingPeriodDimensionTypeMINCURRENT  ChargingPeriodDimensionType = "MIN_CURRENT"
	ChargingPeriodDimensionTypePARKINGTIME ChargingPeriodDimensionType = "PARKING_TIME"
	ChargingPeriodDimensionTypeTIME        ChargingPeriodDimensionType = "TIME"
)

func (e *ChargingPeriodDimensionType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ChargingPeriodDimensionType(s)
	case string:
		*e = ChargingPeriodDimensionType(s)
	default:
		return fmt.Errorf("unsupported scan type for ChargingPeriodDimensionType: %T", src)
	}
	return nil
}

type CommandResponseType string

const (
	CommandResponseTypeREQUESTED      CommandResponseType = "REQUESTED"
	CommandResponseTypeNOTSUPPORTED   CommandResponseType = "NOT_SUPPORTED"
	CommandResponseTypeREJECTED       CommandResponseType = "REJECTED"
	CommandResponseTypeACCEPTED       CommandResponseType = "ACCEPTED"
	CommandResponseTypeTIMEOUT        CommandResponseType = "TIMEOUT"
	CommandResponseTypeUNKNOWNSESSION CommandResponseType = "UNKNOWN_SESSION"
)

func (e *CommandResponseType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CommandResponseType(s)
	case string:
		*e = CommandResponseType(s)
	default:
		return fmt.Errorf("unsupported scan type for CommandResponseType: %T", src)
	}
	return nil
}

type ConnectorFormat string

const (
	ConnectorFormatSOCKET ConnectorFormat = "SOCKET"
	ConnectorFormatCABLE  ConnectorFormat = "CABLE"
)

func (e *ConnectorFormat) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ConnectorFormat(s)
	case string:
		*e = ConnectorFormat(s)
	default:
		return fmt.Errorf("unsupported scan type for ConnectorFormat: %T", src)
	}
	return nil
}

type ConnectorType string

const (
	ConnectorTypeCHADEMO           ConnectorType = "CHADEMO"
	ConnectorTypeDOMESTICA         ConnectorType = "DOMESTIC_A"
	ConnectorTypeDOMESTICB         ConnectorType = "DOMESTIC_B"
	ConnectorTypeDOMESTICC         ConnectorType = "DOMESTIC_C"
	ConnectorTypeDOMESTICD         ConnectorType = "DOMESTIC_D"
	ConnectorTypeDOMESTICE         ConnectorType = "DOMESTIC_E"
	ConnectorTypeDOMESTICF         ConnectorType = "DOMESTIC_F"
	ConnectorTypeDOMESTICG         ConnectorType = "DOMESTIC_G"
	ConnectorTypeDOMESTICH         ConnectorType = "DOMESTIC_H"
	ConnectorTypeDOMESTICI         ConnectorType = "DOMESTIC_I"
	ConnectorTypeDOMESTICJ         ConnectorType = "DOMESTIC_J"
	ConnectorTypeDOMESTICK         ConnectorType = "DOMESTIC_K"
	ConnectorTypeDOMESTICL         ConnectorType = "DOMESTIC_L"
	ConnectorTypeIEC603092Single16 ConnectorType = "IEC_60309_2_single_16"
	ConnectorTypeIEC603092Three16  ConnectorType = "IEC_60309_2_three_16"
	ConnectorTypeIEC603092Three32  ConnectorType = "IEC_60309_2_three_32"
	ConnectorTypeIEC603092Three64  ConnectorType = "IEC_60309_2_three_64"
	ConnectorTypeIEC62196T1        ConnectorType = "IEC_62196_T1"
	ConnectorTypeIEC62196T1COMBO   ConnectorType = "IEC_62196_T1_COMBO"
	ConnectorTypeIEC62196T2        ConnectorType = "IEC_62196_T2"
	ConnectorTypeIEC62196T2COMBO   ConnectorType = "IEC_62196_T2_COMBO"
	ConnectorTypeIEC62196T3A       ConnectorType = "IEC_62196_T3A"
	ConnectorTypeIEC62196T3C       ConnectorType = "IEC_62196_T3C"
	ConnectorTypeTESLAR            ConnectorType = "TESLA_R"
	ConnectorTypeTESLAS            ConnectorType = "TESLA_S"
)

func (e *ConnectorType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ConnectorType(s)
	case string:
		*e = ConnectorType(s)
	default:
		return fmt.Errorf("unsupported scan type for ConnectorType: %T", src)
	}
	return nil
}

type EnergySourceCategory string

const (
	EnergySourceCategoryNUCLEAR       EnergySourceCategory = "NUCLEAR"
	EnergySourceCategoryGENERALFOSSIL EnergySourceCategory = "GENERAL_FOSSIL"
	EnergySourceCategoryCOAL          EnergySourceCategory = "COAL"
	EnergySourceCategoryGAS           EnergySourceCategory = "GAS"
	EnergySourceCategoryGENERALGREEN  EnergySourceCategory = "GENERAL_GREEN"
	EnergySourceCategorySOLAR         EnergySourceCategory = "SOLAR"
	EnergySourceCategoryWIND          EnergySourceCategory = "WIND"
	EnergySourceCategoryWATER         EnergySourceCategory = "WATER"
)

func (e *EnergySourceCategory) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EnergySourceCategory(s)
	case string:
		*e = EnergySourceCategory(s)
	default:
		return fmt.Errorf("unsupported scan type for EnergySourceCategory: %T", src)
	}
	return nil
}

type EnvironmentalImpactCategory string

const (
	EnvironmentalImpactCategoryNUCLEARWASTE  EnvironmentalImpactCategory = "NUCLEAR_WASTE"
	EnvironmentalImpactCategoryCARBONDIOXIDE EnvironmentalImpactCategory = "CARBON_DIOXIDE"
)

func (e *EnvironmentalImpactCategory) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EnvironmentalImpactCategory(s)
	case string:
		*e = EnvironmentalImpactCategory(s)
	default:
		return fmt.Errorf("unsupported scan type for EnvironmentalImpactCategory: %T", src)
	}
	return nil
}

type EvseStatus string

const (
	EvseStatusAVAILABLE   EvseStatus = "AVAILABLE"
	EvseStatusBLOCKED     EvseStatus = "BLOCKED"
	EvseStatusCHARGING    EvseStatus = "CHARGING"
	EvseStatusINOPERATIVE EvseStatus = "INOPERATIVE"
	EvseStatusOUTOFORDER  EvseStatus = "OUTOFORDER"
	EvseStatusPLANNED     EvseStatus = "PLANNED"
	EvseStatusREMOVED     EvseStatus = "REMOVED"
	EvseStatusRESERVED    EvseStatus = "RESERVED"
	EvseStatusUNKNOWN     EvseStatus = "UNKNOWN"
)

func (e *EvseStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EvseStatus(s)
	case string:
		*e = EvseStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for EvseStatus: %T", src)
	}
	return nil
}

type ImageCategory string

const (
	ImageCategoryCHARGER  ImageCategory = "CHARGER"
	ImageCategoryENTRANCE ImageCategory = "ENTRANCE"
	ImageCategoryLOCATION ImageCategory = "LOCATION"
	ImageCategoryNETWORK  ImageCategory = "NETWORK"
	ImageCategoryOPERATOR ImageCategory = "OPERATOR"
	ImageCategoryOTHER    ImageCategory = "OTHER"
	ImageCategoryOWNER    ImageCategory = "OWNER"
)

func (e *ImageCategory) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ImageCategory(s)
	case string:
		*e = ImageCategory(s)
	default:
		return fmt.Errorf("unsupported scan type for ImageCategory: %T", src)
	}
	return nil
}

type LocationType string

const (
	LocationTypeONSTREET          LocationType = "ON_STREET"
	LocationTypePARKINGGARAGE     LocationType = "PARKING_GARAGE"
	LocationTypeUNDERGROUNDGARAGE LocationType = "UNDERGROUND_GARAGE"
	LocationTypePARKINGLOT        LocationType = "PARKING_LOT"
	LocationTypeOTHER             LocationType = "OTHER"
	LocationTypeUNKNOWN           LocationType = "UNKNOWN"
)

func (e *LocationType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = LocationType(s)
	case string:
		*e = LocationType(s)
	default:
		return fmt.Errorf("unsupported scan type for LocationType: %T", src)
	}
	return nil
}

type PeriodType string

const (
	PeriodTypeOPENING PeriodType = "OPENING"
	PeriodTypeCLOSING PeriodType = "CLOSING"
)

func (e *PeriodType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PeriodType(s)
	case string:
		*e = PeriodType(s)
	default:
		return fmt.Errorf("unsupported scan type for PeriodType: %T", src)
	}
	return nil
}

type PowerType string

const (
	PowerTypeAC1PHASE PowerType = "AC_1_PHASE"
	PowerTypeAC3PHASE PowerType = "AC_3_PHASE"
	PowerTypeDC       PowerType = "DC"
)

func (e *PowerType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PowerType(s)
	case string:
		*e = PowerType(s)
	default:
		return fmt.Errorf("unsupported scan type for PowerType: %T", src)
	}
	return nil
}

type SessionStatusType string

const (
	SessionStatusTypeACTIVE    SessionStatusType = "ACTIVE"
	SessionStatusTypeCOMPLETED SessionStatusType = "COMPLETED"
	SessionStatusTypeINVALID   SessionStatusType = "INVALID"
	SessionStatusTypePENDING   SessionStatusType = "PENDING"
)

func (e *SessionStatusType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = SessionStatusType(s)
	case string:
		*e = SessionStatusType(s)
	default:
		return fmt.Errorf("unsupported scan type for SessionStatusType: %T", src)
	}
	return nil
}

type TariffDimension string

const (
	TariffDimensionENERGY      TariffDimension = "ENERGY"
	TariffDimensionFLAT        TariffDimension = "FLAT"
	TariffDimensionPARKINGTIME TariffDimension = "PARKING_TIME"
	TariffDimensionTIME        TariffDimension = "TIME"
)

func (e *TariffDimension) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TariffDimension(s)
	case string:
		*e = TariffDimension(s)
	default:
		return fmt.Errorf("unsupported scan type for TariffDimension: %T", src)
	}
	return nil
}

type TokenAllowedType string

const (
	TokenAllowedTypeALLOWED    TokenAllowedType = "ALLOWED"
	TokenAllowedTypeBLOCKED    TokenAllowedType = "BLOCKED"
	TokenAllowedTypeEXPIRED    TokenAllowedType = "EXPIRED"
	TokenAllowedTypeNOCREDIT   TokenAllowedType = "NO_CREDIT"
	TokenAllowedTypeNOTALLOWED TokenAllowedType = "NOT_ALLOWED"
)

func (e *TokenAllowedType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TokenAllowedType(s)
	case string:
		*e = TokenAllowedType(s)
	default:
		return fmt.Errorf("unsupported scan type for TokenAllowedType: %T", src)
	}
	return nil
}

type TokenType string

const (
	TokenTypeOTHER TokenType = "OTHER"
	TokenTypeRFID  TokenType = "RFID"
)

func (e *TokenType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TokenType(s)
	case string:
		*e = TokenType(s)
	default:
		return fmt.Errorf("unsupported scan type for TokenType: %T", src)
	}
	return nil
}

type TokenWhitelistType string

const (
	TokenWhitelistTypeALWAYS         TokenWhitelistType = "ALWAYS"
	TokenWhitelistTypeALLOWED        TokenWhitelistType = "ALLOWED"
	TokenWhitelistTypeALLOWEDOFFLINE TokenWhitelistType = "ALLOWED_OFFLINE"
	TokenWhitelistTypeNEVER          TokenWhitelistType = "NEVER"
)

func (e *TokenWhitelistType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TokenWhitelistType(s)
	case string:
		*e = TokenWhitelistType(s)
	default:
		return fmt.Errorf("unsupported scan type for TokenWhitelistType: %T", src)
	}
	return nil
}

type Authentication struct {
	ID            int64                 `db:"id" json:"id"`
	Code          string                `db:"code" json:"code"`
	Action        AuthenticationActions `db:"action" json:"action"`
	Challenge     string                `db:"challenge" json:"challenge"`
	Signature     sql.NullString        `db:"signature" json:"signature"`
	LinkingPubkey sql.NullString        `db:"linking_pubkey" json:"linkingPubkey"`
}

type BusinessDetail struct {
	ID      int64          `db:"id" json:"id"`
	Name    string         `db:"name" json:"name"`
	Website sql.NullString `db:"website" json:"website"`
	LogoID  sql.NullInt64  `db:"logo_id" json:"logoID"`
}

type Calibration struct {
	ID                    int64          `db:"id" json:"id"`
	EncodingMethod        string         `db:"encoding_method" json:"encodingMethod"`
	EncodingMethodVersion sql.NullInt32  `db:"encoding_method_version" json:"encodingMethodVersion"`
	PublicKey             sql.NullString `db:"public_key" json:"publicKey"`
	Url                   sql.NullString `db:"url" json:"url"`
}

type CalibrationValue struct {
	ID            int64  `db:"id" json:"id"`
	CalibrationID int64  `db:"calibration_id" json:"calibrationID"`
	Nature        string `db:"nature" json:"nature"`
	PlainData     string `db:"plain_data" json:"plainData"`
	SignedData    string `db:"signed_data" json:"signedData"`
}

type Capability struct {
	ID          int64  `db:"id" json:"id"`
	Text        string `db:"text" json:"text"`
	Description string `db:"description" json:"description"`
}

type Cdr struct {
	ID               int64           `db:"id" json:"id"`
	Uid              string          `db:"uid" json:"uid"`
	AuthorizationID  sql.NullString  `db:"authorization_id" json:"authorizationID"`
	StartDateTime    time.Time       `db:"start_date_time" json:"startDateTime"`
	StopDateTime     sql.NullTime    `db:"stop_date_time" json:"stopDateTime"`
	AuthID           string          `db:"auth_id" json:"authID"`
	AuthMethod       AuthMethodType  `db:"auth_method" json:"authMethod"`
	LocationID       int64           `db:"location_id" json:"locationID"`
	MeterID          sql.NullString  `db:"meter_id" json:"meterID"`
	Currency         string          `db:"currency" json:"currency"`
	CalibrationID    sql.NullInt64   `db:"calibration_id" json:"calibrationID"`
	TotalCost        float64         `db:"total_cost" json:"totalCost"`
	TotalEnergy      float64         `db:"total_energy" json:"totalEnergy"`
	TotalTime        float64         `db:"total_time" json:"totalTime"`
	TotalParkingTime sql.NullFloat64 `db:"total_parking_time" json:"totalParkingTime"`
	Remark           sql.NullString  `db:"remark" json:"remark"`
	LastUpdated      time.Time       `db:"last_updated" json:"lastUpdated"`
}

type CdrChargingPeriod struct {
	CdrID            int64 `db:"cdr_id" json:"cdrID"`
	ChargingPeriodID int64 `db:"charging_period_id" json:"chargingPeriodID"`
}

type ChannelRequest struct {
	ID          int64                `db:"id" json:"id"`
	UserID      int64                `db:"user_id" json:"userID"`
	Status      ChannelRequestStatus `db:"status" json:"status"`
	Pubkey      string               `db:"pubkey" json:"pubkey"`
	PaymentHash []byte               `db:"payment_hash" json:"paymentHash"`
	PaymentAddr []byte               `db:"payment_addr" json:"paymentAddr"`
	AmountMsat  int64                `db:"amount_msat" json:"amountMsat"`
	SettledMsat int64                `db:"settled_msat" json:"settledMsat"`
	FundingTxID []byte               `db:"funding_tx_id" json:"fundingTxID"`
	OutputIndex sql.NullInt64        `db:"output_index" json:"outputIndex"`
	NodeID      int64                `db:"node_id" json:"nodeID"`
}

type ChannelRequestHtlc struct {
	ID               int64 `db:"id" json:"id"`
	ChannelRequestID int64 `db:"channel_request_id" json:"channelRequestID"`
	ChanID           int64 `db:"chan_id" json:"chanID"`
	HtlcID           int64 `db:"htlc_id" json:"htlcID"`
	IsSettled        bool  `db:"is_settled" json:"isSettled"`
}

type ChargingPeriod struct {
	ID            int64     `db:"id" json:"id"`
	StartDateTime time.Time `db:"start_date_time" json:"startDateTime"`
}

type ChargingPeriodDimension struct {
	ID               int64                       `db:"id" json:"id"`
	ChargingPeriodID int64                       `db:"charging_period_id" json:"chargingPeriodID"`
	Type             ChargingPeriodDimensionType `db:"type" json:"type"`
	Volume           float64                     `db:"volume" json:"volume"`
}

type CommandReservation struct {
	ID            int64               `db:"id" json:"id"`
	Status        CommandResponseType `db:"status" json:"status"`
	TokenID       int64               `db:"token_id" json:"tokenID"`
	ExpiryDate    time.Time           `db:"expiry_date" json:"expiryDate"`
	ReservationID int64               `db:"reservation_id" json:"reservationID"`
	LocationID    string              `db:"location_id" json:"locationID"`
	EvseUid       sql.NullString      `db:"evse_uid" json:"evseUid"`
}

type CommandStart struct {
	ID              int64               `db:"id" json:"id"`
	Status          CommandResponseType `db:"status" json:"status"`
	TokenID         int64               `db:"token_id" json:"tokenID"`
	AuthorizationID sql.NullString      `db:"authorization_id" json:"authorizationID"`
	LocationID      string              `db:"location_id" json:"locationID"`
	EvseUid         sql.NullString      `db:"evse_uid" json:"evseUid"`
}

type CommandStop struct {
	ID        int64               `db:"id" json:"id"`
	Status    CommandResponseType `db:"status" json:"status"`
	SessionID string              `db:"session_id" json:"sessionID"`
}

type CommandUnlock struct {
	ID          int64               `db:"id" json:"id"`
	Status      CommandResponseType `db:"status" json:"status"`
	LocationID  string              `db:"location_id" json:"locationID"`
	EvseUid     string              `db:"evse_uid" json:"evseUid"`
	ConnectorID string              `db:"connector_id" json:"connectorID"`
}

type Connector struct {
	ID                 int64           `db:"id" json:"id"`
	EvseID             int64           `db:"evse_id" json:"evseID"`
	Uid                string          `db:"uid" json:"uid"`
	Standard           ConnectorType   `db:"standard" json:"standard"`
	Format             ConnectorFormat `db:"format" json:"format"`
	PowerType          PowerType       `db:"power_type" json:"powerType"`
	Voltage            int32           `db:"voltage" json:"voltage"`
	Amperage           int32           `db:"amperage" json:"amperage"`
	TariffID           sql.NullString  `db:"tariff_id" json:"tariffID"`
	TermsAndConditions sql.NullString  `db:"terms_and_conditions" json:"termsAndConditions"`
	LastUpdated        time.Time       `db:"last_updated" json:"lastUpdated"`
}

type Credential struct {
	ID               int64          `db:"id" json:"id"`
	ClientToken      sql.NullString `db:"client_token" json:"clientToken"`
	ServerToken      sql.NullString `db:"server_token" json:"serverToken"`
	Version          sql.NullString `db:"version" json:"version"`
	Url              string         `db:"url" json:"url"`
	BusinessDetailID int64          `db:"business_detail_id" json:"businessDetailID"`
	PartyID          string         `db:"party_id" json:"partyID"`
	CountryCode      string         `db:"country_code" json:"countryCode"`
	LastUpdated      time.Time      `db:"last_updated" json:"lastUpdated"`
}

type DisplayText struct {
	ID       int64  `db:"id" json:"id"`
	Language string `db:"language" json:"language"`
	Text     string `db:"text" json:"text"`
}

type Element struct {
	ID            int64         `db:"id" json:"id"`
	TariffID      int64         `db:"tariff_id" json:"tariffID"`
	RestrictionID sql.NullInt64 `db:"restriction_id" json:"restrictionID"`
}

type EmailSubscription struct {
	ID               int64     `db:"id" json:"id"`
	Email            string    `db:"email" json:"email"`
	VerificationCode string    `db:"verification_code" json:"verificationCode"`
	UnsubscribeCode  string    `db:"unsubscribe_code" json:"unsubscribeCode"`
	Locale           string    `db:"locale" json:"locale"`
	IsVerified       bool      `db:"is_verified" json:"isVerified"`
	CreatedDate      time.Time `db:"created_date" json:"createdDate"`
}

type EnergyMix struct {
	ID                int64          `db:"id" json:"id"`
	IsGreenEnergy     bool           `db:"is_green_energy" json:"isGreenEnergy"`
	SupplierName      sql.NullString `db:"supplier_name" json:"supplierName"`
	EnergyProductName sql.NullString `db:"energy_product_name" json:"energyProductName"`
}

type EnergySource struct {
	ID          int64                `db:"id" json:"id"`
	EnergyMixID int64                `db:"energy_mix_id" json:"energyMixID"`
	Source      EnergySourceCategory `db:"source" json:"source"`
	Percentage  float64              `db:"percentage" json:"percentage"`
}

type EnvironmentalImpact struct {
	ID          int64                       `db:"id" json:"id"`
	EnergyMixID int64                       `db:"energy_mix_id" json:"energyMixID"`
	Source      EnvironmentalImpactCategory `db:"source" json:"source"`
	Amount      float64                     `db:"amount" json:"amount"`
}

type Evse struct {
	ID                int64          `db:"id" json:"id"`
	LocationID        int64          `db:"location_id" json:"locationID"`
	Uid               string         `db:"uid" json:"uid"`
	EvseID            sql.NullString `db:"evse_id" json:"evseID"`
	Status            EvseStatus     `db:"status" json:"status"`
	FloorLevel        sql.NullString `db:"floor_level" json:"floorLevel"`
	Geom              interface{}    `db:"geom" json:"geom"`
	GeoLocationID     sql.NullInt64  `db:"geo_location_id" json:"geoLocationID"`
	PhysicalReference sql.NullString `db:"physical_reference" json:"physicalReference"`
	LastUpdated       time.Time      `db:"last_updated" json:"lastUpdated"`
}

type EvseCapability struct {
	EvseID       int64 `db:"evse_id" json:"evseID"`
	CapabilityID int64 `db:"capability_id" json:"capabilityID"`
}

type EvseDirection struct {
	EvseID        int64 `db:"evse_id" json:"evseID"`
	DisplayTextID int64 `db:"display_text_id" json:"displayTextID"`
}

type EvseImage struct {
	EvseID  int64 `db:"evse_id" json:"evseID"`
	ImageID int64 `db:"image_id" json:"imageID"`
}

type EvseParkingRestriction struct {
	EvseID               int64 `db:"evse_id" json:"evseID"`
	ParkingRestrictionID int64 `db:"parking_restriction_id" json:"parkingRestrictionID"`
}

type ExceptionalPeriod struct {
	ID            int64      `db:"id" json:"id"`
	OpeningTimeID int64      `db:"opening_time_id" json:"openingTimeID"`
	PeriodType    PeriodType `db:"period_type" json:"periodType"`
	PeriodBegin   time.Time  `db:"period_begin" json:"periodBegin"`
	PeriodEnd     time.Time  `db:"period_end" json:"periodEnd"`
}

type Facility struct {
	ID          int64  `db:"id" json:"id"`
	Text        string `db:"text" json:"text"`
	Description string `db:"description" json:"description"`
}

type GeoLocation struct {
	ID        int64          `db:"id" json:"id"`
	Latitude  string         `db:"latitude" json:"latitude"`
	Longitude string         `db:"longitude" json:"longitude"`
	Name      sql.NullString `db:"name" json:"name"`
}

type Image struct {
	ID        int64          `db:"id" json:"id"`
	Url       string         `db:"url" json:"url"`
	Thumbnail sql.NullString `db:"thumbnail" json:"thumbnail"`
	Category  ImageCategory  `db:"category" json:"category"`
	Type      string         `db:"type" json:"type"`
	Width     sql.NullInt32  `db:"width" json:"width"`
	Height    sql.NullInt32  `db:"height" json:"height"`
}

type Location struct {
	ID                 int64          `db:"id" json:"id"`
	Uid                string         `db:"uid" json:"uid"`
	Type               LocationType   `db:"type" json:"type"`
	Name               sql.NullString `db:"name" json:"name"`
	Address            string         `db:"address" json:"address"`
	City               string         `db:"city" json:"city"`
	PostalCode         string         `db:"postal_code" json:"postalCode"`
	Country            string         `db:"country" json:"country"`
	Geom               geom.Point     `db:"geom" json:"geom"`
	GeoLocationID      int64          `db:"geo_location_id" json:"geoLocationID"`
	OperatorID         sql.NullInt64  `db:"operator_id" json:"operatorID"`
	SuboperatorID      sql.NullInt64  `db:"suboperator_id" json:"suboperatorID"`
	OwnerID            sql.NullInt64  `db:"owner_id" json:"ownerID"`
	TimeZone           sql.NullString `db:"time_zone" json:"timeZone"`
	OpeningTimeID      sql.NullInt64  `db:"opening_time_id" json:"openingTimeID"`
	ChargingWhenClosed bool           `db:"charging_when_closed" json:"chargingWhenClosed"`
	EnergyMixID        sql.NullInt64  `db:"energy_mix_id" json:"energyMixID"`
	LastUpdated        time.Time      `db:"last_updated" json:"lastUpdated"`
}

type LocationDirection struct {
	LocationID    int64 `db:"location_id" json:"locationID"`
	DisplayTextID int64 `db:"display_text_id" json:"displayTextID"`
}

type LocationFacility struct {
	LocationID int64 `db:"location_id" json:"locationID"`
	FacilityID int64 `db:"facility_id" json:"facilityID"`
}

type LocationImage struct {
	LocationID int64 `db:"location_id" json:"locationID"`
	ImageID    int64 `db:"image_id" json:"imageID"`
}

type Node struct {
	ID         int64  `db:"id" json:"id"`
	Pubkey     string `db:"pubkey" json:"pubkey"`
	Addr       string `db:"addr" json:"addr"`
	Alias      string `db:"alias" json:"alias"`
	Color      string `db:"color" json:"color"`
	CommitHash string `db:"commit_hash" json:"commitHash"`
	Version    string `db:"version" json:"version"`
	Channels   int64  `db:"channels" json:"channels"`
	Peers      int64  `db:"peers" json:"peers"`
}

type OpeningTime struct {
	ID              int64 `db:"id" json:"id"`
	Twentyfourseven bool  `db:"twentyfourseven" json:"twentyfourseven"`
}

type ParkingRestriction struct {
	ID          int64  `db:"id" json:"id"`
	Text        string `db:"text" json:"text"`
	Description string `db:"description" json:"description"`
}

type PriceComponent struct {
	ID        int64           `db:"id" json:"id"`
	ElementID int64           `db:"element_id" json:"elementID"`
	Type      TariffDimension `db:"type" json:"type"`
	Price     float64         `db:"price" json:"price"`
	StepSize  int32           `db:"step_size" json:"stepSize"`
}

type RegularHour struct {
	ID            int64  `db:"id" json:"id"`
	OpeningTimeID int64  `db:"opening_time_id" json:"openingTimeID"`
	Weekday       int16  `db:"weekday" json:"weekday"`
	PeriodBegin   string `db:"period_begin" json:"periodBegin"`
	PeriodEnd     string `db:"period_end" json:"periodEnd"`
}

type RelatedLocation struct {
	LocationID    int64 `db:"location_id" json:"locationID"`
	GeoLocationID int64 `db:"geo_location_id" json:"geoLocationID"`
}

type Restriction struct {
	ID          int64           `db:"id" json:"id"`
	StartTime   sql.NullString  `db:"start_time" json:"startTime"`
	EndTime     sql.NullString  `db:"end_time" json:"endTime"`
	StartDate   sql.NullString  `db:"start_date" json:"startDate"`
	EndDate     sql.NullString  `db:"end_date" json:"endDate"`
	MinKwh      sql.NullFloat64 `db:"min_kwh" json:"minKwh"`
	MaxKwh      sql.NullFloat64 `db:"max_kwh" json:"maxKwh"`
	MinPower    sql.NullFloat64 `db:"min_power" json:"minPower"`
	MaxPower    sql.NullFloat64 `db:"max_power" json:"maxPower"`
	MinDuration sql.NullInt32   `db:"min_duration" json:"minDuration"`
	MaxDuration sql.NullInt32   `db:"max_duration" json:"maxDuration"`
}

type RestrictionWeekday struct {
	RestrictionID int64 `db:"restriction_id" json:"restrictionID"`
	WeekdayID     int64 `db:"weekday_id" json:"weekdayID"`
}

type Session struct {
	ID              int64             `db:"id" json:"id"`
	Uid             string            `db:"uid" json:"uid"`
	AuthorizationID sql.NullString    `db:"authorization_id" json:"authorizationID"`
	StartDatetime   time.Time         `db:"start_datetime" json:"startDatetime"`
	EndDatetime     sql.NullTime      `db:"end_datetime" json:"endDatetime"`
	Kwh             float64           `db:"kwh" json:"kwh"`
	AuthID          string            `db:"auth_id" json:"authID"`
	AuthMethod      AuthMethodType    `db:"auth_method" json:"authMethod"`
	LocationID      int64             `db:"location_id" json:"locationID"`
	MeterID         sql.NullString    `db:"meter_id" json:"meterID"`
	Currency        string            `db:"currency" json:"currency"`
	TotalCost       sql.NullFloat64   `db:"total_cost" json:"totalCost"`
	Status          SessionStatusType `db:"status" json:"status"`
	LastUpdated     time.Time         `db:"last_updated" json:"lastUpdated"`
}

type SessionChargingPeriod struct {
	SessionID        int64 `db:"session_id" json:"sessionID"`
	ChargingPeriodID int64 `db:"charging_period_id" json:"chargingPeriodID"`
}

type StatusSchedule struct {
	ID          int64        `db:"id" json:"id"`
	EvseID      int64        `db:"evse_id" json:"evseID"`
	PeriodBegin time.Time    `db:"period_begin" json:"periodBegin"`
	PeriodEnd   sql.NullTime `db:"period_end" json:"periodEnd"`
	Status      EvseStatus   `db:"status" json:"status"`
}

type Tariff struct {
	ID           int64          `db:"id" json:"id"`
	Uid          string         `db:"uid" json:"uid"`
	Currency     string         `db:"currency" json:"currency"`
	TariffAltUrl sql.NullString `db:"tariff_alt_url" json:"tariffAltUrl"`
	EnergyMixID  sql.NullInt64  `db:"energy_mix_id" json:"energyMixID"`
	LastUpdated  time.Time      `db:"last_updated" json:"lastUpdated"`
	CdrID        sql.NullInt64  `db:"cdr_id" json:"cdrID"`
}

type TariffAltText struct {
	TariffID      int64 `db:"tariff_id" json:"tariffID"`
	DisplayTextID int64 `db:"display_text_id" json:"displayTextID"`
}

type Token struct {
	ID           int64              `db:"id" json:"id"`
	Uid          string             `db:"uid" json:"uid"`
	Type         TokenType          `db:"type" json:"type"`
	AuthID       string             `db:"auth_id" json:"authID"`
	VisualNumber sql.NullString     `db:"visual_number" json:"visualNumber"`
	Issuer       string             `db:"issuer" json:"issuer"`
	Allowed      TokenAllowedType   `db:"allowed" json:"allowed"`
	Valid        bool               `db:"valid" json:"valid"`
	Whitelist    TokenWhitelistType `db:"whitelist" json:"whitelist"`
	Language     sql.NullString     `db:"language" json:"language"`
	LastUpdated  time.Time          `db:"last_updated" json:"lastUpdated"`
}

type TokenAuthorization struct {
	ID              int64          `db:"id" json:"id"`
	TokenID         int64          `db:"token_id" json:"tokenID"`
	AuthorizationID string         `db:"authorization_id" json:"authorizationID"`
	LocationID      sql.NullString `db:"location_id" json:"locationID"`
}

type TokenAuthorizationConnector struct {
	TokenAuthorizationID int64  `db:"token_authorization_id" json:"tokenAuthorizationID"`
	ConnectorUid         string `db:"connector_uid" json:"connectorUid"`
}

type TokenAuthorizationEvse struct {
	TokenAuthorizationID int64  `db:"token_authorization_id" json:"tokenAuthorizationID"`
	EvseUid              string `db:"evse_uid" json:"evseUid"`
}

type User struct {
	ID            int64         `db:"id" json:"id"`
	LinkingPubkey string        `db:"linking_pubkey" json:"linkingPubkey"`
	Pubkey        string        `db:"pubkey" json:"pubkey"`
	DeviceToken   string        `db:"device_token" json:"deviceToken"`
	NodeID        sql.NullInt64 `db:"node_id" json:"nodeID"`
	TokenID       sql.NullInt64 `db:"token_id" json:"tokenID"`
}

type Version struct {
	ID           int64  `db:"id" json:"id"`
	CredentialID int64  `db:"credential_id" json:"credentialID"`
	Version      string `db:"version" json:"version"`
	Url          string `db:"url" json:"url"`
}

type VersionEndpoint struct {
	ID         int64  `db:"id" json:"id"`
	VersionID  int64  `db:"version_id" json:"versionID"`
	Identifier string `db:"identifier" json:"identifier"`
	Url        string `db:"url" json:"url"`
}

type Weekday struct {
	ID          int64  `db:"id" json:"id"`
	Text        string `db:"text" json:"text"`
	Description string `db:"description" json:"description"`
}
