package param

import "github.com/satimoto/go-datastore/pkg/db"

func NewUpdateConnectorByUidParams(connector db.Connector) db.UpdateConnectorByUidParams {
	return db.UpdateConnectorByUidParams{
		EvseID:             connector.EvseID,
		Uid:                connector.Uid,
		ConnectorID:        connector.ConnectorID,
		Standard:           connector.Standard,
		Format:             connector.Format,
		PowerType:          connector.PowerType,
		Voltage:            connector.Voltage,
		Amperage:           connector.Amperage,
		Wattage:            connector.Wattage,
		TariffID:           connector.TariffID,
		TermsAndConditions: connector.TermsAndConditions,
		LastUpdated:        connector.LastUpdated,
	}
}
