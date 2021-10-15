package mocks

import "github.com/satimoto/go-datastore/db"

type CapabilitiesMockData struct {
	Capabilities []db.Capability
	Error        error
}
