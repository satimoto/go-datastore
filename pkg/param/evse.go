package param

import (
	"github.com/satimoto/go-datastore/pkg/db"
)

func NewUpdateEvseByUidParams(evse db.Evse) db.UpdateEvseByUidParams {
	return db.UpdateEvseByUidParams{
		Uid:               evse.Uid,
		EvseID:            evse.EvseID,
		Status:            evse.Status,
		FloorLevel:        evse.FloorLevel,
		Geom:              evse.Geom,
		GeoLocationID:     evse.GeoLocationID,
		IsRemoteCapable:   evse.IsRemoteCapable,
		IsRfidCapable:     evse.IsRfidCapable,
		PhysicalReference: evse.PhysicalReference,
		LastUpdated:       evse.LastUpdated,
	}
}

func NewUpdateLocationAvailabilityParams(locationID int64) db.UpdateLocationAvailabilityParams {
	return db.UpdateLocationAvailabilityParams{
		ID:              locationID,
		AvailableEvses:  0,
		TotalEvses:      0,
		IsRemoteCapable: false,
		IsRfidCapable:   false,
	}
}
