package mocks

import "github.com/satimoto/go-datastore/db"

type FacilitiesMockData struct {
	Facilities []db.Facility
	Error      error
}
