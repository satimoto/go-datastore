package mocks

import "github.com/satimoto/go-datastore/db"

type FacilitiesResponse struct {
	Facilities []db.Facility
	Error      error
}
