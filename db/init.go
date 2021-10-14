package db

func NewUpdateConnectorByUidParams(c Connector) UpdateConnectorByUidParams {
	return UpdateConnectorByUidParams{
		EvseID: c.EvseID,
		Uid: c.Uid,
		Standard: c.Standard,
		Format: c.Format,
		PowerType: c.PowerType,
		Voltage: c.Voltage,
		Amperage: c.Amperage,
		TariffID: c.TariffID,
		TermsAndConditions: c.TermsAndConditions,
		LastUpdated: c.LastUpdated,
	}
}