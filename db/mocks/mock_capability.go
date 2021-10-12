package mocks

import "github.com/satimoto/go-datastore/db"

type CapabilitiesPayload struct {
	Capabilities []db.Capability
	Error        error
}
