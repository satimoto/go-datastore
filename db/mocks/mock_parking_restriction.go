package mocks

import "github.com/satimoto/go-datastore/db"

type ParkingRestrictionsResponse struct {
	ParkingRestrictions []db.ParkingRestriction
	Error               error
}
