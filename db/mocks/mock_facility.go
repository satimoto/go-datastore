package mocks

import "github.com/satimoto/go-datastore/db"

type FacilitiesPayload struct {
	Facilities []db.Facility
	Error      error
}
