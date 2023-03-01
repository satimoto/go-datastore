package param

import "github.com/satimoto/go-datastore/pkg/db"

func NewUpdatePartyParams(party db.Party) db.UpdatePartyParams {
	return db.UpdatePartyParams{
		ID:                       party.ID,
		IsIntermediateCdrCapable: party.IsIntermediateCdrCapable,
		PublishLocation:          party.PublishLocation,
		PublishNullTariff:        party.PublishNullTariff,
	}
}
