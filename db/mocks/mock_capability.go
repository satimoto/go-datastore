package mocks

import "github.com/satimoto/go-datastore/db"

type CapabilitiesResponse struct {
	Capabilities []db.Capability
	Error        error
}
