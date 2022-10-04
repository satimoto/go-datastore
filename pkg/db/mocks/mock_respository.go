package mocks

import (
	"database/sql"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
)

type MockRepository interface {
	GetCreateAdditionalGeoLocationMockData() (db.CreateAdditionalGeoLocationParams, error)
	GetCreateAuthenticationMockData() (db.CreateAuthenticationParams, error)
	GetCreateBusinessDetailMockData() (db.CreateBusinessDetailParams, error)
	GetCreateCalibrationMockData() (db.CreateCalibrationParams, error)
	GetCreateCalibrationValueMockData() (db.CreateCalibrationValueParams, error)
	GetCreateCdrMockData() (db.CreateCdrParams, error)
	GetCreateChannelRequestMockData() (db.CreateChannelRequestParams, error)
	GetCreateChannelRequestHtlcMockData() (db.CreateChannelRequestHtlcParams, error)
	GetCreateChargingPeriodMockData() (time.Time, error)
	GetCreateChargingPeriodDimensionMockData() (db.CreateChargingPeriodDimensionParams, error)
	GetCreateConnectorMockData() (db.CreateConnectorParams, error)
	GetCreateCommandReservationMockData() (db.CreateCommandReservationParams, error)
	GetCreateCommandStartMockData() (db.CreateCommandStartParams, error)
	GetCreateCommandStopMockData() (db.CreateCommandStopParams, error)
	GetCreateCommandUnlockMockData() (db.CreateCommandUnlockParams, error)
	GetCreateCountryAccountMockData() (db.CreateCountryAccountParams, error)
	GetCreateCredentialMockData() (db.CreateCredentialParams, error)
	GetCreateDisplayTextMockData() (db.CreateDisplayTextParams, error)
	GetCreateElementMockData() (db.CreateElementParams, error)
	GetCreateElementRestrictionMockData() (db.CreateElementRestrictionParams, error)
	GetCreateEmailSubscriptionMockData() (db.CreateEmailSubscriptionParams, error)
	GetCreateEnergyMixMockData() (db.CreateEnergyMixParams, error)
	GetCreateEnergySourceMockData() (db.CreateEnergySourceParams, error)
	GetCreateEnvironmentalImpactMockData() (db.CreateEnvironmentalImpactParams, error)
	GetCreateEvseMockData() (db.CreateEvseParams, error)
	GetCreateEvseStatusPeriodMockData() (db.CreateEvseStatusPeriodParams, error)
	GetCreateExceptionalPeriodMockData() (db.CreateExceptionalPeriodParams, error)
	GetCreateGeoLocationMockData() (db.CreateGeoLocationParams, error)
	GetCreateImageMockData() (db.CreateImageParams, error)
	GetCreateInvoiceRequestMockData() (db.CreateInvoiceRequestParams, error)
	GetCreateLocationMockData() (db.CreateLocationParams, error)
	GetCreateNodeMockData() (db.CreateNodeParams, error)
	GetCreateOpeningTimeMockData() (bool, error)
	CreatePendingNotification() (db.CreatePendingNotificationParams, error)
	GetCreatePriceComponentMockData() (db.CreatePriceComponentParams, error)
	GetCreatePriceComponentRoundingMockData() (db.CreatePriceComponentRoundingParams, error)
	GetCreatePromotionMockData() (db.CreatePromotionParams, error)
	GetCreatePsbtFundingStateMockData() (db.CreatePsbtFundingStateParams, error)
	GetCreateReferralMockData() (db.CreateReferralParams, error)
	GetCreateRegularHourMockData() (db.CreateRegularHourParams, error)
	GetCreateRoutingEventMockData() (db.CreateRoutingEventParams, error)
	GetCreateSessionMockData() (db.CreateSessionParams, error)
	GetCreateSessionInvoiceMockData() (db.CreateSessionInvoiceParams, error)
	GetCreateStatusScheduleMockData() (db.CreateStatusScheduleParams, error)
	GetCreateTariffMockData() (db.CreateTariffParams, error)
	GetCreateTariffRestrictionMockData() (db.CreateTariffRestrictionParams, error)
	GetCreateTokenMockData() (db.CreateTokenParams, error)
	GetCreateTokenAuthorizationMockData() (db.CreateTokenAuthorizationParams, error)
	GetCreateUserMockData() (db.CreateUserParams, error)
	GetCreateVersionMockData() (db.CreateVersionParams, error)
	GetCreateVersionEndpointMockData() (db.CreateVersionEndpointParams, error)
	GetDeleteAdditionalGeoLocationsMockData() (int64, error)
	GetDeleteAuthenticationMockData() (int64, error)
	GetDeleteBusinessDetailMockData() (int64, error)
	GetDeleteBusinessDetailLogoMockData() (int64, error)
	GetDeleteCdrChargingPeriodsMockData() (int64, error)
	GetDeleteChannelRequestMockData() (int64, error)
	GetDeleteChargingPeriodDimensionsMockData() (int64, error)
	GetDeleteConnectorsMockData() (int64, error)
	GetDeleteDisplayTextMockData() (int64, error)
	GetDeleteElementsMockData() (int64, error)
	GetDeleteElementRestrictionsMockData() (int64, error)
	GetDeleteEnergySourcesMockData() (int64, error)
	GetDeleteEnvironmentalImpactsMockData() (int64, error)
	GetDeleteEvseDirectionsMockData() (int64, error)
	GetDeleteEvseImagesMockData() (int64, error)
	GetDeleteExceptionalClosingPeriodsMockData() (int64, error)
	GetDeleteExceptionalOpeningPeriodsMockData() (int64, error)
	GetDeleteGeoLocationMockData() (int64, error)
	GetDeleteImageMockData() (int64, error)
	GetDeleteInvoiceRequestMockData() (int64, error)
	GetDeleteLocationDirectionsMockData() (int64, error)
	GetDeleteLocationImagesMockData() (int64, error)
	GetDeleteOpeningTimeMockData() (int64, error)
	GetDeletePendingNotificationMockData() (int64, error)
	GetDeletePendingNotificationByInvoiceRequestMockData() (int64, error)
	GetDeletePendingNotificationsMockData() ([]int64, error)
	GetDeletePriceComponentsMockData() (int64, error)
	GetDeletePriceComponentRoundingsMockData() (int64, error)
	GetDeleteRegularHoursMockData() (int64, error)
	GetDeleteRelatedLocationsMockData() (int64, error)
	GetDeleteSessionChargingPeriodsMockData() (int64, error)
	GetDeleteStatusScheduleMockData() (int64, error)
	GetDeleteTariffAltTextsMockData() (int64, error)
	GetDeleteTariffByUidMockData() (string, error)
	GetDeleteTariffRestrictionMockData() (int64, error)
	GetDeleteTokenByUidMockData() (string, error)
	GetDeleteVersionsMockData() (int64, error)
	GetDeleteVersionEndpointsMockData() (int64, error)
	GetSetCdrChargingPeriodMockData() (db.SetCdrChargingPeriodParams, error)
	GetSetElementRestrictionWeekdayMockData() (db.SetElementRestrictionWeekdayParams, error)
	GetSetEvseCapabilityMockData() (db.SetEvseCapabilityParams, error)
	GetSetEvseDirectionMockData() (db.SetEvseDirectionParams, error)
	GetSetEvseImageMockData() (db.SetEvseImageParams, error)
	GetSetEvseParkingRestrictionMockData() (db.SetEvseParkingRestrictionParams, error)
	GetSetLocationDirectionMockData() (db.SetLocationDirectionParams, error)
	GetSetLocationFacilityMockData() (db.SetLocationFacilityParams, error)
	GetSetLocationImageMockData() (db.SetLocationImageParams, error)
	GetSetPsbtFundingStateChannelRequestMockData() (db.SetPsbtFundingStateChannelRequestParams, error)
	GetSetSessionChargingPeriodMockData() (db.SetSessionChargingPeriodParams, error)
	GetSetTariffAltTextMockData() (db.SetTariffAltTextParams, error)
	GetSetTariffRestrictionWeekdayMockData() (db.SetTariffRestrictionWeekdayParams, error)
	GetSetTokenAuthorizationConnectorMockData() (db.SetTokenAuthorizationConnectorParams, error)
	GetSetTokenAuthorizationEvseMockData() (db.SetTokenAuthorizationEvseParams, error)
	GetUnsetElementRestrictionWeekdaysMockData() (int64, error)
	GetUnsetEvseCapabilitiesMockData() (int64, error)
	GetUnsetEvseParkingRestrictionsMockData() (int64, error)
	GetUnsetLocationFacilitiesMockData() (int64, error)
	GetUnsetTariffRestrictionWeekdaysMockData() (int64, error)
	GetUpdateAuthenticationMockData() (db.UpdateAuthenticationParams, error)
	GetUpdateBusinessDetailMockData() (db.UpdateBusinessDetailParams, error)
	GetUpdateConnectorByEvseMockData() (db.UpdateConnectorByEvseParams, error)
	GetUpdateChannelRequestMockData() (db.UpdateChannelRequestParams, error)
	GetUpdateChannelRequestHtlcByCircuitKeysMockData() (db.UpdateChannelRequestHtlcByCircuitKeyParams, error)
	GetUpdatePendingChannelRequestByPubkeyMockData() (db.UpdatePendingChannelRequestByPubkeyParams, error)
	GetUpdateCommandReservationMockData() (db.UpdateCommandReservationParams, error)
	GetUpdateCommandStartMockData() (db.UpdateCommandStartParams, error)
	GetUpdateCommandStopMockData() (db.UpdateCommandStopParams, error)
	GetUpdateCommandUnlockMockData() (db.UpdateCommandUnlockParams, error)
	GetUpdateCredentialMockData() (db.UpdateCredentialParams, error)
	GetUpdateElementRestrictionMockData() (db.UpdateElementRestrictionParams, error)
	GetUpdateEmailSubscriptionMockData() (db.UpdateEmailSubscriptionParams, error)
	GetUpdateEnergyMixMockData() (db.UpdateEnergyMixParams, error)
	GetUpdateEvseByUidMockData() (db.UpdateEvseByUidParams, error)
	GetUpdateEvseLastUpdatedMockData() (db.UpdateEvseLastUpdatedParams, error)
	GetUpdateGeoLocationMockData() (db.UpdateGeoLocationParams, error)
	GetUpdateImageMockData() (db.UpdateImageParams, error)
	GetUpdateInvoiceRequestMockData() (db.UpdateInvoiceRequestParams, error)
	GetUpdateLocationAvailabilityMockData() (db.UpdateLocationAvailabilityParams, error)
	GetUpdateLocationByUidMockData() (db.UpdateLocationByUidParams, error)
	GetUpdateLocationLastUpdatedMockData() (db.UpdateLocationLastUpdatedParams, error)
	GetUpdateLocationPublishMockData() (db.UpdateLocationPublishParams, error)
	GetUpdateLocationsPublishByCredentialMockData() (db.UpdateLocationsPublishByCredentialParams, error)
	GetUpdateLocationsPublishByPartyAndCountryCodeMockData() (db.UpdateLocationsPublishByPartyAndCountryCodeParams, error)
	GetUpdateNodeMockData() (db.UpdateNodeParams, error)
	GetUpdateOpeningTimeMockData() (db.UpdateOpeningTimeParams, error)
	GetUpdatePendingNotificationsMockData() (db.UpdatePendingNotificationsParams, error)
	GetUpdatePendingNotificationsByUserMockData() (db.UpdatePendingNotificationsByUserParams, error)
	GetUpdatePsbtFundingStateMockData() (db.UpdatePsbtFundingStateParams, error)
	GetUpdateRoutingEventMockData() (db.UpdateRoutingEventParams, error)
	GetUpdateSessionByUidMockData() (db.UpdateSessionByUidParams, error)
	GetUpdateSessionInvoiceMockData() (db.UpdateSessionInvoiceParams, error)
	GetUpdateTariffByUidMockData() (db.UpdateTariffByUidParams, error)
	GetUpdateTariffRestrictionMockData() (db.UpdateTariffRestrictionParams, error)
	GetUpdateTokenAuthorizationByAuthorizationIDMockData() (db.UpdateTokenAuthorizationByAuthorizationIDParams, error)
	GetUpdateTokenByUidMockData() (db.UpdateTokenByUidParams, error)
	GetUpdateUserMockData() (db.UpdateUserParams, error)
	GetUpdateUserByPubkeyMockData() (db.UpdateUserByPubkeyParams, error)
	SetGetAuthenticationByCodeMockData(response AuthenticationMockData)
	SetGetAuthenticationByChallengeMockData(response AuthenticationMockData)
	SetGetBusinessDetailMockData(response BusinessDetailMockData)
	SetGetCalibrationMockData(response CalibrationMockData)
	SetGetCdrByLastUpdatedMockData(response CdrMockData)
	SetGetCdrByUidMockData(response CdrMockData)
	SetGetChannelRequestMockData(response ChannelRequestMockData)
	SetGetChannelRequestByChannelPointMockData(response ChannelRequestMockData)
	SetGetChannelRequestByPaymentHashMockData(response ChannelRequestMockData)
	SetGetChannelRequestByPendingChanIdMockData(response ChannelRequestMockData)
	SetGetChannelRequestHtlcMockData(response ChannelRequestHtlcMockData)
	SetGetChannelRequestHtlcByCircuitKeyMockData(response ChannelRequestHtlcMockData)
	SetGetConnectorMockData(response ConnectorMockData)
	SetGetConnectorByIdentifierMockData(response ConnectorMockData)
	SetGetConnectorByEvseMockData(response ConnectorMockData)
	SetGetCommandReservationMockData(response CommandReservationMockData)
	SetGetCommandStartMockData(response CommandStartMockData)
	SetGetCommandStopMockData(response CommandStopMockData)
	SetGetCommandUnlockMockData(response CommandUnlockMockData)
	SetGetCountryAccountByCountryMockData(response CountryAccountMockData)
	SetGetCredentialMockData(response CredentialMockData)
	SetGetCredentialByPartyAndCountryCodeMockData(response CredentialMockData)
	SetGetCredentialByServerTokenMockData(response CredentialMockData)
	SetGetDisplayTextMockData(response DisplayTextMockData)
	SetGetElementRestrictionMockData(response ElementRestrictionMockData)
	SetGetEmailSubscriptionByEmailMockData(response EmailSubscriptionMockData)
	SetGetEnergyMixMockData(response EnergyMixMockData)
	SetGetEvseMockData(response EvseMockData)
	SetGetEvseByIdentifierMockData(response EvseMockData)
	SetGetEvseByUidMockData(response EvseMockData)
	SetGetGeoLocationMockData(response GeoLocationMockData)
	SetGetImageMockData(response ImageMockData)
	SetGetInvoiceRequestMockData(response InvoiceRequestMockData)
	SetGetUnsettledInvoiceRequestMockData(response InvoiceRequestMockData)
	SetGetLocationMockData(response LocationMockData)
	SetGetLocationByLastUpdatedMockData(response LocationMockData)
	SetGetLocationByUidMockData(response LocationMockData)
	SetGetNodeMockData(response NodeMockData)
	SetGetNodeByPubkeyMockData(response NodeMockData)
	SetGetNodeByUserIDMockData(response NodeMockData)
	SetGetOpeningTimeMockData(response OpeningTimeMockData)
	SetGetPriceComponentRoundingMockData(response PriceComponentRoundingMockData)
	SetGetPromotionMockData(response PromotionMockData)
	SetGetPromotionByCodeMockData(response PromotionMockData)
	SetGetPsbtFundingStateMockData(response PsbtFundingStateMockData)
	SetGetUnfundedPsbtFundingStateMockData(response PsbtFundingStateMockData)
	SetGetReferralByIpAddressMockData(response ReferralMockData)
	SetGetSessionMockData(response SessionMockData)
	SetGetSessionByAuthorizationIDMockData(response SessionMockData)
	SetGetSessionByLastUpdatedMockData(response SessionMockData)
	SetGetSessionByUidMockData(response SessionMockData)
	SetGetSessionInvoiceMockData(response SessionInvoiceMockData)
	SetGetSessionInvoiceByPaymentRequestMockData(response SessionInvoiceMockData)
	SetGetTariffMockData(response TariffMockData)
	SetGetTariffByLastUpdatedMockData(response TariffMockData)
	SetGetTariffByUidMockData(response TariffMockData)
	SetGetTariffRestrictionMockData(response TariffRestrictionMockData)
	SetGetTokenAuthorizationByAuthorizationIDMockData(response TokenAuthorizationMockData)
	SetGetTokenMockData(response TokenMockData)
	SetGetTokenByAuthIDMockData(response TokenMockData)
	SetGetTokenByUidMockData(response TokenMockData)
	SetGetTokenByUserIDMockData(response TokenMockData)
	SetGetUserMockData(response UserMockData)
	SetGetUserByDeviceTokenMockData(response UserMockData)
	SetGetUserByLinkingPubkeyMockData(response UserMockData)
	SetGetUserByPubkeyMockData(response UserMockData)
	SetGetUserByReferralCodeMockData(response UserMockData)
	SetGetUserBySessionIDMockData(response UserMockData)
	SetGetUserByTokenIDMockData(response UserMockData)
	SetGetVersionMockData(response VersionMockData)
	SetGetVersionEndpointMockData(response VersionEndpointMockData)
	SetGetVersionEndpointByIdentityMockData(response VersionEndpointMockData)
	SetListAdditionalGeoLocationsMockData(response AdditionalGeoLocationsMockData)
	SetListCalibrationValuesMockData(response CalibrationValuesMockData)
	SetListCapabilitiesMockData(response CapabilitiesMockData)
	SetListCdrChargingPeriodsMockData(response ChargingPeriodsMockData)
	SetListChannelRequestHtlcsMockData(response ChannelRequestHtlcsMockData)
	SetListChargingPeriodDimensionsMockData(response ChargingPeriodDimensionsMockData)
	SetListCredentialsMockData(response CredentialsMockData)
	SetListConnectorsMockData(response ConnectorsMockData)
	SetListElementsMockData(response ElementsMockData)
	SetListElementRestrictionWeekdaysMockData(response WeekdaysMockData)
	SetListEnergySourcesMockData(response EnergySourcesMockData)
	SetListEnvironmentalImpactsMockData(response EnvironmentalImpactsMockData)
	SetListEvsesMockData(response EvsesMockData)
	SetListEvseStatusPeriodsMockData(response EvseStatusPeriodsMockData)
	SetListActiveEvsesMockData(response EvsesMockData)
	SetListEvseCapabilitiesMockData(response CapabilitiesMockData)
	SetListEvseDirectionsMockData(response DisplayTextsMockData)
	SetListEvseImagesMockData(response ImagesMockData)
	SetListEvseParkingRestrictionsMockData(response ParkingRestrictionsMockData)
	SetListExceptionalOpeningPeriodsMockData(response ExceptionalPeriodsMockData)
	SetListExceptionalClosingPeriodsMockData(response ExceptionalPeriodsMockData)
	SetListFacilitiesMockData(response FacilitiesMockData)
	SetListInvoiceRequestsMockData(response FacilitiesMockData)
	SetListLocationDirectionsMockData(response DisplayTextsMockData)
	SetListLocationFacilitiesMockData(response FacilitiesMockData)
	SetListLocationImagesMockData(response ImagesMockData)
	SetListLocationsMockData(response LocationsMockData)
	SetListLocationsByGeomMockData(response LocationsMockData)
	SetListUnfundedPsbtFundingStatesMockData(response PsbtFundingStatesMockData)
	SetListPsbtFundingStateChannelRequestsMockData(response ChannelRequestsMockData)
	SetListNodesMockData(response NodesMockData)
	SetListActiveNodesMockData(response NodesMockData)
	SetListParkingRestrictionsMockData(response ParkingRestrictionsMockData)
	SetListPendingNotificationsMockData(response PendingNotificationsMockData)
	SetListPriceComponentsMockData(response PriceComponentsMockData)
	SetListRegularHoursMockData(response RegularHoursMockData)
	SetListRelatedLocationsMockData(response GeoLocationsMockData)
	SetListSessionChargingPeriodsMockData(response ChargingPeriodsMockData)
	SetListSessionInvoicesMockData(response SessionInvoicesMockData)
	SetListUnsettledSessionInvoicesByUserIDMockData(response SessionInvoicesMockData)
	SetListStatusSchedulesMockData(response StatusSchedulesMockData)
	SetListTariffAltTextsMockData(response DisplayTextsMockData)
	SetListTariffRestrictionWeekdaysMockData(response WeekdaysMockData)
	SetListTariffsByCdrMockData(response TariffsMockData)
	SetListTokensMockData(response TokensMockData)
	SetListTokensByUserIDMockData(response TokensMockData)
	SetListTokenAuthorizationConnectorsMockData(response ConnectorsMockData)
	SetListTokenAuthorizationEvsesMockData(response EvsesMockData)
	SetListVersionsMockData(response VersionsMockData)
	SetListVersionEndpointsMockData(response VersionEndpointsMockData)
	SetListWeekdaysMockData(response WeekdaysMockData)
}

