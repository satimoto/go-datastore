package db

func NewUpdateConnectorByUidParams(c Connector) UpdateConnectorByUidParams {
	return UpdateConnectorByUidParams{
		EvseID:             c.EvseID,
		Uid:                c.Uid,
		Standard:           c.Standard,
		Format:             c.Format,
		PowerType:          c.PowerType,
		Voltage:            c.Voltage,
		Amperage:           c.Amperage,
		TariffID:           c.TariffID,
		TermsAndConditions: c.TermsAndConditions,
		LastUpdated:        c.LastUpdated,
	}
}

func NewUpdateEvseByUidParams(evse Evse) UpdateEvseByUidParams {
	return UpdateEvseByUidParams{
		Uid:               evse.Uid,
		EvseID:            evse.EvseID,
		Status:            evse.Status,
		FloorLevel:        evse.FloorLevel,
		Geom:              evse.Geom,
		GeoLocationID:     evse.GeoLocationID,
		PhysicalReference: evse.PhysicalReference,
		LastUpdated:       evse.LastUpdated,
	}
}

func NewUpdateLocationByUidParams(location Location) UpdateLocationByUidParams {
	return UpdateLocationByUidParams{
		Uid:                location.Uid,
		Type:               location.Type,
		Name:               location.Name,
		Address:            location.Address,
		City:               location.City,
		PostalCode:         location.PostalCode,
		Country:            location.Country,
		Geom:               location.Geom,
		TimeZone:           location.TimeZone,
		ChargingWhenClosed: location.ChargingWhenClosed,
		LastUpdated:        location.LastUpdated,
	}
}
