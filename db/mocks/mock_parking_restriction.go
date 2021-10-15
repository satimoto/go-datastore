package mocks

import "github.com/satimoto/go-datastore/db"

type ParkingRestrictionsMockData struct {
	ParkingRestrictions []db.ParkingRestriction
	Error               error
}
