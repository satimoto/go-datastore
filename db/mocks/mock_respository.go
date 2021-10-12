package mocks

import (
	"github.com/satimoto/go-datastore/db"
)

type MockRepository interface {
	db.Repository

	SetGetConnectorByUidPayload(response ConnectorPayload)
	SetListConnectorsPayload(response ConnectorsPayload)
	SetCreateCredentialPayload(response CredentialPayload)
	SetGetCredentialByPartyAndCountryCodePayload(response CredentialPayload)
	SetCreateEnergyMixPayload(response EnergyMixPayload)
	SetGetEnergyMixPayload(response EnergyMixPayload)
	SetCreateBusinessDetailPayload(response BusinessDetailPayload)
	SetGetBusinessDetailPayload(response BusinessDetailPayload)
	SetGetEvseByUidPayload(response EvsePayload)
	SetListEvsesPayload(response EvsesPayload)
	SetListEvseCapabilitiesPayload(response CapabilitiesPayload)
	SetListEvseDirectionsPayload(response DisplayTextsPayload)
	SetListEvseImagesPayload(response ImagesPayload)
	SetListEvseParkingRestrictionsPayload(response ParkingRestrictionsPayload)
	SetListExceptionalOpeningPeriodsPayload(response ExceptionalPeriodsPayload)
	SetListExceptionalClosingPeriodsPayload(response ExceptionalPeriodsPayload)
	SetCreateGeoLocationPayload(response GeoLocationPayload)
	SetGetGeoLocationPayload(response GeoLocationPayload)
	SetCreateImagePayload(response ImagePayload)
	SetGetImagePayload(response ImagePayload)
	SetGetLocationByUidPayload(response LocationPayload)
	SetListLocationDirectionsPayload(response DisplayTextsPayload)
	SetListLocationFacilitiesPayload(response FacilitiesPayload)
	SetListLocationImagesPayload(response ImagesPayload)
	SetListLocationsPayload(response LocationsPayload)
	SetCreateNodePayload(response NodePayload)
	SetGetNodePayload(response NodePayload)
	SetCreateOpeningTimePayload(response OpeningTimePayload)
	SetGetOpeningTimePayload(response OpeningTimePayload)
	SetListRegularHoursPayload(response RegularHoursPayload)
	SetListRelatedLocationsPayload(response GeoLocationsPayload)
	SetListStatusSchedulesPayload(response StatusSchedulesPayload)
	SetCreateUserPayload(response UserPayload)
	SetDeleteUserPayload(err error)
	SetGetUserPayload(response UserPayload)
	SetListUsersPayload(response UsersPayload)
	SetUpdateUserPayload(response UserPayload)
}

type MockRepositoryService struct {
	getConnectorByUidPayload                  []ConnectorPayload
	listConnectorsPayload                     []ConnectorsPayload
	createCredentialPayload                   []CredentialPayload
	getCredentialByPartyAndCountryCodePayload []CredentialPayload
	createEnergyMixPayload                    []EnergyMixPayload
	getEnergyMixPayload                       []EnergyMixPayload
	createBusinessDetailPayload               []BusinessDetailPayload
	getBusinessDetailPayload                  []BusinessDetailPayload
	getEvseByUidPayload                       []EvsePayload
	listEvsesPayload                          []EvsesPayload
	listEvseCapabilitiesPayload               []CapabilitiesPayload
	listEvseDirectionsPayload                 []DisplayTextsPayload
	listEvseImagesPayload                     []ImagesPayload
	listEvseParkingRestrictionsPayload        []ParkingRestrictionsPayload
	listExceptionalOpeningPeriodsPayload      []ExceptionalPeriodsPayload
	listExceptionalClosingPeriodsPayload      []ExceptionalPeriodsPayload
	createGeoLocationPayload                  []GeoLocationPayload
	getGeoLocationPayload                     []GeoLocationPayload
	createImagePayload                        []ImagePayload
	getImagePayload                           []ImagePayload
	getLocationByUidPayload                   []LocationPayload
	listLocationDirectionsPayload             []DisplayTextsPayload
	listLocationFacilitiesPayload             []FacilitiesPayload
	listLocationImagesPayload                 []ImagesPayload
	listLocationsPayload                      []LocationsPayload
	createNodePayload                         []NodePayload
	getNodePayload                            []NodePayload
	createOpeningTimePayload                  []OpeningTimePayload
	getOpeningTimePayload                     []OpeningTimePayload
	listRegularHoursPayload                   []RegularHoursPayload
	listRelatedLocationsPayload               []GeoLocationsPayload
	listStatusSchedulesPayload                []StatusSchedulesPayload
	createUserPayload                         []UserPayload
	deleteUserPayload                         error
	getUserPayload                            []UserPayload
	listUsersPayload                          []UsersPayload
	updateUserPayload                         []UserPayload
}

func NewMockRepository() MockRepository {
	return &MockRepositoryService{}
}
