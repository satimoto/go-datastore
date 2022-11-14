package param

import (
	"github.com/satimoto/go-datastore/pkg/db"
)

func NewUpdateLocationByUidParams(location db.Location) db.UpdateLocationByUidParams {
	return db.UpdateLocationByUidParams{
		Uid:                location.Uid,
		CountryCode:        location.CountryCode,
		PartyID:            location.PartyID,
		Type:               location.Type,
		Name:               location.Name,
		Address:            location.Address,
		City:               location.City,
		PostalCode:         location.PostalCode,
		Country:            location.Country,
		Geom:               location.Geom,
		GeoLocationID:      location.GeoLocationID,
		AvailableEvses:     location.AvailableEvses,
		TotalEvses:         location.TotalEvses,
		IsRemoteCapable:    location.IsRemoteCapable,
		IsRfidCapable:      location.IsRfidCapable,
		OperatorID:         location.OperatorID,
		SuboperatorID:      location.SuboperatorID,
		OwnerID:            location.OwnerID,
		TimeZone:           location.TimeZone,
		OpeningTimeID:      location.OpeningTimeID,
		ChargingWhenClosed: location.ChargingWhenClosed,
		EnergyMixID:        location.EnergyMixID,
		LastUpdated:        location.LastUpdated,
	}
}

func NewUpdateLocationLastUpdatedParams(location db.Location) db.UpdateLocationLastUpdatedParams {
	return db.UpdateLocationLastUpdatedParams{
		ID:              location.ID,
		AvailableEvses:  location.AvailableEvses,
		TotalEvses:      location.TotalEvses,
		IsRemoteCapable: location.IsRemoteCapable,
		IsRfidCapable:   location.IsRfidCapable,
		LastUpdated:     location.LastUpdated,
	}
}
