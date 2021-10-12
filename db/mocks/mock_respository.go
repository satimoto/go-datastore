package mocks

import (
	"github.com/satimoto/go-datastore/db"
)

type MockRepository interface {
	db.Repository

	SetListConnectorsResponse(response ConnectorsResponse)
	SetCreateCredentialResponse(response CredentialResponse)
	SetGetCredentialByPartyAndCountryCodeResponse(response CredentialResponse)
	SetCreateEnergyMixResponse(response EnergyMixResponse)
	SetGetEnergyMixResponse(response EnergyMixResponse)
	SetCreateBusinessDetailResponse(response BusinessDetailResponse)
	SetGetBusinessDetailResponse(response BusinessDetailResponse)
	SetListEvsesResponse(response EvsesResponse)
	SetListEvseCapabilitiesResponse(response CapabilitiesResponse)
	SetListEvseDirectionsResponse(response DisplayTextsResponse)
	SetListEvseImagesResponse(response ImagesResponse)
	SetListEvseParkingRestrictionsResponse(response ParkingRestrictionsResponse)
	SetListExceptionalOpeningPeriodsResponse(response ExceptionalPeriodsResponse)
	SetListExceptionalClosingPeriodsResponse(response ExceptionalPeriodsResponse)
	SetCreateGeoLocationResponse(response GeoLocationResponse)
	SetGetGeoLocationResponse(response GeoLocationResponse)
	SetCreateImageResponse(response ImageResponse)
	SetGetImageResponse(response ImageResponse)
	SetListLocationDirectionsResponse(response DisplayTextsResponse)
	SetListLocationFacilitiesResponse(response FacilitiesResponse)
	SetListLocationImagesResponse(response ImagesResponse)
	SetListLocationsResponse(response LocationsResponse)
	SetCreateNodeResponse(response NodeResponse)
	SetGetNodeResponse(response NodeResponse)
	SetCreateOpeningTimeResponse(response OpeningTimeResponse)
	SetGetOpeningTimeResponse(response OpeningTimeResponse)
	SetListRegularHoursResponse(response RegularHoursResponse)
	SetListRelatedLocationsResponse(response GeoLocationsResponse)
	SetListStatusSchedulesResponse(response StatusSchedulesResponse)
	SetCreateUserResponse(response UserResponse)
	SetDeleteUserResponse(err error)
	SetGetUserResponse(response UserResponse)
	SetListUsersResponse(response UsersResponse)
	SetUpdateUserResponse(response UserResponse)
}

type MockRepositoryService struct {
	listConnectorsResponse                     []ConnectorsResponse
	createCredentialResponse                   []CredentialResponse
	getCredentialByPartyAndCountryCodeResponse []CredentialResponse
	createEnergyMixResponse                    []EnergyMixResponse
	getEnergyMixResponse                       []EnergyMixResponse
	createBusinessDetailResponse               []BusinessDetailResponse
	getBusinessDetailResponse                  []BusinessDetailResponse
	listEvsesResponse                          []EvsesResponse
	listEvseCapabilitiesResponse               []CapabilitiesResponse
	listEvseDirectionsResponse                 []DisplayTextsResponse
	listEvseImagesResponse                     []ImagesResponse
	listEvseParkingRestrictionsResponse        []ParkingRestrictionsResponse
	listExceptionalOpeningPeriodsResponse      []ExceptionalPeriodsResponse
	listExceptionalClosingPeriodsResponse      []ExceptionalPeriodsResponse
	createGeoLocationResponse                  []GeoLocationResponse
	getGeoLocationResponse                     []GeoLocationResponse
	createImageResponse                        []ImageResponse
	getImageResponse                           []ImageResponse
	listLocationDirectionsResponse             []DisplayTextsResponse
	listLocationFacilitiesResponse             []FacilitiesResponse
	listLocationImagesResponse                 []ImagesResponse
	listLocationsResponse                      []LocationsResponse
	createNodeResponse                         []NodeResponse
	getNodeResponse                            []NodeResponse
	createOpeningTimeResponse                  []OpeningTimeResponse
	getOpeningTimeResponse                     []OpeningTimeResponse
	listRegularHoursResponse                   []RegularHoursResponse
	listRelatedLocationsResponse               []GeoLocationsResponse
	listStatusSchedulesResponse                []StatusSchedulesResponse
	createUserResponse                         []UserResponse
	deleteUserResponse                         error
	getUserResponse                            []UserResponse
	listUsersResponse                          []UsersResponse
	updateUserResponse                         []UserResponse
}

func NewMockRepository() MockRepository {
	return &MockRepositoryService{}
}