type MockRepositoryService struct {
	createAdditionalGeoLocationMockData                 []db.CreateAdditionalGeoLocationParams
	createAuthenticationMockData                        []db.CreateAuthenticationParams
	createBusinessDetailMockData                        []db.CreateBusinessDetailParams
	createCalibrationMockData                           []db.CreateCalibrationParams
	createCalibrationValueMockData                      []db.CreateCalibrationValueParams
	createCdrMockData                                   []db.CreateCdrParams
	createChannelRequestMockData                        []db.CreateChannelRequestParams
	createChannelRequestHtlcMockData                    []db.CreateChannelRequestHtlcParams
	createChargingPeriodMockData                        []time.Time
	createChargingPeriodDimensionMockData               []db.CreateChargingPeriodDimensionParams
	createConnectorMockData                             []db.CreateConnectorParams
	createCommandReservationMockData                    []db.CreateCommandReservationParams
	createCommandStartMockData                          []db.CreateCommandStartParams
	createCommandStopMockData                           []db.CreateCommandStopParams
	createCommandUnlockMockData                         []db.CreateCommandUnlockParams
	createCountryAccountMockData                        []db.CreateCountryAccountParams
	createCredentialMockData                            []db.CreateCredentialParams
	createDisplayTextMockData                           []db.CreateDisplayTextParams
	createElementMockData                               []db.CreateElementParams
	createElementRestrictionMockData                    []db.CreateElementRestrictionParams
	createEmailSubscriptionMockData                     []db.CreateEmailSubscriptionParams
	createEnergyMixMockData                             []db.CreateEnergyMixParams
	createEnergySourceMockData                          []db.CreateEnergySourceParams
	createEnvironmentalImpactMockData                   []db.CreateEnvironmentalImpactParams
	createEvseMockData                                  []db.CreateEvseParams
	createEvseStatusPeriodMockData                      []db.CreateEvseStatusPeriodParams
	createExceptionalPeriodMockData                     []db.CreateExceptionalPeriodParams
	createGeoLocationMockData                           []db.CreateGeoLocationParams
	createImageMockData                                 []db.CreateImageParams
	createInvoiceRequestMockData                        []db.CreateInvoiceRequestParams
	createLocationMockData                              []db.CreateLocationParams
	createOpeningTimeMockData                           []bool
	createNodeMockData                                  []db.CreateNodeParams
	createPendingNotificationMockData                   []db.CreatePendingNotificationParams
	createPriceComponentMockData                        []db.CreatePriceComponentParams
	createPriceComponentRoundingMockData                []db.CreatePriceComponentRoundingParams
	createPromotionMockData                             []db.CreatePromotionParams
	createPsbtFundingStateMockData                      []db.CreatePsbtFundingStateParams
	createReferralMockData                              []db.CreateReferralParams
	createRegularHourMockData                           []db.CreateRegularHourParams
	createRoutingEventMockData                          []db.CreateRoutingEventParams
	createSessionMockData                               []db.CreateSessionParams
	createSessionInvoiceMockData                        []db.CreateSessionInvoiceParams
	createStatusScheduleMockData                        []db.CreateStatusScheduleParams
	createTariffMockData                                []db.CreateTariffParams
	createTariffRestrictionMockData                     []db.CreateTariffRestrictionParams
	createTokenMockData                                 []db.CreateTokenParams
	createTokenAuthorizationMockData                    []db.CreateTokenAuthorizationParams
	createUserMockData                                  []db.CreateUserParams
	createVersionMockData                               []db.CreateVersionParams
	createVersionEndpointMockData                       []db.CreateVersionEndpointParams
	deleteAdditionalGeoLocationsMockData                []int64
	deleteAuthenticationMockData                        []int64
	deleteBusinessDetailMockData                        []int64
	deleteBusinessDetailLogoMockData                    []int64
	deleteCdrChargingPeriodsMockData                    []int64
	deleteChannelRequestMockData                        []int64
	deleteChargingPeriodDimensionsMockData              []int64
	deleteConnectorsMockData                            []int64
	deleteDisplayTextMockData                           []int64
	deleteElementsMockData                              []int64
	deleteElementRestrictionsMockData                   []int64
	deleteEnergySourcesMockData                         []int64
	deleteEnvironmentalImpactsMockData                  []int64
	deleteEvseDirectionsMockData                        []int64
	deleteEvseImagesMockData                            []int64
	deleteExceptionalClosingPeriodsMockData             []int64
	deleteExceptionalOpeningPeriodsMockData             []int64
	deleteGeoLocationMockData                           []int64
	deleteImageMockData                                 []int64
	deleteInvoiceRequestMockData                        []int64
	deleteLocationDirectionsMockData                    []int64
	deleteLocationImagesMockData                        []int64
	deleteOpeningTimeMockData                           []int64
	deletePendingNotificationMockData                   []int64
	deletePendingNotificationByInvoiceRequestMockData   []sql.NullInt64
	deletePendingNotificationsMockData                  [][]int64
	deletePriceComponentsMockData                       []int64
	deletePriceComponentRoundingsMockData               []int64
	deleteRegularHoursMockData                          []int64
	deleteRelatedLocationsMockData                      []int64
	deleteSessionChargingPeriodsMockData                []int64
	deleteStatusScheduleMockData                        []int64
	deleteTariffByUidMockData                           []string
	deleteTariffAltTextsMockData                        []int64
	deleteTariffRestrictionMockData                     []int64
	deleteTokenByUidMockData                            []string
	deleteVersionsMockData                              []int64
	deleteVersionEndpointsMockData                      []int64
	getAuthenticationByCodeMockData                     []AuthenticationMockData
	getAuthenticationByChallengeMockData                []AuthenticationMockData
	getBusinessDetailMockData                           []BusinessDetailMockData
	getCalibrationMockData                              []CalibrationMockData
	getCdrByLastUpdatedMockData                         []CdrMockData
	getCdrByUidMockData                                 []CdrMockData
	getChannelRequestMockData                           []ChannelRequestMockData
	getChannelRequestByChannelPointMockData             []ChannelRequestMockData
	getChannelRequestByPaymentHashMockData              []ChannelRequestMockData
	getChannelRequestByPendingChanIdMockData            []ChannelRequestMockData
	getChannelRequestHtlcMockData                       []ChannelRequestHtlcMockData
	getChannelRequestHtlcByCircuitKeyMockData           []ChannelRequestHtlcMockData
	getConnectorMockData                                []ConnectorMockData
	getConnectorByIdentifierMockData                    []ConnectorMockData
	getConnectorByEvseMockData                          []ConnectorMockData
	getCommandReservationMockData                       []CommandReservationMockData
	getCommandStartMockData                             []CommandStartMockData
	getCommandStopMockData                              []CommandStopMockData
	getCommandUnlockMockData                            []CommandUnlockMockData
	getCountryAccountByCountryMockData                  []CountryAccountMockData
	getCredentialMockData                               []CredentialMockData
	getCredentialByPartyAndCountryCodeMockData          []CredentialMockData
	getCredentialByServerTokenMockData                  []CredentialMockData
	getDisplayTextMockData                              []DisplayTextMockData
	getElementRestrictionMockData                       []ElementRestrictionMockData
	getEmailSubscriptionByEmailMockData                 []EmailSubscriptionMockData
	getEnergyMixMockData                                []EnergyMixMockData
	getEvseMockData                                     []EvseMockData
	getEvseByIdentifierMockData                         []EvseMockData
	getEvseByUidMockData                                []EvseMockData
	getGeoLocationMockData                              []GeoLocationMockData
	getImageMockData                                    []ImageMockData
	getInvoiceRequestMockData                           []InvoiceRequestMockData
	getUnsettledInvoiceRequestMockData                  []InvoiceRequestMockData
	getLocationMockData                                 []LocationMockData
	getLocationByLastUpdatedMockData                    []LocationMockData
	getLocationByUidMockData                            []LocationMockData
	getNodeMockData                                     []NodeMockData
	getNodeByPubkeyMockData                             []NodeMockData
	getNodeByUserIDMockData                             []NodeMockData
	getOpeningTimeMockData                              []OpeningTimeMockData
	getPriceComponentRoundingMockData                   []PriceComponentRoundingMockData
	getPromotionMockData                                []PromotionMockData
	getPromotionByCodeMockData                          []PromotionMockData
	getPsbtFundingStateMockData                         []PsbtFundingStateMockData
	getUnfundedPsbtFundingStateMockData                 []PsbtFundingStateMockData
	getReferralByIpAddressMockData                      []ReferralMockData
	getSessionMockData                                  []SessionMockData
	getSessionByAuthorizationIDMockData                 []SessionMockData
	getSessionByLastUpdatedMockData                     []SessionMockData
	getSessionByUidMockData                             []SessionMockData
	getSessionInvoiceMockData                           []SessionInvoiceMockData
	getSessionInvoiceByPaymentRequestMockData           []SessionInvoiceMockData
	getTariffByLastUpdatedMockData                      []TariffMockData
	getTariffMockData                                   []TariffMockData
	getTariffByUidMockData                              []TariffMockData
	getTariffRestrictionMockData                        []TariffRestrictionMockData
	getTokenAuthorizationByAuthorizationIDMockData      []TokenAuthorizationMockData
	getTokenMockData                                    []TokenMockData
	getTokenByAuthIDMockData                            []TokenMockData
	getTokenByUidMockData                               []TokenMockData
	getTokenByUserIDMockData                            []TokenMockData
	getUserMockData                                     []UserMockData
	getUserByDeviceTokenMockData                        []UserMockData
	getUserByLinkingPubkeyMockData                      []UserMockData
	getUserByPubkeyMockData                             []UserMockData
	getUserByReferralCodeMockData                       []UserMockData
	getUserBySessionIDMockData                          []UserMockData
	getUserByTokenIDMockData                            []UserMockData
	getVersionMockData                                  []VersionMockData
	getVersionEndpointMockData                          []VersionEndpointMockData
	getVersionEndpointByIdentityMockData                []VersionEndpointMockData
	listAdditionalGeoLocationsMockData                  []AdditionalGeoLocationsMockData
	listCalibrationValuesMockData                       []CalibrationValuesMockData
	listCapabilitiesMockData                            []CapabilitiesMockData
	listCdrChargingPeriodsMockData                      []ChargingPeriodsMockData
	listChannelRequestHtlcsMockData                     []ChannelRequestHtlcsMockData
	listChargingPeriodDimensionsMockData                []ChargingPeriodDimensionsMockData
	listCredentialsMockData                             []CredentialsMockData
	listConnectorsMockData                              []ConnectorsMockData
	listElementsMockData                                []ElementsMockData
	listElementRestrictionWeekdaysMockData              []WeekdaysMockData
	listEnergySourcesMockData                           []EnergySourcesMockData
	listEnvironmentalImpactsMockData                    []EnvironmentalImpactsMockData
	listEvsesMockData                                   []EvsesMockData
	listEvseStatusPeriodsMockData                       []EvseStatusPeriodsMockData
	listActiveEvsesMockData                             []EvsesMockData
	listEvseCapabilitiesMockData                        []CapabilitiesMockData
	listEvseDirectionsMockData                          []DisplayTextsMockData
	listEvseImagesMockData                              []ImagesMockData
	listEvseParkingRestrictionsMockData                 []ParkingRestrictionsMockData
	listExceptionalOpeningPeriodsMockData               []ExceptionalPeriodsMockData
	listExceptionalClosingPeriodsMockData               []ExceptionalPeriodsMockData
	listFacilitiesMockData                              []FacilitiesMockData
	listUnsettledInvoiceRequestsMockData                []InvoiceRequestsMockData
	listLocationDirectionsMockData                      []DisplayTextsMockData
	listLocationFacilitiesMockData                      []FacilitiesMockData
	listLocationImagesMockData                          []ImagesMockData
	listLocationsMockData                               []LocationsMockData
	listLocationsByGeomMockData                         []LocationsMockData
	listNodesMockData                                   []NodesMockData
	listActiveNodesMockData                             []NodesMockData
	listParkingRestrictionsMockData                     []ParkingRestrictionsMockData
	listPendingNotificationsMockData                    []PendingNotificationsMockData
	listPriceComponentsMockData                         []PriceComponentsMockData
	listUnfundedPsbtFundingStatesMockData               []PsbtFundingStatesMockData
	listPsbtFundingStateChannelRequestsMockData         []ChannelRequestsMockData
	listRegularHoursMockData                            []RegularHoursMockData
	listRelatedLocationsMockData                        []GeoLocationsMockData
	listSessionChargingPeriodsMockData                  []ChargingPeriodsMockData
	listSessionInvoicesMockData                         []SessionInvoicesMockData
	listUnsettledSessionInvoicesByUserIDMockData        []SessionInvoicesMockData
	listStatusSchedulesMockData                         []StatusSchedulesMockData
	listTariffAltTextsMockData                          []DisplayTextsMockData
	listTariffRestrictionWeekdaysMockData               []WeekdaysMockData
	listTariffsByCdrMockData                            []TariffsMockData
	listTokensMockData                                  []TokensMockData
	listTokensByUserIDMockData                          []TokensMockData
	listTokenAuthorizationConnectorsMockData            []ConnectorsMockData
	listTokenAuthorizationEvsesMockData                 []EvsesMockData
	listVersionsMockData                                []VersionsMockData
	listVersionEndpointsMockData                        []VersionEndpointsMockData
	listWeekdaysMockData                                []WeekdaysMockData
	setCdrChargingPeriodMockData                        []db.SetCdrChargingPeriodParams
	setElementRestrictionWeekdayMockData                []db.SetElementRestrictionWeekdayParams
	setEvseCapabilityMockData                           []db.SetEvseCapabilityParams
	setEvseDirectionMockData                            []db.SetEvseDirectionParams
	setEvseImageMockData                                []db.SetEvseImageParams
	setEvseParkingRestrictionMockData                   []db.SetEvseParkingRestrictionParams
	setLocationDirectionMockData                        []db.SetLocationDirectionParams
	setLocationFacilityMockData                         []db.SetLocationFacilityParams
	setLocationImageMockData                            []db.SetLocationImageParams
	setPsbtFundingStateChannelRequestMockData           []db.SetPsbtFundingStateChannelRequestParams
	setSessionChargingPeriodMockData                    []db.SetSessionChargingPeriodParams
	setTariffAltTextMockData                            []db.SetTariffAltTextParams
	setTariffRestrictionWeekdayMockData                 []db.SetTariffRestrictionWeekdayParams
	setTokenAuthorizationConnectorMockData              []db.SetTokenAuthorizationConnectorParams
	setTokenAuthorizationEvseMockData                   []db.SetTokenAuthorizationEvseParams
	unsetElementRestrictionWeekdaysMockData             []int64
	unsetEvseCapabilitiesMockData                       []int64
	unsetEvseParkingRestrictionsMockData                []int64
	unsetLocationFacilitiesMockData                     []int64
	unsetTariffRestrictionWeekdaysMockData              []int64
	updateAuthenticationMockData                        []db.UpdateAuthenticationParams
	updateBusinessDetailMockData                        []db.UpdateBusinessDetailParams
	updateChannelRequestMockData                        []db.UpdateChannelRequestParams
	updateChannelRequestHtlcByCircuitKeyMockData        []db.UpdateChannelRequestHtlcByCircuitKeyParams
	updatePendingChannelRequestByPubkeyMockData         []db.UpdatePendingChannelRequestByPubkeyParams
	updateConnectorByEvseMockData                       []db.UpdateConnectorByEvseParams
	updateCommandReservationMockData                    []db.UpdateCommandReservationParams
	updateCommandStartMockData                          []db.UpdateCommandStartParams
	updateCommandStopMockData                           []db.UpdateCommandStopParams
	updateCommandUnlockMockData                         []db.UpdateCommandUnlockParams
	updateCredentialMockData                            []db.UpdateCredentialParams
	updateElementRestrictionMockData                    []db.UpdateElementRestrictionParams
	updateEmailSubscriptionMockData                     []db.UpdateEmailSubscriptionParams
	updateEnergyMixMockData                             []db.UpdateEnergyMixParams
	updateEvseByUidMockData                             []db.UpdateEvseByUidParams
	updateEvseLastUpdatedMockData                       []db.UpdateEvseLastUpdatedParams
	updateGeoLocationMockData                           []db.UpdateGeoLocationParams
	updateImageMockData                                 []db.UpdateImageParams
	updateInvoiceRequestMockData                        []db.UpdateInvoiceRequestParams
	updateLocationAvailabilityMockData                  []db.UpdateLocationAvailabilityParams
	updateLocationByUidMockData                         []db.UpdateLocationByUidParams
	updateLocationLastUpdatedMockData                   []db.UpdateLocationLastUpdatedParams
	updateLocationPublishMockData                       []db.UpdateLocationPublishParams
	updateLocationsPublishByCredentialMockData          []db.UpdateLocationsPublishByCredentialParams
	updateLocationsPublishByPartyAndCountryCodeMockData []db.UpdateLocationsPublishByPartyAndCountryCodeParams
	updateNodeMockData                                  []db.UpdateNodeParams
	updateOpeningTimeMockData                           []db.UpdateOpeningTimeParams
	updatePendingNotificationsMockData                  []db.UpdatePendingNotificationsParams
	updatePendingNotificationsByUserMockData            []db.UpdatePendingNotificationsByUserParams
	updatePsbtFundingStateMockData                      []db.UpdatePsbtFundingStateParams
	updateUserMockData                                  []db.UpdateUserParams
	updateUserByPubkeyMockData                          []db.UpdateUserByPubkeyParams
	updateRoutingEventMockData                          []db.UpdateRoutingEventParams
	updateSessionByUidMockData                          []db.UpdateSessionByUidParams
	updateSessionInvoiceMockData                        []db.UpdateSessionInvoiceParams
	updateTariffByUidMockData                           []db.UpdateTariffByUidParams
	updateTariffRestrictionMockData                     []db.UpdateTariffRestrictionParams
	updateTokenByUidMockData                            []db.UpdateTokenByUidParams
	updateTokenAuthorizationByAuthorizationIDMockData   []db.UpdateTokenAuthorizationByAuthorizationIDParams
}

func NewMockRepositoryService() *MockRepositoryService {
	return &MockRepositoryService{}
}
