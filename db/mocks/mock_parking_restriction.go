package mocks

import "github.com/satimoto/go-datastore/db"

type ParkingRestrictionsPayload struct {
	ParkingRestrictions []db.ParkingRestriction
	Error               error
}
