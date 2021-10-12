package mocks

import (
	"github.com/satimoto/go-datastore/db"
)

type MockRepository interface {
	db.Repository

	SetCreateBusinessDetailPayload(response BusinessDetailPayload)
	SetCreateCredentialPayload(response CredentialPayload)
	SetCreateEnergyMixPayload(response EnergyMixPayload)
	SetCreateGeoLocationPayload(response GeoLocationPayload)
	SetCreateImagePayload(response ImagePayload)
	SetCreateNodePayload(response NodePayload)
	SetCreateOpeningTimePayload(response OpeningTimePayload)
	SetCreateUserPayload(response UserPayload)
	SetDeleteUserPayload(err error)
	SetGetBusinessDetailPayload(response BusinessDetailPayload)
	SetGetConnectorByUidPayload(response ConnectorPayload)
	SetGetCredentialByPartyAndCountryCodePayload(response CredentialPayload)
	SetGetEnergyMixPayload(response EnergyMixPayload)
	SetGetEvseByUidPayload(response EvsePayload)
	SetGetGeoLocationPayload(response GeoLocationPayload)
	SetGetImagePayload(response ImagePayload)
	SetGetLocationByUidPayload(response LocationPayload)
	SetGetNodePayload(response NodePayload)
	SetGetOpeningTimePayload(response OpeningTimePayload)
	SetGetUserPayload(response UserPayload)
	SetListConnectorsPayload(response ConnectorsPayload)
	SetListEvsesPayload(response EvsesPayload)
	SetListEvseCapabilitiesPayload(response CapabilitiesPayload)
	SetListEvseDirectionsPayload(response DisplayTextsPayload)
	SetListEvseImagesPayload(response ImagesPayload)
	SetListEvseParkingRestrictionsPayload(response ParkingRestrictionsPayload)
	SetListExceptionalOpeningPeriodsPayload(response ExceptionalPeriodsPayload)
	SetListExceptionalClosingPeriodsPayload(response ExceptionalPeriodsPayload)
	SetListLocationDirectionsPayload(response DisplayTextsPayload)
	SetListLocationFacilitiesPayload(response FacilitiesPayload)
	SetListLocationImagesPayload(response ImagesPayload)
	SetListLocationsPayload(response LocationsPayload)
	SetListRegularHoursPayload(response RegularHoursPayload)
	SetListRelatedLocationsPayload(response GeoLocationsPayload)
	SetListStatusSchedulesPayload(response StatusSchedulesPayload)
	SetListUsersPayload(response UsersPayload)
	SetUpdateConnectorByUidPayload(response ConnectorPayload)
	SetUpdateEvseByUidPayload(response EvsePayload)
	SetUpdateLocationByUidPayload(response LocationPayload)
	SetUpdateUserPayload(response UserPayload)
}

type MockRepositoryService struct {
	createBusinessDetailPayload               []BusinessDetailPayload
	createCredentialPayload                   []CredentialPayload
	createEnergyMixPayload                    []EnergyMixPayload
	createGeoLocationPayload                  []GeoLocationPayload
	createImagePayload                        []ImagePayload
	createNodePayload                         []NodePayload
	createOpeningTimePayload                  []OpeningTimePayload
	createUserPayload                         []UserPayload
	deleteUserPayload                         error
	getBusinessDetailPayload                  []BusinessDetailPayload
	getConnectorByUidPayload                  []ConnectorPayload
	getCredentialByPartyAndCountryCodePayload []CredentialPayload
	getEnergyMixPayload                       []EnergyMixPayload
	getEvseByUidPayload                       []EvsePayload
	getGeoLocationPayload                     []GeoLocationPayload
	getImagePayload                           []ImagePayload
	getLocationByUidPayload                   []LocationPayload
	getNodePayload                            []NodePayload
	getOpeningTimePayload                     []OpeningTimePayload
	getUserPayload                            []UserPayload
	listConnectorsPayload                     []ConnectorsPayload
	listEvsesPayload                          []EvsesPayload
	listEvseCapabilitiesPayload               []CapabilitiesPayload
	listEvseDirectionsPayload                 []DisplayTextsPayload
	listEvseImagesPayload                     []ImagesPayload
	listEvseParkingRestrictionsPayload        []ParkingRestrictionsPayload
	listExceptionalOpeningPeriodsPayload      []ExceptionalPeriodsPayload
	listExceptionalClosingPeriodsPayload      []ExceptionalPeriodsPayload
	listLocationDirectionsPayload             []DisplayTextsPayload
	listLocationFacilitiesPayload             []FacilitiesPayload
	listLocationImagesPayload                 []ImagesPayload
	listLocationsPayload                      []LocationsPayload
	listRegularHoursPayload                   []RegularHoursPayload
	listRelatedLocationsPayload               []GeoLocationsPayload
	listStatusSchedulesPayload                []StatusSchedulesPayload
	listUsersPayload                          []UsersPayload
	updateConnectorByUidPayload               []ConnectorPayload
	updateEvseByUidPayload                    []EvsePayload
	updateLocationByUidPayload                []LocationPayload
	updateUserPayload                         []UserPayload
}

func NewMockRepository() MockRepository {
	return &MockRepositoryService{}
}
