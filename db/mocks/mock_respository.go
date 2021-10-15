package mocks

import (
	"github.com/satimoto/go-datastore/db"
)

type MockRepository interface {
	db.Repository

	GetCreateBusinessDetailMockData() (db.CreateBusinessDetailParams, error)
	GetCreateConnectorMockData() (db.CreateConnectorParams, error)
	GetCreateCredentialMockData() (db.CreateCredentialParams, error)
	GetCreateDisplayTextMockData() (db.CreateDisplayTextParams, error)
	GetCreateEnergyMixMockData() (db.CreateEnergyMixParams, error)
	GetCreateEvseMockData() (db.CreateEvseParams, error)
	GetCreateGeoLocationMockData() (db.CreateGeoLocationParams, error)
	GetCreateImageMockData() (db.CreateImageParams, error)
	GetCreateNodeMockData() (db.CreateNodeParams, error)
	GetCreateOpeningTimeMockData() (bool, error)
	GetCreateStatusScheduleMockData() (db.CreateStatusScheduleParams, error)
	GetCreateUserMockData() (db.CreateUserParams, error)
	GetSetEvseDirectionMockData() (db.SetEvseDirectionParams, error)
	GetSetEvseImageMockData() (db.SetEvseImageParams, error)
	SetDeleteUserMockData(err error)
	SetGetBusinessDetailMockData(response BusinessDetailMockData)
	SetGetConnectorByUidMockData(response ConnectorMockData)
	SetGetCredentialByPartyAndCountryCodeMockData(response CredentialMockData)
	SetGetEnergyMixMockData(response EnergyMixMockData)
	SetGetEvseMockData(response EvseMockData)
	SetGetEvseByUidMockData(response EvseMockData)
	SetGetGeoLocationMockData(response GeoLocationMockData)
	SetGetImageMockData(response ImageMockData)
	SetGetLocationMockData(response LocationMockData)
	SetGetLocationByUidMockData(response LocationMockData)
	SetGetNodeMockData(response NodeMockData)
	SetGetOpeningTimeMockData(response OpeningTimeMockData)
	SetGetUserMockData(response UserMockData)
	SetListConnectorsMockData(response ConnectorsMockData)
	SetListEvsesMockData(response EvsesMockData)
	SetListEvseCapabilitiesMockData(response CapabilitiesMockData)
	SetListEvseDirectionsMockData(response DisplayTextsMockData)
	SetListEvseImagesMockData(response ImagesMockData)
	SetListEvseParkingRestrictionsMockData(response ParkingRestrictionsMockData)
	SetListExceptionalOpeningPeriodsMockData(response ExceptionalPeriodsMockData)
	SetListExceptionalClosingPeriodsMockData(response ExceptionalPeriodsMockData)
	SetListLocationDirectionsMockData(response DisplayTextsMockData)
	SetListLocationFacilitiesMockData(response FacilitiesMockData)
	SetListLocationImagesMockData(response ImagesMockData)
	SetListLocationsMockData(response LocationsMockData)
	SetListRegularHoursMockData(response RegularHoursMockData)
	SetListRelatedLocationsMockData(response GeoLocationsMockData)
	SetListStatusSchedulesMockData(response StatusSchedulesMockData)
	SetListUsersMockData(response UsersMockData)
	SetUpdateConnectorByUidMockData(response ConnectorMockData)
	SetUpdateEvseByUidMockData(response EvseMockData)
	SetUpdateEvseLastUpdatedMockData(err error)
	SetUpdateLocationByUidMockData(response LocationMockData)
	SetUpdateLocationLastUpdatedMockData(err error)
	SetUpdateUserMockData(response UserMockData)
}

type MockRepositoryService struct {
	createBusinessDetailMockData               []db.CreateBusinessDetailParams
	createConnectorMockData                    []db.CreateConnectorParams
	createCredentialMockData                   []db.CreateCredentialParams
	createDisplayTextMockData                  []db.CreateDisplayTextParams
	createEnergyMixMockData                    []db.CreateEnergyMixParams
	createEvseMockData                         []db.CreateEvseParams
	createGeoLocationMockData                  []db.CreateGeoLocationParams
	createImageMockData                        []db.CreateImageParams
	createNodeMockData                         []db.CreateNodeParams
	createOpeningTimeMockData                  []bool
	createStatusScheduleMockData               []db.CreateStatusScheduleParams
	createUserMockData                         []db.CreateUserParams
	deleteUserMockData                         []error
	getBusinessDetailMockData                  []BusinessDetailMockData
	getConnectorByUidMockData                  []ConnectorMockData
	getCredentialByPartyAndCountryCodeMockData []CredentialMockData
	getEnergyMixMockData                       []EnergyMixMockData
	getEvseMockData                            []EvseMockData
	getEvseByUidMockData                       []EvseMockData
	getGeoLocationMockData                     []GeoLocationMockData
	getImageMockData                           []ImageMockData
	getLocationMockData                        []LocationMockData
	getLocationByUidMockData                   []LocationMockData
	getNodeMockData                            []NodeMockData
	getOpeningTimeMockData                     []OpeningTimeMockData
	getUserMockData                            []UserMockData
	listConnectorsMockData                     []ConnectorsMockData
	listEvsesMockData                          []EvsesMockData
	listEvseCapabilitiesMockData               []CapabilitiesMockData
	listEvseDirectionsMockData                 []DisplayTextsMockData
	listEvseImagesMockData                     []ImagesMockData
	listEvseParkingRestrictionsMockData        []ParkingRestrictionsMockData
	listExceptionalOpeningPeriodsMockData      []ExceptionalPeriodsMockData
	listExceptionalClosingPeriodsMockData      []ExceptionalPeriodsMockData
	listLocationDirectionsMockData             []DisplayTextsMockData
	listLocationFacilitiesMockData             []FacilitiesMockData
	listLocationImagesMockData                 []ImagesMockData
	listLocationsMockData                      []LocationsMockData
	listRegularHoursMockData                   []RegularHoursMockData
	listRelatedLocationsMockData               []GeoLocationsMockData
	listStatusSchedulesMockData                []StatusSchedulesMockData
	listUsersMockData                          []UsersMockData
	setEvseDirectionMockData                   []db.SetEvseDirectionParams
	setEvseImageMockData                       []db.SetEvseImageParams
	updateConnectorByUidMockData               []ConnectorMockData
	updateEvseByUidMockData                    []EvseMockData
	updateEvseLastUpdatedMockData              []error
	updateLocationByUidMockData                []LocationMockData
	updateLocationLastUpdatedMockData              []error
	updateUserMockData                         []UserMockData
}

func NewMockRepository() MockRepository {
	return &MockRepositoryService{}
}
